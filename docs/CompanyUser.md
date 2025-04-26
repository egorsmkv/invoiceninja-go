# CompanyUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **string** | The user permissions for this company in a comma separated list | [optional] 
**Settings** | Pointer to **map[string]interface{}** | Settings that are used for the flutter applications to store user preferences / metadata | [optional] [readonly] 
**ReactSettings** | Pointer to **map[string]interface{}** | Dedicated settings object for the react web application | [optional] [readonly] 
**IsOwner** | Pointer to **bool** | Determines whether the user owns this company | [optional] [readonly] 
**IsAdmin** | Pointer to **bool** | Determines whether the user is the admin of this company | [optional] [readonly] 
**IsLocked** | Pointer to **bool** | Determines whether the users access to this company has been locked | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | The last time the record was modified, format Unix Timestamp | [optional] 
**DeletedAt** | Pointer to **int32** | Timestamp when the user was archived, format Unix Timestamp | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**User** | Pointer to [**UserRef**](UserRef.md) |  | [optional] 
**Token** | Pointer to [**CompanyToken**](CompanyToken.md) |  | [optional] 

## Methods

### NewCompanyUser

`func NewCompanyUser() *CompanyUser`

NewCompanyUser instantiates a new CompanyUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyUserWithDefaults

`func NewCompanyUserWithDefaults() *CompanyUser`

NewCompanyUserWithDefaults instantiates a new CompanyUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *CompanyUser) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CompanyUser) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CompanyUser) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CompanyUser) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSettings

`func (o *CompanyUser) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CompanyUser) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CompanyUser) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CompanyUser) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetReactSettings

`func (o *CompanyUser) GetReactSettings() map[string]interface{}`

GetReactSettings returns the ReactSettings field if non-nil, zero value otherwise.

### GetReactSettingsOk

`func (o *CompanyUser) GetReactSettingsOk() (*map[string]interface{}, bool)`

GetReactSettingsOk returns a tuple with the ReactSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactSettings

`func (o *CompanyUser) SetReactSettings(v map[string]interface{})`

SetReactSettings sets ReactSettings field to given value.

### HasReactSettings

`func (o *CompanyUser) HasReactSettings() bool`

HasReactSettings returns a boolean if a field has been set.

### GetIsOwner

`func (o *CompanyUser) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *CompanyUser) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *CompanyUser) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *CompanyUser) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsAdmin

`func (o *CompanyUser) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CompanyUser) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CompanyUser) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CompanyUser) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsLocked

`func (o *CompanyUser) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *CompanyUser) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *CompanyUser) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *CompanyUser) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CompanyUser) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyUser) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyUser) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CompanyUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CompanyUser) GetDeletedAt() int32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CompanyUser) GetDeletedAtOk() (*int32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CompanyUser) SetDeletedAt(v int32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CompanyUser) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAccount

`func (o *CompanyUser) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CompanyUser) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CompanyUser) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CompanyUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyUser) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyUser) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyUser) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyUser) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetUser

`func (o *CompanyUser) GetUser() UserRef`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CompanyUser) GetUserOk() (*UserRef, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CompanyUser) SetUser(v UserRef)`

SetUser sets User field to given value.

### HasUser

`func (o *CompanyUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetToken

`func (o *CompanyUser) GetToken() CompanyToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CompanyUser) GetTokenOk() (*CompanyToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CompanyUser) SetToken(v CompanyToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *CompanyUser) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


