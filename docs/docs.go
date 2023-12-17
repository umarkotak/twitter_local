// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request_contract.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request_contract.UserRegister": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "photo_url",
                "username"
            ],
            "properties": {
                "about": {
                    "type": "string",
                    "example": "hello world"
                },
                "email": {
                    "type": "string",
                    "example": "jhone.doe@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "jhone doe"
                },
                "password": {
                    "type": "string",
                    "example": "rahasia"
                },
                "photo_url": {
                    "type": "string",
                    "example": "https://picsum.photos/300/300"
                },
                "username": {
                    "type": "string",
                    "example": "jhone_doe"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:17000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "Twitter local API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
