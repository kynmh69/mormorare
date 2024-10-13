package domain

type ErrorJson struct {
	Error string `json:"error"`
}

func NewErrorJson(err string) *ErrorJson {
	return &ErrorJson{Error: err}
}
