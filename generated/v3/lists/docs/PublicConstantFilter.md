# PublicConstantFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "CONSTANT"]
**ShouldAccept** | **bool** |  | 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicConstantFilter

`func NewPublicConstantFilter(filterType string, shouldAccept bool, ) *PublicConstantFilter`

NewPublicConstantFilter instantiates a new PublicConstantFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConstantFilterWithDefaults

`func NewPublicConstantFilterWithDefaults() *PublicConstantFilter`

NewPublicConstantFilterWithDefaults instantiates a new PublicConstantFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicConstantFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicConstantFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicConstantFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetShouldAccept

`func (o *PublicConstantFilter) GetShouldAccept() bool`

GetShouldAccept returns the ShouldAccept field if non-nil, zero value otherwise.

### GetShouldAcceptOk

`func (o *PublicConstantFilter) GetShouldAcceptOk() (*bool, bool)`

GetShouldAcceptOk returns a tuple with the ShouldAccept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAccept

`func (o *PublicConstantFilter) SetShouldAccept(v bool)`

SetShouldAccept sets ShouldAccept field to given value.


### GetSource

`func (o *PublicConstantFilter) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PublicConstantFilter) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PublicConstantFilter) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PublicConstantFilter) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


