package zettle

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	BasePath    string
	ClientId    string
	ApiKey      string
	accessToken *string
}

type ServiceNewParams struct {
	ClientId string
	ApiKey   string
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type updateProductParams struct {
	UpdateProductParams
	productUuid string
}

func New(params ServiceNewParams) (*Service, error) {
	return &Service{
		BasePath: "/organizations/self",
		ClientId: params.ClientId,
		ApiKey:   params.ApiKey,
	}, nil
}

func (service Service) getAccessToken() (*string, error) {
	if service.accessToken != nil {
		token, err := jwt.Parse(*service.accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte{}, nil
		}, jwt.WithLeeway(5*time.Minute))

		if token.Valid {
			return service.accessToken, nil
		}

		// if expired, fetch a new token
		if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
			slog.Error(err.Error())
			return nil, err
		}
	}

	client := &http.Client{}
	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	data.Set("client_id", service.ClientId)
	data.Set("assertion", service.ApiKey)

	request, err := http.NewRequest(http.MethodPost, "https://oauth.zettle.com/token", strings.NewReader(data.Encode()))
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code from getAccessToken %d", response.StatusCode)
	}

	var accessTokenResponse *accessTokenResponse
	if err := json.NewDecoder(response.Body).Decode(&accessTokenResponse); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	service.accessToken = &accessTokenResponse.AccessToken

	return service.accessToken, nil
}

func (service Service) GetProducts() (*ProductResponse, error) {
	accessToken, err := service.getAccessToken()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/products/v2", service.BasePath), nil)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *accessToken))

	response, err := client.Do(request)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	var products *ProductResponse
	if err := json.NewDecoder(response.Body).Decode(&products); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return products, nil
}

// TODO: Actually we don't need it, but I want to talk to Pit about it <3
func (service Service) UpdateProduct(params updateProductParams, productUpdate FullProductUpdateRequest) error {
	accessToken, err := service.getAccessToken()
	if err != nil {
		return err
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/products/v2/%s", service.BasePath, params.productUuid), nil)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *accessToken))

	response, err := client.Do(request)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusBadRequest {
		var errorResponse *ErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
			slog.Error(err.Error())
			return err
		}

		// TODO: How to print violations? Should it be in the error log?
		err := fmt.Errorf("Unable to update product. Got message %s", *errorResponse.DeveloperMessage)
		slog.Error(err.Error())
		return err
	}

	if response.StatusCode == http.StatusPreconditionFailed {
		err := errors.New("Unable to update product. Invalid ETag provided.")
		slog.Error(err.Error())
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		err := errors.New("Unable to update product. Unexpected status code.")
		slog.Error(err.Error())
		return err
	}

	// TODO: Update product with ETag and Location
	return nil
}
