package pkg

import "github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal"

type Response struct {
	Message string            `json:"message"`
	Data    *internal.Product `json:"data"`
}
