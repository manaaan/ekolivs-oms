package zettle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct{}

const getProductsPath = "/organizations/{organizationUuid}/products/v2"

func New() (*Service, error) {
	return &Service{}, nil
}

func (Service) GetProducts() (ProductResponse, error) {
	response, err := http.Get(getProductsPath)
	if err != nil {
		fmt.Println(err)
		return ProductResponse{}, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ProductResponse{}, err
	}

	var products ProductResponse
	json.Unmarshal(responseData, &products)

	return products, nil
}
