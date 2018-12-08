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
	"testing"
)

// TestClient_Authenticate tests if the function Authenticate works.
func TestClient_Authenticate(t *testing.T) {
	c := &Client{
		Username: "123@fakemail.com",
		Password: "12345678",
		BaseUrl:  BaseUrl,
	}

	err := c.Authenticate()
	if err.Error() != "bad request, the user credentials were incorrect" {
		panic(err)
	}
}

// TestClient_Refresh tests if the function Refresh works.
func TestClient_Refresh(t *testing.T) {
	c := &Client{
		Username: "123@fakemail.com",
		Password: "12345678",
		BaseUrl:  BaseUrl,
	}

	err := c.Refresh()
	if err.Error() != "bad request, the user credentials were incorrect" {
		panic(err)
	}
}
