// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transaction": {
            "post": {
                "description": "Add a new transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Add transaction",
                "parameters": [
                    {
                        "description": "Transaction input",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.TransactionInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.InputValidationErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "get": {
                "description": "Get all transactions with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Required page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Required per page",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.TransactionResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "main.InputValidationErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.ValidationError"
                    }
                }
            }
        },
        "main.TransactionResponse": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Transaction"
                    }
                }
            }
        },
        "types.Transaction": {
            "type": "object",
            "required": [
                "accountNumber",
                "amount",
                "iban"
            ],
            "properties": {
                "accountName": {
                    "type": "string"
                },
                "accountNumber": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer",
                    "minimum": 1
                },
                "iban": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "transactionType": {
                    "$ref": "#/definitions/types.TransactionType"
                }
            }
        },
        "types.TransactionInput": {
            "type": "object",
            "required": [
                "accountNumber",
                "amount",
                "iban"
            ],
            "properties": {
                "accountName": {
                    "type": "string"
                },
                "accountNumber": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer",
                    "minimum": 1
                },
                "iban": {
                    "type": "string"
                },
                "transactionType": {
                    "$ref": "#/definitions/types.TransactionType"
                }
            }
        },
        "types.TransactionType": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "TransactionTypeSent",
                "TransactionTypeReceived"
            ]
        },
        "types.ValidationError": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}