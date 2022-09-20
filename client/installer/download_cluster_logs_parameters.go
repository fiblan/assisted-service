// Code generated by go-swagger; DO NOT EDIT.

package installer

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
)

// NewDownloadClusterLogsParams creates a new DownloadClusterLogsParams object
// with the default values initialized.
func NewDownloadClusterLogsParams() *DownloadClusterLogsParams {
	var ()
	return &DownloadClusterLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadClusterLogsParamsWithTimeout creates a new DownloadClusterLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDownloadClusterLogsParamsWithTimeout(timeout time.Duration) *DownloadClusterLogsParams {
	var ()
	return &DownloadClusterLogsParams{

		timeout: timeout,
	}
}

// NewDownloadClusterLogsParamsWithContext creates a new DownloadClusterLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDownloadClusterLogsParamsWithContext(ctx context.Context) *DownloadClusterLogsParams {
	var ()
	return &DownloadClusterLogsParams{

		Context: ctx,
	}
}

// NewDownloadClusterLogsParamsWithHTTPClient creates a new DownloadClusterLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDownloadClusterLogsParamsWithHTTPClient(client *http.Client) *DownloadClusterLogsParams {
	var ()
	return &DownloadClusterLogsParams{
		HTTPClient: client,
	}
}

/*DownloadClusterLogsParams contains all the parameters to send to the API endpoint
for the download cluster logs operation typically these are written to a http.Request
*/
type DownloadClusterLogsParams struct {

	/*ClusterID
	  The cluster whose logs should be downloaded.

	*/
	ClusterID strfmt.UUID
	/*HostID
	  A specific host in the cluster whose logs should be downloaded.

	*/
	HostID *strfmt.UUID
	/*LogsType
	  The type of logs to be downloaded.

	*/
	LogsType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the download cluster logs params
func (o *DownloadClusterLogsParams) WithTimeout(timeout time.Duration) *DownloadClusterLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download cluster logs params
func (o *DownloadClusterLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download cluster logs params
func (o *DownloadClusterLogsParams) WithContext(ctx context.Context) *DownloadClusterLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download cluster logs params
func (o *DownloadClusterLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download cluster logs params
func (o *DownloadClusterLogsParams) WithHTTPClient(client *http.Client) *DownloadClusterLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download cluster logs params
func (o *DownloadClusterLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the download cluster logs params
func (o *DownloadClusterLogsParams) WithClusterID(clusterID strfmt.UUID) *DownloadClusterLogsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the download cluster logs params
func (o *DownloadClusterLogsParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the download cluster logs params
func (o *DownloadClusterLogsParams) WithHostID(hostID *strfmt.UUID) *DownloadClusterLogsParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the download cluster logs params
func (o *DownloadClusterLogsParams) SetHostID(hostID *strfmt.UUID) {
	o.HostID = hostID
}

// WithLogsType adds the logsType to the download cluster logs params
func (o *DownloadClusterLogsParams) WithLogsType(logsType *string) *DownloadClusterLogsParams {
	o.SetLogsType(logsType)
	return o
}

// SetLogsType adds the logsType to the download cluster logs params
func (o *DownloadClusterLogsParams) SetLogsType(logsType *string) {
	o.LogsType = logsType
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadClusterLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.HostID != nil {

		// query param host_id
		var qrHostID strfmt.UUID
		if o.HostID != nil {
			qrHostID = *o.HostID
		}
		qHostID := qrHostID.String()
		if qHostID != "" {
			if err := r.SetQueryParam("host_id", qHostID); err != nil {
				return err
			}
		}

	}

	if o.LogsType != nil {

		// query param logs_type
		var qrLogsType string
		if o.LogsType != nil {
			qrLogsType = *o.LogsType
		}
		qLogsType := qrLogsType
		if qLogsType != "" {
			if err := r.SetQueryParam("logs_type", qLogsType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}