package handler

import (
	"errors"
	"strings"
)

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// Handler is a concrete implementation of StringService
type Handler struct{}

func (Handler) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToUpper(s), nil
}

func (Handler) Lowercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToLower(s), nil
}

func (Handler) Concatenate(s string, c string) (string, error) {
	if s == "" || c == "" {
		return "", ErrEmpty
	}

	return s + c, nil
}

func (Handler) Split(s string, k string) ([]string, error) {
	if s == "" || k == "" {
		return nil, ErrEmpty
	}

	return strings.Split(s, k), nil
}

func (Handler) Count(s string) int {
	return len(s)
}
