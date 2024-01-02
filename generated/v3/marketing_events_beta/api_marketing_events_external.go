/*
Marketing Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MarketingEventsExternalAPIService MarketingEventsExternalAPI service
type MarketingEventsExternalAPIService service

type ApiExternalCompleteRequest struct {
	ctx                                 context.Context
	ApiService                          *MarketingEventsExternalAPIService
	externalEventId                     string
	externalAccountId                   *string
	marketingEventCompleteRequestParams *MarketingEventCompleteRequestParams
}

func (r ApiExternalCompleteRequest) ExternalAccountId(externalAccountId string) ApiExternalCompleteRequest {
	r.externalAccountId = &externalAccountId
	return r
}

func (r ApiExternalCompleteRequest) MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams MarketingEventCompleteRequestParams) ApiExternalCompleteRequest {
	r.marketingEventCompleteRequestParams = &marketingEventCompleteRequestParams
	return r
}

func (r ApiExternalCompleteRequest) Execute() (*MarketingEventDefaultResponse, *http.Response, error) {
	return r.ApiService.ExternalCompleteExecute(r)
}

/*
ExternalComplete Method for ExternalComplete

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalEventId
	@return ApiExternalCompleteRequest
*/
func (a *MarketingEventsExternalAPIService) ExternalComplete(ctx context.Context, externalEventId string) ApiExternalCompleteRequest {
	return ApiExternalCompleteRequest{
		ApiService:      a,
		ctx:             ctx,
		externalEventId: externalEventId,
	}
}

// Execute executes the request
//
//	@return MarketingEventDefaultResponse
func (a *MarketingEventsExternalAPIService) ExternalCompleteExecute(r ApiExternalCompleteRequest) (*MarketingEventDefaultResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MarketingEventDefaultResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingEventsExternalAPIService.ExternalComplete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/events/{externalEventId}/complete"
	localVarPath = strings.Replace(localVarPath, "{"+"externalEventId"+"}", url.PathEscape(parameterValueToString(r.externalEventId, "externalEventId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalAccountId == nil {
		return localVarReturnValue, nil, reportError("externalAccountId is required and must be specified")
	}
	if r.marketingEventCompleteRequestParams == nil {
		return localVarReturnValue, nil, reportError("marketingEventCompleteRequestParams is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "externalAccountId", r.externalAccountId, "")
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
	localVarPostBody = r.marketingEventCompleteRequestParams
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
