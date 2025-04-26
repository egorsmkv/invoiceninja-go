package openapi

import (
	"encoding/json"
)

// checks if the GetWebhooks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhooks200Response{}

// GetWebhooks200Response struct for GetWebhooks200Response
type GetWebhooks200Response struct {
	Data []Webhook `json:"data,omitempty"`
	Meta Meta      `json:"meta,omitempty"`
}

// NewGetWebhooks200Response instantiates a new GetWebhooks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhooks200Response() *GetWebhooks200Response {
	this := GetWebhooks200Response{}
	return &this
}

// NewGetWebhooks200ResponseWithDefaults instantiates a new GetWebhooks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhooks200ResponseWithDefaults() *GetWebhooks200Response {
	this := GetWebhooks200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetWebhooks200Response) GetData() []Webhook {
	if o == nil || IsNil(o.Data) {
		var ret []Webhook
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhooks200Response) GetDataOk() ([]Webhook, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetWebhooks200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Webhook and assigns it to the Data field.
func (o *GetWebhooks200Response) SetData(v []Webhook) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetWebhooks200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhooks200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetWebhooks200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetWebhooks200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetWebhooks200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhooks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetWebhooks200Response struct {
	value *GetWebhooks200Response
	isSet bool
}

func (v NullableGetWebhooks200Response) Get() *GetWebhooks200Response {
	return v.value
}

func (v *NullableGetWebhooks200Response) Set(val *GetWebhooks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhooks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhooks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhooks200Response(val *GetWebhooks200Response) *NullableGetWebhooks200Response {
	return &NullableGetWebhooks200Response{value: val, isSet: true}
}

func (v NullableGetWebhooks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhooks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
