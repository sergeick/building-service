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
        "/buildings": {
            "get": {
                "description": "Get list of buildings",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "List buildings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Year built",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Floors count",
                        "name": "floors",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Building"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new building",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Create a new building",
                "parameters": [
                    {
                        "description": "Building object",
                        "name": "building",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/docs": {
            "get": {
                "description": "Get Swagger API documentation",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "docs"
                ],
                "summary": "Swagger API documentation",
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Building": {
            "type": "object",
            "required": [
                "city",
                "floors",
                "name",
                "year"
            ],
            "properties": {
                "city": {
                    "type": "string"
                },
                "floors": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Building Service API",
	Description:      "Service for managing buildings",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
