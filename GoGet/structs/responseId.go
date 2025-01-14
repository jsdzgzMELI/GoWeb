package structs

type ResponseId struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
}
