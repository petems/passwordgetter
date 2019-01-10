package cmd_test

import (
	"errors"
	"testing"

	"github.com/petems/passwordgetter/cmd"
	"github.com/stretchr/testify/assert"
)

type stubPasswordReader struct {
	Password    string
	ReturnError bool
}

func (pr stubPasswordReader) ReadPassword() (string, error) {
	if pr.ReturnError {
		return "", errors.New("stubbed error")
	}
	return pr.Password, nil
}

func TestRunReturnsErrorWhenReadPasswordFails(t *testing.T) {
	pr := stubPasswordReader{ReturnError: true}
	result, err := cmd.Run(pr)
	assert.Error(t, err)
	assert.Equal(t, errors.New("stubbed error"), err)
	assert.Equal(t, "", result)
}

func TestRunReturnsPasswordInput(t *testing.T) {
	pr := stubPasswordReader{Password: "password"}
	result, err := cmd.Run(pr)
	assert.NoError(t, err)
	assert.Equal(t, "password", result)
}
