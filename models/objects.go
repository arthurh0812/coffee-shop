package models

import (
	"errors"
	"math/rand"
)

var ErrCannotConvertToObjectID = errors.New("cannot convert given string to ObjectID")

const letterAndDigitRunes = "abcdefghijklmnopqvwxyzABCDEFGHIJKLMNOPQVWXYZ1234567890"

// ObjectID is a 16-byte sequence
type ObjectID string

func NewObjectID() ObjectID {
	b := make([]byte, 16, 16)
	ln := len(letterAndDigitRunes)
	for i := range b {
		b[i] = letterAndDigitRunes[rand.Intn(ln)]
	}
	return ObjectID(b)
}

func ToObjectID(source string) (ObjectID, error) {
	bytes := source[:]
	if len(bytes) != 16 {
		return "", ErrCannotConvertToObjectID
	}
	var id = make([]byte, 16, 16)
	copy(id, bytes)
	return ObjectID(id), nil
}
