package pkg

import "github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal"

type ResponseGet struct {
	Message string              `json:"message"`
	Data    *[]internal.Product `json:"data"`
}
