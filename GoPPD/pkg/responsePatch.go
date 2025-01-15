package pkg

import "github.com/jsdzgzMELI/GoWeb/GoPPD/internal"

type Response struct {
	Message string            `json:"message"`
	Data    *internal.Product `json:"data"`
}
