/*
Quotes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quotes

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the StandardError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardError{}

// StandardError struct for StandardError
type StandardError struct {
	SubCategory map[string]interface{} `json:"subCategory,omitempty"`
	Context     map[string][]string    `json:"context"`
	Links       map[string]string      `json:"links"`
	Id          *string                `json:"id,omitempty"`
	Category    string                 `json:"category"`
	Message     string                 `json:"message"`
	Errors      []ErrorDetail          `json:"errors"`
	Status      string                 `json:"status"`
}

type _StandardError StandardError

// NewStandardError instantiates a new StandardError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardError(context map[string][]string, links map[string]string, category string, message string, errors []ErrorDetail, status string) *StandardError {
	this := StandardError{}
	this.Context = context
	this.Links = links
	this.Category = category
	this.Message = message
	this.Errors = errors
	this.Status = status
	return &this
}

// NewStandardErrorWithDefaults instantiates a new StandardError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardErrorWithDefaults() *StandardError {
	this := StandardError{}
	return &this
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise.
func (o *StandardError) GetSubCategory() map[string]interface{} {
	if o == nil || IsNil(o.SubCategory) {
		var ret map[string]interface{}
		return ret
	}
	return o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardError) GetSubCategoryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SubCategory) {
		return map[string]interface{}{}, false
	}
	return o.SubCategory, true
}

// HasSubCategory returns a boolean if a field has been set.
func (o *StandardError) HasSubCategory() bool {
	if o != nil && !IsNil(o.SubCategory) {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given map[string]interface{} and assigns it to the SubCategory field.
func (o *StandardError) SetSubCategory(v map[string]interface{}) {
	o.SubCategory = v
}

// GetContext returns the Context field value
func (o *StandardError) GetContext() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetContextOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *StandardError) SetContext(v map[string][]string) {
	o.Context = v
}

// GetLinks returns the Links field value
func (o *StandardError) GetLinks() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetLinksOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *StandardError) SetLinks(v map[string]string) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandardError) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardError) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandardError) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StandardError) SetId(v string) {
	o.Id = &v
}

// GetCategory returns the Category field value
func (o *StandardError) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *StandardError) SetCategory(v string) {
	o.Category = v
}

// GetMessage returns the Message field value
func (o *StandardError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *StandardError) SetMessage(v string) {
	o.Message = v
}

// GetErrors returns the Errors field value
func (o *StandardError) GetErrors() []ErrorDetail {
	if o == nil {
		var ret []ErrorDetail
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetErrorsOk() ([]ErrorDetail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *StandardError) SetErrors(v []ErrorDetail) {
	o.Errors = v
}

// GetStatus returns the Status field value
func (o *StandardError) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StandardError) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StandardError) SetStatus(v string) {
	o.Status = v
}

func (o StandardError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubCategory) {
		toSerialize["subCategory"] = o.SubCategory
	}
	toSerialize["context"] = o.Context
	toSerialize["links"] = o.Links
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["category"] = o.Category
	toSerialize["message"] = o.Message
	toSerialize["errors"] = o.Errors
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *StandardError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"context",
		"links",
		"category",
		"message",
		"errors",
		"status",
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

	varStandardError := _StandardError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStandardError)

	if err != nil {
		return err
	}

	*o = StandardError(varStandardError)

	return err
}

type NullableStandardError struct {
	value *StandardError
	isSet bool
}

func (v NullableStandardError) Get() *StandardError {
	return v.value
}

func (v *NullableStandardError) Set(val *StandardError) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardError) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardError(val *StandardError) *NullableStandardError {
	return &NullableStandardError{value: val, isSet: true}
}

func (v NullableStandardError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
