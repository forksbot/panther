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

// SuppressReader is a Reader for the Suppress structure.
type SuppressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuppressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuppressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSuppressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSuppressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSuppressOK creates a SuppressOK with default headers values
func NewSuppressOK() *SuppressOK {
	return &SuppressOK{}
}

/*SuppressOK handles this case with default header values.

OK
*/
type SuppressOK struct {
}

func (o *SuppressOK) Error() string {
	return fmt.Sprintf("[POST /suppress][%d] suppressOK ", 200)
}

func (o *SuppressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSuppressBadRequest creates a SuppressBadRequest with default headers values
func NewSuppressBadRequest() *SuppressBadRequest {
	return &SuppressBadRequest{}
}

/*SuppressBadRequest handles this case with default header values.

Bad request
*/
type SuppressBadRequest struct {
	Payload *models.Error
}

func (o *SuppressBadRequest) Error() string {
	return fmt.Sprintf("[POST /suppress][%d] suppressBadRequest  %+v", 400, o.Payload)
}

func (o *SuppressBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SuppressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuppressInternalServerError creates a SuppressInternalServerError with default headers values
func NewSuppressInternalServerError() *SuppressInternalServerError {
	return &SuppressInternalServerError{}
}

/*SuppressInternalServerError handles this case with default header values.

Internal server error
*/
type SuppressInternalServerError struct {
}

func (o *SuppressInternalServerError) Error() string {
	return fmt.Sprintf("[POST /suppress][%d] suppressInternalServerError ", 500)
}

func (o *SuppressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
