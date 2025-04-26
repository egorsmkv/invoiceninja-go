package openapi

import (
	"encoding/json"
)

// checks if the GetBankIntegrations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBankIntegrations200Response{}

// GetBankIntegrations200Response struct for GetBankIntegrations200Response
type GetBankIntegrations200Response struct {
	Data []BankIntegration `json:"data,omitempty"`
	Meta Meta              `json:"meta,omitempty"`
}

// NewGetBankIntegrations200Response instantiates a new GetBankIntegrations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBankIntegrations200Response() *GetBankIntegrations200Response {
	this := GetBankIntegrations200Response{}
	return &this
}

// NewGetBankIntegrations200ResponseWithDefaults instantiates a new GetBankIntegrations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBankIntegrations200ResponseWithDefaults() *GetBankIntegrations200Response {
	this := GetBankIntegrations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetBankIntegrations200Response) GetData() []BankIntegration {
	if o == nil || IsNil(o.Data) {
		var ret []BankIntegration
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBankIntegrations200Response) GetDataOk() ([]BankIntegration, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetBankIntegrations200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BankIntegration and assigns it to the Data field.
func (o *GetBankIntegrations200Response) SetData(v []BankIntegration) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetBankIntegrations200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBankIntegrations200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetBankIntegrations200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetBankIntegrations200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetBankIntegrations200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBankIntegrations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetBankIntegrations200Response struct {
	value *GetBankIntegrations200Response
	isSet bool
}

func (v NullableGetBankIntegrations200Response) Get() *GetBankIntegrations200Response {
	return v.value
}

func (v *NullableGetBankIntegrations200Response) Set(val *GetBankIntegrations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBankIntegrations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBankIntegrations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBankIntegrations200Response(val *GetBankIntegrations200Response) *NullableGetBankIntegrations200Response {
	return &NullableGetBankIntegrations200Response{value: val, isSet: true}
}

func (v NullableGetBankIntegrations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBankIntegrations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
