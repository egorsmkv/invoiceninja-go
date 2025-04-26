package openapi

import (
	"encoding/json"
)

// checks if the GetTokens200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTokens200Response{}

// GetTokens200Response struct for GetTokens200Response
type GetTokens200Response struct {
	Data []CompanyToken `json:"data,omitempty"`
	Meta Meta           `json:"meta,omitempty"`
}

// NewGetTokens200Response instantiates a new GetTokens200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTokens200Response() *GetTokens200Response {
	this := GetTokens200Response{}
	return &this
}

// NewGetTokens200ResponseWithDefaults instantiates a new GetTokens200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTokens200ResponseWithDefaults() *GetTokens200Response {
	this := GetTokens200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTokens200Response) GetData() []CompanyToken {
	if o == nil || IsNil(o.Data) {
		var ret []CompanyToken
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTokens200Response) GetDataOk() ([]CompanyToken, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTokens200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CompanyToken and assigns it to the Data field.
func (o *GetTokens200Response) SetData(v []CompanyToken) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetTokens200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTokens200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetTokens200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetTokens200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetTokens200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTokens200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetTokens200Response struct {
	value *GetTokens200Response
	isSet bool
}

func (v NullableGetTokens200Response) Get() *GetTokens200Response {
	return v.value
}

func (v *NullableGetTokens200Response) Set(val *GetTokens200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTokens200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTokens200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTokens200Response(val *GetTokens200Response) *NullableGetTokens200Response {
	return &NullableGetTokens200Response{value: val, isSet: true}
}

func (v NullableGetTokens200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTokens200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
