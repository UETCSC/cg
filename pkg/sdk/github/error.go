package github

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrVersionNotFound = errors.New("version not found")
	ErrPackageNotFound = errors.New("installation package not found")
)

var (
	ErrUnsupportedChecksumAlgorithm = errors.New("unsupported checksum algorithm")
	ErrChecksumNotMatched = errors.New("file checksum does not match the computed checksum")
	ErrChecksumFileNotFound = errors.New("checksum file not found")
)

type URLUnreachableError struct {
	err error
	url string
}

func NewURLUnreachableError(url string, err error) error {
	return &URLUnreachableError{
		err: err,
		url: url,
	}
}

func (e URLUnreachableError) Error() string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("URL %q is unreachable", e.url))
	if e.err != nil {
		buf.WriteString(" ==> " + e.err.Error())
	}
	return buf.String()
}

func (e URLUnreachableError) Err() error {
	return e.err
}

func (e URLUnreachableError) URL() string {
	return e.url
}

type DownloadError struct {
	url string
	err error
}

func NewDownloadError(url string, err error) error {
	return &DownloadError{
		url: url,
		err: err,
	}
}

func (e DownloadError) Error() string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("Resource(%s) download failed", e.url))
	if e.err != nil {
		buf.WriteString(" ==> " + e.err.Error())
	}
	return buf.String()
}

func (e DownloadError) Err() error {
	return e.err
}

func (e DownloadError) URL() string {
	return e.url
}
