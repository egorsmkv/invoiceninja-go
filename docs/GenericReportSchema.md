# GenericReportSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | Pointer to **string** | The string representation of the date range of data to be returned | [optional] 
**DateKey** | Pointer to **string** | The date column to search between. | [optional] 
**StartDate** | Pointer to **string** | The start date to search between | [optional] 
**EndDate** | Pointer to **string** | The end date to search between | [optional] 
**ReportKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGenericReportSchema

`func NewGenericReportSchema() *GenericReportSchema`

NewGenericReportSchema instantiates a new GenericReportSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericReportSchemaWithDefaults

`func NewGenericReportSchemaWithDefaults() *GenericReportSchema`

NewGenericReportSchemaWithDefaults instantiates a new GenericReportSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *GenericReportSchema) GetDateRange() string`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *GenericReportSchema) GetDateRangeOk() (*string, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *GenericReportSchema) SetDateRange(v string)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *GenericReportSchema) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDateKey

`func (o *GenericReportSchema) GetDateKey() string`

GetDateKey returns the DateKey field if non-nil, zero value otherwise.

### GetDateKeyOk

`func (o *GenericReportSchema) GetDateKeyOk() (*string, bool)`

GetDateKeyOk returns a tuple with the DateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateKey

`func (o *GenericReportSchema) SetDateKey(v string)`

SetDateKey sets DateKey field to given value.

### HasDateKey

`func (o *GenericReportSchema) HasDateKey() bool`

HasDateKey returns a boolean if a field has been set.

### GetStartDate

`func (o *GenericReportSchema) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GenericReportSchema) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GenericReportSchema) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GenericReportSchema) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GenericReportSchema) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GenericReportSchema) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GenericReportSchema) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GenericReportSchema) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetReportKeys

`func (o *GenericReportSchema) GetReportKeys() []string`

GetReportKeys returns the ReportKeys field if non-nil, zero value otherwise.

### GetReportKeysOk

`func (o *GenericReportSchema) GetReportKeysOk() (*[]string, bool)`

GetReportKeysOk returns a tuple with the ReportKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportKeys

`func (o *GenericReportSchema) SetReportKeys(v []string)`

SetReportKeys sets ReportKeys field to given value.

### HasReportKeys

`func (o *GenericReportSchema) HasReportKeys() bool`

HasReportKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


