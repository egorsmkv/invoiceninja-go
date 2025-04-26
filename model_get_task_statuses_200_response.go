package openapi

import (
	"encoding/json"
)

// checks if the GetTaskStatuses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTaskStatuses200Response{}

// GetTaskStatuses200Response struct for GetTaskStatuses200Response
type GetTaskStatuses200Response struct {
	Data []TaskStatus `json:"data,omitempty"`
	Meta Meta         `json:"meta,omitempty"`
}

// NewGetTaskStatuses200Response instantiates a new GetTaskStatuses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaskStatuses200Response() *GetTaskStatuses200Response {
	this := GetTaskStatuses200Response{}
	return &this
}

// NewGetTaskStatuses200ResponseWithDefaults instantiates a new GetTaskStatuses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaskStatuses200ResponseWithDefaults() *GetTaskStatuses200Response {
	this := GetTaskStatuses200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTaskStatuses200Response) GetData() []TaskStatus {
	if o == nil || IsNil(o.Data) {
		var ret []TaskStatus
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTaskStatuses200Response) GetDataOk() ([]TaskStatus, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTaskStatuses200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TaskStatus and assigns it to the Data field.
func (o *GetTaskStatuses200Response) SetData(v []TaskStatus) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetTaskStatuses200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTaskStatuses200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetTaskStatuses200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetTaskStatuses200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetTaskStatuses200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTaskStatuses200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetTaskStatuses200Response struct {
	value *GetTaskStatuses200Response
	isSet bool
}

func (v NullableGetTaskStatuses200Response) Get() *GetTaskStatuses200Response {
	return v.value
}

func (v *NullableGetTaskStatuses200Response) Set(val *GetTaskStatuses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaskStatuses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaskStatuses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaskStatuses200Response(val *GetTaskStatuses200Response) *NullableGetTaskStatuses200Response {
	return &NullableGetTaskStatuses200Response{value: val, isSet: true}
}

func (v NullableGetTaskStatuses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaskStatuses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
