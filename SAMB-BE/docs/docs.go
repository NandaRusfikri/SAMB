// Code generated by swaggo/swag. DO NOT EDIT.

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
            "email": "nandarusfikri@gmail.com"
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
        "/v1/barang_masuk/detail": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "CreatePenerimaanBarangDetailController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PenerimaanBarang"
                ],
                "summary": "Create Penerimaan Barang Detail",
                "operationId": "CreatePenerimaanBarangDetailController",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.PenerimaanBarangDetailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/barang_masuk/detail/{TrxInPK}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ListBarangMasukDetailController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PenerimaanBarang"
                ],
                "summary": "ListBarangMasukDetailController",
                "operationId": "ListBarangMasukDetailController",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TrxInPK",
                        "name": "TrxInPK",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/barang_masuk/header": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ListBarangMasukHeaderController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PenerimaanBarang"
                ],
                "summary": "ListBarangMasukHeaderController",
                "operationId": "ListBarangMasukHeaderController",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "CreatePenerimaanBarangHeaderController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PenerimaanBarang"
                ],
                "summary": "Create Penerimaan Barang Header",
                "operationId": "CreatePenerimaanBarangHeaderController",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.PenerimaanBarangHeaderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/barang_masuk/header/{TrxInPK}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetByIdBarangMasukHeaderController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PenerimaanBarang"
                ],
                "summary": "GetByIdBarangMasukHeaderController",
                "operationId": "GetByIdBarangMasukHeaderController",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TrxInPK",
                        "name": "TrxInPK",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/products": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ListProductController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MasterProduct"
                ],
                "summary": "ListProductController",
                "operationId": "ListProductController",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/suppliers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ListSupplierController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MasterSupplier"
                ],
                "summary": "ListSupplierController",
                "operationId": "ListSupplierController",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/warehouses": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ListWareHouseController",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MasterWareHouse"
                ],
                "summary": "ListWareHouseController",
                "operationId": "ListWareHouseController",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.PenerimaanBarangDetailRequest": {
            "type": "object",
            "properties": {
                "TrxInDProductIdf": {
                    "type": "integer"
                },
                "TrxInDQtyDus": {
                    "type": "integer"
                },
                "TrxInDQtyPcs": {
                    "type": "integer"
                },
                "TrxInIDF": {
                    "type": "integer"
                }
            }
        },
        "schemas.PenerimaanBarangHeaderRequest": {
            "type": "object",
            "properties": {
                "TrxInNo": {
                    "type": "string",
                    "example": "TRX-001"
                },
                "TrxInNotes": {
                    "type": "string",
                    "example": "catatan"
                },
                "TrxInSuppIdf": {
                    "type": "integer",
                    "example": 1
                },
                "WhsIdf": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header",
            "x-extension-openapi": "{\"example\": \"value on a json format\"}"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "SAMB-BE",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
