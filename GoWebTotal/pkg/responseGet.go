package pkg

import "github.com/jsdzgzMELI/GoWeb/GoWebTotal/internal/domain"

type ResponseGet struct {
	Message string            `json:"message"`
	Data    *[]domain.Product `json:"data"`
}
