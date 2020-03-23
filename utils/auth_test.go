package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePassword(t *testing.T) {
	password := "control123"
	hash := CreatePassword(password)

	assert.NotNil(t, hash)
}

func TestCompareHashAndPassword(t *testing.T) {
	password := "control123"
	hash := CreatePassword(password)

	result := CompareHashAndPassword([]byte(hash), []byte(password))

	assert.Equal(t, true, result)
}

func TestCompareHashAndPassword_InvalidPass(t *testing.T) {
	password := "control123"
	hash := CreatePassword(password)
	invalidPass := "testinvalid"
	result := CompareHashAndPassword([]byte(hash), []byte(invalidPass))

	assert.Equal(t, false, result)
}
