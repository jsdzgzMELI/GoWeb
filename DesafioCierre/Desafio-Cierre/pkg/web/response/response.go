package response

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
