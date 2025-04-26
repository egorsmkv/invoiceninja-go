# UpdateTaskSchedulerSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paused** | Pointer to **bool** | The scheduler paused state | [optional] 
**RepeatEvery** | Pointer to **string** | Accepted values (DAY,WEEK,MONTH,3MONTHS,YEAR) | [optional] 
**StartFrom** | Pointer to **int32** | Timestamp when we should start the scheduler, default is today | [optional] 
**Job** | Pointer to **string** | Job, we can find list of available jobs in Scheduler model | [optional] 
**DateRange** | Pointer to **string** | The string representation of the date range of data to be returned | [optional] 
**DateKey** | Pointer to **string** | The date column to search between. | [optional] 
**StartDate** | Pointer to **string** | The start date to search between | [optional] 
**EndDate** | Pointer to **string** | The end date to search between | [optional] 

## Methods

### NewUpdateTaskSchedulerSchema

`func NewUpdateTaskSchedulerSchema() *UpdateTaskSchedulerSchema`

NewUpdateTaskSchedulerSchema instantiates a new UpdateTaskSchedulerSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskSchedulerSchemaWithDefaults

`func NewUpdateTaskSchedulerSchemaWithDefaults() *UpdateTaskSchedulerSchema`

NewUpdateTaskSchedulerSchemaWithDefaults instantiates a new UpdateTaskSchedulerSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaused

`func (o *UpdateTaskSchedulerSchema) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *UpdateTaskSchedulerSchema) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *UpdateTaskSchedulerSchema) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *UpdateTaskSchedulerSchema) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetRepeatEvery

`func (o *UpdateTaskSchedulerSchema) GetRepeatEvery() string`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *UpdateTaskSchedulerSchema) GetRepeatEveryOk() (*string, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *UpdateTaskSchedulerSchema) SetRepeatEvery(v string)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *UpdateTaskSchedulerSchema) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### GetStartFrom

`func (o *UpdateTaskSchedulerSchema) GetStartFrom() int32`

GetStartFrom returns the StartFrom field if non-nil, zero value otherwise.

### GetStartFromOk

`func (o *UpdateTaskSchedulerSchema) GetStartFromOk() (*int32, bool)`

GetStartFromOk returns a tuple with the StartFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFrom

`func (o *UpdateTaskSchedulerSchema) SetStartFrom(v int32)`

SetStartFrom sets StartFrom field to given value.

### HasStartFrom

`func (o *UpdateTaskSchedulerSchema) HasStartFrom() bool`

HasStartFrom returns a boolean if a field has been set.

### GetJob

`func (o *UpdateTaskSchedulerSchema) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *UpdateTaskSchedulerSchema) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *UpdateTaskSchedulerSchema) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *UpdateTaskSchedulerSchema) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetDateRange

`func (o *UpdateTaskSchedulerSchema) GetDateRange() string`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *UpdateTaskSchedulerSchema) GetDateRangeOk() (*string, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *UpdateTaskSchedulerSchema) SetDateRange(v string)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *UpdateTaskSchedulerSchema) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDateKey

`func (o *UpdateTaskSchedulerSchema) GetDateKey() string`

GetDateKey returns the DateKey field if non-nil, zero value otherwise.

### GetDateKeyOk

`func (o *UpdateTaskSchedulerSchema) GetDateKeyOk() (*string, bool)`

GetDateKeyOk returns a tuple with the DateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateKey

`func (o *UpdateTaskSchedulerSchema) SetDateKey(v string)`

SetDateKey sets DateKey field to given value.

### HasDateKey

`func (o *UpdateTaskSchedulerSchema) HasDateKey() bool`

HasDateKey returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateTaskSchedulerSchema) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateTaskSchedulerSchema) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateTaskSchedulerSchema) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateTaskSchedulerSchema) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateTaskSchedulerSchema) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateTaskSchedulerSchema) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateTaskSchedulerSchema) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateTaskSchedulerSchema) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


