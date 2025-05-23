package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CompanySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanySettings{}

// CompanySettings struct for CompanySettings
type CompanySettings struct {
	// The timezone id
	TimezoneId *string `json:"timezone_id,omitempty"`
	// The date format id
	DateFormatId *string `json:"date_format_id,omitempty"`
	// Toggles 12/24 hour time
	MilitaryTime *bool `json:"military_time,omitempty"`
	// The language id
	LanguageId *string `json:"language_id,omitempty"`
	// Toggles whether the currency symbol or code is shown
	ShowCurrencyCode *bool `json:"show_currency_code,omitempty"`
	// -1 sets no payment term, 0 sets payment due immediately, positive integers indicates payment terms in days
	PaymentTerms *int32 `json:"payment_terms,omitempty"`
	// A commad separate list of available gateways
	CompanyGatewayIds *string `json:"company_gateway_ids,omitempty"`
	// A Custom Label
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A Custom Label
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A Custom Label
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A Custom Label
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The default task rate
	DefaultTaskRate *float32 `json:"default_task_rate,omitempty"`
	// Toggles whether reminders are sent
	SendReminders *bool `json:"send_reminders,omitempty"`
	// Show/hide the tasks panel in the client portal
	EnableClientPortalTasks *bool `json:"enable_client_portal_tasks,omitempty"`
	// options include plain,light,dark,custom
	EmailStyle *string `json:"email_style,omitempty"`
	// The reply to email address
	ReplyToEmail *string `json:"reply_to_email,omitempty"`
	// A comma separate list of BCC emails
	BccEmail *string `json:"bcc_email,omitempty"`
	// Toggles whether to attach PDF as attachment
	PdfEmailAttachment *bool `json:"pdf_email_attachment,omitempty"`
	// Toggles whether to attach UBL as attachment
	UblEmailAttachment *bool `json:"ubl_email_attachment,omitempty"`
	// The custom template
	EmailStyleCustom *string `json:"email_style_custom,omitempty"`
	// enum when the invoice number counter is set
	CounterNumberApplied *string `json:"counter_number_applied,omitempty"`
	// enum when the quote number counter is set
	QuoteNumberApplied *string `json:"quote_number_applied,omitempty"`
	// A custom message which is displayed on the dashboard
	CustomMessageDashboard *string `json:"custom_message_dashboard,omitempty"`
	// A custom message which is displayed in the client portal when a client is viewing a unpaid invoice.
	CustomMessageUnpaidInvoice *string `json:"custom_message_unpaid_invoice,omitempty"`
	// A custom message which is displayed in the client portal when a client is viewing a paid invoice.
	CustomMessagePaidInvoice *string `json:"custom_message_paid_invoice,omitempty"`
	// A custom message which is displayed in the client portal when a client is viewing a unapproved quote.
	CustomMessageUnapprovedQuote *string `json:"custom_message_unapproved_quote,omitempty"`
	// Toggles whether invoices are locked once sent and cannot be modified further
	LockInvoices *string `json:"lock_invoices,omitempty"`
	// Toggles whether a invoice is archived immediately following payment
	AutoArchiveInvoice *bool `json:"auto_archive_invoice,omitempty"`
	// Toggles whether a quote is archived after being converted to a invoice
	AutoArchiveQuote *bool `json:"auto_archive_quote,omitempty"`
	// Toggles whether a quote is converted to a invoice when approved
	AutoConvertQuote *bool `json:"auto_convert_quote,omitempty"`
	// Boolean flag determining whether inclusive or exclusive taxes are used
	InclusiveTaxes *bool `json:"inclusive_taxes,omitempty"`
	// JSON payload of customized translations
	Translations map[string]any `json:"translations,omitempty"`
	// Allows customisation of the task number pattern
	TaskNumberPattern *string `json:"task_number_pattern,omitempty"`
	// The incrementing counter for tasks
	TaskNumberCounter *int32 `json:"task_number_counter,omitempty"`
	// Time from UTC +0 when the email will be sent to the client
	ReminderSendTime *int32 `json:"reminder_send_time,omitempty"`
	// Allows customisation of the expense number pattern
	ExpenseNumberPattern *string `json:"expense_number_pattern,omitempty"`
	// The incrementing counter for expenses
	ExpenseNumberCounter *int32 `json:"expense_number_counter,omitempty"`
	// Allows customisation of the vendor number pattern
	VendorNumberPattern *string `json:"vendor_number_pattern,omitempty"`
	// The incrementing counter for vendors
	VendorNumberCounter *int32 `json:"vendor_number_counter,omitempty"`
	// Allows customisation of the ticket number pattern
	TicketNumberPattern *string `json:"ticket_number_pattern,omitempty"`
	// The incrementing counter for tickets
	TicketNumberCounter *int32 `json:"ticket_number_counter,omitempty"`
	// Allows customisation of the payment number pattern
	PaymentNumberPattern *string `json:"payment_number_pattern,omitempty"`
	// The incrementing counter for payments
	PaymentNumberCounter *int32 `json:"payment_number_counter,omitempty"`
	// Allows customisation of the invoice number pattern
	InvoiceNumberPattern *string `json:"invoice_number_pattern,omitempty"`
	// The incrementing counter for invoices
	InvoiceNumberCounter *int32 `json:"invoice_number_counter,omitempty"`
	// Allows customisation of the quote number pattern
	QuoteNumberPattern *string `json:"quote_number_pattern,omitempty"`
	// The incrementing counter for quotes
	QuoteNumberCounter *int32 `json:"quote_number_counter,omitempty"`
	// Allows customisation of the client number pattern
	ClientNumberPattern *string `json:"client_number_pattern,omitempty"`
	// The incrementing counter for clients
	ClientNumberCounter *int32 `json:"client_number_counter,omitempty"`
	// Allows customisation of the credit number pattern
	CreditNumberPattern *string `json:"credit_number_pattern,omitempty"`
	// The incrementing counter for credits
	CreditNumberCounter *int32 `json:"credit_number_counter,omitempty"`
	// This string is prepended to the recurring invoice number
	RecurringInvoiceNumberPrefix *string `json:"recurring_invoice_number_prefix,omitempty"`
	// CONSTANT which is used to apply the frequency which the counters are reset
	ResetCounterFrequencyId *int32 `json:"reset_counter_frequency_id,omitempty"`
	// The explicit date which is used to reset counters
	ResetCounterDate *string `json:"reset_counter_date,omitempty"`
	// Pads the counter with leading zeros
	CounterPadding *int32 `json:"counter_padding,omitempty"`
	// Flags whether to share the counter for invoices and quotes
	SharedInvoiceQuoteCounter *bool `json:"shared_invoice_quote_counter,omitempty"`
	// Determines if client fields are updated from third party APIs
	UpdateProducts *bool `json:"update_products,omitempty"`
	//
	ConvertProducts *bool `json:"convert_products,omitempty"`
	// Automatically fill products based on product_key
	FillProducts *bool `json:"fill_products,omitempty"`
	// The default invoice terms
	InvoiceTerms *string `json:"invoice_terms,omitempty"`
	// The default quote terms
	QuoteTerms *string `json:"quote_terms,omitempty"`
	// The number of Taxes that can be applied to the invoice
	InvoiceTaxes *int32 `json:"invoice_taxes,omitempty"`
	// The default design id (invoice, quote etc)
	InvoiceDesignId *string `json:"invoice_design_id,omitempty"`
	// The default design id (invoice, quote etc)
	QuoteDesignId *string `json:"quote_design_id,omitempty"`
	// The default invoice footer
	InvoiceFooter *string `json:"invoice_footer,omitempty"`
	// JSON string of invoice labels
	InvoiceLabels *string `json:"invoice_labels,omitempty"`
	// The tax rate (float)
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// The tax name
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The tax rate (float)
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// The tax name
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The tax rate (float)
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// The tax name
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The default payment type id
	PaymentTypeId *string `json:"payment_type_id,omitempty"`
	// JSON string of custom fields
	CustomFields *string `json:"custom_fields,omitempty"`
	// The default email footer
	EmailFooter *string `json:"email_footer,omitempty"`
	// The email driver to use to send email
	EmailSendingMethod *string `json:"email_sending_method,omitempty"`
	// The hashed_id of the user account to send email from
	GmailSendingUserId *string `json:"gmail_sending_user_id,omitempty"`
	//
	EmailSubjectInvoice *string `json:"email_subject_invoice,omitempty"`
	//
	EmailSubjectQuote *string `json:"email_subject_quote,omitempty"`
	//
	EmailSubjectPayment *string `json:"email_subject_payment,omitempty"`
	// The full template for invoice emails
	EmailTemplateInvoice *string `json:"email_template_invoice,omitempty"`
	// The full template for quote emails
	EmailTemplateQuote *string `json:"email_template_quote,omitempty"`
	// The full template for payment emails
	EmailTemplatePayment *string `json:"email_template_payment,omitempty"`
	// Email subject for Reminder
	EmailSubjectReminder1 *string `json:"email_subject_reminder1,omitempty"`
	// Email subject for Reminder
	EmailSubjectReminder2 *string `json:"email_subject_reminder2,omitempty"`
	// Email subject for Reminder
	EmailSubjectReminder3 *string `json:"email_subject_reminder3,omitempty"`
	// Email subject for endless reminders
	EmailSubjectReminderEndless *string `json:"email_subject_reminder_endless,omitempty"`
	// The full template for Reminder 1
	EmailTemplateReminder1 *string `json:"email_template_reminder1,omitempty"`
	// The full template for Reminder 2
	EmailTemplateReminder2 *string `json:"email_template_reminder2,omitempty"`
	// The full template for Reminder 3
	EmailTemplateReminder3 *string `json:"email_template_reminder3,omitempty"`
	// The full template for enless reminders
	EmailTemplateReminderEndless *string `json:"email_template_reminder_endless,omitempty"`
	// Toggles whether a password is required to log into the client portal
	EnablePortalPassword *bool `json:"enable_portal_password,omitempty"`
	// Toggles whether the terms dialogue is shown to the client
	ShowAcceptInvoiceTerms *bool `json:"show_accept_invoice_terms,omitempty"`
	// Toggles whether the terms dialogue is shown to the client
	ShowAcceptQuoteTerms *bool `json:"show_accept_quote_terms,omitempty"`
	// Toggles whether a invoice signature is required
	RequireInvoiceSignature *bool `json:"require_invoice_signature,omitempty"`
	// Toggles whether a quote signature is required
	RequireQuoteSignature *bool `json:"require_quote_signature,omitempty"`
	// The company name
	Name *string `json:"name,omitempty"`
	// The company logo file
	CompanyLogo *string `json:"company_logo,omitempty"`
	// The company website URL
	Website *string `json:"website,omitempty"`
	// The company address line 1
	Address1 *string `json:"address1,omitempty"`
	// The company address line 2
	Address2 *string `json:"address2,omitempty"`
	// The company city
	City *string `json:"city,omitempty"`
	// The company state
	State *string `json:"state,omitempty"`
	// The company zip/postal code
	PostalCode *string `json:"postal_code,omitempty"`
	// The company phone
	Phone *string `json:"phone,omitempty"`
	// The company email
	Email *string `json:"email,omitempty"`
	// The country ID
	CountryId *string `json:"country_id,omitempty"`
	// The company VAT/TAX ID number
	VatNumber *string `json:"vat_number,omitempty"`
	// The default page size
	PageSize *string `json:"page_size,omitempty"`
	// The font size
	FontSize *float32 `json:"font_size,omitempty"`
	// The primary font
	PrimaryFont *string `json:"primary_font,omitempty"`
	// The secondary font
	SecondaryFont *string `json:"secondary_font,omitempty"`
	// Flags whether to hide the paid to date field
	HidePaidToDate *bool `json:"hide_paid_to_date,omitempty"`
	// Toggled whether to embed documents in the PDF
	EmbedDocuments *bool `json:"embed_documents,omitempty"`
	// The header for the PDF
	AllPagesHeader *bool `json:"all_pages_header,omitempty"`
	// The footer for the PDF
	AllPagesFooter *bool `json:"all_pages_footer,omitempty"`
	// Toggles whether to attach documents in the email
	DocumentEmailAttachment *bool `json:"document_email_attachment,omitempty"`
	// Toggles password protection of the client portal
	EnableClientPortalPassword *bool `json:"enable_client_portal_password,omitempty"`
	// Toggles the use of markdown in emails
	EnableEmailMarkup *bool `json:"enable_email_markup,omitempty"`
	// Toggles whether the client dashboard is shown in the client portal
	EnableClientPortalDashboard *bool `json:"enable_client_portal_dashboard,omitempty"`
	// Toggles whether the entire client portal is displayed to the client, or only the context
	EnableClientPortal *bool `json:"enable_client_portal,omitempty"`
	// The body of the email for statements
	EmailTemplateStatement *string `json:"email_template_statement,omitempty"`
	// The subject of the email for statements
	EmailSubjectStatement *string `json:"email_subject_statement,omitempty"`
	// Toggles whether the signature (if available) is displayed on the PDF
	SignatureOnPdf *bool `json:"signature_on_pdf,omitempty"`
	// The default quote footer
	QuoteFooter *string `json:"quote_footer,omitempty"`
	// Custom reminder template subject
	EmailSubjectCustom1 *string `json:"email_subject_custom1,omitempty"`
	// Custom reminder template subject
	EmailSubjectCustom2 *string `json:"email_subject_custom2,omitempty"`
	// Custom reminder template subject
	EmailSubjectCustom3 *string `json:"email_subject_custom3,omitempty"`
	// Custom reminder template body
	EmailTemplateCustom1 *string `json:"email_template_custom1,omitempty"`
	// Custom reminder template body
	EmailTemplateCustom2 *string `json:"email_template_custom2,omitempty"`
	// Custom reminder template body
	EmailTemplateCustom3 *string `json:"email_template_custom3,omitempty"`
	// Toggles whether this reminder is enabled
	EnableReminder1 *bool `json:"enable_reminder1,omitempty"`
	// Toggles whether this reminder is enabled
	EnableReminder2 *bool `json:"enable_reminder2,omitempty"`
	// Toggles whether this reminder is enabled
	EnableReminder3 *bool `json:"enable_reminder3,omitempty"`
	// The Reminder interval
	NumDaysReminder1 *int32 `json:"num_days_reminder1,omitempty"`
	// The Reminder interval
	NumDaysReminder2 *int32 `json:"num_days_reminder2,omitempty"`
	// The Reminder interval
	NumDaysReminder3 *int32 `json:"num_days_reminder3,omitempty"`
	// (enum: after_invoice_date, before_due_date, after_due_date)
	ScheduleReminder1 *string `json:"schedule_reminder1,omitempty"`
	// (enum: after_invoice_date, before_due_date, after_due_date)
	ScheduleReminder2 *string `json:"schedule_reminder2,omitempty"`
	// (enum: after_invoice_date, before_due_date, after_due_date)
	ScheduleReminder3 *string `json:"schedule_reminder3,omitempty"`
	// The late fee amount for reminder 1
	LateFeeAmount1 *float32 `json:"late_fee_amount1,omitempty"`
	// The late fee amount for reminder 2
	LateFeeAmount2 *float32 `json:"late_fee_amount2,omitempty"`
	// The late fee amount for reminder 2
	LateFeeAmount3 *float32 `json:"late_fee_amount3,omitempty"`
	// The frequency id of the endless reminder
	EndlessReminderFrequencyId *string `json:"endless_reminder_frequency_id,omitempty"`
	// Determines if a client should receive the notification for a online payment
	ClientOnlinePaymentNotification *bool `json:"client_online_payment_notification,omitempty"`
	// Determines if a client should receive the notification for a manually entered payment
	ClientManualPaymentNotification *bool `json:"client_manual_payment_notification,omitempty"`
	// Determines if e-invoicing is enabled
	EnableEInvoice *bool `json:"enable_e_invoice,omitempty"`
	// The default payment type for expenses
	DefaultExpensePaymentTypeId *string `json:"default_expense_payment_type_id,omitempty"`
	// The e-invoice type
	EInvoiceType *string `json:"e_invoice_type,omitempty"`
	// The mailgun endpoint - used to determine whether US or EU endpoints are used
	MailgunEndpoint *string `json:"mailgun_endpoint,omitempty"`
	// Determines if clients can initiate payments directly from the client portal
	ClientInitiatedPayments *bool `json:"client_initiated_payments,omitempty"`
	// The minimum amount a client can pay
	ClientInitiatedPaymentsMinimum *float32 `json:"client_initiated_payments_minimum,omitempty"`
	// Determines if invoice and quote columns are synced for the PDF rendering, or if they use their own columns
	SyncInvoiceQuoteColumns *bool `json:"sync_invoice_quote_columns,omitempty"`
	// Determines if the task item description is shown on the invoice
	ShowTaskItemDescription *bool `json:"show_task_item_description,omitempty"`
	// Determines if task items can be marked as billable
	AllowBillableTaskItems *bool `json:"allow_billable_task_items,omitempty"`
	// Determines if clients can approve quotes and also pass through a PO Number reference
	AcceptClientInputQuoteApproval *bool `json:"accept_client_input_quote_approval,omitempty"`
	// When using Mailgun or Postmark, the FROM email address can be customized using this setting.
	CustomSendingEmail *string `json:"custom_sending_email,omitempty"`
	// Determines if the PAID stamp is shown on the invoice
	ShowPaidStamp *bool `json:"show_paid_stamp,omitempty"`
	// Determines if the shipping address is shown on the invoice
	ShowShippingAddress *bool `json:"show_shipping_address,omitempty"`
	// The size of the company logo on the PDF - percentage value between 0 and 100
	CompanyLogoSize *int32 `json:"company_logo_size,omitempty"`
	// Determines if the email footer is shown on emails
	ShowEmailFooter *bool `json:"show_email_footer,omitempty"`
	// The alignment of the email body text, options include left / center / right
	EmailAlignment *string `json:"email_alignment,omitempty"`
	// Determines if standard invoices are automatically billed when they are created or due
	AutoBillStandardInvoices *bool `json:"auto_bill_standard_invoices,omitempty"`
	// The Postmark secret API key
	PostmarkSecret *string `json:"postmark_secret,omitempty"`
	// The Mailgun secret API key
	MailgunSecret *string `json:"mailgun_secret,omitempty"`
	// The Mailgun domain
	MailgunDomain *string `json:"mailgun_domain,omitempty"`
	// Determines if an email is sent when an invoice is marked as paid
	SendEmailOnMarkPaid *bool `json:"send_email_on_mark_paid,omitempty"`
	// Determines if vendors can upload files to the portal
	VendorPortalEnableUploads *bool `json:"vendor_portal_enable_uploads,omitempty"`
	// The BESR ID
	BesrId *string `json:"besr_id,omitempty"`
	// The IBAN for the QR code
	QrIban *string `json:"qr_iban,omitempty"`
	// The email subject for purchase orders
	EmailSubjectPurchaseOrder *string `json:"email_subject_purchase_order,omitempty"`
	// The email template for purchase orders
	EmailTemplatePurchaseOrder *string `json:"email_template_purchase_order,omitempty"`
	// Determines if a signature is required on purchase orders
	RequirePurchaseOrderSignature *bool `json:"require_purchase_order_signature,omitempty"`
	// The public notes for purchase orders
	PurchaseOrderPublicNotes *string `json:"purchase_order_public_notes,omitempty"`
	// The terms for purchase orders
	PurchaseOrderTerms *string `json:"purchase_order_terms,omitempty"`
	// The footer for purchase orders
	PurchaseOrderFooter *string `json:"purchase_order_footer,omitempty"`
	// The design id for purchase orders
	PurchaseOrderDesignId *string `json:"purchase_order_design_id,omitempty"`
	// The pattern for purchase order numbers
	PurchaseOrderNumberPattern *string `json:"purchase_order_number_pattern,omitempty"`
	// The counter for purchase order numbers
	PurchaseOrderNumberCounter *int32 `json:"purchase_order_number_counter,omitempty"`
	// The alignment for page numbering: options include left / center / right
	PageNumberingAlignment *string `json:"page_numbering_alignment,omitempty"`
	// Determines if page numbering is enabled on Document PDFs
	PageNumbering *bool `json:"page_numbering,omitempty"`
	// Determines if invoices are automatically archived when they are cancelled
	AutoArchiveInvoiceCancelled *bool `json:"auto_archive_invoice_cancelled,omitempty"`
	// The FROM name for emails when using Custom emailers
	EmailFromName *string `json:"email_from_name,omitempty"`
	// Determines if all tasks are shown on the client portal
	ShowAllTasksClientPortal *bool `json:"show_all_tasks_client_portal,omitempty"`
	// The time that emails are sent. The time is localized to the clients locale, integer values from 1 - 24
	EntitySendTime *int32 `json:"entity_send_time,omitempty"`
	// Determines if the invoice and credit counter are shared
	SharedInvoiceCreditCounter *bool `json:"shared_invoice_credit_counter,omitempty"`
	// The reply to name for emails
	ReplyToName *string `json:"reply_to_name,omitempty"`
	// Determines if empty columns are hidden on PDFs
	HideEmptyColumnsOnPdf *bool `json:"hide_empty_columns_on_pdf,omitempty"`
	// Determines if endless reminders are enabled
	EnableReminderEndless *bool `json:"enable_reminder_endless,omitempty"`
	// Determines if credits can be used as a payment method
	UseCreditsPayment *bool `json:"use_credits_payment,omitempty"`
	// The pattern for recurring invoice numbers
	RecurringInvoiceNumberPattern *string `json:"recurring_invoice_number_pattern,omitempty"`
	// The counter for recurring invoice numbers
	RecurringInvoiceNumberCounter *int32 `json:"recurring_invoice_number_counter,omitempty"`
	// The minimum payment payment
	ClientPortalUnderPaymentMinimum *float32 `json:"client_portal_under_payment_minimum,omitempty"`
	// Determines when the invoices are auto billed
	AutoBillDate *string `json:"auto_bill_date,omitempty"`
	// The primary color for the client portal / document highlights
	PrimaryColor *string `json:"primary_color,omitempty" validate:"regexp=^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"`
	// The secondary color for the client portal / document highlights
	SecondaryColor *string `json:"secondary_color,omitempty" validate:"regexp=^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"`
	// Determines if clients can pay invoices under the invoice amount due
	ClientPortalAllowUnderPayment *bool `json:"client_portal_allow_under_payment,omitempty"`
	// Determines if clients can pay invoices over the invoice amount
	ClientPortalAllowOverPayment *bool `json:"client_portal_allow_over_payment,omitempty"`
	// Determines how autobilling is applied for recurring invoices
	AutoBill *string `json:"auto_bill,omitempty"`
	// The terms which are displayed on the client portal
	ClientPortalTerms *string `json:"client_portal_terms,omitempty"`
	// The privacy policy which is displayed on the client portal
	ClientPortalPrivacyPolicy *string `json:"client_portal_privacy_policy,omitempty"`
	// Determines if clients can register on the client portal
	ClientCanRegister *bool `json:"client_can_register,omitempty"`
	// The design id for the client portal
	PortalDesignId *string `json:"portal_design_id,omitempty"`
	// The late fee percentage for endless late fees
	LateFeeEndlessPercent *float32 `json:"late_fee_endless_percent,omitempty"`
	// The late fee amount for endless late fees
	LateFeeEndlessAmount *float32 `json:"late_fee_endless_amount,omitempty"`
	// Determines if invoices are automatically emailed when they are created
	AutoEmailInvoice *bool `json:"auto_email_invoice,omitempty"`
	// The email signature for emails
	EmailSignature *string `json:"email_signature,omitempty"`
	// The classification for the company
	Classification *string `json:"classification,omitempty"`
	// Default validity period for quotes
	ValidUntil *string `json:"valid_until,omitempty"`
	// Default terms for credit notes
	CreditTerms *string `json:"credit_terms,omitempty"`
	// Default footer for credit notes
	CreditFooter *string `json:"credit_footer,omitempty"`
	// The default currency id
	CurrencyId string `json:"currency_id"`
}

type _CompanySettings CompanySettings

// NewCompanySettings instantiates a new CompanySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanySettings(currencyId string) *CompanySettings {
	this := CompanySettings{}
	this.CurrencyId = currencyId
	return &this
}

// NewCompanySettingsWithDefaults instantiates a new CompanySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanySettingsWithDefaults() *CompanySettings {
	this := CompanySettings{}
	return &this
}

// GetCurrencyId returns the CurrencyId field value
func (o *CompanySettings) GetCurrencyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCurrencyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyId, true
}

// SetCurrencyId sets field value
func (o *CompanySettings) SetCurrencyId(v string) {
	o.CurrencyId = v
}

// GetTimezoneId returns the TimezoneId field value if set, zero value otherwise.
func (o *CompanySettings) GetTimezoneId() string {
	if o == nil || IsNil(o.TimezoneId) {
		var ret string
		return ret
	}
	return *o.TimezoneId
}

// GetTimezoneIdOk returns a tuple with the TimezoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTimezoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimezoneId) {
		return nil, false
	}
	return o.TimezoneId, true
}

// HasTimezoneId returns a boolean if a field has been set.
func (o *CompanySettings) HasTimezoneId() bool {
	if o != nil && !IsNil(o.TimezoneId) {
		return true
	}

	return false
}

// SetTimezoneId gets a reference to the given string and assigns it to the TimezoneId field.
func (o *CompanySettings) SetTimezoneId(v string) {
	o.TimezoneId = &v
}

// GetDateFormatId returns the DateFormatId field value if set, zero value otherwise.
func (o *CompanySettings) GetDateFormatId() string {
	if o == nil || IsNil(o.DateFormatId) {
		var ret string
		return ret
	}
	return *o.DateFormatId
}

// GetDateFormatIdOk returns a tuple with the DateFormatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetDateFormatIdOk() (*string, bool) {
	if o == nil || IsNil(o.DateFormatId) {
		return nil, false
	}
	return o.DateFormatId, true
}

// HasDateFormatId returns a boolean if a field has been set.
func (o *CompanySettings) HasDateFormatId() bool {
	if o != nil && !IsNil(o.DateFormatId) {
		return true
	}

	return false
}

// SetDateFormatId gets a reference to the given string and assigns it to the DateFormatId field.
func (o *CompanySettings) SetDateFormatId(v string) {
	o.DateFormatId = &v
}

// GetMilitaryTime returns the MilitaryTime field value if set, zero value otherwise.
func (o *CompanySettings) GetMilitaryTime() bool {
	if o == nil || IsNil(o.MilitaryTime) {
		var ret bool
		return ret
	}
	return *o.MilitaryTime
}

// GetMilitaryTimeOk returns a tuple with the MilitaryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetMilitaryTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.MilitaryTime) {
		return nil, false
	}
	return o.MilitaryTime, true
}

// HasMilitaryTime returns a boolean if a field has been set.
func (o *CompanySettings) HasMilitaryTime() bool {
	if o != nil && !IsNil(o.MilitaryTime) {
		return true
	}

	return false
}

// SetMilitaryTime gets a reference to the given bool and assigns it to the MilitaryTime field.
func (o *CompanySettings) SetMilitaryTime(v bool) {
	o.MilitaryTime = &v
}

// GetLanguageId returns the LanguageId field value if set, zero value otherwise.
func (o *CompanySettings) GetLanguageId() string {
	if o == nil || IsNil(o.LanguageId) {
		var ret string
		return ret
	}
	return *o.LanguageId
}

// GetLanguageIdOk returns a tuple with the LanguageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetLanguageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageId) {
		return nil, false
	}
	return o.LanguageId, true
}

// HasLanguageId returns a boolean if a field has been set.
func (o *CompanySettings) HasLanguageId() bool {
	if o != nil && !IsNil(o.LanguageId) {
		return true
	}

	return false
}

// SetLanguageId gets a reference to the given string and assigns it to the LanguageId field.
func (o *CompanySettings) SetLanguageId(v string) {
	o.LanguageId = &v
}

// GetShowCurrencyCode returns the ShowCurrencyCode field value if set, zero value otherwise.
func (o *CompanySettings) GetShowCurrencyCode() bool {
	if o == nil || IsNil(o.ShowCurrencyCode) {
		var ret bool
		return ret
	}
	return *o.ShowCurrencyCode
}

// GetShowCurrencyCodeOk returns a tuple with the ShowCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowCurrencyCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCurrencyCode) {
		return nil, false
	}
	return o.ShowCurrencyCode, true
}

// HasShowCurrencyCode returns a boolean if a field has been set.
func (o *CompanySettings) HasShowCurrencyCode() bool {
	if o != nil && !IsNil(o.ShowCurrencyCode) {
		return true
	}

	return false
}

// SetShowCurrencyCode gets a reference to the given bool and assigns it to the ShowCurrencyCode field.
func (o *CompanySettings) SetShowCurrencyCode(v bool) {
	o.ShowCurrencyCode = &v
}

// GetPaymentTerms returns the PaymentTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetPaymentTerms() int32 {
	if o == nil || IsNil(o.PaymentTerms) {
		var ret int32
		return ret
	}
	return *o.PaymentTerms
}

// GetPaymentTermsOk returns a tuple with the PaymentTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPaymentTermsOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentTerms) {
		return nil, false
	}
	return o.PaymentTerms, true
}

// HasPaymentTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasPaymentTerms() bool {
	if o != nil && !IsNil(o.PaymentTerms) {
		return true
	}

	return false
}

// SetPaymentTerms gets a reference to the given int32 and assigns it to the PaymentTerms field.
func (o *CompanySettings) SetPaymentTerms(v int32) {
	o.PaymentTerms = &v
}

// GetCompanyGatewayIds returns the CompanyGatewayIds field value if set, zero value otherwise.
func (o *CompanySettings) GetCompanyGatewayIds() string {
	if o == nil || IsNil(o.CompanyGatewayIds) {
		var ret string
		return ret
	}
	return *o.CompanyGatewayIds
}

// GetCompanyGatewayIdsOk returns a tuple with the CompanyGatewayIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCompanyGatewayIdsOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyGatewayIds) {
		return nil, false
	}
	return o.CompanyGatewayIds, true
}

// HasCompanyGatewayIds returns a boolean if a field has been set.
func (o *CompanySettings) HasCompanyGatewayIds() bool {
	if o != nil && !IsNil(o.CompanyGatewayIds) {
		return true
	}

	return false
}

// SetCompanyGatewayIds gets a reference to the given string and assigns it to the CompanyGatewayIds field.
func (o *CompanySettings) SetCompanyGatewayIds(v string) {
	o.CompanyGatewayIds = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *CompanySettings) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *CompanySettings) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *CompanySettings) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *CompanySettings) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetDefaultTaskRate returns the DefaultTaskRate field value if set, zero value otherwise.
func (o *CompanySettings) GetDefaultTaskRate() float32 {
	if o == nil || IsNil(o.DefaultTaskRate) {
		var ret float32
		return ret
	}
	return *o.DefaultTaskRate
}

// GetDefaultTaskRateOk returns a tuple with the DefaultTaskRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetDefaultTaskRateOk() (*float32, bool) {
	if o == nil || IsNil(o.DefaultTaskRate) {
		return nil, false
	}
	return o.DefaultTaskRate, true
}

// HasDefaultTaskRate returns a boolean if a field has been set.
func (o *CompanySettings) HasDefaultTaskRate() bool {
	if o != nil && !IsNil(o.DefaultTaskRate) {
		return true
	}

	return false
}

// SetDefaultTaskRate gets a reference to the given float32 and assigns it to the DefaultTaskRate field.
func (o *CompanySettings) SetDefaultTaskRate(v float32) {
	o.DefaultTaskRate = &v
}

// GetSendReminders returns the SendReminders field value if set, zero value otherwise.
func (o *CompanySettings) GetSendReminders() bool {
	if o == nil || IsNil(o.SendReminders) {
		var ret bool
		return ret
	}
	return *o.SendReminders
}

// GetSendRemindersOk returns a tuple with the SendReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSendRemindersOk() (*bool, bool) {
	if o == nil || IsNil(o.SendReminders) {
		return nil, false
	}
	return o.SendReminders, true
}

// HasSendReminders returns a boolean if a field has been set.
func (o *CompanySettings) HasSendReminders() bool {
	if o != nil && !IsNil(o.SendReminders) {
		return true
	}

	return false
}

// SetSendReminders gets a reference to the given bool and assigns it to the SendReminders field.
func (o *CompanySettings) SetSendReminders(v bool) {
	o.SendReminders = &v
}

// GetEnableClientPortalTasks returns the EnableClientPortalTasks field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableClientPortalTasks() bool {
	if o == nil || IsNil(o.EnableClientPortalTasks) {
		var ret bool
		return ret
	}
	return *o.EnableClientPortalTasks
}

// GetEnableClientPortalTasksOk returns a tuple with the EnableClientPortalTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableClientPortalTasksOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableClientPortalTasks) {
		return nil, false
	}
	return o.EnableClientPortalTasks, true
}

// HasEnableClientPortalTasks returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableClientPortalTasks() bool {
	if o != nil && !IsNil(o.EnableClientPortalTasks) {
		return true
	}

	return false
}

// SetEnableClientPortalTasks gets a reference to the given bool and assigns it to the EnableClientPortalTasks field.
func (o *CompanySettings) SetEnableClientPortalTasks(v bool) {
	o.EnableClientPortalTasks = &v
}

// GetEmailStyle returns the EmailStyle field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailStyle() string {
	if o == nil || IsNil(o.EmailStyle) {
		var ret string
		return ret
	}
	return *o.EmailStyle
}

// GetEmailStyleOk returns a tuple with the EmailStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailStyleOk() (*string, bool) {
	if o == nil || IsNil(o.EmailStyle) {
		return nil, false
	}
	return o.EmailStyle, true
}

// HasEmailStyle returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailStyle() bool {
	if o != nil && !IsNil(o.EmailStyle) {
		return true
	}

	return false
}

// SetEmailStyle gets a reference to the given string and assigns it to the EmailStyle field.
func (o *CompanySettings) SetEmailStyle(v string) {
	o.EmailStyle = &v
}

// GetReplyToEmail returns the ReplyToEmail field value if set, zero value otherwise.
func (o *CompanySettings) GetReplyToEmail() string {
	if o == nil || IsNil(o.ReplyToEmail) {
		var ret string
		return ret
	}
	return *o.ReplyToEmail
}

// GetReplyToEmailOk returns a tuple with the ReplyToEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetReplyToEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyToEmail) {
		return nil, false
	}
	return o.ReplyToEmail, true
}

// HasReplyToEmail returns a boolean if a field has been set.
func (o *CompanySettings) HasReplyToEmail() bool {
	if o != nil && !IsNil(o.ReplyToEmail) {
		return true
	}

	return false
}

// SetReplyToEmail gets a reference to the given string and assigns it to the ReplyToEmail field.
func (o *CompanySettings) SetReplyToEmail(v string) {
	o.ReplyToEmail = &v
}

// GetBccEmail returns the BccEmail field value if set, zero value otherwise.
func (o *CompanySettings) GetBccEmail() string {
	if o == nil || IsNil(o.BccEmail) {
		var ret string
		return ret
	}
	return *o.BccEmail
}

// GetBccEmailOk returns a tuple with the BccEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetBccEmailOk() (*string, bool) {
	if o == nil || IsNil(o.BccEmail) {
		return nil, false
	}
	return o.BccEmail, true
}

// HasBccEmail returns a boolean if a field has been set.
func (o *CompanySettings) HasBccEmail() bool {
	if o != nil && !IsNil(o.BccEmail) {
		return true
	}

	return false
}

// SetBccEmail gets a reference to the given string and assigns it to the BccEmail field.
func (o *CompanySettings) SetBccEmail(v string) {
	o.BccEmail = &v
}

// GetPdfEmailAttachment returns the PdfEmailAttachment field value if set, zero value otherwise.
func (o *CompanySettings) GetPdfEmailAttachment() bool {
	if o == nil || IsNil(o.PdfEmailAttachment) {
		var ret bool
		return ret
	}
	return *o.PdfEmailAttachment
}

// GetPdfEmailAttachmentOk returns a tuple with the PdfEmailAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPdfEmailAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.PdfEmailAttachment) {
		return nil, false
	}
	return o.PdfEmailAttachment, true
}

// HasPdfEmailAttachment returns a boolean if a field has been set.
func (o *CompanySettings) HasPdfEmailAttachment() bool {
	if o != nil && !IsNil(o.PdfEmailAttachment) {
		return true
	}

	return false
}

// SetPdfEmailAttachment gets a reference to the given bool and assigns it to the PdfEmailAttachment field.
func (o *CompanySettings) SetPdfEmailAttachment(v bool) {
	o.PdfEmailAttachment = &v
}

// GetUblEmailAttachment returns the UblEmailAttachment field value if set, zero value otherwise.
func (o *CompanySettings) GetUblEmailAttachment() bool {
	if o == nil || IsNil(o.UblEmailAttachment) {
		var ret bool
		return ret
	}
	return *o.UblEmailAttachment
}

// GetUblEmailAttachmentOk returns a tuple with the UblEmailAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetUblEmailAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.UblEmailAttachment) {
		return nil, false
	}
	return o.UblEmailAttachment, true
}

// HasUblEmailAttachment returns a boolean if a field has been set.
func (o *CompanySettings) HasUblEmailAttachment() bool {
	if o != nil && !IsNil(o.UblEmailAttachment) {
		return true
	}

	return false
}

// SetUblEmailAttachment gets a reference to the given bool and assigns it to the UblEmailAttachment field.
func (o *CompanySettings) SetUblEmailAttachment(v bool) {
	o.UblEmailAttachment = &v
}

// GetEmailStyleCustom returns the EmailStyleCustom field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailStyleCustom() string {
	if o == nil || IsNil(o.EmailStyleCustom) {
		var ret string
		return ret
	}
	return *o.EmailStyleCustom
}

// GetEmailStyleCustomOk returns a tuple with the EmailStyleCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailStyleCustomOk() (*string, bool) {
	if o == nil || IsNil(o.EmailStyleCustom) {
		return nil, false
	}
	return o.EmailStyleCustom, true
}

// HasEmailStyleCustom returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailStyleCustom() bool {
	if o != nil && !IsNil(o.EmailStyleCustom) {
		return true
	}

	return false
}

// SetEmailStyleCustom gets a reference to the given string and assigns it to the EmailStyleCustom field.
func (o *CompanySettings) SetEmailStyleCustom(v string) {
	o.EmailStyleCustom = &v
}

// GetCounterNumberApplied returns the CounterNumberApplied field value if set, zero value otherwise.
func (o *CompanySettings) GetCounterNumberApplied() string {
	if o == nil || IsNil(o.CounterNumberApplied) {
		var ret string
		return ret
	}
	return *o.CounterNumberApplied
}

// GetCounterNumberAppliedOk returns a tuple with the CounterNumberApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCounterNumberAppliedOk() (*string, bool) {
	if o == nil || IsNil(o.CounterNumberApplied) {
		return nil, false
	}
	return o.CounterNumberApplied, true
}

// HasCounterNumberApplied returns a boolean if a field has been set.
func (o *CompanySettings) HasCounterNumberApplied() bool {
	if o != nil && !IsNil(o.CounterNumberApplied) {
		return true
	}

	return false
}

// SetCounterNumberApplied gets a reference to the given string and assigns it to the CounterNumberApplied field.
func (o *CompanySettings) SetCounterNumberApplied(v string) {
	o.CounterNumberApplied = &v
}

// GetQuoteNumberApplied returns the QuoteNumberApplied field value if set, zero value otherwise.
func (o *CompanySettings) GetQuoteNumberApplied() string {
	if o == nil || IsNil(o.QuoteNumberApplied) {
		var ret string
		return ret
	}
	return *o.QuoteNumberApplied
}

// GetQuoteNumberAppliedOk returns a tuple with the QuoteNumberApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetQuoteNumberAppliedOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumberApplied) {
		return nil, false
	}
	return o.QuoteNumberApplied, true
}

// HasQuoteNumberApplied returns a boolean if a field has been set.
func (o *CompanySettings) HasQuoteNumberApplied() bool {
	if o != nil && !IsNil(o.QuoteNumberApplied) {
		return true
	}

	return false
}

// SetQuoteNumberApplied gets a reference to the given string and assigns it to the QuoteNumberApplied field.
func (o *CompanySettings) SetQuoteNumberApplied(v string) {
	o.QuoteNumberApplied = &v
}

// GetCustomMessageDashboard returns the CustomMessageDashboard field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomMessageDashboard() string {
	if o == nil || IsNil(o.CustomMessageDashboard) {
		var ret string
		return ret
	}
	return *o.CustomMessageDashboard
}

// GetCustomMessageDashboardOk returns a tuple with the CustomMessageDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomMessageDashboardOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessageDashboard) {
		return nil, false
	}
	return o.CustomMessageDashboard, true
}

// HasCustomMessageDashboard returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomMessageDashboard() bool {
	if o != nil && !IsNil(o.CustomMessageDashboard) {
		return true
	}

	return false
}

// SetCustomMessageDashboard gets a reference to the given string and assigns it to the CustomMessageDashboard field.
func (o *CompanySettings) SetCustomMessageDashboard(v string) {
	o.CustomMessageDashboard = &v
}

// GetCustomMessageUnpaidInvoice returns the CustomMessageUnpaidInvoice field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomMessageUnpaidInvoice() string {
	if o == nil || IsNil(o.CustomMessageUnpaidInvoice) {
		var ret string
		return ret
	}
	return *o.CustomMessageUnpaidInvoice
}

// GetCustomMessageUnpaidInvoiceOk returns a tuple with the CustomMessageUnpaidInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomMessageUnpaidInvoiceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessageUnpaidInvoice) {
		return nil, false
	}
	return o.CustomMessageUnpaidInvoice, true
}

// HasCustomMessageUnpaidInvoice returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomMessageUnpaidInvoice() bool {
	if o != nil && !IsNil(o.CustomMessageUnpaidInvoice) {
		return true
	}

	return false
}

// SetCustomMessageUnpaidInvoice gets a reference to the given string and assigns it to the CustomMessageUnpaidInvoice field.
func (o *CompanySettings) SetCustomMessageUnpaidInvoice(v string) {
	o.CustomMessageUnpaidInvoice = &v
}

// GetCustomMessagePaidInvoice returns the CustomMessagePaidInvoice field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomMessagePaidInvoice() string {
	if o == nil || IsNil(o.CustomMessagePaidInvoice) {
		var ret string
		return ret
	}
	return *o.CustomMessagePaidInvoice
}

// GetCustomMessagePaidInvoiceOk returns a tuple with the CustomMessagePaidInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomMessagePaidInvoiceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessagePaidInvoice) {
		return nil, false
	}
	return o.CustomMessagePaidInvoice, true
}

// HasCustomMessagePaidInvoice returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomMessagePaidInvoice() bool {
	if o != nil && !IsNil(o.CustomMessagePaidInvoice) {
		return true
	}

	return false
}

// SetCustomMessagePaidInvoice gets a reference to the given string and assigns it to the CustomMessagePaidInvoice field.
func (o *CompanySettings) SetCustomMessagePaidInvoice(v string) {
	o.CustomMessagePaidInvoice = &v
}

// GetCustomMessageUnapprovedQuote returns the CustomMessageUnapprovedQuote field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomMessageUnapprovedQuote() string {
	if o == nil || IsNil(o.CustomMessageUnapprovedQuote) {
		var ret string
		return ret
	}
	return *o.CustomMessageUnapprovedQuote
}

// GetCustomMessageUnapprovedQuoteOk returns a tuple with the CustomMessageUnapprovedQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomMessageUnapprovedQuoteOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessageUnapprovedQuote) {
		return nil, false
	}
	return o.CustomMessageUnapprovedQuote, true
}

// HasCustomMessageUnapprovedQuote returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomMessageUnapprovedQuote() bool {
	if o != nil && !IsNil(o.CustomMessageUnapprovedQuote) {
		return true
	}

	return false
}

// SetCustomMessageUnapprovedQuote gets a reference to the given string and assigns it to the CustomMessageUnapprovedQuote field.
func (o *CompanySettings) SetCustomMessageUnapprovedQuote(v string) {
	o.CustomMessageUnapprovedQuote = &v
}

// GetLockInvoices returns the LockInvoices field value if set, zero value otherwise.
func (o *CompanySettings) GetLockInvoices() string {
	if o == nil || IsNil(o.LockInvoices) {
		var ret string
		return ret
	}
	return *o.LockInvoices
}

// GetLockInvoicesOk returns a tuple with the LockInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetLockInvoicesOk() (*string, bool) {
	if o == nil || IsNil(o.LockInvoices) {
		return nil, false
	}
	return o.LockInvoices, true
}

// HasLockInvoices returns a boolean if a field has been set.
func (o *CompanySettings) HasLockInvoices() bool {
	if o != nil && !IsNil(o.LockInvoices) {
		return true
	}

	return false
}

// SetLockInvoices gets a reference to the given string and assigns it to the LockInvoices field.
func (o *CompanySettings) SetLockInvoices(v string) {
	o.LockInvoices = &v
}

// GetAutoArchiveInvoice returns the AutoArchiveInvoice field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoArchiveInvoice() bool {
	if o == nil || IsNil(o.AutoArchiveInvoice) {
		var ret bool
		return ret
	}
	return *o.AutoArchiveInvoice
}

// GetAutoArchiveInvoiceOk returns a tuple with the AutoArchiveInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoArchiveInvoiceOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoArchiveInvoice) {
		return nil, false
	}
	return o.AutoArchiveInvoice, true
}

// HasAutoArchiveInvoice returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoArchiveInvoice() bool {
	if o != nil && !IsNil(o.AutoArchiveInvoice) {
		return true
	}

	return false
}

// SetAutoArchiveInvoice gets a reference to the given bool and assigns it to the AutoArchiveInvoice field.
func (o *CompanySettings) SetAutoArchiveInvoice(v bool) {
	o.AutoArchiveInvoice = &v
}

// GetAutoArchiveQuote returns the AutoArchiveQuote field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoArchiveQuote() bool {
	if o == nil || IsNil(o.AutoArchiveQuote) {
		var ret bool
		return ret
	}
	return *o.AutoArchiveQuote
}

// GetAutoArchiveQuoteOk returns a tuple with the AutoArchiveQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoArchiveQuoteOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoArchiveQuote) {
		return nil, false
	}
	return o.AutoArchiveQuote, true
}

// HasAutoArchiveQuote returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoArchiveQuote() bool {
	if o != nil && !IsNil(o.AutoArchiveQuote) {
		return true
	}

	return false
}

// SetAutoArchiveQuote gets a reference to the given bool and assigns it to the AutoArchiveQuote field.
func (o *CompanySettings) SetAutoArchiveQuote(v bool) {
	o.AutoArchiveQuote = &v
}

// GetAutoConvertQuote returns the AutoConvertQuote field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoConvertQuote() bool {
	if o == nil || IsNil(o.AutoConvertQuote) {
		var ret bool
		return ret
	}
	return *o.AutoConvertQuote
}

// GetAutoConvertQuoteOk returns a tuple with the AutoConvertQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoConvertQuoteOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoConvertQuote) {
		return nil, false
	}
	return o.AutoConvertQuote, true
}

// HasAutoConvertQuote returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoConvertQuote() bool {
	if o != nil && !IsNil(o.AutoConvertQuote) {
		return true
	}

	return false
}

// SetAutoConvertQuote gets a reference to the given bool and assigns it to the AutoConvertQuote field.
func (o *CompanySettings) SetAutoConvertQuote(v bool) {
	o.AutoConvertQuote = &v
}

// GetInclusiveTaxes returns the InclusiveTaxes field value if set, zero value otherwise.
func (o *CompanySettings) GetInclusiveTaxes() bool {
	if o == nil || IsNil(o.InclusiveTaxes) {
		var ret bool
		return ret
	}
	return *o.InclusiveTaxes
}

// GetInclusiveTaxesOk returns a tuple with the InclusiveTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInclusiveTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.InclusiveTaxes) {
		return nil, false
	}
	return o.InclusiveTaxes, true
}

// HasInclusiveTaxes returns a boolean if a field has been set.
func (o *CompanySettings) HasInclusiveTaxes() bool {
	if o != nil && !IsNil(o.InclusiveTaxes) {
		return true
	}

	return false
}

// SetInclusiveTaxes gets a reference to the given bool and assigns it to the InclusiveTaxes field.
func (o *CompanySettings) SetInclusiveTaxes(v bool) {
	o.InclusiveTaxes = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *CompanySettings) GetTranslations() map[string]any {
	if o == nil || IsNil(o.Translations) {
		var ret map[string]any
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTranslationsOk() (map[string]any, bool) {
	if o == nil || IsNil(o.Translations) {
		return map[string]any{}, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *CompanySettings) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given map[string]any and assigns it to the Translations field.
func (o *CompanySettings) SetTranslations(v map[string]any) {
	o.Translations = v
}

// GetTaskNumberPattern returns the TaskNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetTaskNumberPattern() string {
	if o == nil || IsNil(o.TaskNumberPattern) {
		var ret string
		return ret
	}
	return *o.TaskNumberPattern
}

// GetTaskNumberPatternOk returns a tuple with the TaskNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaskNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.TaskNumberPattern) {
		return nil, false
	}
	return o.TaskNumberPattern, true
}

// HasTaskNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasTaskNumberPattern() bool {
	if o != nil && !IsNil(o.TaskNumberPattern) {
		return true
	}

	return false
}

// SetTaskNumberPattern gets a reference to the given string and assigns it to the TaskNumberPattern field.
func (o *CompanySettings) SetTaskNumberPattern(v string) {
	o.TaskNumberPattern = &v
}

// GetTaskNumberCounter returns the TaskNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetTaskNumberCounter() int32 {
	if o == nil || IsNil(o.TaskNumberCounter) {
		var ret int32
		return ret
	}
	return *o.TaskNumberCounter
}

// GetTaskNumberCounterOk returns a tuple with the TaskNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaskNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskNumberCounter) {
		return nil, false
	}
	return o.TaskNumberCounter, true
}

// HasTaskNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasTaskNumberCounter() bool {
	if o != nil && !IsNil(o.TaskNumberCounter) {
		return true
	}

	return false
}

// SetTaskNumberCounter gets a reference to the given int32 and assigns it to the TaskNumberCounter field.
func (o *CompanySettings) SetTaskNumberCounter(v int32) {
	o.TaskNumberCounter = &v
}

// GetReminderSendTime returns the ReminderSendTime field value if set, zero value otherwise.
func (o *CompanySettings) GetReminderSendTime() int32 {
	if o == nil || IsNil(o.ReminderSendTime) {
		var ret int32
		return ret
	}
	return *o.ReminderSendTime
}

// GetReminderSendTimeOk returns a tuple with the ReminderSendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetReminderSendTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ReminderSendTime) {
		return nil, false
	}
	return o.ReminderSendTime, true
}

// HasReminderSendTime returns a boolean if a field has been set.
func (o *CompanySettings) HasReminderSendTime() bool {
	if o != nil && !IsNil(o.ReminderSendTime) {
		return true
	}

	return false
}

// SetReminderSendTime gets a reference to the given int32 and assigns it to the ReminderSendTime field.
func (o *CompanySettings) SetReminderSendTime(v int32) {
	o.ReminderSendTime = &v
}

// GetExpenseNumberPattern returns the ExpenseNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetExpenseNumberPattern() string {
	if o == nil || IsNil(o.ExpenseNumberPattern) {
		var ret string
		return ret
	}
	return *o.ExpenseNumberPattern
}

// GetExpenseNumberPatternOk returns a tuple with the ExpenseNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetExpenseNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseNumberPattern) {
		return nil, false
	}
	return o.ExpenseNumberPattern, true
}

// HasExpenseNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasExpenseNumberPattern() bool {
	if o != nil && !IsNil(o.ExpenseNumberPattern) {
		return true
	}

	return false
}

// SetExpenseNumberPattern gets a reference to the given string and assigns it to the ExpenseNumberPattern field.
func (o *CompanySettings) SetExpenseNumberPattern(v string) {
	o.ExpenseNumberPattern = &v
}

// GetExpenseNumberCounter returns the ExpenseNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetExpenseNumberCounter() int32 {
	if o == nil || IsNil(o.ExpenseNumberCounter) {
		var ret int32
		return ret
	}
	return *o.ExpenseNumberCounter
}

// GetExpenseNumberCounterOk returns a tuple with the ExpenseNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetExpenseNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpenseNumberCounter) {
		return nil, false
	}
	return o.ExpenseNumberCounter, true
}

// HasExpenseNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasExpenseNumberCounter() bool {
	if o != nil && !IsNil(o.ExpenseNumberCounter) {
		return true
	}

	return false
}

// SetExpenseNumberCounter gets a reference to the given int32 and assigns it to the ExpenseNumberCounter field.
func (o *CompanySettings) SetExpenseNumberCounter(v int32) {
	o.ExpenseNumberCounter = &v
}

// GetVendorNumberPattern returns the VendorNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetVendorNumberPattern() string {
	if o == nil || IsNil(o.VendorNumberPattern) {
		var ret string
		return ret
	}
	return *o.VendorNumberPattern
}

// GetVendorNumberPatternOk returns a tuple with the VendorNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetVendorNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.VendorNumberPattern) {
		return nil, false
	}
	return o.VendorNumberPattern, true
}

// HasVendorNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasVendorNumberPattern() bool {
	if o != nil && !IsNil(o.VendorNumberPattern) {
		return true
	}

	return false
}

// SetVendorNumberPattern gets a reference to the given string and assigns it to the VendorNumberPattern field.
func (o *CompanySettings) SetVendorNumberPattern(v string) {
	o.VendorNumberPattern = &v
}

// GetVendorNumberCounter returns the VendorNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetVendorNumberCounter() int32 {
	if o == nil || IsNil(o.VendorNumberCounter) {
		var ret int32
		return ret
	}
	return *o.VendorNumberCounter
}

// GetVendorNumberCounterOk returns a tuple with the VendorNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetVendorNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.VendorNumberCounter) {
		return nil, false
	}
	return o.VendorNumberCounter, true
}

// HasVendorNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasVendorNumberCounter() bool {
	if o != nil && !IsNil(o.VendorNumberCounter) {
		return true
	}

	return false
}

// SetVendorNumberCounter gets a reference to the given int32 and assigns it to the VendorNumberCounter field.
func (o *CompanySettings) SetVendorNumberCounter(v int32) {
	o.VendorNumberCounter = &v
}

// GetTicketNumberPattern returns the TicketNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetTicketNumberPattern() string {
	if o == nil || IsNil(o.TicketNumberPattern) {
		var ret string
		return ret
	}
	return *o.TicketNumberPattern
}

// GetTicketNumberPatternOk returns a tuple with the TicketNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTicketNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.TicketNumberPattern) {
		return nil, false
	}
	return o.TicketNumberPattern, true
}

// HasTicketNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasTicketNumberPattern() bool {
	if o != nil && !IsNil(o.TicketNumberPattern) {
		return true
	}

	return false
}

// SetTicketNumberPattern gets a reference to the given string and assigns it to the TicketNumberPattern field.
func (o *CompanySettings) SetTicketNumberPattern(v string) {
	o.TicketNumberPattern = &v
}

// GetTicketNumberCounter returns the TicketNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetTicketNumberCounter() int32 {
	if o == nil || IsNil(o.TicketNumberCounter) {
		var ret int32
		return ret
	}
	return *o.TicketNumberCounter
}

// GetTicketNumberCounterOk returns a tuple with the TicketNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTicketNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.TicketNumberCounter) {
		return nil, false
	}
	return o.TicketNumberCounter, true
}

// HasTicketNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasTicketNumberCounter() bool {
	if o != nil && !IsNil(o.TicketNumberCounter) {
		return true
	}

	return false
}

// SetTicketNumberCounter gets a reference to the given int32 and assigns it to the TicketNumberCounter field.
func (o *CompanySettings) SetTicketNumberCounter(v int32) {
	o.TicketNumberCounter = &v
}

// GetPaymentNumberPattern returns the PaymentNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetPaymentNumberPattern() string {
	if o == nil || IsNil(o.PaymentNumberPattern) {
		var ret string
		return ret
	}
	return *o.PaymentNumberPattern
}

// GetPaymentNumberPatternOk returns a tuple with the PaymentNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPaymentNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentNumberPattern) {
		return nil, false
	}
	return o.PaymentNumberPattern, true
}

// HasPaymentNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasPaymentNumberPattern() bool {
	if o != nil && !IsNil(o.PaymentNumberPattern) {
		return true
	}

	return false
}

// SetPaymentNumberPattern gets a reference to the given string and assigns it to the PaymentNumberPattern field.
func (o *CompanySettings) SetPaymentNumberPattern(v string) {
	o.PaymentNumberPattern = &v
}

// GetPaymentNumberCounter returns the PaymentNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetPaymentNumberCounter() int32 {
	if o == nil || IsNil(o.PaymentNumberCounter) {
		var ret int32
		return ret
	}
	return *o.PaymentNumberCounter
}

// GetPaymentNumberCounterOk returns a tuple with the PaymentNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPaymentNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentNumberCounter) {
		return nil, false
	}
	return o.PaymentNumberCounter, true
}

// HasPaymentNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasPaymentNumberCounter() bool {
	if o != nil && !IsNil(o.PaymentNumberCounter) {
		return true
	}

	return false
}

// SetPaymentNumberCounter gets a reference to the given int32 and assigns it to the PaymentNumberCounter field.
func (o *CompanySettings) SetPaymentNumberCounter(v int32) {
	o.PaymentNumberCounter = &v
}

// GetInvoiceNumberPattern returns the InvoiceNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetInvoiceNumberPattern() string {
	if o == nil || IsNil(o.InvoiceNumberPattern) {
		var ret string
		return ret
	}
	return *o.InvoiceNumberPattern
}

// GetInvoiceNumberPatternOk returns a tuple with the InvoiceNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInvoiceNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumberPattern) {
		return nil, false
	}
	return o.InvoiceNumberPattern, true
}

// HasInvoiceNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasInvoiceNumberPattern() bool {
	if o != nil && !IsNil(o.InvoiceNumberPattern) {
		return true
	}

	return false
}

// SetInvoiceNumberPattern gets a reference to the given string and assigns it to the InvoiceNumberPattern field.
func (o *CompanySettings) SetInvoiceNumberPattern(v string) {
	o.InvoiceNumberPattern = &v
}

// GetInvoiceNumberCounter returns the InvoiceNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetInvoiceNumberCounter() int32 {
	if o == nil || IsNil(o.InvoiceNumberCounter) {
		var ret int32
		return ret
	}
	return *o.InvoiceNumberCounter
}

// GetInvoiceNumberCounterOk returns a tuple with the InvoiceNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInvoiceNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.InvoiceNumberCounter) {
		return nil, false
	}
	return o.InvoiceNumberCounter, true
}

// HasInvoiceNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasInvoiceNumberCounter() bool {
	if o != nil && !IsNil(o.InvoiceNumberCounter) {
		return true
	}

	return false
}

// SetInvoiceNumberCounter gets a reference to the given int32 and assigns it to the InvoiceNumberCounter field.
func (o *CompanySettings) SetInvoiceNumberCounter(v int32) {
	o.InvoiceNumberCounter = &v
}

// GetQuoteNumberPattern returns the QuoteNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetQuoteNumberPattern() string {
	if o == nil || IsNil(o.QuoteNumberPattern) {
		var ret string
		return ret
	}
	return *o.QuoteNumberPattern
}

// GetQuoteNumberPatternOk returns a tuple with the QuoteNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetQuoteNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumberPattern) {
		return nil, false
	}
	return o.QuoteNumberPattern, true
}

// HasQuoteNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasQuoteNumberPattern() bool {
	if o != nil && !IsNil(o.QuoteNumberPattern) {
		return true
	}

	return false
}

// SetQuoteNumberPattern gets a reference to the given string and assigns it to the QuoteNumberPattern field.
func (o *CompanySettings) SetQuoteNumberPattern(v string) {
	o.QuoteNumberPattern = &v
}

// GetQuoteNumberCounter returns the QuoteNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetQuoteNumberCounter() int32 {
	if o == nil || IsNil(o.QuoteNumberCounter) {
		var ret int32
		return ret
	}
	return *o.QuoteNumberCounter
}

// GetQuoteNumberCounterOk returns a tuple with the QuoteNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetQuoteNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.QuoteNumberCounter) {
		return nil, false
	}
	return o.QuoteNumberCounter, true
}

// HasQuoteNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasQuoteNumberCounter() bool {
	if o != nil && !IsNil(o.QuoteNumberCounter) {
		return true
	}

	return false
}

// SetQuoteNumberCounter gets a reference to the given int32 and assigns it to the QuoteNumberCounter field.
func (o *CompanySettings) SetQuoteNumberCounter(v int32) {
	o.QuoteNumberCounter = &v
}

// GetClientNumberPattern returns the ClientNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetClientNumberPattern() string {
	if o == nil || IsNil(o.ClientNumberPattern) {
		var ret string
		return ret
	}
	return *o.ClientNumberPattern
}

// GetClientNumberPatternOk returns a tuple with the ClientNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.ClientNumberPattern) {
		return nil, false
	}
	return o.ClientNumberPattern, true
}

// HasClientNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasClientNumberPattern() bool {
	if o != nil && !IsNil(o.ClientNumberPattern) {
		return true
	}

	return false
}

// SetClientNumberPattern gets a reference to the given string and assigns it to the ClientNumberPattern field.
func (o *CompanySettings) SetClientNumberPattern(v string) {
	o.ClientNumberPattern = &v
}

// GetClientNumberCounter returns the ClientNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetClientNumberCounter() int32 {
	if o == nil || IsNil(o.ClientNumberCounter) {
		var ret int32
		return ret
	}
	return *o.ClientNumberCounter
}

// GetClientNumberCounterOk returns a tuple with the ClientNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.ClientNumberCounter) {
		return nil, false
	}
	return o.ClientNumberCounter, true
}

// HasClientNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasClientNumberCounter() bool {
	if o != nil && !IsNil(o.ClientNumberCounter) {
		return true
	}

	return false
}

// SetClientNumberCounter gets a reference to the given int32 and assigns it to the ClientNumberCounter field.
func (o *CompanySettings) SetClientNumberCounter(v int32) {
	o.ClientNumberCounter = &v
}

// GetCreditNumberPattern returns the CreditNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetCreditNumberPattern() string {
	if o == nil || IsNil(o.CreditNumberPattern) {
		var ret string
		return ret
	}
	return *o.CreditNumberPattern
}

// GetCreditNumberPatternOk returns a tuple with the CreditNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCreditNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.CreditNumberPattern) {
		return nil, false
	}
	return o.CreditNumberPattern, true
}

// HasCreditNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasCreditNumberPattern() bool {
	if o != nil && !IsNil(o.CreditNumberPattern) {
		return true
	}

	return false
}

// SetCreditNumberPattern gets a reference to the given string and assigns it to the CreditNumberPattern field.
func (o *CompanySettings) SetCreditNumberPattern(v string) {
	o.CreditNumberPattern = &v
}

// GetCreditNumberCounter returns the CreditNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetCreditNumberCounter() int32 {
	if o == nil || IsNil(o.CreditNumberCounter) {
		var ret int32
		return ret
	}
	return *o.CreditNumberCounter
}

// GetCreditNumberCounterOk returns a tuple with the CreditNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCreditNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.CreditNumberCounter) {
		return nil, false
	}
	return o.CreditNumberCounter, true
}

// HasCreditNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasCreditNumberCounter() bool {
	if o != nil && !IsNil(o.CreditNumberCounter) {
		return true
	}

	return false
}

// SetCreditNumberCounter gets a reference to the given int32 and assigns it to the CreditNumberCounter field.
func (o *CompanySettings) SetCreditNumberCounter(v int32) {
	o.CreditNumberCounter = &v
}

// GetRecurringInvoiceNumberPrefix returns the RecurringInvoiceNumberPrefix field value if set, zero value otherwise.
func (o *CompanySettings) GetRecurringInvoiceNumberPrefix() string {
	if o == nil || IsNil(o.RecurringInvoiceNumberPrefix) {
		var ret string
		return ret
	}
	return *o.RecurringInvoiceNumberPrefix
}

// GetRecurringInvoiceNumberPrefixOk returns a tuple with the RecurringInvoiceNumberPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetRecurringInvoiceNumberPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringInvoiceNumberPrefix) {
		return nil, false
	}
	return o.RecurringInvoiceNumberPrefix, true
}

// HasRecurringInvoiceNumberPrefix returns a boolean if a field has been set.
func (o *CompanySettings) HasRecurringInvoiceNumberPrefix() bool {
	if o != nil && !IsNil(o.RecurringInvoiceNumberPrefix) {
		return true
	}

	return false
}

// SetRecurringInvoiceNumberPrefix gets a reference to the given string and assigns it to the RecurringInvoiceNumberPrefix field.
func (o *CompanySettings) SetRecurringInvoiceNumberPrefix(v string) {
	o.RecurringInvoiceNumberPrefix = &v
}

// GetResetCounterFrequencyId returns the ResetCounterFrequencyId field value if set, zero value otherwise.
func (o *CompanySettings) GetResetCounterFrequencyId() int32 {
	if o == nil || IsNil(o.ResetCounterFrequencyId) {
		var ret int32
		return ret
	}
	return *o.ResetCounterFrequencyId
}

// GetResetCounterFrequencyIdOk returns a tuple with the ResetCounterFrequencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetResetCounterFrequencyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ResetCounterFrequencyId) {
		return nil, false
	}
	return o.ResetCounterFrequencyId, true
}

// HasResetCounterFrequencyId returns a boolean if a field has been set.
func (o *CompanySettings) HasResetCounterFrequencyId() bool {
	if o != nil && !IsNil(o.ResetCounterFrequencyId) {
		return true
	}

	return false
}

// SetResetCounterFrequencyId gets a reference to the given int32 and assigns it to the ResetCounterFrequencyId field.
func (o *CompanySettings) SetResetCounterFrequencyId(v int32) {
	o.ResetCounterFrequencyId = &v
}

// GetResetCounterDate returns the ResetCounterDate field value if set, zero value otherwise.
func (o *CompanySettings) GetResetCounterDate() string {
	if o == nil || IsNil(o.ResetCounterDate) {
		var ret string
		return ret
	}
	return *o.ResetCounterDate
}

// GetResetCounterDateOk returns a tuple with the ResetCounterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetResetCounterDateOk() (*string, bool) {
	if o == nil || IsNil(o.ResetCounterDate) {
		return nil, false
	}
	return o.ResetCounterDate, true
}

// HasResetCounterDate returns a boolean if a field has been set.
func (o *CompanySettings) HasResetCounterDate() bool {
	if o != nil && !IsNil(o.ResetCounterDate) {
		return true
	}

	return false
}

// SetResetCounterDate gets a reference to the given string and assigns it to the ResetCounterDate field.
func (o *CompanySettings) SetResetCounterDate(v string) {
	o.ResetCounterDate = &v
}

// GetCounterPadding returns the CounterPadding field value if set, zero value otherwise.
func (o *CompanySettings) GetCounterPadding() int32 {
	if o == nil || IsNil(o.CounterPadding) {
		var ret int32
		return ret
	}
	return *o.CounterPadding
}

// GetCounterPaddingOk returns a tuple with the CounterPadding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCounterPaddingOk() (*int32, bool) {
	if o == nil || IsNil(o.CounterPadding) {
		return nil, false
	}
	return o.CounterPadding, true
}

// HasCounterPadding returns a boolean if a field has been set.
func (o *CompanySettings) HasCounterPadding() bool {
	if o != nil && !IsNil(o.CounterPadding) {
		return true
	}

	return false
}

// SetCounterPadding gets a reference to the given int32 and assigns it to the CounterPadding field.
func (o *CompanySettings) SetCounterPadding(v int32) {
	o.CounterPadding = &v
}

// GetSharedInvoiceQuoteCounter returns the SharedInvoiceQuoteCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetSharedInvoiceQuoteCounter() bool {
	if o == nil || IsNil(o.SharedInvoiceQuoteCounter) {
		var ret bool
		return ret
	}
	return *o.SharedInvoiceQuoteCounter
}

// GetSharedInvoiceQuoteCounterOk returns a tuple with the SharedInvoiceQuoteCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSharedInvoiceQuoteCounterOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedInvoiceQuoteCounter) {
		return nil, false
	}
	return o.SharedInvoiceQuoteCounter, true
}

// HasSharedInvoiceQuoteCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasSharedInvoiceQuoteCounter() bool {
	if o != nil && !IsNil(o.SharedInvoiceQuoteCounter) {
		return true
	}

	return false
}

// SetSharedInvoiceQuoteCounter gets a reference to the given bool and assigns it to the SharedInvoiceQuoteCounter field.
func (o *CompanySettings) SetSharedInvoiceQuoteCounter(v bool) {
	o.SharedInvoiceQuoteCounter = &v
}

// GetUpdateProducts returns the UpdateProducts field value if set, zero value otherwise.
func (o *CompanySettings) GetUpdateProducts() bool {
	if o == nil || IsNil(o.UpdateProducts) {
		var ret bool
		return ret
	}
	return *o.UpdateProducts
}

// GetUpdateProductsOk returns a tuple with the UpdateProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetUpdateProductsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateProducts) {
		return nil, false
	}
	return o.UpdateProducts, true
}

// HasUpdateProducts returns a boolean if a field has been set.
func (o *CompanySettings) HasUpdateProducts() bool {
	if o != nil && !IsNil(o.UpdateProducts) {
		return true
	}

	return false
}

// SetUpdateProducts gets a reference to the given bool and assigns it to the UpdateProducts field.
func (o *CompanySettings) SetUpdateProducts(v bool) {
	o.UpdateProducts = &v
}

// GetConvertProducts returns the ConvertProducts field value if set, zero value otherwise.
func (o *CompanySettings) GetConvertProducts() bool {
	if o == nil || IsNil(o.ConvertProducts) {
		var ret bool
		return ret
	}
	return *o.ConvertProducts
}

// GetConvertProductsOk returns a tuple with the ConvertProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetConvertProductsOk() (*bool, bool) {
	if o == nil || IsNil(o.ConvertProducts) {
		return nil, false
	}
	return o.ConvertProducts, true
}

// HasConvertProducts returns a boolean if a field has been set.
func (o *CompanySettings) HasConvertProducts() bool {
	if o != nil && !IsNil(o.ConvertProducts) {
		return true
	}

	return false
}

// SetConvertProducts gets a reference to the given bool and assigns it to the ConvertProducts field.
func (o *CompanySettings) SetConvertProducts(v bool) {
	o.ConvertProducts = &v
}

// GetFillProducts returns the FillProducts field value if set, zero value otherwise.
func (o *CompanySettings) GetFillProducts() bool {
	if o == nil || IsNil(o.FillProducts) {
		var ret bool
		return ret
	}
	return *o.FillProducts
}

// GetFillProductsOk returns a tuple with the FillProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetFillProductsOk() (*bool, bool) {
	if o == nil || IsNil(o.FillProducts) {
		return nil, false
	}
	return o.FillProducts, true
}

// HasFillProducts returns a boolean if a field has been set.
func (o *CompanySettings) HasFillProducts() bool {
	if o != nil && !IsNil(o.FillProducts) {
		return true
	}

	return false
}

// SetFillProducts gets a reference to the given bool and assigns it to the FillProducts field.
func (o *CompanySettings) SetFillProducts(v bool) {
	o.FillProducts = &v
}

// GetInvoiceTerms returns the InvoiceTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetInvoiceTerms() string {
	if o == nil || IsNil(o.InvoiceTerms) {
		var ret string
		return ret
	}
	return *o.InvoiceTerms
}

// GetInvoiceTermsOk returns a tuple with the InvoiceTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInvoiceTermsOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceTerms) {
		return nil, false
	}
	return o.InvoiceTerms, true
}

// HasInvoiceTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasInvoiceTerms() bool {
	if o != nil && !IsNil(o.InvoiceTerms) {
		return true
	}

	return false
}

// SetInvoiceTerms gets a reference to the given string and assigns it to the InvoiceTerms field.
func (o *CompanySettings) SetInvoiceTerms(v string) {
	o.InvoiceTerms = &v
}

// GetQuoteTerms returns the QuoteTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetQuoteTerms() string {
	if o == nil || IsNil(o.QuoteTerms) {
		var ret string
		return ret
	}
	return *o.QuoteTerms
}

// GetQuoteTermsOk returns a tuple with the QuoteTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetQuoteTermsOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteTerms) {
		return nil, false
	}
	return o.QuoteTerms, true
}

// HasQuoteTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasQuoteTerms() bool {
	if o != nil && !IsNil(o.QuoteTerms) {
		return true
	}

	return false
}

// SetQuoteTerms gets a reference to the given string and assigns it to the QuoteTerms field.
func (o *CompanySettings) SetQuoteTerms(v string) {
	o.QuoteTerms = &v
}

// GetInvoiceTaxes returns the InvoiceTaxes field value if set, zero value otherwise.
func (o *CompanySettings) GetInvoiceTaxes() int32 {
	if o == nil || IsNil(o.InvoiceTaxes) {
		var ret int32
		return ret
	}
	return *o.InvoiceTaxes
}

// GetInvoiceTaxesOk returns a tuple with the InvoiceTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInvoiceTaxesOk() (*int32, bool) {
	if o == nil || IsNil(o.InvoiceTaxes) {
		return nil, false
	}
	return o.InvoiceTaxes, true
}

// HasInvoiceTaxes returns a boolean if a field has been set.
func (o *CompanySettings) HasInvoiceTaxes() bool {
	if o != nil && !IsNil(o.InvoiceTaxes) {
		return true
	}

	return false
}

// SetInvoiceTaxes gets a reference to the given int32 and assigns it to the InvoiceTaxes field.
func (o *CompanySettings) SetInvoiceTaxes(v int32) {
	o.InvoiceTaxes = &v
}

// GetInvoiceDesignId returns the InvoiceDesignId field value if set, zero value otherwise.
func (o *CompanySettings) GetInvoiceDesignId() string {
	if o == nil || IsNil(o.InvoiceDesignId) {
		var ret string
		return ret
	}
	return *o.InvoiceDesignId
}

// GetInvoiceDesignIdOk returns a tuple with the InvoiceDesignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInvoiceDesignIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDesignId) {
		return nil, false
	}
	return o.InvoiceDesignId, true
}

// HasInvoiceDesignId returns a boolean if a field has been set.
func (o *CompanySettings) HasInvoiceDesignId() bool {
	if o != nil && !IsNil(o.InvoiceDesignId) {
		return true
	}

	return false
}

// SetInvoiceDesignId gets a reference to the given string and assigns it to the InvoiceDesignId field.
func (o *CompanySettings) SetInvoiceDesignId(v string) {
	o.InvoiceDesignId = &v
}

// GetQuoteDesignId returns the QuoteDesignId field value if set, zero value otherwise.
func (o *CompanySettings) GetQuoteDesignId() string {
	if o == nil || IsNil(o.QuoteDesignId) {
		var ret string
		return ret
	}
	return *o.QuoteDesignId
}

// GetQuoteDesignIdOk returns a tuple with the QuoteDesignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetQuoteDesignIdOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteDesignId) {
		return nil, false
	}
	return o.QuoteDesignId, true
}

// HasQuoteDesignId returns a boolean if a field has been set.
func (o *CompanySettings) HasQuoteDesignId() bool {
	if o != nil && !IsNil(o.QuoteDesignId) {
		return true
	}

	return false
}

// SetQuoteDesignId gets a reference to the given string and assigns it to the QuoteDesignId field.
func (o *CompanySettings) SetQuoteDesignId(v string) {
	o.QuoteDesignId = &v
}

// GetInvoiceFooter returns the InvoiceFooter field value if set, zero value otherwise.
func (o *CompanySettings) GetInvoiceFooter() string {
	if o == nil || IsNil(o.InvoiceFooter) {
		var ret string
		return ret
	}
	return *o.InvoiceFooter
}

// GetInvoiceFooterOk returns a tuple with the InvoiceFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInvoiceFooterOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceFooter) {
		return nil, false
	}
	return o.InvoiceFooter, true
}

// HasInvoiceFooter returns a boolean if a field has been set.
func (o *CompanySettings) HasInvoiceFooter() bool {
	if o != nil && !IsNil(o.InvoiceFooter) {
		return true
	}

	return false
}

// SetInvoiceFooter gets a reference to the given string and assigns it to the InvoiceFooter field.
func (o *CompanySettings) SetInvoiceFooter(v string) {
	o.InvoiceFooter = &v
}

// GetInvoiceLabels returns the InvoiceLabels field value if set, zero value otherwise.
func (o *CompanySettings) GetInvoiceLabels() string {
	if o == nil || IsNil(o.InvoiceLabels) {
		var ret string
		return ret
	}
	return *o.InvoiceLabels
}

// GetInvoiceLabelsOk returns a tuple with the InvoiceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetInvoiceLabelsOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceLabels) {
		return nil, false
	}
	return o.InvoiceLabels, true
}

// HasInvoiceLabels returns a boolean if a field has been set.
func (o *CompanySettings) HasInvoiceLabels() bool {
	if o != nil && !IsNil(o.InvoiceLabels) {
		return true
	}

	return false
}

// SetInvoiceLabels gets a reference to the given string and assigns it to the InvoiceLabels field.
func (o *CompanySettings) SetInvoiceLabels(v string) {
	o.InvoiceLabels = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *CompanySettings) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *CompanySettings) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *CompanySettings) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *CompanySettings) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *CompanySettings) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *CompanySettings) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *CompanySettings) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *CompanySettings) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *CompanySettings) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *CompanySettings) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *CompanySettings) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *CompanySettings) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *CompanySettings) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *CompanySettings) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *CompanySettings) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *CompanySettings) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *CompanySettings) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *CompanySettings) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetPaymentTypeId returns the PaymentTypeId field value if set, zero value otherwise.
func (o *CompanySettings) GetPaymentTypeId() string {
	if o == nil || IsNil(o.PaymentTypeId) {
		var ret string
		return ret
	}
	return *o.PaymentTypeId
}

// GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPaymentTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTypeId) {
		return nil, false
	}
	return o.PaymentTypeId, true
}

// HasPaymentTypeId returns a boolean if a field has been set.
func (o *CompanySettings) HasPaymentTypeId() bool {
	if o != nil && !IsNil(o.PaymentTypeId) {
		return true
	}

	return false
}

// SetPaymentTypeId gets a reference to the given string and assigns it to the PaymentTypeId field.
func (o *CompanySettings) SetPaymentTypeId(v string) {
	o.PaymentTypeId = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomFields() string {
	if o == nil || IsNil(o.CustomFields) {
		var ret string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given string and assigns it to the CustomFields field.
func (o *CompanySettings) SetCustomFields(v string) {
	o.CustomFields = &v
}

// GetEmailFooter returns the EmailFooter field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailFooter() string {
	if o == nil || IsNil(o.EmailFooter) {
		var ret string
		return ret
	}
	return *o.EmailFooter
}

// GetEmailFooterOk returns a tuple with the EmailFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailFooterOk() (*string, bool) {
	if o == nil || IsNil(o.EmailFooter) {
		return nil, false
	}
	return o.EmailFooter, true
}

// HasEmailFooter returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailFooter() bool {
	if o != nil && !IsNil(o.EmailFooter) {
		return true
	}

	return false
}

// SetEmailFooter gets a reference to the given string and assigns it to the EmailFooter field.
func (o *CompanySettings) SetEmailFooter(v string) {
	o.EmailFooter = &v
}

// GetEmailSendingMethod returns the EmailSendingMethod field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSendingMethod() string {
	if o == nil || IsNil(o.EmailSendingMethod) {
		var ret string
		return ret
	}
	return *o.EmailSendingMethod
}

// GetEmailSendingMethodOk returns a tuple with the EmailSendingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSendingMethodOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSendingMethod) {
		return nil, false
	}
	return o.EmailSendingMethod, true
}

// HasEmailSendingMethod returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSendingMethod() bool {
	if o != nil && !IsNil(o.EmailSendingMethod) {
		return true
	}

	return false
}

// SetEmailSendingMethod gets a reference to the given string and assigns it to the EmailSendingMethod field.
func (o *CompanySettings) SetEmailSendingMethod(v string) {
	o.EmailSendingMethod = &v
}

// GetGmailSendingUserId returns the GmailSendingUserId field value if set, zero value otherwise.
func (o *CompanySettings) GetGmailSendingUserId() string {
	if o == nil || IsNil(o.GmailSendingUserId) {
		var ret string
		return ret
	}
	return *o.GmailSendingUserId
}

// GetGmailSendingUserIdOk returns a tuple with the GmailSendingUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetGmailSendingUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.GmailSendingUserId) {
		return nil, false
	}
	return o.GmailSendingUserId, true
}

// HasGmailSendingUserId returns a boolean if a field has been set.
func (o *CompanySettings) HasGmailSendingUserId() bool {
	if o != nil && !IsNil(o.GmailSendingUserId) {
		return true
	}

	return false
}

// SetGmailSendingUserId gets a reference to the given string and assigns it to the GmailSendingUserId field.
func (o *CompanySettings) SetGmailSendingUserId(v string) {
	o.GmailSendingUserId = &v
}

// GetEmailSubjectInvoice returns the EmailSubjectInvoice field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectInvoice() string {
	if o == nil || IsNil(o.EmailSubjectInvoice) {
		var ret string
		return ret
	}
	return *o.EmailSubjectInvoice
}

// GetEmailSubjectInvoiceOk returns a tuple with the EmailSubjectInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectInvoiceOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectInvoice) {
		return nil, false
	}
	return o.EmailSubjectInvoice, true
}

// HasEmailSubjectInvoice returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectInvoice() bool {
	if o != nil && !IsNil(o.EmailSubjectInvoice) {
		return true
	}

	return false
}

// SetEmailSubjectInvoice gets a reference to the given string and assigns it to the EmailSubjectInvoice field.
func (o *CompanySettings) SetEmailSubjectInvoice(v string) {
	o.EmailSubjectInvoice = &v
}

// GetEmailSubjectQuote returns the EmailSubjectQuote field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectQuote() string {
	if o == nil || IsNil(o.EmailSubjectQuote) {
		var ret string
		return ret
	}
	return *o.EmailSubjectQuote
}

// GetEmailSubjectQuoteOk returns a tuple with the EmailSubjectQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectQuoteOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectQuote) {
		return nil, false
	}
	return o.EmailSubjectQuote, true
}

// HasEmailSubjectQuote returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectQuote() bool {
	if o != nil && !IsNil(o.EmailSubjectQuote) {
		return true
	}

	return false
}

// SetEmailSubjectQuote gets a reference to the given string and assigns it to the EmailSubjectQuote field.
func (o *CompanySettings) SetEmailSubjectQuote(v string) {
	o.EmailSubjectQuote = &v
}

// GetEmailSubjectPayment returns the EmailSubjectPayment field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectPayment() string {
	if o == nil || IsNil(o.EmailSubjectPayment) {
		var ret string
		return ret
	}
	return *o.EmailSubjectPayment
}

// GetEmailSubjectPaymentOk returns a tuple with the EmailSubjectPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectPaymentOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectPayment) {
		return nil, false
	}
	return o.EmailSubjectPayment, true
}

// HasEmailSubjectPayment returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectPayment() bool {
	if o != nil && !IsNil(o.EmailSubjectPayment) {
		return true
	}

	return false
}

// SetEmailSubjectPayment gets a reference to the given string and assigns it to the EmailSubjectPayment field.
func (o *CompanySettings) SetEmailSubjectPayment(v string) {
	o.EmailSubjectPayment = &v
}

// GetEmailTemplateInvoice returns the EmailTemplateInvoice field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateInvoice() string {
	if o == nil || IsNil(o.EmailTemplateInvoice) {
		var ret string
		return ret
	}
	return *o.EmailTemplateInvoice
}

// GetEmailTemplateInvoiceOk returns a tuple with the EmailTemplateInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateInvoiceOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateInvoice) {
		return nil, false
	}
	return o.EmailTemplateInvoice, true
}

// HasEmailTemplateInvoice returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateInvoice() bool {
	if o != nil && !IsNil(o.EmailTemplateInvoice) {
		return true
	}

	return false
}

// SetEmailTemplateInvoice gets a reference to the given string and assigns it to the EmailTemplateInvoice field.
func (o *CompanySettings) SetEmailTemplateInvoice(v string) {
	o.EmailTemplateInvoice = &v
}

// GetEmailTemplateQuote returns the EmailTemplateQuote field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateQuote() string {
	if o == nil || IsNil(o.EmailTemplateQuote) {
		var ret string
		return ret
	}
	return *o.EmailTemplateQuote
}

// GetEmailTemplateQuoteOk returns a tuple with the EmailTemplateQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateQuoteOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateQuote) {
		return nil, false
	}
	return o.EmailTemplateQuote, true
}

// HasEmailTemplateQuote returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateQuote() bool {
	if o != nil && !IsNil(o.EmailTemplateQuote) {
		return true
	}

	return false
}

// SetEmailTemplateQuote gets a reference to the given string and assigns it to the EmailTemplateQuote field.
func (o *CompanySettings) SetEmailTemplateQuote(v string) {
	o.EmailTemplateQuote = &v
}

// GetEmailTemplatePayment returns the EmailTemplatePayment field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplatePayment() string {
	if o == nil || IsNil(o.EmailTemplatePayment) {
		var ret string
		return ret
	}
	return *o.EmailTemplatePayment
}

// GetEmailTemplatePaymentOk returns a tuple with the EmailTemplatePayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplatePaymentOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplatePayment) {
		return nil, false
	}
	return o.EmailTemplatePayment, true
}

// HasEmailTemplatePayment returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplatePayment() bool {
	if o != nil && !IsNil(o.EmailTemplatePayment) {
		return true
	}

	return false
}

// SetEmailTemplatePayment gets a reference to the given string and assigns it to the EmailTemplatePayment field.
func (o *CompanySettings) SetEmailTemplatePayment(v string) {
	o.EmailTemplatePayment = &v
}

// GetEmailSubjectReminder1 returns the EmailSubjectReminder1 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectReminder1() string {
	if o == nil || IsNil(o.EmailSubjectReminder1) {
		var ret string
		return ret
	}
	return *o.EmailSubjectReminder1
}

// GetEmailSubjectReminder1Ok returns a tuple with the EmailSubjectReminder1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectReminder1Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectReminder1) {
		return nil, false
	}
	return o.EmailSubjectReminder1, true
}

// HasEmailSubjectReminder1 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectReminder1() bool {
	if o != nil && !IsNil(o.EmailSubjectReminder1) {
		return true
	}

	return false
}

// SetEmailSubjectReminder1 gets a reference to the given string and assigns it to the EmailSubjectReminder1 field.
func (o *CompanySettings) SetEmailSubjectReminder1(v string) {
	o.EmailSubjectReminder1 = &v
}

// GetEmailSubjectReminder2 returns the EmailSubjectReminder2 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectReminder2() string {
	if o == nil || IsNil(o.EmailSubjectReminder2) {
		var ret string
		return ret
	}
	return *o.EmailSubjectReminder2
}

// GetEmailSubjectReminder2Ok returns a tuple with the EmailSubjectReminder2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectReminder2Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectReminder2) {
		return nil, false
	}
	return o.EmailSubjectReminder2, true
}

// HasEmailSubjectReminder2 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectReminder2() bool {
	if o != nil && !IsNil(o.EmailSubjectReminder2) {
		return true
	}

	return false
}

// SetEmailSubjectReminder2 gets a reference to the given string and assigns it to the EmailSubjectReminder2 field.
func (o *CompanySettings) SetEmailSubjectReminder2(v string) {
	o.EmailSubjectReminder2 = &v
}

// GetEmailSubjectReminder3 returns the EmailSubjectReminder3 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectReminder3() string {
	if o == nil || IsNil(o.EmailSubjectReminder3) {
		var ret string
		return ret
	}
	return *o.EmailSubjectReminder3
}

// GetEmailSubjectReminder3Ok returns a tuple with the EmailSubjectReminder3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectReminder3Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectReminder3) {
		return nil, false
	}
	return o.EmailSubjectReminder3, true
}

// HasEmailSubjectReminder3 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectReminder3() bool {
	if o != nil && !IsNil(o.EmailSubjectReminder3) {
		return true
	}

	return false
}

// SetEmailSubjectReminder3 gets a reference to the given string and assigns it to the EmailSubjectReminder3 field.
func (o *CompanySettings) SetEmailSubjectReminder3(v string) {
	o.EmailSubjectReminder3 = &v
}

// GetEmailSubjectReminderEndless returns the EmailSubjectReminderEndless field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectReminderEndless() string {
	if o == nil || IsNil(o.EmailSubjectReminderEndless) {
		var ret string
		return ret
	}
	return *o.EmailSubjectReminderEndless
}

// GetEmailSubjectReminderEndlessOk returns a tuple with the EmailSubjectReminderEndless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectReminderEndlessOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectReminderEndless) {
		return nil, false
	}
	return o.EmailSubjectReminderEndless, true
}

// HasEmailSubjectReminderEndless returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectReminderEndless() bool {
	if o != nil && !IsNil(o.EmailSubjectReminderEndless) {
		return true
	}

	return false
}

// SetEmailSubjectReminderEndless gets a reference to the given string and assigns it to the EmailSubjectReminderEndless field.
func (o *CompanySettings) SetEmailSubjectReminderEndless(v string) {
	o.EmailSubjectReminderEndless = &v
}

// GetEmailTemplateReminder1 returns the EmailTemplateReminder1 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateReminder1() string {
	if o == nil || IsNil(o.EmailTemplateReminder1) {
		var ret string
		return ret
	}
	return *o.EmailTemplateReminder1
}

// GetEmailTemplateReminder1Ok returns a tuple with the EmailTemplateReminder1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateReminder1Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateReminder1) {
		return nil, false
	}
	return o.EmailTemplateReminder1, true
}

// HasEmailTemplateReminder1 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateReminder1() bool {
	if o != nil && !IsNil(o.EmailTemplateReminder1) {
		return true
	}

	return false
}

// SetEmailTemplateReminder1 gets a reference to the given string and assigns it to the EmailTemplateReminder1 field.
func (o *CompanySettings) SetEmailTemplateReminder1(v string) {
	o.EmailTemplateReminder1 = &v
}

// GetEmailTemplateReminder2 returns the EmailTemplateReminder2 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateReminder2() string {
	if o == nil || IsNil(o.EmailTemplateReminder2) {
		var ret string
		return ret
	}
	return *o.EmailTemplateReminder2
}

// GetEmailTemplateReminder2Ok returns a tuple with the EmailTemplateReminder2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateReminder2Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateReminder2) {
		return nil, false
	}
	return o.EmailTemplateReminder2, true
}

// HasEmailTemplateReminder2 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateReminder2() bool {
	if o != nil && !IsNil(o.EmailTemplateReminder2) {
		return true
	}

	return false
}

// SetEmailTemplateReminder2 gets a reference to the given string and assigns it to the EmailTemplateReminder2 field.
func (o *CompanySettings) SetEmailTemplateReminder2(v string) {
	o.EmailTemplateReminder2 = &v
}

// GetEmailTemplateReminder3 returns the EmailTemplateReminder3 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateReminder3() string {
	if o == nil || IsNil(o.EmailTemplateReminder3) {
		var ret string
		return ret
	}
	return *o.EmailTemplateReminder3
}

// GetEmailTemplateReminder3Ok returns a tuple with the EmailTemplateReminder3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateReminder3Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateReminder3) {
		return nil, false
	}
	return o.EmailTemplateReminder3, true
}

// HasEmailTemplateReminder3 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateReminder3() bool {
	if o != nil && !IsNil(o.EmailTemplateReminder3) {
		return true
	}

	return false
}

// SetEmailTemplateReminder3 gets a reference to the given string and assigns it to the EmailTemplateReminder3 field.
func (o *CompanySettings) SetEmailTemplateReminder3(v string) {
	o.EmailTemplateReminder3 = &v
}

// GetEmailTemplateReminderEndless returns the EmailTemplateReminderEndless field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateReminderEndless() string {
	if o == nil || IsNil(o.EmailTemplateReminderEndless) {
		var ret string
		return ret
	}
	return *o.EmailTemplateReminderEndless
}

// GetEmailTemplateReminderEndlessOk returns a tuple with the EmailTemplateReminderEndless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateReminderEndlessOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateReminderEndless) {
		return nil, false
	}
	return o.EmailTemplateReminderEndless, true
}

// HasEmailTemplateReminderEndless returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateReminderEndless() bool {
	if o != nil && !IsNil(o.EmailTemplateReminderEndless) {
		return true
	}

	return false
}

// SetEmailTemplateReminderEndless gets a reference to the given string and assigns it to the EmailTemplateReminderEndless field.
func (o *CompanySettings) SetEmailTemplateReminderEndless(v string) {
	o.EmailTemplateReminderEndless = &v
}

// GetEnablePortalPassword returns the EnablePortalPassword field value if set, zero value otherwise.
func (o *CompanySettings) GetEnablePortalPassword() bool {
	if o == nil || IsNil(o.EnablePortalPassword) {
		var ret bool
		return ret
	}
	return *o.EnablePortalPassword
}

// GetEnablePortalPasswordOk returns a tuple with the EnablePortalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnablePortalPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePortalPassword) {
		return nil, false
	}
	return o.EnablePortalPassword, true
}

// HasEnablePortalPassword returns a boolean if a field has been set.
func (o *CompanySettings) HasEnablePortalPassword() bool {
	if o != nil && !IsNil(o.EnablePortalPassword) {
		return true
	}

	return false
}

// SetEnablePortalPassword gets a reference to the given bool and assigns it to the EnablePortalPassword field.
func (o *CompanySettings) SetEnablePortalPassword(v bool) {
	o.EnablePortalPassword = &v
}

// GetShowAcceptInvoiceTerms returns the ShowAcceptInvoiceTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetShowAcceptInvoiceTerms() bool {
	if o == nil || IsNil(o.ShowAcceptInvoiceTerms) {
		var ret bool
		return ret
	}
	return *o.ShowAcceptInvoiceTerms
}

// GetShowAcceptInvoiceTermsOk returns a tuple with the ShowAcceptInvoiceTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowAcceptInvoiceTermsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowAcceptInvoiceTerms) {
		return nil, false
	}
	return o.ShowAcceptInvoiceTerms, true
}

// HasShowAcceptInvoiceTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasShowAcceptInvoiceTerms() bool {
	if o != nil && !IsNil(o.ShowAcceptInvoiceTerms) {
		return true
	}

	return false
}

// SetShowAcceptInvoiceTerms gets a reference to the given bool and assigns it to the ShowAcceptInvoiceTerms field.
func (o *CompanySettings) SetShowAcceptInvoiceTerms(v bool) {
	o.ShowAcceptInvoiceTerms = &v
}

// GetShowAcceptQuoteTerms returns the ShowAcceptQuoteTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetShowAcceptQuoteTerms() bool {
	if o == nil || IsNil(o.ShowAcceptQuoteTerms) {
		var ret bool
		return ret
	}
	return *o.ShowAcceptQuoteTerms
}

// GetShowAcceptQuoteTermsOk returns a tuple with the ShowAcceptQuoteTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowAcceptQuoteTermsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowAcceptQuoteTerms) {
		return nil, false
	}
	return o.ShowAcceptQuoteTerms, true
}

// HasShowAcceptQuoteTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasShowAcceptQuoteTerms() bool {
	if o != nil && !IsNil(o.ShowAcceptQuoteTerms) {
		return true
	}

	return false
}

// SetShowAcceptQuoteTerms gets a reference to the given bool and assigns it to the ShowAcceptQuoteTerms field.
func (o *CompanySettings) SetShowAcceptQuoteTerms(v bool) {
	o.ShowAcceptQuoteTerms = &v
}

// GetRequireInvoiceSignature returns the RequireInvoiceSignature field value if set, zero value otherwise.
func (o *CompanySettings) GetRequireInvoiceSignature() bool {
	if o == nil || IsNil(o.RequireInvoiceSignature) {
		var ret bool
		return ret
	}
	return *o.RequireInvoiceSignature
}

// GetRequireInvoiceSignatureOk returns a tuple with the RequireInvoiceSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetRequireInvoiceSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireInvoiceSignature) {
		return nil, false
	}
	return o.RequireInvoiceSignature, true
}

// HasRequireInvoiceSignature returns a boolean if a field has been set.
func (o *CompanySettings) HasRequireInvoiceSignature() bool {
	if o != nil && !IsNil(o.RequireInvoiceSignature) {
		return true
	}

	return false
}

// SetRequireInvoiceSignature gets a reference to the given bool and assigns it to the RequireInvoiceSignature field.
func (o *CompanySettings) SetRequireInvoiceSignature(v bool) {
	o.RequireInvoiceSignature = &v
}

// GetRequireQuoteSignature returns the RequireQuoteSignature field value if set, zero value otherwise.
func (o *CompanySettings) GetRequireQuoteSignature() bool {
	if o == nil || IsNil(o.RequireQuoteSignature) {
		var ret bool
		return ret
	}
	return *o.RequireQuoteSignature
}

// GetRequireQuoteSignatureOk returns a tuple with the RequireQuoteSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetRequireQuoteSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireQuoteSignature) {
		return nil, false
	}
	return o.RequireQuoteSignature, true
}

// HasRequireQuoteSignature returns a boolean if a field has been set.
func (o *CompanySettings) HasRequireQuoteSignature() bool {
	if o != nil && !IsNil(o.RequireQuoteSignature) {
		return true
	}

	return false
}

// SetRequireQuoteSignature gets a reference to the given bool and assigns it to the RequireQuoteSignature field.
func (o *CompanySettings) SetRequireQuoteSignature(v bool) {
	o.RequireQuoteSignature = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanySettings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanySettings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanySettings) SetName(v string) {
	o.Name = &v
}

// GetCompanyLogo returns the CompanyLogo field value if set, zero value otherwise.
func (o *CompanySettings) GetCompanyLogo() string {
	if o == nil || IsNil(o.CompanyLogo) {
		var ret string
		return ret
	}
	return *o.CompanyLogo
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCompanyLogoOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyLogo) {
		return nil, false
	}
	return o.CompanyLogo, true
}

// HasCompanyLogo returns a boolean if a field has been set.
func (o *CompanySettings) HasCompanyLogo() bool {
	if o != nil && !IsNil(o.CompanyLogo) {
		return true
	}

	return false
}

// SetCompanyLogo gets a reference to the given string and assigns it to the CompanyLogo field.
func (o *CompanySettings) SetCompanyLogo(v string) {
	o.CompanyLogo = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *CompanySettings) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *CompanySettings) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *CompanySettings) SetWebsite(v string) {
	o.Website = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *CompanySettings) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *CompanySettings) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *CompanySettings) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *CompanySettings) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *CompanySettings) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *CompanySettings) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CompanySettings) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CompanySettings) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CompanySettings) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CompanySettings) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CompanySettings) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CompanySettings) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CompanySettings) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CompanySettings) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CompanySettings) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CompanySettings) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CompanySettings) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CompanySettings) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CompanySettings) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CompanySettings) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CompanySettings) SetEmail(v string) {
	o.Email = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *CompanySettings) GetCountryId() string {
	if o == nil || IsNil(o.CountryId) {
		var ret string
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCountryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *CompanySettings) HasCountryId() bool {
	if o != nil && !IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given string and assigns it to the CountryId field.
func (o *CompanySettings) SetCountryId(v string) {
	o.CountryId = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *CompanySettings) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *CompanySettings) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *CompanySettings) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CompanySettings) GetPageSize() string {
	if o == nil || IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPageSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CompanySettings) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *CompanySettings) SetPageSize(v string) {
	o.PageSize = &v
}

// GetFontSize returns the FontSize field value if set, zero value otherwise.
func (o *CompanySettings) GetFontSize() float32 {
	if o == nil || IsNil(o.FontSize) {
		var ret float32
		return ret
	}
	return *o.FontSize
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetFontSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.FontSize) {
		return nil, false
	}
	return o.FontSize, true
}

// HasFontSize returns a boolean if a field has been set.
func (o *CompanySettings) HasFontSize() bool {
	if o != nil && !IsNil(o.FontSize) {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given float32 and assigns it to the FontSize field.
func (o *CompanySettings) SetFontSize(v float32) {
	o.FontSize = &v
}

// GetPrimaryFont returns the PrimaryFont field value if set, zero value otherwise.
func (o *CompanySettings) GetPrimaryFont() string {
	if o == nil || IsNil(o.PrimaryFont) {
		var ret string
		return ret
	}
	return *o.PrimaryFont
}

// GetPrimaryFontOk returns a tuple with the PrimaryFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPrimaryFontOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryFont) {
		return nil, false
	}
	return o.PrimaryFont, true
}

// HasPrimaryFont returns a boolean if a field has been set.
func (o *CompanySettings) HasPrimaryFont() bool {
	if o != nil && !IsNil(o.PrimaryFont) {
		return true
	}

	return false
}

// SetPrimaryFont gets a reference to the given string and assigns it to the PrimaryFont field.
func (o *CompanySettings) SetPrimaryFont(v string) {
	o.PrimaryFont = &v
}

// GetSecondaryFont returns the SecondaryFont field value if set, zero value otherwise.
func (o *CompanySettings) GetSecondaryFont() string {
	if o == nil || IsNil(o.SecondaryFont) {
		var ret string
		return ret
	}
	return *o.SecondaryFont
}

// GetSecondaryFontOk returns a tuple with the SecondaryFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSecondaryFontOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryFont) {
		return nil, false
	}
	return o.SecondaryFont, true
}

// HasSecondaryFont returns a boolean if a field has been set.
func (o *CompanySettings) HasSecondaryFont() bool {
	if o != nil && !IsNil(o.SecondaryFont) {
		return true
	}

	return false
}

// SetSecondaryFont gets a reference to the given string and assigns it to the SecondaryFont field.
func (o *CompanySettings) SetSecondaryFont(v string) {
	o.SecondaryFont = &v
}

// GetHidePaidToDate returns the HidePaidToDate field value if set, zero value otherwise.
func (o *CompanySettings) GetHidePaidToDate() bool {
	if o == nil || IsNil(o.HidePaidToDate) {
		var ret bool
		return ret
	}
	return *o.HidePaidToDate
}

// GetHidePaidToDateOk returns a tuple with the HidePaidToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetHidePaidToDateOk() (*bool, bool) {
	if o == nil || IsNil(o.HidePaidToDate) {
		return nil, false
	}
	return o.HidePaidToDate, true
}

// HasHidePaidToDate returns a boolean if a field has been set.
func (o *CompanySettings) HasHidePaidToDate() bool {
	if o != nil && !IsNil(o.HidePaidToDate) {
		return true
	}

	return false
}

// SetHidePaidToDate gets a reference to the given bool and assigns it to the HidePaidToDate field.
func (o *CompanySettings) SetHidePaidToDate(v bool) {
	o.HidePaidToDate = &v
}

// GetEmbedDocuments returns the EmbedDocuments field value if set, zero value otherwise.
func (o *CompanySettings) GetEmbedDocuments() bool {
	if o == nil || IsNil(o.EmbedDocuments) {
		var ret bool
		return ret
	}
	return *o.EmbedDocuments
}

// GetEmbedDocumentsOk returns a tuple with the EmbedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmbedDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.EmbedDocuments) {
		return nil, false
	}
	return o.EmbedDocuments, true
}

// HasEmbedDocuments returns a boolean if a field has been set.
func (o *CompanySettings) HasEmbedDocuments() bool {
	if o != nil && !IsNil(o.EmbedDocuments) {
		return true
	}

	return false
}

// SetEmbedDocuments gets a reference to the given bool and assigns it to the EmbedDocuments field.
func (o *CompanySettings) SetEmbedDocuments(v bool) {
	o.EmbedDocuments = &v
}

// GetAllPagesHeader returns the AllPagesHeader field value if set, zero value otherwise.
func (o *CompanySettings) GetAllPagesHeader() bool {
	if o == nil || IsNil(o.AllPagesHeader) {
		var ret bool
		return ret
	}
	return *o.AllPagesHeader
}

// GetAllPagesHeaderOk returns a tuple with the AllPagesHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAllPagesHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.AllPagesHeader) {
		return nil, false
	}
	return o.AllPagesHeader, true
}

// HasAllPagesHeader returns a boolean if a field has been set.
func (o *CompanySettings) HasAllPagesHeader() bool {
	if o != nil && !IsNil(o.AllPagesHeader) {
		return true
	}

	return false
}

// SetAllPagesHeader gets a reference to the given bool and assigns it to the AllPagesHeader field.
func (o *CompanySettings) SetAllPagesHeader(v bool) {
	o.AllPagesHeader = &v
}

// GetAllPagesFooter returns the AllPagesFooter field value if set, zero value otherwise.
func (o *CompanySettings) GetAllPagesFooter() bool {
	if o == nil || IsNil(o.AllPagesFooter) {
		var ret bool
		return ret
	}
	return *o.AllPagesFooter
}

// GetAllPagesFooterOk returns a tuple with the AllPagesFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAllPagesFooterOk() (*bool, bool) {
	if o == nil || IsNil(o.AllPagesFooter) {
		return nil, false
	}
	return o.AllPagesFooter, true
}

// HasAllPagesFooter returns a boolean if a field has been set.
func (o *CompanySettings) HasAllPagesFooter() bool {
	if o != nil && !IsNil(o.AllPagesFooter) {
		return true
	}

	return false
}

// SetAllPagesFooter gets a reference to the given bool and assigns it to the AllPagesFooter field.
func (o *CompanySettings) SetAllPagesFooter(v bool) {
	o.AllPagesFooter = &v
}

// GetDocumentEmailAttachment returns the DocumentEmailAttachment field value if set, zero value otherwise.
func (o *CompanySettings) GetDocumentEmailAttachment() bool {
	if o == nil || IsNil(o.DocumentEmailAttachment) {
		var ret bool
		return ret
	}
	return *o.DocumentEmailAttachment
}

// GetDocumentEmailAttachmentOk returns a tuple with the DocumentEmailAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetDocumentEmailAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.DocumentEmailAttachment) {
		return nil, false
	}
	return o.DocumentEmailAttachment, true
}

// HasDocumentEmailAttachment returns a boolean if a field has been set.
func (o *CompanySettings) HasDocumentEmailAttachment() bool {
	if o != nil && !IsNil(o.DocumentEmailAttachment) {
		return true
	}

	return false
}

// SetDocumentEmailAttachment gets a reference to the given bool and assigns it to the DocumentEmailAttachment field.
func (o *CompanySettings) SetDocumentEmailAttachment(v bool) {
	o.DocumentEmailAttachment = &v
}

// GetEnableClientPortalPassword returns the EnableClientPortalPassword field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableClientPortalPassword() bool {
	if o == nil || IsNil(o.EnableClientPortalPassword) {
		var ret bool
		return ret
	}
	return *o.EnableClientPortalPassword
}

// GetEnableClientPortalPasswordOk returns a tuple with the EnableClientPortalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableClientPortalPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableClientPortalPassword) {
		return nil, false
	}
	return o.EnableClientPortalPassword, true
}

// HasEnableClientPortalPassword returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableClientPortalPassword() bool {
	if o != nil && !IsNil(o.EnableClientPortalPassword) {
		return true
	}

	return false
}

// SetEnableClientPortalPassword gets a reference to the given bool and assigns it to the EnableClientPortalPassword field.
func (o *CompanySettings) SetEnableClientPortalPassword(v bool) {
	o.EnableClientPortalPassword = &v
}

// GetEnableEmailMarkup returns the EnableEmailMarkup field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableEmailMarkup() bool {
	if o == nil || IsNil(o.EnableEmailMarkup) {
		var ret bool
		return ret
	}
	return *o.EnableEmailMarkup
}

// GetEnableEmailMarkupOk returns a tuple with the EnableEmailMarkup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableEmailMarkupOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEmailMarkup) {
		return nil, false
	}
	return o.EnableEmailMarkup, true
}

// HasEnableEmailMarkup returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableEmailMarkup() bool {
	if o != nil && !IsNil(o.EnableEmailMarkup) {
		return true
	}

	return false
}

// SetEnableEmailMarkup gets a reference to the given bool and assigns it to the EnableEmailMarkup field.
func (o *CompanySettings) SetEnableEmailMarkup(v bool) {
	o.EnableEmailMarkup = &v
}

// GetEnableClientPortalDashboard returns the EnableClientPortalDashboard field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableClientPortalDashboard() bool {
	if o == nil || IsNil(o.EnableClientPortalDashboard) {
		var ret bool
		return ret
	}
	return *o.EnableClientPortalDashboard
}

// GetEnableClientPortalDashboardOk returns a tuple with the EnableClientPortalDashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableClientPortalDashboardOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableClientPortalDashboard) {
		return nil, false
	}
	return o.EnableClientPortalDashboard, true
}

// HasEnableClientPortalDashboard returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableClientPortalDashboard() bool {
	if o != nil && !IsNil(o.EnableClientPortalDashboard) {
		return true
	}

	return false
}

// SetEnableClientPortalDashboard gets a reference to the given bool and assigns it to the EnableClientPortalDashboard field.
func (o *CompanySettings) SetEnableClientPortalDashboard(v bool) {
	o.EnableClientPortalDashboard = &v
}

// GetEnableClientPortal returns the EnableClientPortal field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableClientPortal() bool {
	if o == nil || IsNil(o.EnableClientPortal) {
		var ret bool
		return ret
	}
	return *o.EnableClientPortal
}

// GetEnableClientPortalOk returns a tuple with the EnableClientPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableClientPortalOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableClientPortal) {
		return nil, false
	}
	return o.EnableClientPortal, true
}

// HasEnableClientPortal returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableClientPortal() bool {
	if o != nil && !IsNil(o.EnableClientPortal) {
		return true
	}

	return false
}

// SetEnableClientPortal gets a reference to the given bool and assigns it to the EnableClientPortal field.
func (o *CompanySettings) SetEnableClientPortal(v bool) {
	o.EnableClientPortal = &v
}

// GetEmailTemplateStatement returns the EmailTemplateStatement field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateStatement() string {
	if o == nil || IsNil(o.EmailTemplateStatement) {
		var ret string
		return ret
	}
	return *o.EmailTemplateStatement
}

// GetEmailTemplateStatementOk returns a tuple with the EmailTemplateStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateStatementOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateStatement) {
		return nil, false
	}
	return o.EmailTemplateStatement, true
}

// HasEmailTemplateStatement returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateStatement() bool {
	if o != nil && !IsNil(o.EmailTemplateStatement) {
		return true
	}

	return false
}

// SetEmailTemplateStatement gets a reference to the given string and assigns it to the EmailTemplateStatement field.
func (o *CompanySettings) SetEmailTemplateStatement(v string) {
	o.EmailTemplateStatement = &v
}

// GetEmailSubjectStatement returns the EmailSubjectStatement field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectStatement() string {
	if o == nil || IsNil(o.EmailSubjectStatement) {
		var ret string
		return ret
	}
	return *o.EmailSubjectStatement
}

// GetEmailSubjectStatementOk returns a tuple with the EmailSubjectStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectStatementOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectStatement) {
		return nil, false
	}
	return o.EmailSubjectStatement, true
}

// HasEmailSubjectStatement returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectStatement() bool {
	if o != nil && !IsNil(o.EmailSubjectStatement) {
		return true
	}

	return false
}

// SetEmailSubjectStatement gets a reference to the given string and assigns it to the EmailSubjectStatement field.
func (o *CompanySettings) SetEmailSubjectStatement(v string) {
	o.EmailSubjectStatement = &v
}

// GetSignatureOnPdf returns the SignatureOnPdf field value if set, zero value otherwise.
func (o *CompanySettings) GetSignatureOnPdf() bool {
	if o == nil || IsNil(o.SignatureOnPdf) {
		var ret bool
		return ret
	}
	return *o.SignatureOnPdf
}

// GetSignatureOnPdfOk returns a tuple with the SignatureOnPdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSignatureOnPdfOk() (*bool, bool) {
	if o == nil || IsNil(o.SignatureOnPdf) {
		return nil, false
	}
	return o.SignatureOnPdf, true
}

// HasSignatureOnPdf returns a boolean if a field has been set.
func (o *CompanySettings) HasSignatureOnPdf() bool {
	if o != nil && !IsNil(o.SignatureOnPdf) {
		return true
	}

	return false
}

// SetSignatureOnPdf gets a reference to the given bool and assigns it to the SignatureOnPdf field.
func (o *CompanySettings) SetSignatureOnPdf(v bool) {
	o.SignatureOnPdf = &v
}

// GetQuoteFooter returns the QuoteFooter field value if set, zero value otherwise.
func (o *CompanySettings) GetQuoteFooter() string {
	if o == nil || IsNil(o.QuoteFooter) {
		var ret string
		return ret
	}
	return *o.QuoteFooter
}

// GetQuoteFooterOk returns a tuple with the QuoteFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetQuoteFooterOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteFooter) {
		return nil, false
	}
	return o.QuoteFooter, true
}

// HasQuoteFooter returns a boolean if a field has been set.
func (o *CompanySettings) HasQuoteFooter() bool {
	if o != nil && !IsNil(o.QuoteFooter) {
		return true
	}

	return false
}

// SetQuoteFooter gets a reference to the given string and assigns it to the QuoteFooter field.
func (o *CompanySettings) SetQuoteFooter(v string) {
	o.QuoteFooter = &v
}

// GetEmailSubjectCustom1 returns the EmailSubjectCustom1 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectCustom1() string {
	if o == nil || IsNil(o.EmailSubjectCustom1) {
		var ret string
		return ret
	}
	return *o.EmailSubjectCustom1
}

// GetEmailSubjectCustom1Ok returns a tuple with the EmailSubjectCustom1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectCustom1Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectCustom1) {
		return nil, false
	}
	return o.EmailSubjectCustom1, true
}

// HasEmailSubjectCustom1 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectCustom1() bool {
	if o != nil && !IsNil(o.EmailSubjectCustom1) {
		return true
	}

	return false
}

// SetEmailSubjectCustom1 gets a reference to the given string and assigns it to the EmailSubjectCustom1 field.
func (o *CompanySettings) SetEmailSubjectCustom1(v string) {
	o.EmailSubjectCustom1 = &v
}

// GetEmailSubjectCustom2 returns the EmailSubjectCustom2 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectCustom2() string {
	if o == nil || IsNil(o.EmailSubjectCustom2) {
		var ret string
		return ret
	}
	return *o.EmailSubjectCustom2
}

// GetEmailSubjectCustom2Ok returns a tuple with the EmailSubjectCustom2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectCustom2Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectCustom2) {
		return nil, false
	}
	return o.EmailSubjectCustom2, true
}

// HasEmailSubjectCustom2 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectCustom2() bool {
	if o != nil && !IsNil(o.EmailSubjectCustom2) {
		return true
	}

	return false
}

// SetEmailSubjectCustom2 gets a reference to the given string and assigns it to the EmailSubjectCustom2 field.
func (o *CompanySettings) SetEmailSubjectCustom2(v string) {
	o.EmailSubjectCustom2 = &v
}

// GetEmailSubjectCustom3 returns the EmailSubjectCustom3 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectCustom3() string {
	if o == nil || IsNil(o.EmailSubjectCustom3) {
		var ret string
		return ret
	}
	return *o.EmailSubjectCustom3
}

// GetEmailSubjectCustom3Ok returns a tuple with the EmailSubjectCustom3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectCustom3Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectCustom3) {
		return nil, false
	}
	return o.EmailSubjectCustom3, true
}

// HasEmailSubjectCustom3 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectCustom3() bool {
	if o != nil && !IsNil(o.EmailSubjectCustom3) {
		return true
	}

	return false
}

// SetEmailSubjectCustom3 gets a reference to the given string and assigns it to the EmailSubjectCustom3 field.
func (o *CompanySettings) SetEmailSubjectCustom3(v string) {
	o.EmailSubjectCustom3 = &v
}

// GetEmailTemplateCustom1 returns the EmailTemplateCustom1 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateCustom1() string {
	if o == nil || IsNil(o.EmailTemplateCustom1) {
		var ret string
		return ret
	}
	return *o.EmailTemplateCustom1
}

// GetEmailTemplateCustom1Ok returns a tuple with the EmailTemplateCustom1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateCustom1Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateCustom1) {
		return nil, false
	}
	return o.EmailTemplateCustom1, true
}

// HasEmailTemplateCustom1 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateCustom1() bool {
	if o != nil && !IsNil(o.EmailTemplateCustom1) {
		return true
	}

	return false
}

// SetEmailTemplateCustom1 gets a reference to the given string and assigns it to the EmailTemplateCustom1 field.
func (o *CompanySettings) SetEmailTemplateCustom1(v string) {
	o.EmailTemplateCustom1 = &v
}

// GetEmailTemplateCustom2 returns the EmailTemplateCustom2 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateCustom2() string {
	if o == nil || IsNil(o.EmailTemplateCustom2) {
		var ret string
		return ret
	}
	return *o.EmailTemplateCustom2
}

// GetEmailTemplateCustom2Ok returns a tuple with the EmailTemplateCustom2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateCustom2Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateCustom2) {
		return nil, false
	}
	return o.EmailTemplateCustom2, true
}

// HasEmailTemplateCustom2 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateCustom2() bool {
	if o != nil && !IsNil(o.EmailTemplateCustom2) {
		return true
	}

	return false
}

// SetEmailTemplateCustom2 gets a reference to the given string and assigns it to the EmailTemplateCustom2 field.
func (o *CompanySettings) SetEmailTemplateCustom2(v string) {
	o.EmailTemplateCustom2 = &v
}

// GetEmailTemplateCustom3 returns the EmailTemplateCustom3 field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplateCustom3() string {
	if o == nil || IsNil(o.EmailTemplateCustom3) {
		var ret string
		return ret
	}
	return *o.EmailTemplateCustom3
}

// GetEmailTemplateCustom3Ok returns a tuple with the EmailTemplateCustom3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplateCustom3Ok() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplateCustom3) {
		return nil, false
	}
	return o.EmailTemplateCustom3, true
}

// HasEmailTemplateCustom3 returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplateCustom3() bool {
	if o != nil && !IsNil(o.EmailTemplateCustom3) {
		return true
	}

	return false
}

// SetEmailTemplateCustom3 gets a reference to the given string and assigns it to the EmailTemplateCustom3 field.
func (o *CompanySettings) SetEmailTemplateCustom3(v string) {
	o.EmailTemplateCustom3 = &v
}

// GetEnableReminder1 returns the EnableReminder1 field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableReminder1() bool {
	if o == nil || IsNil(o.EnableReminder1) {
		var ret bool
		return ret
	}
	return *o.EnableReminder1
}

// GetEnableReminder1Ok returns a tuple with the EnableReminder1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableReminder1Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableReminder1) {
		return nil, false
	}
	return o.EnableReminder1, true
}

// HasEnableReminder1 returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableReminder1() bool {
	if o != nil && !IsNil(o.EnableReminder1) {
		return true
	}

	return false
}

// SetEnableReminder1 gets a reference to the given bool and assigns it to the EnableReminder1 field.
func (o *CompanySettings) SetEnableReminder1(v bool) {
	o.EnableReminder1 = &v
}

// GetEnableReminder2 returns the EnableReminder2 field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableReminder2() bool {
	if o == nil || IsNil(o.EnableReminder2) {
		var ret bool
		return ret
	}
	return *o.EnableReminder2
}

// GetEnableReminder2Ok returns a tuple with the EnableReminder2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableReminder2Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableReminder2) {
		return nil, false
	}
	return o.EnableReminder2, true
}

// HasEnableReminder2 returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableReminder2() bool {
	if o != nil && !IsNil(o.EnableReminder2) {
		return true
	}

	return false
}

// SetEnableReminder2 gets a reference to the given bool and assigns it to the EnableReminder2 field.
func (o *CompanySettings) SetEnableReminder2(v bool) {
	o.EnableReminder2 = &v
}

// GetEnableReminder3 returns the EnableReminder3 field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableReminder3() bool {
	if o == nil || IsNil(o.EnableReminder3) {
		var ret bool
		return ret
	}
	return *o.EnableReminder3
}

// GetEnableReminder3Ok returns a tuple with the EnableReminder3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableReminder3Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableReminder3) {
		return nil, false
	}
	return o.EnableReminder3, true
}

// HasEnableReminder3 returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableReminder3() bool {
	if o != nil && !IsNil(o.EnableReminder3) {
		return true
	}

	return false
}

// SetEnableReminder3 gets a reference to the given bool and assigns it to the EnableReminder3 field.
func (o *CompanySettings) SetEnableReminder3(v bool) {
	o.EnableReminder3 = &v
}

// GetNumDaysReminder1 returns the NumDaysReminder1 field value if set, zero value otherwise.
func (o *CompanySettings) GetNumDaysReminder1() int32 {
	if o == nil || IsNil(o.NumDaysReminder1) {
		var ret int32
		return ret
	}
	return *o.NumDaysReminder1
}

// GetNumDaysReminder1Ok returns a tuple with the NumDaysReminder1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetNumDaysReminder1Ok() (*int32, bool) {
	if o == nil || IsNil(o.NumDaysReminder1) {
		return nil, false
	}
	return o.NumDaysReminder1, true
}

// HasNumDaysReminder1 returns a boolean if a field has been set.
func (o *CompanySettings) HasNumDaysReminder1() bool {
	if o != nil && !IsNil(o.NumDaysReminder1) {
		return true
	}

	return false
}

// SetNumDaysReminder1 gets a reference to the given int32 and assigns it to the NumDaysReminder1 field.
func (o *CompanySettings) SetNumDaysReminder1(v int32) {
	o.NumDaysReminder1 = &v
}

// GetNumDaysReminder2 returns the NumDaysReminder2 field value if set, zero value otherwise.
func (o *CompanySettings) GetNumDaysReminder2() int32 {
	if o == nil || IsNil(o.NumDaysReminder2) {
		var ret int32
		return ret
	}
	return *o.NumDaysReminder2
}

// GetNumDaysReminder2Ok returns a tuple with the NumDaysReminder2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetNumDaysReminder2Ok() (*int32, bool) {
	if o == nil || IsNil(o.NumDaysReminder2) {
		return nil, false
	}
	return o.NumDaysReminder2, true
}

// HasNumDaysReminder2 returns a boolean if a field has been set.
func (o *CompanySettings) HasNumDaysReminder2() bool {
	if o != nil && !IsNil(o.NumDaysReminder2) {
		return true
	}

	return false
}

// SetNumDaysReminder2 gets a reference to the given int32 and assigns it to the NumDaysReminder2 field.
func (o *CompanySettings) SetNumDaysReminder2(v int32) {
	o.NumDaysReminder2 = &v
}

// GetNumDaysReminder3 returns the NumDaysReminder3 field value if set, zero value otherwise.
func (o *CompanySettings) GetNumDaysReminder3() int32 {
	if o == nil || IsNil(o.NumDaysReminder3) {
		var ret int32
		return ret
	}
	return *o.NumDaysReminder3
}

// GetNumDaysReminder3Ok returns a tuple with the NumDaysReminder3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetNumDaysReminder3Ok() (*int32, bool) {
	if o == nil || IsNil(o.NumDaysReminder3) {
		return nil, false
	}
	return o.NumDaysReminder3, true
}

// HasNumDaysReminder3 returns a boolean if a field has been set.
func (o *CompanySettings) HasNumDaysReminder3() bool {
	if o != nil && !IsNil(o.NumDaysReminder3) {
		return true
	}

	return false
}

// SetNumDaysReminder3 gets a reference to the given int32 and assigns it to the NumDaysReminder3 field.
func (o *CompanySettings) SetNumDaysReminder3(v int32) {
	o.NumDaysReminder3 = &v
}

// GetScheduleReminder1 returns the ScheduleReminder1 field value if set, zero value otherwise.
func (o *CompanySettings) GetScheduleReminder1() string {
	if o == nil || IsNil(o.ScheduleReminder1) {
		var ret string
		return ret
	}
	return *o.ScheduleReminder1
}

// GetScheduleReminder1Ok returns a tuple with the ScheduleReminder1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetScheduleReminder1Ok() (*string, bool) {
	if o == nil || IsNil(o.ScheduleReminder1) {
		return nil, false
	}
	return o.ScheduleReminder1, true
}

// HasScheduleReminder1 returns a boolean if a field has been set.
func (o *CompanySettings) HasScheduleReminder1() bool {
	if o != nil && !IsNil(o.ScheduleReminder1) {
		return true
	}

	return false
}

// SetScheduleReminder1 gets a reference to the given string and assigns it to the ScheduleReminder1 field.
func (o *CompanySettings) SetScheduleReminder1(v string) {
	o.ScheduleReminder1 = &v
}

// GetScheduleReminder2 returns the ScheduleReminder2 field value if set, zero value otherwise.
func (o *CompanySettings) GetScheduleReminder2() string {
	if o == nil || IsNil(o.ScheduleReminder2) {
		var ret string
		return ret
	}
	return *o.ScheduleReminder2
}

// GetScheduleReminder2Ok returns a tuple with the ScheduleReminder2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetScheduleReminder2Ok() (*string, bool) {
	if o == nil || IsNil(o.ScheduleReminder2) {
		return nil, false
	}
	return o.ScheduleReminder2, true
}

// HasScheduleReminder2 returns a boolean if a field has been set.
func (o *CompanySettings) HasScheduleReminder2() bool {
	if o != nil && !IsNil(o.ScheduleReminder2) {
		return true
	}

	return false
}

// SetScheduleReminder2 gets a reference to the given string and assigns it to the ScheduleReminder2 field.
func (o *CompanySettings) SetScheduleReminder2(v string) {
	o.ScheduleReminder2 = &v
}

// GetScheduleReminder3 returns the ScheduleReminder3 field value if set, zero value otherwise.
func (o *CompanySettings) GetScheduleReminder3() string {
	if o == nil || IsNil(o.ScheduleReminder3) {
		var ret string
		return ret
	}
	return *o.ScheduleReminder3
}

// GetScheduleReminder3Ok returns a tuple with the ScheduleReminder3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetScheduleReminder3Ok() (*string, bool) {
	if o == nil || IsNil(o.ScheduleReminder3) {
		return nil, false
	}
	return o.ScheduleReminder3, true
}

// HasScheduleReminder3 returns a boolean if a field has been set.
func (o *CompanySettings) HasScheduleReminder3() bool {
	if o != nil && !IsNil(o.ScheduleReminder3) {
		return true
	}

	return false
}

// SetScheduleReminder3 gets a reference to the given string and assigns it to the ScheduleReminder3 field.
func (o *CompanySettings) SetScheduleReminder3(v string) {
	o.ScheduleReminder3 = &v
}

// GetLateFeeAmount1 returns the LateFeeAmount1 field value if set, zero value otherwise.
func (o *CompanySettings) GetLateFeeAmount1() float32 {
	if o == nil || IsNil(o.LateFeeAmount1) {
		var ret float32
		return ret
	}
	return *o.LateFeeAmount1
}

// GetLateFeeAmount1Ok returns a tuple with the LateFeeAmount1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetLateFeeAmount1Ok() (*float32, bool) {
	if o == nil || IsNil(o.LateFeeAmount1) {
		return nil, false
	}
	return o.LateFeeAmount1, true
}

// HasLateFeeAmount1 returns a boolean if a field has been set.
func (o *CompanySettings) HasLateFeeAmount1() bool {
	if o != nil && !IsNil(o.LateFeeAmount1) {
		return true
	}

	return false
}

// SetLateFeeAmount1 gets a reference to the given float32 and assigns it to the LateFeeAmount1 field.
func (o *CompanySettings) SetLateFeeAmount1(v float32) {
	o.LateFeeAmount1 = &v
}

// GetLateFeeAmount2 returns the LateFeeAmount2 field value if set, zero value otherwise.
func (o *CompanySettings) GetLateFeeAmount2() float32 {
	if o == nil || IsNil(o.LateFeeAmount2) {
		var ret float32
		return ret
	}
	return *o.LateFeeAmount2
}

// GetLateFeeAmount2Ok returns a tuple with the LateFeeAmount2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetLateFeeAmount2Ok() (*float32, bool) {
	if o == nil || IsNil(o.LateFeeAmount2) {
		return nil, false
	}
	return o.LateFeeAmount2, true
}

// HasLateFeeAmount2 returns a boolean if a field has been set.
func (o *CompanySettings) HasLateFeeAmount2() bool {
	if o != nil && !IsNil(o.LateFeeAmount2) {
		return true
	}

	return false
}

// SetLateFeeAmount2 gets a reference to the given float32 and assigns it to the LateFeeAmount2 field.
func (o *CompanySettings) SetLateFeeAmount2(v float32) {
	o.LateFeeAmount2 = &v
}

// GetLateFeeAmount3 returns the LateFeeAmount3 field value if set, zero value otherwise.
func (o *CompanySettings) GetLateFeeAmount3() float32 {
	if o == nil || IsNil(o.LateFeeAmount3) {
		var ret float32
		return ret
	}
	return *o.LateFeeAmount3
}

// GetLateFeeAmount3Ok returns a tuple with the LateFeeAmount3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetLateFeeAmount3Ok() (*float32, bool) {
	if o == nil || IsNil(o.LateFeeAmount3) {
		return nil, false
	}
	return o.LateFeeAmount3, true
}

// HasLateFeeAmount3 returns a boolean if a field has been set.
func (o *CompanySettings) HasLateFeeAmount3() bool {
	if o != nil && !IsNil(o.LateFeeAmount3) {
		return true
	}

	return false
}

// SetLateFeeAmount3 gets a reference to the given float32 and assigns it to the LateFeeAmount3 field.
func (o *CompanySettings) SetLateFeeAmount3(v float32) {
	o.LateFeeAmount3 = &v
}

// GetEndlessReminderFrequencyId returns the EndlessReminderFrequencyId field value if set, zero value otherwise.
func (o *CompanySettings) GetEndlessReminderFrequencyId() string {
	if o == nil || IsNil(o.EndlessReminderFrequencyId) {
		var ret string
		return ret
	}
	return *o.EndlessReminderFrequencyId
}

// GetEndlessReminderFrequencyIdOk returns a tuple with the EndlessReminderFrequencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEndlessReminderFrequencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndlessReminderFrequencyId) {
		return nil, false
	}
	return o.EndlessReminderFrequencyId, true
}

// HasEndlessReminderFrequencyId returns a boolean if a field has been set.
func (o *CompanySettings) HasEndlessReminderFrequencyId() bool {
	if o != nil && !IsNil(o.EndlessReminderFrequencyId) {
		return true
	}

	return false
}

// SetEndlessReminderFrequencyId gets a reference to the given string and assigns it to the EndlessReminderFrequencyId field.
func (o *CompanySettings) SetEndlessReminderFrequencyId(v string) {
	o.EndlessReminderFrequencyId = &v
}

// GetClientOnlinePaymentNotification returns the ClientOnlinePaymentNotification field value if set, zero value otherwise.
func (o *CompanySettings) GetClientOnlinePaymentNotification() bool {
	if o == nil || IsNil(o.ClientOnlinePaymentNotification) {
		var ret bool
		return ret
	}
	return *o.ClientOnlinePaymentNotification
}

// GetClientOnlinePaymentNotificationOk returns a tuple with the ClientOnlinePaymentNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientOnlinePaymentNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientOnlinePaymentNotification) {
		return nil, false
	}
	return o.ClientOnlinePaymentNotification, true
}

// HasClientOnlinePaymentNotification returns a boolean if a field has been set.
func (o *CompanySettings) HasClientOnlinePaymentNotification() bool {
	if o != nil && !IsNil(o.ClientOnlinePaymentNotification) {
		return true
	}

	return false
}

// SetClientOnlinePaymentNotification gets a reference to the given bool and assigns it to the ClientOnlinePaymentNotification field.
func (o *CompanySettings) SetClientOnlinePaymentNotification(v bool) {
	o.ClientOnlinePaymentNotification = &v
}

// GetClientManualPaymentNotification returns the ClientManualPaymentNotification field value if set, zero value otherwise.
func (o *CompanySettings) GetClientManualPaymentNotification() bool {
	if o == nil || IsNil(o.ClientManualPaymentNotification) {
		var ret bool
		return ret
	}
	return *o.ClientManualPaymentNotification
}

// GetClientManualPaymentNotificationOk returns a tuple with the ClientManualPaymentNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientManualPaymentNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientManualPaymentNotification) {
		return nil, false
	}
	return o.ClientManualPaymentNotification, true
}

// HasClientManualPaymentNotification returns a boolean if a field has been set.
func (o *CompanySettings) HasClientManualPaymentNotification() bool {
	if o != nil && !IsNil(o.ClientManualPaymentNotification) {
		return true
	}

	return false
}

// SetClientManualPaymentNotification gets a reference to the given bool and assigns it to the ClientManualPaymentNotification field.
func (o *CompanySettings) SetClientManualPaymentNotification(v bool) {
	o.ClientManualPaymentNotification = &v
}

// GetEnableEInvoice returns the EnableEInvoice field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableEInvoice() bool {
	if o == nil || IsNil(o.EnableEInvoice) {
		var ret bool
		return ret
	}
	return *o.EnableEInvoice
}

// GetEnableEInvoiceOk returns a tuple with the EnableEInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableEInvoiceOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEInvoice) {
		return nil, false
	}
	return o.EnableEInvoice, true
}

// HasEnableEInvoice returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableEInvoice() bool {
	if o != nil && !IsNil(o.EnableEInvoice) {
		return true
	}

	return false
}

// SetEnableEInvoice gets a reference to the given bool and assigns it to the EnableEInvoice field.
func (o *CompanySettings) SetEnableEInvoice(v bool) {
	o.EnableEInvoice = &v
}

// GetDefaultExpensePaymentTypeId returns the DefaultExpensePaymentTypeId field value if set, zero value otherwise.
func (o *CompanySettings) GetDefaultExpensePaymentTypeId() string {
	if o == nil || IsNil(o.DefaultExpensePaymentTypeId) {
		var ret string
		return ret
	}
	return *o.DefaultExpensePaymentTypeId
}

// GetDefaultExpensePaymentTypeIdOk returns a tuple with the DefaultExpensePaymentTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetDefaultExpensePaymentTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultExpensePaymentTypeId) {
		return nil, false
	}
	return o.DefaultExpensePaymentTypeId, true
}

// HasDefaultExpensePaymentTypeId returns a boolean if a field has been set.
func (o *CompanySettings) HasDefaultExpensePaymentTypeId() bool {
	if o != nil && !IsNil(o.DefaultExpensePaymentTypeId) {
		return true
	}

	return false
}

// SetDefaultExpensePaymentTypeId gets a reference to the given string and assigns it to the DefaultExpensePaymentTypeId field.
func (o *CompanySettings) SetDefaultExpensePaymentTypeId(v string) {
	o.DefaultExpensePaymentTypeId = &v
}

// GetEInvoiceType returns the EInvoiceType field value if set, zero value otherwise.
func (o *CompanySettings) GetEInvoiceType() string {
	if o == nil || IsNil(o.EInvoiceType) {
		var ret string
		return ret
	}
	return *o.EInvoiceType
}

// GetEInvoiceTypeOk returns a tuple with the EInvoiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEInvoiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EInvoiceType) {
		return nil, false
	}
	return o.EInvoiceType, true
}

// HasEInvoiceType returns a boolean if a field has been set.
func (o *CompanySettings) HasEInvoiceType() bool {
	if o != nil && !IsNil(o.EInvoiceType) {
		return true
	}

	return false
}

// SetEInvoiceType gets a reference to the given string and assigns it to the EInvoiceType field.
func (o *CompanySettings) SetEInvoiceType(v string) {
	o.EInvoiceType = &v
}

// GetMailgunEndpoint returns the MailgunEndpoint field value if set, zero value otherwise.
func (o *CompanySettings) GetMailgunEndpoint() string {
	if o == nil || IsNil(o.MailgunEndpoint) {
		var ret string
		return ret
	}
	return *o.MailgunEndpoint
}

// GetMailgunEndpointOk returns a tuple with the MailgunEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetMailgunEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.MailgunEndpoint) {
		return nil, false
	}
	return o.MailgunEndpoint, true
}

// HasMailgunEndpoint returns a boolean if a field has been set.
func (o *CompanySettings) HasMailgunEndpoint() bool {
	if o != nil && !IsNil(o.MailgunEndpoint) {
		return true
	}

	return false
}

// SetMailgunEndpoint gets a reference to the given string and assigns it to the MailgunEndpoint field.
func (o *CompanySettings) SetMailgunEndpoint(v string) {
	o.MailgunEndpoint = &v
}

// GetClientInitiatedPayments returns the ClientInitiatedPayments field value if set, zero value otherwise.
func (o *CompanySettings) GetClientInitiatedPayments() bool {
	if o == nil || IsNil(o.ClientInitiatedPayments) {
		var ret bool
		return ret
	}
	return *o.ClientInitiatedPayments
}

// GetClientInitiatedPaymentsOk returns a tuple with the ClientInitiatedPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientInitiatedPaymentsOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientInitiatedPayments) {
		return nil, false
	}
	return o.ClientInitiatedPayments, true
}

// HasClientInitiatedPayments returns a boolean if a field has been set.
func (o *CompanySettings) HasClientInitiatedPayments() bool {
	if o != nil && !IsNil(o.ClientInitiatedPayments) {
		return true
	}

	return false
}

// SetClientInitiatedPayments gets a reference to the given bool and assigns it to the ClientInitiatedPayments field.
func (o *CompanySettings) SetClientInitiatedPayments(v bool) {
	o.ClientInitiatedPayments = &v
}

// GetClientInitiatedPaymentsMinimum returns the ClientInitiatedPaymentsMinimum field value if set, zero value otherwise.
func (o *CompanySettings) GetClientInitiatedPaymentsMinimum() float32 {
	if o == nil || IsNil(o.ClientInitiatedPaymentsMinimum) {
		var ret float32
		return ret
	}
	return *o.ClientInitiatedPaymentsMinimum
}

// GetClientInitiatedPaymentsMinimumOk returns a tuple with the ClientInitiatedPaymentsMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientInitiatedPaymentsMinimumOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientInitiatedPaymentsMinimum) {
		return nil, false
	}
	return o.ClientInitiatedPaymentsMinimum, true
}

// HasClientInitiatedPaymentsMinimum returns a boolean if a field has been set.
func (o *CompanySettings) HasClientInitiatedPaymentsMinimum() bool {
	if o != nil && !IsNil(o.ClientInitiatedPaymentsMinimum) {
		return true
	}

	return false
}

// SetClientInitiatedPaymentsMinimum gets a reference to the given float32 and assigns it to the ClientInitiatedPaymentsMinimum field.
func (o *CompanySettings) SetClientInitiatedPaymentsMinimum(v float32) {
	o.ClientInitiatedPaymentsMinimum = &v
}

// GetSyncInvoiceQuoteColumns returns the SyncInvoiceQuoteColumns field value if set, zero value otherwise.
func (o *CompanySettings) GetSyncInvoiceQuoteColumns() bool {
	if o == nil || IsNil(o.SyncInvoiceQuoteColumns) {
		var ret bool
		return ret
	}
	return *o.SyncInvoiceQuoteColumns
}

// GetSyncInvoiceQuoteColumnsOk returns a tuple with the SyncInvoiceQuoteColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSyncInvoiceQuoteColumnsOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncInvoiceQuoteColumns) {
		return nil, false
	}
	return o.SyncInvoiceQuoteColumns, true
}

// HasSyncInvoiceQuoteColumns returns a boolean if a field has been set.
func (o *CompanySettings) HasSyncInvoiceQuoteColumns() bool {
	if o != nil && !IsNil(o.SyncInvoiceQuoteColumns) {
		return true
	}

	return false
}

// SetSyncInvoiceQuoteColumns gets a reference to the given bool and assigns it to the SyncInvoiceQuoteColumns field.
func (o *CompanySettings) SetSyncInvoiceQuoteColumns(v bool) {
	o.SyncInvoiceQuoteColumns = &v
}

// GetShowTaskItemDescription returns the ShowTaskItemDescription field value if set, zero value otherwise.
func (o *CompanySettings) GetShowTaskItemDescription() bool {
	if o == nil || IsNil(o.ShowTaskItemDescription) {
		var ret bool
		return ret
	}
	return *o.ShowTaskItemDescription
}

// GetShowTaskItemDescriptionOk returns a tuple with the ShowTaskItemDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowTaskItemDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowTaskItemDescription) {
		return nil, false
	}
	return o.ShowTaskItemDescription, true
}

// HasShowTaskItemDescription returns a boolean if a field has been set.
func (o *CompanySettings) HasShowTaskItemDescription() bool {
	if o != nil && !IsNil(o.ShowTaskItemDescription) {
		return true
	}

	return false
}

// SetShowTaskItemDescription gets a reference to the given bool and assigns it to the ShowTaskItemDescription field.
func (o *CompanySettings) SetShowTaskItemDescription(v bool) {
	o.ShowTaskItemDescription = &v
}

// GetAllowBillableTaskItems returns the AllowBillableTaskItems field value if set, zero value otherwise.
func (o *CompanySettings) GetAllowBillableTaskItems() bool {
	if o == nil || IsNil(o.AllowBillableTaskItems) {
		var ret bool
		return ret
	}
	return *o.AllowBillableTaskItems
}

// GetAllowBillableTaskItemsOk returns a tuple with the AllowBillableTaskItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAllowBillableTaskItemsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowBillableTaskItems) {
		return nil, false
	}
	return o.AllowBillableTaskItems, true
}

// HasAllowBillableTaskItems returns a boolean if a field has been set.
func (o *CompanySettings) HasAllowBillableTaskItems() bool {
	if o != nil && !IsNil(o.AllowBillableTaskItems) {
		return true
	}

	return false
}

// SetAllowBillableTaskItems gets a reference to the given bool and assigns it to the AllowBillableTaskItems field.
func (o *CompanySettings) SetAllowBillableTaskItems(v bool) {
	o.AllowBillableTaskItems = &v
}

// GetAcceptClientInputQuoteApproval returns the AcceptClientInputQuoteApproval field value if set, zero value otherwise.
func (o *CompanySettings) GetAcceptClientInputQuoteApproval() bool {
	if o == nil || IsNil(o.AcceptClientInputQuoteApproval) {
		var ret bool
		return ret
	}
	return *o.AcceptClientInputQuoteApproval
}

// GetAcceptClientInputQuoteApprovalOk returns a tuple with the AcceptClientInputQuoteApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAcceptClientInputQuoteApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptClientInputQuoteApproval) {
		return nil, false
	}
	return o.AcceptClientInputQuoteApproval, true
}

// HasAcceptClientInputQuoteApproval returns a boolean if a field has been set.
func (o *CompanySettings) HasAcceptClientInputQuoteApproval() bool {
	if o != nil && !IsNil(o.AcceptClientInputQuoteApproval) {
		return true
	}

	return false
}

// SetAcceptClientInputQuoteApproval gets a reference to the given bool and assigns it to the AcceptClientInputQuoteApproval field.
func (o *CompanySettings) SetAcceptClientInputQuoteApproval(v bool) {
	o.AcceptClientInputQuoteApproval = &v
}

// GetCustomSendingEmail returns the CustomSendingEmail field value if set, zero value otherwise.
func (o *CompanySettings) GetCustomSendingEmail() string {
	if o == nil || IsNil(o.CustomSendingEmail) {
		var ret string
		return ret
	}
	return *o.CustomSendingEmail
}

// GetCustomSendingEmailOk returns a tuple with the CustomSendingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCustomSendingEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CustomSendingEmail) {
		return nil, false
	}
	return o.CustomSendingEmail, true
}

// HasCustomSendingEmail returns a boolean if a field has been set.
func (o *CompanySettings) HasCustomSendingEmail() bool {
	if o != nil && !IsNil(o.CustomSendingEmail) {
		return true
	}

	return false
}

// SetCustomSendingEmail gets a reference to the given string and assigns it to the CustomSendingEmail field.
func (o *CompanySettings) SetCustomSendingEmail(v string) {
	o.CustomSendingEmail = &v
}

// GetShowPaidStamp returns the ShowPaidStamp field value if set, zero value otherwise.
func (o *CompanySettings) GetShowPaidStamp() bool {
	if o == nil || IsNil(o.ShowPaidStamp) {
		var ret bool
		return ret
	}
	return *o.ShowPaidStamp
}

// GetShowPaidStampOk returns a tuple with the ShowPaidStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowPaidStampOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPaidStamp) {
		return nil, false
	}
	return o.ShowPaidStamp, true
}

// HasShowPaidStamp returns a boolean if a field has been set.
func (o *CompanySettings) HasShowPaidStamp() bool {
	if o != nil && !IsNil(o.ShowPaidStamp) {
		return true
	}

	return false
}

// SetShowPaidStamp gets a reference to the given bool and assigns it to the ShowPaidStamp field.
func (o *CompanySettings) SetShowPaidStamp(v bool) {
	o.ShowPaidStamp = &v
}

// GetShowShippingAddress returns the ShowShippingAddress field value if set, zero value otherwise.
func (o *CompanySettings) GetShowShippingAddress() bool {
	if o == nil || IsNil(o.ShowShippingAddress) {
		var ret bool
		return ret
	}
	return *o.ShowShippingAddress
}

// GetShowShippingAddressOk returns a tuple with the ShowShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowShippingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowShippingAddress) {
		return nil, false
	}
	return o.ShowShippingAddress, true
}

// HasShowShippingAddress returns a boolean if a field has been set.
func (o *CompanySettings) HasShowShippingAddress() bool {
	if o != nil && !IsNil(o.ShowShippingAddress) {
		return true
	}

	return false
}

// SetShowShippingAddress gets a reference to the given bool and assigns it to the ShowShippingAddress field.
func (o *CompanySettings) SetShowShippingAddress(v bool) {
	o.ShowShippingAddress = &v
}

// GetCompanyLogoSize returns the CompanyLogoSize field value if set, zero value otherwise.
func (o *CompanySettings) GetCompanyLogoSize() int32 {
	if o == nil || IsNil(o.CompanyLogoSize) {
		var ret int32
		return ret
	}
	return *o.CompanyLogoSize
}

// GetCompanyLogoSizeOk returns a tuple with the CompanyLogoSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCompanyLogoSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.CompanyLogoSize) {
		return nil, false
	}
	return o.CompanyLogoSize, true
}

// HasCompanyLogoSize returns a boolean if a field has been set.
func (o *CompanySettings) HasCompanyLogoSize() bool {
	if o != nil && !IsNil(o.CompanyLogoSize) {
		return true
	}

	return false
}

// SetCompanyLogoSize gets a reference to the given int32 and assigns it to the CompanyLogoSize field.
func (o *CompanySettings) SetCompanyLogoSize(v int32) {
	o.CompanyLogoSize = &v
}

// GetShowEmailFooter returns the ShowEmailFooter field value if set, zero value otherwise.
func (o *CompanySettings) GetShowEmailFooter() bool {
	if o == nil || IsNil(o.ShowEmailFooter) {
		var ret bool
		return ret
	}
	return *o.ShowEmailFooter
}

// GetShowEmailFooterOk returns a tuple with the ShowEmailFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowEmailFooterOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowEmailFooter) {
		return nil, false
	}
	return o.ShowEmailFooter, true
}

// HasShowEmailFooter returns a boolean if a field has been set.
func (o *CompanySettings) HasShowEmailFooter() bool {
	if o != nil && !IsNil(o.ShowEmailFooter) {
		return true
	}

	return false
}

// SetShowEmailFooter gets a reference to the given bool and assigns it to the ShowEmailFooter field.
func (o *CompanySettings) SetShowEmailFooter(v bool) {
	o.ShowEmailFooter = &v
}

// GetEmailAlignment returns the EmailAlignment field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailAlignment() string {
	if o == nil || IsNil(o.EmailAlignment) {
		var ret string
		return ret
	}
	return *o.EmailAlignment
}

// GetEmailAlignmentOk returns a tuple with the EmailAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailAlignmentOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAlignment) {
		return nil, false
	}
	return o.EmailAlignment, true
}

// HasEmailAlignment returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailAlignment() bool {
	if o != nil && !IsNil(o.EmailAlignment) {
		return true
	}

	return false
}

// SetEmailAlignment gets a reference to the given string and assigns it to the EmailAlignment field.
func (o *CompanySettings) SetEmailAlignment(v string) {
	o.EmailAlignment = &v
}

// GetAutoBillStandardInvoices returns the AutoBillStandardInvoices field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoBillStandardInvoices() bool {
	if o == nil || IsNil(o.AutoBillStandardInvoices) {
		var ret bool
		return ret
	}
	return *o.AutoBillStandardInvoices
}

// GetAutoBillStandardInvoicesOk returns a tuple with the AutoBillStandardInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoBillStandardInvoicesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoBillStandardInvoices) {
		return nil, false
	}
	return o.AutoBillStandardInvoices, true
}

// HasAutoBillStandardInvoices returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoBillStandardInvoices() bool {
	if o != nil && !IsNil(o.AutoBillStandardInvoices) {
		return true
	}

	return false
}

// SetAutoBillStandardInvoices gets a reference to the given bool and assigns it to the AutoBillStandardInvoices field.
func (o *CompanySettings) SetAutoBillStandardInvoices(v bool) {
	o.AutoBillStandardInvoices = &v
}

// GetPostmarkSecret returns the PostmarkSecret field value if set, zero value otherwise.
func (o *CompanySettings) GetPostmarkSecret() string {
	if o == nil || IsNil(o.PostmarkSecret) {
		var ret string
		return ret
	}
	return *o.PostmarkSecret
}

// GetPostmarkSecretOk returns a tuple with the PostmarkSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPostmarkSecretOk() (*string, bool) {
	if o == nil || IsNil(o.PostmarkSecret) {
		return nil, false
	}
	return o.PostmarkSecret, true
}

// HasPostmarkSecret returns a boolean if a field has been set.
func (o *CompanySettings) HasPostmarkSecret() bool {
	if o != nil && !IsNil(o.PostmarkSecret) {
		return true
	}

	return false
}

// SetPostmarkSecret gets a reference to the given string and assigns it to the PostmarkSecret field.
func (o *CompanySettings) SetPostmarkSecret(v string) {
	o.PostmarkSecret = &v
}

// GetMailgunSecret returns the MailgunSecret field value if set, zero value otherwise.
func (o *CompanySettings) GetMailgunSecret() string {
	if o == nil || IsNil(o.MailgunSecret) {
		var ret string
		return ret
	}
	return *o.MailgunSecret
}

// GetMailgunSecretOk returns a tuple with the MailgunSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetMailgunSecretOk() (*string, bool) {
	if o == nil || IsNil(o.MailgunSecret) {
		return nil, false
	}
	return o.MailgunSecret, true
}

// HasMailgunSecret returns a boolean if a field has been set.
func (o *CompanySettings) HasMailgunSecret() bool {
	if o != nil && !IsNil(o.MailgunSecret) {
		return true
	}

	return false
}

// SetMailgunSecret gets a reference to the given string and assigns it to the MailgunSecret field.
func (o *CompanySettings) SetMailgunSecret(v string) {
	o.MailgunSecret = &v
}

// GetMailgunDomain returns the MailgunDomain field value if set, zero value otherwise.
func (o *CompanySettings) GetMailgunDomain() string {
	if o == nil || IsNil(o.MailgunDomain) {
		var ret string
		return ret
	}
	return *o.MailgunDomain
}

// GetMailgunDomainOk returns a tuple with the MailgunDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetMailgunDomainOk() (*string, bool) {
	if o == nil || IsNil(o.MailgunDomain) {
		return nil, false
	}
	return o.MailgunDomain, true
}

// HasMailgunDomain returns a boolean if a field has been set.
func (o *CompanySettings) HasMailgunDomain() bool {
	if o != nil && !IsNil(o.MailgunDomain) {
		return true
	}

	return false
}

// SetMailgunDomain gets a reference to the given string and assigns it to the MailgunDomain field.
func (o *CompanySettings) SetMailgunDomain(v string) {
	o.MailgunDomain = &v
}

// GetSendEmailOnMarkPaid returns the SendEmailOnMarkPaid field value if set, zero value otherwise.
func (o *CompanySettings) GetSendEmailOnMarkPaid() bool {
	if o == nil || IsNil(o.SendEmailOnMarkPaid) {
		var ret bool
		return ret
	}
	return *o.SendEmailOnMarkPaid
}

// GetSendEmailOnMarkPaidOk returns a tuple with the SendEmailOnMarkPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSendEmailOnMarkPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.SendEmailOnMarkPaid) {
		return nil, false
	}
	return o.SendEmailOnMarkPaid, true
}

// HasSendEmailOnMarkPaid returns a boolean if a field has been set.
func (o *CompanySettings) HasSendEmailOnMarkPaid() bool {
	if o != nil && !IsNil(o.SendEmailOnMarkPaid) {
		return true
	}

	return false
}

// SetSendEmailOnMarkPaid gets a reference to the given bool and assigns it to the SendEmailOnMarkPaid field.
func (o *CompanySettings) SetSendEmailOnMarkPaid(v bool) {
	o.SendEmailOnMarkPaid = &v
}

// GetVendorPortalEnableUploads returns the VendorPortalEnableUploads field value if set, zero value otherwise.
func (o *CompanySettings) GetVendorPortalEnableUploads() bool {
	if o == nil || IsNil(o.VendorPortalEnableUploads) {
		var ret bool
		return ret
	}
	return *o.VendorPortalEnableUploads
}

// GetVendorPortalEnableUploadsOk returns a tuple with the VendorPortalEnableUploads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetVendorPortalEnableUploadsOk() (*bool, bool) {
	if o == nil || IsNil(o.VendorPortalEnableUploads) {
		return nil, false
	}
	return o.VendorPortalEnableUploads, true
}

// HasVendorPortalEnableUploads returns a boolean if a field has been set.
func (o *CompanySettings) HasVendorPortalEnableUploads() bool {
	if o != nil && !IsNil(o.VendorPortalEnableUploads) {
		return true
	}

	return false
}

// SetVendorPortalEnableUploads gets a reference to the given bool and assigns it to the VendorPortalEnableUploads field.
func (o *CompanySettings) SetVendorPortalEnableUploads(v bool) {
	o.VendorPortalEnableUploads = &v
}

// GetBesrId returns the BesrId field value if set, zero value otherwise.
func (o *CompanySettings) GetBesrId() string {
	if o == nil || IsNil(o.BesrId) {
		var ret string
		return ret
	}
	return *o.BesrId
}

// GetBesrIdOk returns a tuple with the BesrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetBesrIdOk() (*string, bool) {
	if o == nil || IsNil(o.BesrId) {
		return nil, false
	}
	return o.BesrId, true
}

// HasBesrId returns a boolean if a field has been set.
func (o *CompanySettings) HasBesrId() bool {
	if o != nil && !IsNil(o.BesrId) {
		return true
	}

	return false
}

// SetBesrId gets a reference to the given string and assigns it to the BesrId field.
func (o *CompanySettings) SetBesrId(v string) {
	o.BesrId = &v
}

// GetQrIban returns the QrIban field value if set, zero value otherwise.
func (o *CompanySettings) GetQrIban() string {
	if o == nil || IsNil(o.QrIban) {
		var ret string
		return ret
	}
	return *o.QrIban
}

// GetQrIbanOk returns a tuple with the QrIban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetQrIbanOk() (*string, bool) {
	if o == nil || IsNil(o.QrIban) {
		return nil, false
	}
	return o.QrIban, true
}

// HasQrIban returns a boolean if a field has been set.
func (o *CompanySettings) HasQrIban() bool {
	if o != nil && !IsNil(o.QrIban) {
		return true
	}

	return false
}

// SetQrIban gets a reference to the given string and assigns it to the QrIban field.
func (o *CompanySettings) SetQrIban(v string) {
	o.QrIban = &v
}

// GetEmailSubjectPurchaseOrder returns the EmailSubjectPurchaseOrder field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSubjectPurchaseOrder() string {
	if o == nil || IsNil(o.EmailSubjectPurchaseOrder) {
		var ret string
		return ret
	}
	return *o.EmailSubjectPurchaseOrder
}

// GetEmailSubjectPurchaseOrderOk returns a tuple with the EmailSubjectPurchaseOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSubjectPurchaseOrderOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubjectPurchaseOrder) {
		return nil, false
	}
	return o.EmailSubjectPurchaseOrder, true
}

// HasEmailSubjectPurchaseOrder returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSubjectPurchaseOrder() bool {
	if o != nil && !IsNil(o.EmailSubjectPurchaseOrder) {
		return true
	}

	return false
}

// SetEmailSubjectPurchaseOrder gets a reference to the given string and assigns it to the EmailSubjectPurchaseOrder field.
func (o *CompanySettings) SetEmailSubjectPurchaseOrder(v string) {
	o.EmailSubjectPurchaseOrder = &v
}

// GetEmailTemplatePurchaseOrder returns the EmailTemplatePurchaseOrder field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailTemplatePurchaseOrder() string {
	if o == nil || IsNil(o.EmailTemplatePurchaseOrder) {
		var ret string
		return ret
	}
	return *o.EmailTemplatePurchaseOrder
}

// GetEmailTemplatePurchaseOrderOk returns a tuple with the EmailTemplatePurchaseOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailTemplatePurchaseOrderOk() (*string, bool) {
	if o == nil || IsNil(o.EmailTemplatePurchaseOrder) {
		return nil, false
	}
	return o.EmailTemplatePurchaseOrder, true
}

// HasEmailTemplatePurchaseOrder returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailTemplatePurchaseOrder() bool {
	if o != nil && !IsNil(o.EmailTemplatePurchaseOrder) {
		return true
	}

	return false
}

// SetEmailTemplatePurchaseOrder gets a reference to the given string and assigns it to the EmailTemplatePurchaseOrder field.
func (o *CompanySettings) SetEmailTemplatePurchaseOrder(v string) {
	o.EmailTemplatePurchaseOrder = &v
}

// GetRequirePurchaseOrderSignature returns the RequirePurchaseOrderSignature field value if set, zero value otherwise.
func (o *CompanySettings) GetRequirePurchaseOrderSignature() bool {
	if o == nil || IsNil(o.RequirePurchaseOrderSignature) {
		var ret bool
		return ret
	}
	return *o.RequirePurchaseOrderSignature
}

// GetRequirePurchaseOrderSignatureOk returns a tuple with the RequirePurchaseOrderSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetRequirePurchaseOrderSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirePurchaseOrderSignature) {
		return nil, false
	}
	return o.RequirePurchaseOrderSignature, true
}

// HasRequirePurchaseOrderSignature returns a boolean if a field has been set.
func (o *CompanySettings) HasRequirePurchaseOrderSignature() bool {
	if o != nil && !IsNil(o.RequirePurchaseOrderSignature) {
		return true
	}

	return false
}

// SetRequirePurchaseOrderSignature gets a reference to the given bool and assigns it to the RequirePurchaseOrderSignature field.
func (o *CompanySettings) SetRequirePurchaseOrderSignature(v bool) {
	o.RequirePurchaseOrderSignature = &v
}

// GetPurchaseOrderPublicNotes returns the PurchaseOrderPublicNotes field value if set, zero value otherwise.
func (o *CompanySettings) GetPurchaseOrderPublicNotes() string {
	if o == nil || IsNil(o.PurchaseOrderPublicNotes) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderPublicNotes
}

// GetPurchaseOrderPublicNotesOk returns a tuple with the PurchaseOrderPublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPurchaseOrderPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderPublicNotes) {
		return nil, false
	}
	return o.PurchaseOrderPublicNotes, true
}

// HasPurchaseOrderPublicNotes returns a boolean if a field has been set.
func (o *CompanySettings) HasPurchaseOrderPublicNotes() bool {
	if o != nil && !IsNil(o.PurchaseOrderPublicNotes) {
		return true
	}

	return false
}

// SetPurchaseOrderPublicNotes gets a reference to the given string and assigns it to the PurchaseOrderPublicNotes field.
func (o *CompanySettings) SetPurchaseOrderPublicNotes(v string) {
	o.PurchaseOrderPublicNotes = &v
}

// GetPurchaseOrderTerms returns the PurchaseOrderTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetPurchaseOrderTerms() string {
	if o == nil || IsNil(o.PurchaseOrderTerms) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderTerms
}

// GetPurchaseOrderTermsOk returns a tuple with the PurchaseOrderTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPurchaseOrderTermsOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderTerms) {
		return nil, false
	}
	return o.PurchaseOrderTerms, true
}

// HasPurchaseOrderTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasPurchaseOrderTerms() bool {
	if o != nil && !IsNil(o.PurchaseOrderTerms) {
		return true
	}

	return false
}

// SetPurchaseOrderTerms gets a reference to the given string and assigns it to the PurchaseOrderTerms field.
func (o *CompanySettings) SetPurchaseOrderTerms(v string) {
	o.PurchaseOrderTerms = &v
}

// GetPurchaseOrderFooter returns the PurchaseOrderFooter field value if set, zero value otherwise.
func (o *CompanySettings) GetPurchaseOrderFooter() string {
	if o == nil || IsNil(o.PurchaseOrderFooter) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderFooter
}

// GetPurchaseOrderFooterOk returns a tuple with the PurchaseOrderFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPurchaseOrderFooterOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderFooter) {
		return nil, false
	}
	return o.PurchaseOrderFooter, true
}

// HasPurchaseOrderFooter returns a boolean if a field has been set.
func (o *CompanySettings) HasPurchaseOrderFooter() bool {
	if o != nil && !IsNil(o.PurchaseOrderFooter) {
		return true
	}

	return false
}

// SetPurchaseOrderFooter gets a reference to the given string and assigns it to the PurchaseOrderFooter field.
func (o *CompanySettings) SetPurchaseOrderFooter(v string) {
	o.PurchaseOrderFooter = &v
}

// GetPurchaseOrderDesignId returns the PurchaseOrderDesignId field value if set, zero value otherwise.
func (o *CompanySettings) GetPurchaseOrderDesignId() string {
	if o == nil || IsNil(o.PurchaseOrderDesignId) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderDesignId
}

// GetPurchaseOrderDesignIdOk returns a tuple with the PurchaseOrderDesignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPurchaseOrderDesignIdOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderDesignId) {
		return nil, false
	}
	return o.PurchaseOrderDesignId, true
}

// HasPurchaseOrderDesignId returns a boolean if a field has been set.
func (o *CompanySettings) HasPurchaseOrderDesignId() bool {
	if o != nil && !IsNil(o.PurchaseOrderDesignId) {
		return true
	}

	return false
}

// SetPurchaseOrderDesignId gets a reference to the given string and assigns it to the PurchaseOrderDesignId field.
func (o *CompanySettings) SetPurchaseOrderDesignId(v string) {
	o.PurchaseOrderDesignId = &v
}

// GetPurchaseOrderNumberPattern returns the PurchaseOrderNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetPurchaseOrderNumberPattern() string {
	if o == nil || IsNil(o.PurchaseOrderNumberPattern) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumberPattern
}

// GetPurchaseOrderNumberPatternOk returns a tuple with the PurchaseOrderNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPurchaseOrderNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderNumberPattern) {
		return nil, false
	}
	return o.PurchaseOrderNumberPattern, true
}

// HasPurchaseOrderNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasPurchaseOrderNumberPattern() bool {
	if o != nil && !IsNil(o.PurchaseOrderNumberPattern) {
		return true
	}

	return false
}

// SetPurchaseOrderNumberPattern gets a reference to the given string and assigns it to the PurchaseOrderNumberPattern field.
func (o *CompanySettings) SetPurchaseOrderNumberPattern(v string) {
	o.PurchaseOrderNumberPattern = &v
}

// GetPurchaseOrderNumberCounter returns the PurchaseOrderNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetPurchaseOrderNumberCounter() int32 {
	if o == nil || IsNil(o.PurchaseOrderNumberCounter) {
		var ret int32
		return ret
	}
	return *o.PurchaseOrderNumberCounter
}

// GetPurchaseOrderNumberCounterOk returns a tuple with the PurchaseOrderNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPurchaseOrderNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.PurchaseOrderNumberCounter) {
		return nil, false
	}
	return o.PurchaseOrderNumberCounter, true
}

// HasPurchaseOrderNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasPurchaseOrderNumberCounter() bool {
	if o != nil && !IsNil(o.PurchaseOrderNumberCounter) {
		return true
	}

	return false
}

// SetPurchaseOrderNumberCounter gets a reference to the given int32 and assigns it to the PurchaseOrderNumberCounter field.
func (o *CompanySettings) SetPurchaseOrderNumberCounter(v int32) {
	o.PurchaseOrderNumberCounter = &v
}

// GetPageNumberingAlignment returns the PageNumberingAlignment field value if set, zero value otherwise.
func (o *CompanySettings) GetPageNumberingAlignment() string {
	if o == nil || IsNil(o.PageNumberingAlignment) {
		var ret string
		return ret
	}
	return *o.PageNumberingAlignment
}

// GetPageNumberingAlignmentOk returns a tuple with the PageNumberingAlignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPageNumberingAlignmentOk() (*string, bool) {
	if o == nil || IsNil(o.PageNumberingAlignment) {
		return nil, false
	}
	return o.PageNumberingAlignment, true
}

// HasPageNumberingAlignment returns a boolean if a field has been set.
func (o *CompanySettings) HasPageNumberingAlignment() bool {
	if o != nil && !IsNil(o.PageNumberingAlignment) {
		return true
	}

	return false
}

// SetPageNumberingAlignment gets a reference to the given string and assigns it to the PageNumberingAlignment field.
func (o *CompanySettings) SetPageNumberingAlignment(v string) {
	o.PageNumberingAlignment = &v
}

// GetPageNumbering returns the PageNumbering field value if set, zero value otherwise.
func (o *CompanySettings) GetPageNumbering() bool {
	if o == nil || IsNil(o.PageNumbering) {
		var ret bool
		return ret
	}
	return *o.PageNumbering
}

// GetPageNumberingOk returns a tuple with the PageNumbering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPageNumberingOk() (*bool, bool) {
	if o == nil || IsNil(o.PageNumbering) {
		return nil, false
	}
	return o.PageNumbering, true
}

// HasPageNumbering returns a boolean if a field has been set.
func (o *CompanySettings) HasPageNumbering() bool {
	if o != nil && !IsNil(o.PageNumbering) {
		return true
	}

	return false
}

// SetPageNumbering gets a reference to the given bool and assigns it to the PageNumbering field.
func (o *CompanySettings) SetPageNumbering(v bool) {
	o.PageNumbering = &v
}

// GetAutoArchiveInvoiceCancelled returns the AutoArchiveInvoiceCancelled field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoArchiveInvoiceCancelled() bool {
	if o == nil || IsNil(o.AutoArchiveInvoiceCancelled) {
		var ret bool
		return ret
	}
	return *o.AutoArchiveInvoiceCancelled
}

// GetAutoArchiveInvoiceCancelledOk returns a tuple with the AutoArchiveInvoiceCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoArchiveInvoiceCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoArchiveInvoiceCancelled) {
		return nil, false
	}
	return o.AutoArchiveInvoiceCancelled, true
}

// HasAutoArchiveInvoiceCancelled returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoArchiveInvoiceCancelled() bool {
	if o != nil && !IsNil(o.AutoArchiveInvoiceCancelled) {
		return true
	}

	return false
}

// SetAutoArchiveInvoiceCancelled gets a reference to the given bool and assigns it to the AutoArchiveInvoiceCancelled field.
func (o *CompanySettings) SetAutoArchiveInvoiceCancelled(v bool) {
	o.AutoArchiveInvoiceCancelled = &v
}

// GetEmailFromName returns the EmailFromName field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailFromName() string {
	if o == nil || IsNil(o.EmailFromName) {
		var ret string
		return ret
	}
	return *o.EmailFromName
}

// GetEmailFromNameOk returns a tuple with the EmailFromName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailFromNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmailFromName) {
		return nil, false
	}
	return o.EmailFromName, true
}

// HasEmailFromName returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailFromName() bool {
	if o != nil && !IsNil(o.EmailFromName) {
		return true
	}

	return false
}

// SetEmailFromName gets a reference to the given string and assigns it to the EmailFromName field.
func (o *CompanySettings) SetEmailFromName(v string) {
	o.EmailFromName = &v
}

// GetShowAllTasksClientPortal returns the ShowAllTasksClientPortal field value if set, zero value otherwise.
func (o *CompanySettings) GetShowAllTasksClientPortal() bool {
	if o == nil || IsNil(o.ShowAllTasksClientPortal) {
		var ret bool
		return ret
	}
	return *o.ShowAllTasksClientPortal
}

// GetShowAllTasksClientPortalOk returns a tuple with the ShowAllTasksClientPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetShowAllTasksClientPortalOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowAllTasksClientPortal) {
		return nil, false
	}
	return o.ShowAllTasksClientPortal, true
}

// HasShowAllTasksClientPortal returns a boolean if a field has been set.
func (o *CompanySettings) HasShowAllTasksClientPortal() bool {
	if o != nil && !IsNil(o.ShowAllTasksClientPortal) {
		return true
	}

	return false
}

// SetShowAllTasksClientPortal gets a reference to the given bool and assigns it to the ShowAllTasksClientPortal field.
func (o *CompanySettings) SetShowAllTasksClientPortal(v bool) {
	o.ShowAllTasksClientPortal = &v
}

// GetEntitySendTime returns the EntitySendTime field value if set, zero value otherwise.
func (o *CompanySettings) GetEntitySendTime() int32 {
	if o == nil || IsNil(o.EntitySendTime) {
		var ret int32
		return ret
	}
	return *o.EntitySendTime
}

// GetEntitySendTimeOk returns a tuple with the EntitySendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEntitySendTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EntitySendTime) {
		return nil, false
	}
	return o.EntitySendTime, true
}

// HasEntitySendTime returns a boolean if a field has been set.
func (o *CompanySettings) HasEntitySendTime() bool {
	if o != nil && !IsNil(o.EntitySendTime) {
		return true
	}

	return false
}

// SetEntitySendTime gets a reference to the given int32 and assigns it to the EntitySendTime field.
func (o *CompanySettings) SetEntitySendTime(v int32) {
	o.EntitySendTime = &v
}

// GetSharedInvoiceCreditCounter returns the SharedInvoiceCreditCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetSharedInvoiceCreditCounter() bool {
	if o == nil || IsNil(o.SharedInvoiceCreditCounter) {
		var ret bool
		return ret
	}
	return *o.SharedInvoiceCreditCounter
}

// GetSharedInvoiceCreditCounterOk returns a tuple with the SharedInvoiceCreditCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSharedInvoiceCreditCounterOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedInvoiceCreditCounter) {
		return nil, false
	}
	return o.SharedInvoiceCreditCounter, true
}

// HasSharedInvoiceCreditCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasSharedInvoiceCreditCounter() bool {
	if o != nil && !IsNil(o.SharedInvoiceCreditCounter) {
		return true
	}

	return false
}

// SetSharedInvoiceCreditCounter gets a reference to the given bool and assigns it to the SharedInvoiceCreditCounter field.
func (o *CompanySettings) SetSharedInvoiceCreditCounter(v bool) {
	o.SharedInvoiceCreditCounter = &v
}

// GetReplyToName returns the ReplyToName field value if set, zero value otherwise.
func (o *CompanySettings) GetReplyToName() string {
	if o == nil || IsNil(o.ReplyToName) {
		var ret string
		return ret
	}
	return *o.ReplyToName
}

// GetReplyToNameOk returns a tuple with the ReplyToName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetReplyToNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyToName) {
		return nil, false
	}
	return o.ReplyToName, true
}

// HasReplyToName returns a boolean if a field has been set.
func (o *CompanySettings) HasReplyToName() bool {
	if o != nil && !IsNil(o.ReplyToName) {
		return true
	}

	return false
}

// SetReplyToName gets a reference to the given string and assigns it to the ReplyToName field.
func (o *CompanySettings) SetReplyToName(v string) {
	o.ReplyToName = &v
}

// GetHideEmptyColumnsOnPdf returns the HideEmptyColumnsOnPdf field value if set, zero value otherwise.
func (o *CompanySettings) GetHideEmptyColumnsOnPdf() bool {
	if o == nil || IsNil(o.HideEmptyColumnsOnPdf) {
		var ret bool
		return ret
	}
	return *o.HideEmptyColumnsOnPdf
}

// GetHideEmptyColumnsOnPdfOk returns a tuple with the HideEmptyColumnsOnPdf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetHideEmptyColumnsOnPdfOk() (*bool, bool) {
	if o == nil || IsNil(o.HideEmptyColumnsOnPdf) {
		return nil, false
	}
	return o.HideEmptyColumnsOnPdf, true
}

// HasHideEmptyColumnsOnPdf returns a boolean if a field has been set.
func (o *CompanySettings) HasHideEmptyColumnsOnPdf() bool {
	if o != nil && !IsNil(o.HideEmptyColumnsOnPdf) {
		return true
	}

	return false
}

// SetHideEmptyColumnsOnPdf gets a reference to the given bool and assigns it to the HideEmptyColumnsOnPdf field.
func (o *CompanySettings) SetHideEmptyColumnsOnPdf(v bool) {
	o.HideEmptyColumnsOnPdf = &v
}

// GetEnableReminderEndless returns the EnableReminderEndless field value if set, zero value otherwise.
func (o *CompanySettings) GetEnableReminderEndless() bool {
	if o == nil || IsNil(o.EnableReminderEndless) {
		var ret bool
		return ret
	}
	return *o.EnableReminderEndless
}

// GetEnableReminderEndlessOk returns a tuple with the EnableReminderEndless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEnableReminderEndlessOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableReminderEndless) {
		return nil, false
	}
	return o.EnableReminderEndless, true
}

// HasEnableReminderEndless returns a boolean if a field has been set.
func (o *CompanySettings) HasEnableReminderEndless() bool {
	if o != nil && !IsNil(o.EnableReminderEndless) {
		return true
	}

	return false
}

// SetEnableReminderEndless gets a reference to the given bool and assigns it to the EnableReminderEndless field.
func (o *CompanySettings) SetEnableReminderEndless(v bool) {
	o.EnableReminderEndless = &v
}

// GetUseCreditsPayment returns the UseCreditsPayment field value if set, zero value otherwise.
func (o *CompanySettings) GetUseCreditsPayment() bool {
	if o == nil || IsNil(o.UseCreditsPayment) {
		var ret bool
		return ret
	}
	return *o.UseCreditsPayment
}

// GetUseCreditsPaymentOk returns a tuple with the UseCreditsPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetUseCreditsPaymentOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCreditsPayment) {
		return nil, false
	}
	return o.UseCreditsPayment, true
}

// HasUseCreditsPayment returns a boolean if a field has been set.
func (o *CompanySettings) HasUseCreditsPayment() bool {
	if o != nil && !IsNil(o.UseCreditsPayment) {
		return true
	}

	return false
}

// SetUseCreditsPayment gets a reference to the given bool and assigns it to the UseCreditsPayment field.
func (o *CompanySettings) SetUseCreditsPayment(v bool) {
	o.UseCreditsPayment = &v
}

// GetRecurringInvoiceNumberPattern returns the RecurringInvoiceNumberPattern field value if set, zero value otherwise.
func (o *CompanySettings) GetRecurringInvoiceNumberPattern() string {
	if o == nil || IsNil(o.RecurringInvoiceNumberPattern) {
		var ret string
		return ret
	}
	return *o.RecurringInvoiceNumberPattern
}

// GetRecurringInvoiceNumberPatternOk returns a tuple with the RecurringInvoiceNumberPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetRecurringInvoiceNumberPatternOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringInvoiceNumberPattern) {
		return nil, false
	}
	return o.RecurringInvoiceNumberPattern, true
}

// HasRecurringInvoiceNumberPattern returns a boolean if a field has been set.
func (o *CompanySettings) HasRecurringInvoiceNumberPattern() bool {
	if o != nil && !IsNil(o.RecurringInvoiceNumberPattern) {
		return true
	}

	return false
}

// SetRecurringInvoiceNumberPattern gets a reference to the given string and assigns it to the RecurringInvoiceNumberPattern field.
func (o *CompanySettings) SetRecurringInvoiceNumberPattern(v string) {
	o.RecurringInvoiceNumberPattern = &v
}

// GetRecurringInvoiceNumberCounter returns the RecurringInvoiceNumberCounter field value if set, zero value otherwise.
func (o *CompanySettings) GetRecurringInvoiceNumberCounter() int32 {
	if o == nil || IsNil(o.RecurringInvoiceNumberCounter) {
		var ret int32
		return ret
	}
	return *o.RecurringInvoiceNumberCounter
}

// GetRecurringInvoiceNumberCounterOk returns a tuple with the RecurringInvoiceNumberCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetRecurringInvoiceNumberCounterOk() (*int32, bool) {
	if o == nil || IsNil(o.RecurringInvoiceNumberCounter) {
		return nil, false
	}
	return o.RecurringInvoiceNumberCounter, true
}

// HasRecurringInvoiceNumberCounter returns a boolean if a field has been set.
func (o *CompanySettings) HasRecurringInvoiceNumberCounter() bool {
	if o != nil && !IsNil(o.RecurringInvoiceNumberCounter) {
		return true
	}

	return false
}

// SetRecurringInvoiceNumberCounter gets a reference to the given int32 and assigns it to the RecurringInvoiceNumberCounter field.
func (o *CompanySettings) SetRecurringInvoiceNumberCounter(v int32) {
	o.RecurringInvoiceNumberCounter = &v
}

// GetClientPortalUnderPaymentMinimum returns the ClientPortalUnderPaymentMinimum field value if set, zero value otherwise.
func (o *CompanySettings) GetClientPortalUnderPaymentMinimum() float32 {
	if o == nil || IsNil(o.ClientPortalUnderPaymentMinimum) {
		var ret float32
		return ret
	}
	return *o.ClientPortalUnderPaymentMinimum
}

// GetClientPortalUnderPaymentMinimumOk returns a tuple with the ClientPortalUnderPaymentMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientPortalUnderPaymentMinimumOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientPortalUnderPaymentMinimum) {
		return nil, false
	}
	return o.ClientPortalUnderPaymentMinimum, true
}

// HasClientPortalUnderPaymentMinimum returns a boolean if a field has been set.
func (o *CompanySettings) HasClientPortalUnderPaymentMinimum() bool {
	if o != nil && !IsNil(o.ClientPortalUnderPaymentMinimum) {
		return true
	}

	return false
}

// SetClientPortalUnderPaymentMinimum gets a reference to the given float32 and assigns it to the ClientPortalUnderPaymentMinimum field.
func (o *CompanySettings) SetClientPortalUnderPaymentMinimum(v float32) {
	o.ClientPortalUnderPaymentMinimum = &v
}

// GetAutoBillDate returns the AutoBillDate field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoBillDate() string {
	if o == nil || IsNil(o.AutoBillDate) {
		var ret string
		return ret
	}
	return *o.AutoBillDate
}

// GetAutoBillDateOk returns a tuple with the AutoBillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoBillDateOk() (*string, bool) {
	if o == nil || IsNil(o.AutoBillDate) {
		return nil, false
	}
	return o.AutoBillDate, true
}

// HasAutoBillDate returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoBillDate() bool {
	if o != nil && !IsNil(o.AutoBillDate) {
		return true
	}

	return false
}

// SetAutoBillDate gets a reference to the given string and assigns it to the AutoBillDate field.
func (o *CompanySettings) SetAutoBillDate(v string) {
	o.AutoBillDate = &v
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise.
func (o *CompanySettings) GetPrimaryColor() string {
	if o == nil || IsNil(o.PrimaryColor) {
		var ret string
		return ret
	}
	return *o.PrimaryColor
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPrimaryColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryColor) {
		return nil, false
	}
	return o.PrimaryColor, true
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *CompanySettings) HasPrimaryColor() bool {
	if o != nil && !IsNil(o.PrimaryColor) {
		return true
	}

	return false
}

// SetPrimaryColor gets a reference to the given string and assigns it to the PrimaryColor field.
func (o *CompanySettings) SetPrimaryColor(v string) {
	o.PrimaryColor = &v
}

// GetSecondaryColor returns the SecondaryColor field value if set, zero value otherwise.
func (o *CompanySettings) GetSecondaryColor() string {
	if o == nil || IsNil(o.SecondaryColor) {
		var ret string
		return ret
	}
	return *o.SecondaryColor
}

// GetSecondaryColorOk returns a tuple with the SecondaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetSecondaryColorOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryColor) {
		return nil, false
	}
	return o.SecondaryColor, true
}

// HasSecondaryColor returns a boolean if a field has been set.
func (o *CompanySettings) HasSecondaryColor() bool {
	if o != nil && !IsNil(o.SecondaryColor) {
		return true
	}

	return false
}

// SetSecondaryColor gets a reference to the given string and assigns it to the SecondaryColor field.
func (o *CompanySettings) SetSecondaryColor(v string) {
	o.SecondaryColor = &v
}

// GetClientPortalAllowUnderPayment returns the ClientPortalAllowUnderPayment field value if set, zero value otherwise.
func (o *CompanySettings) GetClientPortalAllowUnderPayment() bool {
	if o == nil || IsNil(o.ClientPortalAllowUnderPayment) {
		var ret bool
		return ret
	}
	return *o.ClientPortalAllowUnderPayment
}

// GetClientPortalAllowUnderPaymentOk returns a tuple with the ClientPortalAllowUnderPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientPortalAllowUnderPaymentOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientPortalAllowUnderPayment) {
		return nil, false
	}
	return o.ClientPortalAllowUnderPayment, true
}

// HasClientPortalAllowUnderPayment returns a boolean if a field has been set.
func (o *CompanySettings) HasClientPortalAllowUnderPayment() bool {
	if o != nil && !IsNil(o.ClientPortalAllowUnderPayment) {
		return true
	}

	return false
}

// SetClientPortalAllowUnderPayment gets a reference to the given bool and assigns it to the ClientPortalAllowUnderPayment field.
func (o *CompanySettings) SetClientPortalAllowUnderPayment(v bool) {
	o.ClientPortalAllowUnderPayment = &v
}

// GetClientPortalAllowOverPayment returns the ClientPortalAllowOverPayment field value if set, zero value otherwise.
func (o *CompanySettings) GetClientPortalAllowOverPayment() bool {
	if o == nil || IsNil(o.ClientPortalAllowOverPayment) {
		var ret bool
		return ret
	}
	return *o.ClientPortalAllowOverPayment
}

// GetClientPortalAllowOverPaymentOk returns a tuple with the ClientPortalAllowOverPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientPortalAllowOverPaymentOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientPortalAllowOverPayment) {
		return nil, false
	}
	return o.ClientPortalAllowOverPayment, true
}

// HasClientPortalAllowOverPayment returns a boolean if a field has been set.
func (o *CompanySettings) HasClientPortalAllowOverPayment() bool {
	if o != nil && !IsNil(o.ClientPortalAllowOverPayment) {
		return true
	}

	return false
}

// SetClientPortalAllowOverPayment gets a reference to the given bool and assigns it to the ClientPortalAllowOverPayment field.
func (o *CompanySettings) SetClientPortalAllowOverPayment(v bool) {
	o.ClientPortalAllowOverPayment = &v
}

// GetAutoBill returns the AutoBill field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoBill() string {
	if o == nil || IsNil(o.AutoBill) {
		var ret string
		return ret
	}
	return *o.AutoBill
}

// GetAutoBillOk returns a tuple with the AutoBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoBillOk() (*string, bool) {
	if o == nil || IsNil(o.AutoBill) {
		return nil, false
	}
	return o.AutoBill, true
}

// HasAutoBill returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoBill() bool {
	if o != nil && !IsNil(o.AutoBill) {
		return true
	}

	return false
}

// SetAutoBill gets a reference to the given string and assigns it to the AutoBill field.
func (o *CompanySettings) SetAutoBill(v string) {
	o.AutoBill = &v
}

// GetClientPortalTerms returns the ClientPortalTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetClientPortalTerms() string {
	if o == nil || IsNil(o.ClientPortalTerms) {
		var ret string
		return ret
	}
	return *o.ClientPortalTerms
}

// GetClientPortalTermsOk returns a tuple with the ClientPortalTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientPortalTermsOk() (*string, bool) {
	if o == nil || IsNil(o.ClientPortalTerms) {
		return nil, false
	}
	return o.ClientPortalTerms, true
}

// HasClientPortalTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasClientPortalTerms() bool {
	if o != nil && !IsNil(o.ClientPortalTerms) {
		return true
	}

	return false
}

// SetClientPortalTerms gets a reference to the given string and assigns it to the ClientPortalTerms field.
func (o *CompanySettings) SetClientPortalTerms(v string) {
	o.ClientPortalTerms = &v
}

// GetClientPortalPrivacyPolicy returns the ClientPortalPrivacyPolicy field value if set, zero value otherwise.
func (o *CompanySettings) GetClientPortalPrivacyPolicy() string {
	if o == nil || IsNil(o.ClientPortalPrivacyPolicy) {
		var ret string
		return ret
	}
	return *o.ClientPortalPrivacyPolicy
}

// GetClientPortalPrivacyPolicyOk returns a tuple with the ClientPortalPrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientPortalPrivacyPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.ClientPortalPrivacyPolicy) {
		return nil, false
	}
	return o.ClientPortalPrivacyPolicy, true
}

// HasClientPortalPrivacyPolicy returns a boolean if a field has been set.
func (o *CompanySettings) HasClientPortalPrivacyPolicy() bool {
	if o != nil && !IsNil(o.ClientPortalPrivacyPolicy) {
		return true
	}

	return false
}

// SetClientPortalPrivacyPolicy gets a reference to the given string and assigns it to the ClientPortalPrivacyPolicy field.
func (o *CompanySettings) SetClientPortalPrivacyPolicy(v string) {
	o.ClientPortalPrivacyPolicy = &v
}

// GetClientCanRegister returns the ClientCanRegister field value if set, zero value otherwise.
func (o *CompanySettings) GetClientCanRegister() bool {
	if o == nil || IsNil(o.ClientCanRegister) {
		var ret bool
		return ret
	}
	return *o.ClientCanRegister
}

// GetClientCanRegisterOk returns a tuple with the ClientCanRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClientCanRegisterOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientCanRegister) {
		return nil, false
	}
	return o.ClientCanRegister, true
}

// HasClientCanRegister returns a boolean if a field has been set.
func (o *CompanySettings) HasClientCanRegister() bool {
	if o != nil && !IsNil(o.ClientCanRegister) {
		return true
	}

	return false
}

// SetClientCanRegister gets a reference to the given bool and assigns it to the ClientCanRegister field.
func (o *CompanySettings) SetClientCanRegister(v bool) {
	o.ClientCanRegister = &v
}

// GetPortalDesignId returns the PortalDesignId field value if set, zero value otherwise.
func (o *CompanySettings) GetPortalDesignId() string {
	if o == nil || IsNil(o.PortalDesignId) {
		var ret string
		return ret
	}
	return *o.PortalDesignId
}

// GetPortalDesignIdOk returns a tuple with the PortalDesignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetPortalDesignIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortalDesignId) {
		return nil, false
	}
	return o.PortalDesignId, true
}

// HasPortalDesignId returns a boolean if a field has been set.
func (o *CompanySettings) HasPortalDesignId() bool {
	if o != nil && !IsNil(o.PortalDesignId) {
		return true
	}

	return false
}

// SetPortalDesignId gets a reference to the given string and assigns it to the PortalDesignId field.
func (o *CompanySettings) SetPortalDesignId(v string) {
	o.PortalDesignId = &v
}

// GetLateFeeEndlessPercent returns the LateFeeEndlessPercent field value if set, zero value otherwise.
func (o *CompanySettings) GetLateFeeEndlessPercent() float32 {
	if o == nil || IsNil(o.LateFeeEndlessPercent) {
		var ret float32
		return ret
	}
	return *o.LateFeeEndlessPercent
}

// GetLateFeeEndlessPercentOk returns a tuple with the LateFeeEndlessPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetLateFeeEndlessPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.LateFeeEndlessPercent) {
		return nil, false
	}
	return o.LateFeeEndlessPercent, true
}

// HasLateFeeEndlessPercent returns a boolean if a field has been set.
func (o *CompanySettings) HasLateFeeEndlessPercent() bool {
	if o != nil && !IsNil(o.LateFeeEndlessPercent) {
		return true
	}

	return false
}

// SetLateFeeEndlessPercent gets a reference to the given float32 and assigns it to the LateFeeEndlessPercent field.
func (o *CompanySettings) SetLateFeeEndlessPercent(v float32) {
	o.LateFeeEndlessPercent = &v
}

// GetLateFeeEndlessAmount returns the LateFeeEndlessAmount field value if set, zero value otherwise.
func (o *CompanySettings) GetLateFeeEndlessAmount() float32 {
	if o == nil || IsNil(o.LateFeeEndlessAmount) {
		var ret float32
		return ret
	}
	return *o.LateFeeEndlessAmount
}

// GetLateFeeEndlessAmountOk returns a tuple with the LateFeeEndlessAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetLateFeeEndlessAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.LateFeeEndlessAmount) {
		return nil, false
	}
	return o.LateFeeEndlessAmount, true
}

// HasLateFeeEndlessAmount returns a boolean if a field has been set.
func (o *CompanySettings) HasLateFeeEndlessAmount() bool {
	if o != nil && !IsNil(o.LateFeeEndlessAmount) {
		return true
	}

	return false
}

// SetLateFeeEndlessAmount gets a reference to the given float32 and assigns it to the LateFeeEndlessAmount field.
func (o *CompanySettings) SetLateFeeEndlessAmount(v float32) {
	o.LateFeeEndlessAmount = &v
}

// GetAutoEmailInvoice returns the AutoEmailInvoice field value if set, zero value otherwise.
func (o *CompanySettings) GetAutoEmailInvoice() bool {
	if o == nil || IsNil(o.AutoEmailInvoice) {
		var ret bool
		return ret
	}
	return *o.AutoEmailInvoice
}

// GetAutoEmailInvoiceOk returns a tuple with the AutoEmailInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetAutoEmailInvoiceOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoEmailInvoice) {
		return nil, false
	}
	return o.AutoEmailInvoice, true
}

// HasAutoEmailInvoice returns a boolean if a field has been set.
func (o *CompanySettings) HasAutoEmailInvoice() bool {
	if o != nil && !IsNil(o.AutoEmailInvoice) {
		return true
	}

	return false
}

// SetAutoEmailInvoice gets a reference to the given bool and assigns it to the AutoEmailInvoice field.
func (o *CompanySettings) SetAutoEmailInvoice(v bool) {
	o.AutoEmailInvoice = &v
}

// GetEmailSignature returns the EmailSignature field value if set, zero value otherwise.
func (o *CompanySettings) GetEmailSignature() string {
	if o == nil || IsNil(o.EmailSignature) {
		var ret string
		return ret
	}
	return *o.EmailSignature
}

// GetEmailSignatureOk returns a tuple with the EmailSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetEmailSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSignature) {
		return nil, false
	}
	return o.EmailSignature, true
}

// HasEmailSignature returns a boolean if a field has been set.
func (o *CompanySettings) HasEmailSignature() bool {
	if o != nil && !IsNil(o.EmailSignature) {
		return true
	}

	return false
}

// SetEmailSignature gets a reference to the given string and assigns it to the EmailSignature field.
func (o *CompanySettings) SetEmailSignature(v string) {
	o.EmailSignature = &v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *CompanySettings) GetClassification() string {
	if o == nil || IsNil(o.Classification) {
		var ret string
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetClassificationOk() (*string, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *CompanySettings) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given string and assigns it to the Classification field.
func (o *CompanySettings) SetClassification(v string) {
	o.Classification = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *CompanySettings) GetValidUntil() string {
	if o == nil || IsNil(o.ValidUntil) {
		var ret string
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetValidUntilOk() (*string, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *CompanySettings) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given string and assigns it to the ValidUntil field.
func (o *CompanySettings) SetValidUntil(v string) {
	o.ValidUntil = &v
}

// GetCreditTerms returns the CreditTerms field value if set, zero value otherwise.
func (o *CompanySettings) GetCreditTerms() string {
	if o == nil || IsNil(o.CreditTerms) {
		var ret string
		return ret
	}
	return *o.CreditTerms
}

// GetCreditTermsOk returns a tuple with the CreditTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCreditTermsOk() (*string, bool) {
	if o == nil || IsNil(o.CreditTerms) {
		return nil, false
	}
	return o.CreditTerms, true
}

// HasCreditTerms returns a boolean if a field has been set.
func (o *CompanySettings) HasCreditTerms() bool {
	if o != nil && !IsNil(o.CreditTerms) {
		return true
	}

	return false
}

// SetCreditTerms gets a reference to the given string and assigns it to the CreditTerms field.
func (o *CompanySettings) SetCreditTerms(v string) {
	o.CreditTerms = &v
}

// GetCreditFooter returns the CreditFooter field value if set, zero value otherwise.
func (o *CompanySettings) GetCreditFooter() string {
	if o == nil || IsNil(o.CreditFooter) {
		var ret string
		return ret
	}
	return *o.CreditFooter
}

// GetCreditFooterOk returns a tuple with the CreditFooter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanySettings) GetCreditFooterOk() (*string, bool) {
	if o == nil || IsNil(o.CreditFooter) {
		return nil, false
	}
	return o.CreditFooter, true
}

// HasCreditFooter returns a boolean if a field has been set.
func (o *CompanySettings) HasCreditFooter() bool {
	if o != nil && !IsNil(o.CreditFooter) {
		return true
	}

	return false
}

// SetCreditFooter gets a reference to the given string and assigns it to the CreditFooter field.
func (o *CompanySettings) SetCreditFooter(v string) {
	o.CreditFooter = &v
}

func (o CompanySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanySettings) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	toSerialize["currency_id"] = o.CurrencyId
	if !IsNil(o.TimezoneId) {
		toSerialize["timezone_id"] = o.TimezoneId
	}
	if !IsNil(o.DateFormatId) {
		toSerialize["date_format_id"] = o.DateFormatId
	}
	if !IsNil(o.MilitaryTime) {
		toSerialize["military_time"] = o.MilitaryTime
	}
	if !IsNil(o.LanguageId) {
		toSerialize["language_id"] = o.LanguageId
	}
	if !IsNil(o.ShowCurrencyCode) {
		toSerialize["show_currency_code"] = o.ShowCurrencyCode
	}
	if !IsNil(o.PaymentTerms) {
		toSerialize["payment_terms"] = o.PaymentTerms
	}
	if !IsNil(o.CompanyGatewayIds) {
		toSerialize["company_gateway_ids"] = o.CompanyGatewayIds
	}
	if !IsNil(o.CustomValue1) {
		toSerialize["custom_value1"] = o.CustomValue1
	}
	if !IsNil(o.CustomValue2) {
		toSerialize["custom_value2"] = o.CustomValue2
	}
	if !IsNil(o.CustomValue3) {
		toSerialize["custom_value3"] = o.CustomValue3
	}
	if !IsNil(o.CustomValue4) {
		toSerialize["custom_value4"] = o.CustomValue4
	}
	if !IsNil(o.DefaultTaskRate) {
		toSerialize["default_task_rate"] = o.DefaultTaskRate
	}
	if !IsNil(o.SendReminders) {
		toSerialize["send_reminders"] = o.SendReminders
	}
	if !IsNil(o.EnableClientPortalTasks) {
		toSerialize["enable_client_portal_tasks"] = o.EnableClientPortalTasks
	}
	if !IsNil(o.EmailStyle) {
		toSerialize["email_style"] = o.EmailStyle
	}
	if !IsNil(o.ReplyToEmail) {
		toSerialize["reply_to_email"] = o.ReplyToEmail
	}
	if !IsNil(o.BccEmail) {
		toSerialize["bcc_email"] = o.BccEmail
	}
	if !IsNil(o.PdfEmailAttachment) {
		toSerialize["pdf_email_attachment"] = o.PdfEmailAttachment
	}
	if !IsNil(o.UblEmailAttachment) {
		toSerialize["ubl_email_attachment"] = o.UblEmailAttachment
	}
	if !IsNil(o.EmailStyleCustom) {
		toSerialize["email_style_custom"] = o.EmailStyleCustom
	}
	if !IsNil(o.CounterNumberApplied) {
		toSerialize["counter_number_applied"] = o.CounterNumberApplied
	}
	if !IsNil(o.QuoteNumberApplied) {
		toSerialize["quote_number_applied"] = o.QuoteNumberApplied
	}
	if !IsNil(o.CustomMessageDashboard) {
		toSerialize["custom_message_dashboard"] = o.CustomMessageDashboard
	}
	if !IsNil(o.CustomMessageUnpaidInvoice) {
		toSerialize["custom_message_unpaid_invoice"] = o.CustomMessageUnpaidInvoice
	}
	if !IsNil(o.CustomMessagePaidInvoice) {
		toSerialize["custom_message_paid_invoice"] = o.CustomMessagePaidInvoice
	}
	if !IsNil(o.CustomMessageUnapprovedQuote) {
		toSerialize["custom_message_unapproved_quote"] = o.CustomMessageUnapprovedQuote
	}
	if !IsNil(o.LockInvoices) {
		toSerialize["lock_invoices"] = o.LockInvoices
	}
	if !IsNil(o.AutoArchiveInvoice) {
		toSerialize["auto_archive_invoice"] = o.AutoArchiveInvoice
	}
	if !IsNil(o.AutoArchiveQuote) {
		toSerialize["auto_archive_quote"] = o.AutoArchiveQuote
	}
	if !IsNil(o.AutoConvertQuote) {
		toSerialize["auto_convert_quote"] = o.AutoConvertQuote
	}
	if !IsNil(o.InclusiveTaxes) {
		toSerialize["inclusive_taxes"] = o.InclusiveTaxes
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	if !IsNil(o.TaskNumberPattern) {
		toSerialize["task_number_pattern"] = o.TaskNumberPattern
	}
	if !IsNil(o.TaskNumberCounter) {
		toSerialize["task_number_counter"] = o.TaskNumberCounter
	}
	if !IsNil(o.ReminderSendTime) {
		toSerialize["reminder_send_time"] = o.ReminderSendTime
	}
	if !IsNil(o.ExpenseNumberPattern) {
		toSerialize["expense_number_pattern"] = o.ExpenseNumberPattern
	}
	if !IsNil(o.ExpenseNumberCounter) {
		toSerialize["expense_number_counter"] = o.ExpenseNumberCounter
	}
	if !IsNil(o.VendorNumberPattern) {
		toSerialize["vendor_number_pattern"] = o.VendorNumberPattern
	}
	if !IsNil(o.VendorNumberCounter) {
		toSerialize["vendor_number_counter"] = o.VendorNumberCounter
	}
	if !IsNil(o.TicketNumberPattern) {
		toSerialize["ticket_number_pattern"] = o.TicketNumberPattern
	}
	if !IsNil(o.TicketNumberCounter) {
		toSerialize["ticket_number_counter"] = o.TicketNumberCounter
	}
	if !IsNil(o.PaymentNumberPattern) {
		toSerialize["payment_number_pattern"] = o.PaymentNumberPattern
	}
	if !IsNil(o.PaymentNumberCounter) {
		toSerialize["payment_number_counter"] = o.PaymentNumberCounter
	}
	if !IsNil(o.InvoiceNumberPattern) {
		toSerialize["invoice_number_pattern"] = o.InvoiceNumberPattern
	}
	if !IsNil(o.InvoiceNumberCounter) {
		toSerialize["invoice_number_counter"] = o.InvoiceNumberCounter
	}
	if !IsNil(o.QuoteNumberPattern) {
		toSerialize["quote_number_pattern"] = o.QuoteNumberPattern
	}
	if !IsNil(o.QuoteNumberCounter) {
		toSerialize["quote_number_counter"] = o.QuoteNumberCounter
	}
	if !IsNil(o.ClientNumberPattern) {
		toSerialize["client_number_pattern"] = o.ClientNumberPattern
	}
	if !IsNil(o.ClientNumberCounter) {
		toSerialize["client_number_counter"] = o.ClientNumberCounter
	}
	if !IsNil(o.CreditNumberPattern) {
		toSerialize["credit_number_pattern"] = o.CreditNumberPattern
	}
	if !IsNil(o.CreditNumberCounter) {
		toSerialize["credit_number_counter"] = o.CreditNumberCounter
	}
	if !IsNil(o.RecurringInvoiceNumberPrefix) {
		toSerialize["recurring_invoice_number_prefix"] = o.RecurringInvoiceNumberPrefix
	}
	if !IsNil(o.ResetCounterFrequencyId) {
		toSerialize["reset_counter_frequency_id"] = o.ResetCounterFrequencyId
	}
	if !IsNil(o.ResetCounterDate) {
		toSerialize["reset_counter_date"] = o.ResetCounterDate
	}
	if !IsNil(o.CounterPadding) {
		toSerialize["counter_padding"] = o.CounterPadding
	}
	if !IsNil(o.SharedInvoiceQuoteCounter) {
		toSerialize["shared_invoice_quote_counter"] = o.SharedInvoiceQuoteCounter
	}
	if !IsNil(o.UpdateProducts) {
		toSerialize["update_products"] = o.UpdateProducts
	}
	if !IsNil(o.ConvertProducts) {
		toSerialize["convert_products"] = o.ConvertProducts
	}
	if !IsNil(o.FillProducts) {
		toSerialize["fill_products"] = o.FillProducts
	}
	if !IsNil(o.InvoiceTerms) {
		toSerialize["invoice_terms"] = o.InvoiceTerms
	}
	if !IsNil(o.QuoteTerms) {
		toSerialize["quote_terms"] = o.QuoteTerms
	}
	if !IsNil(o.InvoiceTaxes) {
		toSerialize["invoice_taxes"] = o.InvoiceTaxes
	}
	if !IsNil(o.InvoiceDesignId) {
		toSerialize["invoice_design_id"] = o.InvoiceDesignId
	}
	if !IsNil(o.QuoteDesignId) {
		toSerialize["quote_design_id"] = o.QuoteDesignId
	}
	if !IsNil(o.InvoiceFooter) {
		toSerialize["invoice_footer"] = o.InvoiceFooter
	}
	if !IsNil(o.InvoiceLabels) {
		toSerialize["invoice_labels"] = o.InvoiceLabels
	}
	if !IsNil(o.TaxRate1) {
		toSerialize["tax_rate1"] = o.TaxRate1
	}
	if !IsNil(o.TaxName1) {
		toSerialize["tax_name1"] = o.TaxName1
	}
	if !IsNil(o.TaxRate2) {
		toSerialize["tax_rate2"] = o.TaxRate2
	}
	if !IsNil(o.TaxName2) {
		toSerialize["tax_name2"] = o.TaxName2
	}
	if !IsNil(o.TaxRate3) {
		toSerialize["tax_rate3"] = o.TaxRate3
	}
	if !IsNil(o.TaxName3) {
		toSerialize["tax_name3"] = o.TaxName3
	}
	if !IsNil(o.PaymentTypeId) {
		toSerialize["payment_type_id"] = o.PaymentTypeId
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.EmailFooter) {
		toSerialize["email_footer"] = o.EmailFooter
	}
	if !IsNil(o.EmailSendingMethod) {
		toSerialize["email_sending_method"] = o.EmailSendingMethod
	}
	if !IsNil(o.GmailSendingUserId) {
		toSerialize["gmail_sending_user_id"] = o.GmailSendingUserId
	}
	if !IsNil(o.EmailSubjectInvoice) {
		toSerialize["email_subject_invoice"] = o.EmailSubjectInvoice
	}
	if !IsNil(o.EmailSubjectQuote) {
		toSerialize["email_subject_quote"] = o.EmailSubjectQuote
	}
	if !IsNil(o.EmailSubjectPayment) {
		toSerialize["email_subject_payment"] = o.EmailSubjectPayment
	}
	if !IsNil(o.EmailTemplateInvoice) {
		toSerialize["email_template_invoice"] = o.EmailTemplateInvoice
	}
	if !IsNil(o.EmailTemplateQuote) {
		toSerialize["email_template_quote"] = o.EmailTemplateQuote
	}
	if !IsNil(o.EmailTemplatePayment) {
		toSerialize["email_template_payment"] = o.EmailTemplatePayment
	}
	if !IsNil(o.EmailSubjectReminder1) {
		toSerialize["email_subject_reminder1"] = o.EmailSubjectReminder1
	}
	if !IsNil(o.EmailSubjectReminder2) {
		toSerialize["email_subject_reminder2"] = o.EmailSubjectReminder2
	}
	if !IsNil(o.EmailSubjectReminder3) {
		toSerialize["email_subject_reminder3"] = o.EmailSubjectReminder3
	}
	if !IsNil(o.EmailSubjectReminderEndless) {
		toSerialize["email_subject_reminder_endless"] = o.EmailSubjectReminderEndless
	}
	if !IsNil(o.EmailTemplateReminder1) {
		toSerialize["email_template_reminder1"] = o.EmailTemplateReminder1
	}
	if !IsNil(o.EmailTemplateReminder2) {
		toSerialize["email_template_reminder2"] = o.EmailTemplateReminder2
	}
	if !IsNil(o.EmailTemplateReminder3) {
		toSerialize["email_template_reminder3"] = o.EmailTemplateReminder3
	}
	if !IsNil(o.EmailTemplateReminderEndless) {
		toSerialize["email_template_reminder_endless"] = o.EmailTemplateReminderEndless
	}
	if !IsNil(o.EnablePortalPassword) {
		toSerialize["enable_portal_password"] = o.EnablePortalPassword
	}
	if !IsNil(o.ShowAcceptInvoiceTerms) {
		toSerialize["show_accept_invoice_terms"] = o.ShowAcceptInvoiceTerms
	}
	if !IsNil(o.ShowAcceptQuoteTerms) {
		toSerialize["show_accept_quote_terms"] = o.ShowAcceptQuoteTerms
	}
	if !IsNil(o.RequireInvoiceSignature) {
		toSerialize["require_invoice_signature"] = o.RequireInvoiceSignature
	}
	if !IsNil(o.RequireQuoteSignature) {
		toSerialize["require_quote_signature"] = o.RequireQuoteSignature
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CompanyLogo) {
		toSerialize["company_logo"] = o.CompanyLogo
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.CountryId) {
		toSerialize["country_id"] = o.CountryId
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vat_number"] = o.VatNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.FontSize) {
		toSerialize["font_size"] = o.FontSize
	}
	if !IsNil(o.PrimaryFont) {
		toSerialize["primary_font"] = o.PrimaryFont
	}
	if !IsNil(o.SecondaryFont) {
		toSerialize["secondary_font"] = o.SecondaryFont
	}
	if !IsNil(o.HidePaidToDate) {
		toSerialize["hide_paid_to_date"] = o.HidePaidToDate
	}
	if !IsNil(o.EmbedDocuments) {
		toSerialize["embed_documents"] = o.EmbedDocuments
	}
	if !IsNil(o.AllPagesHeader) {
		toSerialize["all_pages_header"] = o.AllPagesHeader
	}
	if !IsNil(o.AllPagesFooter) {
		toSerialize["all_pages_footer"] = o.AllPagesFooter
	}
	if !IsNil(o.DocumentEmailAttachment) {
		toSerialize["document_email_attachment"] = o.DocumentEmailAttachment
	}
	if !IsNil(o.EnableClientPortalPassword) {
		toSerialize["enable_client_portal_password"] = o.EnableClientPortalPassword
	}
	if !IsNil(o.EnableEmailMarkup) {
		toSerialize["enable_email_markup"] = o.EnableEmailMarkup
	}
	if !IsNil(o.EnableClientPortalDashboard) {
		toSerialize["enable_client_portal_dashboard"] = o.EnableClientPortalDashboard
	}
	if !IsNil(o.EnableClientPortal) {
		toSerialize["enable_client_portal"] = o.EnableClientPortal
	}
	if !IsNil(o.EmailTemplateStatement) {
		toSerialize["email_template_statement"] = o.EmailTemplateStatement
	}
	if !IsNil(o.EmailSubjectStatement) {
		toSerialize["email_subject_statement"] = o.EmailSubjectStatement
	}
	if !IsNil(o.SignatureOnPdf) {
		toSerialize["signature_on_pdf"] = o.SignatureOnPdf
	}
	if !IsNil(o.QuoteFooter) {
		toSerialize["quote_footer"] = o.QuoteFooter
	}
	if !IsNil(o.EmailSubjectCustom1) {
		toSerialize["email_subject_custom1"] = o.EmailSubjectCustom1
	}
	if !IsNil(o.EmailSubjectCustom2) {
		toSerialize["email_subject_custom2"] = o.EmailSubjectCustom2
	}
	if !IsNil(o.EmailSubjectCustom3) {
		toSerialize["email_subject_custom3"] = o.EmailSubjectCustom3
	}
	if !IsNil(o.EmailTemplateCustom1) {
		toSerialize["email_template_custom1"] = o.EmailTemplateCustom1
	}
	if !IsNil(o.EmailTemplateCustom2) {
		toSerialize["email_template_custom2"] = o.EmailTemplateCustom2
	}
	if !IsNil(o.EmailTemplateCustom3) {
		toSerialize["email_template_custom3"] = o.EmailTemplateCustom3
	}
	if !IsNil(o.EnableReminder1) {
		toSerialize["enable_reminder1"] = o.EnableReminder1
	}
	if !IsNil(o.EnableReminder2) {
		toSerialize["enable_reminder2"] = o.EnableReminder2
	}
	if !IsNil(o.EnableReminder3) {
		toSerialize["enable_reminder3"] = o.EnableReminder3
	}
	if !IsNil(o.NumDaysReminder1) {
		toSerialize["num_days_reminder1"] = o.NumDaysReminder1
	}
	if !IsNil(o.NumDaysReminder2) {
		toSerialize["num_days_reminder2"] = o.NumDaysReminder2
	}
	if !IsNil(o.NumDaysReminder3) {
		toSerialize["num_days_reminder3"] = o.NumDaysReminder3
	}
	if !IsNil(o.ScheduleReminder1) {
		toSerialize["schedule_reminder1"] = o.ScheduleReminder1
	}
	if !IsNil(o.ScheduleReminder2) {
		toSerialize["schedule_reminder2"] = o.ScheduleReminder2
	}
	if !IsNil(o.ScheduleReminder3) {
		toSerialize["schedule_reminder3"] = o.ScheduleReminder3
	}
	if !IsNil(o.LateFeeAmount1) {
		toSerialize["late_fee_amount1"] = o.LateFeeAmount1
	}
	if !IsNil(o.LateFeeAmount2) {
		toSerialize["late_fee_amount2"] = o.LateFeeAmount2
	}
	if !IsNil(o.LateFeeAmount3) {
		toSerialize["late_fee_amount3"] = o.LateFeeAmount3
	}
	if !IsNil(o.EndlessReminderFrequencyId) {
		toSerialize["endless_reminder_frequency_id"] = o.EndlessReminderFrequencyId
	}
	if !IsNil(o.ClientOnlinePaymentNotification) {
		toSerialize["client_online_payment_notification"] = o.ClientOnlinePaymentNotification
	}
	if !IsNil(o.ClientManualPaymentNotification) {
		toSerialize["client_manual_payment_notification"] = o.ClientManualPaymentNotification
	}
	if !IsNil(o.EnableEInvoice) {
		toSerialize["enable_e_invoice"] = o.EnableEInvoice
	}
	if !IsNil(o.DefaultExpensePaymentTypeId) {
		toSerialize["default_expense_payment_type_id"] = o.DefaultExpensePaymentTypeId
	}
	if !IsNil(o.EInvoiceType) {
		toSerialize["e_invoice_type"] = o.EInvoiceType
	}
	if !IsNil(o.MailgunEndpoint) {
		toSerialize["mailgun_endpoint"] = o.MailgunEndpoint
	}
	if !IsNil(o.ClientInitiatedPayments) {
		toSerialize["client_initiated_payments"] = o.ClientInitiatedPayments
	}
	if !IsNil(o.ClientInitiatedPaymentsMinimum) {
		toSerialize["client_initiated_payments_minimum"] = o.ClientInitiatedPaymentsMinimum
	}
	if !IsNil(o.SyncInvoiceQuoteColumns) {
		toSerialize["sync_invoice_quote_columns"] = o.SyncInvoiceQuoteColumns
	}
	if !IsNil(o.ShowTaskItemDescription) {
		toSerialize["show_task_item_description"] = o.ShowTaskItemDescription
	}
	if !IsNil(o.AllowBillableTaskItems) {
		toSerialize["allow_billable_task_items"] = o.AllowBillableTaskItems
	}
	if !IsNil(o.AcceptClientInputQuoteApproval) {
		toSerialize["accept_client_input_quote_approval"] = o.AcceptClientInputQuoteApproval
	}
	if !IsNil(o.CustomSendingEmail) {
		toSerialize["custom_sending_email"] = o.CustomSendingEmail
	}
	if !IsNil(o.ShowPaidStamp) {
		toSerialize["show_paid_stamp"] = o.ShowPaidStamp
	}
	if !IsNil(o.ShowShippingAddress) {
		toSerialize["show_shipping_address"] = o.ShowShippingAddress
	}
	if !IsNil(o.CompanyLogoSize) {
		toSerialize["company_logo_size"] = o.CompanyLogoSize
	}
	if !IsNil(o.ShowEmailFooter) {
		toSerialize["show_email_footer"] = o.ShowEmailFooter
	}
	if !IsNil(o.EmailAlignment) {
		toSerialize["email_alignment"] = o.EmailAlignment
	}
	if !IsNil(o.AutoBillStandardInvoices) {
		toSerialize["auto_bill_standard_invoices"] = o.AutoBillStandardInvoices
	}
	if !IsNil(o.PostmarkSecret) {
		toSerialize["postmark_secret"] = o.PostmarkSecret
	}
	if !IsNil(o.MailgunSecret) {
		toSerialize["mailgun_secret"] = o.MailgunSecret
	}
	if !IsNil(o.MailgunDomain) {
		toSerialize["mailgun_domain"] = o.MailgunDomain
	}
	if !IsNil(o.SendEmailOnMarkPaid) {
		toSerialize["send_email_on_mark_paid"] = o.SendEmailOnMarkPaid
	}
	if !IsNil(o.VendorPortalEnableUploads) {
		toSerialize["vendor_portal_enable_uploads"] = o.VendorPortalEnableUploads
	}
	if !IsNil(o.BesrId) {
		toSerialize["besr_id"] = o.BesrId
	}
	if !IsNil(o.QrIban) {
		toSerialize["qr_iban"] = o.QrIban
	}
	if !IsNil(o.EmailSubjectPurchaseOrder) {
		toSerialize["email_subject_purchase_order"] = o.EmailSubjectPurchaseOrder
	}
	if !IsNil(o.EmailTemplatePurchaseOrder) {
		toSerialize["email_template_purchase_order"] = o.EmailTemplatePurchaseOrder
	}
	if !IsNil(o.RequirePurchaseOrderSignature) {
		toSerialize["require_purchase_order_signature"] = o.RequirePurchaseOrderSignature
	}
	if !IsNil(o.PurchaseOrderPublicNotes) {
		toSerialize["purchase_order_public_notes"] = o.PurchaseOrderPublicNotes
	}
	if !IsNil(o.PurchaseOrderTerms) {
		toSerialize["purchase_order_terms"] = o.PurchaseOrderTerms
	}
	if !IsNil(o.PurchaseOrderFooter) {
		toSerialize["purchase_order_footer"] = o.PurchaseOrderFooter
	}
	if !IsNil(o.PurchaseOrderDesignId) {
		toSerialize["purchase_order_design_id"] = o.PurchaseOrderDesignId
	}
	if !IsNil(o.PurchaseOrderNumberPattern) {
		toSerialize["purchase_order_number_pattern"] = o.PurchaseOrderNumberPattern
	}
	if !IsNil(o.PurchaseOrderNumberCounter) {
		toSerialize["purchase_order_number_counter"] = o.PurchaseOrderNumberCounter
	}
	if !IsNil(o.PageNumberingAlignment) {
		toSerialize["page_numbering_alignment"] = o.PageNumberingAlignment
	}
	if !IsNil(o.PageNumbering) {
		toSerialize["page_numbering"] = o.PageNumbering
	}
	if !IsNil(o.AutoArchiveInvoiceCancelled) {
		toSerialize["auto_archive_invoice_cancelled"] = o.AutoArchiveInvoiceCancelled
	}
	if !IsNil(o.EmailFromName) {
		toSerialize["email_from_name"] = o.EmailFromName
	}
	if !IsNil(o.ShowAllTasksClientPortal) {
		toSerialize["show_all_tasks_client_portal"] = o.ShowAllTasksClientPortal
	}
	if !IsNil(o.EntitySendTime) {
		toSerialize["entity_send_time"] = o.EntitySendTime
	}
	if !IsNil(o.SharedInvoiceCreditCounter) {
		toSerialize["shared_invoice_credit_counter"] = o.SharedInvoiceCreditCounter
	}
	if !IsNil(o.ReplyToName) {
		toSerialize["reply_to_name"] = o.ReplyToName
	}
	if !IsNil(o.HideEmptyColumnsOnPdf) {
		toSerialize["hide_empty_columns_on_pdf"] = o.HideEmptyColumnsOnPdf
	}
	if !IsNil(o.EnableReminderEndless) {
		toSerialize["enable_reminder_endless"] = o.EnableReminderEndless
	}
	if !IsNil(o.UseCreditsPayment) {
		toSerialize["use_credits_payment"] = o.UseCreditsPayment
	}
	if !IsNil(o.RecurringInvoiceNumberPattern) {
		toSerialize["recurring_invoice_number_pattern"] = o.RecurringInvoiceNumberPattern
	}
	if !IsNil(o.RecurringInvoiceNumberCounter) {
		toSerialize["recurring_invoice_number_counter"] = o.RecurringInvoiceNumberCounter
	}
	if !IsNil(o.ClientPortalUnderPaymentMinimum) {
		toSerialize["client_portal_under_payment_minimum"] = o.ClientPortalUnderPaymentMinimum
	}
	if !IsNil(o.AutoBillDate) {
		toSerialize["auto_bill_date"] = o.AutoBillDate
	}
	if !IsNil(o.PrimaryColor) {
		toSerialize["primary_color"] = o.PrimaryColor
	}
	if !IsNil(o.SecondaryColor) {
		toSerialize["secondary_color"] = o.SecondaryColor
	}
	if !IsNil(o.ClientPortalAllowUnderPayment) {
		toSerialize["client_portal_allow_under_payment"] = o.ClientPortalAllowUnderPayment
	}
	if !IsNil(o.ClientPortalAllowOverPayment) {
		toSerialize["client_portal_allow_over_payment"] = o.ClientPortalAllowOverPayment
	}
	if !IsNil(o.AutoBill) {
		toSerialize["auto_bill"] = o.AutoBill
	}
	if !IsNil(o.ClientPortalTerms) {
		toSerialize["client_portal_terms"] = o.ClientPortalTerms
	}
	if !IsNil(o.ClientPortalPrivacyPolicy) {
		toSerialize["client_portal_privacy_policy"] = o.ClientPortalPrivacyPolicy
	}
	if !IsNil(o.ClientCanRegister) {
		toSerialize["client_can_register"] = o.ClientCanRegister
	}
	if !IsNil(o.PortalDesignId) {
		toSerialize["portal_design_id"] = o.PortalDesignId
	}
	if !IsNil(o.LateFeeEndlessPercent) {
		toSerialize["late_fee_endless_percent"] = o.LateFeeEndlessPercent
	}
	if !IsNil(o.LateFeeEndlessAmount) {
		toSerialize["late_fee_endless_amount"] = o.LateFeeEndlessAmount
	}
	if !IsNil(o.AutoEmailInvoice) {
		toSerialize["auto_email_invoice"] = o.AutoEmailInvoice
	}
	if !IsNil(o.EmailSignature) {
		toSerialize["email_signature"] = o.EmailSignature
	}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !IsNil(o.ValidUntil) {
		toSerialize["valid_until"] = o.ValidUntil
	}
	if !IsNil(o.CreditTerms) {
		toSerialize["credit_terms"] = o.CreditTerms
	}
	if !IsNil(o.CreditFooter) {
		toSerialize["credit_footer"] = o.CreditFooter
	}
	return toSerialize, nil
}

func (o *CompanySettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency_id",
	}

	allProperties := make(map[string]any)

	err = json.Unmarshal(data, &allProperties)
	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCompanySettings := _CompanySettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanySettings)
	if err != nil {
		return err
	}

	*o = CompanySettings(varCompanySettings)

	return err
}

type NullableCompanySettings struct {
	value *CompanySettings
	isSet bool
}

func (v NullableCompanySettings) Get() *CompanySettings {
	return v.value
}

func (v *NullableCompanySettings) Set(val *CompanySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanySettings(val *CompanySettings) *NullableCompanySettings {
	return &NullableCompanySettings{value: val, isSet: true}
}

func (v NullableCompanySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
