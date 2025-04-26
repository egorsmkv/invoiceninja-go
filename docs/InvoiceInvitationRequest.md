# InvoiceInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The entity invitation hashed id | [optional] [readonly] 
**ClientContactId** | **string** | The client contact hashed id | 
**Key** | Pointer to **string** | The invitation key | [optional] [readonly] 
**Link** | Pointer to **string** | The invitation link | [optional] [readonly] 
**SentDate** | Pointer to **time.Time** | The invitation sent date | [optional] [readonly] 
**ViewedDate** | Pointer to **time.Time** | The invitation viewed date | [optional] [readonly] 
**OpenedDate** | Pointer to **time.Time** | The invitation opened date | [optional] [readonly] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] [readonly] 
**EmailError** | Pointer to **string** | The email error | [optional] [readonly] 
**EmailStatus** | Pointer to **string** | The email status | [optional] [readonly] 

## Methods

### NewInvoiceInvitationRequest

`func NewInvoiceInvitationRequest(clientContactId string, ) *InvoiceInvitationRequest`

NewInvoiceInvitationRequest instantiates a new InvoiceInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceInvitationRequestWithDefaults

`func NewInvoiceInvitationRequestWithDefaults() *InvoiceInvitationRequest`

NewInvoiceInvitationRequestWithDefaults instantiates a new InvoiceInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceInvitationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceInvitationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceInvitationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceInvitationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientContactId

`func (o *InvoiceInvitationRequest) GetClientContactId() string`

GetClientContactId returns the ClientContactId field if non-nil, zero value otherwise.

### GetClientContactIdOk

`func (o *InvoiceInvitationRequest) GetClientContactIdOk() (*string, bool)`

GetClientContactIdOk returns a tuple with the ClientContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContactId

`func (o *InvoiceInvitationRequest) SetClientContactId(v string)`

SetClientContactId sets ClientContactId field to given value.


### GetKey

`func (o *InvoiceInvitationRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InvoiceInvitationRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InvoiceInvitationRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InvoiceInvitationRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLink

`func (o *InvoiceInvitationRequest) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *InvoiceInvitationRequest) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *InvoiceInvitationRequest) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *InvoiceInvitationRequest) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSentDate

`func (o *InvoiceInvitationRequest) GetSentDate() time.Time`

GetSentDate returns the SentDate field if non-nil, zero value otherwise.

### GetSentDateOk

`func (o *InvoiceInvitationRequest) GetSentDateOk() (*time.Time, bool)`

GetSentDateOk returns a tuple with the SentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDate

`func (o *InvoiceInvitationRequest) SetSentDate(v time.Time)`

SetSentDate sets SentDate field to given value.

### HasSentDate

`func (o *InvoiceInvitationRequest) HasSentDate() bool`

HasSentDate returns a boolean if a field has been set.

### GetViewedDate

`func (o *InvoiceInvitationRequest) GetViewedDate() time.Time`

GetViewedDate returns the ViewedDate field if non-nil, zero value otherwise.

### GetViewedDateOk

`func (o *InvoiceInvitationRequest) GetViewedDateOk() (*time.Time, bool)`

GetViewedDateOk returns a tuple with the ViewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedDate

`func (o *InvoiceInvitationRequest) SetViewedDate(v time.Time)`

SetViewedDate sets ViewedDate field to given value.

### HasViewedDate

`func (o *InvoiceInvitationRequest) HasViewedDate() bool`

HasViewedDate returns a boolean if a field has been set.

### GetOpenedDate

`func (o *InvoiceInvitationRequest) GetOpenedDate() time.Time`

GetOpenedDate returns the OpenedDate field if non-nil, zero value otherwise.

### GetOpenedDateOk

`func (o *InvoiceInvitationRequest) GetOpenedDateOk() (*time.Time, bool)`

GetOpenedDateOk returns a tuple with the OpenedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedDate

`func (o *InvoiceInvitationRequest) SetOpenedDate(v time.Time)`

SetOpenedDate sets OpenedDate field to given value.

### HasOpenedDate

`func (o *InvoiceInvitationRequest) HasOpenedDate() bool`

HasOpenedDate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InvoiceInvitationRequest) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceInvitationRequest) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceInvitationRequest) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvoiceInvitationRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *InvoiceInvitationRequest) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *InvoiceInvitationRequest) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *InvoiceInvitationRequest) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *InvoiceInvitationRequest) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetEmailError

`func (o *InvoiceInvitationRequest) GetEmailError() string`

GetEmailError returns the EmailError field if non-nil, zero value otherwise.

### GetEmailErrorOk

`func (o *InvoiceInvitationRequest) GetEmailErrorOk() (*string, bool)`

GetEmailErrorOk returns a tuple with the EmailError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailError

`func (o *InvoiceInvitationRequest) SetEmailError(v string)`

SetEmailError sets EmailError field to given value.

### HasEmailError

`func (o *InvoiceInvitationRequest) HasEmailError() bool`

HasEmailError returns a boolean if a field has been set.

### GetEmailStatus

`func (o *InvoiceInvitationRequest) GetEmailStatus() string`

GetEmailStatus returns the EmailStatus field if non-nil, zero value otherwise.

### GetEmailStatusOk

`func (o *InvoiceInvitationRequest) GetEmailStatusOk() (*string, bool)`

GetEmailStatusOk returns a tuple with the EmailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStatus

`func (o *InvoiceInvitationRequest) SetEmailStatus(v string)`

SetEmailStatus sets EmailStatus field to given value.

### HasEmailStatus

`func (o *InvoiceInvitationRequest) HasEmailStatus() bool`

HasEmailStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


