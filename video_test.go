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
	"fmt"
	"testing"
)

// TestUploadVideo_ToJson tests if the function ToJson works.
func TestUploadVideo_ToJson(t *testing.T) {
	u := &UploadVideo{
		Title:       "test",
		Description: "test 123",
		Tags:        []string{"science", "finance", "stuff ig"},
		Metadata: []struct {
			Key   string `json:"key"`
			Value string `json:""`
		}{
			{Key: "test key 1", Value: "test value 1"},
			{Key: "test key 2", Value: "test value 2"},
		},
	}

	_, err := u.ToJson()
	if err != nil {
		panic(err)
	}
}

// TestClient_CreateVideo tests if the function CreateVideo works.
func TestClient_CreateVideo(t *testing.T) {
	// Correct input, should work.
	c := LoadClientFromEnv()
	if c.Password != "" && c.Username != "" {
		err := c.Authenticate()
		if err != nil {
			panic(err)
		}

		video, err := c.CreateVideo(&UploadVideo{
			Title:  "test",
			Source: "https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_10mb.mp4",
			Tags:   []string{"test", "api"},
			Metadata: []struct {
				Key   string `json:"key"`
				Value string `json:""`
			}{
				{Key: "Author", Value: "Fris"},
			},
		}, true)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success.")
		fmt.Println(video)
	}
}
