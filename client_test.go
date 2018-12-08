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
	"os"
	"testing"
)

// TestLoadClientFromEnv tests if the function LoadClientFromEnv works.
func TestLoadClientFromEnv(t *testing.T) {
	oldUsername := os.Getenv("APIVIDEO_USERNAME")
	oldPassword := os.Getenv("APIVIDEO_PASSWORD")

	err := os.Setenv("APIVIDEO_USERNAME", "fakeUsername")
	if err != nil {
		panic(err)
	}

	err = os.Setenv("APIVIDEO_PASSWORD", "fakePassword")
	if err != nil {
		panic(err)
	}

	c := LoadClientFromEnv()

	err = os.Setenv("APIVIDEO_USERNAME", oldUsername)
	if err != nil {
		panic(err)
	}

	err = os.Setenv("APIVIDEO_PASSWORD", oldPassword)
	if err != nil {
		panic(err)
	}
	if c.Username != "fakeUsername" || c.Password != "fakePassword" {
		t.Fatal("Username and/or Password does not match the provided data.")
		t.FailNow()
	}
}

// TestClient_Authenticate tests if the function Authenticate works.
func TestClient_Authenticate(t *testing.T) {
	// Correct input, should work.
	c := LoadClientFromEnv()
	if c.Password != "" && c.Username != "" {
		err := c.Authenticate()
		if err != nil {
			panic(err)
		}
	}

	// Fake input, should not work.
	c = &Client{
		Username: "123@fakemail.com",
		Password: "12345678",
		BaseUri:  BaseUri,
	}

	err := c.Authenticate()
	if err.Error() != "bad request, the user credentials were incorrect" {
		panic(err)
	}
}

// TestClient_Refresh tests if the function Refresh works.
func TestClient_Refresh(t *testing.T) {
	// Correct input, should work.
	c := LoadClientFromEnv()
	if c.Password != "" && c.Username != "" {
		err := c.Authenticate()
		if err != nil {
			panic(err)
		}

		err = c.Refresh()
		if err != nil {
			panic(err)
		}
	}

	// Fake input, should not work.
	c = &Client{
		Username: "123@fakemail.com",
		Password: "12345678",
		BaseUri:  BaseUri,
	}

	err := c.Refresh()
	if err.Error() != "bad request, the user credentials were incorrect" {
		panic(err)
	}
}
