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
        "/api/v1/people": {
            "get": {
                "description": "Get filtered and paginated list of people",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Get people by query parameters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by age",
                        "name": "age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by gender",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by nationality",
                        "name": "nationality",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort by (name, surname, age, gender, nationality)",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "Sort in descending order",
                        "name": "sort_desc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPeopleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/person": {
            "post": {
                "description": "Create new person with provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Create new person",
                "parameters": [
                    {
                        "description": "Person data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePersonRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePersonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update person with provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Update person by ID",
                "parameters": [
                    {
                        "description": "Person data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePersonRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePersonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/person/{id}": {
            "get": {
                "description": "Get person details by person ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Get person by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPersonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete person from database by person ID",
                "tags": [
                    "persons"
                ],
                "summary": "Delete person by ID",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreatePersonRequest": {
            "type": "object",
            "required": [
                "name",
                "surname"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePersonResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "string"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "dto.GetPeopleResponse": {
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/dto.Meta"
                },
                "people": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Person"
                    }
                }
            }
        },
        "dto.GetPersonResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nation": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "dto.Meta": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "dto.Person": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nation": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "dto.UpdatePersonRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "minimum": 0
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ]
                },
                "id": {
                    "type": "integer",
                    "minimum": 1
                },
                "name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "dto.UpdatePersonResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
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
	BasePath:         "",
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
