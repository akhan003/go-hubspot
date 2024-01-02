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
)

// checks if the ListSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSearchRequest{}

// ListSearchRequest The request object used for searching through lists.
type ListSearchRequest struct {
	// Value used to paginate through lists. The `offset` provided in the response can be used in the next request to fetch the next page of results. Defaults to `0` if no offset is provided.
	Offset int32 `json:"offset"`
	// The `query` that will be used to search for lists by list name. If no `query` is provided, then the results will include all lists.
	Query *string `json:"query,omitempty"`
	// The number of lists to include in the response. Defaults to `20` if no value is provided. The max `count` is `500`.
	Count *int32 `json:"count,omitempty"`
	// The property names of any additional list properties to include in the response. Properties that do not exist or that are empty for a particular list are not included in the response.  By default, all requests will fetch the following properties for each list: `hs_list_size`, `hs_last_record_added_at`, `hs_last_record_removed_at`, `hs_folder_name`, and `hs_list_reference_count`.
	AdditionalPropertiesField []string `json:"additionalProperties"`
}

type _ListSearchRequest ListSearchRequest

// NewListSearchRequest instantiates a new ListSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSearchRequest(offset int32, additionalPropertiesField []string) *ListSearchRequest {
	this := ListSearchRequest{}
	this.Offset = offset
	this.AdditionalPropertiesField = additionalPropertiesField
	return &this
}

// NewListSearchRequestWithDefaults instantiates a new ListSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSearchRequestWithDefaults() *ListSearchRequest {
	this := ListSearchRequest{}
	return &this
}

// GetOffset returns the Offset field value
func (o *ListSearchRequest) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListSearchRequest) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListSearchRequest) SetOffset(v int32) {
	o.Offset = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ListSearchRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSearchRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ListSearchRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ListSearchRequest) SetQuery(v string) {
	o.Query = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListSearchRequest) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSearchRequest) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListSearchRequest) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ListSearchRequest) SetCount(v int32) {
	o.Count = &v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value
func (o *ListSearchRequest) GetAdditionalPropertiesField() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value
// and a boolean to check if the value has been set.
func (o *ListSearchRequest) GetAdditionalPropertiesFieldOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// SetAdditionalPropertiesField sets field value
func (o *ListSearchRequest) SetAdditionalPropertiesField(v []string) {
	o.AdditionalPropertiesField = v
}

func (o ListSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offset"] = o.Offset
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	return toSerialize, nil
}

func (o *ListSearchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offset",
		"additionalProperties",
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

	varListSearchRequest := _ListSearchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListSearchRequest)

	if err != nil {
		return err
	}

	*o = ListSearchRequest(varListSearchRequest)

	return err
}

type NullableListSearchRequest struct {
	value *ListSearchRequest
	isSet bool
}

func (v NullableListSearchRequest) Get() *ListSearchRequest {
	return v.value
}

func (v *NullableListSearchRequest) Set(val *ListSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSearchRequest(val *ListSearchRequest) *NullableListSearchRequest {
	return &NullableListSearchRequest{value: val, isSet: true}
}

func (v NullableListSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
