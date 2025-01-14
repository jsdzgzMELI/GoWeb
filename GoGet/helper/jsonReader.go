package helper

import (
	"encoding/json"
	"os"

	"github.com/jsdzgzMELI/GoWebGet/structs"
)

func LoadProducts(filename string) ([]structs.Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// var products []structs.Product
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&structs.Products)
	if err != nil {
		return nil, err
	}

	return structs.Products, nil
}
