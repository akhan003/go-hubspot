# ListSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | The total number of lists that match the search criteria. | 
**Offset** | **int32** | Value to be passed in a future request to paginate through list search results. | 
**Lists** | [**[]PublicObjectListSearchResult**](PublicObjectListSearchResult.md) | The lists that matched the search criteria. | 
**HasMore** | **bool** | Whether or not there are more results to page through. | 

## Methods

### NewListSearchResponse

`func NewListSearchResponse(total int32, offset int32, lists []PublicObjectListSearchResult, hasMore bool, ) *ListSearchResponse`

NewListSearchResponse instantiates a new ListSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSearchResponseWithDefaults

`func NewListSearchResponseWithDefaults() *ListSearchResponse`

NewListSearchResponseWithDefaults instantiates a new ListSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListSearchResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListSearchResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListSearchResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetOffset

`func (o *ListSearchResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListSearchResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListSearchResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLists

`func (o *ListSearchResponse) GetLists() []PublicObjectListSearchResult`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *ListSearchResponse) GetListsOk() (*[]PublicObjectListSearchResult, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *ListSearchResponse) SetLists(v []PublicObjectListSearchResult)`

SetLists sets Lists field to given value.


### GetHasMore

`func (o *ListSearchResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListSearchResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListSearchResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


