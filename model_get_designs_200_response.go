package openapi

import (
	"encoding/json"
)

// checks if the GetDesigns200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDesigns200Response{}

// GetDesigns200Response struct for GetDesigns200Response
type GetDesigns200Response struct {
	Data []Design `json:"data,omitempty"`
	Meta Meta     `json:"meta,omitempty"`
}

// NewGetDesigns200Response instantiates a new GetDesigns200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDesigns200Response() *GetDesigns200Response {
	this := GetDesigns200Response{}
	return &this
}

// NewGetDesigns200ResponseWithDefaults instantiates a new GetDesigns200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDesigns200ResponseWithDefaults() *GetDesigns200Response {
	this := GetDesigns200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDesigns200Response) GetData() []Design {
	if o == nil || IsNil(o.Data) {
		var ret []Design
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDesigns200Response) GetDataOk() ([]Design, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDesigns200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Design and assigns it to the Data field.
func (o *GetDesigns200Response) SetData(v []Design) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetDesigns200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDesigns200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetDesigns200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetDesigns200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetDesigns200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDesigns200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetDesigns200Response struct {
	value *GetDesigns200Response
	isSet bool
}

func (v NullableGetDesigns200Response) Get() *GetDesigns200Response {
	return v.value
}

func (v *NullableGetDesigns200Response) Set(val *GetDesigns200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDesigns200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDesigns200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDesigns200Response(val *GetDesigns200Response) *NullableGetDesigns200Response {
	return &NullableGetDesigns200Response{value: val, isSet: true}
}

func (v NullableGetDesigns200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDesigns200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
