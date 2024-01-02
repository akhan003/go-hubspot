# PublicTimeOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OffsetDirection** | **string** |  | 
**TimeUnit** | **string** |  | 
**Amount** | **int64** |  | 

## Methods

### NewPublicTimeOffset

`func NewPublicTimeOffset(offsetDirection string, timeUnit string, amount int64, ) *PublicTimeOffset`

NewPublicTimeOffset instantiates a new PublicTimeOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTimeOffsetWithDefaults

`func NewPublicTimeOffsetWithDefaults() *PublicTimeOffset`

NewPublicTimeOffsetWithDefaults instantiates a new PublicTimeOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffsetDirection

`func (o *PublicTimeOffset) GetOffsetDirection() string`

GetOffsetDirection returns the OffsetDirection field if non-nil, zero value otherwise.

### GetOffsetDirectionOk

`func (o *PublicTimeOffset) GetOffsetDirectionOk() (*string, bool)`

GetOffsetDirectionOk returns a tuple with the OffsetDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDirection

`func (o *PublicTimeOffset) SetOffsetDirection(v string)`

SetOffsetDirection sets OffsetDirection field to given value.


### GetTimeUnit

`func (o *PublicTimeOffset) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *PublicTimeOffset) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *PublicTimeOffset) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetAmount

`func (o *PublicTimeOffset) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PublicTimeOffset) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PublicTimeOffset) SetAmount(v int64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


