/*
Calling Extensions

Provides a way for apps to add custom calling options to a contact record. This works in conjunction with the [Calling SDK](#), which is used to build your phone/calling UI. The endpoints here allow your service to appear as an option to HubSpot users when they access the *Call* action on a contact record. Once accessed, your custom phone/calling UI will be displayed in an iframe at the specified URL with the specified dimensions on that record.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calling

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RecordingSettingsAPIService RecordingSettingsAPI service
type RecordingSettingsAPIService service

type ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingRequest struct {
	ctx        context.Context
	ApiService *RecordingSettingsAPIService
	appId      int32
}

func (r ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) Execute() (*RecordingSettingsResponse, *http.Response, error) {
	return r.ApiService.GetCrmV3ExtensionsCallingAppIdSettingsRecordingExecute(r)
}

/*
GetCrmV3ExtensionsCallingAppIdSettingsRecording Method for GetCrmV3ExtensionsCallingAppIdSettingsRecording

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingRequest
*/
func (a *RecordingSettingsAPIService) GetCrmV3ExtensionsCallingAppIdSettingsRecording(ctx context.Context, appId int32) ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingRequest {
	return ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return RecordingSettingsResponse
func (a *RecordingSettingsAPIService) GetCrmV3ExtensionsCallingAppIdSettingsRecordingExecute(r ApiGetCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) (*RecordingSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RecordingSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingSettingsAPIService.GetCrmV3ExtensionsCallingAppIdSettingsRecording")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings/recording"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest struct {
	ctx                           context.Context
	ApiService                    *RecordingSettingsAPIService
	appId                         int32
	recordingSettingsPatchRequest *RecordingSettingsPatchRequest
}

func (r ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) RecordingSettingsPatchRequest(recordingSettingsPatchRequest RecordingSettingsPatchRequest) ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest {
	r.recordingSettingsPatchRequest = &recordingSettingsPatchRequest
	return r
}

func (r ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) Execute() (*RecordingSettingsResponse, *http.Response, error) {
	return r.ApiService.PatchCrmV3ExtensionsCallingAppIdSettingsRecordingExecute(r)
}

/*
PatchCrmV3ExtensionsCallingAppIdSettingsRecording Method for PatchCrmV3ExtensionsCallingAppIdSettingsRecording

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest
*/
func (a *RecordingSettingsAPIService) PatchCrmV3ExtensionsCallingAppIdSettingsRecording(ctx context.Context, appId int32) ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest {
	return ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return RecordingSettingsResponse
func (a *RecordingSettingsAPIService) PatchCrmV3ExtensionsCallingAppIdSettingsRecordingExecute(r ApiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) (*RecordingSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RecordingSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingSettingsAPIService.PatchCrmV3ExtensionsCallingAppIdSettingsRecording")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings/recording"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordingSettingsPatchRequest == nil {
		return localVarReturnValue, nil, reportError("recordingSettingsPatchRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.recordingSettingsPatchRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest struct {
	ctx                      context.Context
	ApiService               *RecordingSettingsAPIService
	appId                    int32
	recordingSettingsRequest *RecordingSettingsRequest
}

func (r ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) RecordingSettingsRequest(recordingSettingsRequest RecordingSettingsRequest) ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest {
	r.recordingSettingsRequest = &recordingSettingsRequest
	return r
}

func (r ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) Execute() (*RecordingSettingsResponse, *http.Response, error) {
	return r.ApiService.PostCrmV3ExtensionsCallingAppIdSettingsRecordingExecute(r)
}

/*
PostCrmV3ExtensionsCallingAppIdSettingsRecording Method for PostCrmV3ExtensionsCallingAppIdSettingsRecording

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest
*/
func (a *RecordingSettingsAPIService) PostCrmV3ExtensionsCallingAppIdSettingsRecording(ctx context.Context, appId int32) ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest {
	return ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return RecordingSettingsResponse
func (a *RecordingSettingsAPIService) PostCrmV3ExtensionsCallingAppIdSettingsRecordingExecute(r ApiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest) (*RecordingSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RecordingSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingSettingsAPIService.PostCrmV3ExtensionsCallingAppIdSettingsRecording")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/extensions/calling/{appId}/settings/recording"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recordingSettingsRequest == nil {
		return localVarReturnValue, nil, reportError("recordingSettingsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.recordingSettingsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["developer_hapikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("hapikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}