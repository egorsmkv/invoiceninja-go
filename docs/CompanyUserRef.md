# CompanyUserRef

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

### NewCompanyUserRef

`func NewCompanyUserRef() *CompanyUserRef`

NewCompanyUserRef instantiates a new CompanyUserRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyUserRefWithDefaults

`func NewCompanyUserRefWithDefaults() *CompanyUserRef`

NewCompanyUserRefWithDefaults instantiates a new CompanyUserRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *CompanyUserRef) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CompanyUserRef) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CompanyUserRef) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CompanyUserRef) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSettings

`func (o *CompanyUserRef) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CompanyUserRef) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CompanyUserRef) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CompanyUserRef) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetReactSettings

`func (o *CompanyUserRef) GetReactSettings() map[string]interface{}`

GetReactSettings returns the ReactSettings field if non-nil, zero value otherwise.

### GetReactSettingsOk

`func (o *CompanyUserRef) GetReactSettingsOk() (*map[string]interface{}, bool)`

GetReactSettingsOk returns a tuple with the ReactSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactSettings

`func (o *CompanyUserRef) SetReactSettings(v map[string]interface{})`

SetReactSettings sets ReactSettings field to given value.

### HasReactSettings

`func (o *CompanyUserRef) HasReactSettings() bool`

HasReactSettings returns a boolean if a field has been set.

### GetIsOwner

`func (o *CompanyUserRef) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *CompanyUserRef) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *CompanyUserRef) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *CompanyUserRef) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsAdmin

`func (o *CompanyUserRef) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CompanyUserRef) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CompanyUserRef) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CompanyUserRef) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsLocked

`func (o *CompanyUserRef) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *CompanyUserRef) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *CompanyUserRef) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *CompanyUserRef) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CompanyUserRef) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyUserRef) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyUserRef) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CompanyUserRef) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CompanyUserRef) GetDeletedAt() int32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CompanyUserRef) GetDeletedAtOk() (*int32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CompanyUserRef) SetDeletedAt(v int32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CompanyUserRef) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetAccount

`func (o *CompanyUserRef) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CompanyUserRef) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CompanyUserRef) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CompanyUserRef) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyUserRef) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyUserRef) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyUserRef) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyUserRef) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetUser

`func (o *CompanyUserRef) GetUser() UserRef`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CompanyUserRef) GetUserOk() (*UserRef, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CompanyUserRef) SetUser(v UserRef)`

SetUser sets User field to given value.

### HasUser

`func (o *CompanyUserRef) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetToken

`func (o *CompanyUserRef) GetToken() CompanyToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CompanyUserRef) GetTokenOk() (*CompanyToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CompanyUserRef) SetToken(v CompanyToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *CompanyUserRef) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


