// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/orders/": {
            "get": {
                "description": "Get orders list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "List orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Start of the period used to filter orders, in datetime format (e.g. 2021-10-23T21:00:00.000Z)",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End of the period used to filter orders, in datetime format (e.g. 2021-10-24T20:59:59.999Z)",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "date_ordered",
                            "date_closed"
                        ],
                        "type": "string",
                        "description": "Date used to filter orders",
                        "name": "dateColumn",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Query used to filter orders, might be customer name, order number or buyer name",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "itemsPerPage",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "Order status (Создан 1, В обработке 2, На согласовании 3, На сборке 10, В пути 21, Доставлен 15, Приёмка 20, Принят 22, Завершён 24, Отказ/Не согласован 4)",
                        "name": "selectedStatuses[]",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Orders"
                        }
                    }
                }
            }
        },
        "/orders/{orderId}": {
            "get": {
                "description": "Get order itemslist",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "List order lines",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order Id",
                        "name": "orderId",
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
                                "$ref": "#/definitions/main.OrderLine"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.OrderDetails": {
            "type": "object",
            "properties": {
                "accepted_date": {
                    "type": "string"
                },
                "buyer": {
                    "type": "string"
                },
                "buyer_id": {
                    "type": "integer"
                },
                "closed_date": {
                    "type": "string"
                },
                "consignee_name": {
                    "type": "string"
                },
                "contractor_number": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "customer_name": {
                    "type": "string"
                },
                "delivered_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "on_order_coupon": {
                    "type": "number"
                },
                "on_order_coupon_fixed": {
                    "type": "number"
                },
                "ordered_date": {
                    "type": "string"
                },
                "seller_id": {
                    "type": "integer"
                },
                "seller_name": {
                    "type": "string"
                },
                "shipped_date": {
                    "type": "string"
                },
                "shipping_date_est": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "sum": {
                    "type": "number"
                },
                "sum_with_tax": {
                    "type": "number"
                }
            }
        },
        "main.OrderLine": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "comment": {
                    "type": "string"
                },
                "count": {
                    "type": "number"
                },
                "coupon_fixed": {
                    "type": "number"
                },
                "coupon_percent": {
                    "type": "number"
                },
                "coupon_value": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "nds": {
                    "type": "number"
                },
                "price": {
                    "type": "number"
                },
                "product_id": {
                    "type": "integer"
                },
                "sum": {
                    "type": "number"
                },
                "sum_with_tax": {
                    "type": "number"
                },
                "tax": {
                    "type": "number"
                },
                "warehouse": {
                    "type": "string"
                }
            }
        },
        "main.Orders": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.OrderDetails"
                    }
                },
                "sum": {
                    "type": "number"
                },
                "sum_with_tax": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
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
	Version:     "1.0",
	Host:        "industrial.market",
	BasePath:    "/crutch/methods",
	Schemes:     []string{},
	Title:       "Industrial.Market API",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
