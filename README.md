# api.video Go-sdk [![Build Status](https://travis-ci.com/OGFris/api.video.svg?branch=master)](https://travis-ci.com/OGFris/api.video) [![Go Report Card](https://goreportcard.com/badge/github.com/OGFris/api.video)](https://goreportcard.com/report/github.com/OGFris/api.video)
 Golang sdk for the api of [api.video](https://api.video), maintained by [Fris](https://twitter.com/FrisXYZ).

## Documentations
 [GoDocs](https://godoc.org/github.com/OGFris/api.video).
 
## Example
 *NOTICE:* **Make you sure to edit the login credentials with yours when using the example code** *!*
 
 ```go
package main

import "github.com/OGFris/api.video"

func main() {
	c := &apiVideo.Client{
 	    Username: "youremail@example.com",
 	    Password: "12345678",
        BaseUrl:  apiVideo.BaseUrl,
    }
 
    // Authenticate to get the tokens.
    err := c.Authenticate()
    if err != nil {
 	    panic(err)
    }
    
    
}
 ```

## License
 GoStats is under the [MIT License](https://github.com/OGFris/api.video/blob/master/LICENSE).

    MIT License

    Copyright (c) 2018 Fris

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:
 
    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.
    
    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
