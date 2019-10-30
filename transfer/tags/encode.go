package tags

import (
	"errors"

	"github.com/eknkc/basex"
)

// ErrImageTooLong represents a too long fqdn error
var ErrImageTooLong = errors.New("Image FQDN is too long to encode")

const dictionary string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Encode will encode a fqdn into a compatible tag
func Encode(fqdn string) (string, error) {
	bx, err := basex.NewEncoding(dictionary)
	if err != nil {
		return "", err
	}
	encoded := bx.Encode([]byte(fqdn))
	if len(encoded) > 128 {
		return "", ErrImageTooLong
	}
	return encoded, nil
}

// Decode will decode a compatible tag into a fqdn
func Decode(tag string) (string, error) {
	bx, err := basex.NewEncoding(dictionary)
	if err != nil {
		return "", err
	}
	decoded, err := bx.Decode(tag)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
