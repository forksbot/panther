package api

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/panther-labs/panther/api/lambda/outputs/models"
	"github.com/panther-labs/panther/internal/core/outputs_api/table"
)

func TestGetOutput(t *testing.T) {
	mockOutputsTable := &mockOutputTable{}
	outputsTable = mockOutputsTable
	mockEncryptionKey := &mockEncryptionKey{}
	encryptionKey = mockEncryptionKey

	mockGetOutputInput := &models.GetOutputInput{
		OutputID: aws.String("outputId"),
	}

	alertOutputItem := &table.AlertOutputItem{
		OutputID:           aws.String("outputId"),
		DisplayName:        aws.String("displayName"),
		CreatedBy:          aws.String("createdBy"),
		CreationTime:       aws.String("creationTime"),
		LastModifiedBy:     aws.String("lastModifiedBy"),
		LastModifiedTime:   aws.String("lastModifiedTime"),
		OutputType:         aws.String("slack"),
		EncryptedConfig:    make([]byte, 1),
		DefaultForSeverity: aws.StringSlice([]string{"HIGH", "CRITICAL"}),
	}
	mockEncryptionKey.On("DecryptConfig", make([]byte, 1), mock.Anything).Return(nil)
	mockOutputsTable.On("GetOutput", aws.String("outputId")).Return(alertOutputItem, nil)

	expectedAlertOutput := &models.AlertOutput{
		OutputID:           aws.String("outputId"),
		OutputType:         aws.String("slack"),
		CreatedBy:          aws.String("createdBy"),
		CreationTime:       aws.String("creationTime"),
		DisplayName:        aws.String("displayName"),
		LastModifiedBy:     aws.String("lastModifiedBy"),
		LastModifiedTime:   aws.String("lastModifiedTime"),
		OutputConfig:       &models.OutputConfig{},
		DefaultForSeverity: aws.StringSlice([]string{"HIGH", "CRITICAL"}),
	}

	result, err := (API{}).GetOutput(mockGetOutputInput)

	require.NoError(t, err)
	assert.Equal(t, expectedAlertOutput, result)
	mockOutputsTable.AssertExpectations(t)
	mockEncryptionKey.AssertExpectations(t)
}
