package structs

type ResponseGet struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
}
