package openapi

import (
	"encoding/json"
)

// checks if the GetCompanyGateways200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompanyGateways200Response{}

// GetCompanyGateways200Response struct for GetCompanyGateways200Response
type GetCompanyGateways200Response struct {
	Data []CompanyGateway `json:"data,omitempty"`
	Meta Meta             `json:"meta,omitempty"`
}

// NewGetCompanyGateways200Response instantiates a new GetCompanyGateways200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompanyGateways200Response() *GetCompanyGateways200Response {
	this := GetCompanyGateways200Response{}
	return &this
}

// NewGetCompanyGateways200ResponseWithDefaults instantiates a new GetCompanyGateways200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompanyGateways200ResponseWithDefaults() *GetCompanyGateways200Response {
	this := GetCompanyGateways200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCompanyGateways200Response) GetData() []CompanyGateway {
	if o == nil || IsNil(o.Data) {
		var ret []CompanyGateway
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompanyGateways200Response) GetDataOk() ([]CompanyGateway, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCompanyGateways200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CompanyGateway and assigns it to the Data field.
func (o *GetCompanyGateways200Response) SetData(v []CompanyGateway) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetCompanyGateways200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompanyGateways200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetCompanyGateways200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetCompanyGateways200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetCompanyGateways200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompanyGateways200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetCompanyGateways200Response struct {
	value *GetCompanyGateways200Response
	isSet bool
}

func (v NullableGetCompanyGateways200Response) Get() *GetCompanyGateways200Response {
	return v.value
}

func (v *NullableGetCompanyGateways200Response) Set(val *GetCompanyGateways200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompanyGateways200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompanyGateways200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompanyGateways200Response(val *GetCompanyGateways200Response) *NullableGetCompanyGateways200Response {
	return &NullableGetCompanyGateways200Response{value: val, isSet: true}
}

func (v NullableGetCompanyGateways200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompanyGateways200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
