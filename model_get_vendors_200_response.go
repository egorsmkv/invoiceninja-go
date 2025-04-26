package openapi

import (
	"encoding/json"
)

// checks if the GetVendors200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVendors200Response{}

// GetVendors200Response struct for GetVendors200Response
type GetVendors200Response struct {
	Meta Meta     `json:"meta,omitempty"`
	Data []Vendor `json:"data,omitempty"`
}

// NewGetVendors200Response instantiates a new GetVendors200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVendors200Response() *GetVendors200Response {
	this := GetVendors200Response{}
	return &this
}

// NewGetVendors200ResponseWithDefaults instantiates a new GetVendors200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVendors200ResponseWithDefaults() *GetVendors200Response {
	this := GetVendors200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetVendors200Response) GetData() []Vendor {
	if o == nil || IsNil(o.Data) {
		var ret []Vendor
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVendors200Response) GetDataOk() ([]Vendor, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetVendors200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Vendor and assigns it to the Data field.
func (o *GetVendors200Response) SetData(v []Vendor) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetVendors200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVendors200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetVendors200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetVendors200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetVendors200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVendors200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetVendors200Response struct {
	value *GetVendors200Response
	isSet bool
}

func (v NullableGetVendors200Response) Get() *GetVendors200Response {
	return v.value
}

func (v *NullableGetVendors200Response) Set(val *GetVendors200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVendors200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVendors200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVendors200Response(val *GetVendors200Response) *NullableGetVendors200Response {
	return &NullableGetVendors200Response{value: val, isSet: true}
}

func (v NullableGetVendors200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVendors200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
