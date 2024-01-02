/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefinitionsAPIService DefinitionsAPI service
type DefinitionsAPIService service

type ApiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveRequest struct {
	ctx               context.Context
	ApiService        *DefinitionsAPIService
	fromObjectType    string
	toObjectType      string
	associationTypeId int32
}

func (r ApiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveExecute(r)
}

/*
DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive Delete

Deletes an association definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@param associationTypeId
	@return ApiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveRequest
*/
func (a *DefinitionsAPIService) DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive(ctx context.Context, fromObjectType string, toObjectType string, associationTypeId int32) ApiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveRequest {
	return ApiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveRequest{
		ApiService:        a,
		ctx:               ctx,
		fromObjectType:    fromObjectType,
		toObjectType:      toObjectType,
		associationTypeId: associationTypeId,
	}
}

// Execute executes the request
func (a *DefinitionsAPIService) DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveExecute(r ApiDeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsAPIService.DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/labels/{associationTypeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterValueToString(r.fromObjectType, "fromObjectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterValueToString(r.toObjectType, "toObjectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"associationTypeId"+"}", url.PathEscape(parameterValueToString(r.associationTypeId, "associationTypeId")), -1)

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

type ApiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllRequest struct {
	ctx            context.Context
	ApiService     *DefinitionsAPIService
	fromObjectType string
	toObjectType   string
}

func (r ApiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllRequest) Execute() (*CollectionResponseAssociationSpecWithLabelNoPaging, *http.Response, error) {
	return r.ApiService.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllExecute(r)
}

/*
GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll Read

Returns all association types between two object types

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllRequest
*/
func (a *DefinitionsAPIService) GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll(ctx context.Context, fromObjectType string, toObjectType string) ApiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllRequest {
	return ApiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//
//	@return CollectionResponseAssociationSpecWithLabelNoPaging
func (a *DefinitionsAPIService) GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllExecute(r ApiGetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAllRequest) (*CollectionResponseAssociationSpecWithLabelNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseAssociationSpecWithLabelNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsAPIService.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/labels"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterValueToString(r.fromObjectType, "fromObjectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterValueToString(r.toObjectType, "toObjectType")), -1)

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

type ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest struct {
	ctx                                      context.Context
	ApiService                               *DefinitionsAPIService
	fromObjectType                           string
	toObjectType                             string
	publicAssociationDefinitionCreateRequest *PublicAssociationDefinitionCreateRequest
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest) PublicAssociationDefinitionCreateRequest(publicAssociationDefinitionCreateRequest PublicAssociationDefinitionCreateRequest) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest {
	r.publicAssociationDefinitionCreateRequest = &publicAssociationDefinitionCreateRequest
	return r
}

func (r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest) Execute() (*CollectionResponseAssociationSpecWithLabelNoPaging, *http.Response, error) {
	return r.ApiService.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateExecute(r)
}

/*
PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate Create

Create a user defined association definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest
*/
func (a *DefinitionsAPIService) PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate(ctx context.Context, fromObjectType string, toObjectType string) ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest {
	return ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
//
//	@return CollectionResponseAssociationSpecWithLabelNoPaging
func (a *DefinitionsAPIService) PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateExecute(r ApiPostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreateRequest) (*CollectionResponseAssociationSpecWithLabelNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseAssociationSpecWithLabelNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsAPIService.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/labels"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterValueToString(r.fromObjectType, "fromObjectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterValueToString(r.toObjectType, "toObjectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicAssociationDefinitionCreateRequest == nil {
		return localVarReturnValue, nil, reportError("publicAssociationDefinitionCreateRequest is required and must be specified")
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
	localVarPostBody = r.publicAssociationDefinitionCreateRequest
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

type ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest struct {
	ctx                                      context.Context
	ApiService                               *DefinitionsAPIService
	fromObjectType                           string
	toObjectType                             string
	publicAssociationDefinitionUpdateRequest *PublicAssociationDefinitionUpdateRequest
}

func (r ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest) PublicAssociationDefinitionUpdateRequest(publicAssociationDefinitionUpdateRequest PublicAssociationDefinitionUpdateRequest) ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest {
	r.publicAssociationDefinitionUpdateRequest = &publicAssociationDefinitionUpdateRequest
	return r
}

func (r ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateExecute(r)
}

/*
PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate Update

Update a user defined association definition

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fromObjectType
	@param toObjectType
	@return ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest
*/
func (a *DefinitionsAPIService) PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate(ctx context.Context, fromObjectType string, toObjectType string) ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest {
	return ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest{
		ApiService:     a,
		ctx:            ctx,
		fromObjectType: fromObjectType,
		toObjectType:   toObjectType,
	}
}

// Execute executes the request
func (a *DefinitionsAPIService) PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateExecute(r ApiPutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefinitionsAPIService.PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v4/associations/{fromObjectType}/{toObjectType}/labels"
	localVarPath = strings.Replace(localVarPath, "{"+"fromObjectType"+"}", url.PathEscape(parameterValueToString(r.fromObjectType, "fromObjectType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toObjectType"+"}", url.PathEscape(parameterValueToString(r.toObjectType, "toObjectType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publicAssociationDefinitionUpdateRequest == nil {
		return nil, reportError("publicAssociationDefinitionUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.publicAssociationDefinitionUpdateRequest
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
