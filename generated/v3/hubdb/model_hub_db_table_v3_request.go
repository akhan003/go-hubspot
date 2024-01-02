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

// checks if the HubDbTableV3Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HubDbTableV3Request{}

// HubDbTableV3Request struct for HubDbTableV3Request
type HubDbTableV3Request struct {
	// Specifies the key value pairs of the metadata fields with the associated column ids
	DynamicMetaTags *map[string]int32 `json:"dynamicMetaTags,omitempty"`
	// Specifies whether the table can be read by public without authorization
	AllowPublicApiAccess *bool `json:"allowPublicApiAccess,omitempty"`
	// Specifies whether the table can be used for creation of dynamic pages
	UseForPages *bool `json:"useForPages,omitempty"`
	// List of columns in the table
	Columns []ColumnRequest `json:"columns,omitempty"`
	// Name of the table
	Name string `json:"name"`
	// Specifies creation of multi-level dynamic pages using child tables
	EnableChildTablePages *bool `json:"enableChildTablePages,omitempty"`
	// Label of the table
	Label string `json:"label"`
	// Specifies whether child tables can be created
	AllowChildTables *bool `json:"allowChildTables,omitempty"`
}

type _HubDbTableV3Request HubDbTableV3Request

// NewHubDbTableV3Request instantiates a new HubDbTableV3Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHubDbTableV3Request(name string, label string) *HubDbTableV3Request {
	this := HubDbTableV3Request{}
	this.Name = name
	this.Label = label
	return &this
}

// NewHubDbTableV3RequestWithDefaults instantiates a new HubDbTableV3Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHubDbTableV3RequestWithDefaults() *HubDbTableV3Request {
	this := HubDbTableV3Request{}
	return &this
}

// GetDynamicMetaTags returns the DynamicMetaTags field value if set, zero value otherwise.
func (o *HubDbTableV3Request) GetDynamicMetaTags() map[string]int32 {
	if o == nil || IsNil(o.DynamicMetaTags) {
		var ret map[string]int32
		return ret
	}
	return *o.DynamicMetaTags
}

// GetDynamicMetaTagsOk returns a tuple with the DynamicMetaTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetDynamicMetaTagsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.DynamicMetaTags) {
		return nil, false
	}
	return o.DynamicMetaTags, true
}

// HasDynamicMetaTags returns a boolean if a field has been set.
func (o *HubDbTableV3Request) HasDynamicMetaTags() bool {
	if o != nil && !IsNil(o.DynamicMetaTags) {
		return true
	}

	return false
}

// SetDynamicMetaTags gets a reference to the given map[string]int32 and assigns it to the DynamicMetaTags field.
func (o *HubDbTableV3Request) SetDynamicMetaTags(v map[string]int32) {
	o.DynamicMetaTags = &v
}

// GetAllowPublicApiAccess returns the AllowPublicApiAccess field value if set, zero value otherwise.
func (o *HubDbTableV3Request) GetAllowPublicApiAccess() bool {
	if o == nil || IsNil(o.AllowPublicApiAccess) {
		var ret bool
		return ret
	}
	return *o.AllowPublicApiAccess
}

// GetAllowPublicApiAccessOk returns a tuple with the AllowPublicApiAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetAllowPublicApiAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPublicApiAccess) {
		return nil, false
	}
	return o.AllowPublicApiAccess, true
}

// HasAllowPublicApiAccess returns a boolean if a field has been set.
func (o *HubDbTableV3Request) HasAllowPublicApiAccess() bool {
	if o != nil && !IsNil(o.AllowPublicApiAccess) {
		return true
	}

	return false
}

// SetAllowPublicApiAccess gets a reference to the given bool and assigns it to the AllowPublicApiAccess field.
func (o *HubDbTableV3Request) SetAllowPublicApiAccess(v bool) {
	o.AllowPublicApiAccess = &v
}

// GetUseForPages returns the UseForPages field value if set, zero value otherwise.
func (o *HubDbTableV3Request) GetUseForPages() bool {
	if o == nil || IsNil(o.UseForPages) {
		var ret bool
		return ret
	}
	return *o.UseForPages
}

// GetUseForPagesOk returns a tuple with the UseForPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetUseForPagesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseForPages) {
		return nil, false
	}
	return o.UseForPages, true
}

// HasUseForPages returns a boolean if a field has been set.
func (o *HubDbTableV3Request) HasUseForPages() bool {
	if o != nil && !IsNil(o.UseForPages) {
		return true
	}

	return false
}

// SetUseForPages gets a reference to the given bool and assigns it to the UseForPages field.
func (o *HubDbTableV3Request) SetUseForPages(v bool) {
	o.UseForPages = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *HubDbTableV3Request) GetColumns() []ColumnRequest {
	if o == nil || IsNil(o.Columns) {
		var ret []ColumnRequest
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetColumnsOk() ([]ColumnRequest, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *HubDbTableV3Request) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ColumnRequest and assigns it to the Columns field.
func (o *HubDbTableV3Request) SetColumns(v []ColumnRequest) {
	o.Columns = v
}

// GetName returns the Name field value
func (o *HubDbTableV3Request) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HubDbTableV3Request) SetName(v string) {
	o.Name = v
}

// GetEnableChildTablePages returns the EnableChildTablePages field value if set, zero value otherwise.
func (o *HubDbTableV3Request) GetEnableChildTablePages() bool {
	if o == nil || IsNil(o.EnableChildTablePages) {
		var ret bool
		return ret
	}
	return *o.EnableChildTablePages
}

// GetEnableChildTablePagesOk returns a tuple with the EnableChildTablePages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetEnableChildTablePagesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableChildTablePages) {
		return nil, false
	}
	return o.EnableChildTablePages, true
}

// HasEnableChildTablePages returns a boolean if a field has been set.
func (o *HubDbTableV3Request) HasEnableChildTablePages() bool {
	if o != nil && !IsNil(o.EnableChildTablePages) {
		return true
	}

	return false
}

// SetEnableChildTablePages gets a reference to the given bool and assigns it to the EnableChildTablePages field.
func (o *HubDbTableV3Request) SetEnableChildTablePages(v bool) {
	o.EnableChildTablePages = &v
}

// GetLabel returns the Label field value
func (o *HubDbTableV3Request) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *HubDbTableV3Request) SetLabel(v string) {
	o.Label = v
}

// GetAllowChildTables returns the AllowChildTables field value if set, zero value otherwise.
func (o *HubDbTableV3Request) GetAllowChildTables() bool {
	if o == nil || IsNil(o.AllowChildTables) {
		var ret bool
		return ret
	}
	return *o.AllowChildTables
}

// GetAllowChildTablesOk returns a tuple with the AllowChildTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HubDbTableV3Request) GetAllowChildTablesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChildTables) {
		return nil, false
	}
	return o.AllowChildTables, true
}

// HasAllowChildTables returns a boolean if a field has been set.
func (o *HubDbTableV3Request) HasAllowChildTables() bool {
	if o != nil && !IsNil(o.AllowChildTables) {
		return true
	}

	return false
}

// SetAllowChildTables gets a reference to the given bool and assigns it to the AllowChildTables field.
func (o *HubDbTableV3Request) SetAllowChildTables(v bool) {
	o.AllowChildTables = &v
}

func (o HubDbTableV3Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HubDbTableV3Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DynamicMetaTags) {
		toSerialize["dynamicMetaTags"] = o.DynamicMetaTags
	}
	if !IsNil(o.AllowPublicApiAccess) {
		toSerialize["allowPublicApiAccess"] = o.AllowPublicApiAccess
	}
	if !IsNil(o.UseForPages) {
		toSerialize["useForPages"] = o.UseForPages
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.EnableChildTablePages) {
		toSerialize["enableChildTablePages"] = o.EnableChildTablePages
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.AllowChildTables) {
		toSerialize["allowChildTables"] = o.AllowChildTables
	}
	return toSerialize, nil
}

func (o *HubDbTableV3Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"label",
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

	varHubDbTableV3Request := _HubDbTableV3Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHubDbTableV3Request)

	if err != nil {
		return err
	}

	*o = HubDbTableV3Request(varHubDbTableV3Request)

	return err
}

type NullableHubDbTableV3Request struct {
	value *HubDbTableV3Request
	isSet bool
}

func (v NullableHubDbTableV3Request) Get() *HubDbTableV3Request {
	return v.value
}

func (v *NullableHubDbTableV3Request) Set(val *HubDbTableV3Request) {
	v.value = val
	v.isSet = true
}

func (v NullableHubDbTableV3Request) IsSet() bool {
	return v.isSet
}

func (v *NullableHubDbTableV3Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHubDbTableV3Request(val *HubDbTableV3Request) *NullableHubDbTableV3Request {
	return &NullableHubDbTableV3Request{value: val, isSet: true}
}

func (v NullableHubDbTableV3Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHubDbTableV3Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
