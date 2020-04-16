// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// TestPolicyReader is a Reader for the TestPolicy structure.
type TestPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTestPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestPolicyOK creates a TestPolicyOK with default headers values
func NewTestPolicyOK() *TestPolicyOK {
	return &TestPolicyOK{}
}

/*TestPolicyOK handles this case with default header values.

OK
*/
type TestPolicyOK struct {
	Payload *models.TestPolicyResult
}

func (o *TestPolicyOK) Error() string {
	return fmt.Sprintf("[POST /test][%d] testPolicyOK  %+v", 200, o.Payload)
}

func (o *TestPolicyOK) GetPayload() *models.TestPolicyResult {
	return o.Payload
}

func (o *TestPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestPolicyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestPolicyBadRequest creates a TestPolicyBadRequest with default headers values
func NewTestPolicyBadRequest() *TestPolicyBadRequest {
	return &TestPolicyBadRequest{}
}

/*TestPolicyBadRequest handles this case with default header values.

Bad request
*/
type TestPolicyBadRequest struct {
	Payload *models.Error
}

func (o *TestPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /test][%d] testPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *TestPolicyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestPolicyInternalServerError creates a TestPolicyInternalServerError with default headers values
func NewTestPolicyInternalServerError() *TestPolicyInternalServerError {
	return &TestPolicyInternalServerError{}
}

/*TestPolicyInternalServerError handles this case with default header values.

Internal server error
*/
type TestPolicyInternalServerError struct {
}

func (o *TestPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /test][%d] testPolicyInternalServerError ", 500)
}

func (o *TestPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
