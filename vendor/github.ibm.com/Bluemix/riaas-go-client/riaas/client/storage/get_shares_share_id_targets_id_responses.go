// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetSharesShareIDTargetsIDReader is a Reader for the GetSharesShareIDTargetsID structure.
type GetSharesShareIDTargetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSharesShareIDTargetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSharesShareIDTargetsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSharesShareIDTargetsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSharesShareIDTargetsIDOK creates a GetSharesShareIDTargetsIDOK with default headers values
func NewGetSharesShareIDTargetsIDOK() *GetSharesShareIDTargetsIDOK {
	return &GetSharesShareIDTargetsIDOK{}
}

/*GetSharesShareIDTargetsIDOK handles this case with default header values.

The the specified tag exists on the specified resource.
*/
type GetSharesShareIDTargetsIDOK struct {
	Payload *models.Sharetarget
}

func (o *GetSharesShareIDTargetsIDOK) Error() string {
	return fmt.Sprintf("[GET /shares/{share_id}/targets/{id}][%d] getSharesShareIdTargetsIdOK  %+v", 200, o.Payload)
}

func (o *GetSharesShareIDTargetsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sharetarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSharesShareIDTargetsIDNotFound creates a GetSharesShareIDTargetsIDNotFound with default headers values
func NewGetSharesShareIDTargetsIDNotFound() *GetSharesShareIDTargetsIDNotFound {
	return &GetSharesShareIDTargetsIDNotFound{}
}

/*GetSharesShareIDTargetsIDNotFound handles this case with default header values.

error
*/
type GetSharesShareIDTargetsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetSharesShareIDTargetsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /shares/{share_id}/targets/{id}][%d] getSharesShareIdTargetsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSharesShareIDTargetsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
