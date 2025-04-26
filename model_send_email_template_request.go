package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SendEmailTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendEmailTemplateRequest{}

// SendEmailTemplateRequest struct for SendEmailTemplateRequest
type SendEmailTemplateRequest struct {
	// The email subject
	Subject *string `json:"subject,omitempty"`
	// The email body
	Body *string `json:"body,omitempty"`
	// The entity name
	Entity string `json:"entity"`
	// The entity_id
	EntityId string `json:"entity_id"`
	// The email address of a user to be CC'd on the email
	CcEmail *string `json:"cc_email,omitempty"`
	// The template required
	Template string `json:"template"`
}

type _SendEmailTemplateRequest SendEmailTemplateRequest

// NewSendEmailTemplateRequest instantiates a new SendEmailTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendEmailTemplateRequest(entity string, entityId string, template string) *SendEmailTemplateRequest {
	this := SendEmailTemplateRequest{}
	this.Entity = entity
	this.EntityId = entityId
	this.Template = template
	return &this
}

// NewSendEmailTemplateRequestWithDefaults instantiates a new SendEmailTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendEmailTemplateRequestWithDefaults() *SendEmailTemplateRequest {
	this := SendEmailTemplateRequest{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SendEmailTemplateRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendEmailTemplateRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SendEmailTemplateRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SendEmailTemplateRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SendEmailTemplateRequest) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendEmailTemplateRequest) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SendEmailTemplateRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SendEmailTemplateRequest) SetBody(v string) {
	o.Body = &v
}

// GetEntity returns the Entity field value
func (o *SendEmailTemplateRequest) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *SendEmailTemplateRequest) GetEntityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *SendEmailTemplateRequest) SetEntity(v string) {
	o.Entity = v
}

// GetEntityId returns the EntityId field value
func (o *SendEmailTemplateRequest) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *SendEmailTemplateRequest) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *SendEmailTemplateRequest) SetEntityId(v string) {
	o.EntityId = v
}

// GetCcEmail returns the CcEmail field value if set, zero value otherwise.
func (o *SendEmailTemplateRequest) GetCcEmail() string {
	if o == nil || IsNil(o.CcEmail) {
		var ret string
		return ret
	}
	return *o.CcEmail
}

// GetCcEmailOk returns a tuple with the CcEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendEmailTemplateRequest) GetCcEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CcEmail) {
		return nil, false
	}
	return o.CcEmail, true
}

// HasCcEmail returns a boolean if a field has been set.
func (o *SendEmailTemplateRequest) HasCcEmail() bool {
	if o != nil && !IsNil(o.CcEmail) {
		return true
	}

	return false
}

// SetCcEmail gets a reference to the given string and assigns it to the CcEmail field.
func (o *SendEmailTemplateRequest) SetCcEmail(v string) {
	o.CcEmail = &v
}

// GetTemplate returns the Template field value
func (o *SendEmailTemplateRequest) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *SendEmailTemplateRequest) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *SendEmailTemplateRequest) SetTemplate(v string) {
	o.Template = v
}

func (o SendEmailTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendEmailTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	toSerialize["entity"] = o.Entity
	toSerialize["entity_id"] = o.EntityId
	if !IsNil(o.CcEmail) {
		toSerialize["cc_email"] = o.CcEmail
	}
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *SendEmailTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity",
		"entity_id",
		"template",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)
	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSendEmailTemplateRequest := _SendEmailTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendEmailTemplateRequest)
	if err != nil {
		return err
	}

	*o = SendEmailTemplateRequest(varSendEmailTemplateRequest)

	return err
}

type NullableSendEmailTemplateRequest struct {
	value *SendEmailTemplateRequest
	isSet bool
}

func (v NullableSendEmailTemplateRequest) Get() *SendEmailTemplateRequest {
	return v.value
}

func (v *NullableSendEmailTemplateRequest) Set(val *SendEmailTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendEmailTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendEmailTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendEmailTemplateRequest(val *SendEmailTemplateRequest) *NullableSendEmailTemplateRequest {
	return &NullableSendEmailTemplateRequest{value: val, isSet: true}
}

func (v NullableSendEmailTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendEmailTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
