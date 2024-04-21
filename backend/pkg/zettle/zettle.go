package zettle

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Service struct{}

const getProductsPath = "/organizations/{organizationUuid}/products/v2"

func New() (*Service, error) {
	return &Service{}, nil
}

func (Service) GetProducts() (*ProductResponse, error) {
	response, err := http.Get(getProductsPath)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	var products *ProductResponse
	if err := json.Unmarshal(responseData, &products); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return products, nil
}
