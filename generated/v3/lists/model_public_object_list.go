/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the PublicObjectList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicObjectList{}

// PublicObjectList An object list definition.
type PublicObjectList struct {
	// The processing type of the list.
	ProcessingType string `json:"processingType"`
	// The object type of the list.
	ObjectTypeId string `json:"objectTypeId"`
	// The ID of the user that last updated the list.
	UpdatedById *int32 `json:"updatedById,omitempty"`
	// The time when the filters for this list were last updated.
	FiltersUpdatedAt *time.Time `json:"filtersUpdatedAt,omitempty"`
	// The **ILS ID** of the list.
	ListId int32 `json:"listId"`
	// The time when the list was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The processing status of the list.
	ProcessingStatus string `json:"processingStatus"`
	// The time when the list was deleted.
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The version of the list.
	ListVersion int32 `json:"listVersion"`
	// The name of the list.
	Name string `json:"name"`
	// The ID of the user that created the list.
	CreatedById  *int32                                                    `json:"createdById,omitempty"`
	FilterBranch *PublicPropertyAssociationFilterBranchFilterBranchesInner `json:"filterBranch,omitempty"`
	// The time the list was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _PublicObjectList PublicObjectList

// NewPublicObjectList instantiates a new PublicObjectList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicObjectList(processingType string, objectTypeId string, listId int32, processingStatus string, listVersion int32, name string) *PublicObjectList {
	this := PublicObjectList{}
	this.ProcessingType = processingType
	this.ObjectTypeId = objectTypeId
	this.ListId = listId
	this.ProcessingStatus = processingStatus
	this.ListVersion = listVersion
	this.Name = name
	return &this
}

// NewPublicObjectListWithDefaults instantiates a new PublicObjectList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicObjectListWithDefaults() *PublicObjectList {
	this := PublicObjectList{}
	return &this
}

// GetProcessingType returns the ProcessingType field value
func (o *PublicObjectList) GetProcessingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingType
}

// GetProcessingTypeOk returns a tuple with the ProcessingType field value
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetProcessingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingType, true
}

// SetProcessingType sets field value
func (o *PublicObjectList) SetProcessingType(v string) {
	o.ProcessingType = v
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *PublicObjectList) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *PublicObjectList) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetUpdatedById returns the UpdatedById field value if set, zero value otherwise.
func (o *PublicObjectList) GetUpdatedById() int32 {
	if o == nil || IsNil(o.UpdatedById) {
		var ret int32
		return ret
	}
	return *o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetUpdatedByIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedById) {
		return nil, false
	}
	return o.UpdatedById, true
}

// HasUpdatedById returns a boolean if a field has been set.
func (o *PublicObjectList) HasUpdatedById() bool {
	if o != nil && !IsNil(o.UpdatedById) {
		return true
	}

	return false
}

// SetUpdatedById gets a reference to the given int32 and assigns it to the UpdatedById field.
func (o *PublicObjectList) SetUpdatedById(v int32) {
	o.UpdatedById = &v
}

// GetFiltersUpdatedAt returns the FiltersUpdatedAt field value if set, zero value otherwise.
func (o *PublicObjectList) GetFiltersUpdatedAt() time.Time {
	if o == nil || IsNil(o.FiltersUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.FiltersUpdatedAt
}

// GetFiltersUpdatedAtOk returns a tuple with the FiltersUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetFiltersUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FiltersUpdatedAt) {
		return nil, false
	}
	return o.FiltersUpdatedAt, true
}

// HasFiltersUpdatedAt returns a boolean if a field has been set.
func (o *PublicObjectList) HasFiltersUpdatedAt() bool {
	if o != nil && !IsNil(o.FiltersUpdatedAt) {
		return true
	}

	return false
}

// SetFiltersUpdatedAt gets a reference to the given time.Time and assigns it to the FiltersUpdatedAt field.
func (o *PublicObjectList) SetFiltersUpdatedAt(v time.Time) {
	o.FiltersUpdatedAt = &v
}

// GetListId returns the ListId field value
func (o *PublicObjectList) GetListId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetListIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *PublicObjectList) SetListId(v int32) {
	o.ListId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PublicObjectList) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PublicObjectList) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PublicObjectList) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetProcessingStatus returns the ProcessingStatus field value
func (o *PublicObjectList) GetProcessingStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessingStatus
}

// GetProcessingStatusOk returns a tuple with the ProcessingStatus field value
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetProcessingStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingStatus, true
}

// SetProcessingStatus sets field value
func (o *PublicObjectList) SetProcessingStatus(v string) {
	o.ProcessingStatus = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *PublicObjectList) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *PublicObjectList) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *PublicObjectList) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetListVersion returns the ListVersion field value
func (o *PublicObjectList) GetListVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListVersion
}

// GetListVersionOk returns a tuple with the ListVersion field value
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetListVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListVersion, true
}

// SetListVersion sets field value
func (o *PublicObjectList) SetListVersion(v int32) {
	o.ListVersion = v
}

// GetName returns the Name field value
func (o *PublicObjectList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PublicObjectList) SetName(v string) {
	o.Name = v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *PublicObjectList) GetCreatedById() int32 {
	if o == nil || IsNil(o.CreatedById) {
		var ret int32
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetCreatedByIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *PublicObjectList) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given int32 and assigns it to the CreatedById field.
func (o *PublicObjectList) SetCreatedById(v int32) {
	o.CreatedById = &v
}

// GetFilterBranch returns the FilterBranch field value if set, zero value otherwise.
func (o *PublicObjectList) GetFilterBranch() PublicPropertyAssociationFilterBranchFilterBranchesInner {
	if o == nil || IsNil(o.FilterBranch) {
		var ret PublicPropertyAssociationFilterBranchFilterBranchesInner
		return ret
	}
	return *o.FilterBranch
}

// GetFilterBranchOk returns a tuple with the FilterBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetFilterBranchOk() (*PublicPropertyAssociationFilterBranchFilterBranchesInner, bool) {
	if o == nil || IsNil(o.FilterBranch) {
		return nil, false
	}
	return o.FilterBranch, true
}

// HasFilterBranch returns a boolean if a field has been set.
func (o *PublicObjectList) HasFilterBranch() bool {
	if o != nil && !IsNil(o.FilterBranch) {
		return true
	}

	return false
}

// SetFilterBranch gets a reference to the given PublicPropertyAssociationFilterBranchFilterBranchesInner and assigns it to the FilterBranch field.
func (o *PublicObjectList) SetFilterBranch(v PublicPropertyAssociationFilterBranchFilterBranchesInner) {
	o.FilterBranch = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PublicObjectList) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicObjectList) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PublicObjectList) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PublicObjectList) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PublicObjectList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicObjectList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["processingType"] = o.ProcessingType
	toSerialize["objectTypeId"] = o.ObjectTypeId
	if !IsNil(o.UpdatedById) {
		toSerialize["updatedById"] = o.UpdatedById
	}
	if !IsNil(o.FiltersUpdatedAt) {
		toSerialize["filtersUpdatedAt"] = o.FiltersUpdatedAt
	}
	toSerialize["listId"] = o.ListId
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["processingStatus"] = o.ProcessingStatus
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	toSerialize["listVersion"] = o.ListVersion
	toSerialize["name"] = o.Name
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !IsNil(o.FilterBranch) {
		toSerialize["filterBranch"] = o.FilterBranch
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *PublicObjectList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"processingType",
		"objectTypeId",
		"listId",
		"processingStatus",
		"listVersion",
		"name",
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

	varPublicObjectList := _PublicObjectList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicObjectList)

	if err != nil {
		return err
	}

	*o = PublicObjectList(varPublicObjectList)

	return err
}

type NullablePublicObjectList struct {
	value *PublicObjectList
	isSet bool
}

func (v NullablePublicObjectList) Get() *PublicObjectList {
	return v.value
}

func (v *NullablePublicObjectList) Set(val *PublicObjectList) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicObjectList) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicObjectList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicObjectList(val *PublicObjectList) *NullablePublicObjectList {
	return &NullablePublicObjectList{value: val, isSet: true}
}

func (v NullablePublicObjectList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicObjectList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
