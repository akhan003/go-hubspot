/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MembershipsAPIService MembershipsAPI service
type MembershipsAPIService service

type ApiAddAndRemoveMembershipsRequest struct {
	ctx                     context.Context
	ApiService              *MembershipsAPIService
	listId                  int32
	membershipChangeRequest *MembershipChangeRequest
}

// The IDs of the records to add and/or remove from the list.
func (r ApiAddAndRemoveMembershipsRequest) MembershipChangeRequest(membershipChangeRequest MembershipChangeRequest) ApiAddAndRemoveMembershipsRequest {
	r.membershipChangeRequest = &membershipChangeRequest
	return r
}

func (r ApiAddAndRemoveMembershipsRequest) Execute() (*MembershipsUpdateResponse, *http.Response, error) {
	return r.ApiService.AddAndRemoveMembershipsExecute(r)
}

/*
AddAndRemoveMemberships Add and/or Remove Records from a List

Add and/or remove records that have already been created in the system to and/or from a list.

This endpoint only works for lists that have a `processingType` of `MANUAL` or `SNAPSHOT`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	@return ApiAddAndRemoveMembershipsRequest
*/
func (a *MembershipsAPIService) AddAndRemoveMemberships(ctx context.Context, listId int32) ApiAddAndRemoveMembershipsRequest {
	return ApiAddAndRemoveMembershipsRequest{
		ApiService: a,
		ctx:        ctx,
		listId:     listId,
	}
}

// Execute executes the request
//
//	@return MembershipsUpdateResponse
func (a *MembershipsAPIService) AddAndRemoveMembershipsExecute(r ApiAddAndRemoveMembershipsRequest) (*MembershipsUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MembershipsUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MembershipsAPIService.AddAndRemoveMemberships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/lists/{listId}/memberships/add-and-remove"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.membershipChangeRequest == nil {
		return localVarReturnValue, nil, reportError("membershipChangeRequest is required and must be specified")
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
	localVarPostBody = r.membershipChangeRequest
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

type ApiAddMembershipsRequest struct {
	ctx         context.Context
	ApiService  *MembershipsAPIService
	listId      int32
	requestBody *[]int64
}

// The IDs of the records to add to the list.
func (r ApiAddMembershipsRequest) RequestBody(requestBody []int64) ApiAddMembershipsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiAddMembershipsRequest) Execute() (*MembershipsUpdateResponse, *http.Response, error) {
	return r.ApiService.AddMembershipsExecute(r)
}

/*
AddMemberships Add Records to a List

Add the records provided to the list. Records that do not exist or that are already members of the list are ignored.

This endpoint only works for lists that have a `processingType` of `MANUAL` or `SNAPSHOT`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	@return ApiAddMembershipsRequest
*/
func (a *MembershipsAPIService) AddMemberships(ctx context.Context, listId int32) ApiAddMembershipsRequest {
	return ApiAddMembershipsRequest{
		ApiService: a,
		ctx:        ctx,
		listId:     listId,
	}
}

// Execute executes the request
//
//	@return MembershipsUpdateResponse
func (a *MembershipsAPIService) AddMembershipsExecute(r ApiAddMembershipsRequest) (*MembershipsUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MembershipsUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MembershipsAPIService.AddMemberships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/lists/{listId}/memberships/add"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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

type ApiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest struct {
	ctx        context.Context
	ApiService *MembershipsAPIService
	listId     int32
}

func (r ApiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3ListsListIdMembershipsRemoveAllExecute(r)
}

/*
DeleteCrmV3ListsListIdMembershipsRemoveAll Delete All Records from a List

Remove **all** of the records from a list. ***Note:*** *The list is not deleted.*

This endpoint only works for lists that have a `processingType` of `MANUAL` or `SNAPSHOT`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	@return ApiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest
*/
func (a *MembershipsAPIService) DeleteCrmV3ListsListIdMembershipsRemoveAll(ctx context.Context, listId int32) ApiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest {
	return ApiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest{
		ApiService: a,
		ctx:        ctx,
		listId:     listId,
	}
}

// Execute executes the request
func (a *MembershipsAPIService) DeleteCrmV3ListsListIdMembershipsRemoveAllExecute(r ApiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MembershipsAPIService.DeleteCrmV3ListsListIdMembershipsRemoveAll")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/lists/{listId}/memberships"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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

type ApiGetCrmV3ListsListIdMembershipsGetPageRequest struct {
	ctx        context.Context
	ApiService *MembershipsAPIService
	listId     int32
	after      *string
	before     *string
	limit      *int32
}

// The paging offset token for the page that comes &#x60;after&#x60; the previously requested records.  If provided, then the records in the response will be the records following the offset, sorted in *ascending* order. Takes precedence over the &#x60;before&#x60; offset.
func (r ApiGetCrmV3ListsListIdMembershipsGetPageRequest) After(after string) ApiGetCrmV3ListsListIdMembershipsGetPageRequest {
	r.after = &after
	return r
}

// The paging offset token for the page that comes &#x60;before&#x60; the previously requested records.  If provided, then the records in the response will be the records preceding the offset, sorted in *descending* order.
func (r ApiGetCrmV3ListsListIdMembershipsGetPageRequest) Before(before string) ApiGetCrmV3ListsListIdMembershipsGetPageRequest {
	r.before = &before
	return r
}

// The number of records to return in the response. The maximum &#x60;limit&#x60; is 250.
func (r ApiGetCrmV3ListsListIdMembershipsGetPageRequest) Limit(limit int32) ApiGetCrmV3ListsListIdMembershipsGetPageRequest {
	r.limit = &limit
	return r
}

func (r ApiGetCrmV3ListsListIdMembershipsGetPageRequest) Execute() (*CollectionResponseLong, *http.Response, error) {
	return r.ApiService.GetCrmV3ListsListIdMembershipsGetPageExecute(r)
}

/*
GetCrmV3ListsListIdMembershipsGetPage Fetch List Memberships Ordered by ID

Fetch the memberships of a list in order sorted by the `recordId` of the records in the list.

The `recordId`s are sorted in *ascending* order if an `after` offset or no offset is provided. If only a `before` offset is provided, then the records are sorted in *descending* order.

The `after` offset parameter will take precedence over the `before` offset in a case where both are provided.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId The **ILS ID** of the list.
	@return ApiGetCrmV3ListsListIdMembershipsGetPageRequest
*/
func (a *MembershipsAPIService) GetCrmV3ListsListIdMembershipsGetPage(ctx context.Context, listId int32) ApiGetCrmV3ListsListIdMembershipsGetPageRequest {
	return ApiGetCrmV3ListsListIdMembershipsGetPageRequest{
		ApiService: a,
		ctx:        ctx,
		listId:     listId,
	}
}

// Execute executes the request
//
//	@return CollectionResponseLong
func (a *MembershipsAPIService) GetCrmV3ListsListIdMembershipsGetPageExecute(r ApiGetCrmV3ListsListIdMembershipsGetPageRequest) (*CollectionResponseLong, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponseLong
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MembershipsAPIService.GetCrmV3ListsListIdMembershipsGetPage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/lists/{listId}/memberships"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
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

type ApiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest struct {
	ctx          context.Context
	ApiService   *MembershipsAPIService
	listId       int32
	sourceListId int32
}

func (r ApiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListExecute(r)
}

/*
PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList Add All Records from a Source List to a Destination List

Add all of the records from a *source list* (specified by the `sourceListId`) to a *destination list* (specified by the `listId`). Records that are already members of the *destination list* will be ignored. The *destination* and *source list* IDs must be different. The *destination* and *source lists* must contain records of the same type (e.g. contacts, companies, etc.).

This endpoint only works for *destination lists* that have a `processingType` of `MANUAL` or `SNAPSHOT`. The *source list* can have any `processingType`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId The **ILS ID** of the `MANUAL` or `SNAPSHOT` *destination list*, which the *source list* records are added to.
	@param sourceListId The **ILS ID** of the *source list* to grab the records from, which are then added to the *destination list*.
	@return ApiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest
*/
func (a *MembershipsAPIService) PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList(ctx context.Context, listId int32, sourceListId int32) ApiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest {
	return ApiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest{
		ApiService:   a,
		ctx:          ctx,
		listId:       listId,
		sourceListId: sourceListId,
	}
}

// Execute executes the request
func (a *MembershipsAPIService) PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListExecute(r ApiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MembershipsAPIService.PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/lists/{listId}/memberships/add-from/{sourceListId}"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sourceListId"+"}", url.PathEscape(parameterValueToString(r.sourceListId, "sourceListId")), -1)

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

type ApiRemoveMembershipsRequest struct {
	ctx         context.Context
	ApiService  *MembershipsAPIService
	listId      int32
	requestBody *[]int64
}

// The IDs of the records to remove from the list.
func (r ApiRemoveMembershipsRequest) RequestBody(requestBody []int64) ApiRemoveMembershipsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiRemoveMembershipsRequest) Execute() (*MembershipsUpdateResponse, *http.Response, error) {
	return r.ApiService.RemoveMembershipsExecute(r)
}

/*
RemoveMemberships Remove Records from a List

Remove the records provided from the list. Records that do not exist or that are not members of the list are ignored.

This endpoint only works for lists that have a `processingType` of `MANUAL` or `SNAPSHOT`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	@return ApiRemoveMembershipsRequest
*/
func (a *MembershipsAPIService) RemoveMemberships(ctx context.Context, listId int32) ApiRemoveMembershipsRequest {
	return ApiRemoveMembershipsRequest{
		ApiService: a,
		ctx:        ctx,
		listId:     listId,
	}
}

// Execute executes the request
//
//	@return MembershipsUpdateResponse
func (a *MembershipsAPIService) RemoveMembershipsExecute(r ApiRemoveMembershipsRequest) (*MembershipsUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MembershipsUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MembershipsAPIService.RemoveMemberships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/lists/{listId}/memberships/remove"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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