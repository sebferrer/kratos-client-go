// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/models"
)

// GetSelfServiceBrowserRecoveryRequestReader is a Reader for the GetSelfServiceBrowserRecoveryRequest structure.
type GetSelfServiceBrowserRecoveryRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceBrowserRecoveryRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceBrowserRecoveryRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSelfServiceBrowserRecoveryRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSelfServiceBrowserRecoveryRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSelfServiceBrowserRecoveryRequestGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceBrowserRecoveryRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSelfServiceBrowserRecoveryRequestOK creates a GetSelfServiceBrowserRecoveryRequestOK with default headers values
func NewGetSelfServiceBrowserRecoveryRequestOK() *GetSelfServiceBrowserRecoveryRequestOK {
	return &GetSelfServiceBrowserRecoveryRequestOK{}
}

/*GetSelfServiceBrowserRecoveryRequestOK handles this case with default header values.

recoveryRequest
*/
type GetSelfServiceBrowserRecoveryRequestOK struct {
	Payload *models.RecoveryRequest
}

func (o *GetSelfServiceBrowserRecoveryRequestOK) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/recovery][%d] getSelfServiceBrowserRecoveryRequestOK  %+v", 200, o.Payload)
}

func (o *GetSelfServiceBrowserRecoveryRequestOK) GetPayload() *models.RecoveryRequest {
	return o.Payload
}

func (o *GetSelfServiceBrowserRecoveryRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecoveryRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRecoveryRequestForbidden creates a GetSelfServiceBrowserRecoveryRequestForbidden with default headers values
func NewGetSelfServiceBrowserRecoveryRequestForbidden() *GetSelfServiceBrowserRecoveryRequestForbidden {
	return &GetSelfServiceBrowserRecoveryRequestForbidden{}
}

/*GetSelfServiceBrowserRecoveryRequestForbidden handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRecoveryRequestForbidden struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRecoveryRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/recovery][%d] getSelfServiceBrowserRecoveryRequestForbidden  %+v", 403, o.Payload)
}

func (o *GetSelfServiceBrowserRecoveryRequestForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRecoveryRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRecoveryRequestNotFound creates a GetSelfServiceBrowserRecoveryRequestNotFound with default headers values
func NewGetSelfServiceBrowserRecoveryRequestNotFound() *GetSelfServiceBrowserRecoveryRequestNotFound {
	return &GetSelfServiceBrowserRecoveryRequestNotFound{}
}

/*GetSelfServiceBrowserRecoveryRequestNotFound handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRecoveryRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRecoveryRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/recovery][%d] getSelfServiceBrowserRecoveryRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetSelfServiceBrowserRecoveryRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRecoveryRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRecoveryRequestGone creates a GetSelfServiceBrowserRecoveryRequestGone with default headers values
func NewGetSelfServiceBrowserRecoveryRequestGone() *GetSelfServiceBrowserRecoveryRequestGone {
	return &GetSelfServiceBrowserRecoveryRequestGone{}
}

/*GetSelfServiceBrowserRecoveryRequestGone handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRecoveryRequestGone struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRecoveryRequestGone) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/recovery][%d] getSelfServiceBrowserRecoveryRequestGone  %+v", 410, o.Payload)
}

func (o *GetSelfServiceBrowserRecoveryRequestGone) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRecoveryRequestGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRecoveryRequestInternalServerError creates a GetSelfServiceBrowserRecoveryRequestInternalServerError with default headers values
func NewGetSelfServiceBrowserRecoveryRequestInternalServerError() *GetSelfServiceBrowserRecoveryRequestInternalServerError {
	return &GetSelfServiceBrowserRecoveryRequestInternalServerError{}
}

/*GetSelfServiceBrowserRecoveryRequestInternalServerError handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRecoveryRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRecoveryRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/recovery][%d] getSelfServiceBrowserRecoveryRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSelfServiceBrowserRecoveryRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRecoveryRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
