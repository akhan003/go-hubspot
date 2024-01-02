# MembershipChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordIdsToRemove** | **[]int64** |  | 
**RecordIdsToAdd** | **[]int64** |  | 

## Methods

### NewMembershipChangeRequest

`func NewMembershipChangeRequest(recordIdsToRemove []int64, recordIdsToAdd []int64, ) *MembershipChangeRequest`

NewMembershipChangeRequest instantiates a new MembershipChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipChangeRequestWithDefaults

`func NewMembershipChangeRequestWithDefaults() *MembershipChangeRequest`

NewMembershipChangeRequestWithDefaults instantiates a new MembershipChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordIdsToRemove

`func (o *MembershipChangeRequest) GetRecordIdsToRemove() []int64`

GetRecordIdsToRemove returns the RecordIdsToRemove field if non-nil, zero value otherwise.

### GetRecordIdsToRemoveOk

`func (o *MembershipChangeRequest) GetRecordIdsToRemoveOk() (*[]int64, bool)`

GetRecordIdsToRemoveOk returns a tuple with the RecordIdsToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdsToRemove

`func (o *MembershipChangeRequest) SetRecordIdsToRemove(v []int64)`

SetRecordIdsToRemove sets RecordIdsToRemove field to given value.


### GetRecordIdsToAdd

`func (o *MembershipChangeRequest) GetRecordIdsToAdd() []int64`

GetRecordIdsToAdd returns the RecordIdsToAdd field if non-nil, zero value otherwise.

### GetRecordIdsToAddOk

`func (o *MembershipChangeRequest) GetRecordIdsToAddOk() (*[]int64, bool)`

GetRecordIdsToAddOk returns a tuple with the RecordIdsToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdsToAdd

`func (o *MembershipChangeRequest) SetRecordIdsToAdd(v []int64)`

SetRecordIdsToAdd sets RecordIdsToAdd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


