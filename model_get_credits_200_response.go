package openapi

import (
	"encoding/json"
)

// checks if the GetCredits200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCredits200Response{}

// GetCredits200Response struct for GetCredits200Response
type GetCredits200Response struct {
	Data []Credit `json:"data,omitempty"`
	Meta Meta     `json:"meta,omitempty"`
}

// NewGetCredits200Response instantiates a new GetCredits200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCredits200Response() *GetCredits200Response {
	this := GetCredits200Response{}
	return &this
}

// NewGetCredits200ResponseWithDefaults instantiates a new GetCredits200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCredits200ResponseWithDefaults() *GetCredits200Response {
	this := GetCredits200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCredits200Response) GetData() []Credit {
	if o == nil || IsNil(o.Data) {
		var ret []Credit
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCredits200Response) GetDataOk() ([]Credit, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCredits200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Credit and assigns it to the Data field.
func (o *GetCredits200Response) SetData(v []Credit) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetCredits200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCredits200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetCredits200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetCredits200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetCredits200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCredits200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetCredits200Response struct {
	value *GetCredits200Response
	isSet bool
}

func (v NullableGetCredits200Response) Get() *GetCredits200Response {
	return v.value
}

func (v *NullableGetCredits200Response) Set(val *GetCredits200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCredits200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCredits200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCredits200Response(val *GetCredits200Response) *NullableGetCredits200Response {
	return &NullableGetCredits200Response{value: val, isSet: true}
}

func (v NullableGetCredits200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCredits200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
