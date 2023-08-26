package schemas

type SchemaDatabaseError struct {
	Message string
	Code    int
	Error   error
}

type SchemaErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

type SchemaUnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}
