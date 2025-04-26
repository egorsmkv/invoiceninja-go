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

// checks if the UpdateTaskSchedulerSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTaskSchedulerSchema{}

// UpdateTaskSchedulerSchema struct for UpdateTaskSchedulerSchema
type UpdateTaskSchedulerSchema struct {
	// The scheduler paused state
	Paused *bool `json:"paused,omitempty"`
	// Accepted values (DAY,WEEK,MONTH,3MONTHS,YEAR)
	RepeatEvery *string `json:"repeat_every,omitempty"`
	// Timestamp when we should start the scheduler, default is today
	StartFrom *int32 `json:"start_from,omitempty"`
	// Job, we can find list of available jobs in Scheduler model
	Job *string `json:"job,omitempty"`
	// The string representation of the date range of data to be returned
	DateRange *string `json:"date_range,omitempty"`
	// The date column to search between.
	DateKey *string `json:"date_key,omitempty"`
	// The start date to search between
	StartDate *string `json:"start_date,omitempty"`
	// The end date to search between
	EndDate *string `json:"end_date,omitempty"`
}

// NewUpdateTaskSchedulerSchema instantiates a new UpdateTaskSchedulerSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTaskSchedulerSchema() *UpdateTaskSchedulerSchema {
	this := UpdateTaskSchedulerSchema{}
	return &this
}

// NewUpdateTaskSchedulerSchemaWithDefaults instantiates a new UpdateTaskSchedulerSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTaskSchedulerSchemaWithDefaults() *UpdateTaskSchedulerSchema {
	this := UpdateTaskSchedulerSchema{}
	return &this
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetPaused() bool {
	if o == nil || IsNil(o.Paused) {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.Paused) {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasPaused() bool {
	if o != nil && !IsNil(o.Paused) {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *UpdateTaskSchedulerSchema) SetPaused(v bool) {
	o.Paused = &v
}

// GetRepeatEvery returns the RepeatEvery field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetRepeatEvery() string {
	if o == nil || IsNil(o.RepeatEvery) {
		var ret string
		return ret
	}
	return *o.RepeatEvery
}

// GetRepeatEveryOk returns a tuple with the RepeatEvery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetRepeatEveryOk() (*string, bool) {
	if o == nil || IsNil(o.RepeatEvery) {
		return nil, false
	}
	return o.RepeatEvery, true
}

// HasRepeatEvery returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasRepeatEvery() bool {
	if o != nil && !IsNil(o.RepeatEvery) {
		return true
	}

	return false
}

// SetRepeatEvery gets a reference to the given string and assigns it to the RepeatEvery field.
func (o *UpdateTaskSchedulerSchema) SetRepeatEvery(v string) {
	o.RepeatEvery = &v
}

// GetStartFrom returns the StartFrom field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetStartFrom() int32 {
	if o == nil || IsNil(o.StartFrom) {
		var ret int32
		return ret
	}
	return *o.StartFrom
}

// GetStartFromOk returns a tuple with the StartFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetStartFromOk() (*int32, bool) {
	if o == nil || IsNil(o.StartFrom) {
		return nil, false
	}
	return o.StartFrom, true
}

// HasStartFrom returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasStartFrom() bool {
	if o != nil && !IsNil(o.StartFrom) {
		return true
	}

	return false
}

// SetStartFrom gets a reference to the given int32 and assigns it to the StartFrom field.
func (o *UpdateTaskSchedulerSchema) SetStartFrom(v int32) {
	o.StartFrom = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *UpdateTaskSchedulerSchema) SetJob(v string) {
	o.Job = &v
}

// GetDateRange returns the DateRange field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetDateRange() string {
	if o == nil || IsNil(o.DateRange) {
		var ret string
		return ret
	}
	return *o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetDateRangeOk() (*string, bool) {
	if o == nil || IsNil(o.DateRange) {
		return nil, false
	}
	return o.DateRange, true
}

// HasDateRange returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasDateRange() bool {
	if o != nil && !IsNil(o.DateRange) {
		return true
	}

	return false
}

// SetDateRange gets a reference to the given string and assigns it to the DateRange field.
func (o *UpdateTaskSchedulerSchema) SetDateRange(v string) {
	o.DateRange = &v
}

// GetDateKey returns the DateKey field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetDateKey() string {
	if o == nil || IsNil(o.DateKey) {
		var ret string
		return ret
	}
	return *o.DateKey
}

// GetDateKeyOk returns a tuple with the DateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetDateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DateKey) {
		return nil, false
	}
	return o.DateKey, true
}

// HasDateKey returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasDateKey() bool {
	if o != nil && !IsNil(o.DateKey) {
		return true
	}

	return false
}

// SetDateKey gets a reference to the given string and assigns it to the DateKey field.
func (o *UpdateTaskSchedulerSchema) SetDateKey(v string) {
	o.DateKey = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *UpdateTaskSchedulerSchema) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UpdateTaskSchedulerSchema) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTaskSchedulerSchema) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UpdateTaskSchedulerSchema) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *UpdateTaskSchedulerSchema) SetEndDate(v string) {
	o.EndDate = &v
}

func (o UpdateTaskSchedulerSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTaskSchedulerSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paused) {
		toSerialize["paused"] = o.Paused
	}
	if !IsNil(o.RepeatEvery) {
		toSerialize["repeat_every"] = o.RepeatEvery
	}
	if !IsNil(o.StartFrom) {
		toSerialize["start_from"] = o.StartFrom
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.DateRange) {
		toSerialize["date_range"] = o.DateRange
	}
	if !IsNil(o.DateKey) {
		toSerialize["date_key"] = o.DateKey
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableUpdateTaskSchedulerSchema struct {
	value *UpdateTaskSchedulerSchema
	isSet bool
}

func (v NullableUpdateTaskSchedulerSchema) Get() *UpdateTaskSchedulerSchema {
	return v.value
}

func (v *NullableUpdateTaskSchedulerSchema) Set(val *UpdateTaskSchedulerSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTaskSchedulerSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTaskSchedulerSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTaskSchedulerSchema(val *UpdateTaskSchedulerSchema) *NullableUpdateTaskSchedulerSchema {
	return &NullableUpdateTaskSchedulerSchema{value: val, isSet: true}
}

func (v NullableUpdateTaskSchedulerSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTaskSchedulerSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


