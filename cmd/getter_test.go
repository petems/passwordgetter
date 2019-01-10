package cmd_test

import (
	"testing"
	"errors"

	"github.com/golang/mock/gomock"

	"github.com/petems/passwordgetter/mocks"
	"github.com/petems/passwordgetter/cmd"
	"github.com/stretchr/testify/assert"
)

func TestRunReturnsStringFromReadPassword(t *testing.T) {

	mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockPasswordReader := mocks.NewMockPasswordReader(mockCtrl)

  mockPasswordReader.EXPECT().ReadPassword().Return("hunter2", nil).Times(1)

	result, err := cmd.Run(mockPasswordReader)
	assert.NoError(t, err)
	assert.Equal(t, "hunter2", result)
}

func TestRunReturnsErrorWhenEmptyString(t *testing.T) {

	mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockPasswordReader := mocks.NewMockPasswordReader(mockCtrl)

  mockPasswordReader.EXPECT().ReadPassword().Return("", nil).Times(1)

	result, err := cmd.Run(mockPasswordReader)
	assert.Error(t, err)
	assert.Equal(t, "", result)
}

func TestRunReturnsErrorWhenReadPasswordReturnsError(t *testing.T) {

	mockCtrl := gomock.NewController(t)
  defer mockCtrl.Finish()

  mockPasswordReader := mocks.NewMockPasswordReader(mockCtrl)

  mockPasswordReader.EXPECT().ReadPassword().Return("", errors.New("stubbed error")).Times(1)

	result, err := cmd.Run(mockPasswordReader)
	assert.Error(t, err)
	assert.Equal(t, "", result)
}



