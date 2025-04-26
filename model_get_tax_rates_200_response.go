package openapi

import (
	"encoding/json"
)

// checks if the GetTaxRates200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTaxRates200Response{}

// GetTaxRates200Response struct for GetTaxRates200Response
type GetTaxRates200Response struct {
	Meta Meta      `json:"meta,omitempty"`
	Data []TaxRate `json:"data,omitempty"`
}

// NewGetTaxRates200Response instantiates a new GetTaxRates200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaxRates200Response() *GetTaxRates200Response {
	this := GetTaxRates200Response{}
	return &this
}

// NewGetTaxRates200ResponseWithDefaults instantiates a new GetTaxRates200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaxRates200ResponseWithDefaults() *GetTaxRates200Response {
	this := GetTaxRates200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTaxRates200Response) GetData() []TaxRate {
	if o == nil || IsNil(o.Data) {
		var ret []TaxRate
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTaxRates200Response) GetDataOk() ([]TaxRate, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTaxRates200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TaxRate and assigns it to the Data field.
func (o *GetTaxRates200Response) SetData(v []TaxRate) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetTaxRates200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTaxRates200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetTaxRates200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetTaxRates200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetTaxRates200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTaxRates200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetTaxRates200Response struct {
	value *GetTaxRates200Response
	isSet bool
}

func (v NullableGetTaxRates200Response) Get() *GetTaxRates200Response {
	return v.value
}

func (v *NullableGetTaxRates200Response) Set(val *GetTaxRates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaxRates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaxRates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaxRates200Response(val *GetTaxRates200Response) *NullableGetTaxRates200Response {
	return &NullableGetTaxRates200Response{value: val, isSet: true}
}

func (v NullableGetTaxRates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaxRates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
