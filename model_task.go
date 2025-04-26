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

// checks if the Task type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Task{}

// Task struct for Task
type Task struct {
	// The hashed id of the task
	Id *string `json:"id,omitempty"`
	// The hashed id of the user who created the task
	UserId *string `json:"user_id,omitempty"`
	// The assigned user of the task
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The hashed if of the client
	ClientId *string `json:"client_id,omitempty"`
	// The hashed id of the invoice associated with the task
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The hashed id of the project associated with the task
	ProjectId *string `json:"project_id,omitempty"`
	// The number of the task
	Number *string `json:"number,omitempty"`
	// An array of unix time stamps defining the start and end times of the task
	TimeLog *string `json:"time_log,omitempty"`
	// Determines if the task is still running
	IsRunning *bool `json:"is_running,omitempty"`
	// Boolean flag determining if the task has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The hashed id of the task status
	TaskStatusId *string `json:"task_status_id,omitempty"`
	// The task description
	Description *string `json:"description,omitempty"`
	// The task duration in seconds
	Duration *int32 `json:"duration,omitempty"`
	// The order of the task
	TaskStatusOrder *int32 `json:"task_status_order,omitempty"`
	// The task rate
	Rate *float32 `json:"rate,omitempty"`
	// A custom value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A custom value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A custom value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A custom value
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// Boolean flag determining if the task is date based
	IsDateBased *bool `json:"is_date_based,omitempty"`
	// The calculated start date of the task
	CalculatedStartDate *string `json:"calculated_start_date,omitempty"`
	// Boolean flags which determines whether to include the task documents on the invoice
	InvoiceDocuments *bool `json:"invoice_documents,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Task) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Task) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Task) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Task) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Task) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Task) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Task) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Task) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Task) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Task) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Task) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Task) SetClientId(v string) {
	o.ClientId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *Task) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Task) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *Task) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Task) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Task) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *Task) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Task) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Task) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Task) SetNumber(v string) {
	o.Number = &v
}

// GetTimeLog returns the TimeLog field value if set, zero value otherwise.
func (o *Task) GetTimeLog() string {
	if o == nil || IsNil(o.TimeLog) {
		var ret string
		return ret
	}
	return *o.TimeLog
}

// GetTimeLogOk returns a tuple with the TimeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTimeLogOk() (*string, bool) {
	if o == nil || IsNil(o.TimeLog) {
		return nil, false
	}
	return o.TimeLog, true
}

// HasTimeLog returns a boolean if a field has been set.
func (o *Task) HasTimeLog() bool {
	if o != nil && !IsNil(o.TimeLog) {
		return true
	}

	return false
}

// SetTimeLog gets a reference to the given string and assigns it to the TimeLog field.
func (o *Task) SetTimeLog(v string) {
	o.TimeLog = &v
}

// GetIsRunning returns the IsRunning field value if set, zero value otherwise.
func (o *Task) GetIsRunning() bool {
	if o == nil || IsNil(o.IsRunning) {
		var ret bool
		return ret
	}
	return *o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIsRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRunning) {
		return nil, false
	}
	return o.IsRunning, true
}

// HasIsRunning returns a boolean if a field has been set.
func (o *Task) HasIsRunning() bool {
	if o != nil && !IsNil(o.IsRunning) {
		return true
	}

	return false
}

// SetIsRunning gets a reference to the given bool and assigns it to the IsRunning field.
func (o *Task) SetIsRunning(v bool) {
	o.IsRunning = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Task) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Task) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Task) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetTaskStatusId returns the TaskStatusId field value if set, zero value otherwise.
func (o *Task) GetTaskStatusId() string {
	if o == nil || IsNil(o.TaskStatusId) {
		var ret string
		return ret
	}
	return *o.TaskStatusId
}

// GetTaskStatusIdOk returns a tuple with the TaskStatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTaskStatusIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskStatusId) {
		return nil, false
	}
	return o.TaskStatusId, true
}

// HasTaskStatusId returns a boolean if a field has been set.
func (o *Task) HasTaskStatusId() bool {
	if o != nil && !IsNil(o.TaskStatusId) {
		return true
	}

	return false
}

// SetTaskStatusId gets a reference to the given string and assigns it to the TaskStatusId field.
func (o *Task) SetTaskStatusId(v string) {
	o.TaskStatusId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Task) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Task) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Task) SetDescription(v string) {
	o.Description = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Task) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Task) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *Task) SetDuration(v int32) {
	o.Duration = &v
}

// GetTaskStatusOrder returns the TaskStatusOrder field value if set, zero value otherwise.
func (o *Task) GetTaskStatusOrder() int32 {
	if o == nil || IsNil(o.TaskStatusOrder) {
		var ret int32
		return ret
	}
	return *o.TaskStatusOrder
}

// GetTaskStatusOrderOk returns a tuple with the TaskStatusOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTaskStatusOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskStatusOrder) {
		return nil, false
	}
	return o.TaskStatusOrder, true
}

// HasTaskStatusOrder returns a boolean if a field has been set.
func (o *Task) HasTaskStatusOrder() bool {
	if o != nil && !IsNil(o.TaskStatusOrder) {
		return true
	}

	return false
}

// SetTaskStatusOrder gets a reference to the given int32 and assigns it to the TaskStatusOrder field.
func (o *Task) SetTaskStatusOrder(v int32) {
	o.TaskStatusOrder = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *Task) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *Task) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *Task) SetRate(v float32) {
	o.Rate = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *Task) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Task) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *Task) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *Task) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Task) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *Task) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *Task) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Task) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *Task) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *Task) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Task) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *Task) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetIsDateBased returns the IsDateBased field value if set, zero value otherwise.
func (o *Task) GetIsDateBased() bool {
	if o == nil || IsNil(o.IsDateBased) {
		var ret bool
		return ret
	}
	return *o.IsDateBased
}

// GetIsDateBasedOk returns a tuple with the IsDateBased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIsDateBasedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDateBased) {
		return nil, false
	}
	return o.IsDateBased, true
}

// HasIsDateBased returns a boolean if a field has been set.
func (o *Task) HasIsDateBased() bool {
	if o != nil && !IsNil(o.IsDateBased) {
		return true
	}

	return false
}

// SetIsDateBased gets a reference to the given bool and assigns it to the IsDateBased field.
func (o *Task) SetIsDateBased(v bool) {
	o.IsDateBased = &v
}

// GetCalculatedStartDate returns the CalculatedStartDate field value if set, zero value otherwise.
func (o *Task) GetCalculatedStartDate() string {
	if o == nil || IsNil(o.CalculatedStartDate) {
		var ret string
		return ret
	}
	return *o.CalculatedStartDate
}

// GetCalculatedStartDateOk returns a tuple with the CalculatedStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCalculatedStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.CalculatedStartDate) {
		return nil, false
	}
	return o.CalculatedStartDate, true
}

// HasCalculatedStartDate returns a boolean if a field has been set.
func (o *Task) HasCalculatedStartDate() bool {
	if o != nil && !IsNil(o.CalculatedStartDate) {
		return true
	}

	return false
}

// SetCalculatedStartDate gets a reference to the given string and assigns it to the CalculatedStartDate field.
func (o *Task) SetCalculatedStartDate(v string) {
	o.CalculatedStartDate = &v
}

// GetInvoiceDocuments returns the InvoiceDocuments field value if set, zero value otherwise.
func (o *Task) GetInvoiceDocuments() bool {
	if o == nil || IsNil(o.InvoiceDocuments) {
		var ret bool
		return ret
	}
	return *o.InvoiceDocuments
}

// GetInvoiceDocumentsOk returns a tuple with the InvoiceDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetInvoiceDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceDocuments) {
		return nil, false
	}
	return o.InvoiceDocuments, true
}

// HasInvoiceDocuments returns a boolean if a field has been set.
func (o *Task) HasInvoiceDocuments() bool {
	if o != nil && !IsNil(o.InvoiceDocuments) {
		return true
	}

	return false
}

// SetInvoiceDocuments gets a reference to the given bool and assigns it to the InvoiceDocuments field.
func (o *Task) SetInvoiceDocuments(v bool) {
	o.InvoiceDocuments = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Task) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Task) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *Task) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Task) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Task) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Task) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Task) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Task) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *Task) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Task) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.TimeLog) {
		toSerialize["time_log"] = o.TimeLog
	}
	if !IsNil(o.IsRunning) {
		toSerialize["is_running"] = o.IsRunning
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.TaskStatusId) {
		toSerialize["task_status_id"] = o.TaskStatusId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.TaskStatusOrder) {
		toSerialize["task_status_order"] = o.TaskStatusOrder
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
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
	if !IsNil(o.IsDateBased) {
		toSerialize["is_date_based"] = o.IsDateBased
	}
	if !IsNil(o.CalculatedStartDate) {
		toSerialize["calculated_start_date"] = o.CalculatedStartDate
	}
	if !IsNil(o.InvoiceDocuments) {
		toSerialize["invoice_documents"] = o.InvoiceDocuments
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	return toSerialize, nil
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


