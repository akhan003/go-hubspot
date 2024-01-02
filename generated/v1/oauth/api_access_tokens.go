/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oauth

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AccessTokensAPIService AccessTokensAPI service
type AccessTokensAPIService service

type ApiGetOauthV1AccessTokensTokenGetRequest struct {
	ctx        context.Context
	ApiService *AccessTokensAPIService
	token      string
}

func (r ApiGetOauthV1AccessTokensTokenGetRequest) Execute() (*AccessTokenInfoResponse, *http.Response, error) {
	return r.ApiService.GetOauthV1AccessTokensTokenGetExecute(r)
}

/*
GetOauthV1AccessTokensTokenGet Method for GetOauthV1AccessTokensTokenGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param token
	@return ApiGetOauthV1AccessTokensTokenGetRequest
*/
func (a *AccessTokensAPIService) GetOauthV1AccessTokensTokenGet(ctx context.Context, token string) ApiGetOauthV1AccessTokensTokenGetRequest {
	return ApiGetOauthV1AccessTokensTokenGetRequest{
		ApiService: a,
		ctx:        ctx,
		token:      token,
	}
}

// Execute executes the request
//
//	@return AccessTokenInfoResponse
func (a *AccessTokensAPIService) GetOauthV1AccessTokensTokenGetExecute(r ApiGetOauthV1AccessTokensTokenGetRequest) (*AccessTokenInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTokensAPIService.GetOauthV1AccessTokensTokenGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/v1/access-tokens/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

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
