package openapi

import (
	"encoding/json"
)

// checks if the GetProjects200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjects200Response{}

// GetProjects200Response struct for GetProjects200Response
type GetProjects200Response struct {
	Data []Project `json:"data,omitempty"`
	Meta Meta      `json:"meta,omitempty"`
}

// NewGetProjects200Response instantiates a new GetProjects200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjects200Response() *GetProjects200Response {
	this := GetProjects200Response{}
	return &this
}

// NewGetProjects200ResponseWithDefaults instantiates a new GetProjects200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjects200ResponseWithDefaults() *GetProjects200Response {
	this := GetProjects200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetProjects200Response) GetData() []Project {
	if o == nil || IsNil(o.Data) {
		var ret []Project
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjects200Response) GetDataOk() ([]Project, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetProjects200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Project and assigns it to the Data field.
func (o *GetProjects200Response) SetData(v []Project) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetProjects200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjects200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetProjects200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetProjects200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetProjects200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjects200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetProjects200Response struct {
	value *GetProjects200Response
	isSet bool
}

func (v NullableGetProjects200Response) Get() *GetProjects200Response {
	return v.value
}

func (v *NullableGetProjects200Response) Set(val *GetProjects200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjects200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjects200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjects200Response(val *GetProjects200Response) *NullableGetProjects200Response {
	return &NullableGetProjects200Response{value: val, isSet: true}
}

func (v NullableGetProjects200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjects200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
