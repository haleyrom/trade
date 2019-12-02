// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-02 17:27:11.361097977 +0800 CST m=+0.041174625

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/team/create": {
            "post": {
                "description": "创建团队",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "1. CreateTeam"
                ],
                "summary": "创建团队接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "团队名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/team/dismiss": {
            "post": {
                "description": "查看团队列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "5. ReadListTeam"
                ],
                "summary": "查看团队列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页数默认为1",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页数数量默认为20",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/team/exit": {
            "post": {
                "description": "退出团队",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "3. ExitTeam"
                ],
                "summary": "退出团队接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "团队id",
                        "name": "tid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/team/info": {
            "post": {
                "description": "查看团队信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "6. ReadInfoTeam"
                ],
                "summary": "查看团队信息接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "团队id",
                        "name": "tid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/team/join": {
            "post": {
                "description": "加入团队",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "2. JoinTeam"
                ],
                "summary": "加入团队接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "团队id",
                        "name": "tid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/team/user_list": {
            "post": {
                "description": "查看团队成员列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "7. ReadUserListTeam"
                ],
                "summary": "查看团队成员列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "页数默认为1",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页数数量默认为20",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
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
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "测试",
	Description: "测试",
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