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

// NewGetPresignedForClusterFilesParams creates a new GetPresignedForClusterFilesParams object
// with the default values initialized.
func NewGetPresignedForClusterFilesParams() *GetPresignedForClusterFilesParams {
	var ()
	return &GetPresignedForClusterFilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPresignedForClusterFilesParamsWithTimeout creates a new GetPresignedForClusterFilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPresignedForClusterFilesParamsWithTimeout(timeout time.Duration) *GetPresignedForClusterFilesParams {
	var ()
	return &GetPresignedForClusterFilesParams{

		timeout: timeout,
	}
}

// NewGetPresignedForClusterFilesParamsWithContext creates a new GetPresignedForClusterFilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPresignedForClusterFilesParamsWithContext(ctx context.Context) *GetPresignedForClusterFilesParams {
	var ()
	return &GetPresignedForClusterFilesParams{

		Context: ctx,
	}
}

// NewGetPresignedForClusterFilesParamsWithHTTPClient creates a new GetPresignedForClusterFilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPresignedForClusterFilesParamsWithHTTPClient(client *http.Client) *GetPresignedForClusterFilesParams {
	var ()
	return &GetPresignedForClusterFilesParams{
		HTTPClient: client,
	}
}

/*GetPresignedForClusterFilesParams contains all the parameters to send to the API endpoint
for the get presigned for cluster files operation typically these are written to a http.Request
*/
type GetPresignedForClusterFilesParams struct {

	/*AdditionalName
	  If downloading a manifest, the file name, prefaced with folder name, for example, openshift/99-openshift-xyz.yaml.

	*/
	AdditionalName *string
	/*ClusterID
	  The cluster that owns the file that should be downloaded.

	*/
	ClusterID strfmt.UUID
	/*FileName
	  The file to be downloaded.

	*/
	FileName string
	/*HostID
	  If downloading a file related to a host, the relevant host.

	*/
	HostID *strfmt.UUID
	/*LogsType
	  If downloading logs, the type of logs to download.

	*/
	LogsType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithTimeout(timeout time.Duration) *GetPresignedForClusterFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithContext(ctx context.Context) *GetPresignedForClusterFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithHTTPClient(client *http.Client) *GetPresignedForClusterFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdditionalName adds the additionalName to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithAdditionalName(additionalName *string) *GetPresignedForClusterFilesParams {
	o.SetAdditionalName(additionalName)
	return o
}

// SetAdditionalName adds the additionalName to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetAdditionalName(additionalName *string) {
	o.AdditionalName = additionalName
}

// WithClusterID adds the clusterID to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithClusterID(clusterID strfmt.UUID) *GetPresignedForClusterFilesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithFileName adds the fileName to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithFileName(fileName string) *GetPresignedForClusterFilesParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithHostID adds the hostID to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithHostID(hostID *strfmt.UUID) *GetPresignedForClusterFilesParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetHostID(hostID *strfmt.UUID) {
	o.HostID = hostID
}

// WithLogsType adds the logsType to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) WithLogsType(logsType *string) *GetPresignedForClusterFilesParams {
	o.SetLogsType(logsType)
	return o
}

// SetLogsType adds the logsType to the get presigned for cluster files params
func (o *GetPresignedForClusterFilesParams) SetLogsType(logsType *string) {
	o.LogsType = logsType
}

// WriteToRequest writes these params to a swagger request
func (o *GetPresignedForClusterFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdditionalName != nil {

		// query param additional_name
		var qrAdditionalName string
		if o.AdditionalName != nil {
			qrAdditionalName = *o.AdditionalName
		}
		qAdditionalName := qrAdditionalName
		if qAdditionalName != "" {
			if err := r.SetQueryParam("additional_name", qAdditionalName); err != nil {
				return err
			}
		}

	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// query param file_name
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {
		if err := r.SetQueryParam("file_name", qFileName); err != nil {
			return err
		}
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
