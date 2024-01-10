# PublicPropertyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "PROPERTY"]
**Property** | **string** |  | 
**Operation** | [**PublicPropertyFilterOperation**](PublicPropertyFilterOperation.md) |  | 

## Methods

### NewPublicPropertyFilter

`func NewPublicPropertyFilter(filterType string, property string, operation PublicPropertyFilterOperation, ) *PublicPropertyFilter`

NewPublicPropertyFilter instantiates a new PublicPropertyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPropertyFilterWithDefaults

`func NewPublicPropertyFilterWithDefaults() *PublicPropertyFilter`

NewPublicPropertyFilterWithDefaults instantiates a new PublicPropertyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicPropertyFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicPropertyFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicPropertyFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetProperty

`func (o *PublicPropertyFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PublicPropertyFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PublicPropertyFilter) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetOperation

`func (o *PublicPropertyFilter) GetOperation() PublicPropertyFilterOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PublicPropertyFilter) GetOperationOk() (*PublicPropertyFilterOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PublicPropertyFilter) SetOperation(v PublicPropertyFilterOperation)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


