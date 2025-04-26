# SystemLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The account hashed id | [optional] 
**UserId** | Pointer to **string** | The user_id hashed id | [optional] 
**ClientId** | Pointer to **string** | The client_id hashed id | [optional] 
**EventId** | Pointer to **int32** | The Log Type ID | [optional] 
**CategoryId** | Pointer to **int32** | The Category Type ID | [optional] 
**TypeId** | Pointer to **int32** | The Type Type ID | [optional] 
**Log** | Pointer to **map[string]interface{}** | The json object of the error | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp | [optional] 
**CreatedAt** | Pointer to **string** | Timestamp | [optional] 

## Methods

### NewSystemLog

`func NewSystemLog() *SystemLog`

NewSystemLog instantiates a new SystemLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemLogWithDefaults

`func NewSystemLogWithDefaults() *SystemLog`

NewSystemLogWithDefaults instantiates a new SystemLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SystemLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *SystemLog) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SystemLog) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SystemLog) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SystemLog) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetClientId

`func (o *SystemLog) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SystemLog) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SystemLog) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SystemLog) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetEventId

`func (o *SystemLog) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SystemLog) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SystemLog) SetEventId(v int32)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *SystemLog) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetCategoryId

`func (o *SystemLog) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *SystemLog) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *SystemLog) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *SystemLog) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetTypeId

`func (o *SystemLog) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *SystemLog) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *SystemLog) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *SystemLog) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetLog

`func (o *SystemLog) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SystemLog) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SystemLog) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *SystemLog) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SystemLog) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SystemLog) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SystemLog) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SystemLog) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SystemLog) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SystemLog) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SystemLog) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SystemLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


