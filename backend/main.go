package main

import (
	"accounting/client"
	_ "accounting/docs"
	"accounting/helpers"
	"accounting/types"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type TransactionResponse struct {
	Transactions []types.Transaction `json:"transactions"`
	Total        int                 `json:"total"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type InputValidationErrorResponse struct {
	Errors []types.ValidationError `json:"errors"`
}

// @BasePath /api/v1

// @Summary Get transactions
// @Description Get all transactions with pagination
// @Tags transactions
// @Accept json
// @Produce json
// @Param page query int false "Required page"
// @Param per_page query int false "Required per page"
// @Success 200 {object} TransactionResponse
// @Failure 404 {object} ErrorResponse
// @Router /transactions [get]
func getTransactions(clientInstance *client.ImmuDBClient) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get total number of transactions to enable pagination
		transactionCount := clientInstance.GetTransactionCount()

		if transactionCount == 0 {
			c.JSON(http.StatusOK, gin.H{"transactions": []types.Transaction{}, "total": 0})
			return
		}

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			error := ErrorResponse{
				Errors: []string{"Invalid page number"},
			}
			c.JSON(http.StatusBadRequest, error)
			return
		}
		perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "100"))
		if err != nil {
			error := ErrorResponse{
				Errors: []string{"Invalid per page number"},
			}
			c.JSON(http.StatusBadRequest, error)
			return
		}

		res, err := clientInstance.GetTransactions(page, perPage)
		if err != nil {
			error := ErrorResponse{
				Errors: []string{err.Error()},
			}
			c.JSON(http.StatusInternalServerError, error)
			return
		}
		transactionResponse := TransactionResponse{
			Transactions: res,
			Total:        transactionCount,
		}
		c.JSON(http.StatusOK, transactionResponse)
	}
}

// @Summary Add transaction
// @Description Add a new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body types.TransactionInput true "Transaction input"
// @Success 200
// @Failure 400 {object} InputValidationErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [post]
func addTransaction(clientInstance *client.ImmuDBClient) func(*gin.Context) {
	return func(c *gin.Context) {
		var account types.TransactionInput
		if err := c.ShouldBindJSON(&account); err != nil {
			userFacingErrors := helpers.MakeAddTransactionUserFacingErrors(err, account)
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": userFacingErrors,
			})
			return
		}
		if err := clientInstance.AddTransaction(account); err != nil {
			error := ErrorResponse{
				Errors: []string{err.Error()},
			}
			c.JSON(http.StatusInternalServerError, error)
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}

func main() {
	helpers.LoadEnvFile()

	clientInstance := client.NewImmuDBClient()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	v1.GET("/transactions", getTransactions(clientInstance))
	v1.POST("/transaction", addTransaction(clientInstance))

	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
