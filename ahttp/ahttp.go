// Copyright (c) Jeevanandam M (https://github.com/jeevatkm)
// go-aah/aah source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

// Package ahttp is to cater HTTP helper methods for aah framework.
// Like parse HTTP headers, ResponseWriter, content type, etc.
package ahttp

import "net/http"

// HTTP Method names
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
)

type (
	// Locale value is negotiated from HTTP header `Accept-Language`
	Locale struct {
		Raw      string
		Language string
		Region   string
	}

	// ContentType is represents request and response content type values
	ContentType struct {
		Raw    string
		Mime   string
		Exts   []string
		Params map[string]string
	}

	// Request is extends `http.Request` for aah framework
	Request struct {
		*http.Request
	}

	// Response is extends `http.Response` for aah framework
	Response struct {
		http.Response
	}
)