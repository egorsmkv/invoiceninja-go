/*
Invoice Ninja API Reference.

---   ![Invoice Ninja](https://invoicing.co/images/new_logo.png)   ## Introduction   Welcome to the Invoice Ninja API documentation, your comprehensive guide to integrating Invoice Ninja's powerful features into your applications. Whether you're building a custom client, automating workflows, or integrating with other systems, our API provides the tools you need to streamline your invoicing and billing processes.   ### What is Invoice Ninja?   Invoice Ninja is a robust source-available platform designed to simplify invoicing, billing, and payment management for freelancers, small businesses, and enterprises alike. With a user-friendly interface, customizable templates, and a suite of powerful features, Invoice Ninja empowers businesses to create professional invoices, track expenses, manage clients, and get paid faster.   ### Why use the Invoice Ninja API?   The Invoice Ninja API allows developers to extend the functionality of Invoice Ninja by programmatically accessing and manipulating data within their Invoice Ninja accounts. With the API, you can automate repetitive tasks, integrate with third-party services, and build custom solutions tailored to your specific business needs.   ### Getting Started   To get started with the Invoice Ninja API, you'll need an active Invoice Ninja account (or your own self hosted installation) and API credentials. If you haven't already done so, sign up for an account at Invoice Ninja and generate your API keys from the settings section.    Once you have your API credentials, you can start exploring the API endpoints, authentication methods, request and response formats, and more using the documentation provided here.   ### Explore the Documentation     This documentation is organized into sections to help you navigate and understand the various aspects of the Invoice Ninja API:    - Authentication: Learn how to authenticate your requests to the API using API tokens.   - Endpoints: Explore the available API endpoints for managing invoices, clients, payments, expenses, and more.   - Request and Response Formats: Understand the structure of API requests and responses, including parameters, headers, and payloads.   - Error Handling: Learn about error codes, status messages, and best practices for handling errors gracefully.   - Code Examples: Find code examples and tutorials to help you get started with integrating the Invoice Ninja API into your applications.         ### Need Help?         If you have any questions, encounter any issues, or need assistance with using the Invoice Ninja API, don't hesitate to reach out to our support team or join our community forums. We're here to help you succeed with Invoice Ninja and make the most of our API.        Let's start building together!   ### Endpoints      <div style=\"background-color: #2D394E; color: #fff padding: 20px; border-radius: 5px; border: 4px solid #212A3B; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);\">       <p style=\"padding:10px; color: #DBDBDB;\"\">Production: https://invoicing.co</p>       <p style=\"padding:10px; color: #DBDBDB;\">Demo: https://demo.invoiceninja.com</p>   </div>    ### Client Libraries    PHP SDK can be found [here](https://github.com/invoiceninja/sdk-php)   ### Authentication:    Invoice Ninja uses API tokens to authenticate requests. You can view and manage your API keys in Settings > Account Management > Integrations > API tokens    API requests must be made over HTTPS. Calls made to HTTP will fail.   ### Errors:    Invoice Ninja uses standard HTTP response codes to indicate the success or failure of a request. below is a table of standard status codes and responses    | Status Code | Explanation                                                                 |   |-------------|-----------------------------------------------------------------------------|   | 200         | OK: The request has succeeded. The information returned with the response is dependent on the method used in the request. |   | 301         | Moved Permanently: This and all future requests should be directed to the given URI. |   | 303         | See Other: The response to the request can be found under another URI using the GET method. |   | 400         | Bad Request: The server cannot or will not process the request due to an apparent client error. |   | 401         | Unauthorized: Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided. |   | 403         | Forbidden: The request was valid, but the server is refusing action. |   | 404         | Not Found: The requested resource could not be found but may be available in the future. |   | 405         | Method Not Allowed: A request method is not supported for the requested resource. |   | 409         | Conflict: Indicates that the request could not be processed because of conflict in the request. |   | 422         | Unprocessable Entity: The request was well-formed but was unable to be followed due to semantic errors. |   | 429         | Too Many Requests: The user has sent too many requests in a given amount of time (\"rate limiting\"). |   | 500         | Internal Server Error: A generic error message, given when an unexpected condition was encountered and no more specific message is suitable. |   ### Pagination    When using index routes to retrieve lists of data, by default we limit the number of records returned to 20. You can using standard pagination to paginate results, ie: ?per_page=50 

API version: 5.11.48
Contact: contact@invoiceninja.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
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
	toSerialize,err := o.ToMap()
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
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
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


