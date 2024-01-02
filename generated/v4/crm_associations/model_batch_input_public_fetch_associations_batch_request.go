/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchInputPublicFetchAssociationsBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputPublicFetchAssociationsBatchRequest{}

// BatchInputPublicFetchAssociationsBatchRequest struct for BatchInputPublicFetchAssociationsBatchRequest
type BatchInputPublicFetchAssociationsBatchRequest struct {
	Inputs []PublicFetchAssociationsBatchRequest `json:"inputs"`
}

type _BatchInputPublicFetchAssociationsBatchRequest BatchInputPublicFetchAssociationsBatchRequest

// NewBatchInputPublicFetchAssociationsBatchRequest instantiates a new BatchInputPublicFetchAssociationsBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputPublicFetchAssociationsBatchRequest(inputs []PublicFetchAssociationsBatchRequest) *BatchInputPublicFetchAssociationsBatchRequest {
	this := BatchInputPublicFetchAssociationsBatchRequest{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputPublicFetchAssociationsBatchRequestWithDefaults instantiates a new BatchInputPublicFetchAssociationsBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputPublicFetchAssociationsBatchRequestWithDefaults() *BatchInputPublicFetchAssociationsBatchRequest {
	this := BatchInputPublicFetchAssociationsBatchRequest{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputPublicFetchAssociationsBatchRequest) GetInputs() []PublicFetchAssociationsBatchRequest {
	if o == nil {
		var ret []PublicFetchAssociationsBatchRequest
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputPublicFetchAssociationsBatchRequest) GetInputsOk() ([]PublicFetchAssociationsBatchRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputPublicFetchAssociationsBatchRequest) SetInputs(v []PublicFetchAssociationsBatchRequest) {
	o.Inputs = v
}

func (o BatchInputPublicFetchAssociationsBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputPublicFetchAssociationsBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputPublicFetchAssociationsBatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inputs",
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

	varBatchInputPublicFetchAssociationsBatchRequest := _BatchInputPublicFetchAssociationsBatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputPublicFetchAssociationsBatchRequest)

	if err != nil {
		return err
	}

	*o = BatchInputPublicFetchAssociationsBatchRequest(varBatchInputPublicFetchAssociationsBatchRequest)

	return err
}

type NullableBatchInputPublicFetchAssociationsBatchRequest struct {
	value *BatchInputPublicFetchAssociationsBatchRequest
	isSet bool
}

func (v NullableBatchInputPublicFetchAssociationsBatchRequest) Get() *BatchInputPublicFetchAssociationsBatchRequest {
	return v.value
}

func (v *NullableBatchInputPublicFetchAssociationsBatchRequest) Set(val *BatchInputPublicFetchAssociationsBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputPublicFetchAssociationsBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputPublicFetchAssociationsBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputPublicFetchAssociationsBatchRequest(val *BatchInputPublicFetchAssociationsBatchRequest) *NullableBatchInputPublicFetchAssociationsBatchRequest {
	return &NullableBatchInputPublicFetchAssociationsBatchRequest{value: val, isSet: true}
}

func (v NullableBatchInputPublicFetchAssociationsBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputPublicFetchAssociationsBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
