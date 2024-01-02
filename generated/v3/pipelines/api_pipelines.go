/*
Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// PipelinesAPIService PipelinesAPI service
type PipelinesAPIService service

type ApiArchiveRequest struct {
	ctx                                 context.Context
	ApiService                          *PipelinesAPIService
	objectType                          string
	pipelineId                          string
	validateReferencesBeforeDelete      *bool
	validateDealStageUsagesBeforeDelete *bool
}

func (r ApiArchiveRequest) ValidateReferencesBeforeDelete(validateReferencesBeforeDelete bool) ApiArchiveRequest {
	r.validateReferencesBeforeDelete = &validateReferencesBeforeDelete
	return r
}

func (r ApiArchiveRequest) ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete bool) ApiArchiveRequest {
	r.validateDealStageUsagesBeforeDelete = &validateDealStageUsagesBeforeDelete
	return r
}

func (r ApiArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.ArchiveExecute(r)
}

/*
Archive Delete a pipeline

Delete the pipeline identified by `{pipelineId}`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@param pipelineId
	@return ApiArchiveRequest
*/
func (a *PipelinesAPIService) Archive(ctx context.Context, objectType string, pipelineId string) ApiArchiveRequest {
	return ApiArchiveRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
func (a *PipelinesAPIService) ArchiveExecute(r ApiArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesAPIService.Archive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterValueToString(r.pipelineId, "pipelineId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.validateReferencesBeforeDelete != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validateReferencesBeforeDelete", r.validateReferencesBeforeDelete, "")
	} else {
		var defaultValue bool = false
		r.validateReferencesBeforeDelete = &defaultValue
	}
	if r.validateDealStageUsagesBeforeDelete != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validateDealStageUsagesBeforeDelete", r.validateDealStageUsagesBeforeDelete, "")
	} else {
		var defaultValue bool = false
		r.validateDealStageUsagesBeforeDelete = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateRequest struct {
	ctx           context.Context
	ApiService    *PipelinesAPIService
	objectType    string
	pipelineInput *PipelineInput
}

func (r ApiCreateRequest) PipelineInput(pipelineInput PipelineInput) ApiCreateRequest {
	r.pipelineInput = &pipelineInput
	return r
}

func (r ApiCreateRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a pipeline

Create a new pipeline with the provided property values. The entire pipeline object, including its unique ID, will be returned in the response.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@return ApiCreateRequest
*/
func (a *PipelinesAPIService) Create(ctx context.Context, objectType string) ApiCreateRequest {
	return ApiCreateRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//
//	@return Pipeline
func (a *PipelinesAPIService) CreateExecute(r ApiCreateRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesAPIService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pipelineInput == nil {
		return localVarReturnValue, nil, reportError("pipelineInput is required and must be specified")
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
	localVarPostBody = r.pipelineInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
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

type ApiGetAllRequest struct {
	ctx        context.Context
	ApiService *PipelinesAPIService
	objectType string
}

func (r ApiGetAllRequest) Execute() (*CollectionResponsePipelineNoPaging, *http.Response, error) {
	return r.ApiService.GetAllExecute(r)
}

/*
GetAll Retrieve all pipelines

Return all pipelines for the object type specified by `{objectType}`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@return ApiGetAllRequest
*/
func (a *PipelinesAPIService) GetAll(ctx context.Context, objectType string) ApiGetAllRequest {
	return ApiGetAllRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//
//	@return CollectionResponsePipelineNoPaging
func (a *PipelinesAPIService) GetAllExecute(r ApiGetAllRequest) (*CollectionResponsePipelineNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePipelineNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesAPIService.GetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)

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
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
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

type ApiGetByIDRequest struct {
	ctx        context.Context
	ApiService *PipelinesAPIService
	objectType string
	pipelineId string
}

func (r ApiGetByIDRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.GetByIDExecute(r)
}

/*
GetByID Return a pipeline by ID

Return a single pipeline object identified by its unique `{pipelineId}`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@param pipelineId
	@return ApiGetByIDRequest
*/
func (a *PipelinesAPIService) GetByID(ctx context.Context, objectType string, pipelineId string) ApiGetByIDRequest {
	return ApiGetByIDRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
//
//	@return Pipeline
func (a *PipelinesAPIService) GetByIDExecute(r ApiGetByIDRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesAPIService.GetByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterValueToString(r.pipelineId, "pipelineId")), -1)

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
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
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

type ApiReplaceRequest struct {
	ctx                                 context.Context
	ApiService                          *PipelinesAPIService
	objectType                          string
	pipelineId                          string
	pipelineInput                       *PipelineInput
	validateReferencesBeforeDelete      *bool
	validateDealStageUsagesBeforeDelete *bool
}

func (r ApiReplaceRequest) PipelineInput(pipelineInput PipelineInput) ApiReplaceRequest {
	r.pipelineInput = &pipelineInput
	return r
}

func (r ApiReplaceRequest) ValidateReferencesBeforeDelete(validateReferencesBeforeDelete bool) ApiReplaceRequest {
	r.validateReferencesBeforeDelete = &validateReferencesBeforeDelete
	return r
}

func (r ApiReplaceRequest) ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete bool) ApiReplaceRequest {
	r.validateDealStageUsagesBeforeDelete = &validateDealStageUsagesBeforeDelete
	return r
}

func (r ApiReplaceRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.ReplaceExecute(r)
}

/*
Replace Replace a pipeline

Replace all the properties of an existing pipeline with the values provided. This will overwrite any existing pipeline stages. The updated pipeline will be returned in the response.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@param pipelineId
	@return ApiReplaceRequest
*/
func (a *PipelinesAPIService) Replace(ctx context.Context, objectType string, pipelineId string) ApiReplaceRequest {
	return ApiReplaceRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
//
//	@return Pipeline
func (a *PipelinesAPIService) ReplaceExecute(r ApiReplaceRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesAPIService.Replace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterValueToString(r.pipelineId, "pipelineId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pipelineInput == nil {
		return localVarReturnValue, nil, reportError("pipelineInput is required and must be specified")
	}

	if r.validateReferencesBeforeDelete != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validateReferencesBeforeDelete", r.validateReferencesBeforeDelete, "")
	} else {
		var defaultValue bool = false
		r.validateReferencesBeforeDelete = &defaultValue
	}
	if r.validateDealStageUsagesBeforeDelete != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validateDealStageUsagesBeforeDelete", r.validateDealStageUsagesBeforeDelete, "")
	} else {
		var defaultValue bool = false
		r.validateDealStageUsagesBeforeDelete = &defaultValue
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
	localVarPostBody = r.pipelineInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
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

type ApiUpdateRequest struct {
	ctx                                 context.Context
	ApiService                          *PipelinesAPIService
	objectType                          string
	pipelineId                          string
	pipelinePatchInput                  *PipelinePatchInput
	validateReferencesBeforeDelete      *bool
	validateDealStageUsagesBeforeDelete *bool
}

func (r ApiUpdateRequest) PipelinePatchInput(pipelinePatchInput PipelinePatchInput) ApiUpdateRequest {
	r.pipelinePatchInput = &pipelinePatchInput
	return r
}

func (r ApiUpdateRequest) ValidateReferencesBeforeDelete(validateReferencesBeforeDelete bool) ApiUpdateRequest {
	r.validateReferencesBeforeDelete = &validateReferencesBeforeDelete
	return r
}

func (r ApiUpdateRequest) ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete bool) ApiUpdateRequest {
	r.validateDealStageUsagesBeforeDelete = &validateDealStageUsagesBeforeDelete
	return r
}

func (r ApiUpdateRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update a pipeline

Perform a partial update of the pipeline identified by `{pipelineId}`. The updated pipeline will be returned in the response.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param objectType
	@param pipelineId
	@return ApiUpdateRequest
*/
func (a *PipelinesAPIService) Update(ctx context.Context, objectType string, pipelineId string) ApiUpdateRequest {
	return ApiUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
//
//	@return Pipeline
func (a *PipelinesAPIService) UpdateExecute(r ApiUpdateRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesAPIService.Update")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterValueToString(r.objectType, "objectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterValueToString(r.pipelineId, "pipelineId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pipelinePatchInput == nil {
		return localVarReturnValue, nil, reportError("pipelinePatchInput is required and must be specified")
	}

	if r.validateReferencesBeforeDelete != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validateReferencesBeforeDelete", r.validateReferencesBeforeDelete, "")
	} else {
		var defaultValue bool = false
		r.validateReferencesBeforeDelete = &defaultValue
	}
	if r.validateDealStageUsagesBeforeDelete != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "validateDealStageUsagesBeforeDelete", r.validateDealStageUsagesBeforeDelete, "")
	} else {
		var defaultValue bool = false
		r.validateDealStageUsagesBeforeDelete = &defaultValue
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
	localVarPostBody = r.pipelinePatchInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app"] = key
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
