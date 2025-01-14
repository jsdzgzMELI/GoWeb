package structs

type ResponsePrice struct {
	Message string     `json:"message"`
	Data    *[]Product `json:"data"`
	Total   int        `json:"total"`
}
