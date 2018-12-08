// MIT License
//
// Copyright (c) 2018 Fris
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package apiVideo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// BaseUrl is the main url to the api.
	BaseUrl = "https://ws.api.video"

	// ApiKeyPath is the path for the authentication.
	ApiKeyPath = "/auth/api-key"

	// RefreshPath is the path for the token refresh.
	RefreshPath = "/auth/refresh"
)

// Client is the instance that is used to communicate with the api.
type Client struct {
	// Username stands for the email you used to register your account.
	Username string

	// Password stands for the api token you received in the registration email.
	Password string

	// BaseUrl should always be the constant BaseUrl.
	BaseUrl string

	// TokenType is provided inside the response to the authentication.
	TokenType string

	// ExpiresIn is provided inside the response to the authentication, it is the timeout for the AccessToken.
	ExpiresIn time.Time

	// AccessToken is provided inside the response to the authentication, it is used to make requests to the api.
	AccessToken string

	// RefreshToken is provided inside the response to the authentication, it is used to update the AccessToken when-
	// expired.
	RefreshToken string
}

// AuthResponse is the model of the authentication and refresh responses.
type AuthResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Authenticate has to be executed in order to use the api, it authenticates the client to get the token.
func (c *Client) Authenticate() error {
	req, err := http.NewRequest("POST", c.BaseUrl+ApiKeyPath, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-type", "application/json")

	req.PostForm = url.Values{"apiKey": {c.Password}}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusBadRequest {
		return errors.New("bad request, the user credentials were incorrect")
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("status code does not seem to be correct")
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var data AuthResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	c.TokenType = data.TokenType
	c.ExpiresIn = time.Now().Add(time.Duration(data.ExpiresIn * 1000))
	c.AccessToken = data.AccessToken
	c.RefreshToken = data.RefreshToken

	return nil
}

// Refresh is used to update the access token using the refresh token on the api
func (c *Client) Refresh() error {
	req, err := http.NewRequest("POST", c.BaseUrl+RefreshPath, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-type", "application/json")

	req.PostForm = url.Values{"refreshToken": {c.RefreshToken}}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusBadRequest {
		return errors.New("bad request, the user credentials were incorrect")
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("status code does not seem to be correct")
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var data AuthResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	c.TokenType = data.TokenType
	c.ExpiresIn = time.Now().Add(time.Duration(data.ExpiresIn * 1000))
	c.AccessToken = data.AccessToken
	c.RefreshToken = data.RefreshToken

	return nil
}
