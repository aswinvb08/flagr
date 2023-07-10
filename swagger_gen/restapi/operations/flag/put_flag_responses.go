// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openflagr/flagr/swagger_gen/models"
)

// PutFlagOKCode is the HTTP code returned for type PutFlagOK
const PutFlagOKCode int = 200

/*
PutFlagOK returns the flag

swagger:response putFlagOK
*/
type PutFlagOK struct {

	/*
	  In: Body
	*/
	Payload *models.Flag `json:"body,omitempty"`
}

// NewPutFlagOK creates PutFlagOK with default headers values
func NewPutFlagOK() *PutFlagOK {

	return &PutFlagOK{}
}

// WithPayload adds the payload to the put flag o k response
func (o *PutFlagOK) WithPayload(payload *models.Flag) *PutFlagOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put flag o k response
func (o *PutFlagOK) SetPayload(payload *models.Flag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutFlagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PutFlagDefault generic error response

swagger:response putFlagDefault
*/
type PutFlagDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutFlagDefault creates PutFlagDefault with default headers values
func NewPutFlagDefault(code int) *PutFlagDefault {
	if code <= 0 {
		code = 500
	}

	return &PutFlagDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put flag default response
func (o *PutFlagDefault) WithStatusCode(code int) *PutFlagDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put flag default response
func (o *PutFlagDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put flag default response
func (o *PutFlagDefault) WithPayload(payload *models.Error) *PutFlagDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put flag default response
func (o *PutFlagDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutFlagDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
