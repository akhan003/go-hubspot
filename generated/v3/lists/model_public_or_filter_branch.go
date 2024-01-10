/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"encoding/json"
	"fmt"
)

// checks if the PublicOrFilterBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicOrFilterBranch{}

// PublicOrFilterBranch struct for PublicOrFilterBranch
type PublicOrFilterBranch struct {
	FilterBranchType     string                                                     `json:"filterBranchType"`
	FilterBranches       []PublicPropertyAssociationFilterBranchFilterBranchesInner `json:"filterBranches"`
	Filters              []PublicPropertyAssociationFilterBranchFiltersInner        `json:"filters"`
	FilterBranchOperator string                                                     `json:"filterBranchOperator"`
	AdditionalProperties map[string]interface{}
}

type _PublicOrFilterBranch PublicOrFilterBranch

// NewPublicOrFilterBranch instantiates a new PublicOrFilterBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicOrFilterBranch(filterBranchType string, filterBranches []PublicPropertyAssociationFilterBranchFilterBranchesInner, filters []PublicPropertyAssociationFilterBranchFiltersInner, filterBranchOperator string) *PublicOrFilterBranch {
	this := PublicOrFilterBranch{}
	this.FilterBranchType = filterBranchType
	this.FilterBranches = filterBranches
	this.Filters = filters
	this.FilterBranchOperator = filterBranchOperator
	return &this
}

// NewPublicOrFilterBranchWithDefaults instantiates a new PublicOrFilterBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicOrFilterBranchWithDefaults() *PublicOrFilterBranch {
	this := PublicOrFilterBranch{}
	var filterBranchType string = "OR"
	this.FilterBranchType = filterBranchType
	return &this
}

// GetFilterBranchType returns the FilterBranchType field value
func (o *PublicOrFilterBranch) GetFilterBranchType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterBranchType
}

// GetFilterBranchTypeOk returns a tuple with the FilterBranchType field value
// and a boolean to check if the value has been set.
func (o *PublicOrFilterBranch) GetFilterBranchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterBranchType, true
}

// SetFilterBranchType sets field value
func (o *PublicOrFilterBranch) SetFilterBranchType(v string) {
	o.FilterBranchType = v
}

// GetFilterBranches returns the FilterBranches field value
func (o *PublicOrFilterBranch) GetFilterBranches() []PublicPropertyAssociationFilterBranchFilterBranchesInner {
	if o == nil {
		var ret []PublicPropertyAssociationFilterBranchFilterBranchesInner
		return ret
	}

	return o.FilterBranches
}

// GetFilterBranchesOk returns a tuple with the FilterBranches field value
// and a boolean to check if the value has been set.
func (o *PublicOrFilterBranch) GetFilterBranchesOk() ([]PublicPropertyAssociationFilterBranchFilterBranchesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterBranches, true
}

// SetFilterBranches sets field value
func (o *PublicOrFilterBranch) SetFilterBranches(v []PublicPropertyAssociationFilterBranchFilterBranchesInner) {
	o.FilterBranches = v
}

// GetFilters returns the Filters field value
func (o *PublicOrFilterBranch) GetFilters() []PublicPropertyAssociationFilterBranchFiltersInner {
	if o == nil {
		var ret []PublicPropertyAssociationFilterBranchFiltersInner
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *PublicOrFilterBranch) GetFiltersOk() ([]PublicPropertyAssociationFilterBranchFiltersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *PublicOrFilterBranch) SetFilters(v []PublicPropertyAssociationFilterBranchFiltersInner) {
	o.Filters = v
}

// GetFilterBranchOperator returns the FilterBranchOperator field value
func (o *PublicOrFilterBranch) GetFilterBranchOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterBranchOperator
}

// GetFilterBranchOperatorOk returns a tuple with the FilterBranchOperator field value
// and a boolean to check if the value has been set.
func (o *PublicOrFilterBranch) GetFilterBranchOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterBranchOperator, true
}

// SetFilterBranchOperator sets field value
func (o *PublicOrFilterBranch) SetFilterBranchOperator(v string) {
	o.FilterBranchOperator = v
}

func (o PublicOrFilterBranch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicOrFilterBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterBranchType"] = o.FilterBranchType
	toSerialize["filterBranches"] = o.FilterBranches
	toSerialize["filters"] = o.Filters
	toSerialize["filterBranchOperator"] = o.FilterBranchOperator

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PublicOrFilterBranch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterBranchType",
		"filterBranches",
		"filters",
		"filterBranchOperator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicOrFilterBranch := _PublicOrFilterBranch{}

	err = json.Unmarshal(data, &varPublicOrFilterBranch)

	if err != nil {
		return err
	}

	*o = PublicOrFilterBranch(varPublicOrFilterBranch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filterBranchType")
		delete(additionalProperties, "filterBranches")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "filterBranchOperator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicOrFilterBranch struct {
	value *PublicOrFilterBranch
	isSet bool
}

func (v NullablePublicOrFilterBranch) Get() *PublicOrFilterBranch {
	return v.value
}

func (v *NullablePublicOrFilterBranch) Set(val *PublicOrFilterBranch) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicOrFilterBranch) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicOrFilterBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicOrFilterBranch(val *PublicOrFilterBranch) *NullablePublicOrFilterBranch {
	return &NullablePublicOrFilterBranch{value: val, isSet: true}
}

func (v NullablePublicOrFilterBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicOrFilterBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
