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

// checks if the CompanyUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyUser{}

// CompanyUser struct for CompanyUser
type CompanyUser struct {
	// The user permissions for this company in a comma separated list
	Permissions *string `json:"permissions,omitempty"`
	// Settings that are used for the flutter applications to store user preferences / metadata
	Settings map[string]interface{} `json:"settings,omitempty"`
	// Dedicated settings object for the react web application
	ReactSettings map[string]interface{} `json:"react_settings,omitempty"`
	// Determines whether the user owns this company
	IsOwner *bool `json:"is_owner,omitempty"`
	// Determines whether the user is the admin of this company
	IsAdmin *bool `json:"is_admin,omitempty"`
	// Determines whether the users access to this company has been locked
	IsLocked *bool `json:"is_locked,omitempty"`
	// The last time the record was modified, format Unix Timestamp
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// Timestamp when the user was archived, format Unix Timestamp
	DeletedAt *int32 `json:"deleted_at,omitempty"`
	Account *Account `json:"account,omitempty"`
	Company *Company `json:"company,omitempty"`
	User *UserRef `json:"user,omitempty"`
	Token *CompanyToken `json:"token,omitempty"`
}

// NewCompanyUser instantiates a new CompanyUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyUser() *CompanyUser {
	this := CompanyUser{}
	return &this
}

// NewCompanyUserWithDefaults instantiates a new CompanyUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyUserWithDefaults() *CompanyUser {
	this := CompanyUser{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CompanyUser) GetPermissions() string {
	if o == nil || IsNil(o.Permissions) {
		var ret string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetPermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CompanyUser) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given string and assigns it to the Permissions field.
func (o *CompanyUser) SetPermissions(v string) {
	o.Permissions = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CompanyUser) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CompanyUser) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *CompanyUser) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetReactSettings returns the ReactSettings field value if set, zero value otherwise.
func (o *CompanyUser) GetReactSettings() map[string]interface{} {
	if o == nil || IsNil(o.ReactSettings) {
		var ret map[string]interface{}
		return ret
	}
	return o.ReactSettings
}

// GetReactSettingsOk returns a tuple with the ReactSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetReactSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReactSettings) {
		return map[string]interface{}{}, false
	}
	return o.ReactSettings, true
}

// HasReactSettings returns a boolean if a field has been set.
func (o *CompanyUser) HasReactSettings() bool {
	if o != nil && !IsNil(o.ReactSettings) {
		return true
	}

	return false
}

// SetReactSettings gets a reference to the given map[string]interface{} and assigns it to the ReactSettings field.
func (o *CompanyUser) SetReactSettings(v map[string]interface{}) {
	o.ReactSettings = v
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise.
func (o *CompanyUser) GetIsOwner() bool {
	if o == nil || IsNil(o.IsOwner) {
		var ret bool
		return ret
	}
	return *o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetIsOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOwner) {
		return nil, false
	}
	return o.IsOwner, true
}

// HasIsOwner returns a boolean if a field has been set.
func (o *CompanyUser) HasIsOwner() bool {
	if o != nil && !IsNil(o.IsOwner) {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given bool and assigns it to the IsOwner field.
func (o *CompanyUser) SetIsOwner(v bool) {
	o.IsOwner = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *CompanyUser) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *CompanyUser) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *CompanyUser) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *CompanyUser) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *CompanyUser) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *CompanyUser) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CompanyUser) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CompanyUser) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *CompanyUser) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *CompanyUser) GetDeletedAt() int32 {
	if o == nil || IsNil(o.DeletedAt) {
		var ret int32
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetDeletedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CompanyUser) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given int32 and assigns it to the DeletedAt field.
func (o *CompanyUser) SetDeletedAt(v int32) {
	o.DeletedAt = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CompanyUser) GetAccount() Account {
	if o == nil || IsNil(o.Account) {
		var ret Account
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetAccountOk() (*Account, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CompanyUser) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given Account and assigns it to the Account field.
func (o *CompanyUser) SetAccount(v Account) {
	o.Account = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CompanyUser) GetCompany() Company {
	if o == nil || IsNil(o.Company) {
		var ret Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetCompanyOk() (*Company, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CompanyUser) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given Company and assigns it to the Company field.
func (o *CompanyUser) SetCompany(v Company) {
	o.Company = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CompanyUser) GetUser() UserRef {
	if o == nil || IsNil(o.User) {
		var ret UserRef
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetUserOk() (*UserRef, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CompanyUser) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserRef and assigns it to the User field.
func (o *CompanyUser) SetUser(v UserRef) {
	o.User = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CompanyUser) GetToken() CompanyToken {
	if o == nil || IsNil(o.Token) {
		var ret CompanyToken
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyUser) GetTokenOk() (*CompanyToken, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CompanyUser) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given CompanyToken and assigns it to the Token field.
func (o *CompanyUser) SetToken(v CompanyToken) {
	o.Token = &v
}

func (o CompanyUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.ReactSettings) {
		toSerialize["react_settings"] = o.ReactSettings
	}
	if !IsNil(o.IsOwner) {
		toSerialize["is_owner"] = o.IsOwner
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["is_admin"] = o.IsAdmin
	}
	if !IsNil(o.IsLocked) {
		toSerialize["is_locked"] = o.IsLocked
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableCompanyUser struct {
	value *CompanyUser
	isSet bool
}

func (v NullableCompanyUser) Get() *CompanyUser {
	return v.value
}

func (v *NullableCompanyUser) Set(val *CompanyUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyUser(val *CompanyUser) *NullableCompanyUser {
	return &NullableCompanyUser{value: val, isSet: true}
}

func (v NullableCompanyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


