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

// checks if the Subscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscription{}

// Subscription struct for Subscription
type Subscription struct {
	// Unique identifier for the subscription
	Id *string `json:"id,omitempty"`
	// Unique identifier for the user associated with the subscription
	UserId *string `json:"user_id,omitempty"`
	// Unique identifier for the product associated with the subscription
	ProductId *string `json:"product_id,omitempty"`
	// Unique identifier for the recurring invoice associated with the subscription
	RecurringInvoiceId *string `json:"recurring_invoice_id,omitempty"`
	// Indicates whether the subscription is recurring
	IsRecurring *bool `json:"is_recurring,omitempty"`
	// integer const representation of the frequency
	FrequencyId *string `json:"frequency_id,omitempty"`
	// enum setting
	AutoBill *string `json:"auto_bill,omitempty"`
	// Promotional code applied to the subscription
	PromoCode *string `json:"promo_code,omitempty"`
	// Discount percentage or amount applied to the subscription
	PromoDiscount *float32 `json:"promo_discount,omitempty"`
	// Indicates whether the discount is a fixed amount
	IsAmountDiscount *bool `json:"is_amount_discount,omitempty"`
	// Indicates whether the subscription can be cancelled
	AllowCancellation *bool `json:"allow_cancellation,omitempty"`
	// Indicates whether the subscription pricing is per seat
	PerSeatEnabled *bool `json:"per_seat_enabled,omitempty"`
	// Unique identifier for the currency used in the subscription
	CurrencyId *int32 `json:"currency_id,omitempty"`
	// Maximum number of seats allowed for the subscription
	MaxSeatsLimit *int32 `json:"max_seats_limit,omitempty"`
	// Indicates whether the subscription has a trial period
	TrialEnabled *bool `json:"trial_enabled,omitempty"`
	// Duration of the trial period in days
	TrialDuration *int32 `json:"trial_duration,omitempty"`
	// Indicates whether query overrides are allowed for the subscription
	AllowQueryOverrides *bool `json:"allow_query_overrides,omitempty"`
	// Indicates whether plan changes are allowed for the subscription
	AllowPlanChanges *bool `json:"allow_plan_changes,omitempty"`
	// Number of days within which refunds can be requested
	RefundPeriod *int32 `json:"refund_period,omitempty"`
	// Webhook configuration for the subscription
	WebhookConfiguration *string `json:"webhook_configuration,omitempty"`
	// Indicates whether the subscription has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription() *Subscription {
	this := Subscription{}
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subscription) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subscription) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Subscription) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Subscription) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Subscription) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Subscription) SetUserId(v string) {
	o.UserId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *Subscription) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *Subscription) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *Subscription) SetProductId(v string) {
	o.ProductId = &v
}

// GetRecurringInvoiceId returns the RecurringInvoiceId field value if set, zero value otherwise.
func (o *Subscription) GetRecurringInvoiceId() string {
	if o == nil || IsNil(o.RecurringInvoiceId) {
		var ret string
		return ret
	}
	return *o.RecurringInvoiceId
}

// GetRecurringInvoiceIdOk returns a tuple with the RecurringInvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetRecurringInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringInvoiceId) {
		return nil, false
	}
	return o.RecurringInvoiceId, true
}

// HasRecurringInvoiceId returns a boolean if a field has been set.
func (o *Subscription) HasRecurringInvoiceId() bool {
	if o != nil && !IsNil(o.RecurringInvoiceId) {
		return true
	}

	return false
}

// SetRecurringInvoiceId gets a reference to the given string and assigns it to the RecurringInvoiceId field.
func (o *Subscription) SetRecurringInvoiceId(v string) {
	o.RecurringInvoiceId = &v
}

// GetIsRecurring returns the IsRecurring field value if set, zero value otherwise.
func (o *Subscription) GetIsRecurring() bool {
	if o == nil || IsNil(o.IsRecurring) {
		var ret bool
		return ret
	}
	return *o.IsRecurring
}

// GetIsRecurringOk returns a tuple with the IsRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIsRecurringOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRecurring) {
		return nil, false
	}
	return o.IsRecurring, true
}

// HasIsRecurring returns a boolean if a field has been set.
func (o *Subscription) HasIsRecurring() bool {
	if o != nil && !IsNil(o.IsRecurring) {
		return true
	}

	return false
}

// SetIsRecurring gets a reference to the given bool and assigns it to the IsRecurring field.
func (o *Subscription) SetIsRecurring(v bool) {
	o.IsRecurring = &v
}

// GetFrequencyId returns the FrequencyId field value if set, zero value otherwise.
func (o *Subscription) GetFrequencyId() string {
	if o == nil || IsNil(o.FrequencyId) {
		var ret string
		return ret
	}
	return *o.FrequencyId
}

// GetFrequencyIdOk returns a tuple with the FrequencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetFrequencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyId) {
		return nil, false
	}
	return o.FrequencyId, true
}

// HasFrequencyId returns a boolean if a field has been set.
func (o *Subscription) HasFrequencyId() bool {
	if o != nil && !IsNil(o.FrequencyId) {
		return true
	}

	return false
}

// SetFrequencyId gets a reference to the given string and assigns it to the FrequencyId field.
func (o *Subscription) SetFrequencyId(v string) {
	o.FrequencyId = &v
}

// GetAutoBill returns the AutoBill field value if set, zero value otherwise.
func (o *Subscription) GetAutoBill() string {
	if o == nil || IsNil(o.AutoBill) {
		var ret string
		return ret
	}
	return *o.AutoBill
}

// GetAutoBillOk returns a tuple with the AutoBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetAutoBillOk() (*string, bool) {
	if o == nil || IsNil(o.AutoBill) {
		return nil, false
	}
	return o.AutoBill, true
}

// HasAutoBill returns a boolean if a field has been set.
func (o *Subscription) HasAutoBill() bool {
	if o != nil && !IsNil(o.AutoBill) {
		return true
	}

	return false
}

// SetAutoBill gets a reference to the given string and assigns it to the AutoBill field.
func (o *Subscription) SetAutoBill(v string) {
	o.AutoBill = &v
}

// GetPromoCode returns the PromoCode field value if set, zero value otherwise.
func (o *Subscription) GetPromoCode() string {
	if o == nil || IsNil(o.PromoCode) {
		var ret string
		return ret
	}
	return *o.PromoCode
}

// GetPromoCodeOk returns a tuple with the PromoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPromoCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PromoCode) {
		return nil, false
	}
	return o.PromoCode, true
}

// HasPromoCode returns a boolean if a field has been set.
func (o *Subscription) HasPromoCode() bool {
	if o != nil && !IsNil(o.PromoCode) {
		return true
	}

	return false
}

// SetPromoCode gets a reference to the given string and assigns it to the PromoCode field.
func (o *Subscription) SetPromoCode(v string) {
	o.PromoCode = &v
}

// GetPromoDiscount returns the PromoDiscount field value if set, zero value otherwise.
func (o *Subscription) GetPromoDiscount() float32 {
	if o == nil || IsNil(o.PromoDiscount) {
		var ret float32
		return ret
	}
	return *o.PromoDiscount
}

// GetPromoDiscountOk returns a tuple with the PromoDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPromoDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.PromoDiscount) {
		return nil, false
	}
	return o.PromoDiscount, true
}

// HasPromoDiscount returns a boolean if a field has been set.
func (o *Subscription) HasPromoDiscount() bool {
	if o != nil && !IsNil(o.PromoDiscount) {
		return true
	}

	return false
}

// SetPromoDiscount gets a reference to the given float32 and assigns it to the PromoDiscount field.
func (o *Subscription) SetPromoDiscount(v float32) {
	o.PromoDiscount = &v
}

// GetIsAmountDiscount returns the IsAmountDiscount field value if set, zero value otherwise.
func (o *Subscription) GetIsAmountDiscount() bool {
	if o == nil || IsNil(o.IsAmountDiscount) {
		var ret bool
		return ret
	}
	return *o.IsAmountDiscount
}

// GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIsAmountDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountDiscount) {
		return nil, false
	}
	return o.IsAmountDiscount, true
}

// HasIsAmountDiscount returns a boolean if a field has been set.
func (o *Subscription) HasIsAmountDiscount() bool {
	if o != nil && !IsNil(o.IsAmountDiscount) {
		return true
	}

	return false
}

// SetIsAmountDiscount gets a reference to the given bool and assigns it to the IsAmountDiscount field.
func (o *Subscription) SetIsAmountDiscount(v bool) {
	o.IsAmountDiscount = &v
}

// GetAllowCancellation returns the AllowCancellation field value if set, zero value otherwise.
func (o *Subscription) GetAllowCancellation() bool {
	if o == nil || IsNil(o.AllowCancellation) {
		var ret bool
		return ret
	}
	return *o.AllowCancellation
}

// GetAllowCancellationOk returns a tuple with the AllowCancellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetAllowCancellationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCancellation) {
		return nil, false
	}
	return o.AllowCancellation, true
}

// HasAllowCancellation returns a boolean if a field has been set.
func (o *Subscription) HasAllowCancellation() bool {
	if o != nil && !IsNil(o.AllowCancellation) {
		return true
	}

	return false
}

// SetAllowCancellation gets a reference to the given bool and assigns it to the AllowCancellation field.
func (o *Subscription) SetAllowCancellation(v bool) {
	o.AllowCancellation = &v
}

// GetPerSeatEnabled returns the PerSeatEnabled field value if set, zero value otherwise.
func (o *Subscription) GetPerSeatEnabled() bool {
	if o == nil || IsNil(o.PerSeatEnabled) {
		var ret bool
		return ret
	}
	return *o.PerSeatEnabled
}

// GetPerSeatEnabledOk returns a tuple with the PerSeatEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPerSeatEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PerSeatEnabled) {
		return nil, false
	}
	return o.PerSeatEnabled, true
}

// HasPerSeatEnabled returns a boolean if a field has been set.
func (o *Subscription) HasPerSeatEnabled() bool {
	if o != nil && !IsNil(o.PerSeatEnabled) {
		return true
	}

	return false
}

// SetPerSeatEnabled gets a reference to the given bool and assigns it to the PerSeatEnabled field.
func (o *Subscription) SetPerSeatEnabled(v bool) {
	o.PerSeatEnabled = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *Subscription) GetCurrencyId() int32 {
	if o == nil || IsNil(o.CurrencyId) {
		var ret int32
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrencyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *Subscription) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given int32 and assigns it to the CurrencyId field.
func (o *Subscription) SetCurrencyId(v int32) {
	o.CurrencyId = &v
}

// GetMaxSeatsLimit returns the MaxSeatsLimit field value if set, zero value otherwise.
func (o *Subscription) GetMaxSeatsLimit() int32 {
	if o == nil || IsNil(o.MaxSeatsLimit) {
		var ret int32
		return ret
	}
	return *o.MaxSeatsLimit
}

// GetMaxSeatsLimitOk returns a tuple with the MaxSeatsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetMaxSeatsLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSeatsLimit) {
		return nil, false
	}
	return o.MaxSeatsLimit, true
}

// HasMaxSeatsLimit returns a boolean if a field has been set.
func (o *Subscription) HasMaxSeatsLimit() bool {
	if o != nil && !IsNil(o.MaxSeatsLimit) {
		return true
	}

	return false
}

// SetMaxSeatsLimit gets a reference to the given int32 and assigns it to the MaxSeatsLimit field.
func (o *Subscription) SetMaxSeatsLimit(v int32) {
	o.MaxSeatsLimit = &v
}

// GetTrialEnabled returns the TrialEnabled field value if set, zero value otherwise.
func (o *Subscription) GetTrialEnabled() bool {
	if o == nil || IsNil(o.TrialEnabled) {
		var ret bool
		return ret
	}
	return *o.TrialEnabled
}

// GetTrialEnabledOk returns a tuple with the TrialEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTrialEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TrialEnabled) {
		return nil, false
	}
	return o.TrialEnabled, true
}

// HasTrialEnabled returns a boolean if a field has been set.
func (o *Subscription) HasTrialEnabled() bool {
	if o != nil && !IsNil(o.TrialEnabled) {
		return true
	}

	return false
}

// SetTrialEnabled gets a reference to the given bool and assigns it to the TrialEnabled field.
func (o *Subscription) SetTrialEnabled(v bool) {
	o.TrialEnabled = &v
}

// GetTrialDuration returns the TrialDuration field value if set, zero value otherwise.
func (o *Subscription) GetTrialDuration() int32 {
	if o == nil || IsNil(o.TrialDuration) {
		var ret int32
		return ret
	}
	return *o.TrialDuration
}

// GetTrialDurationOk returns a tuple with the TrialDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTrialDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialDuration) {
		return nil, false
	}
	return o.TrialDuration, true
}

// HasTrialDuration returns a boolean if a field has been set.
func (o *Subscription) HasTrialDuration() bool {
	if o != nil && !IsNil(o.TrialDuration) {
		return true
	}

	return false
}

// SetTrialDuration gets a reference to the given int32 and assigns it to the TrialDuration field.
func (o *Subscription) SetTrialDuration(v int32) {
	o.TrialDuration = &v
}

// GetAllowQueryOverrides returns the AllowQueryOverrides field value if set, zero value otherwise.
func (o *Subscription) GetAllowQueryOverrides() bool {
	if o == nil || IsNil(o.AllowQueryOverrides) {
		var ret bool
		return ret
	}
	return *o.AllowQueryOverrides
}

// GetAllowQueryOverridesOk returns a tuple with the AllowQueryOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetAllowQueryOverridesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowQueryOverrides) {
		return nil, false
	}
	return o.AllowQueryOverrides, true
}

// HasAllowQueryOverrides returns a boolean if a field has been set.
func (o *Subscription) HasAllowQueryOverrides() bool {
	if o != nil && !IsNil(o.AllowQueryOverrides) {
		return true
	}

	return false
}

// SetAllowQueryOverrides gets a reference to the given bool and assigns it to the AllowQueryOverrides field.
func (o *Subscription) SetAllowQueryOverrides(v bool) {
	o.AllowQueryOverrides = &v
}

// GetAllowPlanChanges returns the AllowPlanChanges field value if set, zero value otherwise.
func (o *Subscription) GetAllowPlanChanges() bool {
	if o == nil || IsNil(o.AllowPlanChanges) {
		var ret bool
		return ret
	}
	return *o.AllowPlanChanges
}

// GetAllowPlanChangesOk returns a tuple with the AllowPlanChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetAllowPlanChangesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPlanChanges) {
		return nil, false
	}
	return o.AllowPlanChanges, true
}

// HasAllowPlanChanges returns a boolean if a field has been set.
func (o *Subscription) HasAllowPlanChanges() bool {
	if o != nil && !IsNil(o.AllowPlanChanges) {
		return true
	}

	return false
}

// SetAllowPlanChanges gets a reference to the given bool and assigns it to the AllowPlanChanges field.
func (o *Subscription) SetAllowPlanChanges(v bool) {
	o.AllowPlanChanges = &v
}

// GetRefundPeriod returns the RefundPeriod field value if set, zero value otherwise.
func (o *Subscription) GetRefundPeriod() int32 {
	if o == nil || IsNil(o.RefundPeriod) {
		var ret int32
		return ret
	}
	return *o.RefundPeriod
}

// GetRefundPeriodOk returns a tuple with the RefundPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetRefundPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.RefundPeriod) {
		return nil, false
	}
	return o.RefundPeriod, true
}

// HasRefundPeriod returns a boolean if a field has been set.
func (o *Subscription) HasRefundPeriod() bool {
	if o != nil && !IsNil(o.RefundPeriod) {
		return true
	}

	return false
}

// SetRefundPeriod gets a reference to the given int32 and assigns it to the RefundPeriod field.
func (o *Subscription) SetRefundPeriod(v int32) {
	o.RefundPeriod = &v
}

// GetWebhookConfiguration returns the WebhookConfiguration field value if set, zero value otherwise.
func (o *Subscription) GetWebhookConfiguration() string {
	if o == nil || IsNil(o.WebhookConfiguration) {
		var ret string
		return ret
	}
	return *o.WebhookConfiguration
}

// GetWebhookConfigurationOk returns a tuple with the WebhookConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetWebhookConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookConfiguration) {
		return nil, false
	}
	return o.WebhookConfiguration, true
}

// HasWebhookConfiguration returns a boolean if a field has been set.
func (o *Subscription) HasWebhookConfiguration() bool {
	if o != nil && !IsNil(o.WebhookConfiguration) {
		return true
	}

	return false
}

// SetWebhookConfiguration gets a reference to the given string and assigns it to the WebhookConfiguration field.
func (o *Subscription) SetWebhookConfiguration(v string) {
	o.WebhookConfiguration = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Subscription) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Subscription) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Subscription) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Subscription) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Subscription) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *Subscription) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Subscription) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Subscription) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *Subscription) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Subscription) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Subscription) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Subscription) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.RecurringInvoiceId) {
		toSerialize["recurring_invoice_id"] = o.RecurringInvoiceId
	}
	if !IsNil(o.IsRecurring) {
		toSerialize["is_recurring"] = o.IsRecurring
	}
	if !IsNil(o.FrequencyId) {
		toSerialize["frequency_id"] = o.FrequencyId
	}
	if !IsNil(o.AutoBill) {
		toSerialize["auto_bill"] = o.AutoBill
	}
	if !IsNil(o.PromoCode) {
		toSerialize["promo_code"] = o.PromoCode
	}
	if !IsNil(o.PromoDiscount) {
		toSerialize["promo_discount"] = o.PromoDiscount
	}
	if !IsNil(o.IsAmountDiscount) {
		toSerialize["is_amount_discount"] = o.IsAmountDiscount
	}
	if !IsNil(o.AllowCancellation) {
		toSerialize["allow_cancellation"] = o.AllowCancellation
	}
	if !IsNil(o.PerSeatEnabled) {
		toSerialize["per_seat_enabled"] = o.PerSeatEnabled
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.MaxSeatsLimit) {
		toSerialize["max_seats_limit"] = o.MaxSeatsLimit
	}
	if !IsNil(o.TrialEnabled) {
		toSerialize["trial_enabled"] = o.TrialEnabled
	}
	if !IsNil(o.TrialDuration) {
		toSerialize["trial_duration"] = o.TrialDuration
	}
	if !IsNil(o.AllowQueryOverrides) {
		toSerialize["allow_query_overrides"] = o.AllowQueryOverrides
	}
	if !IsNil(o.AllowPlanChanges) {
		toSerialize["allow_plan_changes"] = o.AllowPlanChanges
	}
	if !IsNil(o.RefundPeriod) {
		toSerialize["refund_period"] = o.RefundPeriod
	}
	if !IsNil(o.WebhookConfiguration) {
		toSerialize["webhook_configuration"] = o.WebhookConfiguration
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


