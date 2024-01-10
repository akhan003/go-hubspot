# PublicEmailEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "EMAIL_EVENT"]
**Level** | **string** |  | 
**EmailId** | **int32** |  | 
**AppId** | **int32** |  | 
**Operator** | **string** |  | 
**ClickUrl** | Pointer to **string** |  | [optional] 
**PruningRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 

## Methods

### NewPublicEmailEventFilter

`func NewPublicEmailEventFilter(filterType string, level string, emailId int32, appId int32, operator string, ) *PublicEmailEventFilter`

NewPublicEmailEventFilter instantiates a new PublicEmailEventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEmailEventFilterWithDefaults

`func NewPublicEmailEventFilterWithDefaults() *PublicEmailEventFilter`

NewPublicEmailEventFilterWithDefaults instantiates a new PublicEmailEventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicEmailEventFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicEmailEventFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicEmailEventFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetLevel

`func (o *PublicEmailEventFilter) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PublicEmailEventFilter) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PublicEmailEventFilter) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetEmailId

`func (o *PublicEmailEventFilter) GetEmailId() int32`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *PublicEmailEventFilter) GetEmailIdOk() (*int32, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *PublicEmailEventFilter) SetEmailId(v int32)`

SetEmailId sets EmailId field to given value.


### GetAppId

`func (o *PublicEmailEventFilter) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *PublicEmailEventFilter) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *PublicEmailEventFilter) SetAppId(v int32)`

SetAppId sets AppId field to given value.


### GetOperator

`func (o *PublicEmailEventFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicEmailEventFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicEmailEventFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetClickUrl

`func (o *PublicEmailEventFilter) GetClickUrl() string`

GetClickUrl returns the ClickUrl field if non-nil, zero value otherwise.

### GetClickUrlOk

`func (o *PublicEmailEventFilter) GetClickUrlOk() (*string, bool)`

GetClickUrlOk returns a tuple with the ClickUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickUrl

`func (o *PublicEmailEventFilter) SetClickUrl(v string)`

SetClickUrl sets ClickUrl field to given value.

### HasClickUrl

`func (o *PublicEmailEventFilter) HasClickUrl() bool`

HasClickUrl returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *PublicEmailEventFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicEmailEventFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicEmailEventFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicEmailEventFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


