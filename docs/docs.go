// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Oyamo Brian",
            "url": "https://github.com/oyamo",
            "email": "oyamo.xyz@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "login user, returns user and set session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login new user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "logout user removing session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "description": "Get current user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get user by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "register new user, returns user and token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/auth/token": {
            "get": {
                "description": "Get CSRF token, required auth session cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get CSRF token",
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "get user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "put": {
                "description": "update existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "some description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Delete user account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/transactions/bywallet/{wallet_id}": {
            "get": {
                "description": "Get transactions by wallet id. Transactions are logs of money actions e.g debit and credit.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get transactions by wallet id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Wallet"
                            }
                        }
                    }
                }
            }
        },
        "/wallets": {
            "post": {
                "description": "Set up a new wallet for the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Create a new wallet",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Wallet"
                        }
                    }
                }
            }
        },
        "/wallets/{wallet_id}": {
            "get": {
                "description": "Lookup a wallet by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Get a wallet by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Wallet"
                        }
                    }
                }
            },
            "put": {
                "description": "Change the value of a wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Update a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Wallet"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a wallet from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Delete a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Wallet"
                        }
                    }
                }
            }
        },
        "/wallets/{wallet_id}/balance": {
            "get": {
                "description": "Get a wallet balance by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Get a wallet balance by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Wallet"
                        }
                    }
                }
            }
        },
        "/wallets/{wallet_id}/credit": {
            "post": {
                "description": "credit money into a wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Put money into a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Wallet"
                        }
                    }
                }
            }
        },
        "/wallets/{wallet_id}/debit": {
            "post": {
                "description": "debit money from a wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "withdraw money from a wallet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "wallet_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Wallet"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password"
            ],
            "properties": {
                "about": {
                    "type": "string",
                    "maxLength": 1024
                },
                "address": {
                    "type": "string",
                    "maxLength": 250
                },
                "avatar": {
                    "type": "string",
                    "maxLength": 512
                },
                "birthday": {
                    "type": "string",
                    "maxLength": 10
                },
                "city": {
                    "type": "string",
                    "maxLength": 24
                },
                "country": {
                    "type": "string",
                    "maxLength": 24
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 60
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 30
                },
                "gender": {
                    "type": "string",
                    "maxLength": 10
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 30
                },
                "login_date": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "phone_number": {
                    "type": "string",
                    "maxLength": 20
                },
                "postcode": {
                    "type": "integer"
                },
                "role": {
                    "type": "string",
                    "maxLength": 10
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.Wallet": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "balance": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "wallet_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Wallet API",
	Description:      "An API for managing transactions.go of players",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}