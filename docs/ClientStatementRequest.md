# ClientStatementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | The start date of the statement period - format Y-m-d | [optional] 
**EndDate** | Pointer to **string** | The start date of the statement period - format Y-m-d | [optional] 
**ClientId** | Pointer to **string** | The hashed ID of the client | [optional] 
**ShowPaymentsTable** | Pointer to **bool** | Flag which determines if the payments table is shown | [optional] 
**ShowCreditsTable** | Pointer to **bool** | Flag which determines if the credits table is shown | [optional] 
**ShowAgingTable** | Pointer to **bool** | Flag which determines if the aging table is shown | [optional] 

## Methods

### NewClientStatementRequest

`func NewClientStatementRequest() *ClientStatementRequest`

NewClientStatementRequest instantiates a new ClientStatementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStatementRequestWithDefaults

`func NewClientStatementRequestWithDefaults() *ClientStatementRequest`

NewClientStatementRequestWithDefaults instantiates a new ClientStatementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ClientStatementRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ClientStatementRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ClientStatementRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ClientStatementRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ClientStatementRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ClientStatementRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ClientStatementRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ClientStatementRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetClientId

`func (o *ClientStatementRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientStatementRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientStatementRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientStatementRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetShowPaymentsTable

`func (o *ClientStatementRequest) GetShowPaymentsTable() bool`

GetShowPaymentsTable returns the ShowPaymentsTable field if non-nil, zero value otherwise.

### GetShowPaymentsTableOk

`func (o *ClientStatementRequest) GetShowPaymentsTableOk() (*bool, bool)`

GetShowPaymentsTableOk returns a tuple with the ShowPaymentsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPaymentsTable

`func (o *ClientStatementRequest) SetShowPaymentsTable(v bool)`

SetShowPaymentsTable sets ShowPaymentsTable field to given value.

### HasShowPaymentsTable

`func (o *ClientStatementRequest) HasShowPaymentsTable() bool`

HasShowPaymentsTable returns a boolean if a field has been set.

### GetShowCreditsTable

`func (o *ClientStatementRequest) GetShowCreditsTable() bool`

GetShowCreditsTable returns the ShowCreditsTable field if non-nil, zero value otherwise.

### GetShowCreditsTableOk

`func (o *ClientStatementRequest) GetShowCreditsTableOk() (*bool, bool)`

GetShowCreditsTableOk returns a tuple with the ShowCreditsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCreditsTable

`func (o *ClientStatementRequest) SetShowCreditsTable(v bool)`

SetShowCreditsTable sets ShowCreditsTable field to given value.

### HasShowCreditsTable

`func (o *ClientStatementRequest) HasShowCreditsTable() bool`

HasShowCreditsTable returns a boolean if a field has been set.

### GetShowAgingTable

`func (o *ClientStatementRequest) GetShowAgingTable() bool`

GetShowAgingTable returns the ShowAgingTable field if non-nil, zero value otherwise.

### GetShowAgingTableOk

`func (o *ClientStatementRequest) GetShowAgingTableOk() (*bool, bool)`

GetShowAgingTableOk returns a tuple with the ShowAgingTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAgingTable

`func (o *ClientStatementRequest) SetShowAgingTable(v bool)`

SetShowAgingTable sets ShowAgingTable field to given value.

### HasShowAgingTable

`func (o *ClientStatementRequest) HasShowAgingTable() bool`

HasShowAgingTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


