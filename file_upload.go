package grequests

import (
	"io"
)

// FileUpload is a struct that is used to specify the file that a User
// wishes to upload.
type FileUpload struct {
	// Filename is the name of the file that you wish to upload. We use this to guess the mimetype as well as pass it onto the server
	FileName string

	// FileContents is happy as long as you pass it a io.ReadCloser (which most file use anyways)
	FileContents io.ReadCloser

	// FieldName is form field name
	FieldName string

	// FileMime represents which mimetime should be sent along with the file.
	// When empty, defaults to application/octet-stream
	FileMime string
}

// FileUploadFromDisk allows you to create a FileUpload struct slice by just specifying a location on the disk
func FileUploadFromDisk(fileName string) ([]FileUpload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FileUploadFromGlob allows you to create a FileUpload struct slice by just specifying a glob location on the disk
// this function will gloss over all errors in the files and only upload the files that don't return errors from the glob
func FileUploadFromGlob(fileSystemGlob string) ([]FileUpload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ignoring error because I can stat the file
