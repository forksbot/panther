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

	"github.com/panther-labs/panther/api/gateway/resources/models"
)

// ListResourcesReader is a Reader for the ListResources structure.
type ListResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListResourcesOK creates a ListResourcesOK with default headers values
func NewListResourcesOK() *ListResourcesOK {
	return &ListResourcesOK{}
}

/*ListResourcesOK handles this case with default header values.

OK
*/
type ListResourcesOK struct {
	Payload *models.ResourceList
}

func (o *ListResourcesOK) Error() string {
	return fmt.Sprintf("[GET /list][%d] listResourcesOK  %+v", 200, o.Payload)
}

func (o *ListResourcesOK) GetPayload() *models.ResourceList {
	return o.Payload
}

func (o *ListResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourcesBadRequest creates a ListResourcesBadRequest with default headers values
func NewListResourcesBadRequest() *ListResourcesBadRequest {
	return &ListResourcesBadRequest{}
}

/*ListResourcesBadRequest handles this case with default header values.

Bad request
*/
type ListResourcesBadRequest struct {
	Payload *models.Error
}

func (o *ListResourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /list][%d] listResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *ListResourcesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourcesInternalServerError creates a ListResourcesInternalServerError with default headers values
func NewListResourcesInternalServerError() *ListResourcesInternalServerError {
	return &ListResourcesInternalServerError{}
}

/*ListResourcesInternalServerError handles this case with default header values.

Internal server error
*/
type ListResourcesInternalServerError struct {
}

func (o *ListResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /list][%d] listResourcesInternalServerError ", 500)
}

func (o *ListResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
