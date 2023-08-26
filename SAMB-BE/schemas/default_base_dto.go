package schemas

import (
	"errors"
	"mime/multipart"
)

type FilterBaseDto struct {
	FilterID int `json:"filter_id" form:"filter_id"`
	//IsLookup    bool   `json:"is_lookup" form:"is_lookup"`
	SearchText  string `json:"search_text" form:"search_text" example:"Search example name,email or code"`
	OrderField  string `json:"order_field" form:"order_field" example:"id|desc"`
	FilterPage  int    `json:"filter_page" form:"filter_page" example:"1"`
	FilterLimit int    `json:"filter_limit" form:"filter_limit" example:"10"`
}

type LookupBaseDto struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ResponseBaseDto struct {
	Code         int    `json:"code"`
	Message      string `json:"message,omitempty"`
	MessageError string `json:"message_error,omitempty"`
}

type CallAPIDto struct {
	Method       string
	Url          string
	ContentType  string
	Headers      map[string]interface{}
	BodyRequest  string
	BodyResponse string
	HttpCode     int
}

func (d *CallAPIDto) Validate() error {
	if d.Method == "" {
		return errors.New("method required")
	}
	if d.Url == "" {
		return errors.New("url required")
	}

	return nil
}

type ExportBase64ResponseDto struct {
	ResponseBaseDto
	Extension string `json:"extension"`
	Base64    string `json:"base_64"`
}
type UploadFileRequestDto struct {
	File []*multipart.FileHeader `json:"-" form:"file" binding:"required"`
	Path string                  `json:"path"`
}
type UploadFileResponseDto struct {
	ResponseBaseDto
	Count int64                   `json:"count,omitempty"`
	Items []*UploadFileRequestDto `json:"items,omitempty"`
}
