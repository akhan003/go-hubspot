# PublicPropertyAssociationFilterBranchFilterBranchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBranchType** | **string** |  | [default to "ASSOCIATION"]
**FilterBranches** | [**[]PublicPropertyAssociationFilterBranchFilterBranchesInner**](PublicPropertyAssociationFilterBranchFilterBranchesInner.md) |  | 
**Filters** | [**[]PublicPropertyAssociationFilterBranchFiltersInner**](PublicPropertyAssociationFilterBranchFiltersInner.md) |  | 
**FilterBranchOperator** | **string** |  | 
**EventTypeId** | **string** |  | 
**Operator** | **string** |  | 
**ObjectTypeId** | **string** |  | 
**PropertyWithObjectId** | **string** |  | 
**AssociationTypeId** | **int32** |  | 
**AssociationCategory** | **string** |  | 

## Methods

### NewPublicPropertyAssociationFilterBranchFilterBranchesInner

`func NewPublicPropertyAssociationFilterBranchFilterBranchesInner(filterBranchType string, filterBranches []PublicPropertyAssociationFilterBranchFilterBranchesInner, filters []PublicPropertyAssociationFilterBranchFiltersInner, filterBranchOperator string, eventTypeId string, operator string, objectTypeId string, propertyWithObjectId string, associationTypeId int32, associationCategory string, ) *PublicPropertyAssociationFilterBranchFilterBranchesInner`

NewPublicPropertyAssociationFilterBranchFilterBranchesInner instantiates a new PublicPropertyAssociationFilterBranchFilterBranchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPropertyAssociationFilterBranchFilterBranchesInnerWithDefaults

`func NewPublicPropertyAssociationFilterBranchFilterBranchesInnerWithDefaults() *PublicPropertyAssociationFilterBranchFilterBranchesInner`

NewPublicPropertyAssociationFilterBranchFilterBranchesInnerWithDefaults instantiates a new PublicPropertyAssociationFilterBranchFilterBranchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBranchType

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFilterBranchType() string`

GetFilterBranchType returns the FilterBranchType field if non-nil, zero value otherwise.

### GetFilterBranchTypeOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFilterBranchTypeOk() (*string, bool)`

GetFilterBranchTypeOk returns a tuple with the FilterBranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchType

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetFilterBranchType(v string)`

SetFilterBranchType sets FilterBranchType field to given value.


### GetFilterBranches

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFilterBranches() []PublicPropertyAssociationFilterBranchFilterBranchesInner`

GetFilterBranches returns the FilterBranches field if non-nil, zero value otherwise.

### GetFilterBranchesOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFilterBranchesOk() (*[]PublicPropertyAssociationFilterBranchFilterBranchesInner, bool)`

GetFilterBranchesOk returns a tuple with the FilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranches

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetFilterBranches(v []PublicPropertyAssociationFilterBranchFilterBranchesInner)`

SetFilterBranches sets FilterBranches field to given value.


### GetFilters

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFiltersOk() (*[]PublicPropertyAssociationFilterBranchFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner)`

SetFilters sets Filters field to given value.


### GetFilterBranchOperator

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFilterBranchOperator() string`

GetFilterBranchOperator returns the FilterBranchOperator field if non-nil, zero value otherwise.

### GetFilterBranchOperatorOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetFilterBranchOperatorOk() (*string, bool)`

GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchOperator

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetFilterBranchOperator(v string)`

SetFilterBranchOperator sets FilterBranchOperator field to given value.


### GetEventTypeId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetEventTypeId() string`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetEventTypeIdOk() (*string, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetEventTypeId(v string)`

SetEventTypeId sets EventTypeId field to given value.


### GetOperator

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetObjectTypeId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetPropertyWithObjectId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetPropertyWithObjectId() string`

GetPropertyWithObjectId returns the PropertyWithObjectId field if non-nil, zero value otherwise.

### GetPropertyWithObjectIdOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetPropertyWithObjectIdOk() (*string, bool)`

GetPropertyWithObjectIdOk returns a tuple with the PropertyWithObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyWithObjectId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetPropertyWithObjectId(v string)`

SetPropertyWithObjectId sets PropertyWithObjectId field to given value.


### GetAssociationTypeId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *PublicPropertyAssociationFilterBranchFilterBranchesInner) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


