// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-01-25 17:16:41.7490111 +0800 CST m=+0.087732801

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "An example of gin",
        "title": "Golang Gin API",
        "termsOfService": "https://github.com/hequan2017/go-admin",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/hequan2017/go-admin/master/LICENSE"
        },
        "version": "1.0"
    },
    "paths": {
        "/auth": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取登录token 信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": { \"token\": \"xxx\" }, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
