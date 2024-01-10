# PublicAssociationFilterBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBranchType** | **string** |  | [default to "ASSOCIATION"]
**FilterBranches** | [**[]PublicPropertyAssociationFilterBranchFilterBranchesInner**](PublicPropertyAssociationFilterBranchFilterBranchesInner.md) |  | 
**Filters** | [**[]PublicPropertyAssociationFilterBranchFiltersInner**](PublicPropertyAssociationFilterBranchFiltersInner.md) |  | 
**ObjectTypeId** | **string** |  | 
**Operator** | **string** |  | 
**AssociationTypeId** | **int32** |  | 
**AssociationCategory** | **string** |  | 
**FilterBranchOperator** | **string** |  | 

## Methods

### NewPublicAssociationFilterBranch

`func NewPublicAssociationFilterBranch(filterBranchType string, filterBranches []PublicPropertyAssociationFilterBranchFilterBranchesInner, filters []PublicPropertyAssociationFilterBranchFiltersInner, objectTypeId string, operator string, associationTypeId int32, associationCategory string, filterBranchOperator string, ) *PublicAssociationFilterBranch`

NewPublicAssociationFilterBranch instantiates a new PublicAssociationFilterBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationFilterBranchWithDefaults

`func NewPublicAssociationFilterBranchWithDefaults() *PublicAssociationFilterBranch`

NewPublicAssociationFilterBranchWithDefaults instantiates a new PublicAssociationFilterBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBranchType

`func (o *PublicAssociationFilterBranch) GetFilterBranchType() string`

GetFilterBranchType returns the FilterBranchType field if non-nil, zero value otherwise.

### GetFilterBranchTypeOk

`func (o *PublicAssociationFilterBranch) GetFilterBranchTypeOk() (*string, bool)`

GetFilterBranchTypeOk returns a tuple with the FilterBranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchType

`func (o *PublicAssociationFilterBranch) SetFilterBranchType(v string)`

SetFilterBranchType sets FilterBranchType field to given value.


### GetFilterBranches

`func (o *PublicAssociationFilterBranch) GetFilterBranches() []PublicPropertyAssociationFilterBranchFilterBranchesInner`

GetFilterBranches returns the FilterBranches field if non-nil, zero value otherwise.

### GetFilterBranchesOk

`func (o *PublicAssociationFilterBranch) GetFilterBranchesOk() (*[]PublicPropertyAssociationFilterBranchFilterBranchesInner, bool)`

GetFilterBranchesOk returns a tuple with the FilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranches

`func (o *PublicAssociationFilterBranch) SetFilterBranches(v []PublicPropertyAssociationFilterBranchFilterBranchesInner)`

SetFilterBranches sets FilterBranches field to given value.


### GetFilters

`func (o *PublicAssociationFilterBranch) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PublicAssociationFilterBranch) GetFiltersOk() (*[]PublicPropertyAssociationFilterBranchFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PublicAssociationFilterBranch) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner)`

SetFilters sets Filters field to given value.


### GetObjectTypeId

`func (o *PublicAssociationFilterBranch) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *PublicAssociationFilterBranch) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *PublicAssociationFilterBranch) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetOperator

`func (o *PublicAssociationFilterBranch) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicAssociationFilterBranch) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicAssociationFilterBranch) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAssociationTypeId

`func (o *PublicAssociationFilterBranch) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *PublicAssociationFilterBranch) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *PublicAssociationFilterBranch) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *PublicAssociationFilterBranch) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *PublicAssociationFilterBranch) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *PublicAssociationFilterBranch) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetFilterBranchOperator

`func (o *PublicAssociationFilterBranch) GetFilterBranchOperator() string`

GetFilterBranchOperator returns the FilterBranchOperator field if non-nil, zero value otherwise.

### GetFilterBranchOperatorOk

`func (o *PublicAssociationFilterBranch) GetFilterBranchOperatorOk() (*string, bool)`

GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchOperator

`func (o *PublicAssociationFilterBranch) SetFilterBranchOperator(v string)`

SetFilterBranchOperator sets FilterBranchOperator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


