/*
Actions V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// PublicActionDefinitionsAPIService PublicActionDefinitionsAPI service
type PublicActionDefinitionsAPIService service

type ApiDeleteAutomationV4ActionsAppIdDefinitionIdRequest struct {
	ctx          context.Context
	ApiService   *PublicActionDefinitionsAPIService
	definitionId string
	appId        int32
}

func (r ApiDeleteAutomationV4ActionsAppIdDefinitionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAutomationV4ActionsAppIdDefinitionIdExecute(r)
}

/*
DeleteAutomationV4ActionsAppIdDefinitionId Method for DeleteAutomationV4ActionsAppIdDefinitionId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param appId
	@return ApiDeleteAutomationV4ActionsAppIdDefinitionIdRequest
*/
func (a *PublicActionDefinitionsAPIService) DeleteAutomationV4ActionsAppIdDefinitionId(ctx context.Context, definitionId string, appId int32) ApiDeleteAutomationV4ActionsAppIdDefinitionIdRequest {
	return ApiDeleteAutomationV4ActionsAppIdDefinitionIdRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
func (a *PublicActionDefinitionsAPIService) DeleteAutomationV4ActionsAppIdDefinitionIdExecute(r ApiDeleteAutomationV4ActionsAppIdDefinitionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicActionDefinitionsAPIService.DeleteAutomationV4ActionsAppIdDefinitionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterValueToString(r.definitionId, "definitionId")), -1)
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
	localVarHTTPHeaderAccepts := []string{"*/*"}

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

type ApiGetAutomationV4ActionsAppIdRequest struct {
	ctx        context.Context
	ApiService *PublicActionDefinitionsAPIService
	appId      int32
	limit      *int32
	after      *string
	archived   *bool
}

// The maximum number of results to display per page.
func (r ApiGetAutomationV4ActionsAppIdRequest) Limit(limit int32) ApiGetAutomationV4ActionsAppIdRequest {
	r.limit = &limit
	return r
}

// The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results.
func (r ApiGetAutomationV4ActionsAppIdRequest) After(after string) ApiGetAutomationV4ActionsAppIdRequest {
	r.after = &after
	return r
}

// Whether to return only results that have been archived.
func (r ApiGetAutomationV4ActionsAppIdRequest) Archived(archived bool) ApiGetAutomationV4ActionsAppIdRequest {
	r.archived = &archived
	return r
}

func (r ApiGetAutomationV4ActionsAppIdRequest) Execute() (*CollectionResponsePublicActionDefinitionForwardPaging, *http.Response, error) {
	return r.ApiService.GetAutomationV4ActionsAppIdExecute(r)
}

/*
GetAutomationV4ActionsAppId Method for GetAutomationV4ActionsAppId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiGetAutomationV4ActionsAppIdRequest
*/
func (a *PublicActionDefinitionsAPIService) GetAutomationV4ActionsAppId(ctx context.Context, appId int32) ApiGetAutomationV4ActionsAppIdRequest {
	return ApiGetAutomationV4ActionsAppIdRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return CollectionResponsePublicActionDefinitionForwardPaging
func (a *PublicActionDefinitionsAPIService) GetAutomationV4ActionsAppIdExecute(r ApiGetAutomationV4ActionsAppIdRequest) (*CollectionResponsePublicActionDefinitionForwardPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePublicActionDefinitionForwardPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicActionDefinitionsAPIService.GetAutomationV4ActionsAppId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	}
	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
	}
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

type ApiGetAutomationV4ActionsAppIdDefinitionIdRequest struct {
	ctx          context.Context
	ApiService   *PublicActionDefinitionsAPIService
	definitionId string
	appId        int32
	archived     *bool
}

// Whether to return only results that have been archived.
func (r ApiGetAutomationV4ActionsAppIdDefinitionIdRequest) Archived(archived bool) ApiGetAutomationV4ActionsAppIdDefinitionIdRequest {
	r.archived = &archived
	return r
}

func (r ApiGetAutomationV4ActionsAppIdDefinitionIdRequest) Execute() (*PublicActionDefinition, *http.Response, error) {
	return r.ApiService.GetAutomationV4ActionsAppIdDefinitionIdExecute(r)
}

/*
GetAutomationV4ActionsAppIdDefinitionId Method for GetAutomationV4ActionsAppIdDefinitionId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param appId
	@return ApiGetAutomationV4ActionsAppIdDefinitionIdRequest
*/
func (a *PublicActionDefinitionsAPIService) GetAutomationV4ActionsAppIdDefinitionId(ctx context.Context, definitionId string, appId int32) ApiGetAutomationV4ActionsAppIdDefinitionIdRequest {
	return ApiGetAutomationV4ActionsAppIdDefinitionIdRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return PublicActionDefinition
func (a *PublicActionDefinitionsAPIService) GetAutomationV4ActionsAppIdDefinitionIdExecute(r ApiGetAutomationV4ActionsAppIdDefinitionIdRequest) (*PublicActionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicActionDefinitionsAPIService.GetAutomationV4ActionsAppIdDefinitionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterValueToString(r.definitionId, "definitionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
	}
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

type ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest struct {
	ctx                         context.Context
	ApiService                  *PublicActionDefinitionsAPIService
	definitionId                string
	appId                       int32
	publicActionDefinitionPatch *PublicActionDefinitionPatch
}

func (r ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest) PublicActionDefinitionPatch(publicActionDefinitionPatch PublicActionDefinitionPatch) ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest {
	r.publicActionDefinitionPatch = &publicActionDefinitionPatch
	return r
}

func (r ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest) Execute() (*PublicActionDefinition, *http.Response, error) {
	return r.ApiService.PatchAutomationV4ActionsAppIdDefinitionIdExecute(r)
}

/*
PatchAutomationV4ActionsAppIdDefinitionId Method for PatchAutomationV4ActionsAppIdDefinitionId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param definitionId
	@param appId
	@return ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest
*/
func (a *PublicActionDefinitionsAPIService) PatchAutomationV4ActionsAppIdDefinitionId(ctx context.Context, definitionId string, appId int32) ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest {
	return ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest{
		ApiService:   a,
		ctx:          ctx,
		definitionId: definitionId,
		appId:        appId,
	}
}

// Execute executes the request
//
//	@return PublicActionDefinition
func (a *PublicActionDefinitionsAPIService) PatchAutomationV4ActionsAppIdDefinitionIdExecute(r ApiPatchAutomationV4ActionsAppIdDefinitionIdRequest) (*PublicActionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicActionDefinitionsAPIService.PatchAutomationV4ActionsAppIdDefinitionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}/{definitionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"definitionId"+"}", url.PathEscape(parameterValueToString(r.definitionId, "definitionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicActionDefinitionPatch == nil {
		return localVarReturnValue, nil, reportError("publicActionDefinitionPatch is required and must be specified")
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
	localVarPostBody = r.publicActionDefinitionPatch
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

type ApiPostAutomationV4ActionsAppIdRequest struct {
	ctx                       context.Context
	ApiService                *PublicActionDefinitionsAPIService
	appId                     int32
	publicActionDefinitionEgg *PublicActionDefinitionEgg
}

func (r ApiPostAutomationV4ActionsAppIdRequest) PublicActionDefinitionEgg(publicActionDefinitionEgg PublicActionDefinitionEgg) ApiPostAutomationV4ActionsAppIdRequest {
	r.publicActionDefinitionEgg = &publicActionDefinitionEgg
	return r
}

func (r ApiPostAutomationV4ActionsAppIdRequest) Execute() (*PublicActionDefinition, *http.Response, error) {
	return r.ApiService.PostAutomationV4ActionsAppIdExecute(r)
}

/*
PostAutomationV4ActionsAppId Method for PostAutomationV4ActionsAppId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param appId
	@return ApiPostAutomationV4ActionsAppIdRequest
*/
func (a *PublicActionDefinitionsAPIService) PostAutomationV4ActionsAppId(ctx context.Context, appId int32) ApiPostAutomationV4ActionsAppIdRequest {
	return ApiPostAutomationV4ActionsAppIdRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
	}
}

// Execute executes the request
//
//	@return PublicActionDefinition
func (a *PublicActionDefinitionsAPIService) PostAutomationV4ActionsAppIdExecute(r ApiPostAutomationV4ActionsAppIdRequest) (*PublicActionDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublicActionDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicActionDefinitionsAPIService.PostAutomationV4ActionsAppId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/automation/v4/actions/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicActionDefinitionEgg == nil {
		return localVarReturnValue, nil, reportError("publicActionDefinitionEgg is required and must be specified")
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
	localVarPostBody = r.publicActionDefinitionEgg
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