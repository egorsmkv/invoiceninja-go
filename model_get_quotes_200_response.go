package openapi

import (
	"encoding/json"
)

// checks if the GetQuotes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuotes200Response{}

// GetQuotes200Response struct for GetQuotes200Response
type GetQuotes200Response struct {
	Meta Meta    `json:"meta,omitempty"`
	Data []Quote `json:"data,omitempty"`
}

// NewGetQuotes200Response instantiates a new GetQuotes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuotes200Response() *GetQuotes200Response {
	this := GetQuotes200Response{}
	return &this
}

// NewGetQuotes200ResponseWithDefaults instantiates a new GetQuotes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuotes200ResponseWithDefaults() *GetQuotes200Response {
	this := GetQuotes200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetQuotes200Response) GetData() []Quote {
	if o == nil || IsNil(o.Data) {
		var ret []Quote
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuotes200Response) GetDataOk() ([]Quote, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetQuotes200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Quote and assigns it to the Data field.
func (o *GetQuotes200Response) SetData(v []Quote) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetQuotes200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuotes200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetQuotes200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetQuotes200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetQuotes200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuotes200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetQuotes200Response struct {
	value *GetQuotes200Response
	isSet bool
}

func (v NullableGetQuotes200Response) Get() *GetQuotes200Response {
	return v.value
}

func (v *NullableGetQuotes200Response) Set(val *GetQuotes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuotes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuotes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuotes200Response(val *GetQuotes200Response) *NullableGetQuotes200Response {
	return &NullableGetQuotes200Response{value: val, isSet: true}
}

func (v NullableGetQuotes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuotes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
