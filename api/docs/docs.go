// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Sean Cheng",
            "url": "https://blog.seancheng.space",
            "email": "blackhorseya@gmail.com"
        },
        "license": {
            "name": "GPL-3.0",
            "url": "https://spdx.org/licenses/GPL-3.0-only.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/liveness": {
            "get": {
                "description": "to know when to restart an application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/readiness": {
            "get": {
                "description": "Show application was ready to start accepting traffic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "user profile",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.Profile"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/auth/signup": {
            "post": {
                "description": "Signup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Signup",
                "parameters": [
                    {
                        "description": "new user profile",
                        "name": "newUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.Profile"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/goals": {
            "get": {
                "description": "List all objectives",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "size of page",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a objective",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "parameters": [
                    {
                        "description": "created goal",
                        "name": "created",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/okr.Goal"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/goals/{id}": {
            "get": {
                "description": "Get a objective by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of objective",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Get a objective by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of objective",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/goals/{id}/results": {
            "get": {
                "description": "Get key result by goal id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Results"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of goal",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/goals/{id}/title": {
            "patch": {
                "description": "Modify title of goal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of goal",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated goal",
                        "name": "updated",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/okr.Goal"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/results": {
            "get": {
                "description": "List all key results",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Results"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "size of page",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a key result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Results"
                ],
                "parameters": [
                    {
                        "description": "created key result",
                        "name": "created",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/okr.Result"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/results/{id}": {
            "get": {
                "description": "Get a key result by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Results"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of key result",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a key result by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Results"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of key result",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/results/{id}/title": {
            "patch": {
                "description": "Modify title of result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Results"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of result",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated result",
                        "name": "updated",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/okr.Result"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/tasks": {
            "get": {
                "description": "List all tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "size of page",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "description": "created task",
                        "name": "created",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/tasks/{id}": {
            "get": {
                "description": "Get a task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of task",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of task",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/tasks/{id}/status": {
            "patch": {
                "description": "UpdateStatus a status of task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of task",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated task",
                        "name": "updated",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        },
        "/v1/tasks/{id}/title": {
            "patch": {
                "description": "ModifyTitle a status of task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of task",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated task",
                        "name": "updated",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/er.APPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "er.APPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "okr.Goal": {
            "type": "object",
            "properties": {
                "create_at": {
                    "description": "CreateAt describe the objective create milliseconds",
                    "type": "integer"
                },
                "end_at": {
                    "description": "EndAt describe the objective end timex milliseconds",
                    "type": "integer"
                },
                "id": {
                    "description": "ID describe the unique identify code of objective",
                    "type": "string"
                },
                "key_results": {
                    "description": "KeyResults describe key results of objective",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/okr.Result"
                    }
                },
                "start_at": {
                    "description": "StartAt describe the objective start timex milliseconds",
                    "type": "integer"
                },
                "title": {
                    "description": "Title describe the title of objective",
                    "type": "string"
                }
            }
        },
        "okr.Result": {
            "type": "object",
            "properties": {
                "actual": {
                    "description": "Actual describe the actual of key result",
                    "type": "integer"
                },
                "create_at": {
                    "description": "CreateAt describe the key result create milliseconds",
                    "type": "integer"
                },
                "goal_id": {
                    "description": "GoalID describe the parent goal's id",
                    "type": "string"
                },
                "id": {
                    "description": "ID describe the unique identify code of key result",
                    "type": "string"
                },
                "target": {
                    "description": "Target describe the target of key result",
                    "type": "integer"
                },
                "title": {
                    "description": "Title describe the title of key result",
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "task.Task": {
            "type": "object",
            "properties": {
                "completed": {
                    "description": "Completed describe the completed of task",
                    "type": "boolean"
                },
                "create_at": {
                    "description": "CreateAt describe the task create milliseconds",
                    "type": "integer"
                },
                "id": {
                    "description": "ID describe the unique identify code of task",
                    "type": "string"
                },
                "result_id": {
                    "description": "ResultID describe the parent key result's id",
                    "type": "string"
                },
                "status": {
                    "description": "Status describe the status of task",
                    "type": "integer"
                },
                "title": {
                    "description": "Title describe the title of task",
                    "type": "string"
                }
            }
        },
        "user.Profile": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "AccessToken describe this user's accessToken",
                    "type": "string"
                },
                "email": {
                    "description": "Email describe user's email to login system",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Lobster API",
	Description: "Lobster API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
