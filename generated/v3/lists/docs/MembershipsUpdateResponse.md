# MembershipsUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordIdsRemoved** | **[]int64** | The IDs of the records that were &#x60;removed&#x60; from the list. | 
**RecordsIdsAdded** | **[]int64** |  | 
**RecordIdsMissing** | **[]int64** | The IDs of the records that were &#x60;missing&#x60; (e.g. did not exist in the portal) and so were not &#x60;added&#x60; or &#x60;removed&#x60;. | 

## Methods

### NewMembershipsUpdateResponse

`func NewMembershipsUpdateResponse(recordIdsRemoved []int64, recordsIdsAdded []int64, recordIdsMissing []int64, ) *MembershipsUpdateResponse`

NewMembershipsUpdateResponse instantiates a new MembershipsUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipsUpdateResponseWithDefaults

`func NewMembershipsUpdateResponseWithDefaults() *MembershipsUpdateResponse`

NewMembershipsUpdateResponseWithDefaults instantiates a new MembershipsUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordIdsRemoved

`func (o *MembershipsUpdateResponse) GetRecordIdsRemoved() []int64`

GetRecordIdsRemoved returns the RecordIdsRemoved field if non-nil, zero value otherwise.

### GetRecordIdsRemovedOk

`func (o *MembershipsUpdateResponse) GetRecordIdsRemovedOk() (*[]int64, bool)`

GetRecordIdsRemovedOk returns a tuple with the RecordIdsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdsRemoved

`func (o *MembershipsUpdateResponse) SetRecordIdsRemoved(v []int64)`

SetRecordIdsRemoved sets RecordIdsRemoved field to given value.


### GetRecordsIdsAdded

`func (o *MembershipsUpdateResponse) GetRecordsIdsAdded() []int64`

GetRecordsIdsAdded returns the RecordsIdsAdded field if non-nil, zero value otherwise.

### GetRecordsIdsAddedOk

`func (o *MembershipsUpdateResponse) GetRecordsIdsAddedOk() (*[]int64, bool)`

GetRecordsIdsAddedOk returns a tuple with the RecordsIdsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsIdsAdded

`func (o *MembershipsUpdateResponse) SetRecordsIdsAdded(v []int64)`

SetRecordsIdsAdded sets RecordsIdsAdded field to given value.


### GetRecordIdsMissing

`func (o *MembershipsUpdateResponse) GetRecordIdsMissing() []int64`

GetRecordIdsMissing returns the RecordIdsMissing field if non-nil, zero value otherwise.

### GetRecordIdsMissingOk

`func (o *MembershipsUpdateResponse) GetRecordIdsMissingOk() (*[]int64, bool)`

GetRecordIdsMissingOk returns a tuple with the RecordIdsMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdsMissing

`func (o *MembershipsUpdateResponse) SetRecordIdsMissing(v []int64)`

SetRecordIdsMissing sets RecordIdsMissing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


