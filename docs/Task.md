# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed id of the task | [optional] 
**UserId** | Pointer to **string** | The hashed id of the user who created the task | [optional] 
**AssignedUserId** | Pointer to **string** | The assigned user of the task | [optional] 
**ClientId** | Pointer to **string** | The hashed if of the client | [optional] 
**InvoiceId** | Pointer to **string** | The hashed id of the invoice associated with the task | [optional] 
**ProjectId** | Pointer to **string** | The hashed id of the project associated with the task | [optional] 
**Number** | Pointer to **string** | The number of the task | [optional] 
**TimeLog** | Pointer to **string** | An array of unix time stamps defining the start and end times of the task | [optional] 
**IsRunning** | Pointer to **bool** | Determines if the task is still running | [optional] 
**IsDeleted** | Pointer to **bool** | Boolean flag determining if the task has been deleted | [optional] 
**TaskStatusId** | Pointer to **string** | The hashed id of the task status | [optional] 
**Description** | Pointer to **string** | The task description | [optional] 
**Duration** | Pointer to **int32** | The task duration in seconds | [optional] 
**TaskStatusOrder** | Pointer to **int32** | The order of the task | [optional] 
**Rate** | Pointer to **float32** | The task rate | [optional] 
**CustomValue1** | Pointer to **string** | A custom value | [optional] 
**CustomValue2** | Pointer to **string** | A custom value | [optional] 
**CustomValue3** | Pointer to **string** | A custom value | [optional] 
**CustomValue4** | Pointer to **string** | A custom value | [optional] 
**IsDateBased** | Pointer to **bool** | Boolean flag determining if the task is date based | [optional] 
**CalculatedStartDate** | Pointer to **string** | The calculated start date of the task | [optional] [readonly] 
**InvoiceDocuments** | Pointer to **bool** | Boolean flags which determines whether to include the task documents on the invoice | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Task) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Task) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Task) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Task) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Task) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Task) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Task) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Task) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *Task) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Task) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Task) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Task) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Task) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Task) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Task) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Task) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetProjectId

`func (o *Task) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Task) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Task) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Task) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetNumber

`func (o *Task) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Task) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Task) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Task) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetTimeLog

`func (o *Task) GetTimeLog() string`

GetTimeLog returns the TimeLog field if non-nil, zero value otherwise.

### GetTimeLogOk

`func (o *Task) GetTimeLogOk() (*string, bool)`

GetTimeLogOk returns a tuple with the TimeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLog

`func (o *Task) SetTimeLog(v string)`

SetTimeLog sets TimeLog field to given value.

### HasTimeLog

`func (o *Task) HasTimeLog() bool`

HasTimeLog returns a boolean if a field has been set.

### GetIsRunning

`func (o *Task) GetIsRunning() bool`

GetIsRunning returns the IsRunning field if non-nil, zero value otherwise.

### GetIsRunningOk

`func (o *Task) GetIsRunningOk() (*bool, bool)`

GetIsRunningOk returns a tuple with the IsRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRunning

`func (o *Task) SetIsRunning(v bool)`

SetIsRunning sets IsRunning field to given value.

### HasIsRunning

`func (o *Task) HasIsRunning() bool`

HasIsRunning returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Task) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Task) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Task) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Task) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetTaskStatusId

`func (o *Task) GetTaskStatusId() string`

GetTaskStatusId returns the TaskStatusId field if non-nil, zero value otherwise.

### GetTaskStatusIdOk

`func (o *Task) GetTaskStatusIdOk() (*string, bool)`

GetTaskStatusIdOk returns a tuple with the TaskStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatusId

`func (o *Task) SetTaskStatusId(v string)`

SetTaskStatusId sets TaskStatusId field to given value.

### HasTaskStatusId

`func (o *Task) HasTaskStatusId() bool`

HasTaskStatusId returns a boolean if a field has been set.

### GetDescription

`func (o *Task) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Task) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Task) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Task) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *Task) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Task) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Task) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Task) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTaskStatusOrder

`func (o *Task) GetTaskStatusOrder() int32`

GetTaskStatusOrder returns the TaskStatusOrder field if non-nil, zero value otherwise.

### GetTaskStatusOrderOk

`func (o *Task) GetTaskStatusOrderOk() (*int32, bool)`

GetTaskStatusOrderOk returns a tuple with the TaskStatusOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatusOrder

`func (o *Task) SetTaskStatusOrder(v int32)`

SetTaskStatusOrder sets TaskStatusOrder field to given value.

### HasTaskStatusOrder

`func (o *Task) HasTaskStatusOrder() bool`

HasTaskStatusOrder returns a boolean if a field has been set.

### GetRate

`func (o *Task) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Task) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Task) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *Task) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCustomValue1

`func (o *Task) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *Task) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *Task) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *Task) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *Task) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *Task) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *Task) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *Task) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *Task) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *Task) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *Task) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *Task) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *Task) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *Task) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *Task) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *Task) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetIsDateBased

`func (o *Task) GetIsDateBased() bool`

GetIsDateBased returns the IsDateBased field if non-nil, zero value otherwise.

### GetIsDateBasedOk

`func (o *Task) GetIsDateBasedOk() (*bool, bool)`

GetIsDateBasedOk returns a tuple with the IsDateBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDateBased

`func (o *Task) SetIsDateBased(v bool)`

SetIsDateBased sets IsDateBased field to given value.

### HasIsDateBased

`func (o *Task) HasIsDateBased() bool`

HasIsDateBased returns a boolean if a field has been set.

### GetCalculatedStartDate

`func (o *Task) GetCalculatedStartDate() string`

GetCalculatedStartDate returns the CalculatedStartDate field if non-nil, zero value otherwise.

### GetCalculatedStartDateOk

`func (o *Task) GetCalculatedStartDateOk() (*string, bool)`

GetCalculatedStartDateOk returns a tuple with the CalculatedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedStartDate

`func (o *Task) SetCalculatedStartDate(v string)`

SetCalculatedStartDate sets CalculatedStartDate field to given value.

### HasCalculatedStartDate

`func (o *Task) HasCalculatedStartDate() bool`

HasCalculatedStartDate returns a boolean if a field has been set.

### GetInvoiceDocuments

`func (o *Task) GetInvoiceDocuments() bool`

GetInvoiceDocuments returns the InvoiceDocuments field if non-nil, zero value otherwise.

### GetInvoiceDocumentsOk

`func (o *Task) GetInvoiceDocumentsOk() (*bool, bool)`

GetInvoiceDocumentsOk returns a tuple with the InvoiceDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDocuments

`func (o *Task) SetInvoiceDocuments(v bool)`

SetInvoiceDocuments sets InvoiceDocuments field to given value.

### HasInvoiceDocuments

`func (o *Task) HasInvoiceDocuments() bool`

HasInvoiceDocuments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Task) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Task) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Task) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Task) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Task) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


