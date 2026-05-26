package grequests

import (
	"bytes"
	"io"
	"net/http"
)

// Response is what is returned to a user when they fire off a request
type Response struct {

	// Ok is a boolean flag that validates that the server returned a 2xx code
	Ok bool

	// This is the Go error flag – if something went wrong within the request, this flag will be set.
	Error error

	// We want to abstract (at least at the moment) the Go http.Response object away. So we are going to make use of it
	// internal but not give the user access
	RawResponse *http.Response

	// StatusCode is the HTTP Status Code returned by the HTTP Response. Taken from resp.StatusCode
	StatusCode int

	// Header is a net/http/Header structure
	Header http.Header

	internalByteBuffer *bytes.Buffer
}

func buildResponse(resp *http.Response, err error) (*Response, error) {
	_ = "STUB: not implemented"
	// If the connection didn't succeed we just return a blank response
	return nil, nil
}

// If your code is within the 2xx range – the response is considered `Ok`

// EnsureResponseFinalized(goodResp) This will come back in 1.0

// Read is part of our ability to support io.ReadCloser if someone wants to make use of the raw body
func (r *Response) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close is part of our ability to support io.ReadCloser if someone wants to make use of the raw body
func (r *Response) Close() error { _ = "STUB: not implemented"; return nil }

// DownloadToFile allows you to download the contents of the response to a file
func (r *Response) DownloadToFile(fileName string) error { _ = "STUB: not implemented"; return nil }

// This is a noop if we use the internal ByteBuffer

// getInternalReader because we implement io.ReadCloser and optionally hold a large buffer of the response (created by
// the user's request)
func (r *Response) getInternalReader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// XML is a method that will populate a struct that is provided `userStruct` with the XML returned within the
// response body
func (r *Response) XML(userStruct interface{}, charsetReader XMLCharDecoder) error {
	_ = "STUB: not implemented"
	return nil
}

// JSON is a method that will populate a struct that is provided `userStruct` with the JSON returned within the
// response body
func (r *Response) JSON(userStruct interface{}) error { _ = "STUB: not implemented"; return nil }

// createResponseBytesBuffer is a utility method that will populate the internal byte reader – this is largely used for .String()
// and .Bytes()
func (r *Response) populateResponseByteBuffer() {
	_ = "STUB: not implemented"

	// Have I done this already?
	return
}

// Is there any content?

// Did the server tell us how big the response is going to be?

// Bytes returns the response as a byte array
func (r *Response) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// Are we still empty?

// String returns the response as a string
func (r *Response) String() string { _ = "STUB: not implemented"; return "" }

// ClearInternalBuffer is a function that will clear the internal buffer that we use to hold the .String() and .Bytes()
// data. Once you have used these functions – you may want to free up the memory.
func (r *Response) ClearInternalBuffer() { _ = "STUB: not implemented"; return }
