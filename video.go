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
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	// VideosPath is the path for creating videos.
	VideosPath = "/videos"
)

// Video is the model for the response received from the api.
type Video struct {
	VideoId     string   `json:"videoId"`
	PlayerId    string   `json:"playerId"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Public      bool     `json:"public"`
	Tags        []string `json:"tags"`
	Metadata    []struct {
		Key   string `json:"key"`
		Value string `json:""`
	} `json:"metadata"`
	PublishedAt string `json:"publishedAt"`
	Source      struct {
		Uri string `json:"uri"`
	} `json:"source"`
	Assets struct {
		Iframe    string `json:"iframe"`
		Player    string `json:"player"`
		Hls       string `json:"hls"`
		Thumbnail string `json:"thumbnail"`
	} `json:"assets"`
}

// UploadVideo is the model of the post form that is sent to the api when creating a new video.
type UploadVideo struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Metadata    []struct {
		Key   string `json:"key"`
		Value string `json:""`
	} `json:"metadata"`
	Public   bool   `json:"public"`
	Source   string `json:"source"`
	PlayerId string `json:"playerId"`
}

// CreateVideo creates a new video using the client.
func (c *Client) CreateVideo(u *UploadVideo, fromSource bool) (*Video, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", c.TokenType+" "+c.AccessToken)
	req.SetRequestURI(c.BaseUri + VideosPath)

	if data, err := u.ToJson(); err == nil {
		req.SetBodyString(data)
	} else {
		return &Video{}, err
	}

	client := fasthttp.Client{}
	response := fasthttp.AcquireResponse()

	err := client.Do(req, response)
	if err != nil {
		return &Video{}, err
	}

	switch response.StatusCode() {

	case fasthttp.StatusCreated:
		// Video created
		var data Video

		err = json.Unmarshal(response.Body(), &data)
		if err != nil {
			return &Video{}, err
		}

		break

	case fasthttp.StatusAccepted:
		// Source video accepted and is being downloaded
		var data Video

		err = json.Unmarshal(response.Body(), &data)
		if err != nil {
			return &Video{}, err
		}

		break

	case fasthttp.StatusBadRequest:
		// Error
		fmt.Println(string(response.Body()))
		return &Video{}, errors.New("bad request")

	default:
		fmt.Println(response.StatusCode())
		fmt.Println(string(response.Body()))
		return &Video{}, errors.New("unexpected response status, report it to api.video")
	}

	// This line of code should never be reached
	return &Video{}, errors.New("line of code should never be reached")
}

// ToJson converts the UploadVideo type to a json encoded string.
func (u *UploadVideo) ToJson() (string, error) {
	bytes, err := json.Marshal(&u)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
