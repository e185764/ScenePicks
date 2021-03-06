// Code generated by go-swagger; DO NOT EDIT.

package scenepicks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
)

// PostToTwitterOKCode is the HTTP code returned for type PostToTwitterOK
const PostToTwitterOKCode int = 200

/*PostToTwitterOK posted to twitter

swagger:response postToTwitterOK
*/
type PostToTwitterOK struct {

	/*
	  In: Body
	*/
	Payload *PostToTwitterOKBody `json:"body,omitempty"`
}

// NewPostToTwitterOK creates PostToTwitterOK with default headers values
func NewPostToTwitterOK() *PostToTwitterOK {

	return &PostToTwitterOK{}
}

// WithPayload adds the payload to the post to twitter o k response
func (o *PostToTwitterOK) WithPayload(payload *PostToTwitterOKBody) *PostToTwitterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post to twitter o k response
func (o *PostToTwitterOK) SetPayload(payload *PostToTwitterOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostToTwitterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostToTwitterDefault generic error response

swagger:response postToTwitterDefault
*/
type PostToTwitterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostToTwitterDefault creates PostToTwitterDefault with default headers values
func NewPostToTwitterDefault(code int) *PostToTwitterDefault {
	if code <= 0 {
		code = 500
	}

	return &PostToTwitterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post to twitter default response
func (o *PostToTwitterDefault) WithStatusCode(code int) *PostToTwitterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post to twitter default response
func (o *PostToTwitterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post to twitter default response
func (o *PostToTwitterDefault) WithPayload(payload models.Error) *PostToTwitterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post to twitter default response
func (o *PostToTwitterDefault) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostToTwitterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
