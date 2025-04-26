package openapi

import (
	"encoding/json"
)

// checks if the GetCompanies200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompanies200Response{}

// GetCompanies200Response struct for GetCompanies200Response
type GetCompanies200Response struct {
	Data []Company `json:"data,omitempty"`
	Meta Meta      `json:"meta,omitempty"`
}

// NewGetCompanies200Response instantiates a new GetCompanies200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompanies200Response() *GetCompanies200Response {
	this := GetCompanies200Response{}
	return &this
}

// NewGetCompanies200ResponseWithDefaults instantiates a new GetCompanies200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompanies200ResponseWithDefaults() *GetCompanies200Response {
	this := GetCompanies200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCompanies200Response) GetData() []Company {
	if o == nil || IsNil(o.Data) {
		var ret []Company
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompanies200Response) GetDataOk() ([]Company, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCompanies200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Company and assigns it to the Data field.
func (o *GetCompanies200Response) SetData(v []Company) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetCompanies200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompanies200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetCompanies200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetCompanies200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetCompanies200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompanies200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetCompanies200Response struct {
	value *GetCompanies200Response
	isSet bool
}

func (v NullableGetCompanies200Response) Get() *GetCompanies200Response {
	return v.value
}

func (v *NullableGetCompanies200Response) Set(val *GetCompanies200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompanies200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompanies200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompanies200Response(val *GetCompanies200Response) *NullableGetCompanies200Response {
	return &NullableGetCompanies200Response{value: val, isSet: true}
}

func (v NullableGetCompanies200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompanies200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
