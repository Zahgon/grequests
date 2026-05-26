// Package grequests implements a friendly API over Go's existing net/http library
package grequests

import "context"

// Get takes 2 parameters and returns a Response Struct. These two options are:
//  1. A URL
//  2. A set of options for the request
func Get(ctx context.Context, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A set of options for the request
//
// If you do not intend to use the `RequestOptions` you can just pass nil
func Put(ctx context.Context, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Patch takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A set of options for the request
//
// If you do not intend to use the `RequestOptions` you can just pass nil
func Patch(ctx context.Context, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A set of options for the request
//
// If you do not intend to use the `RequestOptions` you can just pass nil
func Delete(ctx context.Context, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Post takes 2 parameters and returns a Response channel. These two options are:
//  1. A URL
//  2. A set of options for the request
//
// If you do not intend to use the `RequestOptions` you can just pass nil
func Post(ctx context.Context, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Head takes 2 parameters and returns a Response channel. These two options are:
//  1. A URL
//  2. A set of options for the request
//
// If you do not intend to use the `RequestOptions` you can just pass nil
func Head(ctx context.Context, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Options takes 2 parameters and returns a Response struct. These two options are:
//  1. A URL
//  2. A set of options for the request
//
// If you do not intend to use the `RequestOptions` you can just pass nil
func Options(ctx context.Context, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Request takes 3 parameters and returns a Response Struct. These three options are:
//  1. A verb
//  2. A URL
//  3. A set of options for the request
//
// If you do not intend to use the `RequestOptions` you can just pass nil
func Request(ctx context.Context, verb, url string, options ...Option) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
