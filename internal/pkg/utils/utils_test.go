// Copyright 2020 Hewlett Packard Enterprise Development LP

package utils

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hpe-hcss/iam-lib/pkg/spaces-client"
)

func Test_GetSpaceID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockContext := context.Background()
	mockSpace := "mockSpace"
	mockToken := "mockToken"
	mockSpaceID := "mockID123456"
	mockSpaceClient := spacesclient.NewMockIamSpacesAPI(ctrl)
	mockListSpaceOutput := spacesclient.ListSpaceOutput{Members: []spacesclient.Space{{Name: mockSpace, ID: mockSpaceID}}}
	mockSpaceClient.EXPECT().List(mockContext, mockToken, spacesclient.ListSpaceInput{Name: mockSpace}).Return(mockListSpaceOutput, nil)
	resID, err := GetSpaceID(mockContext, mockSpaceClient, mockToken, mockSpace)
	assert.Nil(t, err)
	assert.Equal(t, mockSpaceID, resID)

	errMsg := "Failed to get Space Details"
	mockSpaceClient.EXPECT().List(mockContext, mockToken, spacesclient.ListSpaceInput{Name: mockSpace}).Return(spacesclient.ListSpaceOutput{},
		errors.New(errMsg))
	resID, err = GetSpaceID(mockContext, mockSpaceClient, mockToken, mockSpace)
	assert.Equal(t, "", resID)
	assert.Contains(t, err.Error(), errMsg)

	mockSpaceClient.EXPECT().List(mockContext, mockToken, spacesclient.ListSpaceInput{Name: "errorSpace"}).Return(mockListSpaceOutput, nil)
	resID, err = GetSpaceID(mockContext, mockSpaceClient, mockToken, "errorSpace")
	assert.NotNil(t, err)
	assert.Equal(t, "", resID)
}
