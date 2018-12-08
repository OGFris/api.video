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
	"io/ioutil"
	"net/http"
	"net/url"
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
	req, err := http.NewRequest("POST", c.BaseUrl+VideosPath, nil)
	if err != nil {
		return &Video{}, err
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprint(c.TokenType, c.AccessToken))
	if values, err := u.ToUrlValues(); err == nil {
		if !fromSource {
			values.Del("source")
		}
		req.PostForm = values
	} else {
		return &Video{}, err
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return &Video{}, err
	}

	switch response.StatusCode {

	case http.StatusCreated:
		// Video created
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return &Video{}, err
		}

		var data Video

		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return &Video{}, err
		}

		break

	case http.StatusAccepted:
		// Source video accepted and is being downloaded
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return &Video{}, err
		}

		var data Video

		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return &Video{}, err
		}

		break

	case http.StatusBadRequest:
		// Error
		return &Video{}, errors.New("bad request")

	default:
		return &Video{}, errors.New("unexpected response status, report it to api.video")
	}

	// This line of code should never be reached
	return &Video{}, errors.New("line of code should never be reached")
}

// ToUrlValues converts the UploadVideo type to url.Values type.
func (u *UploadVideo) ToUrlValues() (url.Values, error) {
	values := url.Values{}
	bytes, err := json.Marshal(&u)
	if err != nil {
		return values, err
	}

	values, err = url.ParseQuery(string(bytes))

	return values, nil
}
