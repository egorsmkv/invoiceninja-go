# ClientSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyId** | **string** | The default currency id | 
**TimezoneId** | Pointer to **string** | The timezone id | [optional] 
**DateFormatId** | Pointer to **string** | The date format id | [optional] 
**MilitaryTime** | Pointer to **bool** | Toggles 12/24 hour time | [optional] 
**LanguageId** | Pointer to **string** | The language id | [optional] 
**ShowCurrencyCode** | Pointer to **bool** | Toggles whether the currency symbol or code is shown | [optional] 
**PaymentTerms** | Pointer to **int32** | -1 sets no payment term, 0 sets payment due immediately, positive integers indicates payment terms in days | [optional] 
**CompanyGatewayIds** | Pointer to **string** | A commad separate list of available gateways | [optional] 
**CustomValue1** | Pointer to **string** | A Custom Label | [optional] 
**CustomValue2** | Pointer to **string** | A Custom Label | [optional] 
**CustomValue3** | Pointer to **string** | A Custom Label | [optional] 
**CustomValue4** | Pointer to **string** | A Custom Label | [optional] 
**DefaultTaskRate** | Pointer to **float32** | The default task rate | [optional] 
**SendReminders** | Pointer to **bool** | Toggles whether reminders are sent | [optional] 
**EnableClientPortalTasks** | Pointer to **bool** | Show/hide the tasks panel in the client portal | [optional] 
**EmailStyle** | Pointer to **string** | options include plain,light,dark,custom | [optional] 
**ReplyToEmail** | Pointer to **string** | The reply to email address | [optional] 
**BccEmail** | Pointer to **string** | A comma separate list of BCC emails | [optional] 
**PdfEmailAttachment** | Pointer to **bool** | Toggles whether to attach PDF as attachment | [optional] 
**UblEmailAttachment** | Pointer to **bool** | Toggles whether to attach UBL as attachment | [optional] 
**EmailStyleCustom** | Pointer to **string** | The custom template | [optional] 
**CounterNumberApplied** | Pointer to **string** | enum when the invoice number counter is set, ie when_saved, when_sent, when_paid | [optional] 
**QuoteNumberApplied** | Pointer to **string** | enum when the quote number counter is set, ie when_saved, when_sent | [optional] 
**CustomMessageDashboard** | Pointer to **string** | A custom message which is displayed on the dashboard | [optional] 
**CustomMessageUnpaidInvoice** | Pointer to **string** | A custom message which is displayed in the client portal when a client is viewing a unpaid invoice. | [optional] 
**CustomMessagePaidInvoice** | Pointer to **string** | A custom message which is displayed in the client portal when a client is viewing a paid invoice. | [optional] 
**CustomMessageUnapprovedQuote** | Pointer to **string** | A custom message which is displayed in the client portal when a client is viewing a unapproved quote. | [optional] 
**LockInvoices** | Pointer to **bool** | Toggles whether invoices are locked once sent and cannot be modified further | [optional] 
**AutoArchiveInvoice** | Pointer to **bool** | Toggles whether a invoice is archived immediately following payment | [optional] 
**AutoArchiveQuote** | Pointer to **bool** | Toggles whether a quote is archived after being converted to a invoice | [optional] 
**AutoConvertQuote** | Pointer to **bool** | Toggles whether a quote is converted to a invoice when approved | [optional] 
**InclusiveTaxes** | Pointer to **bool** | Boolean flag determining whether inclusive or exclusive taxes are used | [optional] 
**TaskNumberPattern** | Pointer to **string** | Allows customisation of the task number pattern | [optional] 
**TaskNumberCounter** | Pointer to **int32** | The incrementing counter for tasks | [optional] 
**ReminderSendTime** | Pointer to **int32** | Time from UTC +0 when the email will be sent to the client | [optional] 
**ExpenseNumberPattern** | Pointer to **string** | Allows customisation of the expense number pattern | [optional] 
**ExpenseNumberCounter** | Pointer to **int32** | The incrementing counter for expenses | [optional] 
**VendorNumberPattern** | Pointer to **string** | Allows customisation of the vendor number pattern | [optional] 
**VendorNumberCounter** | Pointer to **int32** | The incrementing counter for vendors | [optional] 
**TicketNumberPattern** | Pointer to **string** | Allows customisation of the ticket number pattern | [optional] 
**TicketNumberCounter** | Pointer to **int32** | The incrementing counter for tickets | [optional] 
**PaymentNumberPattern** | Pointer to **string** | Allows customisation of the payment number pattern | [optional] 
**PaymentNumberCounter** | Pointer to **int32** | The incrementing counter for payments | [optional] 
**InvoiceNumberPattern** | Pointer to **string** | Allows customisation of the invoice number pattern | [optional] 
**InvoiceNumberCounter** | Pointer to **int32** | The incrementing counter for invoices | [optional] 
**QuoteNumberPattern** | Pointer to **string** | Allows customisation of the quote number pattern | [optional] 
**QuoteNumberCounter** | Pointer to **int32** | The incrementing counter for quotes | [optional] 
**ClientNumberPattern** | Pointer to **string** | Allows customisation of the client number pattern | [optional] 
**ClientNumberCounter** | Pointer to **int32** | The incrementing counter for clients | [optional] 
**CreditNumberPattern** | Pointer to **string** | Allows customisation of the credit number pattern | [optional] 
**CreditNumberCounter** | Pointer to **int32** | The incrementing counter for credits | [optional] 
**RecurringInvoiceNumberPrefix** | Pointer to **string** | This string is prepended to the recurring invoice number | [optional] 
**ResetCounterFrequencyId** | Pointer to **int32** | CONSTANT which is used to apply the frequency which the counters are reset | [optional] 
**ResetCounterDate** | Pointer to **string** | The explicit date which is used to reset counters | [optional] 
**CounterPadding** | Pointer to **int32** | Pads the counter with leading zeros | [optional] 
**SharedInvoiceQuoteCounter** | Pointer to **bool** | Flags whether to share the counter for invoices and quotes | [optional] 
**UpdateProducts** | Pointer to **bool** | Determines if client fields are updated from third party APIs | [optional] 
**ConvertProducts** | Pointer to **bool** |  | [optional] 
**FillProducts** | Pointer to **bool** | Automatically fill products based on product_key | [optional] 
**InvoiceTerms** | Pointer to **string** | The default invoice terms | [optional] 
**QuoteTerms** | Pointer to **string** | The default quote terms | [optional] 
**InvoiceTaxes** | Pointer to **float32** | Taxes can be applied to the invoice | [optional] 
**InvoiceDesignId** | Pointer to **string** | The default design id (invoice, quote etc) | [optional] 
**QuoteDesignId** | Pointer to **string** | The default design id (invoice, quote etc) | [optional] 
**InvoiceFooter** | Pointer to **string** | The default invoice footer | [optional] 
**InvoiceLabels** | Pointer to **string** | JSON string of invoice labels | [optional] 
**TaxRate1** | Pointer to **float32** | The tax rate (float) | [optional] 
**TaxName1** | Pointer to **string** | The tax name | [optional] 
**TaxRate2** | Pointer to **float32** | The tax rate (float) | [optional] 
**TaxName2** | Pointer to **string** | The tax name | [optional] 
**TaxRate3** | Pointer to **float32** | The tax rate (float) | [optional] 
**TaxName3** | Pointer to **string** | The tax name | [optional] 
**PaymentTypeId** | Pointer to **string** | The default payment type id | [optional] 
**CustomFields** | Pointer to **string** | JSON string of custom fields | [optional] 
**EmailFooter** | Pointer to **string** | The default email footer | [optional] 
**EmailSendingMethod** | Pointer to **string** | The email driver to use to send email, options include default, gmail | [optional] 
**GmailSendingUserId** | Pointer to **string** | The hashed_id of the user account to send email from | [optional] 
**EmailSubjectInvoice** | Pointer to **string** |  | [optional] 
**EmailSubjectQuote** | Pointer to **string** |  | [optional] 
**EmailSubjectPayment** | Pointer to **string** |  | [optional] 
**EmailTemplateInvoice** | Pointer to **string** | The full template for invoice emails | [optional] 
**EmailTemplateQuote** | Pointer to **string** | The full template for quote emails | [optional] 
**EmailTemplatePayment** | Pointer to **string** | The full template for payment emails | [optional] 
**EmailSubjectReminder1** | Pointer to **string** | Email subject for Reminder | [optional] 
**EmailSubjectReminder2** | Pointer to **string** | Email subject for Reminder | [optional] 
**EmailSubjectReminder3** | Pointer to **string** | Email subject for Reminder | [optional] 
**EmailSubjectReminderEndless** | Pointer to **string** | Email subject for endless reminders | [optional] 
**EmailTemplateReminder1** | Pointer to **string** | The full template for Reminder 1 | [optional] 
**EmailTemplateReminder2** | Pointer to **string** | The full template for Reminder 2 | [optional] 
**EmailTemplateReminder3** | Pointer to **string** | The full template for Reminder 3 | [optional] 
**EmailTemplateReminderEndless** | Pointer to **string** | The full template for enless reminders | [optional] 
**EnablePortalPassword** | Pointer to **bool** | Toggles whether a password is required to log into the client portal | [optional] 
**ShowAcceptInvoiceTerms** | Pointer to **bool** | Toggles whether the terms dialogue is shown to the client | [optional] 
**ShowAcceptQuoteTerms** | Pointer to **bool** | Toggles whether the terms dialogue is shown to the client | [optional] 
**RequireInvoiceSignature** | Pointer to **bool** | Toggles whether a invoice signature is required | [optional] 
**RequireQuoteSignature** | Pointer to **bool** | Toggles whether a quote signature is required | [optional] 
**Name** | Pointer to **string** | The company name | [optional] 
**CompanyLogo** | Pointer to **map[string]interface{}** | The company logo file | [optional] 
**Website** | Pointer to **string** | The company website URL | [optional] 
**Address1** | Pointer to **string** | The company address line 1 | [optional] 
**Address2** | Pointer to **string** | The company address line 2 | [optional] 
**City** | Pointer to **string** | The company city | [optional] 
**State** | Pointer to **string** | The company state | [optional] 
**PostalCode** | Pointer to **string** | The company zip/postal code | [optional] 
**Phone** | Pointer to **string** | The company phone | [optional] 
**Email** | Pointer to **string** | The company email | [optional] 
**CountryId** | Pointer to **string** | The country ID | [optional] 
**VatNumber** | Pointer to **string** | The company VAT/TAX ID number | [optional] 
**PageSize** | Pointer to **string** | The default page size | [optional] 
**FontSize** | Pointer to **float32** | The font size | [optional] 
**PrimaryFont** | Pointer to **string** | The primary font | [optional] 
**SecondaryFont** | Pointer to **string** | The secondary font | [optional] 
**HidePaidToDate** | Pointer to **bool** | Flags whether to hide the paid to date field | [optional] 
**EmbedDocuments** | Pointer to **bool** | Toggled whether to embed documents in the PDF | [optional] 
**AllPagesHeader** | Pointer to **bool** | The header for the PDF | [optional] 
**AllPagesFooter** | Pointer to **bool** | The footer for the PDF | [optional] 
**DocumentEmailAttachment** | Pointer to **bool** | Toggles whether to attach documents in the email | [optional] 
**EnableClientPortalPassword** | Pointer to **bool** | Toggles password protection of the client portal | [optional] 
**EnableEmailMarkup** | Pointer to **bool** | Toggles the use of markdown in emails | [optional] 
**EnableClientPortalDashboard** | Pointer to **bool** | Toggles whether the client dashboard is shown in the client portal | [optional] 
**EnableClientPortal** | Pointer to **bool** | Toggles whether the entire client portal is displayed to the client, or only the context | [optional] 
**EmailTemplateStatement** | Pointer to **string** | The body of the email for statements | [optional] 
**EmailSubjectStatement** | Pointer to **string** | The subject of the email for statements | [optional] 
**SignatureOnPdf** | Pointer to **bool** | Toggles whether the signature (if available) is displayed on the PDF | [optional] 
**QuoteFooter** | Pointer to **string** | The default quote footer | [optional] 
**EmailSubjectCustom1** | Pointer to **string** | Custom reminder template subject | [optional] 
**EmailSubjectCustom2** | Pointer to **string** | Custom reminder template subject | [optional] 
**EmailSubjectCustom3** | Pointer to **string** | Custom reminder template subject | [optional] 
**EmailTemplateCustom1** | Pointer to **string** | Custom reminder template body | [optional] 
**EmailTemplateCustom2** | Pointer to **string** | Custom reminder template body | [optional] 
**EmailTemplateCustom3** | Pointer to **string** | Custom reminder template body | [optional] 
**EnableReminder1** | Pointer to **bool** | Toggles whether this reminder is enabled | [optional] 
**EnableReminder2** | Pointer to **bool** | Toggles whether this reminder is enabled | [optional] 
**EnableReminder3** | Pointer to **bool** | Toggles whether this reminder is enabled | [optional] 
**NumDaysReminder1** | Pointer to **float32** | The Reminder interval | [optional] 
**NumDaysReminder2** | Pointer to **float32** | The Reminder interval | [optional] 
**NumDaysReminder3** | Pointer to **float32** | The Reminder interval | [optional] 
**ScheduleReminder1** | Pointer to **string** | (enum: after_invoice_date, before_due_date, after_due_date) | [optional] 
**ScheduleReminder2** | Pointer to **string** | (enum: after_invoice_date, before_due_date, after_due_date) | [optional] 
**ScheduleReminder3** | Pointer to **string** | (enum: after_invoice_date, before_due_date, after_due_date) | [optional] 
**LateFeeAmount1** | Pointer to **float32** | The late fee amount for reminder 1 | [optional] 
**LateFeeAmount2** | Pointer to **float32** | The late fee amount for reminder 2 | [optional] 
**LateFeeAmount3** | Pointer to **float32** | The late fee amount for reminder 2 | [optional] 
**EndlessReminderFrequencyId** | Pointer to **string** | The frequency id of the endless reminder | [optional] 
**ClientOnlinePaymentNotification** | Pointer to **bool** | Determines if a client should receive the notification for a online payment | [optional] 
**ClientManualPaymentNotification** | Pointer to **bool** | Determines if a client should receive the notification for a manually entered payment | [optional] 
**EnableEInvoice** | Pointer to **bool** | Determines if e-invoicing is enabled | [optional] 
**DefaultExpensePaymentTypeId** | Pointer to **string** | The default payment type for expenses | [optional] 
**EInvoiceType** | Pointer to **string** | The e-invoice type | [optional] 
**MailgunEndpoint** | Pointer to **string** | The mailgun endpoint - used to determine whether US or EU endpoints are used | [optional] 
**ClientInitiatedPayments** | Pointer to **bool** | Determines if clients can initiate payments directly from the client portal | [optional] 
**ClientInitiatedPaymentsMinimum** | Pointer to **float32** | The minimum amount a client can pay | [optional] 
**SyncInvoiceQuoteColumns** | Pointer to **bool** | Determines if invoice and quote columns are synced for the PDF rendering, or if they use their own columns | [optional] 
**ShowTaskItemDescription** | Pointer to **bool** | Determines if the task item description is shown on the invoice | [optional] 
**AllowBillableTaskItems** | Pointer to **bool** | Determines if task items can be marked as billable | [optional] 
**AcceptClientInputQuoteApproval** | Pointer to **bool** | Determines if clients can approve quotes and also pass through a PO Number reference | [optional] 
**CustomSendingEmail** | Pointer to **string** | When using Mailgun or Postmark, the FROM email address can be customized using this setting. | [optional] 
**ShowPaidStamp** | Pointer to **bool** | Determines if the PAID stamp is shown on the invoice | [optional] 
**ShowShippingAddress** | Pointer to **bool** | Determines if the shipping address is shown on the invoice | [optional] 
**CompanyLogoSize** | Pointer to **float32** | The size of the company logo on the PDF - percentage value between 0 and 100 | [optional] 
**ShowEmailFooter** | Pointer to **bool** | Determines if the email footer is shown on emails | [optional] 
**EmailAlignment** | Pointer to **string** | The alignment of the email body text, options include left / center / right | [optional] 
**AutoBillStandardInvoices** | Pointer to **bool** | Determines if standard invoices are automatically billed when they are created or due | [optional] 
**PostmarkSecret** | Pointer to **string** | The Postmark secret API key | [optional] 
**MailgunSecret** | Pointer to **string** | The Mailgun secret API key | [optional] 
**MailgunDomain** | Pointer to **string** | The Mailgun domain | [optional] 
**SendEmailOnMarkPaid** | Pointer to **bool** | Determines if an email is sent when an invoice is marked as paid | [optional] 
**VendorPortalEnableUploads** | Pointer to **bool** | Determines if vendors can upload files to the portal | [optional] 
**BesrId** | Pointer to **string** | The BESR ID | [optional] 
**QrIban** | Pointer to **string** | The IBAN for the QR code | [optional] 
**EmailSubjectPurchaseOrder** | Pointer to **string** | The email subject for purchase orders | [optional] 
**EmailTemplatePurchaseOrder** | Pointer to **string** | The email template for purchase orders | [optional] 
**RequirePurchaseOrderSignature** | Pointer to **bool** | Determines if a signature is required on purchase orders | [optional] 
**PurchaseOrderPublicNotes** | Pointer to **string** | The public notes for purchase orders | [optional] 
**PurchaseOrderTerms** | Pointer to **string** | The terms for purchase orders | [optional] 
**PurchaseOrderFooter** | Pointer to **string** | The footer for purchase orders | [optional] 
**PurchaseOrderDesignId** | Pointer to **string** | The design id for purchase orders | [optional] 
**PurchaseOrderNumberPattern** | Pointer to **string** | The pattern for purchase order numbers | [optional] 
**PurchaseOrderNumberCounter** | Pointer to **float32** | The counter for purchase order numbers | [optional] 
**PageNumberingAlignment** | Pointer to **string** | The alignment for page numbering: options include left / center / right | [optional] 
**PageNumbering** | Pointer to **bool** | Determines if page numbering is enabled on Document PDFs | [optional] 
**AutoArchiveInvoiceCancelled** | Pointer to **bool** | Determines if invoices are automatically archived when they are cancelled | [optional] 
**EmailFromName** | Pointer to **string** | The FROM name for emails when using Custom emailers | [optional] 
**ShowAllTasksClientPortal** | Pointer to **bool** | Determines if all tasks are shown on the client portal | [optional] 
**EntitySendTime** | Pointer to **int32** | The time that emails are sent. The time is localized to the clients locale, integer values from 1 - 24 | [optional] 
**SharedInvoiceCreditCounter** | Pointer to **bool** | Determines if the invoice and credit counter are shared | [optional] 
**ReplyToName** | Pointer to **string** | The reply to name for emails | [optional] 
**HideEmptyColumnsOnPdf** | Pointer to **bool** | Determines if empty columns are hidden on PDFs | [optional] 
**EnableReminderEndless** | Pointer to **bool** | Determines if endless reminders are enabled | [optional] 
**UseCreditsPayment** | Pointer to **bool** | Determines if credits can be used as a payment method | [optional] 
**RecurringInvoiceNumberPattern** | Pointer to **string** | The pattern for recurring invoice numbers | [optional] 
**RecurringInvoiceNumberCounter** | Pointer to **float32** | The counter for recurring invoice numbers | [optional] 
**ClientPortalUnderPaymentMinimum** | Pointer to **float32** | The minimum payment payment | [optional] 
**AutoBillDate** | Pointer to **string** | Determines when the invoices are auto billed, options are on_send_date (when the invoice is sent) or on_due_date (when the invoice is due)) | [optional] 
**PrimaryColor** | Pointer to **string** | The primary color for the client portal / document highlights | [optional] 
**SecondaryColor** | Pointer to **string** | The secondary color for the client portal / document highlights | [optional] 
**ClientPortalAllowUnderPayment** | Pointer to **bool** | Determines if clients can pay invoices under the invoice amount due | [optional] 
**ClientPortalAllowOverPayment** | Pointer to **bool** | Determines if clients can pay invoices over the invoice amount | [optional] 
**AutoBill** | Pointer to **string** | Determines how autobilling is applied for recurring invoices. off (no auto billed), always (always auto bill), optin (The user must opt in to auto billing), optout (The user must opt out of auto billing | [optional] 
**ClientPortalTerms** | Pointer to **string** | The terms which are displayed on the client portal | [optional] 
**ClientPortalPrivacyPolicy** | Pointer to **string** | The privacy policy which is displayed on the client portal | [optional] 
**ClientCanRegister** | Pointer to **bool** | Determines if clients can register on the client portal | [optional] 
**PortalDesignId** | Pointer to **string** | The design id for the client portal | [optional] 
**LateFeeEndlessPercent** | Pointer to **float32** | The late fee percentage for endless late fees | [optional] 
**LateFeeEndlessAmount** | Pointer to **float32** | The late fee amount for endless late fees | [optional] 
**AutoEmailInvoice** | Pointer to **bool** | Determines if invoices are automatically emailed when they are created | [optional] 
**EmailSignature** | Pointer to **string** | The email signature for emails | [optional] 

## Methods

### NewClientSettings

`func NewClientSettings(currencyId string, ) *ClientSettings`

NewClientSettings instantiates a new ClientSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSettingsWithDefaults

`func NewClientSettingsWithDefaults() *ClientSettings`

NewClientSettingsWithDefaults instantiates a new ClientSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyId

`func (o *ClientSettings) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ClientSettings) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ClientSettings) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetTimezoneId

`func (o *ClientSettings) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ClientSettings) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ClientSettings) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ClientSettings) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetDateFormatId

`func (o *ClientSettings) GetDateFormatId() string`

GetDateFormatId returns the DateFormatId field if non-nil, zero value otherwise.

### GetDateFormatIdOk

`func (o *ClientSettings) GetDateFormatIdOk() (*string, bool)`

GetDateFormatIdOk returns a tuple with the DateFormatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormatId

`func (o *ClientSettings) SetDateFormatId(v string)`

SetDateFormatId sets DateFormatId field to given value.

### HasDateFormatId

`func (o *ClientSettings) HasDateFormatId() bool`

HasDateFormatId returns a boolean if a field has been set.

### GetMilitaryTime

`func (o *ClientSettings) GetMilitaryTime() bool`

GetMilitaryTime returns the MilitaryTime field if non-nil, zero value otherwise.

### GetMilitaryTimeOk

`func (o *ClientSettings) GetMilitaryTimeOk() (*bool, bool)`

GetMilitaryTimeOk returns a tuple with the MilitaryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilitaryTime

`func (o *ClientSettings) SetMilitaryTime(v bool)`

SetMilitaryTime sets MilitaryTime field to given value.

### HasMilitaryTime

`func (o *ClientSettings) HasMilitaryTime() bool`

HasMilitaryTime returns a boolean if a field has been set.

### GetLanguageId

`func (o *ClientSettings) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *ClientSettings) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *ClientSettings) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *ClientSettings) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### GetShowCurrencyCode

`func (o *ClientSettings) GetShowCurrencyCode() bool`

GetShowCurrencyCode returns the ShowCurrencyCode field if non-nil, zero value otherwise.

### GetShowCurrencyCodeOk

`func (o *ClientSettings) GetShowCurrencyCodeOk() (*bool, bool)`

GetShowCurrencyCodeOk returns a tuple with the ShowCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCurrencyCode

`func (o *ClientSettings) SetShowCurrencyCode(v bool)`

SetShowCurrencyCode sets ShowCurrencyCode field to given value.

### HasShowCurrencyCode

`func (o *ClientSettings) HasShowCurrencyCode() bool`

HasShowCurrencyCode returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *ClientSettings) GetPaymentTerms() int32`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *ClientSettings) GetPaymentTermsOk() (*int32, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *ClientSettings) SetPaymentTerms(v int32)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *ClientSettings) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetCompanyGatewayIds

`func (o *ClientSettings) GetCompanyGatewayIds() string`

GetCompanyGatewayIds returns the CompanyGatewayIds field if non-nil, zero value otherwise.

### GetCompanyGatewayIdsOk

`func (o *ClientSettings) GetCompanyGatewayIdsOk() (*string, bool)`

GetCompanyGatewayIdsOk returns a tuple with the CompanyGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyGatewayIds

`func (o *ClientSettings) SetCompanyGatewayIds(v string)`

SetCompanyGatewayIds sets CompanyGatewayIds field to given value.

### HasCompanyGatewayIds

`func (o *ClientSettings) HasCompanyGatewayIds() bool`

HasCompanyGatewayIds returns a boolean if a field has been set.

### GetCustomValue1

`func (o *ClientSettings) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *ClientSettings) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *ClientSettings) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *ClientSettings) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *ClientSettings) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *ClientSettings) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *ClientSettings) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *ClientSettings) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *ClientSettings) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *ClientSettings) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *ClientSettings) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *ClientSettings) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *ClientSettings) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *ClientSettings) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *ClientSettings) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *ClientSettings) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetDefaultTaskRate

`func (o *ClientSettings) GetDefaultTaskRate() float32`

GetDefaultTaskRate returns the DefaultTaskRate field if non-nil, zero value otherwise.

### GetDefaultTaskRateOk

`func (o *ClientSettings) GetDefaultTaskRateOk() (*float32, bool)`

GetDefaultTaskRateOk returns a tuple with the DefaultTaskRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaskRate

`func (o *ClientSettings) SetDefaultTaskRate(v float32)`

SetDefaultTaskRate sets DefaultTaskRate field to given value.

### HasDefaultTaskRate

`func (o *ClientSettings) HasDefaultTaskRate() bool`

HasDefaultTaskRate returns a boolean if a field has been set.

### GetSendReminders

`func (o *ClientSettings) GetSendReminders() bool`

GetSendReminders returns the SendReminders field if non-nil, zero value otherwise.

### GetSendRemindersOk

`func (o *ClientSettings) GetSendRemindersOk() (*bool, bool)`

GetSendRemindersOk returns a tuple with the SendReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReminders

`func (o *ClientSettings) SetSendReminders(v bool)`

SetSendReminders sets SendReminders field to given value.

### HasSendReminders

`func (o *ClientSettings) HasSendReminders() bool`

HasSendReminders returns a boolean if a field has been set.

### GetEnableClientPortalTasks

`func (o *ClientSettings) GetEnableClientPortalTasks() bool`

GetEnableClientPortalTasks returns the EnableClientPortalTasks field if non-nil, zero value otherwise.

### GetEnableClientPortalTasksOk

`func (o *ClientSettings) GetEnableClientPortalTasksOk() (*bool, bool)`

GetEnableClientPortalTasksOk returns a tuple with the EnableClientPortalTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientPortalTasks

`func (o *ClientSettings) SetEnableClientPortalTasks(v bool)`

SetEnableClientPortalTasks sets EnableClientPortalTasks field to given value.

### HasEnableClientPortalTasks

`func (o *ClientSettings) HasEnableClientPortalTasks() bool`

HasEnableClientPortalTasks returns a boolean if a field has been set.

### GetEmailStyle

`func (o *ClientSettings) GetEmailStyle() string`

GetEmailStyle returns the EmailStyle field if non-nil, zero value otherwise.

### GetEmailStyleOk

`func (o *ClientSettings) GetEmailStyleOk() (*string, bool)`

GetEmailStyleOk returns a tuple with the EmailStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStyle

`func (o *ClientSettings) SetEmailStyle(v string)`

SetEmailStyle sets EmailStyle field to given value.

### HasEmailStyle

`func (o *ClientSettings) HasEmailStyle() bool`

HasEmailStyle returns a boolean if a field has been set.

### GetReplyToEmail

`func (o *ClientSettings) GetReplyToEmail() string`

GetReplyToEmail returns the ReplyToEmail field if non-nil, zero value otherwise.

### GetReplyToEmailOk

`func (o *ClientSettings) GetReplyToEmailOk() (*string, bool)`

GetReplyToEmailOk returns a tuple with the ReplyToEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToEmail

`func (o *ClientSettings) SetReplyToEmail(v string)`

SetReplyToEmail sets ReplyToEmail field to given value.

### HasReplyToEmail

`func (o *ClientSettings) HasReplyToEmail() bool`

HasReplyToEmail returns a boolean if a field has been set.

### GetBccEmail

`func (o *ClientSettings) GetBccEmail() string`

GetBccEmail returns the BccEmail field if non-nil, zero value otherwise.

### GetBccEmailOk

`func (o *ClientSettings) GetBccEmailOk() (*string, bool)`

GetBccEmailOk returns a tuple with the BccEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccEmail

`func (o *ClientSettings) SetBccEmail(v string)`

SetBccEmail sets BccEmail field to given value.

### HasBccEmail

`func (o *ClientSettings) HasBccEmail() bool`

HasBccEmail returns a boolean if a field has been set.

### GetPdfEmailAttachment

`func (o *ClientSettings) GetPdfEmailAttachment() bool`

GetPdfEmailAttachment returns the PdfEmailAttachment field if non-nil, zero value otherwise.

### GetPdfEmailAttachmentOk

`func (o *ClientSettings) GetPdfEmailAttachmentOk() (*bool, bool)`

GetPdfEmailAttachmentOk returns a tuple with the PdfEmailAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfEmailAttachment

`func (o *ClientSettings) SetPdfEmailAttachment(v bool)`

SetPdfEmailAttachment sets PdfEmailAttachment field to given value.

### HasPdfEmailAttachment

`func (o *ClientSettings) HasPdfEmailAttachment() bool`

HasPdfEmailAttachment returns a boolean if a field has been set.

### GetUblEmailAttachment

`func (o *ClientSettings) GetUblEmailAttachment() bool`

GetUblEmailAttachment returns the UblEmailAttachment field if non-nil, zero value otherwise.

### GetUblEmailAttachmentOk

`func (o *ClientSettings) GetUblEmailAttachmentOk() (*bool, bool)`

GetUblEmailAttachmentOk returns a tuple with the UblEmailAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUblEmailAttachment

`func (o *ClientSettings) SetUblEmailAttachment(v bool)`

SetUblEmailAttachment sets UblEmailAttachment field to given value.

### HasUblEmailAttachment

`func (o *ClientSettings) HasUblEmailAttachment() bool`

HasUblEmailAttachment returns a boolean if a field has been set.

### GetEmailStyleCustom

`func (o *ClientSettings) GetEmailStyleCustom() string`

GetEmailStyleCustom returns the EmailStyleCustom field if non-nil, zero value otherwise.

### GetEmailStyleCustomOk

`func (o *ClientSettings) GetEmailStyleCustomOk() (*string, bool)`

GetEmailStyleCustomOk returns a tuple with the EmailStyleCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStyleCustom

`func (o *ClientSettings) SetEmailStyleCustom(v string)`

SetEmailStyleCustom sets EmailStyleCustom field to given value.

### HasEmailStyleCustom

`func (o *ClientSettings) HasEmailStyleCustom() bool`

HasEmailStyleCustom returns a boolean if a field has been set.

### GetCounterNumberApplied

`func (o *ClientSettings) GetCounterNumberApplied() string`

GetCounterNumberApplied returns the CounterNumberApplied field if non-nil, zero value otherwise.

### GetCounterNumberAppliedOk

`func (o *ClientSettings) GetCounterNumberAppliedOk() (*string, bool)`

GetCounterNumberAppliedOk returns a tuple with the CounterNumberApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterNumberApplied

`func (o *ClientSettings) SetCounterNumberApplied(v string)`

SetCounterNumberApplied sets CounterNumberApplied field to given value.

### HasCounterNumberApplied

`func (o *ClientSettings) HasCounterNumberApplied() bool`

HasCounterNumberApplied returns a boolean if a field has been set.

### GetQuoteNumberApplied

`func (o *ClientSettings) GetQuoteNumberApplied() string`

GetQuoteNumberApplied returns the QuoteNumberApplied field if non-nil, zero value otherwise.

### GetQuoteNumberAppliedOk

`func (o *ClientSettings) GetQuoteNumberAppliedOk() (*string, bool)`

GetQuoteNumberAppliedOk returns a tuple with the QuoteNumberApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteNumberApplied

`func (o *ClientSettings) SetQuoteNumberApplied(v string)`

SetQuoteNumberApplied sets QuoteNumberApplied field to given value.

### HasQuoteNumberApplied

`func (o *ClientSettings) HasQuoteNumberApplied() bool`

HasQuoteNumberApplied returns a boolean if a field has been set.

### GetCustomMessageDashboard

`func (o *ClientSettings) GetCustomMessageDashboard() string`

GetCustomMessageDashboard returns the CustomMessageDashboard field if non-nil, zero value otherwise.

### GetCustomMessageDashboardOk

`func (o *ClientSettings) GetCustomMessageDashboardOk() (*string, bool)`

GetCustomMessageDashboardOk returns a tuple with the CustomMessageDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessageDashboard

`func (o *ClientSettings) SetCustomMessageDashboard(v string)`

SetCustomMessageDashboard sets CustomMessageDashboard field to given value.

### HasCustomMessageDashboard

`func (o *ClientSettings) HasCustomMessageDashboard() bool`

HasCustomMessageDashboard returns a boolean if a field has been set.

### GetCustomMessageUnpaidInvoice

`func (o *ClientSettings) GetCustomMessageUnpaidInvoice() string`

GetCustomMessageUnpaidInvoice returns the CustomMessageUnpaidInvoice field if non-nil, zero value otherwise.

### GetCustomMessageUnpaidInvoiceOk

`func (o *ClientSettings) GetCustomMessageUnpaidInvoiceOk() (*string, bool)`

GetCustomMessageUnpaidInvoiceOk returns a tuple with the CustomMessageUnpaidInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessageUnpaidInvoice

`func (o *ClientSettings) SetCustomMessageUnpaidInvoice(v string)`

SetCustomMessageUnpaidInvoice sets CustomMessageUnpaidInvoice field to given value.

### HasCustomMessageUnpaidInvoice

`func (o *ClientSettings) HasCustomMessageUnpaidInvoice() bool`

HasCustomMessageUnpaidInvoice returns a boolean if a field has been set.

### GetCustomMessagePaidInvoice

`func (o *ClientSettings) GetCustomMessagePaidInvoice() string`

GetCustomMessagePaidInvoice returns the CustomMessagePaidInvoice field if non-nil, zero value otherwise.

### GetCustomMessagePaidInvoiceOk

`func (o *ClientSettings) GetCustomMessagePaidInvoiceOk() (*string, bool)`

GetCustomMessagePaidInvoiceOk returns a tuple with the CustomMessagePaidInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessagePaidInvoice

`func (o *ClientSettings) SetCustomMessagePaidInvoice(v string)`

SetCustomMessagePaidInvoice sets CustomMessagePaidInvoice field to given value.

### HasCustomMessagePaidInvoice

`func (o *ClientSettings) HasCustomMessagePaidInvoice() bool`

HasCustomMessagePaidInvoice returns a boolean if a field has been set.

### GetCustomMessageUnapprovedQuote

`func (o *ClientSettings) GetCustomMessageUnapprovedQuote() string`

GetCustomMessageUnapprovedQuote returns the CustomMessageUnapprovedQuote field if non-nil, zero value otherwise.

### GetCustomMessageUnapprovedQuoteOk

`func (o *ClientSettings) GetCustomMessageUnapprovedQuoteOk() (*string, bool)`

GetCustomMessageUnapprovedQuoteOk returns a tuple with the CustomMessageUnapprovedQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessageUnapprovedQuote

`func (o *ClientSettings) SetCustomMessageUnapprovedQuote(v string)`

SetCustomMessageUnapprovedQuote sets CustomMessageUnapprovedQuote field to given value.

### HasCustomMessageUnapprovedQuote

`func (o *ClientSettings) HasCustomMessageUnapprovedQuote() bool`

HasCustomMessageUnapprovedQuote returns a boolean if a field has been set.

### GetLockInvoices

`func (o *ClientSettings) GetLockInvoices() bool`

GetLockInvoices returns the LockInvoices field if non-nil, zero value otherwise.

### GetLockInvoicesOk

`func (o *ClientSettings) GetLockInvoicesOk() (*bool, bool)`

GetLockInvoicesOk returns a tuple with the LockInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockInvoices

`func (o *ClientSettings) SetLockInvoices(v bool)`

SetLockInvoices sets LockInvoices field to given value.

### HasLockInvoices

`func (o *ClientSettings) HasLockInvoices() bool`

HasLockInvoices returns a boolean if a field has been set.

### GetAutoArchiveInvoice

`func (o *ClientSettings) GetAutoArchiveInvoice() bool`

GetAutoArchiveInvoice returns the AutoArchiveInvoice field if non-nil, zero value otherwise.

### GetAutoArchiveInvoiceOk

`func (o *ClientSettings) GetAutoArchiveInvoiceOk() (*bool, bool)`

GetAutoArchiveInvoiceOk returns a tuple with the AutoArchiveInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveInvoice

`func (o *ClientSettings) SetAutoArchiveInvoice(v bool)`

SetAutoArchiveInvoice sets AutoArchiveInvoice field to given value.

### HasAutoArchiveInvoice

`func (o *ClientSettings) HasAutoArchiveInvoice() bool`

HasAutoArchiveInvoice returns a boolean if a field has been set.

### GetAutoArchiveQuote

`func (o *ClientSettings) GetAutoArchiveQuote() bool`

GetAutoArchiveQuote returns the AutoArchiveQuote field if non-nil, zero value otherwise.

### GetAutoArchiveQuoteOk

`func (o *ClientSettings) GetAutoArchiveQuoteOk() (*bool, bool)`

GetAutoArchiveQuoteOk returns a tuple with the AutoArchiveQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveQuote

`func (o *ClientSettings) SetAutoArchiveQuote(v bool)`

SetAutoArchiveQuote sets AutoArchiveQuote field to given value.

### HasAutoArchiveQuote

`func (o *ClientSettings) HasAutoArchiveQuote() bool`

HasAutoArchiveQuote returns a boolean if a field has been set.

### GetAutoConvertQuote

`func (o *ClientSettings) GetAutoConvertQuote() bool`

GetAutoConvertQuote returns the AutoConvertQuote field if non-nil, zero value otherwise.

### GetAutoConvertQuoteOk

`func (o *ClientSettings) GetAutoConvertQuoteOk() (*bool, bool)`

GetAutoConvertQuoteOk returns a tuple with the AutoConvertQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConvertQuote

`func (o *ClientSettings) SetAutoConvertQuote(v bool)`

SetAutoConvertQuote sets AutoConvertQuote field to given value.

### HasAutoConvertQuote

`func (o *ClientSettings) HasAutoConvertQuote() bool`

HasAutoConvertQuote returns a boolean if a field has been set.

### GetInclusiveTaxes

`func (o *ClientSettings) GetInclusiveTaxes() bool`

GetInclusiveTaxes returns the InclusiveTaxes field if non-nil, zero value otherwise.

### GetInclusiveTaxesOk

`func (o *ClientSettings) GetInclusiveTaxesOk() (*bool, bool)`

GetInclusiveTaxesOk returns a tuple with the InclusiveTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusiveTaxes

`func (o *ClientSettings) SetInclusiveTaxes(v bool)`

SetInclusiveTaxes sets InclusiveTaxes field to given value.

### HasInclusiveTaxes

`func (o *ClientSettings) HasInclusiveTaxes() bool`

HasInclusiveTaxes returns a boolean if a field has been set.

### GetTaskNumberPattern

`func (o *ClientSettings) GetTaskNumberPattern() string`

GetTaskNumberPattern returns the TaskNumberPattern field if non-nil, zero value otherwise.

### GetTaskNumberPatternOk

`func (o *ClientSettings) GetTaskNumberPatternOk() (*string, bool)`

GetTaskNumberPatternOk returns a tuple with the TaskNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskNumberPattern

`func (o *ClientSettings) SetTaskNumberPattern(v string)`

SetTaskNumberPattern sets TaskNumberPattern field to given value.

### HasTaskNumberPattern

`func (o *ClientSettings) HasTaskNumberPattern() bool`

HasTaskNumberPattern returns a boolean if a field has been set.

### GetTaskNumberCounter

`func (o *ClientSettings) GetTaskNumberCounter() int32`

GetTaskNumberCounter returns the TaskNumberCounter field if non-nil, zero value otherwise.

### GetTaskNumberCounterOk

`func (o *ClientSettings) GetTaskNumberCounterOk() (*int32, bool)`

GetTaskNumberCounterOk returns a tuple with the TaskNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskNumberCounter

`func (o *ClientSettings) SetTaskNumberCounter(v int32)`

SetTaskNumberCounter sets TaskNumberCounter field to given value.

### HasTaskNumberCounter

`func (o *ClientSettings) HasTaskNumberCounter() bool`

HasTaskNumberCounter returns a boolean if a field has been set.

### GetReminderSendTime

`func (o *ClientSettings) GetReminderSendTime() int32`

GetReminderSendTime returns the ReminderSendTime field if non-nil, zero value otherwise.

### GetReminderSendTimeOk

`func (o *ClientSettings) GetReminderSendTimeOk() (*int32, bool)`

GetReminderSendTimeOk returns a tuple with the ReminderSendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderSendTime

`func (o *ClientSettings) SetReminderSendTime(v int32)`

SetReminderSendTime sets ReminderSendTime field to given value.

### HasReminderSendTime

`func (o *ClientSettings) HasReminderSendTime() bool`

HasReminderSendTime returns a boolean if a field has been set.

### GetExpenseNumberPattern

`func (o *ClientSettings) GetExpenseNumberPattern() string`

GetExpenseNumberPattern returns the ExpenseNumberPattern field if non-nil, zero value otherwise.

### GetExpenseNumberPatternOk

`func (o *ClientSettings) GetExpenseNumberPatternOk() (*string, bool)`

GetExpenseNumberPatternOk returns a tuple with the ExpenseNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseNumberPattern

`func (o *ClientSettings) SetExpenseNumberPattern(v string)`

SetExpenseNumberPattern sets ExpenseNumberPattern field to given value.

### HasExpenseNumberPattern

`func (o *ClientSettings) HasExpenseNumberPattern() bool`

HasExpenseNumberPattern returns a boolean if a field has been set.

### GetExpenseNumberCounter

`func (o *ClientSettings) GetExpenseNumberCounter() int32`

GetExpenseNumberCounter returns the ExpenseNumberCounter field if non-nil, zero value otherwise.

### GetExpenseNumberCounterOk

`func (o *ClientSettings) GetExpenseNumberCounterOk() (*int32, bool)`

GetExpenseNumberCounterOk returns a tuple with the ExpenseNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseNumberCounter

`func (o *ClientSettings) SetExpenseNumberCounter(v int32)`

SetExpenseNumberCounter sets ExpenseNumberCounter field to given value.

### HasExpenseNumberCounter

`func (o *ClientSettings) HasExpenseNumberCounter() bool`

HasExpenseNumberCounter returns a boolean if a field has been set.

### GetVendorNumberPattern

`func (o *ClientSettings) GetVendorNumberPattern() string`

GetVendorNumberPattern returns the VendorNumberPattern field if non-nil, zero value otherwise.

### GetVendorNumberPatternOk

`func (o *ClientSettings) GetVendorNumberPatternOk() (*string, bool)`

GetVendorNumberPatternOk returns a tuple with the VendorNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumberPattern

`func (o *ClientSettings) SetVendorNumberPattern(v string)`

SetVendorNumberPattern sets VendorNumberPattern field to given value.

### HasVendorNumberPattern

`func (o *ClientSettings) HasVendorNumberPattern() bool`

HasVendorNumberPattern returns a boolean if a field has been set.

### GetVendorNumberCounter

`func (o *ClientSettings) GetVendorNumberCounter() int32`

GetVendorNumberCounter returns the VendorNumberCounter field if non-nil, zero value otherwise.

### GetVendorNumberCounterOk

`func (o *ClientSettings) GetVendorNumberCounterOk() (*int32, bool)`

GetVendorNumberCounterOk returns a tuple with the VendorNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNumberCounter

`func (o *ClientSettings) SetVendorNumberCounter(v int32)`

SetVendorNumberCounter sets VendorNumberCounter field to given value.

### HasVendorNumberCounter

`func (o *ClientSettings) HasVendorNumberCounter() bool`

HasVendorNumberCounter returns a boolean if a field has been set.

### GetTicketNumberPattern

`func (o *ClientSettings) GetTicketNumberPattern() string`

GetTicketNumberPattern returns the TicketNumberPattern field if non-nil, zero value otherwise.

### GetTicketNumberPatternOk

`func (o *ClientSettings) GetTicketNumberPatternOk() (*string, bool)`

GetTicketNumberPatternOk returns a tuple with the TicketNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumberPattern

`func (o *ClientSettings) SetTicketNumberPattern(v string)`

SetTicketNumberPattern sets TicketNumberPattern field to given value.

### HasTicketNumberPattern

`func (o *ClientSettings) HasTicketNumberPattern() bool`

HasTicketNumberPattern returns a boolean if a field has been set.

### GetTicketNumberCounter

`func (o *ClientSettings) GetTicketNumberCounter() int32`

GetTicketNumberCounter returns the TicketNumberCounter field if non-nil, zero value otherwise.

### GetTicketNumberCounterOk

`func (o *ClientSettings) GetTicketNumberCounterOk() (*int32, bool)`

GetTicketNumberCounterOk returns a tuple with the TicketNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumberCounter

`func (o *ClientSettings) SetTicketNumberCounter(v int32)`

SetTicketNumberCounter sets TicketNumberCounter field to given value.

### HasTicketNumberCounter

`func (o *ClientSettings) HasTicketNumberCounter() bool`

HasTicketNumberCounter returns a boolean if a field has been set.

### GetPaymentNumberPattern

`func (o *ClientSettings) GetPaymentNumberPattern() string`

GetPaymentNumberPattern returns the PaymentNumberPattern field if non-nil, zero value otherwise.

### GetPaymentNumberPatternOk

`func (o *ClientSettings) GetPaymentNumberPatternOk() (*string, bool)`

GetPaymentNumberPatternOk returns a tuple with the PaymentNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentNumberPattern

`func (o *ClientSettings) SetPaymentNumberPattern(v string)`

SetPaymentNumberPattern sets PaymentNumberPattern field to given value.

### HasPaymentNumberPattern

`func (o *ClientSettings) HasPaymentNumberPattern() bool`

HasPaymentNumberPattern returns a boolean if a field has been set.

### GetPaymentNumberCounter

`func (o *ClientSettings) GetPaymentNumberCounter() int32`

GetPaymentNumberCounter returns the PaymentNumberCounter field if non-nil, zero value otherwise.

### GetPaymentNumberCounterOk

`func (o *ClientSettings) GetPaymentNumberCounterOk() (*int32, bool)`

GetPaymentNumberCounterOk returns a tuple with the PaymentNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentNumberCounter

`func (o *ClientSettings) SetPaymentNumberCounter(v int32)`

SetPaymentNumberCounter sets PaymentNumberCounter field to given value.

### HasPaymentNumberCounter

`func (o *ClientSettings) HasPaymentNumberCounter() bool`

HasPaymentNumberCounter returns a boolean if a field has been set.

### GetInvoiceNumberPattern

`func (o *ClientSettings) GetInvoiceNumberPattern() string`

GetInvoiceNumberPattern returns the InvoiceNumberPattern field if non-nil, zero value otherwise.

### GetInvoiceNumberPatternOk

`func (o *ClientSettings) GetInvoiceNumberPatternOk() (*string, bool)`

GetInvoiceNumberPatternOk returns a tuple with the InvoiceNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumberPattern

`func (o *ClientSettings) SetInvoiceNumberPattern(v string)`

SetInvoiceNumberPattern sets InvoiceNumberPattern field to given value.

### HasInvoiceNumberPattern

`func (o *ClientSettings) HasInvoiceNumberPattern() bool`

HasInvoiceNumberPattern returns a boolean if a field has been set.

### GetInvoiceNumberCounter

`func (o *ClientSettings) GetInvoiceNumberCounter() int32`

GetInvoiceNumberCounter returns the InvoiceNumberCounter field if non-nil, zero value otherwise.

### GetInvoiceNumberCounterOk

`func (o *ClientSettings) GetInvoiceNumberCounterOk() (*int32, bool)`

GetInvoiceNumberCounterOk returns a tuple with the InvoiceNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumberCounter

`func (o *ClientSettings) SetInvoiceNumberCounter(v int32)`

SetInvoiceNumberCounter sets InvoiceNumberCounter field to given value.

### HasInvoiceNumberCounter

`func (o *ClientSettings) HasInvoiceNumberCounter() bool`

HasInvoiceNumberCounter returns a boolean if a field has been set.

### GetQuoteNumberPattern

`func (o *ClientSettings) GetQuoteNumberPattern() string`

GetQuoteNumberPattern returns the QuoteNumberPattern field if non-nil, zero value otherwise.

### GetQuoteNumberPatternOk

`func (o *ClientSettings) GetQuoteNumberPatternOk() (*string, bool)`

GetQuoteNumberPatternOk returns a tuple with the QuoteNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteNumberPattern

`func (o *ClientSettings) SetQuoteNumberPattern(v string)`

SetQuoteNumberPattern sets QuoteNumberPattern field to given value.

### HasQuoteNumberPattern

`func (o *ClientSettings) HasQuoteNumberPattern() bool`

HasQuoteNumberPattern returns a boolean if a field has been set.

### GetQuoteNumberCounter

`func (o *ClientSettings) GetQuoteNumberCounter() int32`

GetQuoteNumberCounter returns the QuoteNumberCounter field if non-nil, zero value otherwise.

### GetQuoteNumberCounterOk

`func (o *ClientSettings) GetQuoteNumberCounterOk() (*int32, bool)`

GetQuoteNumberCounterOk returns a tuple with the QuoteNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteNumberCounter

`func (o *ClientSettings) SetQuoteNumberCounter(v int32)`

SetQuoteNumberCounter sets QuoteNumberCounter field to given value.

### HasQuoteNumberCounter

`func (o *ClientSettings) HasQuoteNumberCounter() bool`

HasQuoteNumberCounter returns a boolean if a field has been set.

### GetClientNumberPattern

`func (o *ClientSettings) GetClientNumberPattern() string`

GetClientNumberPattern returns the ClientNumberPattern field if non-nil, zero value otherwise.

### GetClientNumberPatternOk

`func (o *ClientSettings) GetClientNumberPatternOk() (*string, bool)`

GetClientNumberPatternOk returns a tuple with the ClientNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNumberPattern

`func (o *ClientSettings) SetClientNumberPattern(v string)`

SetClientNumberPattern sets ClientNumberPattern field to given value.

### HasClientNumberPattern

`func (o *ClientSettings) HasClientNumberPattern() bool`

HasClientNumberPattern returns a boolean if a field has been set.

### GetClientNumberCounter

`func (o *ClientSettings) GetClientNumberCounter() int32`

GetClientNumberCounter returns the ClientNumberCounter field if non-nil, zero value otherwise.

### GetClientNumberCounterOk

`func (o *ClientSettings) GetClientNumberCounterOk() (*int32, bool)`

GetClientNumberCounterOk returns a tuple with the ClientNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNumberCounter

`func (o *ClientSettings) SetClientNumberCounter(v int32)`

SetClientNumberCounter sets ClientNumberCounter field to given value.

### HasClientNumberCounter

`func (o *ClientSettings) HasClientNumberCounter() bool`

HasClientNumberCounter returns a boolean if a field has been set.

### GetCreditNumberPattern

`func (o *ClientSettings) GetCreditNumberPattern() string`

GetCreditNumberPattern returns the CreditNumberPattern field if non-nil, zero value otherwise.

### GetCreditNumberPatternOk

`func (o *ClientSettings) GetCreditNumberPatternOk() (*string, bool)`

GetCreditNumberPatternOk returns a tuple with the CreditNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNumberPattern

`func (o *ClientSettings) SetCreditNumberPattern(v string)`

SetCreditNumberPattern sets CreditNumberPattern field to given value.

### HasCreditNumberPattern

`func (o *ClientSettings) HasCreditNumberPattern() bool`

HasCreditNumberPattern returns a boolean if a field has been set.

### GetCreditNumberCounter

`func (o *ClientSettings) GetCreditNumberCounter() int32`

GetCreditNumberCounter returns the CreditNumberCounter field if non-nil, zero value otherwise.

### GetCreditNumberCounterOk

`func (o *ClientSettings) GetCreditNumberCounterOk() (*int32, bool)`

GetCreditNumberCounterOk returns a tuple with the CreditNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNumberCounter

`func (o *ClientSettings) SetCreditNumberCounter(v int32)`

SetCreditNumberCounter sets CreditNumberCounter field to given value.

### HasCreditNumberCounter

`func (o *ClientSettings) HasCreditNumberCounter() bool`

HasCreditNumberCounter returns a boolean if a field has been set.

### GetRecurringInvoiceNumberPrefix

`func (o *ClientSettings) GetRecurringInvoiceNumberPrefix() string`

GetRecurringInvoiceNumberPrefix returns the RecurringInvoiceNumberPrefix field if non-nil, zero value otherwise.

### GetRecurringInvoiceNumberPrefixOk

`func (o *ClientSettings) GetRecurringInvoiceNumberPrefixOk() (*string, bool)`

GetRecurringInvoiceNumberPrefixOk returns a tuple with the RecurringInvoiceNumberPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInvoiceNumberPrefix

`func (o *ClientSettings) SetRecurringInvoiceNumberPrefix(v string)`

SetRecurringInvoiceNumberPrefix sets RecurringInvoiceNumberPrefix field to given value.

### HasRecurringInvoiceNumberPrefix

`func (o *ClientSettings) HasRecurringInvoiceNumberPrefix() bool`

HasRecurringInvoiceNumberPrefix returns a boolean if a field has been set.

### GetResetCounterFrequencyId

`func (o *ClientSettings) GetResetCounterFrequencyId() int32`

GetResetCounterFrequencyId returns the ResetCounterFrequencyId field if non-nil, zero value otherwise.

### GetResetCounterFrequencyIdOk

`func (o *ClientSettings) GetResetCounterFrequencyIdOk() (*int32, bool)`

GetResetCounterFrequencyIdOk returns a tuple with the ResetCounterFrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCounterFrequencyId

`func (o *ClientSettings) SetResetCounterFrequencyId(v int32)`

SetResetCounterFrequencyId sets ResetCounterFrequencyId field to given value.

### HasResetCounterFrequencyId

`func (o *ClientSettings) HasResetCounterFrequencyId() bool`

HasResetCounterFrequencyId returns a boolean if a field has been set.

### GetResetCounterDate

`func (o *ClientSettings) GetResetCounterDate() string`

GetResetCounterDate returns the ResetCounterDate field if non-nil, zero value otherwise.

### GetResetCounterDateOk

`func (o *ClientSettings) GetResetCounterDateOk() (*string, bool)`

GetResetCounterDateOk returns a tuple with the ResetCounterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCounterDate

`func (o *ClientSettings) SetResetCounterDate(v string)`

SetResetCounterDate sets ResetCounterDate field to given value.

### HasResetCounterDate

`func (o *ClientSettings) HasResetCounterDate() bool`

HasResetCounterDate returns a boolean if a field has been set.

### GetCounterPadding

`func (o *ClientSettings) GetCounterPadding() int32`

GetCounterPadding returns the CounterPadding field if non-nil, zero value otherwise.

### GetCounterPaddingOk

`func (o *ClientSettings) GetCounterPaddingOk() (*int32, bool)`

GetCounterPaddingOk returns a tuple with the CounterPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterPadding

`func (o *ClientSettings) SetCounterPadding(v int32)`

SetCounterPadding sets CounterPadding field to given value.

### HasCounterPadding

`func (o *ClientSettings) HasCounterPadding() bool`

HasCounterPadding returns a boolean if a field has been set.

### GetSharedInvoiceQuoteCounter

`func (o *ClientSettings) GetSharedInvoiceQuoteCounter() bool`

GetSharedInvoiceQuoteCounter returns the SharedInvoiceQuoteCounter field if non-nil, zero value otherwise.

### GetSharedInvoiceQuoteCounterOk

`func (o *ClientSettings) GetSharedInvoiceQuoteCounterOk() (*bool, bool)`

GetSharedInvoiceQuoteCounterOk returns a tuple with the SharedInvoiceQuoteCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedInvoiceQuoteCounter

`func (o *ClientSettings) SetSharedInvoiceQuoteCounter(v bool)`

SetSharedInvoiceQuoteCounter sets SharedInvoiceQuoteCounter field to given value.

### HasSharedInvoiceQuoteCounter

`func (o *ClientSettings) HasSharedInvoiceQuoteCounter() bool`

HasSharedInvoiceQuoteCounter returns a boolean if a field has been set.

### GetUpdateProducts

`func (o *ClientSettings) GetUpdateProducts() bool`

GetUpdateProducts returns the UpdateProducts field if non-nil, zero value otherwise.

### GetUpdateProductsOk

`func (o *ClientSettings) GetUpdateProductsOk() (*bool, bool)`

GetUpdateProductsOk returns a tuple with the UpdateProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateProducts

`func (o *ClientSettings) SetUpdateProducts(v bool)`

SetUpdateProducts sets UpdateProducts field to given value.

### HasUpdateProducts

`func (o *ClientSettings) HasUpdateProducts() bool`

HasUpdateProducts returns a boolean if a field has been set.

### GetConvertProducts

`func (o *ClientSettings) GetConvertProducts() bool`

GetConvertProducts returns the ConvertProducts field if non-nil, zero value otherwise.

### GetConvertProductsOk

`func (o *ClientSettings) GetConvertProductsOk() (*bool, bool)`

GetConvertProductsOk returns a tuple with the ConvertProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertProducts

`func (o *ClientSettings) SetConvertProducts(v bool)`

SetConvertProducts sets ConvertProducts field to given value.

### HasConvertProducts

`func (o *ClientSettings) HasConvertProducts() bool`

HasConvertProducts returns a boolean if a field has been set.

### GetFillProducts

`func (o *ClientSettings) GetFillProducts() bool`

GetFillProducts returns the FillProducts field if non-nil, zero value otherwise.

### GetFillProductsOk

`func (o *ClientSettings) GetFillProductsOk() (*bool, bool)`

GetFillProductsOk returns a tuple with the FillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillProducts

`func (o *ClientSettings) SetFillProducts(v bool)`

SetFillProducts sets FillProducts field to given value.

### HasFillProducts

`func (o *ClientSettings) HasFillProducts() bool`

HasFillProducts returns a boolean if a field has been set.

### GetInvoiceTerms

`func (o *ClientSettings) GetInvoiceTerms() string`

GetInvoiceTerms returns the InvoiceTerms field if non-nil, zero value otherwise.

### GetInvoiceTermsOk

`func (o *ClientSettings) GetInvoiceTermsOk() (*string, bool)`

GetInvoiceTermsOk returns a tuple with the InvoiceTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTerms

`func (o *ClientSettings) SetInvoiceTerms(v string)`

SetInvoiceTerms sets InvoiceTerms field to given value.

### HasInvoiceTerms

`func (o *ClientSettings) HasInvoiceTerms() bool`

HasInvoiceTerms returns a boolean if a field has been set.

### GetQuoteTerms

`func (o *ClientSettings) GetQuoteTerms() string`

GetQuoteTerms returns the QuoteTerms field if non-nil, zero value otherwise.

### GetQuoteTermsOk

`func (o *ClientSettings) GetQuoteTermsOk() (*string, bool)`

GetQuoteTermsOk returns a tuple with the QuoteTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTerms

`func (o *ClientSettings) SetQuoteTerms(v string)`

SetQuoteTerms sets QuoteTerms field to given value.

### HasQuoteTerms

`func (o *ClientSettings) HasQuoteTerms() bool`

HasQuoteTerms returns a boolean if a field has been set.

### GetInvoiceTaxes

`func (o *ClientSettings) GetInvoiceTaxes() float32`

GetInvoiceTaxes returns the InvoiceTaxes field if non-nil, zero value otherwise.

### GetInvoiceTaxesOk

`func (o *ClientSettings) GetInvoiceTaxesOk() (*float32, bool)`

GetInvoiceTaxesOk returns a tuple with the InvoiceTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaxes

`func (o *ClientSettings) SetInvoiceTaxes(v float32)`

SetInvoiceTaxes sets InvoiceTaxes field to given value.

### HasInvoiceTaxes

`func (o *ClientSettings) HasInvoiceTaxes() bool`

HasInvoiceTaxes returns a boolean if a field has been set.

### GetInvoiceDesignId

`func (o *ClientSettings) GetInvoiceDesignId() string`

GetInvoiceDesignId returns the InvoiceDesignId field if non-nil, zero value otherwise.

### GetInvoiceDesignIdOk

`func (o *ClientSettings) GetInvoiceDesignIdOk() (*string, bool)`

GetInvoiceDesignIdOk returns a tuple with the InvoiceDesignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDesignId

`func (o *ClientSettings) SetInvoiceDesignId(v string)`

SetInvoiceDesignId sets InvoiceDesignId field to given value.

### HasInvoiceDesignId

`func (o *ClientSettings) HasInvoiceDesignId() bool`

HasInvoiceDesignId returns a boolean if a field has been set.

### GetQuoteDesignId

`func (o *ClientSettings) GetQuoteDesignId() string`

GetQuoteDesignId returns the QuoteDesignId field if non-nil, zero value otherwise.

### GetQuoteDesignIdOk

`func (o *ClientSettings) GetQuoteDesignIdOk() (*string, bool)`

GetQuoteDesignIdOk returns a tuple with the QuoteDesignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteDesignId

`func (o *ClientSettings) SetQuoteDesignId(v string)`

SetQuoteDesignId sets QuoteDesignId field to given value.

### HasQuoteDesignId

`func (o *ClientSettings) HasQuoteDesignId() bool`

HasQuoteDesignId returns a boolean if a field has been set.

### GetInvoiceFooter

`func (o *ClientSettings) GetInvoiceFooter() string`

GetInvoiceFooter returns the InvoiceFooter field if non-nil, zero value otherwise.

### GetInvoiceFooterOk

`func (o *ClientSettings) GetInvoiceFooterOk() (*string, bool)`

GetInvoiceFooterOk returns a tuple with the InvoiceFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFooter

`func (o *ClientSettings) SetInvoiceFooter(v string)`

SetInvoiceFooter sets InvoiceFooter field to given value.

### HasInvoiceFooter

`func (o *ClientSettings) HasInvoiceFooter() bool`

HasInvoiceFooter returns a boolean if a field has been set.

### GetInvoiceLabels

`func (o *ClientSettings) GetInvoiceLabels() string`

GetInvoiceLabels returns the InvoiceLabels field if non-nil, zero value otherwise.

### GetInvoiceLabelsOk

`func (o *ClientSettings) GetInvoiceLabelsOk() (*string, bool)`

GetInvoiceLabelsOk returns a tuple with the InvoiceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLabels

`func (o *ClientSettings) SetInvoiceLabels(v string)`

SetInvoiceLabels sets InvoiceLabels field to given value.

### HasInvoiceLabels

`func (o *ClientSettings) HasInvoiceLabels() bool`

HasInvoiceLabels returns a boolean if a field has been set.

### GetTaxRate1

`func (o *ClientSettings) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *ClientSettings) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *ClientSettings) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *ClientSettings) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxName1

`func (o *ClientSettings) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *ClientSettings) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *ClientSettings) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *ClientSettings) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *ClientSettings) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *ClientSettings) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *ClientSettings) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *ClientSettings) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName2

`func (o *ClientSettings) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *ClientSettings) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *ClientSettings) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *ClientSettings) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *ClientSettings) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *ClientSettings) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *ClientSettings) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *ClientSettings) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetTaxName3

`func (o *ClientSettings) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *ClientSettings) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *ClientSettings) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *ClientSettings) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetPaymentTypeId

`func (o *ClientSettings) GetPaymentTypeId() string`

GetPaymentTypeId returns the PaymentTypeId field if non-nil, zero value otherwise.

### GetPaymentTypeIdOk

`func (o *ClientSettings) GetPaymentTypeIdOk() (*string, bool)`

GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeId

`func (o *ClientSettings) SetPaymentTypeId(v string)`

SetPaymentTypeId sets PaymentTypeId field to given value.

### HasPaymentTypeId

`func (o *ClientSettings) HasPaymentTypeId() bool`

HasPaymentTypeId returns a boolean if a field has been set.

### GetCustomFields

`func (o *ClientSettings) GetCustomFields() string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ClientSettings) GetCustomFieldsOk() (*string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ClientSettings) SetCustomFields(v string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ClientSettings) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEmailFooter

`func (o *ClientSettings) GetEmailFooter() string`

GetEmailFooter returns the EmailFooter field if non-nil, zero value otherwise.

### GetEmailFooterOk

`func (o *ClientSettings) GetEmailFooterOk() (*string, bool)`

GetEmailFooterOk returns a tuple with the EmailFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFooter

`func (o *ClientSettings) SetEmailFooter(v string)`

SetEmailFooter sets EmailFooter field to given value.

### HasEmailFooter

`func (o *ClientSettings) HasEmailFooter() bool`

HasEmailFooter returns a boolean if a field has been set.

### GetEmailSendingMethod

`func (o *ClientSettings) GetEmailSendingMethod() string`

GetEmailSendingMethod returns the EmailSendingMethod field if non-nil, zero value otherwise.

### GetEmailSendingMethodOk

`func (o *ClientSettings) GetEmailSendingMethodOk() (*string, bool)`

GetEmailSendingMethodOk returns a tuple with the EmailSendingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSendingMethod

`func (o *ClientSettings) SetEmailSendingMethod(v string)`

SetEmailSendingMethod sets EmailSendingMethod field to given value.

### HasEmailSendingMethod

`func (o *ClientSettings) HasEmailSendingMethod() bool`

HasEmailSendingMethod returns a boolean if a field has been set.

### GetGmailSendingUserId

`func (o *ClientSettings) GetGmailSendingUserId() string`

GetGmailSendingUserId returns the GmailSendingUserId field if non-nil, zero value otherwise.

### GetGmailSendingUserIdOk

`func (o *ClientSettings) GetGmailSendingUserIdOk() (*string, bool)`

GetGmailSendingUserIdOk returns a tuple with the GmailSendingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmailSendingUserId

`func (o *ClientSettings) SetGmailSendingUserId(v string)`

SetGmailSendingUserId sets GmailSendingUserId field to given value.

### HasGmailSendingUserId

`func (o *ClientSettings) HasGmailSendingUserId() bool`

HasGmailSendingUserId returns a boolean if a field has been set.

### GetEmailSubjectInvoice

`func (o *ClientSettings) GetEmailSubjectInvoice() string`

GetEmailSubjectInvoice returns the EmailSubjectInvoice field if non-nil, zero value otherwise.

### GetEmailSubjectInvoiceOk

`func (o *ClientSettings) GetEmailSubjectInvoiceOk() (*string, bool)`

GetEmailSubjectInvoiceOk returns a tuple with the EmailSubjectInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectInvoice

`func (o *ClientSettings) SetEmailSubjectInvoice(v string)`

SetEmailSubjectInvoice sets EmailSubjectInvoice field to given value.

### HasEmailSubjectInvoice

`func (o *ClientSettings) HasEmailSubjectInvoice() bool`

HasEmailSubjectInvoice returns a boolean if a field has been set.

### GetEmailSubjectQuote

`func (o *ClientSettings) GetEmailSubjectQuote() string`

GetEmailSubjectQuote returns the EmailSubjectQuote field if non-nil, zero value otherwise.

### GetEmailSubjectQuoteOk

`func (o *ClientSettings) GetEmailSubjectQuoteOk() (*string, bool)`

GetEmailSubjectQuoteOk returns a tuple with the EmailSubjectQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectQuote

`func (o *ClientSettings) SetEmailSubjectQuote(v string)`

SetEmailSubjectQuote sets EmailSubjectQuote field to given value.

### HasEmailSubjectQuote

`func (o *ClientSettings) HasEmailSubjectQuote() bool`

HasEmailSubjectQuote returns a boolean if a field has been set.

### GetEmailSubjectPayment

`func (o *ClientSettings) GetEmailSubjectPayment() string`

GetEmailSubjectPayment returns the EmailSubjectPayment field if non-nil, zero value otherwise.

### GetEmailSubjectPaymentOk

`func (o *ClientSettings) GetEmailSubjectPaymentOk() (*string, bool)`

GetEmailSubjectPaymentOk returns a tuple with the EmailSubjectPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectPayment

`func (o *ClientSettings) SetEmailSubjectPayment(v string)`

SetEmailSubjectPayment sets EmailSubjectPayment field to given value.

### HasEmailSubjectPayment

`func (o *ClientSettings) HasEmailSubjectPayment() bool`

HasEmailSubjectPayment returns a boolean if a field has been set.

### GetEmailTemplateInvoice

`func (o *ClientSettings) GetEmailTemplateInvoice() string`

GetEmailTemplateInvoice returns the EmailTemplateInvoice field if non-nil, zero value otherwise.

### GetEmailTemplateInvoiceOk

`func (o *ClientSettings) GetEmailTemplateInvoiceOk() (*string, bool)`

GetEmailTemplateInvoiceOk returns a tuple with the EmailTemplateInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateInvoice

`func (o *ClientSettings) SetEmailTemplateInvoice(v string)`

SetEmailTemplateInvoice sets EmailTemplateInvoice field to given value.

### HasEmailTemplateInvoice

`func (o *ClientSettings) HasEmailTemplateInvoice() bool`

HasEmailTemplateInvoice returns a boolean if a field has been set.

### GetEmailTemplateQuote

`func (o *ClientSettings) GetEmailTemplateQuote() string`

GetEmailTemplateQuote returns the EmailTemplateQuote field if non-nil, zero value otherwise.

### GetEmailTemplateQuoteOk

`func (o *ClientSettings) GetEmailTemplateQuoteOk() (*string, bool)`

GetEmailTemplateQuoteOk returns a tuple with the EmailTemplateQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateQuote

`func (o *ClientSettings) SetEmailTemplateQuote(v string)`

SetEmailTemplateQuote sets EmailTemplateQuote field to given value.

### HasEmailTemplateQuote

`func (o *ClientSettings) HasEmailTemplateQuote() bool`

HasEmailTemplateQuote returns a boolean if a field has been set.

### GetEmailTemplatePayment

`func (o *ClientSettings) GetEmailTemplatePayment() string`

GetEmailTemplatePayment returns the EmailTemplatePayment field if non-nil, zero value otherwise.

### GetEmailTemplatePaymentOk

`func (o *ClientSettings) GetEmailTemplatePaymentOk() (*string, bool)`

GetEmailTemplatePaymentOk returns a tuple with the EmailTemplatePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplatePayment

`func (o *ClientSettings) SetEmailTemplatePayment(v string)`

SetEmailTemplatePayment sets EmailTemplatePayment field to given value.

### HasEmailTemplatePayment

`func (o *ClientSettings) HasEmailTemplatePayment() bool`

HasEmailTemplatePayment returns a boolean if a field has been set.

### GetEmailSubjectReminder1

`func (o *ClientSettings) GetEmailSubjectReminder1() string`

GetEmailSubjectReminder1 returns the EmailSubjectReminder1 field if non-nil, zero value otherwise.

### GetEmailSubjectReminder1Ok

`func (o *ClientSettings) GetEmailSubjectReminder1Ok() (*string, bool)`

GetEmailSubjectReminder1Ok returns a tuple with the EmailSubjectReminder1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectReminder1

`func (o *ClientSettings) SetEmailSubjectReminder1(v string)`

SetEmailSubjectReminder1 sets EmailSubjectReminder1 field to given value.

### HasEmailSubjectReminder1

`func (o *ClientSettings) HasEmailSubjectReminder1() bool`

HasEmailSubjectReminder1 returns a boolean if a field has been set.

### GetEmailSubjectReminder2

`func (o *ClientSettings) GetEmailSubjectReminder2() string`

GetEmailSubjectReminder2 returns the EmailSubjectReminder2 field if non-nil, zero value otherwise.

### GetEmailSubjectReminder2Ok

`func (o *ClientSettings) GetEmailSubjectReminder2Ok() (*string, bool)`

GetEmailSubjectReminder2Ok returns a tuple with the EmailSubjectReminder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectReminder2

`func (o *ClientSettings) SetEmailSubjectReminder2(v string)`

SetEmailSubjectReminder2 sets EmailSubjectReminder2 field to given value.

### HasEmailSubjectReminder2

`func (o *ClientSettings) HasEmailSubjectReminder2() bool`

HasEmailSubjectReminder2 returns a boolean if a field has been set.

### GetEmailSubjectReminder3

`func (o *ClientSettings) GetEmailSubjectReminder3() string`

GetEmailSubjectReminder3 returns the EmailSubjectReminder3 field if non-nil, zero value otherwise.

### GetEmailSubjectReminder3Ok

`func (o *ClientSettings) GetEmailSubjectReminder3Ok() (*string, bool)`

GetEmailSubjectReminder3Ok returns a tuple with the EmailSubjectReminder3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectReminder3

`func (o *ClientSettings) SetEmailSubjectReminder3(v string)`

SetEmailSubjectReminder3 sets EmailSubjectReminder3 field to given value.

### HasEmailSubjectReminder3

`func (o *ClientSettings) HasEmailSubjectReminder3() bool`

HasEmailSubjectReminder3 returns a boolean if a field has been set.

### GetEmailSubjectReminderEndless

`func (o *ClientSettings) GetEmailSubjectReminderEndless() string`

GetEmailSubjectReminderEndless returns the EmailSubjectReminderEndless field if non-nil, zero value otherwise.

### GetEmailSubjectReminderEndlessOk

`func (o *ClientSettings) GetEmailSubjectReminderEndlessOk() (*string, bool)`

GetEmailSubjectReminderEndlessOk returns a tuple with the EmailSubjectReminderEndless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectReminderEndless

`func (o *ClientSettings) SetEmailSubjectReminderEndless(v string)`

SetEmailSubjectReminderEndless sets EmailSubjectReminderEndless field to given value.

### HasEmailSubjectReminderEndless

`func (o *ClientSettings) HasEmailSubjectReminderEndless() bool`

HasEmailSubjectReminderEndless returns a boolean if a field has been set.

### GetEmailTemplateReminder1

`func (o *ClientSettings) GetEmailTemplateReminder1() string`

GetEmailTemplateReminder1 returns the EmailTemplateReminder1 field if non-nil, zero value otherwise.

### GetEmailTemplateReminder1Ok

`func (o *ClientSettings) GetEmailTemplateReminder1Ok() (*string, bool)`

GetEmailTemplateReminder1Ok returns a tuple with the EmailTemplateReminder1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateReminder1

`func (o *ClientSettings) SetEmailTemplateReminder1(v string)`

SetEmailTemplateReminder1 sets EmailTemplateReminder1 field to given value.

### HasEmailTemplateReminder1

`func (o *ClientSettings) HasEmailTemplateReminder1() bool`

HasEmailTemplateReminder1 returns a boolean if a field has been set.

### GetEmailTemplateReminder2

`func (o *ClientSettings) GetEmailTemplateReminder2() string`

GetEmailTemplateReminder2 returns the EmailTemplateReminder2 field if non-nil, zero value otherwise.

### GetEmailTemplateReminder2Ok

`func (o *ClientSettings) GetEmailTemplateReminder2Ok() (*string, bool)`

GetEmailTemplateReminder2Ok returns a tuple with the EmailTemplateReminder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateReminder2

`func (o *ClientSettings) SetEmailTemplateReminder2(v string)`

SetEmailTemplateReminder2 sets EmailTemplateReminder2 field to given value.

### HasEmailTemplateReminder2

`func (o *ClientSettings) HasEmailTemplateReminder2() bool`

HasEmailTemplateReminder2 returns a boolean if a field has been set.

### GetEmailTemplateReminder3

`func (o *ClientSettings) GetEmailTemplateReminder3() string`

GetEmailTemplateReminder3 returns the EmailTemplateReminder3 field if non-nil, zero value otherwise.

### GetEmailTemplateReminder3Ok

`func (o *ClientSettings) GetEmailTemplateReminder3Ok() (*string, bool)`

GetEmailTemplateReminder3Ok returns a tuple with the EmailTemplateReminder3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateReminder3

`func (o *ClientSettings) SetEmailTemplateReminder3(v string)`

SetEmailTemplateReminder3 sets EmailTemplateReminder3 field to given value.

### HasEmailTemplateReminder3

`func (o *ClientSettings) HasEmailTemplateReminder3() bool`

HasEmailTemplateReminder3 returns a boolean if a field has been set.

### GetEmailTemplateReminderEndless

`func (o *ClientSettings) GetEmailTemplateReminderEndless() string`

GetEmailTemplateReminderEndless returns the EmailTemplateReminderEndless field if non-nil, zero value otherwise.

### GetEmailTemplateReminderEndlessOk

`func (o *ClientSettings) GetEmailTemplateReminderEndlessOk() (*string, bool)`

GetEmailTemplateReminderEndlessOk returns a tuple with the EmailTemplateReminderEndless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateReminderEndless

`func (o *ClientSettings) SetEmailTemplateReminderEndless(v string)`

SetEmailTemplateReminderEndless sets EmailTemplateReminderEndless field to given value.

### HasEmailTemplateReminderEndless

`func (o *ClientSettings) HasEmailTemplateReminderEndless() bool`

HasEmailTemplateReminderEndless returns a boolean if a field has been set.

### GetEnablePortalPassword

`func (o *ClientSettings) GetEnablePortalPassword() bool`

GetEnablePortalPassword returns the EnablePortalPassword field if non-nil, zero value otherwise.

### GetEnablePortalPasswordOk

`func (o *ClientSettings) GetEnablePortalPasswordOk() (*bool, bool)`

GetEnablePortalPasswordOk returns a tuple with the EnablePortalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePortalPassword

`func (o *ClientSettings) SetEnablePortalPassword(v bool)`

SetEnablePortalPassword sets EnablePortalPassword field to given value.

### HasEnablePortalPassword

`func (o *ClientSettings) HasEnablePortalPassword() bool`

HasEnablePortalPassword returns a boolean if a field has been set.

### GetShowAcceptInvoiceTerms

`func (o *ClientSettings) GetShowAcceptInvoiceTerms() bool`

GetShowAcceptInvoiceTerms returns the ShowAcceptInvoiceTerms field if non-nil, zero value otherwise.

### GetShowAcceptInvoiceTermsOk

`func (o *ClientSettings) GetShowAcceptInvoiceTermsOk() (*bool, bool)`

GetShowAcceptInvoiceTermsOk returns a tuple with the ShowAcceptInvoiceTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAcceptInvoiceTerms

`func (o *ClientSettings) SetShowAcceptInvoiceTerms(v bool)`

SetShowAcceptInvoiceTerms sets ShowAcceptInvoiceTerms field to given value.

### HasShowAcceptInvoiceTerms

`func (o *ClientSettings) HasShowAcceptInvoiceTerms() bool`

HasShowAcceptInvoiceTerms returns a boolean if a field has been set.

### GetShowAcceptQuoteTerms

`func (o *ClientSettings) GetShowAcceptQuoteTerms() bool`

GetShowAcceptQuoteTerms returns the ShowAcceptQuoteTerms field if non-nil, zero value otherwise.

### GetShowAcceptQuoteTermsOk

`func (o *ClientSettings) GetShowAcceptQuoteTermsOk() (*bool, bool)`

GetShowAcceptQuoteTermsOk returns a tuple with the ShowAcceptQuoteTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAcceptQuoteTerms

`func (o *ClientSettings) SetShowAcceptQuoteTerms(v bool)`

SetShowAcceptQuoteTerms sets ShowAcceptQuoteTerms field to given value.

### HasShowAcceptQuoteTerms

`func (o *ClientSettings) HasShowAcceptQuoteTerms() bool`

HasShowAcceptQuoteTerms returns a boolean if a field has been set.

### GetRequireInvoiceSignature

`func (o *ClientSettings) GetRequireInvoiceSignature() bool`

GetRequireInvoiceSignature returns the RequireInvoiceSignature field if non-nil, zero value otherwise.

### GetRequireInvoiceSignatureOk

`func (o *ClientSettings) GetRequireInvoiceSignatureOk() (*bool, bool)`

GetRequireInvoiceSignatureOk returns a tuple with the RequireInvoiceSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireInvoiceSignature

`func (o *ClientSettings) SetRequireInvoiceSignature(v bool)`

SetRequireInvoiceSignature sets RequireInvoiceSignature field to given value.

### HasRequireInvoiceSignature

`func (o *ClientSettings) HasRequireInvoiceSignature() bool`

HasRequireInvoiceSignature returns a boolean if a field has been set.

### GetRequireQuoteSignature

`func (o *ClientSettings) GetRequireQuoteSignature() bool`

GetRequireQuoteSignature returns the RequireQuoteSignature field if non-nil, zero value otherwise.

### GetRequireQuoteSignatureOk

`func (o *ClientSettings) GetRequireQuoteSignatureOk() (*bool, bool)`

GetRequireQuoteSignatureOk returns a tuple with the RequireQuoteSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireQuoteSignature

`func (o *ClientSettings) SetRequireQuoteSignature(v bool)`

SetRequireQuoteSignature sets RequireQuoteSignature field to given value.

### HasRequireQuoteSignature

`func (o *ClientSettings) HasRequireQuoteSignature() bool`

HasRequireQuoteSignature returns a boolean if a field has been set.

### GetName

`func (o *ClientSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompanyLogo

`func (o *ClientSettings) GetCompanyLogo() map[string]interface{}`

GetCompanyLogo returns the CompanyLogo field if non-nil, zero value otherwise.

### GetCompanyLogoOk

`func (o *ClientSettings) GetCompanyLogoOk() (*map[string]interface{}, bool)`

GetCompanyLogoOk returns a tuple with the CompanyLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLogo

`func (o *ClientSettings) SetCompanyLogo(v map[string]interface{})`

SetCompanyLogo sets CompanyLogo field to given value.

### HasCompanyLogo

`func (o *ClientSettings) HasCompanyLogo() bool`

HasCompanyLogo returns a boolean if a field has been set.

### GetWebsite

`func (o *ClientSettings) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ClientSettings) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ClientSettings) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ClientSettings) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetAddress1

`func (o *ClientSettings) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ClientSettings) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ClientSettings) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ClientSettings) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *ClientSettings) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ClientSettings) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ClientSettings) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ClientSettings) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *ClientSettings) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ClientSettings) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ClientSettings) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ClientSettings) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *ClientSettings) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClientSettings) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClientSettings) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClientSettings) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *ClientSettings) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ClientSettings) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ClientSettings) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ClientSettings) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPhone

`func (o *ClientSettings) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ClientSettings) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ClientSettings) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ClientSettings) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *ClientSettings) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ClientSettings) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ClientSettings) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ClientSettings) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCountryId

`func (o *ClientSettings) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ClientSettings) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ClientSettings) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ClientSettings) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetVatNumber

`func (o *ClientSettings) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *ClientSettings) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *ClientSettings) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *ClientSettings) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ClientSettings) GetPageSize() string`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ClientSettings) GetPageSizeOk() (*string, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ClientSettings) SetPageSize(v string)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ClientSettings) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetFontSize

`func (o *ClientSettings) GetFontSize() float32`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *ClientSettings) GetFontSizeOk() (*float32, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *ClientSettings) SetFontSize(v float32)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *ClientSettings) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### GetPrimaryFont

`func (o *ClientSettings) GetPrimaryFont() string`

GetPrimaryFont returns the PrimaryFont field if non-nil, zero value otherwise.

### GetPrimaryFontOk

`func (o *ClientSettings) GetPrimaryFontOk() (*string, bool)`

GetPrimaryFontOk returns a tuple with the PrimaryFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFont

`func (o *ClientSettings) SetPrimaryFont(v string)`

SetPrimaryFont sets PrimaryFont field to given value.

### HasPrimaryFont

`func (o *ClientSettings) HasPrimaryFont() bool`

HasPrimaryFont returns a boolean if a field has been set.

### GetSecondaryFont

`func (o *ClientSettings) GetSecondaryFont() string`

GetSecondaryFont returns the SecondaryFont field if non-nil, zero value otherwise.

### GetSecondaryFontOk

`func (o *ClientSettings) GetSecondaryFontOk() (*string, bool)`

GetSecondaryFontOk returns a tuple with the SecondaryFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFont

`func (o *ClientSettings) SetSecondaryFont(v string)`

SetSecondaryFont sets SecondaryFont field to given value.

### HasSecondaryFont

`func (o *ClientSettings) HasSecondaryFont() bool`

HasSecondaryFont returns a boolean if a field has been set.

### GetHidePaidToDate

`func (o *ClientSettings) GetHidePaidToDate() bool`

GetHidePaidToDate returns the HidePaidToDate field if non-nil, zero value otherwise.

### GetHidePaidToDateOk

`func (o *ClientSettings) GetHidePaidToDateOk() (*bool, bool)`

GetHidePaidToDateOk returns a tuple with the HidePaidToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePaidToDate

`func (o *ClientSettings) SetHidePaidToDate(v bool)`

SetHidePaidToDate sets HidePaidToDate field to given value.

### HasHidePaidToDate

`func (o *ClientSettings) HasHidePaidToDate() bool`

HasHidePaidToDate returns a boolean if a field has been set.

### GetEmbedDocuments

`func (o *ClientSettings) GetEmbedDocuments() bool`

GetEmbedDocuments returns the EmbedDocuments field if non-nil, zero value otherwise.

### GetEmbedDocumentsOk

`func (o *ClientSettings) GetEmbedDocumentsOk() (*bool, bool)`

GetEmbedDocumentsOk returns a tuple with the EmbedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedDocuments

`func (o *ClientSettings) SetEmbedDocuments(v bool)`

SetEmbedDocuments sets EmbedDocuments field to given value.

### HasEmbedDocuments

`func (o *ClientSettings) HasEmbedDocuments() bool`

HasEmbedDocuments returns a boolean if a field has been set.

### GetAllPagesHeader

`func (o *ClientSettings) GetAllPagesHeader() bool`

GetAllPagesHeader returns the AllPagesHeader field if non-nil, zero value otherwise.

### GetAllPagesHeaderOk

`func (o *ClientSettings) GetAllPagesHeaderOk() (*bool, bool)`

GetAllPagesHeaderOk returns a tuple with the AllPagesHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPagesHeader

`func (o *ClientSettings) SetAllPagesHeader(v bool)`

SetAllPagesHeader sets AllPagesHeader field to given value.

### HasAllPagesHeader

`func (o *ClientSettings) HasAllPagesHeader() bool`

HasAllPagesHeader returns a boolean if a field has been set.

### GetAllPagesFooter

`func (o *ClientSettings) GetAllPagesFooter() bool`

GetAllPagesFooter returns the AllPagesFooter field if non-nil, zero value otherwise.

### GetAllPagesFooterOk

`func (o *ClientSettings) GetAllPagesFooterOk() (*bool, bool)`

GetAllPagesFooterOk returns a tuple with the AllPagesFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPagesFooter

`func (o *ClientSettings) SetAllPagesFooter(v bool)`

SetAllPagesFooter sets AllPagesFooter field to given value.

### HasAllPagesFooter

`func (o *ClientSettings) HasAllPagesFooter() bool`

HasAllPagesFooter returns a boolean if a field has been set.

### GetDocumentEmailAttachment

`func (o *ClientSettings) GetDocumentEmailAttachment() bool`

GetDocumentEmailAttachment returns the DocumentEmailAttachment field if non-nil, zero value otherwise.

### GetDocumentEmailAttachmentOk

`func (o *ClientSettings) GetDocumentEmailAttachmentOk() (*bool, bool)`

GetDocumentEmailAttachmentOk returns a tuple with the DocumentEmailAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentEmailAttachment

`func (o *ClientSettings) SetDocumentEmailAttachment(v bool)`

SetDocumentEmailAttachment sets DocumentEmailAttachment field to given value.

### HasDocumentEmailAttachment

`func (o *ClientSettings) HasDocumentEmailAttachment() bool`

HasDocumentEmailAttachment returns a boolean if a field has been set.

### GetEnableClientPortalPassword

`func (o *ClientSettings) GetEnableClientPortalPassword() bool`

GetEnableClientPortalPassword returns the EnableClientPortalPassword field if non-nil, zero value otherwise.

### GetEnableClientPortalPasswordOk

`func (o *ClientSettings) GetEnableClientPortalPasswordOk() (*bool, bool)`

GetEnableClientPortalPasswordOk returns a tuple with the EnableClientPortalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientPortalPassword

`func (o *ClientSettings) SetEnableClientPortalPassword(v bool)`

SetEnableClientPortalPassword sets EnableClientPortalPassword field to given value.

### HasEnableClientPortalPassword

`func (o *ClientSettings) HasEnableClientPortalPassword() bool`

HasEnableClientPortalPassword returns a boolean if a field has been set.

### GetEnableEmailMarkup

`func (o *ClientSettings) GetEnableEmailMarkup() bool`

GetEnableEmailMarkup returns the EnableEmailMarkup field if non-nil, zero value otherwise.

### GetEnableEmailMarkupOk

`func (o *ClientSettings) GetEnableEmailMarkupOk() (*bool, bool)`

GetEnableEmailMarkupOk returns a tuple with the EnableEmailMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailMarkup

`func (o *ClientSettings) SetEnableEmailMarkup(v bool)`

SetEnableEmailMarkup sets EnableEmailMarkup field to given value.

### HasEnableEmailMarkup

`func (o *ClientSettings) HasEnableEmailMarkup() bool`

HasEnableEmailMarkup returns a boolean if a field has been set.

### GetEnableClientPortalDashboard

`func (o *ClientSettings) GetEnableClientPortalDashboard() bool`

GetEnableClientPortalDashboard returns the EnableClientPortalDashboard field if non-nil, zero value otherwise.

### GetEnableClientPortalDashboardOk

`func (o *ClientSettings) GetEnableClientPortalDashboardOk() (*bool, bool)`

GetEnableClientPortalDashboardOk returns a tuple with the EnableClientPortalDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientPortalDashboard

`func (o *ClientSettings) SetEnableClientPortalDashboard(v bool)`

SetEnableClientPortalDashboard sets EnableClientPortalDashboard field to given value.

### HasEnableClientPortalDashboard

`func (o *ClientSettings) HasEnableClientPortalDashboard() bool`

HasEnableClientPortalDashboard returns a boolean if a field has been set.

### GetEnableClientPortal

`func (o *ClientSettings) GetEnableClientPortal() bool`

GetEnableClientPortal returns the EnableClientPortal field if non-nil, zero value otherwise.

### GetEnableClientPortalOk

`func (o *ClientSettings) GetEnableClientPortalOk() (*bool, bool)`

GetEnableClientPortalOk returns a tuple with the EnableClientPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientPortal

`func (o *ClientSettings) SetEnableClientPortal(v bool)`

SetEnableClientPortal sets EnableClientPortal field to given value.

### HasEnableClientPortal

`func (o *ClientSettings) HasEnableClientPortal() bool`

HasEnableClientPortal returns a boolean if a field has been set.

### GetEmailTemplateStatement

`func (o *ClientSettings) GetEmailTemplateStatement() string`

GetEmailTemplateStatement returns the EmailTemplateStatement field if non-nil, zero value otherwise.

### GetEmailTemplateStatementOk

`func (o *ClientSettings) GetEmailTemplateStatementOk() (*string, bool)`

GetEmailTemplateStatementOk returns a tuple with the EmailTemplateStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateStatement

`func (o *ClientSettings) SetEmailTemplateStatement(v string)`

SetEmailTemplateStatement sets EmailTemplateStatement field to given value.

### HasEmailTemplateStatement

`func (o *ClientSettings) HasEmailTemplateStatement() bool`

HasEmailTemplateStatement returns a boolean if a field has been set.

### GetEmailSubjectStatement

`func (o *ClientSettings) GetEmailSubjectStatement() string`

GetEmailSubjectStatement returns the EmailSubjectStatement field if non-nil, zero value otherwise.

### GetEmailSubjectStatementOk

`func (o *ClientSettings) GetEmailSubjectStatementOk() (*string, bool)`

GetEmailSubjectStatementOk returns a tuple with the EmailSubjectStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectStatement

`func (o *ClientSettings) SetEmailSubjectStatement(v string)`

SetEmailSubjectStatement sets EmailSubjectStatement field to given value.

### HasEmailSubjectStatement

`func (o *ClientSettings) HasEmailSubjectStatement() bool`

HasEmailSubjectStatement returns a boolean if a field has been set.

### GetSignatureOnPdf

`func (o *ClientSettings) GetSignatureOnPdf() bool`

GetSignatureOnPdf returns the SignatureOnPdf field if non-nil, zero value otherwise.

### GetSignatureOnPdfOk

`func (o *ClientSettings) GetSignatureOnPdfOk() (*bool, bool)`

GetSignatureOnPdfOk returns a tuple with the SignatureOnPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureOnPdf

`func (o *ClientSettings) SetSignatureOnPdf(v bool)`

SetSignatureOnPdf sets SignatureOnPdf field to given value.

### HasSignatureOnPdf

`func (o *ClientSettings) HasSignatureOnPdf() bool`

HasSignatureOnPdf returns a boolean if a field has been set.

### GetQuoteFooter

`func (o *ClientSettings) GetQuoteFooter() string`

GetQuoteFooter returns the QuoteFooter field if non-nil, zero value otherwise.

### GetQuoteFooterOk

`func (o *ClientSettings) GetQuoteFooterOk() (*string, bool)`

GetQuoteFooterOk returns a tuple with the QuoteFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteFooter

`func (o *ClientSettings) SetQuoteFooter(v string)`

SetQuoteFooter sets QuoteFooter field to given value.

### HasQuoteFooter

`func (o *ClientSettings) HasQuoteFooter() bool`

HasQuoteFooter returns a boolean if a field has been set.

### GetEmailSubjectCustom1

`func (o *ClientSettings) GetEmailSubjectCustom1() string`

GetEmailSubjectCustom1 returns the EmailSubjectCustom1 field if non-nil, zero value otherwise.

### GetEmailSubjectCustom1Ok

`func (o *ClientSettings) GetEmailSubjectCustom1Ok() (*string, bool)`

GetEmailSubjectCustom1Ok returns a tuple with the EmailSubjectCustom1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectCustom1

`func (o *ClientSettings) SetEmailSubjectCustom1(v string)`

SetEmailSubjectCustom1 sets EmailSubjectCustom1 field to given value.

### HasEmailSubjectCustom1

`func (o *ClientSettings) HasEmailSubjectCustom1() bool`

HasEmailSubjectCustom1 returns a boolean if a field has been set.

### GetEmailSubjectCustom2

`func (o *ClientSettings) GetEmailSubjectCustom2() string`

GetEmailSubjectCustom2 returns the EmailSubjectCustom2 field if non-nil, zero value otherwise.

### GetEmailSubjectCustom2Ok

`func (o *ClientSettings) GetEmailSubjectCustom2Ok() (*string, bool)`

GetEmailSubjectCustom2Ok returns a tuple with the EmailSubjectCustom2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectCustom2

`func (o *ClientSettings) SetEmailSubjectCustom2(v string)`

SetEmailSubjectCustom2 sets EmailSubjectCustom2 field to given value.

### HasEmailSubjectCustom2

`func (o *ClientSettings) HasEmailSubjectCustom2() bool`

HasEmailSubjectCustom2 returns a boolean if a field has been set.

### GetEmailSubjectCustom3

`func (o *ClientSettings) GetEmailSubjectCustom3() string`

GetEmailSubjectCustom3 returns the EmailSubjectCustom3 field if non-nil, zero value otherwise.

### GetEmailSubjectCustom3Ok

`func (o *ClientSettings) GetEmailSubjectCustom3Ok() (*string, bool)`

GetEmailSubjectCustom3Ok returns a tuple with the EmailSubjectCustom3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectCustom3

`func (o *ClientSettings) SetEmailSubjectCustom3(v string)`

SetEmailSubjectCustom3 sets EmailSubjectCustom3 field to given value.

### HasEmailSubjectCustom3

`func (o *ClientSettings) HasEmailSubjectCustom3() bool`

HasEmailSubjectCustom3 returns a boolean if a field has been set.

### GetEmailTemplateCustom1

`func (o *ClientSettings) GetEmailTemplateCustom1() string`

GetEmailTemplateCustom1 returns the EmailTemplateCustom1 field if non-nil, zero value otherwise.

### GetEmailTemplateCustom1Ok

`func (o *ClientSettings) GetEmailTemplateCustom1Ok() (*string, bool)`

GetEmailTemplateCustom1Ok returns a tuple with the EmailTemplateCustom1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateCustom1

`func (o *ClientSettings) SetEmailTemplateCustom1(v string)`

SetEmailTemplateCustom1 sets EmailTemplateCustom1 field to given value.

### HasEmailTemplateCustom1

`func (o *ClientSettings) HasEmailTemplateCustom1() bool`

HasEmailTemplateCustom1 returns a boolean if a field has been set.

### GetEmailTemplateCustom2

`func (o *ClientSettings) GetEmailTemplateCustom2() string`

GetEmailTemplateCustom2 returns the EmailTemplateCustom2 field if non-nil, zero value otherwise.

### GetEmailTemplateCustom2Ok

`func (o *ClientSettings) GetEmailTemplateCustom2Ok() (*string, bool)`

GetEmailTemplateCustom2Ok returns a tuple with the EmailTemplateCustom2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateCustom2

`func (o *ClientSettings) SetEmailTemplateCustom2(v string)`

SetEmailTemplateCustom2 sets EmailTemplateCustom2 field to given value.

### HasEmailTemplateCustom2

`func (o *ClientSettings) HasEmailTemplateCustom2() bool`

HasEmailTemplateCustom2 returns a boolean if a field has been set.

### GetEmailTemplateCustom3

`func (o *ClientSettings) GetEmailTemplateCustom3() string`

GetEmailTemplateCustom3 returns the EmailTemplateCustom3 field if non-nil, zero value otherwise.

### GetEmailTemplateCustom3Ok

`func (o *ClientSettings) GetEmailTemplateCustom3Ok() (*string, bool)`

GetEmailTemplateCustom3Ok returns a tuple with the EmailTemplateCustom3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplateCustom3

`func (o *ClientSettings) SetEmailTemplateCustom3(v string)`

SetEmailTemplateCustom3 sets EmailTemplateCustom3 field to given value.

### HasEmailTemplateCustom3

`func (o *ClientSettings) HasEmailTemplateCustom3() bool`

HasEmailTemplateCustom3 returns a boolean if a field has been set.

### GetEnableReminder1

`func (o *ClientSettings) GetEnableReminder1() bool`

GetEnableReminder1 returns the EnableReminder1 field if non-nil, zero value otherwise.

### GetEnableReminder1Ok

`func (o *ClientSettings) GetEnableReminder1Ok() (*bool, bool)`

GetEnableReminder1Ok returns a tuple with the EnableReminder1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReminder1

`func (o *ClientSettings) SetEnableReminder1(v bool)`

SetEnableReminder1 sets EnableReminder1 field to given value.

### HasEnableReminder1

`func (o *ClientSettings) HasEnableReminder1() bool`

HasEnableReminder1 returns a boolean if a field has been set.

### GetEnableReminder2

`func (o *ClientSettings) GetEnableReminder2() bool`

GetEnableReminder2 returns the EnableReminder2 field if non-nil, zero value otherwise.

### GetEnableReminder2Ok

`func (o *ClientSettings) GetEnableReminder2Ok() (*bool, bool)`

GetEnableReminder2Ok returns a tuple with the EnableReminder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReminder2

`func (o *ClientSettings) SetEnableReminder2(v bool)`

SetEnableReminder2 sets EnableReminder2 field to given value.

### HasEnableReminder2

`func (o *ClientSettings) HasEnableReminder2() bool`

HasEnableReminder2 returns a boolean if a field has been set.

### GetEnableReminder3

`func (o *ClientSettings) GetEnableReminder3() bool`

GetEnableReminder3 returns the EnableReminder3 field if non-nil, zero value otherwise.

### GetEnableReminder3Ok

`func (o *ClientSettings) GetEnableReminder3Ok() (*bool, bool)`

GetEnableReminder3Ok returns a tuple with the EnableReminder3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReminder3

`func (o *ClientSettings) SetEnableReminder3(v bool)`

SetEnableReminder3 sets EnableReminder3 field to given value.

### HasEnableReminder3

`func (o *ClientSettings) HasEnableReminder3() bool`

HasEnableReminder3 returns a boolean if a field has been set.

### GetNumDaysReminder1

`func (o *ClientSettings) GetNumDaysReminder1() float32`

GetNumDaysReminder1 returns the NumDaysReminder1 field if non-nil, zero value otherwise.

### GetNumDaysReminder1Ok

`func (o *ClientSettings) GetNumDaysReminder1Ok() (*float32, bool)`

GetNumDaysReminder1Ok returns a tuple with the NumDaysReminder1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDaysReminder1

`func (o *ClientSettings) SetNumDaysReminder1(v float32)`

SetNumDaysReminder1 sets NumDaysReminder1 field to given value.

### HasNumDaysReminder1

`func (o *ClientSettings) HasNumDaysReminder1() bool`

HasNumDaysReminder1 returns a boolean if a field has been set.

### GetNumDaysReminder2

`func (o *ClientSettings) GetNumDaysReminder2() float32`

GetNumDaysReminder2 returns the NumDaysReminder2 field if non-nil, zero value otherwise.

### GetNumDaysReminder2Ok

`func (o *ClientSettings) GetNumDaysReminder2Ok() (*float32, bool)`

GetNumDaysReminder2Ok returns a tuple with the NumDaysReminder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDaysReminder2

`func (o *ClientSettings) SetNumDaysReminder2(v float32)`

SetNumDaysReminder2 sets NumDaysReminder2 field to given value.

### HasNumDaysReminder2

`func (o *ClientSettings) HasNumDaysReminder2() bool`

HasNumDaysReminder2 returns a boolean if a field has been set.

### GetNumDaysReminder3

`func (o *ClientSettings) GetNumDaysReminder3() float32`

GetNumDaysReminder3 returns the NumDaysReminder3 field if non-nil, zero value otherwise.

### GetNumDaysReminder3Ok

`func (o *ClientSettings) GetNumDaysReminder3Ok() (*float32, bool)`

GetNumDaysReminder3Ok returns a tuple with the NumDaysReminder3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDaysReminder3

`func (o *ClientSettings) SetNumDaysReminder3(v float32)`

SetNumDaysReminder3 sets NumDaysReminder3 field to given value.

### HasNumDaysReminder3

`func (o *ClientSettings) HasNumDaysReminder3() bool`

HasNumDaysReminder3 returns a boolean if a field has been set.

### GetScheduleReminder1

`func (o *ClientSettings) GetScheduleReminder1() string`

GetScheduleReminder1 returns the ScheduleReminder1 field if non-nil, zero value otherwise.

### GetScheduleReminder1Ok

`func (o *ClientSettings) GetScheduleReminder1Ok() (*string, bool)`

GetScheduleReminder1Ok returns a tuple with the ScheduleReminder1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleReminder1

`func (o *ClientSettings) SetScheduleReminder1(v string)`

SetScheduleReminder1 sets ScheduleReminder1 field to given value.

### HasScheduleReminder1

`func (o *ClientSettings) HasScheduleReminder1() bool`

HasScheduleReminder1 returns a boolean if a field has been set.

### GetScheduleReminder2

`func (o *ClientSettings) GetScheduleReminder2() string`

GetScheduleReminder2 returns the ScheduleReminder2 field if non-nil, zero value otherwise.

### GetScheduleReminder2Ok

`func (o *ClientSettings) GetScheduleReminder2Ok() (*string, bool)`

GetScheduleReminder2Ok returns a tuple with the ScheduleReminder2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleReminder2

`func (o *ClientSettings) SetScheduleReminder2(v string)`

SetScheduleReminder2 sets ScheduleReminder2 field to given value.

### HasScheduleReminder2

`func (o *ClientSettings) HasScheduleReminder2() bool`

HasScheduleReminder2 returns a boolean if a field has been set.

### GetScheduleReminder3

`func (o *ClientSettings) GetScheduleReminder3() string`

GetScheduleReminder3 returns the ScheduleReminder3 field if non-nil, zero value otherwise.

### GetScheduleReminder3Ok

`func (o *ClientSettings) GetScheduleReminder3Ok() (*string, bool)`

GetScheduleReminder3Ok returns a tuple with the ScheduleReminder3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleReminder3

`func (o *ClientSettings) SetScheduleReminder3(v string)`

SetScheduleReminder3 sets ScheduleReminder3 field to given value.

### HasScheduleReminder3

`func (o *ClientSettings) HasScheduleReminder3() bool`

HasScheduleReminder3 returns a boolean if a field has been set.

### GetLateFeeAmount1

`func (o *ClientSettings) GetLateFeeAmount1() float32`

GetLateFeeAmount1 returns the LateFeeAmount1 field if non-nil, zero value otherwise.

### GetLateFeeAmount1Ok

`func (o *ClientSettings) GetLateFeeAmount1Ok() (*float32, bool)`

GetLateFeeAmount1Ok returns a tuple with the LateFeeAmount1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateFeeAmount1

`func (o *ClientSettings) SetLateFeeAmount1(v float32)`

SetLateFeeAmount1 sets LateFeeAmount1 field to given value.

### HasLateFeeAmount1

`func (o *ClientSettings) HasLateFeeAmount1() bool`

HasLateFeeAmount1 returns a boolean if a field has been set.

### GetLateFeeAmount2

`func (o *ClientSettings) GetLateFeeAmount2() float32`

GetLateFeeAmount2 returns the LateFeeAmount2 field if non-nil, zero value otherwise.

### GetLateFeeAmount2Ok

`func (o *ClientSettings) GetLateFeeAmount2Ok() (*float32, bool)`

GetLateFeeAmount2Ok returns a tuple with the LateFeeAmount2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateFeeAmount2

`func (o *ClientSettings) SetLateFeeAmount2(v float32)`

SetLateFeeAmount2 sets LateFeeAmount2 field to given value.

### HasLateFeeAmount2

`func (o *ClientSettings) HasLateFeeAmount2() bool`

HasLateFeeAmount2 returns a boolean if a field has been set.

### GetLateFeeAmount3

`func (o *ClientSettings) GetLateFeeAmount3() float32`

GetLateFeeAmount3 returns the LateFeeAmount3 field if non-nil, zero value otherwise.

### GetLateFeeAmount3Ok

`func (o *ClientSettings) GetLateFeeAmount3Ok() (*float32, bool)`

GetLateFeeAmount3Ok returns a tuple with the LateFeeAmount3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateFeeAmount3

`func (o *ClientSettings) SetLateFeeAmount3(v float32)`

SetLateFeeAmount3 sets LateFeeAmount3 field to given value.

### HasLateFeeAmount3

`func (o *ClientSettings) HasLateFeeAmount3() bool`

HasLateFeeAmount3 returns a boolean if a field has been set.

### GetEndlessReminderFrequencyId

`func (o *ClientSettings) GetEndlessReminderFrequencyId() string`

GetEndlessReminderFrequencyId returns the EndlessReminderFrequencyId field if non-nil, zero value otherwise.

### GetEndlessReminderFrequencyIdOk

`func (o *ClientSettings) GetEndlessReminderFrequencyIdOk() (*string, bool)`

GetEndlessReminderFrequencyIdOk returns a tuple with the EndlessReminderFrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndlessReminderFrequencyId

`func (o *ClientSettings) SetEndlessReminderFrequencyId(v string)`

SetEndlessReminderFrequencyId sets EndlessReminderFrequencyId field to given value.

### HasEndlessReminderFrequencyId

`func (o *ClientSettings) HasEndlessReminderFrequencyId() bool`

HasEndlessReminderFrequencyId returns a boolean if a field has been set.

### GetClientOnlinePaymentNotification

`func (o *ClientSettings) GetClientOnlinePaymentNotification() bool`

GetClientOnlinePaymentNotification returns the ClientOnlinePaymentNotification field if non-nil, zero value otherwise.

### GetClientOnlinePaymentNotificationOk

`func (o *ClientSettings) GetClientOnlinePaymentNotificationOk() (*bool, bool)`

GetClientOnlinePaymentNotificationOk returns a tuple with the ClientOnlinePaymentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOnlinePaymentNotification

`func (o *ClientSettings) SetClientOnlinePaymentNotification(v bool)`

SetClientOnlinePaymentNotification sets ClientOnlinePaymentNotification field to given value.

### HasClientOnlinePaymentNotification

`func (o *ClientSettings) HasClientOnlinePaymentNotification() bool`

HasClientOnlinePaymentNotification returns a boolean if a field has been set.

### GetClientManualPaymentNotification

`func (o *ClientSettings) GetClientManualPaymentNotification() bool`

GetClientManualPaymentNotification returns the ClientManualPaymentNotification field if non-nil, zero value otherwise.

### GetClientManualPaymentNotificationOk

`func (o *ClientSettings) GetClientManualPaymentNotificationOk() (*bool, bool)`

GetClientManualPaymentNotificationOk returns a tuple with the ClientManualPaymentNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientManualPaymentNotification

`func (o *ClientSettings) SetClientManualPaymentNotification(v bool)`

SetClientManualPaymentNotification sets ClientManualPaymentNotification field to given value.

### HasClientManualPaymentNotification

`func (o *ClientSettings) HasClientManualPaymentNotification() bool`

HasClientManualPaymentNotification returns a boolean if a field has been set.

### GetEnableEInvoice

`func (o *ClientSettings) GetEnableEInvoice() bool`

GetEnableEInvoice returns the EnableEInvoice field if non-nil, zero value otherwise.

### GetEnableEInvoiceOk

`func (o *ClientSettings) GetEnableEInvoiceOk() (*bool, bool)`

GetEnableEInvoiceOk returns a tuple with the EnableEInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEInvoice

`func (o *ClientSettings) SetEnableEInvoice(v bool)`

SetEnableEInvoice sets EnableEInvoice field to given value.

### HasEnableEInvoice

`func (o *ClientSettings) HasEnableEInvoice() bool`

HasEnableEInvoice returns a boolean if a field has been set.

### GetDefaultExpensePaymentTypeId

`func (o *ClientSettings) GetDefaultExpensePaymentTypeId() string`

GetDefaultExpensePaymentTypeId returns the DefaultExpensePaymentTypeId field if non-nil, zero value otherwise.

### GetDefaultExpensePaymentTypeIdOk

`func (o *ClientSettings) GetDefaultExpensePaymentTypeIdOk() (*string, bool)`

GetDefaultExpensePaymentTypeIdOk returns a tuple with the DefaultExpensePaymentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExpensePaymentTypeId

`func (o *ClientSettings) SetDefaultExpensePaymentTypeId(v string)`

SetDefaultExpensePaymentTypeId sets DefaultExpensePaymentTypeId field to given value.

### HasDefaultExpensePaymentTypeId

`func (o *ClientSettings) HasDefaultExpensePaymentTypeId() bool`

HasDefaultExpensePaymentTypeId returns a boolean if a field has been set.

### GetEInvoiceType

`func (o *ClientSettings) GetEInvoiceType() string`

GetEInvoiceType returns the EInvoiceType field if non-nil, zero value otherwise.

### GetEInvoiceTypeOk

`func (o *ClientSettings) GetEInvoiceTypeOk() (*string, bool)`

GetEInvoiceTypeOk returns a tuple with the EInvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoiceType

`func (o *ClientSettings) SetEInvoiceType(v string)`

SetEInvoiceType sets EInvoiceType field to given value.

### HasEInvoiceType

`func (o *ClientSettings) HasEInvoiceType() bool`

HasEInvoiceType returns a boolean if a field has been set.

### GetMailgunEndpoint

`func (o *ClientSettings) GetMailgunEndpoint() string`

GetMailgunEndpoint returns the MailgunEndpoint field if non-nil, zero value otherwise.

### GetMailgunEndpointOk

`func (o *ClientSettings) GetMailgunEndpointOk() (*string, bool)`

GetMailgunEndpointOk returns a tuple with the MailgunEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailgunEndpoint

`func (o *ClientSettings) SetMailgunEndpoint(v string)`

SetMailgunEndpoint sets MailgunEndpoint field to given value.

### HasMailgunEndpoint

`func (o *ClientSettings) HasMailgunEndpoint() bool`

HasMailgunEndpoint returns a boolean if a field has been set.

### GetClientInitiatedPayments

`func (o *ClientSettings) GetClientInitiatedPayments() bool`

GetClientInitiatedPayments returns the ClientInitiatedPayments field if non-nil, zero value otherwise.

### GetClientInitiatedPaymentsOk

`func (o *ClientSettings) GetClientInitiatedPaymentsOk() (*bool, bool)`

GetClientInitiatedPaymentsOk returns a tuple with the ClientInitiatedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientInitiatedPayments

`func (o *ClientSettings) SetClientInitiatedPayments(v bool)`

SetClientInitiatedPayments sets ClientInitiatedPayments field to given value.

### HasClientInitiatedPayments

`func (o *ClientSettings) HasClientInitiatedPayments() bool`

HasClientInitiatedPayments returns a boolean if a field has been set.

### GetClientInitiatedPaymentsMinimum

`func (o *ClientSettings) GetClientInitiatedPaymentsMinimum() float32`

GetClientInitiatedPaymentsMinimum returns the ClientInitiatedPaymentsMinimum field if non-nil, zero value otherwise.

### GetClientInitiatedPaymentsMinimumOk

`func (o *ClientSettings) GetClientInitiatedPaymentsMinimumOk() (*float32, bool)`

GetClientInitiatedPaymentsMinimumOk returns a tuple with the ClientInitiatedPaymentsMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientInitiatedPaymentsMinimum

`func (o *ClientSettings) SetClientInitiatedPaymentsMinimum(v float32)`

SetClientInitiatedPaymentsMinimum sets ClientInitiatedPaymentsMinimum field to given value.

### HasClientInitiatedPaymentsMinimum

`func (o *ClientSettings) HasClientInitiatedPaymentsMinimum() bool`

HasClientInitiatedPaymentsMinimum returns a boolean if a field has been set.

### GetSyncInvoiceQuoteColumns

`func (o *ClientSettings) GetSyncInvoiceQuoteColumns() bool`

GetSyncInvoiceQuoteColumns returns the SyncInvoiceQuoteColumns field if non-nil, zero value otherwise.

### GetSyncInvoiceQuoteColumnsOk

`func (o *ClientSettings) GetSyncInvoiceQuoteColumnsOk() (*bool, bool)`

GetSyncInvoiceQuoteColumnsOk returns a tuple with the SyncInvoiceQuoteColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInvoiceQuoteColumns

`func (o *ClientSettings) SetSyncInvoiceQuoteColumns(v bool)`

SetSyncInvoiceQuoteColumns sets SyncInvoiceQuoteColumns field to given value.

### HasSyncInvoiceQuoteColumns

`func (o *ClientSettings) HasSyncInvoiceQuoteColumns() bool`

HasSyncInvoiceQuoteColumns returns a boolean if a field has been set.

### GetShowTaskItemDescription

`func (o *ClientSettings) GetShowTaskItemDescription() bool`

GetShowTaskItemDescription returns the ShowTaskItemDescription field if non-nil, zero value otherwise.

### GetShowTaskItemDescriptionOk

`func (o *ClientSettings) GetShowTaskItemDescriptionOk() (*bool, bool)`

GetShowTaskItemDescriptionOk returns a tuple with the ShowTaskItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTaskItemDescription

`func (o *ClientSettings) SetShowTaskItemDescription(v bool)`

SetShowTaskItemDescription sets ShowTaskItemDescription field to given value.

### HasShowTaskItemDescription

`func (o *ClientSettings) HasShowTaskItemDescription() bool`

HasShowTaskItemDescription returns a boolean if a field has been set.

### GetAllowBillableTaskItems

`func (o *ClientSettings) GetAllowBillableTaskItems() bool`

GetAllowBillableTaskItems returns the AllowBillableTaskItems field if non-nil, zero value otherwise.

### GetAllowBillableTaskItemsOk

`func (o *ClientSettings) GetAllowBillableTaskItemsOk() (*bool, bool)`

GetAllowBillableTaskItemsOk returns a tuple with the AllowBillableTaskItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBillableTaskItems

`func (o *ClientSettings) SetAllowBillableTaskItems(v bool)`

SetAllowBillableTaskItems sets AllowBillableTaskItems field to given value.

### HasAllowBillableTaskItems

`func (o *ClientSettings) HasAllowBillableTaskItems() bool`

HasAllowBillableTaskItems returns a boolean if a field has been set.

### GetAcceptClientInputQuoteApproval

`func (o *ClientSettings) GetAcceptClientInputQuoteApproval() bool`

GetAcceptClientInputQuoteApproval returns the AcceptClientInputQuoteApproval field if non-nil, zero value otherwise.

### GetAcceptClientInputQuoteApprovalOk

`func (o *ClientSettings) GetAcceptClientInputQuoteApprovalOk() (*bool, bool)`

GetAcceptClientInputQuoteApprovalOk returns a tuple with the AcceptClientInputQuoteApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptClientInputQuoteApproval

`func (o *ClientSettings) SetAcceptClientInputQuoteApproval(v bool)`

SetAcceptClientInputQuoteApproval sets AcceptClientInputQuoteApproval field to given value.

### HasAcceptClientInputQuoteApproval

`func (o *ClientSettings) HasAcceptClientInputQuoteApproval() bool`

HasAcceptClientInputQuoteApproval returns a boolean if a field has been set.

### GetCustomSendingEmail

`func (o *ClientSettings) GetCustomSendingEmail() string`

GetCustomSendingEmail returns the CustomSendingEmail field if non-nil, zero value otherwise.

### GetCustomSendingEmailOk

`func (o *ClientSettings) GetCustomSendingEmailOk() (*string, bool)`

GetCustomSendingEmailOk returns a tuple with the CustomSendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSendingEmail

`func (o *ClientSettings) SetCustomSendingEmail(v string)`

SetCustomSendingEmail sets CustomSendingEmail field to given value.

### HasCustomSendingEmail

`func (o *ClientSettings) HasCustomSendingEmail() bool`

HasCustomSendingEmail returns a boolean if a field has been set.

### GetShowPaidStamp

`func (o *ClientSettings) GetShowPaidStamp() bool`

GetShowPaidStamp returns the ShowPaidStamp field if non-nil, zero value otherwise.

### GetShowPaidStampOk

`func (o *ClientSettings) GetShowPaidStampOk() (*bool, bool)`

GetShowPaidStampOk returns a tuple with the ShowPaidStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPaidStamp

`func (o *ClientSettings) SetShowPaidStamp(v bool)`

SetShowPaidStamp sets ShowPaidStamp field to given value.

### HasShowPaidStamp

`func (o *ClientSettings) HasShowPaidStamp() bool`

HasShowPaidStamp returns a boolean if a field has been set.

### GetShowShippingAddress

`func (o *ClientSettings) GetShowShippingAddress() bool`

GetShowShippingAddress returns the ShowShippingAddress field if non-nil, zero value otherwise.

### GetShowShippingAddressOk

`func (o *ClientSettings) GetShowShippingAddressOk() (*bool, bool)`

GetShowShippingAddressOk returns a tuple with the ShowShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowShippingAddress

`func (o *ClientSettings) SetShowShippingAddress(v bool)`

SetShowShippingAddress sets ShowShippingAddress field to given value.

### HasShowShippingAddress

`func (o *ClientSettings) HasShowShippingAddress() bool`

HasShowShippingAddress returns a boolean if a field has been set.

### GetCompanyLogoSize

`func (o *ClientSettings) GetCompanyLogoSize() float32`

GetCompanyLogoSize returns the CompanyLogoSize field if non-nil, zero value otherwise.

### GetCompanyLogoSizeOk

`func (o *ClientSettings) GetCompanyLogoSizeOk() (*float32, bool)`

GetCompanyLogoSizeOk returns a tuple with the CompanyLogoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLogoSize

`func (o *ClientSettings) SetCompanyLogoSize(v float32)`

SetCompanyLogoSize sets CompanyLogoSize field to given value.

### HasCompanyLogoSize

`func (o *ClientSettings) HasCompanyLogoSize() bool`

HasCompanyLogoSize returns a boolean if a field has been set.

### GetShowEmailFooter

`func (o *ClientSettings) GetShowEmailFooter() bool`

GetShowEmailFooter returns the ShowEmailFooter field if non-nil, zero value otherwise.

### GetShowEmailFooterOk

`func (o *ClientSettings) GetShowEmailFooterOk() (*bool, bool)`

GetShowEmailFooterOk returns a tuple with the ShowEmailFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEmailFooter

`func (o *ClientSettings) SetShowEmailFooter(v bool)`

SetShowEmailFooter sets ShowEmailFooter field to given value.

### HasShowEmailFooter

`func (o *ClientSettings) HasShowEmailFooter() bool`

HasShowEmailFooter returns a boolean if a field has been set.

### GetEmailAlignment

`func (o *ClientSettings) GetEmailAlignment() string`

GetEmailAlignment returns the EmailAlignment field if non-nil, zero value otherwise.

### GetEmailAlignmentOk

`func (o *ClientSettings) GetEmailAlignmentOk() (*string, bool)`

GetEmailAlignmentOk returns a tuple with the EmailAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlignment

`func (o *ClientSettings) SetEmailAlignment(v string)`

SetEmailAlignment sets EmailAlignment field to given value.

### HasEmailAlignment

`func (o *ClientSettings) HasEmailAlignment() bool`

HasEmailAlignment returns a boolean if a field has been set.

### GetAutoBillStandardInvoices

`func (o *ClientSettings) GetAutoBillStandardInvoices() bool`

GetAutoBillStandardInvoices returns the AutoBillStandardInvoices field if non-nil, zero value otherwise.

### GetAutoBillStandardInvoicesOk

`func (o *ClientSettings) GetAutoBillStandardInvoicesOk() (*bool, bool)`

GetAutoBillStandardInvoicesOk returns a tuple with the AutoBillStandardInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBillStandardInvoices

`func (o *ClientSettings) SetAutoBillStandardInvoices(v bool)`

SetAutoBillStandardInvoices sets AutoBillStandardInvoices field to given value.

### HasAutoBillStandardInvoices

`func (o *ClientSettings) HasAutoBillStandardInvoices() bool`

HasAutoBillStandardInvoices returns a boolean if a field has been set.

### GetPostmarkSecret

`func (o *ClientSettings) GetPostmarkSecret() string`

GetPostmarkSecret returns the PostmarkSecret field if non-nil, zero value otherwise.

### GetPostmarkSecretOk

`func (o *ClientSettings) GetPostmarkSecretOk() (*string, bool)`

GetPostmarkSecretOk returns a tuple with the PostmarkSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmarkSecret

`func (o *ClientSettings) SetPostmarkSecret(v string)`

SetPostmarkSecret sets PostmarkSecret field to given value.

### HasPostmarkSecret

`func (o *ClientSettings) HasPostmarkSecret() bool`

HasPostmarkSecret returns a boolean if a field has been set.

### GetMailgunSecret

`func (o *ClientSettings) GetMailgunSecret() string`

GetMailgunSecret returns the MailgunSecret field if non-nil, zero value otherwise.

### GetMailgunSecretOk

`func (o *ClientSettings) GetMailgunSecretOk() (*string, bool)`

GetMailgunSecretOk returns a tuple with the MailgunSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailgunSecret

`func (o *ClientSettings) SetMailgunSecret(v string)`

SetMailgunSecret sets MailgunSecret field to given value.

### HasMailgunSecret

`func (o *ClientSettings) HasMailgunSecret() bool`

HasMailgunSecret returns a boolean if a field has been set.

### GetMailgunDomain

`func (o *ClientSettings) GetMailgunDomain() string`

GetMailgunDomain returns the MailgunDomain field if non-nil, zero value otherwise.

### GetMailgunDomainOk

`func (o *ClientSettings) GetMailgunDomainOk() (*string, bool)`

GetMailgunDomainOk returns a tuple with the MailgunDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailgunDomain

`func (o *ClientSettings) SetMailgunDomain(v string)`

SetMailgunDomain sets MailgunDomain field to given value.

### HasMailgunDomain

`func (o *ClientSettings) HasMailgunDomain() bool`

HasMailgunDomain returns a boolean if a field has been set.

### GetSendEmailOnMarkPaid

`func (o *ClientSettings) GetSendEmailOnMarkPaid() bool`

GetSendEmailOnMarkPaid returns the SendEmailOnMarkPaid field if non-nil, zero value otherwise.

### GetSendEmailOnMarkPaidOk

`func (o *ClientSettings) GetSendEmailOnMarkPaidOk() (*bool, bool)`

GetSendEmailOnMarkPaidOk returns a tuple with the SendEmailOnMarkPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmailOnMarkPaid

`func (o *ClientSettings) SetSendEmailOnMarkPaid(v bool)`

SetSendEmailOnMarkPaid sets SendEmailOnMarkPaid field to given value.

### HasSendEmailOnMarkPaid

`func (o *ClientSettings) HasSendEmailOnMarkPaid() bool`

HasSendEmailOnMarkPaid returns a boolean if a field has been set.

### GetVendorPortalEnableUploads

`func (o *ClientSettings) GetVendorPortalEnableUploads() bool`

GetVendorPortalEnableUploads returns the VendorPortalEnableUploads field if non-nil, zero value otherwise.

### GetVendorPortalEnableUploadsOk

`func (o *ClientSettings) GetVendorPortalEnableUploadsOk() (*bool, bool)`

GetVendorPortalEnableUploadsOk returns a tuple with the VendorPortalEnableUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPortalEnableUploads

`func (o *ClientSettings) SetVendorPortalEnableUploads(v bool)`

SetVendorPortalEnableUploads sets VendorPortalEnableUploads field to given value.

### HasVendorPortalEnableUploads

`func (o *ClientSettings) HasVendorPortalEnableUploads() bool`

HasVendorPortalEnableUploads returns a boolean if a field has been set.

### GetBesrId

`func (o *ClientSettings) GetBesrId() string`

GetBesrId returns the BesrId field if non-nil, zero value otherwise.

### GetBesrIdOk

`func (o *ClientSettings) GetBesrIdOk() (*string, bool)`

GetBesrIdOk returns a tuple with the BesrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBesrId

`func (o *ClientSettings) SetBesrId(v string)`

SetBesrId sets BesrId field to given value.

### HasBesrId

`func (o *ClientSettings) HasBesrId() bool`

HasBesrId returns a boolean if a field has been set.

### GetQrIban

`func (o *ClientSettings) GetQrIban() string`

GetQrIban returns the QrIban field if non-nil, zero value otherwise.

### GetQrIbanOk

`func (o *ClientSettings) GetQrIbanOk() (*string, bool)`

GetQrIbanOk returns a tuple with the QrIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrIban

`func (o *ClientSettings) SetQrIban(v string)`

SetQrIban sets QrIban field to given value.

### HasQrIban

`func (o *ClientSettings) HasQrIban() bool`

HasQrIban returns a boolean if a field has been set.

### GetEmailSubjectPurchaseOrder

`func (o *ClientSettings) GetEmailSubjectPurchaseOrder() string`

GetEmailSubjectPurchaseOrder returns the EmailSubjectPurchaseOrder field if non-nil, zero value otherwise.

### GetEmailSubjectPurchaseOrderOk

`func (o *ClientSettings) GetEmailSubjectPurchaseOrderOk() (*string, bool)`

GetEmailSubjectPurchaseOrderOk returns a tuple with the EmailSubjectPurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubjectPurchaseOrder

`func (o *ClientSettings) SetEmailSubjectPurchaseOrder(v string)`

SetEmailSubjectPurchaseOrder sets EmailSubjectPurchaseOrder field to given value.

### HasEmailSubjectPurchaseOrder

`func (o *ClientSettings) HasEmailSubjectPurchaseOrder() bool`

HasEmailSubjectPurchaseOrder returns a boolean if a field has been set.

### GetEmailTemplatePurchaseOrder

`func (o *ClientSettings) GetEmailTemplatePurchaseOrder() string`

GetEmailTemplatePurchaseOrder returns the EmailTemplatePurchaseOrder field if non-nil, zero value otherwise.

### GetEmailTemplatePurchaseOrderOk

`func (o *ClientSettings) GetEmailTemplatePurchaseOrderOk() (*string, bool)`

GetEmailTemplatePurchaseOrderOk returns a tuple with the EmailTemplatePurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplatePurchaseOrder

`func (o *ClientSettings) SetEmailTemplatePurchaseOrder(v string)`

SetEmailTemplatePurchaseOrder sets EmailTemplatePurchaseOrder field to given value.

### HasEmailTemplatePurchaseOrder

`func (o *ClientSettings) HasEmailTemplatePurchaseOrder() bool`

HasEmailTemplatePurchaseOrder returns a boolean if a field has been set.

### GetRequirePurchaseOrderSignature

`func (o *ClientSettings) GetRequirePurchaseOrderSignature() bool`

GetRequirePurchaseOrderSignature returns the RequirePurchaseOrderSignature field if non-nil, zero value otherwise.

### GetRequirePurchaseOrderSignatureOk

`func (o *ClientSettings) GetRequirePurchaseOrderSignatureOk() (*bool, bool)`

GetRequirePurchaseOrderSignatureOk returns a tuple with the RequirePurchaseOrderSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePurchaseOrderSignature

`func (o *ClientSettings) SetRequirePurchaseOrderSignature(v bool)`

SetRequirePurchaseOrderSignature sets RequirePurchaseOrderSignature field to given value.

### HasRequirePurchaseOrderSignature

`func (o *ClientSettings) HasRequirePurchaseOrderSignature() bool`

HasRequirePurchaseOrderSignature returns a boolean if a field has been set.

### GetPurchaseOrderPublicNotes

`func (o *ClientSettings) GetPurchaseOrderPublicNotes() string`

GetPurchaseOrderPublicNotes returns the PurchaseOrderPublicNotes field if non-nil, zero value otherwise.

### GetPurchaseOrderPublicNotesOk

`func (o *ClientSettings) GetPurchaseOrderPublicNotesOk() (*string, bool)`

GetPurchaseOrderPublicNotesOk returns a tuple with the PurchaseOrderPublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderPublicNotes

`func (o *ClientSettings) SetPurchaseOrderPublicNotes(v string)`

SetPurchaseOrderPublicNotes sets PurchaseOrderPublicNotes field to given value.

### HasPurchaseOrderPublicNotes

`func (o *ClientSettings) HasPurchaseOrderPublicNotes() bool`

HasPurchaseOrderPublicNotes returns a boolean if a field has been set.

### GetPurchaseOrderTerms

`func (o *ClientSettings) GetPurchaseOrderTerms() string`

GetPurchaseOrderTerms returns the PurchaseOrderTerms field if non-nil, zero value otherwise.

### GetPurchaseOrderTermsOk

`func (o *ClientSettings) GetPurchaseOrderTermsOk() (*string, bool)`

GetPurchaseOrderTermsOk returns a tuple with the PurchaseOrderTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderTerms

`func (o *ClientSettings) SetPurchaseOrderTerms(v string)`

SetPurchaseOrderTerms sets PurchaseOrderTerms field to given value.

### HasPurchaseOrderTerms

`func (o *ClientSettings) HasPurchaseOrderTerms() bool`

HasPurchaseOrderTerms returns a boolean if a field has been set.

### GetPurchaseOrderFooter

`func (o *ClientSettings) GetPurchaseOrderFooter() string`

GetPurchaseOrderFooter returns the PurchaseOrderFooter field if non-nil, zero value otherwise.

### GetPurchaseOrderFooterOk

`func (o *ClientSettings) GetPurchaseOrderFooterOk() (*string, bool)`

GetPurchaseOrderFooterOk returns a tuple with the PurchaseOrderFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderFooter

`func (o *ClientSettings) SetPurchaseOrderFooter(v string)`

SetPurchaseOrderFooter sets PurchaseOrderFooter field to given value.

### HasPurchaseOrderFooter

`func (o *ClientSettings) HasPurchaseOrderFooter() bool`

HasPurchaseOrderFooter returns a boolean if a field has been set.

### GetPurchaseOrderDesignId

`func (o *ClientSettings) GetPurchaseOrderDesignId() string`

GetPurchaseOrderDesignId returns the PurchaseOrderDesignId field if non-nil, zero value otherwise.

### GetPurchaseOrderDesignIdOk

`func (o *ClientSettings) GetPurchaseOrderDesignIdOk() (*string, bool)`

GetPurchaseOrderDesignIdOk returns a tuple with the PurchaseOrderDesignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderDesignId

`func (o *ClientSettings) SetPurchaseOrderDesignId(v string)`

SetPurchaseOrderDesignId sets PurchaseOrderDesignId field to given value.

### HasPurchaseOrderDesignId

`func (o *ClientSettings) HasPurchaseOrderDesignId() bool`

HasPurchaseOrderDesignId returns a boolean if a field has been set.

### GetPurchaseOrderNumberPattern

`func (o *ClientSettings) GetPurchaseOrderNumberPattern() string`

GetPurchaseOrderNumberPattern returns the PurchaseOrderNumberPattern field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberPatternOk

`func (o *ClientSettings) GetPurchaseOrderNumberPatternOk() (*string, bool)`

GetPurchaseOrderNumberPatternOk returns a tuple with the PurchaseOrderNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumberPattern

`func (o *ClientSettings) SetPurchaseOrderNumberPattern(v string)`

SetPurchaseOrderNumberPattern sets PurchaseOrderNumberPattern field to given value.

### HasPurchaseOrderNumberPattern

`func (o *ClientSettings) HasPurchaseOrderNumberPattern() bool`

HasPurchaseOrderNumberPattern returns a boolean if a field has been set.

### GetPurchaseOrderNumberCounter

`func (o *ClientSettings) GetPurchaseOrderNumberCounter() float32`

GetPurchaseOrderNumberCounter returns the PurchaseOrderNumberCounter field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberCounterOk

`func (o *ClientSettings) GetPurchaseOrderNumberCounterOk() (*float32, bool)`

GetPurchaseOrderNumberCounterOk returns a tuple with the PurchaseOrderNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumberCounter

`func (o *ClientSettings) SetPurchaseOrderNumberCounter(v float32)`

SetPurchaseOrderNumberCounter sets PurchaseOrderNumberCounter field to given value.

### HasPurchaseOrderNumberCounter

`func (o *ClientSettings) HasPurchaseOrderNumberCounter() bool`

HasPurchaseOrderNumberCounter returns a boolean if a field has been set.

### GetPageNumberingAlignment

`func (o *ClientSettings) GetPageNumberingAlignment() string`

GetPageNumberingAlignment returns the PageNumberingAlignment field if non-nil, zero value otherwise.

### GetPageNumberingAlignmentOk

`func (o *ClientSettings) GetPageNumberingAlignmentOk() (*string, bool)`

GetPageNumberingAlignmentOk returns a tuple with the PageNumberingAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumberingAlignment

`func (o *ClientSettings) SetPageNumberingAlignment(v string)`

SetPageNumberingAlignment sets PageNumberingAlignment field to given value.

### HasPageNumberingAlignment

`func (o *ClientSettings) HasPageNumberingAlignment() bool`

HasPageNumberingAlignment returns a boolean if a field has been set.

### GetPageNumbering

`func (o *ClientSettings) GetPageNumbering() bool`

GetPageNumbering returns the PageNumbering field if non-nil, zero value otherwise.

### GetPageNumberingOk

`func (o *ClientSettings) GetPageNumberingOk() (*bool, bool)`

GetPageNumberingOk returns a tuple with the PageNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumbering

`func (o *ClientSettings) SetPageNumbering(v bool)`

SetPageNumbering sets PageNumbering field to given value.

### HasPageNumbering

`func (o *ClientSettings) HasPageNumbering() bool`

HasPageNumbering returns a boolean if a field has been set.

### GetAutoArchiveInvoiceCancelled

`func (o *ClientSettings) GetAutoArchiveInvoiceCancelled() bool`

GetAutoArchiveInvoiceCancelled returns the AutoArchiveInvoiceCancelled field if non-nil, zero value otherwise.

### GetAutoArchiveInvoiceCancelledOk

`func (o *ClientSettings) GetAutoArchiveInvoiceCancelledOk() (*bool, bool)`

GetAutoArchiveInvoiceCancelledOk returns a tuple with the AutoArchiveInvoiceCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveInvoiceCancelled

`func (o *ClientSettings) SetAutoArchiveInvoiceCancelled(v bool)`

SetAutoArchiveInvoiceCancelled sets AutoArchiveInvoiceCancelled field to given value.

### HasAutoArchiveInvoiceCancelled

`func (o *ClientSettings) HasAutoArchiveInvoiceCancelled() bool`

HasAutoArchiveInvoiceCancelled returns a boolean if a field has been set.

### GetEmailFromName

`func (o *ClientSettings) GetEmailFromName() string`

GetEmailFromName returns the EmailFromName field if non-nil, zero value otherwise.

### GetEmailFromNameOk

`func (o *ClientSettings) GetEmailFromNameOk() (*string, bool)`

GetEmailFromNameOk returns a tuple with the EmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromName

`func (o *ClientSettings) SetEmailFromName(v string)`

SetEmailFromName sets EmailFromName field to given value.

### HasEmailFromName

`func (o *ClientSettings) HasEmailFromName() bool`

HasEmailFromName returns a boolean if a field has been set.

### GetShowAllTasksClientPortal

`func (o *ClientSettings) GetShowAllTasksClientPortal() bool`

GetShowAllTasksClientPortal returns the ShowAllTasksClientPortal field if non-nil, zero value otherwise.

### GetShowAllTasksClientPortalOk

`func (o *ClientSettings) GetShowAllTasksClientPortalOk() (*bool, bool)`

GetShowAllTasksClientPortalOk returns a tuple with the ShowAllTasksClientPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllTasksClientPortal

`func (o *ClientSettings) SetShowAllTasksClientPortal(v bool)`

SetShowAllTasksClientPortal sets ShowAllTasksClientPortal field to given value.

### HasShowAllTasksClientPortal

`func (o *ClientSettings) HasShowAllTasksClientPortal() bool`

HasShowAllTasksClientPortal returns a boolean if a field has been set.

### GetEntitySendTime

`func (o *ClientSettings) GetEntitySendTime() int32`

GetEntitySendTime returns the EntitySendTime field if non-nil, zero value otherwise.

### GetEntitySendTimeOk

`func (o *ClientSettings) GetEntitySendTimeOk() (*int32, bool)`

GetEntitySendTimeOk returns a tuple with the EntitySendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySendTime

`func (o *ClientSettings) SetEntitySendTime(v int32)`

SetEntitySendTime sets EntitySendTime field to given value.

### HasEntitySendTime

`func (o *ClientSettings) HasEntitySendTime() bool`

HasEntitySendTime returns a boolean if a field has been set.

### GetSharedInvoiceCreditCounter

`func (o *ClientSettings) GetSharedInvoiceCreditCounter() bool`

GetSharedInvoiceCreditCounter returns the SharedInvoiceCreditCounter field if non-nil, zero value otherwise.

### GetSharedInvoiceCreditCounterOk

`func (o *ClientSettings) GetSharedInvoiceCreditCounterOk() (*bool, bool)`

GetSharedInvoiceCreditCounterOk returns a tuple with the SharedInvoiceCreditCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedInvoiceCreditCounter

`func (o *ClientSettings) SetSharedInvoiceCreditCounter(v bool)`

SetSharedInvoiceCreditCounter sets SharedInvoiceCreditCounter field to given value.

### HasSharedInvoiceCreditCounter

`func (o *ClientSettings) HasSharedInvoiceCreditCounter() bool`

HasSharedInvoiceCreditCounter returns a boolean if a field has been set.

### GetReplyToName

`func (o *ClientSettings) GetReplyToName() string`

GetReplyToName returns the ReplyToName field if non-nil, zero value otherwise.

### GetReplyToNameOk

`func (o *ClientSettings) GetReplyToNameOk() (*string, bool)`

GetReplyToNameOk returns a tuple with the ReplyToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToName

`func (o *ClientSettings) SetReplyToName(v string)`

SetReplyToName sets ReplyToName field to given value.

### HasReplyToName

`func (o *ClientSettings) HasReplyToName() bool`

HasReplyToName returns a boolean if a field has been set.

### GetHideEmptyColumnsOnPdf

`func (o *ClientSettings) GetHideEmptyColumnsOnPdf() bool`

GetHideEmptyColumnsOnPdf returns the HideEmptyColumnsOnPdf field if non-nil, zero value otherwise.

### GetHideEmptyColumnsOnPdfOk

`func (o *ClientSettings) GetHideEmptyColumnsOnPdfOk() (*bool, bool)`

GetHideEmptyColumnsOnPdfOk returns a tuple with the HideEmptyColumnsOnPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideEmptyColumnsOnPdf

`func (o *ClientSettings) SetHideEmptyColumnsOnPdf(v bool)`

SetHideEmptyColumnsOnPdf sets HideEmptyColumnsOnPdf field to given value.

### HasHideEmptyColumnsOnPdf

`func (o *ClientSettings) HasHideEmptyColumnsOnPdf() bool`

HasHideEmptyColumnsOnPdf returns a boolean if a field has been set.

### GetEnableReminderEndless

`func (o *ClientSettings) GetEnableReminderEndless() bool`

GetEnableReminderEndless returns the EnableReminderEndless field if non-nil, zero value otherwise.

### GetEnableReminderEndlessOk

`func (o *ClientSettings) GetEnableReminderEndlessOk() (*bool, bool)`

GetEnableReminderEndlessOk returns a tuple with the EnableReminderEndless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReminderEndless

`func (o *ClientSettings) SetEnableReminderEndless(v bool)`

SetEnableReminderEndless sets EnableReminderEndless field to given value.

### HasEnableReminderEndless

`func (o *ClientSettings) HasEnableReminderEndless() bool`

HasEnableReminderEndless returns a boolean if a field has been set.

### GetUseCreditsPayment

`func (o *ClientSettings) GetUseCreditsPayment() bool`

GetUseCreditsPayment returns the UseCreditsPayment field if non-nil, zero value otherwise.

### GetUseCreditsPaymentOk

`func (o *ClientSettings) GetUseCreditsPaymentOk() (*bool, bool)`

GetUseCreditsPaymentOk returns a tuple with the UseCreditsPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCreditsPayment

`func (o *ClientSettings) SetUseCreditsPayment(v bool)`

SetUseCreditsPayment sets UseCreditsPayment field to given value.

### HasUseCreditsPayment

`func (o *ClientSettings) HasUseCreditsPayment() bool`

HasUseCreditsPayment returns a boolean if a field has been set.

### GetRecurringInvoiceNumberPattern

`func (o *ClientSettings) GetRecurringInvoiceNumberPattern() string`

GetRecurringInvoiceNumberPattern returns the RecurringInvoiceNumberPattern field if non-nil, zero value otherwise.

### GetRecurringInvoiceNumberPatternOk

`func (o *ClientSettings) GetRecurringInvoiceNumberPatternOk() (*string, bool)`

GetRecurringInvoiceNumberPatternOk returns a tuple with the RecurringInvoiceNumberPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInvoiceNumberPattern

`func (o *ClientSettings) SetRecurringInvoiceNumberPattern(v string)`

SetRecurringInvoiceNumberPattern sets RecurringInvoiceNumberPattern field to given value.

### HasRecurringInvoiceNumberPattern

`func (o *ClientSettings) HasRecurringInvoiceNumberPattern() bool`

HasRecurringInvoiceNumberPattern returns a boolean if a field has been set.

### GetRecurringInvoiceNumberCounter

`func (o *ClientSettings) GetRecurringInvoiceNumberCounter() float32`

GetRecurringInvoiceNumberCounter returns the RecurringInvoiceNumberCounter field if non-nil, zero value otherwise.

### GetRecurringInvoiceNumberCounterOk

`func (o *ClientSettings) GetRecurringInvoiceNumberCounterOk() (*float32, bool)`

GetRecurringInvoiceNumberCounterOk returns a tuple with the RecurringInvoiceNumberCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInvoiceNumberCounter

`func (o *ClientSettings) SetRecurringInvoiceNumberCounter(v float32)`

SetRecurringInvoiceNumberCounter sets RecurringInvoiceNumberCounter field to given value.

### HasRecurringInvoiceNumberCounter

`func (o *ClientSettings) HasRecurringInvoiceNumberCounter() bool`

HasRecurringInvoiceNumberCounter returns a boolean if a field has been set.

### GetClientPortalUnderPaymentMinimum

`func (o *ClientSettings) GetClientPortalUnderPaymentMinimum() float32`

GetClientPortalUnderPaymentMinimum returns the ClientPortalUnderPaymentMinimum field if non-nil, zero value otherwise.

### GetClientPortalUnderPaymentMinimumOk

`func (o *ClientSettings) GetClientPortalUnderPaymentMinimumOk() (*float32, bool)`

GetClientPortalUnderPaymentMinimumOk returns a tuple with the ClientPortalUnderPaymentMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPortalUnderPaymentMinimum

`func (o *ClientSettings) SetClientPortalUnderPaymentMinimum(v float32)`

SetClientPortalUnderPaymentMinimum sets ClientPortalUnderPaymentMinimum field to given value.

### HasClientPortalUnderPaymentMinimum

`func (o *ClientSettings) HasClientPortalUnderPaymentMinimum() bool`

HasClientPortalUnderPaymentMinimum returns a boolean if a field has been set.

### GetAutoBillDate

`func (o *ClientSettings) GetAutoBillDate() string`

GetAutoBillDate returns the AutoBillDate field if non-nil, zero value otherwise.

### GetAutoBillDateOk

`func (o *ClientSettings) GetAutoBillDateOk() (*string, bool)`

GetAutoBillDateOk returns a tuple with the AutoBillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBillDate

`func (o *ClientSettings) SetAutoBillDate(v string)`

SetAutoBillDate sets AutoBillDate field to given value.

### HasAutoBillDate

`func (o *ClientSettings) HasAutoBillDate() bool`

HasAutoBillDate returns a boolean if a field has been set.

### GetPrimaryColor

`func (o *ClientSettings) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *ClientSettings) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *ClientSettings) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.

### HasPrimaryColor

`func (o *ClientSettings) HasPrimaryColor() bool`

HasPrimaryColor returns a boolean if a field has been set.

### GetSecondaryColor

`func (o *ClientSettings) GetSecondaryColor() string`

GetSecondaryColor returns the SecondaryColor field if non-nil, zero value otherwise.

### GetSecondaryColorOk

`func (o *ClientSettings) GetSecondaryColorOk() (*string, bool)`

GetSecondaryColorOk returns a tuple with the SecondaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColor

`func (o *ClientSettings) SetSecondaryColor(v string)`

SetSecondaryColor sets SecondaryColor field to given value.

### HasSecondaryColor

`func (o *ClientSettings) HasSecondaryColor() bool`

HasSecondaryColor returns a boolean if a field has been set.

### GetClientPortalAllowUnderPayment

`func (o *ClientSettings) GetClientPortalAllowUnderPayment() bool`

GetClientPortalAllowUnderPayment returns the ClientPortalAllowUnderPayment field if non-nil, zero value otherwise.

### GetClientPortalAllowUnderPaymentOk

`func (o *ClientSettings) GetClientPortalAllowUnderPaymentOk() (*bool, bool)`

GetClientPortalAllowUnderPaymentOk returns a tuple with the ClientPortalAllowUnderPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPortalAllowUnderPayment

`func (o *ClientSettings) SetClientPortalAllowUnderPayment(v bool)`

SetClientPortalAllowUnderPayment sets ClientPortalAllowUnderPayment field to given value.

### HasClientPortalAllowUnderPayment

`func (o *ClientSettings) HasClientPortalAllowUnderPayment() bool`

HasClientPortalAllowUnderPayment returns a boolean if a field has been set.

### GetClientPortalAllowOverPayment

`func (o *ClientSettings) GetClientPortalAllowOverPayment() bool`

GetClientPortalAllowOverPayment returns the ClientPortalAllowOverPayment field if non-nil, zero value otherwise.

### GetClientPortalAllowOverPaymentOk

`func (o *ClientSettings) GetClientPortalAllowOverPaymentOk() (*bool, bool)`

GetClientPortalAllowOverPaymentOk returns a tuple with the ClientPortalAllowOverPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPortalAllowOverPayment

`func (o *ClientSettings) SetClientPortalAllowOverPayment(v bool)`

SetClientPortalAllowOverPayment sets ClientPortalAllowOverPayment field to given value.

### HasClientPortalAllowOverPayment

`func (o *ClientSettings) HasClientPortalAllowOverPayment() bool`

HasClientPortalAllowOverPayment returns a boolean if a field has been set.

### GetAutoBill

`func (o *ClientSettings) GetAutoBill() string`

GetAutoBill returns the AutoBill field if non-nil, zero value otherwise.

### GetAutoBillOk

`func (o *ClientSettings) GetAutoBillOk() (*string, bool)`

GetAutoBillOk returns a tuple with the AutoBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBill

`func (o *ClientSettings) SetAutoBill(v string)`

SetAutoBill sets AutoBill field to given value.

### HasAutoBill

`func (o *ClientSettings) HasAutoBill() bool`

HasAutoBill returns a boolean if a field has been set.

### GetClientPortalTerms

`func (o *ClientSettings) GetClientPortalTerms() string`

GetClientPortalTerms returns the ClientPortalTerms field if non-nil, zero value otherwise.

### GetClientPortalTermsOk

`func (o *ClientSettings) GetClientPortalTermsOk() (*string, bool)`

GetClientPortalTermsOk returns a tuple with the ClientPortalTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPortalTerms

`func (o *ClientSettings) SetClientPortalTerms(v string)`

SetClientPortalTerms sets ClientPortalTerms field to given value.

### HasClientPortalTerms

`func (o *ClientSettings) HasClientPortalTerms() bool`

HasClientPortalTerms returns a boolean if a field has been set.

### GetClientPortalPrivacyPolicy

`func (o *ClientSettings) GetClientPortalPrivacyPolicy() string`

GetClientPortalPrivacyPolicy returns the ClientPortalPrivacyPolicy field if non-nil, zero value otherwise.

### GetClientPortalPrivacyPolicyOk

`func (o *ClientSettings) GetClientPortalPrivacyPolicyOk() (*string, bool)`

GetClientPortalPrivacyPolicyOk returns a tuple with the ClientPortalPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPortalPrivacyPolicy

`func (o *ClientSettings) SetClientPortalPrivacyPolicy(v string)`

SetClientPortalPrivacyPolicy sets ClientPortalPrivacyPolicy field to given value.

### HasClientPortalPrivacyPolicy

`func (o *ClientSettings) HasClientPortalPrivacyPolicy() bool`

HasClientPortalPrivacyPolicy returns a boolean if a field has been set.

### GetClientCanRegister

`func (o *ClientSettings) GetClientCanRegister() bool`

GetClientCanRegister returns the ClientCanRegister field if non-nil, zero value otherwise.

### GetClientCanRegisterOk

`func (o *ClientSettings) GetClientCanRegisterOk() (*bool, bool)`

GetClientCanRegisterOk returns a tuple with the ClientCanRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCanRegister

`func (o *ClientSettings) SetClientCanRegister(v bool)`

SetClientCanRegister sets ClientCanRegister field to given value.

### HasClientCanRegister

`func (o *ClientSettings) HasClientCanRegister() bool`

HasClientCanRegister returns a boolean if a field has been set.

### GetPortalDesignId

`func (o *ClientSettings) GetPortalDesignId() string`

GetPortalDesignId returns the PortalDesignId field if non-nil, zero value otherwise.

### GetPortalDesignIdOk

`func (o *ClientSettings) GetPortalDesignIdOk() (*string, bool)`

GetPortalDesignIdOk returns a tuple with the PortalDesignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalDesignId

`func (o *ClientSettings) SetPortalDesignId(v string)`

SetPortalDesignId sets PortalDesignId field to given value.

### HasPortalDesignId

`func (o *ClientSettings) HasPortalDesignId() bool`

HasPortalDesignId returns a boolean if a field has been set.

### GetLateFeeEndlessPercent

`func (o *ClientSettings) GetLateFeeEndlessPercent() float32`

GetLateFeeEndlessPercent returns the LateFeeEndlessPercent field if non-nil, zero value otherwise.

### GetLateFeeEndlessPercentOk

`func (o *ClientSettings) GetLateFeeEndlessPercentOk() (*float32, bool)`

GetLateFeeEndlessPercentOk returns a tuple with the LateFeeEndlessPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateFeeEndlessPercent

`func (o *ClientSettings) SetLateFeeEndlessPercent(v float32)`

SetLateFeeEndlessPercent sets LateFeeEndlessPercent field to given value.

### HasLateFeeEndlessPercent

`func (o *ClientSettings) HasLateFeeEndlessPercent() bool`

HasLateFeeEndlessPercent returns a boolean if a field has been set.

### GetLateFeeEndlessAmount

`func (o *ClientSettings) GetLateFeeEndlessAmount() float32`

GetLateFeeEndlessAmount returns the LateFeeEndlessAmount field if non-nil, zero value otherwise.

### GetLateFeeEndlessAmountOk

`func (o *ClientSettings) GetLateFeeEndlessAmountOk() (*float32, bool)`

GetLateFeeEndlessAmountOk returns a tuple with the LateFeeEndlessAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateFeeEndlessAmount

`func (o *ClientSettings) SetLateFeeEndlessAmount(v float32)`

SetLateFeeEndlessAmount sets LateFeeEndlessAmount field to given value.

### HasLateFeeEndlessAmount

`func (o *ClientSettings) HasLateFeeEndlessAmount() bool`

HasLateFeeEndlessAmount returns a boolean if a field has been set.

### GetAutoEmailInvoice

`func (o *ClientSettings) GetAutoEmailInvoice() bool`

GetAutoEmailInvoice returns the AutoEmailInvoice field if non-nil, zero value otherwise.

### GetAutoEmailInvoiceOk

`func (o *ClientSettings) GetAutoEmailInvoiceOk() (*bool, bool)`

GetAutoEmailInvoiceOk returns a tuple with the AutoEmailInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEmailInvoice

`func (o *ClientSettings) SetAutoEmailInvoice(v bool)`

SetAutoEmailInvoice sets AutoEmailInvoice field to given value.

### HasAutoEmailInvoice

`func (o *ClientSettings) HasAutoEmailInvoice() bool`

HasAutoEmailInvoice returns a boolean if a field has been set.

### GetEmailSignature

`func (o *ClientSettings) GetEmailSignature() string`

GetEmailSignature returns the EmailSignature field if non-nil, zero value otherwise.

### GetEmailSignatureOk

`func (o *ClientSettings) GetEmailSignatureOk() (*string, bool)`

GetEmailSignatureOk returns a tuple with the EmailSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSignature

`func (o *ClientSettings) SetEmailSignature(v string)`

SetEmailSignature sets EmailSignature field to given value.

### HasEmailSignature

`func (o *ClientSettings) HasEmailSignature() bool`

HasEmailSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


