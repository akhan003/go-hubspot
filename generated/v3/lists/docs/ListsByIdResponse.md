# ListsByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lists** | [**[]PublicObjectList**](PublicObjectList.md) | The object list definitions. | 

## Methods

### NewListsByIdResponse

`func NewListsByIdResponse(lists []PublicObjectList, ) *ListsByIdResponse`

NewListsByIdResponse instantiates a new ListsByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsByIdResponseWithDefaults

`func NewListsByIdResponseWithDefaults() *ListsByIdResponse`

NewListsByIdResponseWithDefaults instantiates a new ListsByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLists

`func (o *ListsByIdResponse) GetLists() []PublicObjectList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *ListsByIdResponse) GetListsOk() (*[]PublicObjectList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *ListsByIdResponse) SetLists(v []PublicObjectList)`

SetLists sets Lists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


