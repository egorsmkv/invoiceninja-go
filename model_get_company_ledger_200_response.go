package openapi

import (
	"encoding/json"
)

// checks if the GetCompanyLedger200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompanyLedger200Response{}

// GetCompanyLedger200Response struct for GetCompanyLedger200Response
type GetCompanyLedger200Response struct {
	Data []CompanyLedger `json:"data,omitempty"`
	Meta Meta            `json:"meta,omitempty"`
}

// NewGetCompanyLedger200Response instantiates a new GetCompanyLedger200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompanyLedger200Response() *GetCompanyLedger200Response {
	this := GetCompanyLedger200Response{}
	return &this
}

// NewGetCompanyLedger200ResponseWithDefaults instantiates a new GetCompanyLedger200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompanyLedger200ResponseWithDefaults() *GetCompanyLedger200Response {
	this := GetCompanyLedger200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCompanyLedger200Response) GetData() []CompanyLedger {
	if o == nil || IsNil(o.Data) {
		var ret []CompanyLedger
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompanyLedger200Response) GetDataOk() ([]CompanyLedger, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCompanyLedger200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CompanyLedger and assigns it to the Data field.
func (o *GetCompanyLedger200Response) SetData(v []CompanyLedger) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetCompanyLedger200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompanyLedger200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetCompanyLedger200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetCompanyLedger200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetCompanyLedger200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompanyLedger200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetCompanyLedger200Response struct {
	value *GetCompanyLedger200Response
	isSet bool
}

func (v NullableGetCompanyLedger200Response) Get() *GetCompanyLedger200Response {
	return v.value
}

func (v *NullableGetCompanyLedger200Response) Set(val *GetCompanyLedger200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompanyLedger200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompanyLedger200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompanyLedger200Response(val *GetCompanyLedger200Response) *NullableGetCompanyLedger200Response {
	return &NullableGetCompanyLedger200Response{value: val, isSet: true}
}

func (v NullableGetCompanyLedger200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompanyLedger200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
