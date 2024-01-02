# PublicNotAllFilterBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterBranchType** | **string** |  | [default to "NOT_ALL"]
**FilterBranchOperator** | **string** |  | 
**FilterBranches** | [**[]PublicPropertyAssociationFilterBranchFilterBranchesInner**](PublicPropertyAssociationFilterBranchFilterBranchesInner.md) |  | 
**Filters** | [**[]PublicPropertyAssociationFilterBranchFiltersInner**](PublicPropertyAssociationFilterBranchFiltersInner.md) |  | 

## Methods

### NewPublicNotAllFilterBranch

`func NewPublicNotAllFilterBranch(filterBranchType string, filterBranchOperator string, filterBranches []PublicPropertyAssociationFilterBranchFilterBranchesInner, filters []PublicPropertyAssociationFilterBranchFiltersInner, ) *PublicNotAllFilterBranch`

NewPublicNotAllFilterBranch instantiates a new PublicNotAllFilterBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNotAllFilterBranchWithDefaults

`func NewPublicNotAllFilterBranchWithDefaults() *PublicNotAllFilterBranch`

NewPublicNotAllFilterBranchWithDefaults instantiates a new PublicNotAllFilterBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterBranchType

`func (o *PublicNotAllFilterBranch) GetFilterBranchType() string`

GetFilterBranchType returns the FilterBranchType field if non-nil, zero value otherwise.

### GetFilterBranchTypeOk

`func (o *PublicNotAllFilterBranch) GetFilterBranchTypeOk() (*string, bool)`

GetFilterBranchTypeOk returns a tuple with the FilterBranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchType

`func (o *PublicNotAllFilterBranch) SetFilterBranchType(v string)`

SetFilterBranchType sets FilterBranchType field to given value.


### GetFilterBranchOperator

`func (o *PublicNotAllFilterBranch) GetFilterBranchOperator() string`

GetFilterBranchOperator returns the FilterBranchOperator field if non-nil, zero value otherwise.

### GetFilterBranchOperatorOk

`func (o *PublicNotAllFilterBranch) GetFilterBranchOperatorOk() (*string, bool)`

GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranchOperator

`func (o *PublicNotAllFilterBranch) SetFilterBranchOperator(v string)`

SetFilterBranchOperator sets FilterBranchOperator field to given value.


### GetFilterBranches

`func (o *PublicNotAllFilterBranch) GetFilterBranches() []PublicPropertyAssociationFilterBranchFilterBranchesInner`

GetFilterBranches returns the FilterBranches field if non-nil, zero value otherwise.

### GetFilterBranchesOk

`func (o *PublicNotAllFilterBranch) GetFilterBranchesOk() (*[]PublicPropertyAssociationFilterBranchFilterBranchesInner, bool)`

GetFilterBranchesOk returns a tuple with the FilterBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranches

`func (o *PublicNotAllFilterBranch) SetFilterBranches(v []PublicPropertyAssociationFilterBranchFilterBranchesInner)`

SetFilterBranches sets FilterBranches field to given value.


### GetFilters

`func (o *PublicNotAllFilterBranch) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PublicNotAllFilterBranch) GetFiltersOk() (*[]PublicPropertyAssociationFilterBranchFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PublicNotAllFilterBranch) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


