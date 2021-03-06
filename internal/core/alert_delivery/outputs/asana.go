package outputs

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
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"go.uber.org/zap"

	outputmodels "github.com/panther-labs/panther/api/lambda/outputs/models"
	alertmodels "github.com/panther-labs/panther/internal/core/alert_delivery/models"
)

const (
	asanaCreateTaskURL             = "https://app.asana.com/api/1.0/tasks"
	asanaAuthorizationHeaderFormat = "Bearer %s"
)

// Asana creates a task in Asana projects
func (client *OutputClient) Asana(alert *alertmodels.Alert, config *outputmodels.AsanaConfig) *AlertDeliveryError {
	zap.L().Debug("sending alert to Asana")
	payload := map[string]interface{}{
		"data": map[string]interface{}{
			"name":     generateAlertTitle(alert),
			"projects": config.ProjectGids,
			"notes":    generateDetailedAlertMessage(alert),
		},
	}

	postInput := &PostInput{
		url:  asanaCreateTaskURL,
		body: payload,
		headers: map[string]string{
			AuthorizationHTTPHeader: fmt.Sprintf(asanaAuthorizationHeaderFormat, aws.StringValue(config.PersonalAccessToken)),
		},
	}
	return client.httpWrapper.post(postInput)
}
