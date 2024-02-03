// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://swagger.io/",
        "contact": {
            "name": "Trekkstay Team",
            "url": "https://www.trekkstay.com",
            "email": "support@trekkstay.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hotel-emp/create-emp": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create new hotel employee account, require hotel owner permission and hotel profile already created",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hotel Employee"
                ],
                "summary": "Create new hotel employee account",
                "parameters": [
                    {
                        "description": "CreateHotelEmpReq JSON",
                        "name": "CreateHotelEmpReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CreateHotelEmpReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/hotel-emp/create-owner": {
            "post": {
                "description": "Create new hotel owner account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hotel Employee"
                ],
                "summary": "Create new hotel owner account",
                "parameters": [
                    {
                        "description": "CreateHotelOwnerReq JSON",
                        "name": "CreateHotelOwnerReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CreateHotelOwnerReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/region/list-district": {
            "get": {
                "description": "List all districts of a province",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Region"
                ],
                "summary": "List districts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Province code",
                        "name": "province_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/region/list-province": {
            "get": {
                "description": "List all provinces",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Region"
                ],
                "summary": "List provinces",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/region/list-ward": {
            "get": {
                "description": "List all wards of a district",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Region"
                ],
                "summary": "List wards",
                "parameters": [
                    {
                        "type": "string",
                        "description": "District code",
                        "name": "district_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/change-password": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Change password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Change password",
                "parameters": [
                    {
                        "description": "ChangePasswordReq JSON",
                        "name": "ChangePasswordReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ChangePasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login user by email and password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "LoginUserReq JSON",
                        "name": "LoginUserReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.LoginUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/refresh-token": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get new access token and refresh token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Refresh token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/reset-password": {
            "post": {
                "description": "Reset password and send new password to email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "ResetPasswordReq JSON",
                        "name": "ResetPasswordReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ResetPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Register new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "CreateUserReq JSON",
                        "name": "CreateUserReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "patch": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Update user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "UpdateUserReq JSON",
                        "name": "UpdateUserReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UpdateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "req.ChangePasswordReq": {
            "type": "object",
            "required": [
                "new_pwd",
                "old_pwd"
            ],
            "properties": {
                "old_pwd": {
                    "type": "string",
                    "x-order": "1"
                },
                "new_pwd": {
                    "type": "string",
                    "x-order": "2"
                }
            }
        },
        "req.CreateHotelEmpReq": {
            "type": "object",
            "required": [
                "base_salary",
                "contract",
                "email",
                "full_name",
                "phone"
            ],
            "properties": {
                "full_name": {
                    "type": "string",
                    "x-order": "1"
                },
                "email": {
                    "type": "string",
                    "x-order": "2"
                },
                "phone": {
                    "type": "string",
                    "x-order": "3"
                },
                "contract": {
                    "type": "string",
                    "enum": [
                        "FULL_TIME",
                        "PART_TIME",
                        "INTERNSHIP"
                    ],
                    "x-order": "4"
                },
                "base_salary": {
                    "type": "integer",
                    "x-order": "5"
                }
            }
        },
        "req.CreateHotelOwnerReq": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "password",
                "phone"
            ],
            "properties": {
                "full_name": {
                    "type": "string",
                    "x-order": "1"
                },
                "email": {
                    "type": "string",
                    "x-order": "2"
                },
                "phone": {
                    "type": "string",
                    "x-order": "3"
                },
                "password": {
                    "type": "string",
                    "x-order": "4"
                }
            }
        },
        "req.CreateUserReq": {
            "type": "object",
            "required": [
                "email",
                "full_name",
                "password"
            ],
            "properties": {
                "full_name": {
                    "type": "string",
                    "x-order": "1"
                },
                "email": {
                    "type": "string",
                    "x-order": "2"
                },
                "phone": {
                    "type": "string",
                    "x-order": "3"
                },
                "password": {
                    "type": "string",
                    "x-order": "4"
                }
            }
        },
        "req.LoginUserReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "x-order": "1"
                },
                "password": {
                    "type": "string",
                    "x-order": "2"
                }
            }
        },
        "req.ResetPasswordReq": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "req.UpdateUserReq": {
            "type": "object",
            "properties": {
                "full_name": {
                    "type": "string",
                    "x-order": "1"
                },
                "email": {
                    "type": "string",
                    "x-order": "2"
                },
                "phone": {
                    "type": "string",
                    "x-order": "3"
                }
            }
        },
        "res.ErrorResponse": {
            "type": "object",
            "properties": {
                "error_key": {
                    "type": "string"
                },
                "log": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "res.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.1",
	Host:             "52.221.204.232:8888",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Trekkstay - Hotel Booking System API",
	Description:      "API system for Trekkstay - Hotel Booking System",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
