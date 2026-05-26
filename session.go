package grequests

import (
	"context"
	"net/http"
)

// Session allows a user to make use of persistent cookies in between
// HTTP requests
type Session struct {
	// RequestOptions is global options
	RequestOptions *RequestOptions

	// HTTPClient is the client that we will use to request the resources
	HTTPClient *http.Client
}

// NewSession returns a session struct which enables can be used to maintain establish a persistent state with the
// server
// This function will set UseCookieJar to true as that is the purpose of using the session
func NewSession(ro *RequestOptions) *Session { _ = "STUB: not implemented"; return nil }

// Combine session options and request options
// 1. UserAgent
// 2. Host
// 3. Auth
// 4. Headers
func (s *Session) combineRequestOptions(ro *RequestOptions) *RequestOptions {
	_ = "STUB: not implemented"
	return nil
}

// Get takes 2 parameters and returns a Response Struct. These two options are:
//  1. A URL
//  2. A RequestOptions struct
//
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Get(ctx context.Context, url string, ro *RequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A RequestOptions struct
//
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Put(ctx context.Context, url string, ro *RequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Patch takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A RequestOptions struct
//
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Patch(ctx context.Context, url string, ro *RequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A RequestOptions struct
//
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Delete(ctx context.Context, url string, ro *RequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Post takes 2 parameters and returns a Response channel. These two options are:
//  1. A URL
//  2. A RequestOptions struct
//
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Post(ctx context.Context, url string, ro *RequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Head takes 2 parameters and returns a Response channel. These two options are:
//  1. A URL
//  2. A RequestOptions struct
//
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Head(ctx context.Context, url string, ro *RequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Options takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A RequestOptions struct
//
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Options(ctx context.Context, url string, ro *RequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CloseIdleConnections closes the idle connections that a session client may make use of
func (s *Session) CloseIdleConnections() { _ = "STUB: not implemented"; return }
