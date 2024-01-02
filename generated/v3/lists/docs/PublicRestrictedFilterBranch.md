# PublicRestrictedFilterBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBranchType** | **string** |  | [default to "RESTRICTED"]
**FilterBranchOperator** | **string** |  | 
**FilterBranches** | [**[]PublicPropertyAssociationFilterBranchFilterBranchesInner**](PublicPropertyAssociationFilterBranchFilterBranchesInner.md) |  | 
**Filters** | [**[]PublicPropertyAssociationFilterBranchFiltersInner**](PublicPropertyAssociationFilterBranchFiltersInner.md) |  | 

## Methods

### NewPublicRestrictedFilterBranch

`func NewPublicRestrictedFilterBranch(filterBranchType string, filterBranchOperator string, filterBranches []PublicPropertyAssociationFilterBranchFilterBranchesInner, filters []PublicPropertyAssociationFilterBranchFiltersInner, ) *PublicRestrictedFilterBranch`

NewPublicRestrictedFilterBranch instantiates a new PublicRestrictedFilterBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicRestrictedFilterBranchWithDefaults

`func NewPublicRestrictedFilterBranchWithDefaults() *PublicRestrictedFilterBranch`

NewPublicRestrictedFilterBranchWithDefaults instantiates a new PublicRestrictedFilterBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBranchType

`func (o *PublicRestrictedFilterBranch) GetFilterBranchType() string`

GetFilterBranchType returns the FilterBranchType field if non-nil, zero value otherwise.

### GetFilterBranchTypeOk

`func (o *PublicRestrictedFilterBranch) GetFilterBranchTypeOk() (*string, bool)`

GetFilterBranchTypeOk returns a tuple with the FilterBranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchType

`func (o *PublicRestrictedFilterBranch) SetFilterBranchType(v string)`

SetFilterBranchType sets FilterBranchType field to given value.


### GetFilterBranchOperator

`func (o *PublicRestrictedFilterBranch) GetFilterBranchOperator() string`

GetFilterBranchOperator returns the FilterBranchOperator field if non-nil, zero value otherwise.

### GetFilterBranchOperatorOk

`func (o *PublicRestrictedFilterBranch) GetFilterBranchOperatorOk() (*string, bool)`

GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchOperator

`func (o *PublicRestrictedFilterBranch) SetFilterBranchOperator(v string)`

SetFilterBranchOperator sets FilterBranchOperator field to given value.


### GetFilterBranches

`func (o *PublicRestrictedFilterBranch) GetFilterBranches() []PublicPropertyAssociationFilterBranchFilterBranchesInner`

GetFilterBranches returns the FilterBranches field if non-nil, zero value otherwise.

### GetFilterBranchesOk

`func (o *PublicRestrictedFilterBranch) GetFilterBranchesOk() (*[]PublicPropertyAssociationFilterBranchFilterBranchesInner, bool)`

GetFilterBranchesOk returns a tuple with the FilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranches

`func (o *PublicRestrictedFilterBranch) SetFilterBranches(v []PublicPropertyAssociationFilterBranchFilterBranchesInner)`

SetFilterBranches sets FilterBranches field to given value.


### GetFilters

`func (o *PublicRestrictedFilterBranch) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PublicRestrictedFilterBranch) GetFiltersOk() (*[]PublicPropertyAssociationFilterBranchFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PublicRestrictedFilterBranch) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


