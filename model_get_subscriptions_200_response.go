package openapi

import (
	"encoding/json"
)

// checks if the GetSubscriptions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubscriptions200Response{}

// GetSubscriptions200Response struct for GetSubscriptions200Response
type GetSubscriptions200Response struct {
	Meta Meta           `json:"meta,omitempty"`
	Data []Subscription `json:"data,omitempty"`
}

// NewGetSubscriptions200Response instantiates a new GetSubscriptions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubscriptions200Response() *GetSubscriptions200Response {
	this := GetSubscriptions200Response{}
	return &this
}

// NewGetSubscriptions200ResponseWithDefaults instantiates a new GetSubscriptions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubscriptions200ResponseWithDefaults() *GetSubscriptions200Response {
	this := GetSubscriptions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSubscriptions200Response) GetData() []Subscription {
	if o == nil || IsNil(o.Data) {
		var ret []Subscription
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubscriptions200Response) GetDataOk() ([]Subscription, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSubscriptions200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Subscription and assigns it to the Data field.
func (o *GetSubscriptions200Response) SetData(v []Subscription) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetSubscriptions200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubscriptions200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetSubscriptions200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetSubscriptions200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetSubscriptions200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubscriptions200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetSubscriptions200Response struct {
	value *GetSubscriptions200Response
	isSet bool
}

func (v NullableGetSubscriptions200Response) Get() *GetSubscriptions200Response {
	return v.value
}

func (v *NullableGetSubscriptions200Response) Set(val *GetSubscriptions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubscriptions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubscriptions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubscriptions200Response(val *GetSubscriptions200Response) *NullableGetSubscriptions200Response {
	return &NullableGetSubscriptions200Response{value: val, isSet: true}
}

func (v NullableGetSubscriptions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubscriptions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
