// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "name": "1362",
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
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserLogin 登录相关"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginInputDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/name": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Index 依赖注入相关"
                ],
                "summary": "获取最新age",
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.UserLoginInputDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "passWord": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8181",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ExamCodeAboutTeslaByGin",
	Description:      "ExamCodeAboutTeslaByGin",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
