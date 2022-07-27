// Code generated by go-swagger; DO NOT EDIT.

package board

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListMessageOKCode is the HTTP code returned for type ListMessageOK
const ListMessageOKCode int = 200

/*ListMessageOK messages successfully listed.

swagger:response listMessageOK
*/
type ListMessageOK struct {

	/*
	  In: Body
	*/
	Payload []*ListMessageOKBodyItems0 `json:"body,omitempty"`
}

// NewListMessageOK creates ListMessageOK with default headers values
func NewListMessageOK() *ListMessageOK {

	return &ListMessageOK{}
}

// WithPayload adds the payload to the list message o k response
func (o *ListMessageOK) WithPayload(payload []*ListMessageOKBodyItems0) *ListMessageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list message o k response
func (o *ListMessageOK) SetPayload(payload []*ListMessageOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMessageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ListMessageOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListMessageInternalServerErrorCode is the HTTP code returned for type ListMessageInternalServerError
const ListMessageInternalServerErrorCode int = 500

/*ListMessageInternalServerError server error.

swagger:response listMessageInternalServerError
*/
type ListMessageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewListMessageInternalServerError creates ListMessageInternalServerError with default headers values
func NewListMessageInternalServerError() *ListMessageInternalServerError {

	return &ListMessageInternalServerError{}
}

// WithPayload adds the payload to the list message internal server error response
func (o *ListMessageInternalServerError) WithPayload(payload string) *ListMessageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list message internal server error response
func (o *ListMessageInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListMessageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
