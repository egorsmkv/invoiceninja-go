package openapi

import (
	"encoding/json"
)

// checks if the GetShowTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetShowTemplateRequest{}

// GetShowTemplateRequest struct for GetShowTemplateRequest
type GetShowTemplateRequest struct {
	// The email template subject
	Subject *string `json:"subject,omitempty"`
	// The email template body
	Body *string `json:"body,omitempty"`
}

// NewGetShowTemplateRequest instantiates a new GetShowTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetShowTemplateRequest() *GetShowTemplateRequest {
	this := GetShowTemplateRequest{}
	return &this
}

// NewGetShowTemplateRequestWithDefaults instantiates a new GetShowTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetShowTemplateRequestWithDefaults() *GetShowTemplateRequest {
	this := GetShowTemplateRequest{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *GetShowTemplateRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowTemplateRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *GetShowTemplateRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *GetShowTemplateRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *GetShowTemplateRequest) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetShowTemplateRequest) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *GetShowTemplateRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *GetShowTemplateRequest) SetBody(v string) {
	o.Body = &v
}

func (o GetShowTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetShowTemplateRequest) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableGetShowTemplateRequest struct {
	value *GetShowTemplateRequest
	isSet bool
}

func (v NullableGetShowTemplateRequest) Get() *GetShowTemplateRequest {
	return v.value
}

func (v *NullableGetShowTemplateRequest) Set(val *GetShowTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetShowTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetShowTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetShowTemplateRequest(val *GetShowTemplateRequest) *NullableGetShowTemplateRequest {
	return &NullableGetShowTemplateRequest{value: val, isSet: true}
}

func (v NullableGetShowTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetShowTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
