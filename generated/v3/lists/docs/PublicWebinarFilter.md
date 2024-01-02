# PublicWebinarFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "WEBINAR"]
**Operator** | **string** |  | 
**WebinarId** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicWebinarFilter

`func NewPublicWebinarFilter(filterType string, operator string, ) *PublicWebinarFilter`

NewPublicWebinarFilter instantiates a new PublicWebinarFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicWebinarFilterWithDefaults

`func NewPublicWebinarFilterWithDefaults() *PublicWebinarFilter`

NewPublicWebinarFilterWithDefaults instantiates a new PublicWebinarFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicWebinarFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicWebinarFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicWebinarFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetOperator

`func (o *PublicWebinarFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicWebinarFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicWebinarFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetWebinarId

`func (o *PublicWebinarFilter) GetWebinarId() string`

GetWebinarId returns the WebinarId field if non-nil, zero value otherwise.

### GetWebinarIdOk

`func (o *PublicWebinarFilter) GetWebinarIdOk() (*string, bool)`

GetWebinarIdOk returns a tuple with the WebinarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebinarId

`func (o *PublicWebinarFilter) SetWebinarId(v string)`

SetWebinarId sets WebinarId field to given value.

### HasWebinarId

`func (o *PublicWebinarFilter) HasWebinarId() bool`

HasWebinarId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


