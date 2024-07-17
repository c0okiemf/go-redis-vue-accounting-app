package client

import (
	"accounting/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	goCache "github.com/patrickmn/go-cache"
)

var cache = goCache.New(5*time.Minute, 10*time.Minute)

type ImmuDBClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewImmuDBClient() *ImmuDBClient {
	return &ImmuDBClient{
		baseURL: os.Getenv("IMMUDB_BASE_URL"),
		apiKey:  os.Getenv("IMMUDB_API_KEY"),
		client:  &http.Client{},
	}
}

func (c *ImmuDBClient) createRequest(method, path string, body []byte) (*http.Request, error) {
	url := c.baseURL + path

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *ImmuDBClient) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
		} else {
			log.Printf("non-ok status code body: %v", string(bytes))
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp, nil
}

// sends a PUT request to record a new transaction
func (c *ImmuDBClient) AddTransaction(transaction types.TransactionInput) error {
	payloadBytes, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	req, err := c.createRequest("PUT", "/document", payloadBytes)
	if err != nil {
		return err
	}

	defer req.Body.Close()

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	cache.Flush()

	return nil
}

type GetTransactionsOrderBy struct {
	Field string `json:"field"`
	Desc   bool   `json:"desc"`
}

type GetTransactionsQuery struct {
	OrderBy []GetTransactionsOrderBy `json:"orderBy"`
}

type GetTransactionsPayload struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Query  GetTransactionsQuery `json:"query"`
}

// sneds a POST request to get a list of transactions
func (c *ImmuDBClient) GetTransactions(page, perPage int) ([]types.Transaction, error) {
	cacheKey := fmt.Sprintf("page:%d_perPage:%d", page, perPage)
	if cachedData, found := cache.Get(cacheKey); cachedData != nil {
		if found {
			log.Println("Getting transactions from cache")
			return cachedData.([]types.Transaction), nil
		}
	}

	payload := GetTransactionsPayload{
		Page:    page,
		PerPage: perPage,
		Query: GetTransactionsQuery{
			OrderBy: []GetTransactionsOrderBy{
				{
					Field: "_vault_md.ts",
					Desc:  true,
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %v", err)
	}

	req, err := c.createRequest("POST", "/documents/search", payloadBytes)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Revisions []struct {
			Document struct {
				ID              string                 `json:"_id"`
				AccountName     string                 `json:"accountName"`
				AccountNumber   string                 `json:"accountNumber"`
				IBAN            string                 `json:"iban"`
				Address         string                 `json:"address"`
				Amount          int                    `json:"amount"`
				TransactionType types.TransactionType `json:"transactionType"`
			} `json:"document"`
		} `json:"revisions"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	var transactions []types.Transaction
	for _, rev := range response.Revisions {
		transaction := types.Transaction{
			ID: rev.Document.ID,
			TransactionInput: types.TransactionInput{
				AccountNumber:   rev.Document.AccountNumber,
				AccountName:     rev.Document.AccountName,
				IBAN:            rev.Document.IBAN,
				Address:         rev.Document.Address,
				Amount:          rev.Document.Amount,
				TransactionType: rev.Document.TransactionType,
			},
		}
		transactions = append(transactions, transaction)
	}

	cache.Set(cacheKey, transactions, goCache.DefaultExpiration)

	return transactions, nil
}

type GetTransactionCountQuery struct {
	Expressions []interface{} `json:"expressions"`
	OrderBy     []interface{} `json:"orderBy"`
	Limit       int           `json:"limit"`
}

type GetTransactionCountPayload struct {
	Query GetTransactionCountQuery `json:"query"`
}

// sends a POST request to get the total number of transactions
func (c *ImmuDBClient) GetTransactionCount() int {
	const cacheKey = "transactionCount"
	if cachedData, found := cache.Get(cacheKey); cachedData != nil {
		if found {
			log.Println("Getting transaction count from cache")
			return cachedData.(int)
		}
	}

	payload := GetTransactionCountPayload{
		Query: GetTransactionCountQuery{
			Expressions: []interface{}{},
			OrderBy:     []interface{}{},
			Limit:       0,
		},
	}

	payload.Query.Expressions = []any{}
	payload.Query.OrderBy = []any{}
	payload.Query.Limit = 0

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marshalling JSON:", err)
		return 0
	}

	req, err := c.createRequest("POST", "/documents/count", payloadBytes)
	if err != nil {
		log.Println("error creating request:", err)
		return 0
	}
	defer req.Body.Close()

	resp, err := c.doRequest(req)
	if err != nil {
		log.Println("error sending request:", err)
		return 0
	}
	defer resp.Body.Close()

	var response struct {
		Count int `json:"count"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0
	}

	cache.Set(cacheKey, response.Count, goCache.DefaultExpiration)

	return response.Count
}
