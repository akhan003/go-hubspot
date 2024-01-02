/*
Hubdb

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CollectionResponseWithTotalHubDbTableV3ForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseWithTotalHubDbTableV3ForwardPaging{}

// CollectionResponseWithTotalHubDbTableV3ForwardPaging struct for CollectionResponseWithTotalHubDbTableV3ForwardPaging
type CollectionResponseWithTotalHubDbTableV3ForwardPaging struct {
	//
	Total  int32          `json:"total"`
	Paging *ForwardPaging `json:"paging,omitempty"`
	//
	Results []HubDbTableV3 `json:"results"`
}

type _CollectionResponseWithTotalHubDbTableV3ForwardPaging CollectionResponseWithTotalHubDbTableV3ForwardPaging

// NewCollectionResponseWithTotalHubDbTableV3ForwardPaging instantiates a new CollectionResponseWithTotalHubDbTableV3ForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseWithTotalHubDbTableV3ForwardPaging(total int32, results []HubDbTableV3) *CollectionResponseWithTotalHubDbTableV3ForwardPaging {
	this := CollectionResponseWithTotalHubDbTableV3ForwardPaging{}
	this.Total = total
	this.Results = results
	return &this
}

// NewCollectionResponseWithTotalHubDbTableV3ForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalHubDbTableV3ForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseWithTotalHubDbTableV3ForwardPagingWithDefaults() *CollectionResponseWithTotalHubDbTableV3ForwardPaging {
	this := CollectionResponseWithTotalHubDbTableV3ForwardPaging{}
	return &this
}

// GetTotal returns the Total field value
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) SetTotal(v int32) {
	o.Total = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetPaging() ForwardPaging {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

// GetResults returns the Results field value
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetResults() []HubDbTableV3 {
	if o == nil {
		var ret []HubDbTableV3
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) GetResultsOk() ([]HubDbTableV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) SetResults(v []HubDbTableV3) {
	o.Results = v
}

func (o CollectionResponseWithTotalHubDbTableV3ForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseWithTotalHubDbTableV3ForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *CollectionResponseWithTotalHubDbTableV3ForwardPaging) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"results",
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

	varCollectionResponseWithTotalHubDbTableV3ForwardPaging := _CollectionResponseWithTotalHubDbTableV3ForwardPaging{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionResponseWithTotalHubDbTableV3ForwardPaging)

	if err != nil {
		return err
	}

	*o = CollectionResponseWithTotalHubDbTableV3ForwardPaging(varCollectionResponseWithTotalHubDbTableV3ForwardPaging)

	return err
}

type NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging struct {
	value *CollectionResponseWithTotalHubDbTableV3ForwardPaging
	isSet bool
}

func (v NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging) Get() *CollectionResponseWithTotalHubDbTableV3ForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging) Set(val *CollectionResponseWithTotalHubDbTableV3ForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseWithTotalHubDbTableV3ForwardPaging(val *CollectionResponseWithTotalHubDbTableV3ForwardPaging) *NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging {
	return &NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseWithTotalHubDbTableV3ForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
