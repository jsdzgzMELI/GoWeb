package structs

type ResponsePost struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
}
