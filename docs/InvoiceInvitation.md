# InvoiceInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The entity invitation hashed id | [optional] [readonly] 
**ClientContactId** | Pointer to **string** | The client contact hashed id | [optional] 
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

### NewInvoiceInvitation

`func NewInvoiceInvitation() *InvoiceInvitation`

NewInvoiceInvitation instantiates a new InvoiceInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceInvitationWithDefaults

`func NewInvoiceInvitationWithDefaults() *InvoiceInvitation`

NewInvoiceInvitationWithDefaults instantiates a new InvoiceInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientContactId

`func (o *InvoiceInvitation) GetClientContactId() string`

GetClientContactId returns the ClientContactId field if non-nil, zero value otherwise.

### GetClientContactIdOk

`func (o *InvoiceInvitation) GetClientContactIdOk() (*string, bool)`

GetClientContactIdOk returns a tuple with the ClientContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContactId

`func (o *InvoiceInvitation) SetClientContactId(v string)`

SetClientContactId sets ClientContactId field to given value.

### HasClientContactId

`func (o *InvoiceInvitation) HasClientContactId() bool`

HasClientContactId returns a boolean if a field has been set.

### GetKey

`func (o *InvoiceInvitation) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InvoiceInvitation) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InvoiceInvitation) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *InvoiceInvitation) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLink

`func (o *InvoiceInvitation) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *InvoiceInvitation) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *InvoiceInvitation) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *InvoiceInvitation) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSentDate

`func (o *InvoiceInvitation) GetSentDate() time.Time`

GetSentDate returns the SentDate field if non-nil, zero value otherwise.

### GetSentDateOk

`func (o *InvoiceInvitation) GetSentDateOk() (*time.Time, bool)`

GetSentDateOk returns a tuple with the SentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDate

`func (o *InvoiceInvitation) SetSentDate(v time.Time)`

SetSentDate sets SentDate field to given value.

### HasSentDate

`func (o *InvoiceInvitation) HasSentDate() bool`

HasSentDate returns a boolean if a field has been set.

### GetViewedDate

`func (o *InvoiceInvitation) GetViewedDate() time.Time`

GetViewedDate returns the ViewedDate field if non-nil, zero value otherwise.

### GetViewedDateOk

`func (o *InvoiceInvitation) GetViewedDateOk() (*time.Time, bool)`

GetViewedDateOk returns a tuple with the ViewedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedDate

`func (o *InvoiceInvitation) SetViewedDate(v time.Time)`

SetViewedDate sets ViewedDate field to given value.

### HasViewedDate

`func (o *InvoiceInvitation) HasViewedDate() bool`

HasViewedDate returns a boolean if a field has been set.

### GetOpenedDate

`func (o *InvoiceInvitation) GetOpenedDate() time.Time`

GetOpenedDate returns the OpenedDate field if non-nil, zero value otherwise.

### GetOpenedDateOk

`func (o *InvoiceInvitation) GetOpenedDateOk() (*time.Time, bool)`

GetOpenedDateOk returns a tuple with the OpenedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedDate

`func (o *InvoiceInvitation) SetOpenedDate(v time.Time)`

SetOpenedDate sets OpenedDate field to given value.

### HasOpenedDate

`func (o *InvoiceInvitation) HasOpenedDate() bool`

HasOpenedDate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InvoiceInvitation) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceInvitation) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceInvitation) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvoiceInvitation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *InvoiceInvitation) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *InvoiceInvitation) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *InvoiceInvitation) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *InvoiceInvitation) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetEmailError

`func (o *InvoiceInvitation) GetEmailError() string`

GetEmailError returns the EmailError field if non-nil, zero value otherwise.

### GetEmailErrorOk

`func (o *InvoiceInvitation) GetEmailErrorOk() (*string, bool)`

GetEmailErrorOk returns a tuple with the EmailError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailError

`func (o *InvoiceInvitation) SetEmailError(v string)`

SetEmailError sets EmailError field to given value.

### HasEmailError

`func (o *InvoiceInvitation) HasEmailError() bool`

HasEmailError returns a boolean if a field has been set.

### GetEmailStatus

`func (o *InvoiceInvitation) GetEmailStatus() string`

GetEmailStatus returns the EmailStatus field if non-nil, zero value otherwise.

### GetEmailStatusOk

`func (o *InvoiceInvitation) GetEmailStatusOk() (*string, bool)`

GetEmailStatusOk returns a tuple with the EmailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStatus

`func (o *InvoiceInvitation) SetEmailStatus(v string)`

SetEmailStatus sets EmailStatus field to given value.

### HasEmailStatus

`func (o *InvoiceInvitation) HasEmailStatus() bool`

HasEmailStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


