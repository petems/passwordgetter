package cmd_test

import (
	"errors"
	"testing"

	"github.com/petems/passwordgetter/cmd"
	"github.com/stretchr/testify/assert"
)

type errorPasswordReader struct {
}

func (pr errorPasswordReader) ReadPassword() (string, error) {
	return "", errors.New("stubbed error")
}

type stubPasswordReader struct {
	Password string
}

func (pr stubPasswordReader) ReadPassword() (string, error) {
	return pr.Password, nil
}

func TestRunReturnsErrorWhenReadPasswordFails(t *testing.T) {
	var errorReader errorPasswordReader
	result, err := cmd.Run(errorReader)
	assert.Equal(t, errors.New("stubbed error"), err)
	assert.Equal(t, "", result)
}

func TestRunReturnsPasswordInput(t *testing.T) {
	pr := stubPasswordReader{Password: "password"}
	result, err := cmd.Run(pr)
	assert.NoError(t, err)
	assert.Equal(t, "password", result)
}

func TestRunReturnsErrorWhenEmptyPasswordIsProvided(t *testing.T) {
	pr := stubPasswordReader{Password: ""}
	result, err := cmd.Run(pr)
	assert.Error(t, err)
	assert.Equal(t, "", result)
}
