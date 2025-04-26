# SendEmailTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | The email subject | [optional] 
**Body** | Pointer to **string** | The email body | [optional] 
**Entity** | **string** | The entity name | 
**EntityId** | **string** | The entity_id | 
**CcEmail** | Pointer to **string** | The email address of a user to be CC&#39;d on the email | [optional] 
**Template** | **string** | The template required | 

## Methods

### NewSendEmailTemplateRequest

`func NewSendEmailTemplateRequest(entity string, entityId string, template string, ) *SendEmailTemplateRequest`

NewSendEmailTemplateRequest instantiates a new SendEmailTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailTemplateRequestWithDefaults

`func NewSendEmailTemplateRequestWithDefaults() *SendEmailTemplateRequest`

NewSendEmailTemplateRequestWithDefaults instantiates a new SendEmailTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *SendEmailTemplateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendEmailTemplateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendEmailTemplateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SendEmailTemplateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBody

`func (o *SendEmailTemplateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendEmailTemplateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendEmailTemplateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SendEmailTemplateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEntity

`func (o *SendEmailTemplateRequest) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *SendEmailTemplateRequest) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *SendEmailTemplateRequest) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetEntityId

`func (o *SendEmailTemplateRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SendEmailTemplateRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SendEmailTemplateRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetCcEmail

`func (o *SendEmailTemplateRequest) GetCcEmail() string`

GetCcEmail returns the CcEmail field if non-nil, zero value otherwise.

### GetCcEmailOk

`func (o *SendEmailTemplateRequest) GetCcEmailOk() (*string, bool)`

GetCcEmailOk returns a tuple with the CcEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmail

`func (o *SendEmailTemplateRequest) SetCcEmail(v string)`

SetCcEmail sets CcEmail field to given value.

### HasCcEmail

`func (o *SendEmailTemplateRequest) HasCcEmail() bool`

HasCcEmail returns a boolean if a field has been set.

### GetTemplate

`func (o *SendEmailTemplateRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SendEmailTemplateRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SendEmailTemplateRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


