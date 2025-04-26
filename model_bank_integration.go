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

// checks if the BankIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankIntegration{}

// BankIntegration struct for BankIntegration
type BankIntegration struct {
	// The bank integration hashed id
	Id *string `json:"id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The providers bank name
	ProviderBankName *string `json:"provider_bank_name,omitempty"`
	// The bank account id
	BankAccountId *int32 `json:"bank_account_id,omitempty"`
	// The name of the account
	BankAccountName *string `json:"bank_account_name,omitempty"`
	// The account number
	BankAccountNumber *string `json:"bank_account_number,omitempty"`
	// The status of the bank account
	BankAccountStatus *string `json:"bank_account_status,omitempty"`
	// The type of account
	BankAccountType *string `json:"bank_account_type,omitempty"`
	// The current bank balance if available
	Balance *float32 `json:"balance,omitempty"`
	// iso_3166_3 code
	Currency *string `json:"currency,omitempty"`
}

// NewBankIntegration instantiates a new BankIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankIntegration() *BankIntegration {
	this := BankIntegration{}
	return &this
}

// NewBankIntegrationWithDefaults instantiates a new BankIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankIntegrationWithDefaults() *BankIntegration {
	this := BankIntegration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankIntegration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankIntegration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankIntegration) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BankIntegration) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BankIntegration) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BankIntegration) SetUserId(v string) {
	o.UserId = &v
}

// GetProviderBankName returns the ProviderBankName field value if set, zero value otherwise.
func (o *BankIntegration) GetProviderBankName() string {
	if o == nil || IsNil(o.ProviderBankName) {
		var ret string
		return ret
	}
	return *o.ProviderBankName
}

// GetProviderBankNameOk returns a tuple with the ProviderBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetProviderBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderBankName) {
		return nil, false
	}
	return o.ProviderBankName, true
}

// HasProviderBankName returns a boolean if a field has been set.
func (o *BankIntegration) HasProviderBankName() bool {
	if o != nil && !IsNil(o.ProviderBankName) {
		return true
	}

	return false
}

// SetProviderBankName gets a reference to the given string and assigns it to the ProviderBankName field.
func (o *BankIntegration) SetProviderBankName(v string) {
	o.ProviderBankName = &v
}

// GetBankAccountId returns the BankAccountId field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountId() int32 {
	if o == nil || IsNil(o.BankAccountId) {
		var ret int32
		return ret
	}
	return *o.BankAccountId
}

// GetBankAccountIdOk returns a tuple with the BankAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BankAccountId) {
		return nil, false
	}
	return o.BankAccountId, true
}

// HasBankAccountId returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountId() bool {
	if o != nil && !IsNil(o.BankAccountId) {
		return true
	}

	return false
}

// SetBankAccountId gets a reference to the given int32 and assigns it to the BankAccountId field.
func (o *BankIntegration) SetBankAccountId(v int32) {
	o.BankAccountId = &v
}

// GetBankAccountName returns the BankAccountName field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountName() string {
	if o == nil || IsNil(o.BankAccountName) {
		var ret string
		return ret
	}
	return *o.BankAccountName
}

// GetBankAccountNameOk returns a tuple with the BankAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountName) {
		return nil, false
	}
	return o.BankAccountName, true
}

// HasBankAccountName returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountName() bool {
	if o != nil && !IsNil(o.BankAccountName) {
		return true
	}

	return false
}

// SetBankAccountName gets a reference to the given string and assigns it to the BankAccountName field.
func (o *BankIntegration) SetBankAccountName(v string) {
	o.BankAccountName = &v
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountNumber() string {
	if o == nil || IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountNumber() bool {
	if o != nil && !IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *BankIntegration) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetBankAccountStatus returns the BankAccountStatus field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountStatus() string {
	if o == nil || IsNil(o.BankAccountStatus) {
		var ret string
		return ret
	}
	return *o.BankAccountStatus
}

// GetBankAccountStatusOk returns a tuple with the BankAccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountStatus) {
		return nil, false
	}
	return o.BankAccountStatus, true
}

// HasBankAccountStatus returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountStatus() bool {
	if o != nil && !IsNil(o.BankAccountStatus) {
		return true
	}

	return false
}

// SetBankAccountStatus gets a reference to the given string and assigns it to the BankAccountStatus field.
func (o *BankIntegration) SetBankAccountStatus(v string) {
	o.BankAccountStatus = &v
}

// GetBankAccountType returns the BankAccountType field value if set, zero value otherwise.
func (o *BankIntegration) GetBankAccountType() string {
	if o == nil || IsNil(o.BankAccountType) {
		var ret string
		return ret
	}
	return *o.BankAccountType
}

// GetBankAccountTypeOk returns a tuple with the BankAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBankAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountType) {
		return nil, false
	}
	return o.BankAccountType, true
}

// HasBankAccountType returns a boolean if a field has been set.
func (o *BankIntegration) HasBankAccountType() bool {
	if o != nil && !IsNil(o.BankAccountType) {
		return true
	}

	return false
}

// SetBankAccountType gets a reference to the given string and assigns it to the BankAccountType field.
func (o *BankIntegration) SetBankAccountType(v string) {
	o.BankAccountType = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *BankIntegration) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *BankIntegration) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *BankIntegration) SetBalance(v float32) {
	o.Balance = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BankIntegration) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankIntegration) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BankIntegration) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BankIntegration) SetCurrency(v string) {
	o.Currency = &v
}

func (o BankIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ProviderBankName) {
		toSerialize["provider_bank_name"] = o.ProviderBankName
	}
	if !IsNil(o.BankAccountId) {
		toSerialize["bank_account_id"] = o.BankAccountId
	}
	if !IsNil(o.BankAccountName) {
		toSerialize["bank_account_name"] = o.BankAccountName
	}
	if !IsNil(o.BankAccountNumber) {
		toSerialize["bank_account_number"] = o.BankAccountNumber
	}
	if !IsNil(o.BankAccountStatus) {
		toSerialize["bank_account_status"] = o.BankAccountStatus
	}
	if !IsNil(o.BankAccountType) {
		toSerialize["bank_account_type"] = o.BankAccountType
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableBankIntegration struct {
	value *BankIntegration
	isSet bool
}

func (v NullableBankIntegration) Get() *BankIntegration {
	return v.value
}

func (v *NullableBankIntegration) Set(val *BankIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableBankIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableBankIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankIntegration(val *BankIntegration) *NullableBankIntegration {
	return &NullableBankIntegration{value: val, isSet: true}
}

func (v NullableBankIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


