package schemas

type SchemaResponses struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Count   int64       `json:"count,omitempty"`
	Data    interface{} `json:"data"`
	//Items   interface{} `json:"items,omitempty"`
}
