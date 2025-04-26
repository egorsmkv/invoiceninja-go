# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the subscription | [optional] 
**UserId** | Pointer to **string** | Unique identifier for the user associated with the subscription | [optional] 
**ProductId** | Pointer to **string** | Unique identifier for the product associated with the subscription | [optional] 
**RecurringInvoiceId** | Pointer to **string** | Unique identifier for the recurring invoice associated with the subscription | [optional] 
**IsRecurring** | Pointer to **bool** | Indicates whether the subscription is recurring | [optional] 
**FrequencyId** | Pointer to **string** | integer const representation of the frequency | [optional] 
**AutoBill** | Pointer to **string** | enum setting | [optional] 
**PromoCode** | Pointer to **string** | Promotional code applied to the subscription | [optional] 
**PromoDiscount** | Pointer to **float32** | Discount percentage or amount applied to the subscription | [optional] 
**IsAmountDiscount** | Pointer to **bool** | Indicates whether the discount is a fixed amount | [optional] 
**AllowCancellation** | Pointer to **bool** | Indicates whether the subscription can be cancelled | [optional] 
**PerSeatEnabled** | Pointer to **bool** | Indicates whether the subscription pricing is per seat | [optional] 
**CurrencyId** | Pointer to **int32** | Unique identifier for the currency used in the subscription | [optional] 
**MaxSeatsLimit** | Pointer to **int32** | Maximum number of seats allowed for the subscription | [optional] 
**TrialEnabled** | Pointer to **bool** | Indicates whether the subscription has a trial period | [optional] 
**TrialDuration** | Pointer to **int32** | Duration of the trial period in days | [optional] 
**AllowQueryOverrides** | Pointer to **bool** | Indicates whether query overrides are allowed for the subscription | [optional] 
**AllowPlanChanges** | Pointer to **bool** | Indicates whether plan changes are allowed for the subscription | [optional] 
**RefundPeriod** | Pointer to **int32** | Number of days within which refunds can be requested | [optional] 
**WebhookConfiguration** | Pointer to **string** | Webhook configuration for the subscription | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates whether the subscription has been deleted | [optional] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] 
**CreatedAt** | Pointer to **float32** | Timestamp | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Subscription) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Subscription) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Subscription) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Subscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProductId

`func (o *Subscription) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Subscription) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Subscription) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Subscription) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetRecurringInvoiceId

`func (o *Subscription) GetRecurringInvoiceId() string`

GetRecurringInvoiceId returns the RecurringInvoiceId field if non-nil, zero value otherwise.

### GetRecurringInvoiceIdOk

`func (o *Subscription) GetRecurringInvoiceIdOk() (*string, bool)`

GetRecurringInvoiceIdOk returns a tuple with the RecurringInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInvoiceId

`func (o *Subscription) SetRecurringInvoiceId(v string)`

SetRecurringInvoiceId sets RecurringInvoiceId field to given value.

### HasRecurringInvoiceId

`func (o *Subscription) HasRecurringInvoiceId() bool`

HasRecurringInvoiceId returns a boolean if a field has been set.

### GetIsRecurring

`func (o *Subscription) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *Subscription) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *Subscription) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *Subscription) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### GetFrequencyId

`func (o *Subscription) GetFrequencyId() string`

GetFrequencyId returns the FrequencyId field if non-nil, zero value otherwise.

### GetFrequencyIdOk

`func (o *Subscription) GetFrequencyIdOk() (*string, bool)`

GetFrequencyIdOk returns a tuple with the FrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyId

`func (o *Subscription) SetFrequencyId(v string)`

SetFrequencyId sets FrequencyId field to given value.

### HasFrequencyId

`func (o *Subscription) HasFrequencyId() bool`

HasFrequencyId returns a boolean if a field has been set.

### GetAutoBill

`func (o *Subscription) GetAutoBill() string`

GetAutoBill returns the AutoBill field if non-nil, zero value otherwise.

### GetAutoBillOk

`func (o *Subscription) GetAutoBillOk() (*string, bool)`

GetAutoBillOk returns a tuple with the AutoBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBill

`func (o *Subscription) SetAutoBill(v string)`

SetAutoBill sets AutoBill field to given value.

### HasAutoBill

`func (o *Subscription) HasAutoBill() bool`

HasAutoBill returns a boolean if a field has been set.

### GetPromoCode

`func (o *Subscription) GetPromoCode() string`

GetPromoCode returns the PromoCode field if non-nil, zero value otherwise.

### GetPromoCodeOk

`func (o *Subscription) GetPromoCodeOk() (*string, bool)`

GetPromoCodeOk returns a tuple with the PromoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCode

`func (o *Subscription) SetPromoCode(v string)`

SetPromoCode sets PromoCode field to given value.

### HasPromoCode

`func (o *Subscription) HasPromoCode() bool`

HasPromoCode returns a boolean if a field has been set.

### GetPromoDiscount

`func (o *Subscription) GetPromoDiscount() float32`

GetPromoDiscount returns the PromoDiscount field if non-nil, zero value otherwise.

### GetPromoDiscountOk

`func (o *Subscription) GetPromoDiscountOk() (*float32, bool)`

GetPromoDiscountOk returns a tuple with the PromoDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoDiscount

`func (o *Subscription) SetPromoDiscount(v float32)`

SetPromoDiscount sets PromoDiscount field to given value.

### HasPromoDiscount

`func (o *Subscription) HasPromoDiscount() bool`

HasPromoDiscount returns a boolean if a field has been set.

### GetIsAmountDiscount

`func (o *Subscription) GetIsAmountDiscount() bool`

GetIsAmountDiscount returns the IsAmountDiscount field if non-nil, zero value otherwise.

### GetIsAmountDiscountOk

`func (o *Subscription) GetIsAmountDiscountOk() (*bool, bool)`

GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountDiscount

`func (o *Subscription) SetIsAmountDiscount(v bool)`

SetIsAmountDiscount sets IsAmountDiscount field to given value.

### HasIsAmountDiscount

`func (o *Subscription) HasIsAmountDiscount() bool`

HasIsAmountDiscount returns a boolean if a field has been set.

### GetAllowCancellation

`func (o *Subscription) GetAllowCancellation() bool`

GetAllowCancellation returns the AllowCancellation field if non-nil, zero value otherwise.

### GetAllowCancellationOk

`func (o *Subscription) GetAllowCancellationOk() (*bool, bool)`

GetAllowCancellationOk returns a tuple with the AllowCancellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCancellation

`func (o *Subscription) SetAllowCancellation(v bool)`

SetAllowCancellation sets AllowCancellation field to given value.

### HasAllowCancellation

`func (o *Subscription) HasAllowCancellation() bool`

HasAllowCancellation returns a boolean if a field has been set.

### GetPerSeatEnabled

`func (o *Subscription) GetPerSeatEnabled() bool`

GetPerSeatEnabled returns the PerSeatEnabled field if non-nil, zero value otherwise.

### GetPerSeatEnabledOk

`func (o *Subscription) GetPerSeatEnabledOk() (*bool, bool)`

GetPerSeatEnabledOk returns a tuple with the PerSeatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerSeatEnabled

`func (o *Subscription) SetPerSeatEnabled(v bool)`

SetPerSeatEnabled sets PerSeatEnabled field to given value.

### HasPerSeatEnabled

`func (o *Subscription) HasPerSeatEnabled() bool`

HasPerSeatEnabled returns a boolean if a field has been set.

### GetCurrencyId

`func (o *Subscription) GetCurrencyId() int32`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *Subscription) GetCurrencyIdOk() (*int32, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *Subscription) SetCurrencyId(v int32)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *Subscription) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetMaxSeatsLimit

`func (o *Subscription) GetMaxSeatsLimit() int32`

GetMaxSeatsLimit returns the MaxSeatsLimit field if non-nil, zero value otherwise.

### GetMaxSeatsLimitOk

`func (o *Subscription) GetMaxSeatsLimitOk() (*int32, bool)`

GetMaxSeatsLimitOk returns a tuple with the MaxSeatsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSeatsLimit

`func (o *Subscription) SetMaxSeatsLimit(v int32)`

SetMaxSeatsLimit sets MaxSeatsLimit field to given value.

### HasMaxSeatsLimit

`func (o *Subscription) HasMaxSeatsLimit() bool`

HasMaxSeatsLimit returns a boolean if a field has been set.

### GetTrialEnabled

`func (o *Subscription) GetTrialEnabled() bool`

GetTrialEnabled returns the TrialEnabled field if non-nil, zero value otherwise.

### GetTrialEnabledOk

`func (o *Subscription) GetTrialEnabledOk() (*bool, bool)`

GetTrialEnabledOk returns a tuple with the TrialEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnabled

`func (o *Subscription) SetTrialEnabled(v bool)`

SetTrialEnabled sets TrialEnabled field to given value.

### HasTrialEnabled

`func (o *Subscription) HasTrialEnabled() bool`

HasTrialEnabled returns a boolean if a field has been set.

### GetTrialDuration

`func (o *Subscription) GetTrialDuration() int32`

GetTrialDuration returns the TrialDuration field if non-nil, zero value otherwise.

### GetTrialDurationOk

`func (o *Subscription) GetTrialDurationOk() (*int32, bool)`

GetTrialDurationOk returns a tuple with the TrialDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDuration

`func (o *Subscription) SetTrialDuration(v int32)`

SetTrialDuration sets TrialDuration field to given value.

### HasTrialDuration

`func (o *Subscription) HasTrialDuration() bool`

HasTrialDuration returns a boolean if a field has been set.

### GetAllowQueryOverrides

`func (o *Subscription) GetAllowQueryOverrides() bool`

GetAllowQueryOverrides returns the AllowQueryOverrides field if non-nil, zero value otherwise.

### GetAllowQueryOverridesOk

`func (o *Subscription) GetAllowQueryOverridesOk() (*bool, bool)`

GetAllowQueryOverridesOk returns a tuple with the AllowQueryOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryOverrides

`func (o *Subscription) SetAllowQueryOverrides(v bool)`

SetAllowQueryOverrides sets AllowQueryOverrides field to given value.

### HasAllowQueryOverrides

`func (o *Subscription) HasAllowQueryOverrides() bool`

HasAllowQueryOverrides returns a boolean if a field has been set.

### GetAllowPlanChanges

`func (o *Subscription) GetAllowPlanChanges() bool`

GetAllowPlanChanges returns the AllowPlanChanges field if non-nil, zero value otherwise.

### GetAllowPlanChangesOk

`func (o *Subscription) GetAllowPlanChangesOk() (*bool, bool)`

GetAllowPlanChangesOk returns a tuple with the AllowPlanChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPlanChanges

`func (o *Subscription) SetAllowPlanChanges(v bool)`

SetAllowPlanChanges sets AllowPlanChanges field to given value.

### HasAllowPlanChanges

`func (o *Subscription) HasAllowPlanChanges() bool`

HasAllowPlanChanges returns a boolean if a field has been set.

### GetRefundPeriod

`func (o *Subscription) GetRefundPeriod() int32`

GetRefundPeriod returns the RefundPeriod field if non-nil, zero value otherwise.

### GetRefundPeriodOk

`func (o *Subscription) GetRefundPeriodOk() (*int32, bool)`

GetRefundPeriodOk returns a tuple with the RefundPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPeriod

`func (o *Subscription) SetRefundPeriod(v int32)`

SetRefundPeriod sets RefundPeriod field to given value.

### HasRefundPeriod

`func (o *Subscription) HasRefundPeriod() bool`

HasRefundPeriod returns a boolean if a field has been set.

### GetWebhookConfiguration

`func (o *Subscription) GetWebhookConfiguration() string`

GetWebhookConfiguration returns the WebhookConfiguration field if non-nil, zero value otherwise.

### GetWebhookConfigurationOk

`func (o *Subscription) GetWebhookConfigurationOk() (*string, bool)`

GetWebhookConfigurationOk returns a tuple with the WebhookConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookConfiguration

`func (o *Subscription) SetWebhookConfiguration(v string)`

SetWebhookConfiguration sets WebhookConfiguration field to given value.

### HasWebhookConfiguration

`func (o *Subscription) HasWebhookConfiguration() bool`

HasWebhookConfiguration returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Subscription) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Subscription) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Subscription) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Subscription) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Subscription) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Subscription) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Subscription) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Subscription) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscription) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Subscription) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscription) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscription) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Subscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


