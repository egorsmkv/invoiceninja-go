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
)

// checks if the ClientContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientContact{}

// ClientContact struct for ClientContact
type ClientContact struct {
	// The hashed if of the contact
	Id *string `json:"id,omitempty"`
	// The hashed id of the user who created the contact
	UserId *string `json:"user_id,omitempty"`
	// The hashed id of the client
	ClientId *string `json:"client_id,omitempty"`
	// The first name of the contact
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the contact
	LastName *string `json:"last_name,omitempty"`
	// The phone number of the contact
	Phone *string `json:"phone,omitempty"`
	// A Custom field value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A Custom field value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A Custom field value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A Custom field value
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The email of the contact
	Email *string `json:"email,omitempty"`
	// The terms of service which the contact has accpeted
	AcceptedTermsVersion *string `json:"accepted_terms_version,omitempty"`
	// The hashed password of the contact
	Password *string `json:"password,omitempty"`
	// The confirmation code used to authenticate the contacts email address
	ConfirmationCode *string `json:"confirmation_code,omitempty"`
	// A uuid based token.
	Token *string `json:"token,omitempty"`
	// A unique identifier for the contact
	ContactKey *string `json:"contact_key,omitempty"`
	// Defines is this contact is the primary contact for the client
	IsPrimary *bool `json:"is_primary,omitempty"`
	// Boolean value confirms the user has confirmed their account.
	Confirmed *bool `json:"confirmed,omitempty"`
	// Boolean value defines if the contact has been locked out.
	IsLocked *bool `json:"is_locked,omitempty"`
	// Boolean value determines is this contact should receive emails
	SendEmail *bool `json:"send_email,omitempty"`
	// The number of failed logins the contact has had
	FailedLogins *float32 `json:"failed_logins,omitempty"`
	// The date which the contact confirmed their email
	EmailVerifiedAt *float32 `json:"email_verified_at,omitempty"`
	// Timestamp
	LastLogin *float32 `json:"last_login,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	DeletedAt *float32 `json:"deleted_at,omitempty"`
}

// NewClientContact instantiates a new ClientContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientContact() *ClientContact {
	this := ClientContact{}
	return &this
}

// NewClientContactWithDefaults instantiates a new ClientContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientContactWithDefaults() *ClientContact {
	this := ClientContact{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientContact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientContact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientContact) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ClientContact) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ClientContact) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ClientContact) SetUserId(v string) {
	o.UserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientContact) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientContact) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientContact) SetClientId(v string) {
	o.ClientId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ClientContact) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ClientContact) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ClientContact) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ClientContact) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ClientContact) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ClientContact) SetLastName(v string) {
	o.LastName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ClientContact) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ClientContact) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ClientContact) SetPhone(v string) {
	o.Phone = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *ClientContact) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *ClientContact) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *ClientContact) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *ClientContact) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *ClientContact) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *ClientContact) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *ClientContact) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *ClientContact) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *ClientContact) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *ClientContact) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *ClientContact) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *ClientContact) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ClientContact) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ClientContact) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ClientContact) SetEmail(v string) {
	o.Email = &v
}

// GetAcceptedTermsVersion returns the AcceptedTermsVersion field value if set, zero value otherwise.
func (o *ClientContact) GetAcceptedTermsVersion() string {
	if o == nil || IsNil(o.AcceptedTermsVersion) {
		var ret string
		return ret
	}
	return *o.AcceptedTermsVersion
}

// GetAcceptedTermsVersionOk returns a tuple with the AcceptedTermsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetAcceptedTermsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptedTermsVersion) {
		return nil, false
	}
	return o.AcceptedTermsVersion, true
}

// HasAcceptedTermsVersion returns a boolean if a field has been set.
func (o *ClientContact) HasAcceptedTermsVersion() bool {
	if o != nil && !IsNil(o.AcceptedTermsVersion) {
		return true
	}

	return false
}

// SetAcceptedTermsVersion gets a reference to the given string and assigns it to the AcceptedTermsVersion field.
func (o *ClientContact) SetAcceptedTermsVersion(v string) {
	o.AcceptedTermsVersion = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ClientContact) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ClientContact) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ClientContact) SetPassword(v string) {
	o.Password = &v
}

// GetConfirmationCode returns the ConfirmationCode field value if set, zero value otherwise.
func (o *ClientContact) GetConfirmationCode() string {
	if o == nil || IsNil(o.ConfirmationCode) {
		var ret string
		return ret
	}
	return *o.ConfirmationCode
}

// GetConfirmationCodeOk returns a tuple with the ConfirmationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetConfirmationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmationCode) {
		return nil, false
	}
	return o.ConfirmationCode, true
}

// HasConfirmationCode returns a boolean if a field has been set.
func (o *ClientContact) HasConfirmationCode() bool {
	if o != nil && !IsNil(o.ConfirmationCode) {
		return true
	}

	return false
}

// SetConfirmationCode gets a reference to the given string and assigns it to the ConfirmationCode field.
func (o *ClientContact) SetConfirmationCode(v string) {
	o.ConfirmationCode = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ClientContact) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ClientContact) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ClientContact) SetToken(v string) {
	o.Token = &v
}

// GetContactKey returns the ContactKey field value if set, zero value otherwise.
func (o *ClientContact) GetContactKey() string {
	if o == nil || IsNil(o.ContactKey) {
		var ret string
		return ret
	}
	return *o.ContactKey
}

// GetContactKeyOk returns a tuple with the ContactKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetContactKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ContactKey) {
		return nil, false
	}
	return o.ContactKey, true
}

// HasContactKey returns a boolean if a field has been set.
func (o *ClientContact) HasContactKey() bool {
	if o != nil && !IsNil(o.ContactKey) {
		return true
	}

	return false
}

// SetContactKey gets a reference to the given string and assigns it to the ContactKey field.
func (o *ClientContact) SetContactKey(v string) {
	o.ContactKey = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *ClientContact) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *ClientContact) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *ClientContact) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *ClientContact) GetConfirmed() bool {
	if o == nil || IsNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.Confirmed) {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *ClientContact) HasConfirmed() bool {
	if o != nil && !IsNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *ClientContact) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *ClientContact) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *ClientContact) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *ClientContact) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetSendEmail returns the SendEmail field value if set, zero value otherwise.
func (o *ClientContact) GetSendEmail() bool {
	if o == nil || IsNil(o.SendEmail) {
		var ret bool
		return ret
	}
	return *o.SendEmail
}

// GetSendEmailOk returns a tuple with the SendEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetSendEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.SendEmail) {
		return nil, false
	}
	return o.SendEmail, true
}

// HasSendEmail returns a boolean if a field has been set.
func (o *ClientContact) HasSendEmail() bool {
	if o != nil && !IsNil(o.SendEmail) {
		return true
	}

	return false
}

// SetSendEmail gets a reference to the given bool and assigns it to the SendEmail field.
func (o *ClientContact) SetSendEmail(v bool) {
	o.SendEmail = &v
}

// GetFailedLogins returns the FailedLogins field value if set, zero value otherwise.
func (o *ClientContact) GetFailedLogins() float32 {
	if o == nil || IsNil(o.FailedLogins) {
		var ret float32
		return ret
	}
	return *o.FailedLogins
}

// GetFailedLoginsOk returns a tuple with the FailedLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetFailedLoginsOk() (*float32, bool) {
	if o == nil || IsNil(o.FailedLogins) {
		return nil, false
	}
	return o.FailedLogins, true
}

// HasFailedLogins returns a boolean if a field has been set.
func (o *ClientContact) HasFailedLogins() bool {
	if o != nil && !IsNil(o.FailedLogins) {
		return true
	}

	return false
}

// SetFailedLogins gets a reference to the given float32 and assigns it to the FailedLogins field.
func (o *ClientContact) SetFailedLogins(v float32) {
	o.FailedLogins = &v
}

// GetEmailVerifiedAt returns the EmailVerifiedAt field value if set, zero value otherwise.
func (o *ClientContact) GetEmailVerifiedAt() float32 {
	if o == nil || IsNil(o.EmailVerifiedAt) {
		var ret float32
		return ret
	}
	return *o.EmailVerifiedAt
}

// GetEmailVerifiedAtOk returns a tuple with the EmailVerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetEmailVerifiedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.EmailVerifiedAt) {
		return nil, false
	}
	return o.EmailVerifiedAt, true
}

// HasEmailVerifiedAt returns a boolean if a field has been set.
func (o *ClientContact) HasEmailVerifiedAt() bool {
	if o != nil && !IsNil(o.EmailVerifiedAt) {
		return true
	}

	return false
}

// SetEmailVerifiedAt gets a reference to the given float32 and assigns it to the EmailVerifiedAt field.
func (o *ClientContact) SetEmailVerifiedAt(v float32) {
	o.EmailVerifiedAt = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *ClientContact) GetLastLogin() float32 {
	if o == nil || IsNil(o.LastLogin) {
		var ret float32
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetLastLoginOk() (*float32, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *ClientContact) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given float32 and assigns it to the LastLogin field.
func (o *ClientContact) SetLastLogin(v float32) {
	o.LastLogin = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ClientContact) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ClientContact) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *ClientContact) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ClientContact) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ClientContact) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *ClientContact) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *ClientContact) GetDeletedAt() float32 {
	if o == nil || IsNil(o.DeletedAt) {
		var ret float32
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientContact) GetDeletedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ClientContact) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given float32 and assigns it to the DeletedAt field.
func (o *ClientContact) SetDeletedAt(v float32) {
	o.DeletedAt = &v
}

func (o ClientContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.CustomValue1) {
		toSerialize["custom_value1"] = o.CustomValue1
	}
	if !IsNil(o.CustomValue2) {
		toSerialize["custom_value2"] = o.CustomValue2
	}
	if !IsNil(o.CustomValue3) {
		toSerialize["custom_value3"] = o.CustomValue3
	}
	if !IsNil(o.CustomValue4) {
		toSerialize["custom_value4"] = o.CustomValue4
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.AcceptedTermsVersion) {
		toSerialize["accepted_terms_version"] = o.AcceptedTermsVersion
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.ConfirmationCode) {
		toSerialize["confirmation_code"] = o.ConfirmationCode
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.ContactKey) {
		toSerialize["contact_key"] = o.ContactKey
	}
	if !IsNil(o.IsPrimary) {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if !IsNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !IsNil(o.IsLocked) {
		toSerialize["is_locked"] = o.IsLocked
	}
	if !IsNil(o.SendEmail) {
		toSerialize["send_email"] = o.SendEmail
	}
	if !IsNil(o.FailedLogins) {
		toSerialize["failed_logins"] = o.FailedLogins
	}
	if !IsNil(o.EmailVerifiedAt) {
		toSerialize["email_verified_at"] = o.EmailVerifiedAt
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	return toSerialize, nil
}

type NullableClientContact struct {
	value *ClientContact
	isSet bool
}

func (v NullableClientContact) Get() *ClientContact {
	return v.value
}

func (v *NullableClientContact) Set(val *ClientContact) {
	v.value = val
	v.isSet = true
}

func (v NullableClientContact) IsSet() bool {
	return v.isSet
}

func (v *NullableClientContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientContact(val *ClientContact) *NullableClientContact {
	return &NullableClientContact{value: val, isSet: true}
}

func (v NullableClientContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


