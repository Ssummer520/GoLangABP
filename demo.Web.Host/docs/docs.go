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
        "/User": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User 用户相关"
                ],
                "summary": "新增用户信息",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.UserAddInputDto"
                        }
                    }
                ],
                "responses": {}
            }
        },
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
                "responses": {
                    "200": {
                        "description": "{\"Success\": true,\"Token\":\"5545\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\": \"who are you\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserAddInputDto": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "年龄",
                    "type": "integer"
                },
                "name": {
                    "description": "姓名",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "电话",
                    "type": "string"
                },
                "sex": {
                    "description": "性别",
                    "type": "integer"
                }
            }
        },
        "dto.UserLoginInputDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
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
