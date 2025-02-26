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
        "/api/user/register": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Register Request",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.UserRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/public/banner/{subdomain}/{refferal}": {
            "get": {
                "description": "show public of Banner",
                "tags": [
                    "Data Banner"
                ],
                "summary": "Get Public Banner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "subdomain",
                        "name": "subdomain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "referral",
                        "name": "referral",
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/public/gallery/{subdomain}/{referral}": {
            "get": {
                "description": "show detail of Gallery",
                "tags": [
                    "Data Gallery"
                ],
                "summary": "Get Detail Gallery",
                "parameters": [
                    {
                        "type": "string",
                        "description": "subdomain",
                        "name": "subdomain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "referral",
                        "name": "referral",
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/public/social-media/{subdomain}/{referral}": {
            "get": {
                "description": "show list of Public Social Media",
                "tags": [
                    "Data Social Media"
                ],
                "summary": "Get List Public Social Media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "subdomain",
                        "name": "subdomain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "referral",
                        "name": "referral",
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/public/temlate/{subdomain}/{referral}": {
            "get": {
                "description": "show list of Public Template",
                "tags": [
                    "Data Template"
                ],
                "summary": "Get List Public Template",
                "parameters": [
                    {
                        "type": "string",
                        "description": "subdomain",
                        "name": "subdomain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "referral",
                        "name": "referral",
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sales/banner": {
            "get": {
                "description": "show list of Banner",
                "tags": [
                    "Data Banner"
                ],
                "summary": "Get List Banner",
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
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "add data of Banner",
                "tags": [
                    "Data Banner"
                ],
                "summary": "Add Data Banner",
                "parameters": [
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.BannerReq"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sales/banner/{id}": {
            "get": {
                "description": "show detail of Banner",
                "tags": [
                    "Data Banner"
                ],
                "summary": "Get Detail Banner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
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
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Data of Banner",
                "tags": [
                    "Data Banner"
                ],
                "summary": "Update Data Banner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.BannerUpdateReq"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sales/gallery": {
            "get": {
                "description": "show List of Gallery",
                "tags": [
                    "Data Gallery"
                ],
                "summary": "Get List Gallery",
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
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "add data of Template",
                "tags": [
                    "Data Template"
                ],
                "summary": "Add Data Template",
                "parameters": [
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddTemplateReq"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sales/gallery/{id}": {
            "get": {
                "description": "show detail of Gallery",
                "tags": [
                    "Data Gallery"
                ],
                "summary": "Get Detail Gallery",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
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
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Data of Gallery",
                "tags": [
                    "Data Gallery"
                ],
                "summary": "Update Data Gallery",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateGalleryParam"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sales/social-media": {
            "get": {
                "description": "show list of Social Media",
                "tags": [
                    "Data Social Media"
                ],
                "summary": "Get List Social Media",
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
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "add data of Social Media",
                "tags": [
                    "Data Social Media"
                ],
                "summary": "Add Data Social Media",
                "parameters": [
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddSocialMediaReq"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sales/temlate": {
            "get": {
                "description": "show list of Template",
                "tags": [
                    "Data Template"
                ],
                "summary": "Get List Template",
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
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sales/template/{id}": {
            "get": {
                "description": "show detail of Template",
                "tags": [
                    "Data Template"
                ],
                "summary": "Get Detail Template",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
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
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Data of Template",
                "tags": [
                    "Data Template"
                ],
                "summary": "Update Data Template",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateTemplateReq"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/testimony": {
            "get": {
                "description": "show detail of Customer",
                "tags": [
                    "Data Testimoni"
                ],
                "summary": "Get Detail Customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Data of Testimoni",
                "tags": [
                    "Data Testimoni"
                ],
                "summary": "Update Data Testimoni",
                "parameters": [
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Testimoni"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "create data of Testimoni",
                "tags": [
                    "Data Testimoni"
                ],
                "summary": "Create Data Testimoni",
                "parameters": [
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Testimoni"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/testimony/list": {
            "get": {
                "description": "show list of Customer",
                "tags": [
                    "Data Testimoni"
                ],
                "summary": "Get List Customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.RegisterRequest": {
            "type": "object",
            "required": [
                "role"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "description": "fullname",
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "sales",
                        "agent"
                    ]
                },
                "username": {
                    "type": "string"
                },
                "whatsapp_number": {
                    "type": "string"
                }
            }
        },
        "model.TokenRes": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "model.UserRes": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "token_data": {
                    "description": "token response",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.TokenRes"
                        }
                    ]
                },
                "username": {
                    "type": "string"
                },
                "whatsapp_number": {
                    "type": "string"
                }
            }
        },
        "request.AddGalleryParam": {
            "type": "object",
            "properties": {
                "image_url": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "request.AddSocialMediaReq": {
            "type": "object",
            "properties": {
                "data_social_media": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.SocialMediaListReq"
                    }
                }
            }
        },
        "request.AddTemplateReq": {
            "type": "object",
            "properties": {
                "color_plate": {
                    "type": "string"
                }
            }
        },
        "request.BannerReq": {
            "type": "object",
            "properties": {
                "data_banner": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.DataBanner"
                    }
                }
            }
        },
        "request.BannerUpdateReq": {
            "type": "object",
            "required": [
                "image_url"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                }
            }
        },
        "request.DataBanner": {
            "type": "object",
            "required": [
                "image_url"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                }
            }
        },
        "request.SocialMediaListReq": {
            "type": "object",
            "properties": {
                "link_embed": {
                    "type": "string"
                },
                "social_media_name": {
                    "type": "string"
                },
                "user_account": {
                    "type": "string"
                }
            }
        },
        "request.Testimoni": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deskripsi": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "integer"
                },
                "modified_at": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "request.UpdateGalleryParam": {
            "type": "object",
            "required": [
                "image_url"
            ],
            "properties": {
                "image_url": {
                    "type": "string"
                }
            }
        },
        "request.UpdateTemplateReq": {
            "type": "object",
            "required": [
                "color_plate"
            ],
            "properties": {
                "color_plate": {
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
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "url": {
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
