// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// NewRejectLogoutRequestParams creates a new RejectLogoutRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRejectLogoutRequestParams() *RejectLogoutRequestParams {
	return &RejectLogoutRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRejectLogoutRequestParamsWithTimeout creates a new RejectLogoutRequestParams object
// with the ability to set a timeout on a request.
func NewRejectLogoutRequestParamsWithTimeout(timeout time.Duration) *RejectLogoutRequestParams {
	return &RejectLogoutRequestParams{
		timeout: timeout,
	}
}

// NewRejectLogoutRequestParamsWithContext creates a new RejectLogoutRequestParams object
// with the ability to set a context for a request.
func NewRejectLogoutRequestParamsWithContext(ctx context.Context) *RejectLogoutRequestParams {
	return &RejectLogoutRequestParams{
		Context: ctx,
	}
}

// NewRejectLogoutRequestParamsWithHTTPClient creates a new RejectLogoutRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRejectLogoutRequestParamsWithHTTPClient(client *http.Client) *RejectLogoutRequestParams {
	return &RejectLogoutRequestParams{
		HTTPClient: client,
	}
}

/* RejectLogoutRequestParams contains all the parameters to send to the API endpoint
   for the reject logout request operation.

   Typically these are written to a http.Request.
*/
type RejectLogoutRequestParams struct {

	// Body.
	Body *models.RejectRequest

	// LogoutChallenge.
	LogoutChallenge string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reject logout request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RejectLogoutRequestParams) WithDefaults() *RejectLogoutRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reject logout request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RejectLogoutRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reject logout request params
func (o *RejectLogoutRequestParams) WithTimeout(timeout time.Duration) *RejectLogoutRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reject logout request params
func (o *RejectLogoutRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reject logout request params
func (o *RejectLogoutRequestParams) WithContext(ctx context.Context) *RejectLogoutRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reject logout request params
func (o *RejectLogoutRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reject logout request params
func (o *RejectLogoutRequestParams) WithHTTPClient(client *http.Client) *RejectLogoutRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reject logout request params
func (o *RejectLogoutRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reject logout request params
func (o *RejectLogoutRequestParams) WithBody(body *models.RejectRequest) *RejectLogoutRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reject logout request params
func (o *RejectLogoutRequestParams) SetBody(body *models.RejectRequest) {
	o.Body = body
}

// WithLogoutChallenge adds the logoutChallenge to the reject logout request params
func (o *RejectLogoutRequestParams) WithLogoutChallenge(logoutChallenge string) *RejectLogoutRequestParams {
	o.SetLogoutChallenge(logoutChallenge)
	return o
}

// SetLogoutChallenge adds the logoutChallenge to the reject logout request params
func (o *RejectLogoutRequestParams) SetLogoutChallenge(logoutChallenge string) {
	o.LogoutChallenge = logoutChallenge
}

// WriteToRequest writes these params to a swagger request
func (o *RejectLogoutRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param logout_challenge
	qrLogoutChallenge := o.LogoutChallenge
	qLogoutChallenge := qrLogoutChallenge
	if qLogoutChallenge != "" {

		if err := r.SetQueryParam("logout_challenge", qLogoutChallenge); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
