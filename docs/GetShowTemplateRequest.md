# GetShowTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | The email template subject | [optional] 
**Body** | Pointer to **string** | The email template body | [optional] 

## Methods

### NewGetShowTemplateRequest

`func NewGetShowTemplateRequest() *GetShowTemplateRequest`

NewGetShowTemplateRequest instantiates a new GetShowTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShowTemplateRequestWithDefaults

`func NewGetShowTemplateRequestWithDefaults() *GetShowTemplateRequest`

NewGetShowTemplateRequestWithDefaults instantiates a new GetShowTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *GetShowTemplateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetShowTemplateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetShowTemplateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetShowTemplateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBody

`func (o *GetShowTemplateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetShowTemplateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetShowTemplateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetShowTemplateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


