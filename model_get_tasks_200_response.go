package openapi

import (
	"encoding/json"
)

// checks if the GetTasks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTasks200Response{}

// GetTasks200Response struct for GetTasks200Response
type GetTasks200Response struct {
	Data []Task `json:"data,omitempty"`
	Meta Meta   `json:"meta,omitempty"`
}

// NewGetTasks200Response instantiates a new GetTasks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTasks200Response() *GetTasks200Response {
	this := GetTasks200Response{}
	return &this
}

// NewGetTasks200ResponseWithDefaults instantiates a new GetTasks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTasks200ResponseWithDefaults() *GetTasks200Response {
	this := GetTasks200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTasks200Response) GetData() []Task {
	if o == nil || IsNil(o.Data) {
		var ret []Task
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTasks200Response) GetDataOk() ([]Task, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTasks200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Task and assigns it to the Data field.
func (o *GetTasks200Response) SetData(v []Task) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetTasks200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTasks200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetTasks200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetTasks200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetTasks200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTasks200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetTasks200Response struct {
	value *GetTasks200Response
	isSet bool
}

func (v NullableGetTasks200Response) Get() *GetTasks200Response {
	return v.value
}

func (v *NullableGetTasks200Response) Set(val *GetTasks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTasks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTasks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTasks200Response(val *GetTasks200Response) *NullableGetTasks200Response {
	return &NullableGetTasks200Response{value: val, isSet: true}
}

func (v NullableGetTasks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTasks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
