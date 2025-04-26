# GroupSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The group setting hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**Name** | Pointer to **string** | The name of the group | [optional] 
**Settings** | Pointer to **map[string]interface{}** | The settings object | [optional] 

## Methods

### NewGroupSetting

`func NewGroupSetting() *GroupSetting`

NewGroupSetting instantiates a new GroupSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSettingWithDefaults

`func NewGroupSettingWithDefaults() *GroupSetting`

NewGroupSettingWithDefaults instantiates a new GroupSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupSetting) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *GroupSetting) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupSetting) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupSetting) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GroupSetting) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *GroupSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *GroupSetting) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GroupSetting) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GroupSetting) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GroupSetting) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


