package openapi

import (
	"encoding/json"
	"os"
)

// checks if the Company type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Company{}

// Company struct for Company
type Company struct {
	// The unique hashed identifier for the company
	Id *string `json:"id,omitempty"`
	// The unique identifier representing the company's size category
	SizeId *string `json:"size_id,omitempty"`
	// The unique identifier representing the company's industry category
	IndustryId *string `json:"industry_id,omitempty"`
	// The URL for the company's Slack webhook notifications
	SlackWebhookUrl *string `json:"slack_webhook_url,omitempty"`
	// The company's Google Analytics tracking ID
	GoogleAnalyticsKey *string `json:"google_analytics_key,omitempty"`
	// The mode determining how client-facing URLs are structured (e.g., subdomain, domain, or iframe)
	PortalMode *string `json:"portal_mode,omitempty"`
	// The subdomain prefix for the company's domain (e.g., 'acme' in acme.domain.com)
	Subdomain *string `json:"subdomain,omitempty"`
	// The fully qualified domain used for client-facing URLs
	PortalDomain *string `json:"portal_domain,omitempty"`
	// The number of tax rates used per entity
	EnabledTaxRates *int32 `json:"enabled_tax_rates,omitempty"`
	// A flag determining whether to auto-fill product descriptions based on the product key
	FillProducts *bool `json:"fill_products,omitempty"`
	// A flag determining whether to convert products between different types or units
	ConvertProducts *bool `json:"convert_products,omitempty"`
	// A flag determining whether to update product descriptions when the description changes
	UpdateProducts *bool `json:"update_products,omitempty"`
	// A flag determining whether to display product details in the user interface
	ShowProductDetails *bool `json:"show_product_details,omitempty"`
	// A flag determining whether to display product cost is shown in the user interface
	ShowProductCost *bool `json:"show_product_cost,omitempty"`
	// A mapping of custom fields for various objects within the company  **note**  The custom fields object is formatted as follows    key : value   entity(integer) : label|type    for example:   \"company1\": \"Number|single_line_text\",    This defines the first company custom field, with label Number which has a custom field type of a single text line    Supported entity types    invoice   product   client   contact   task   user   project   vendor    expense   payment         Supported input types values    single_line_text   date   switch    For text areas, you only need to supply the label ie \"TextArea\", the | and type values are not required.   For the dropdown the data format is as follows:    label|your,drop,down,values,in,comma,separated,format
	CustomFields map[string]any `json:"custom_fields,omitempty"`
	// A flag determining whether to show or hide the product cost field in the user interface
	EnableProductCost *bool `json:"enable_product_cost,omitempty"`
	// A flag determining whether to show or hide the product quantity field in the user interface
	EnableProductQuantity *bool `json:"enable_product_quantity,omitempty"`
	// A flag determining whether to use a default quantity for products
	DefaultQuantity *bool `json:"default_quantity,omitempty"`
	// A flag determining whether to apply taxes on custom surcharge amounts for the first custom surcharge field
	CustomSurchargeTaxes1 *bool `json:"custom_surcharge_taxes1,omitempty"`
	// A flag determining whether to apply taxes on custom surcharge amounts for the second custom surcharge field
	CustomSurchargeTaxes2 *bool `json:"custom_surcharge_taxes2,omitempty"`
	// A flag determining whether to apply taxes on custom surcharge amounts for the third custom surcharge field
	CustomSurchargeTaxes3 *bool `json:"custom_surcharge_taxes3,omitempty"`
	// A flag determining whether to apply taxes on custom surcharge amounts for the fourth custom
	CustomSurchargeTaxes4 any `json:"custom_surcharge_taxes4,omitempty"`
	// The company logo file in binary format
	Logo **os.File `json:"logo,omitempty"`
	// The static company key hash used to identify the Company
	CompanyKey *string `json:"company_key,omitempty"`
	// A flag determining whether clients can register for the client portal
	ClientCanRegister *bool `json:"client_can_register,omitempty"`
	// Bitmask representation of the modules that are enabled in the application  ``` self::ENTITY_RECURRING_INVOICE => 1, self::ENTITY_CREDIT => 2, self::ENTITY_QUOTE => 4, self::ENTITY_TASK => 8, self::ENTITY_EXPENSE => 16, self::ENTITY_PROJECT => 32, self::ENTITY_VENDOR => 64, self::ENTITY_TICKET => 128, self::ENTITY_PROPOSAL => 256, self::ENTITY_RECURRING_EXPENSE => 512, self::ENTITY_RECURRING_TASK => 1024, self::ENTITY_RECURRING_QUOTE => 2048, ```  The default per_page value is 20.
	EnabledModules *int32  `json:"enabled_modules,omitempty"`
	Db             *string `json:"db,omitempty"`
	// The first day of the week for the company
	FirstDayOfWeek *string `json:"first_day_of_week,omitempty"`
	// The first month for the company financial year
	FirstMonthOfYear *string `json:"first_month_of_year,omitempty"`
	// The number of tax rates used per item
	EnabledItemTaxRates *int32 `json:"enabled_item_tax_rates,omitempty"`
	// A flag determining whether the company is considered large
	IsLarge *bool `json:"is_large,omitempty"`
	// A flag determining whether to auto-bill clients by default  values:  - always - Always auto bill - disabled - Never auto bill - optin - Allow the client to select their auto bill status with the default being disabled - optout -Allow the client to select their auto bill status with the default being enabled
	DefaultAutoBill *string `json:"default_auto_bill,omitempty"`
	// A flag determining whether to mark expenses as invoiceable by default
	MarkExpensesInvoiceable *bool `json:"mark_expenses_invoiceable,omitempty"`
	// A flag determining whether to mark expenses as paid by default
	MarkExpensesPaid *bool `json:"mark_expenses_paid,omitempty"`
	// A flag determining whether to include expense documents on invoices by default
	InvoiceExpenseDocuments *bool `json:"invoice_expense_documents,omitempty"`
	// A flag determining whether to auto-start tasks by default
	AutoStartTasks *bool `json:"auto_start_tasks,omitempty"`
	// A flag determining whether to include task time logs on invoices by default
	InvoiceTaskTimelog *bool `json:"invoice_task_timelog,omitempty"`
	// A flag determining whether to include task documents on invoices by default
	InvoiceTaskDocuments *bool `json:"invoice_task_documents,omitempty"`
	// A flag determining whether to show the tasks table on invoices by default
	ShowTasksTable *bool `json:"show_tasks_table,omitempty"`
	// A flag determining whether the company is disabled
	IsDisabled *bool `json:"is_disabled,omitempty"`
	// A flag determining whether to default tasks to be date-based
	DefaultTaskIsDateBased *bool `json:"default_task_is_date_based,omitempty"`
	// A flag determining whether to show or hide the product discount field in the user interface
	EnableProductDiscount *bool `json:"enable_product_discount,omitempty"`
	// A flag determining whether to calculate expense taxes by amount
	CalculateExpenseTaxByAmount *bool `json:"calculate_expense_tax_by_amount,omitempty"`
	// A flag determining whether to include taxes in the expense amount
	ExpenseInclusiveTaxes *bool `json:"expense_inclusive_taxes,omitempty"`
	// The session timeout for the company
	SessionTimeout *int32 `json:"session_timeout,omitempty"`
	// A flag determining whether to require a password for `dangerous` actions when using OAuth
	OauthPasswordRequired *bool `json:"oauth_password_required,omitempty"`
	// A flag determining whether to include task date logs on invoices by default
	InvoiceTaskDatelog *bool `json:"invoice_task_datelog,omitempty"`
	// The default password timeout for the company
	DefaultPasswordTimeout *int32 `json:"default_password_timeout,omitempty"`
	// A flag determining whether to show the task end date on invoices by default
	ShowTaskEndDate *bool `json:"show_task_end_date,omitempty"`
	// A flag determining whether markdown is enabled for the company
	MarkdownEnabled *bool `json:"markdown_enabled,omitempty"`
	// A flag determining whether to include draft invoices in reports
	ReportIncludeDrafts *bool `json:"report_include_drafts,omitempty"`
	// The client registration fields for the company
	ClientRegistrationFields map[string]any `json:"client_registration_fields,omitempty"`
	// A flag determining whether to stop recurring invoices when they are unpaid
	StopOnUnpaidRecurring *bool `json:"stop_on_unpaid_recurring,omitempty"`
	// A flag determining whether to use quote terms on conversion to an invoice
	UseQuoteTermsOnConversion *bool `json:"use_quote_terms_on_conversion,omitempty"`
	// A flag determining whether to enable applying payments to invoices
	EnableApplyingPayments *bool `json:"enable_applying_payments,omitempty"`
	// A flag determining whether to track inventory for the company
	TrackInventory *bool `json:"track_inventory,omitempty"`
	// The inventory notification threshold for the company
	InventoryNotificationThreshold *int32 `json:"inventory_notification_threshold,omitempty"`
	// A flag determining whether to send stock notifications for the company
	StockNotification *bool `json:"stock_notification,omitempty"`
	// The Matomo URL for the company
	MatomoUrl *string `json:"matomo_url,omitempty"`
	// The Matomo ID for the company
	MatomoId *string `json:"matomo_id,omitempty"`
	// The number of tax rates used per expense
	EnabledExpenseTaxRates *int32 `json:"enabled_expense_tax_rates,omitempty"`
	// A flag determining whether to include the project on invoices by default
	InvoiceTaskProject *bool `json:"invoice_task_project,omitempty"`
	// A flag determining whether to include deleted invoices in reports
	ReportIncludeDeleted *bool `json:"report_include_deleted,omitempty"`
	// A flag determining whether to lock tasks when invoiced
	InvoiceTaskLock *bool `json:"invoice_task_lock,omitempty"`
	// A flag determining whether to convert the payment currency
	ConvertPaymentCurrency *bool `json:"convert_payment_currency,omitempty"`
	// A flag determining whether to convert the expense currency
	ConvertExpenseCurrency *bool `json:"convert_expense_currency,omitempty"`
	// A flag determining whether to notify the vendor when an expense is paid
	NotifyVendorWhenPaid *bool `json:"notify_vendor_when_paid,omitempty"`
	// A flag determining whether to include the task hours on invoices by default
	InvoiceTaskHours *bool `json:"invoice_task_hours,omitempty"`
	// A flag determining whether to calculate taxes for the company
	CalculateTaxes *bool `json:"calculate_taxes,omitempty"`
	// The tax data for the company
	TaxData map[string]any `json:"tax_data,omitempty"`
	// The e-invoice certificate for the company
	EInvoiceCertificate *string `json:"e_invoice_certificate,omitempty"`
	// The e-invoice certificate passphrase for the company
	EInvoiceCertificatePassphrase *string `json:"e_invoice_certificate_passphrase,omitempty"`
	// The origin tax data for the company
	OriginTaxData map[string]any `json:"origin_tax_data,omitempty"`
	// A flag determining whether to include the project header on invoices by default
	InvoiceTaskProjectHeader *bool `json:"invoice_task_project_header,omitempty"`
	// A flag determining whether to include the item description on invoices by default
	InvoiceTaskItemDescription *bool `json:"invoice_task_item_description,omitempty"`
	// The email address for the expense mailbox
	ExpenseMailbox *string `json:"expense_mailbox,omitempty"`
	// Whether the expense mailbox is active
	ExpenseMailboxActive *bool `json:"expense_mailbox_active,omitempty"`
	// Whether company users are allowed to use the inbound mailbox
	InboundMailboxAllowCompanyUsers *bool `json:"inbound_mailbox_allow_company_users,omitempty"`
	// Whether vendors are allowed to use the inbound mailbox
	InboundMailboxAllowVendors *bool `json:"inbound_mailbox_allow_vendors,omitempty"`
	// Whether clients are allowed to use the inbound mailbox
	InboundMailboxAllowClients *bool `json:"inbound_mailbox_allow_clients,omitempty"`
	// Whether unknown senders are allowed to use the inbound mailbox
	InboundMailboxAllowUnknown *bool `json:"inbound_mailbox_allow_unknown,omitempty"`
	// Comma-separated list of whitelisted email addresses for the inbound mailbox
	InboundMailboxWhitelist *string `json:"inbound_mailbox_whitelist,omitempty"`
	// Comma-separated list of blacklisted email addresses for the inbound mailbox
	InboundMailboxBlacklist *string `json:"inbound_mailbox_blacklist,omitempty"`
	// The SMTP host for sending emails
	SmtpHost *string `json:"smtp_host,omitempty"`
	// The SMTP port for sending emails
	SmtpPort *int32 `json:"smtp_port,omitempty"`
	// The encryption method for SMTP
	SmtpEncryption NullableString `json:"smtp_encryption,omitempty"`
	// The local domain for SMTP
	SmtpLocalDomain *string `json:"smtp_local_domain,omitempty"`
	// Whether to verify the SMTP peer
	SmtpVerifyPeer *bool `json:"smtp_verify_peer,omitempty"`
	// E-invoice settings for the company
	EInvoice map[string]any `json:"e_invoice,omitempty"`
	// The ID of the legal entity associated with the company
	LegalEntityId *int32           `json:"legal_entity_id,omitempty"`
	Settings      *CompanySettings `json:"settings,omitempty"`
}

// NewCompany instantiates a new Company object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompany() *Company {
	this := Company{}
	return &this
}

// NewCompanyWithDefaults instantiates a new Company object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyWithDefaults() *Company {
	this := Company{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Company) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Company) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Company) SetId(v string) {
	o.Id = &v
}

// GetSizeId returns the SizeId field value if set, zero value otherwise.
func (o *Company) GetSizeId() string {
	if o == nil || IsNil(o.SizeId) {
		var ret string
		return ret
	}
	return *o.SizeId
}

// GetSizeIdOk returns a tuple with the SizeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSizeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SizeId) {
		return nil, false
	}
	return o.SizeId, true
}

// HasSizeId returns a boolean if a field has been set.
func (o *Company) HasSizeId() bool {
	if o != nil && !IsNil(o.SizeId) {
		return true
	}

	return false
}

// SetSizeId gets a reference to the given string and assigns it to the SizeId field.
func (o *Company) SetSizeId(v string) {
	o.SizeId = &v
}

// GetIndustryId returns the IndustryId field value if set, zero value otherwise.
func (o *Company) GetIndustryId() string {
	if o == nil || IsNil(o.IndustryId) {
		var ret string
		return ret
	}
	return *o.IndustryId
}

// GetIndustryIdOk returns a tuple with the IndustryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetIndustryIdOk() (*string, bool) {
	if o == nil || IsNil(o.IndustryId) {
		return nil, false
	}
	return o.IndustryId, true
}

// HasIndustryId returns a boolean if a field has been set.
func (o *Company) HasIndustryId() bool {
	if o != nil && !IsNil(o.IndustryId) {
		return true
	}

	return false
}

// SetIndustryId gets a reference to the given string and assigns it to the IndustryId field.
func (o *Company) SetIndustryId(v string) {
	o.IndustryId = &v
}

// GetSlackWebhookUrl returns the SlackWebhookUrl field value if set, zero value otherwise.
func (o *Company) GetSlackWebhookUrl() string {
	if o == nil || IsNil(o.SlackWebhookUrl) {
		var ret string
		return ret
	}
	return *o.SlackWebhookUrl
}

// GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSlackWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SlackWebhookUrl) {
		return nil, false
	}
	return o.SlackWebhookUrl, true
}

// HasSlackWebhookUrl returns a boolean if a field has been set.
func (o *Company) HasSlackWebhookUrl() bool {
	if o != nil && !IsNil(o.SlackWebhookUrl) {
		return true
	}

	return false
}

// SetSlackWebhookUrl gets a reference to the given string and assigns it to the SlackWebhookUrl field.
func (o *Company) SetSlackWebhookUrl(v string) {
	o.SlackWebhookUrl = &v
}

// GetGoogleAnalyticsKey returns the GoogleAnalyticsKey field value if set, zero value otherwise.
func (o *Company) GetGoogleAnalyticsKey() string {
	if o == nil || IsNil(o.GoogleAnalyticsKey) {
		var ret string
		return ret
	}
	return *o.GoogleAnalyticsKey
}

// GetGoogleAnalyticsKeyOk returns a tuple with the GoogleAnalyticsKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetGoogleAnalyticsKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleAnalyticsKey) {
		return nil, false
	}
	return o.GoogleAnalyticsKey, true
}

// HasGoogleAnalyticsKey returns a boolean if a field has been set.
func (o *Company) HasGoogleAnalyticsKey() bool {
	if o != nil && !IsNil(o.GoogleAnalyticsKey) {
		return true
	}

	return false
}

// SetGoogleAnalyticsKey gets a reference to the given string and assigns it to the GoogleAnalyticsKey field.
func (o *Company) SetGoogleAnalyticsKey(v string) {
	o.GoogleAnalyticsKey = &v
}

// GetPortalMode returns the PortalMode field value if set, zero value otherwise.
func (o *Company) GetPortalMode() string {
	if o == nil || IsNil(o.PortalMode) {
		var ret string
		return ret
	}
	return *o.PortalMode
}

// GetPortalModeOk returns a tuple with the PortalMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetPortalModeOk() (*string, bool) {
	if o == nil || IsNil(o.PortalMode) {
		return nil, false
	}
	return o.PortalMode, true
}

// HasPortalMode returns a boolean if a field has been set.
func (o *Company) HasPortalMode() bool {
	if o != nil && !IsNil(o.PortalMode) {
		return true
	}

	return false
}

// SetPortalMode gets a reference to the given string and assigns it to the PortalMode field.
func (o *Company) SetPortalMode(v string) {
	o.PortalMode = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *Company) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *Company) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *Company) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetPortalDomain returns the PortalDomain field value if set, zero value otherwise.
func (o *Company) GetPortalDomain() string {
	if o == nil || IsNil(o.PortalDomain) {
		var ret string
		return ret
	}
	return *o.PortalDomain
}

// GetPortalDomainOk returns a tuple with the PortalDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetPortalDomainOk() (*string, bool) {
	if o == nil || IsNil(o.PortalDomain) {
		return nil, false
	}
	return o.PortalDomain, true
}

// HasPortalDomain returns a boolean if a field has been set.
func (o *Company) HasPortalDomain() bool {
	if o != nil && !IsNil(o.PortalDomain) {
		return true
	}

	return false
}

// SetPortalDomain gets a reference to the given string and assigns it to the PortalDomain field.
func (o *Company) SetPortalDomain(v string) {
	o.PortalDomain = &v
}

// GetEnabledTaxRates returns the EnabledTaxRates field value if set, zero value otherwise.
func (o *Company) GetEnabledTaxRates() int32 {
	if o == nil || IsNil(o.EnabledTaxRates) {
		var ret int32
		return ret
	}
	return *o.EnabledTaxRates
}

// GetEnabledTaxRatesOk returns a tuple with the EnabledTaxRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnabledTaxRatesOk() (*int32, bool) {
	if o == nil || IsNil(o.EnabledTaxRates) {
		return nil, false
	}
	return o.EnabledTaxRates, true
}

// HasEnabledTaxRates returns a boolean if a field has been set.
func (o *Company) HasEnabledTaxRates() bool {
	if o != nil && !IsNil(o.EnabledTaxRates) {
		return true
	}

	return false
}

// SetEnabledTaxRates gets a reference to the given int32 and assigns it to the EnabledTaxRates field.
func (o *Company) SetEnabledTaxRates(v int32) {
	o.EnabledTaxRates = &v
}

// GetFillProducts returns the FillProducts field value if set, zero value otherwise.
func (o *Company) GetFillProducts() bool {
	if o == nil || IsNil(o.FillProducts) {
		var ret bool
		return ret
	}
	return *o.FillProducts
}

// GetFillProductsOk returns a tuple with the FillProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetFillProductsOk() (*bool, bool) {
	if o == nil || IsNil(o.FillProducts) {
		return nil, false
	}
	return o.FillProducts, true
}

// HasFillProducts returns a boolean if a field has been set.
func (o *Company) HasFillProducts() bool {
	if o != nil && !IsNil(o.FillProducts) {
		return true
	}

	return false
}

// SetFillProducts gets a reference to the given bool and assigns it to the FillProducts field.
func (o *Company) SetFillProducts(v bool) {
	o.FillProducts = &v
}

// GetConvertProducts returns the ConvertProducts field value if set, zero value otherwise.
func (o *Company) GetConvertProducts() bool {
	if o == nil || IsNil(o.ConvertProducts) {
		var ret bool
		return ret
	}
	return *o.ConvertProducts
}

// GetConvertProductsOk returns a tuple with the ConvertProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetConvertProductsOk() (*bool, bool) {
	if o == nil || IsNil(o.ConvertProducts) {
		return nil, false
	}
	return o.ConvertProducts, true
}

// HasConvertProducts returns a boolean if a field has been set.
func (o *Company) HasConvertProducts() bool {
	if o != nil && !IsNil(o.ConvertProducts) {
		return true
	}

	return false
}

// SetConvertProducts gets a reference to the given bool and assigns it to the ConvertProducts field.
func (o *Company) SetConvertProducts(v bool) {
	o.ConvertProducts = &v
}

// GetUpdateProducts returns the UpdateProducts field value if set, zero value otherwise.
func (o *Company) GetUpdateProducts() bool {
	if o == nil || IsNil(o.UpdateProducts) {
		var ret bool
		return ret
	}
	return *o.UpdateProducts
}

// GetUpdateProductsOk returns a tuple with the UpdateProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetUpdateProductsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateProducts) {
		return nil, false
	}
	return o.UpdateProducts, true
}

// HasUpdateProducts returns a boolean if a field has been set.
func (o *Company) HasUpdateProducts() bool {
	if o != nil && !IsNil(o.UpdateProducts) {
		return true
	}

	return false
}

// SetUpdateProducts gets a reference to the given bool and assigns it to the UpdateProducts field.
func (o *Company) SetUpdateProducts(v bool) {
	o.UpdateProducts = &v
}

// GetShowProductDetails returns the ShowProductDetails field value if set, zero value otherwise.
func (o *Company) GetShowProductDetails() bool {
	if o == nil || IsNil(o.ShowProductDetails) {
		var ret bool
		return ret
	}
	return *o.ShowProductDetails
}

// GetShowProductDetailsOk returns a tuple with the ShowProductDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetShowProductDetailsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowProductDetails) {
		return nil, false
	}
	return o.ShowProductDetails, true
}

// HasShowProductDetails returns a boolean if a field has been set.
func (o *Company) HasShowProductDetails() bool {
	if o != nil && !IsNil(o.ShowProductDetails) {
		return true
	}

	return false
}

// SetShowProductDetails gets a reference to the given bool and assigns it to the ShowProductDetails field.
func (o *Company) SetShowProductDetails(v bool) {
	o.ShowProductDetails = &v
}

// GetShowProductCost returns the ShowProductCost field value if set, zero value otherwise.
func (o *Company) GetShowProductCost() bool {
	if o == nil || IsNil(o.ShowProductCost) {
		var ret bool
		return ret
	}
	return *o.ShowProductCost
}

// GetShowProductCostOk returns a tuple with the ShowProductCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetShowProductCostOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowProductCost) {
		return nil, false
	}
	return o.ShowProductCost, true
}

// HasShowProductCost returns a boolean if a field has been set.
func (o *Company) HasShowProductCost() bool {
	if o != nil && !IsNil(o.ShowProductCost) {
		return true
	}

	return false
}

// SetShowProductCost gets a reference to the given bool and assigns it to the ShowProductCost field.
func (o *Company) SetShowProductCost(v bool) {
	o.ShowProductCost = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Company) GetCustomFields() map[string]any {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]any
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCustomFieldsOk() (map[string]any, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]any{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Company) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Company) SetCustomFields(v map[string]any) {
	o.CustomFields = v
}

// GetEnableProductCost returns the EnableProductCost field value if set, zero value otherwise.
func (o *Company) GetEnableProductCost() bool {
	if o == nil || IsNil(o.EnableProductCost) {
		var ret bool
		return ret
	}
	return *o.EnableProductCost
}

// GetEnableProductCostOk returns a tuple with the EnableProductCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnableProductCostOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableProductCost) {
		return nil, false
	}
	return o.EnableProductCost, true
}

// HasEnableProductCost returns a boolean if a field has been set.
func (o *Company) HasEnableProductCost() bool {
	if o != nil && !IsNil(o.EnableProductCost) {
		return true
	}

	return false
}

// SetEnableProductCost gets a reference to the given bool and assigns it to the EnableProductCost field.
func (o *Company) SetEnableProductCost(v bool) {
	o.EnableProductCost = &v
}

// GetEnableProductQuantity returns the EnableProductQuantity field value if set, zero value otherwise.
func (o *Company) GetEnableProductQuantity() bool {
	if o == nil || IsNil(o.EnableProductQuantity) {
		var ret bool
		return ret
	}
	return *o.EnableProductQuantity
}

// GetEnableProductQuantityOk returns a tuple with the EnableProductQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnableProductQuantityOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableProductQuantity) {
		return nil, false
	}
	return o.EnableProductQuantity, true
}

// HasEnableProductQuantity returns a boolean if a field has been set.
func (o *Company) HasEnableProductQuantity() bool {
	if o != nil && !IsNil(o.EnableProductQuantity) {
		return true
	}

	return false
}

// SetEnableProductQuantity gets a reference to the given bool and assigns it to the EnableProductQuantity field.
func (o *Company) SetEnableProductQuantity(v bool) {
	o.EnableProductQuantity = &v
}

// GetDefaultQuantity returns the DefaultQuantity field value if set, zero value otherwise.
func (o *Company) GetDefaultQuantity() bool {
	if o == nil || IsNil(o.DefaultQuantity) {
		var ret bool
		return ret
	}
	return *o.DefaultQuantity
}

// GetDefaultQuantityOk returns a tuple with the DefaultQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetDefaultQuantityOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultQuantity) {
		return nil, false
	}
	return o.DefaultQuantity, true
}

// HasDefaultQuantity returns a boolean if a field has been set.
func (o *Company) HasDefaultQuantity() bool {
	if o != nil && !IsNil(o.DefaultQuantity) {
		return true
	}

	return false
}

// SetDefaultQuantity gets a reference to the given bool and assigns it to the DefaultQuantity field.
func (o *Company) SetDefaultQuantity(v bool) {
	o.DefaultQuantity = &v
}

// GetCustomSurchargeTaxes1 returns the CustomSurchargeTaxes1 field value if set, zero value otherwise.
func (o *Company) GetCustomSurchargeTaxes1() bool {
	if o == nil || IsNil(o.CustomSurchargeTaxes1) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTaxes1
}

// GetCustomSurchargeTaxes1Ok returns a tuple with the CustomSurchargeTaxes1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCustomSurchargeTaxes1Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTaxes1) {
		return nil, false
	}
	return o.CustomSurchargeTaxes1, true
}

// HasCustomSurchargeTaxes1 returns a boolean if a field has been set.
func (o *Company) HasCustomSurchargeTaxes1() bool {
	if o != nil && !IsNil(o.CustomSurchargeTaxes1) {
		return true
	}

	return false
}

// SetCustomSurchargeTaxes1 gets a reference to the given bool and assigns it to the CustomSurchargeTaxes1 field.
func (o *Company) SetCustomSurchargeTaxes1(v bool) {
	o.CustomSurchargeTaxes1 = &v
}

// GetCustomSurchargeTaxes2 returns the CustomSurchargeTaxes2 field value if set, zero value otherwise.
func (o *Company) GetCustomSurchargeTaxes2() bool {
	if o == nil || IsNil(o.CustomSurchargeTaxes2) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTaxes2
}

// GetCustomSurchargeTaxes2Ok returns a tuple with the CustomSurchargeTaxes2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCustomSurchargeTaxes2Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTaxes2) {
		return nil, false
	}
	return o.CustomSurchargeTaxes2, true
}

// HasCustomSurchargeTaxes2 returns a boolean if a field has been set.
func (o *Company) HasCustomSurchargeTaxes2() bool {
	if o != nil && !IsNil(o.CustomSurchargeTaxes2) {
		return true
	}

	return false
}

// SetCustomSurchargeTaxes2 gets a reference to the given bool and assigns it to the CustomSurchargeTaxes2 field.
func (o *Company) SetCustomSurchargeTaxes2(v bool) {
	o.CustomSurchargeTaxes2 = &v
}

// GetCustomSurchargeTaxes3 returns the CustomSurchargeTaxes3 field value if set, zero value otherwise.
func (o *Company) GetCustomSurchargeTaxes3() bool {
	if o == nil || IsNil(o.CustomSurchargeTaxes3) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTaxes3
}

// GetCustomSurchargeTaxes3Ok returns a tuple with the CustomSurchargeTaxes3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCustomSurchargeTaxes3Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTaxes3) {
		return nil, false
	}
	return o.CustomSurchargeTaxes3, true
}

// HasCustomSurchargeTaxes3 returns a boolean if a field has been set.
func (o *Company) HasCustomSurchargeTaxes3() bool {
	if o != nil && !IsNil(o.CustomSurchargeTaxes3) {
		return true
	}

	return false
}

// SetCustomSurchargeTaxes3 gets a reference to the given bool and assigns it to the CustomSurchargeTaxes3 field.
func (o *Company) SetCustomSurchargeTaxes3(v bool) {
	o.CustomSurchargeTaxes3 = &v
}

// GetCustomSurchargeTaxes4 returns the CustomSurchargeTaxes4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetCustomSurchargeTaxes4() any {
	if o == nil {
		var ret any
		return ret
	}
	return o.CustomSurchargeTaxes4
}

// GetCustomSurchargeTaxes4Ok returns a tuple with the CustomSurchargeTaxes4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetCustomSurchargeTaxes4Ok() (*any, bool) {
	if o == nil || IsNil(o.CustomSurchargeTaxes4) {
		return nil, false
	}
	return &o.CustomSurchargeTaxes4, true
}

// HasCustomSurchargeTaxes4 returns a boolean if a field has been set.
func (o *Company) HasCustomSurchargeTaxes4() bool {
	if o != nil && !IsNil(o.CustomSurchargeTaxes4) {
		return true
	}

	return false
}

// SetCustomSurchargeTaxes4 gets a reference to the given interface{} and assigns it to the CustomSurchargeTaxes4 field.
func (o *Company) SetCustomSurchargeTaxes4(v any) {
	o.CustomSurchargeTaxes4 = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *Company) GetLogo() *os.File {
	if o == nil || IsNil(o.Logo) {
		var ret *os.File
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetLogoOk() (**os.File, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *Company) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given *os.File and assigns it to the Logo field.
func (o *Company) SetLogo(v *os.File) {
	o.Logo = &v
}

// GetCompanyKey returns the CompanyKey field value if set, zero value otherwise.
func (o *Company) GetCompanyKey() string {
	if o == nil || IsNil(o.CompanyKey) {
		var ret string
		return ret
	}
	return *o.CompanyKey
}

// GetCompanyKeyOk returns a tuple with the CompanyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCompanyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyKey) {
		return nil, false
	}
	return o.CompanyKey, true
}

// HasCompanyKey returns a boolean if a field has been set.
func (o *Company) HasCompanyKey() bool {
	if o != nil && !IsNil(o.CompanyKey) {
		return true
	}

	return false
}

// SetCompanyKey gets a reference to the given string and assigns it to the CompanyKey field.
func (o *Company) SetCompanyKey(v string) {
	o.CompanyKey = &v
}

// GetClientCanRegister returns the ClientCanRegister field value if set, zero value otherwise.
func (o *Company) GetClientCanRegister() bool {
	if o == nil || IsNil(o.ClientCanRegister) {
		var ret bool
		return ret
	}
	return *o.ClientCanRegister
}

// GetClientCanRegisterOk returns a tuple with the ClientCanRegister field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetClientCanRegisterOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientCanRegister) {
		return nil, false
	}
	return o.ClientCanRegister, true
}

// HasClientCanRegister returns a boolean if a field has been set.
func (o *Company) HasClientCanRegister() bool {
	if o != nil && !IsNil(o.ClientCanRegister) {
		return true
	}

	return false
}

// SetClientCanRegister gets a reference to the given bool and assigns it to the ClientCanRegister field.
func (o *Company) SetClientCanRegister(v bool) {
	o.ClientCanRegister = &v
}

// GetEnabledModules returns the EnabledModules field value if set, zero value otherwise.
func (o *Company) GetEnabledModules() int32 {
	if o == nil || IsNil(o.EnabledModules) {
		var ret int32
		return ret
	}
	return *o.EnabledModules
}

// GetEnabledModulesOk returns a tuple with the EnabledModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnabledModulesOk() (*int32, bool) {
	if o == nil || IsNil(o.EnabledModules) {
		return nil, false
	}
	return o.EnabledModules, true
}

// HasEnabledModules returns a boolean if a field has been set.
func (o *Company) HasEnabledModules() bool {
	if o != nil && !IsNil(o.EnabledModules) {
		return true
	}

	return false
}

// SetEnabledModules gets a reference to the given int32 and assigns it to the EnabledModules field.
func (o *Company) SetEnabledModules(v int32) {
	o.EnabledModules = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *Company) GetDb() string {
	if o == nil || IsNil(o.Db) {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetDbOk() (*string, bool) {
	if o == nil || IsNil(o.Db) {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *Company) HasDb() bool {
	if o != nil && !IsNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *Company) SetDb(v string) {
	o.Db = &v
}

// GetFirstDayOfWeek returns the FirstDayOfWeek field value if set, zero value otherwise.
func (o *Company) GetFirstDayOfWeek() string {
	if o == nil || IsNil(o.FirstDayOfWeek) {
		var ret string
		return ret
	}
	return *o.FirstDayOfWeek
}

// GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetFirstDayOfWeekOk() (*string, bool) {
	if o == nil || IsNil(o.FirstDayOfWeek) {
		return nil, false
	}
	return o.FirstDayOfWeek, true
}

// HasFirstDayOfWeek returns a boolean if a field has been set.
func (o *Company) HasFirstDayOfWeek() bool {
	if o != nil && !IsNil(o.FirstDayOfWeek) {
		return true
	}

	return false
}

// SetFirstDayOfWeek gets a reference to the given string and assigns it to the FirstDayOfWeek field.
func (o *Company) SetFirstDayOfWeek(v string) {
	o.FirstDayOfWeek = &v
}

// GetFirstMonthOfYear returns the FirstMonthOfYear field value if set, zero value otherwise.
func (o *Company) GetFirstMonthOfYear() string {
	if o == nil || IsNil(o.FirstMonthOfYear) {
		var ret string
		return ret
	}
	return *o.FirstMonthOfYear
}

// GetFirstMonthOfYearOk returns a tuple with the FirstMonthOfYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetFirstMonthOfYearOk() (*string, bool) {
	if o == nil || IsNil(o.FirstMonthOfYear) {
		return nil, false
	}
	return o.FirstMonthOfYear, true
}

// HasFirstMonthOfYear returns a boolean if a field has been set.
func (o *Company) HasFirstMonthOfYear() bool {
	if o != nil && !IsNil(o.FirstMonthOfYear) {
		return true
	}

	return false
}

// SetFirstMonthOfYear gets a reference to the given string and assigns it to the FirstMonthOfYear field.
func (o *Company) SetFirstMonthOfYear(v string) {
	o.FirstMonthOfYear = &v
}

// GetEnabledItemTaxRates returns the EnabledItemTaxRates field value if set, zero value otherwise.
func (o *Company) GetEnabledItemTaxRates() int32 {
	if o == nil || IsNil(o.EnabledItemTaxRates) {
		var ret int32
		return ret
	}
	return *o.EnabledItemTaxRates
}

// GetEnabledItemTaxRatesOk returns a tuple with the EnabledItemTaxRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnabledItemTaxRatesOk() (*int32, bool) {
	if o == nil || IsNil(o.EnabledItemTaxRates) {
		return nil, false
	}
	return o.EnabledItemTaxRates, true
}

// HasEnabledItemTaxRates returns a boolean if a field has been set.
func (o *Company) HasEnabledItemTaxRates() bool {
	if o != nil && !IsNil(o.EnabledItemTaxRates) {
		return true
	}

	return false
}

// SetEnabledItemTaxRates gets a reference to the given int32 and assigns it to the EnabledItemTaxRates field.
func (o *Company) SetEnabledItemTaxRates(v int32) {
	o.EnabledItemTaxRates = &v
}

// GetIsLarge returns the IsLarge field value if set, zero value otherwise.
func (o *Company) GetIsLarge() bool {
	if o == nil || IsNil(o.IsLarge) {
		var ret bool
		return ret
	}
	return *o.IsLarge
}

// GetIsLargeOk returns a tuple with the IsLarge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetIsLargeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLarge) {
		return nil, false
	}
	return o.IsLarge, true
}

// HasIsLarge returns a boolean if a field has been set.
func (o *Company) HasIsLarge() bool {
	if o != nil && !IsNil(o.IsLarge) {
		return true
	}

	return false
}

// SetIsLarge gets a reference to the given bool and assigns it to the IsLarge field.
func (o *Company) SetIsLarge(v bool) {
	o.IsLarge = &v
}

// GetDefaultAutoBill returns the DefaultAutoBill field value if set, zero value otherwise.
func (o *Company) GetDefaultAutoBill() string {
	if o == nil || IsNil(o.DefaultAutoBill) {
		var ret string
		return ret
	}
	return *o.DefaultAutoBill
}

// GetDefaultAutoBillOk returns a tuple with the DefaultAutoBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetDefaultAutoBillOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAutoBill) {
		return nil, false
	}
	return o.DefaultAutoBill, true
}

// HasDefaultAutoBill returns a boolean if a field has been set.
func (o *Company) HasDefaultAutoBill() bool {
	if o != nil && !IsNil(o.DefaultAutoBill) {
		return true
	}

	return false
}

// SetDefaultAutoBill gets a reference to the given string and assigns it to the DefaultAutoBill field.
func (o *Company) SetDefaultAutoBill(v string) {
	o.DefaultAutoBill = &v
}

// GetMarkExpensesInvoiceable returns the MarkExpensesInvoiceable field value if set, zero value otherwise.
func (o *Company) GetMarkExpensesInvoiceable() bool {
	if o == nil || IsNil(o.MarkExpensesInvoiceable) {
		var ret bool
		return ret
	}
	return *o.MarkExpensesInvoiceable
}

// GetMarkExpensesInvoiceableOk returns a tuple with the MarkExpensesInvoiceable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetMarkExpensesInvoiceableOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkExpensesInvoiceable) {
		return nil, false
	}
	return o.MarkExpensesInvoiceable, true
}

// HasMarkExpensesInvoiceable returns a boolean if a field has been set.
func (o *Company) HasMarkExpensesInvoiceable() bool {
	if o != nil && !IsNil(o.MarkExpensesInvoiceable) {
		return true
	}

	return false
}

// SetMarkExpensesInvoiceable gets a reference to the given bool and assigns it to the MarkExpensesInvoiceable field.
func (o *Company) SetMarkExpensesInvoiceable(v bool) {
	o.MarkExpensesInvoiceable = &v
}

// GetMarkExpensesPaid returns the MarkExpensesPaid field value if set, zero value otherwise.
func (o *Company) GetMarkExpensesPaid() bool {
	if o == nil || IsNil(o.MarkExpensesPaid) {
		var ret bool
		return ret
	}
	return *o.MarkExpensesPaid
}

// GetMarkExpensesPaidOk returns a tuple with the MarkExpensesPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetMarkExpensesPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkExpensesPaid) {
		return nil, false
	}
	return o.MarkExpensesPaid, true
}

// HasMarkExpensesPaid returns a boolean if a field has been set.
func (o *Company) HasMarkExpensesPaid() bool {
	if o != nil && !IsNil(o.MarkExpensesPaid) {
		return true
	}

	return false
}

// SetMarkExpensesPaid gets a reference to the given bool and assigns it to the MarkExpensesPaid field.
func (o *Company) SetMarkExpensesPaid(v bool) {
	o.MarkExpensesPaid = &v
}

// GetInvoiceExpenseDocuments returns the InvoiceExpenseDocuments field value if set, zero value otherwise.
func (o *Company) GetInvoiceExpenseDocuments() bool {
	if o == nil || IsNil(o.InvoiceExpenseDocuments) {
		var ret bool
		return ret
	}
	return *o.InvoiceExpenseDocuments
}

// GetInvoiceExpenseDocumentsOk returns a tuple with the InvoiceExpenseDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceExpenseDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceExpenseDocuments) {
		return nil, false
	}
	return o.InvoiceExpenseDocuments, true
}

// HasInvoiceExpenseDocuments returns a boolean if a field has been set.
func (o *Company) HasInvoiceExpenseDocuments() bool {
	if o != nil && !IsNil(o.InvoiceExpenseDocuments) {
		return true
	}

	return false
}

// SetInvoiceExpenseDocuments gets a reference to the given bool and assigns it to the InvoiceExpenseDocuments field.
func (o *Company) SetInvoiceExpenseDocuments(v bool) {
	o.InvoiceExpenseDocuments = &v
}

// GetAutoStartTasks returns the AutoStartTasks field value if set, zero value otherwise.
func (o *Company) GetAutoStartTasks() bool {
	if o == nil || IsNil(o.AutoStartTasks) {
		var ret bool
		return ret
	}
	return *o.AutoStartTasks
}

// GetAutoStartTasksOk returns a tuple with the AutoStartTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetAutoStartTasksOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoStartTasks) {
		return nil, false
	}
	return o.AutoStartTasks, true
}

// HasAutoStartTasks returns a boolean if a field has been set.
func (o *Company) HasAutoStartTasks() bool {
	if o != nil && !IsNil(o.AutoStartTasks) {
		return true
	}

	return false
}

// SetAutoStartTasks gets a reference to the given bool and assigns it to the AutoStartTasks field.
func (o *Company) SetAutoStartTasks(v bool) {
	o.AutoStartTasks = &v
}

// GetInvoiceTaskTimelog returns the InvoiceTaskTimelog field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskTimelog() bool {
	if o == nil || IsNil(o.InvoiceTaskTimelog) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskTimelog
}

// GetInvoiceTaskTimelogOk returns a tuple with the InvoiceTaskTimelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskTimelogOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskTimelog) {
		return nil, false
	}
	return o.InvoiceTaskTimelog, true
}

// HasInvoiceTaskTimelog returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskTimelog() bool {
	if o != nil && !IsNil(o.InvoiceTaskTimelog) {
		return true
	}

	return false
}

// SetInvoiceTaskTimelog gets a reference to the given bool and assigns it to the InvoiceTaskTimelog field.
func (o *Company) SetInvoiceTaskTimelog(v bool) {
	o.InvoiceTaskTimelog = &v
}

// GetInvoiceTaskDocuments returns the InvoiceTaskDocuments field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskDocuments() bool {
	if o == nil || IsNil(o.InvoiceTaskDocuments) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskDocuments
}

// GetInvoiceTaskDocumentsOk returns a tuple with the InvoiceTaskDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskDocuments) {
		return nil, false
	}
	return o.InvoiceTaskDocuments, true
}

// HasInvoiceTaskDocuments returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskDocuments() bool {
	if o != nil && !IsNil(o.InvoiceTaskDocuments) {
		return true
	}

	return false
}

// SetInvoiceTaskDocuments gets a reference to the given bool and assigns it to the InvoiceTaskDocuments field.
func (o *Company) SetInvoiceTaskDocuments(v bool) {
	o.InvoiceTaskDocuments = &v
}

// GetShowTasksTable returns the ShowTasksTable field value if set, zero value otherwise.
func (o *Company) GetShowTasksTable() bool {
	if o == nil || IsNil(o.ShowTasksTable) {
		var ret bool
		return ret
	}
	return *o.ShowTasksTable
}

// GetShowTasksTableOk returns a tuple with the ShowTasksTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetShowTasksTableOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowTasksTable) {
		return nil, false
	}
	return o.ShowTasksTable, true
}

// HasShowTasksTable returns a boolean if a field has been set.
func (o *Company) HasShowTasksTable() bool {
	if o != nil && !IsNil(o.ShowTasksTable) {
		return true
	}

	return false
}

// SetShowTasksTable gets a reference to the given bool and assigns it to the ShowTasksTable field.
func (o *Company) SetShowTasksTable(v bool) {
	o.ShowTasksTable = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *Company) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *Company) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *Company) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetDefaultTaskIsDateBased returns the DefaultTaskIsDateBased field value if set, zero value otherwise.
func (o *Company) GetDefaultTaskIsDateBased() bool {
	if o == nil || IsNil(o.DefaultTaskIsDateBased) {
		var ret bool
		return ret
	}
	return *o.DefaultTaskIsDateBased
}

// GetDefaultTaskIsDateBasedOk returns a tuple with the DefaultTaskIsDateBased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetDefaultTaskIsDateBasedOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultTaskIsDateBased) {
		return nil, false
	}
	return o.DefaultTaskIsDateBased, true
}

// HasDefaultTaskIsDateBased returns a boolean if a field has been set.
func (o *Company) HasDefaultTaskIsDateBased() bool {
	if o != nil && !IsNil(o.DefaultTaskIsDateBased) {
		return true
	}

	return false
}

// SetDefaultTaskIsDateBased gets a reference to the given bool and assigns it to the DefaultTaskIsDateBased field.
func (o *Company) SetDefaultTaskIsDateBased(v bool) {
	o.DefaultTaskIsDateBased = &v
}

// GetEnableProductDiscount returns the EnableProductDiscount field value if set, zero value otherwise.
func (o *Company) GetEnableProductDiscount() bool {
	if o == nil || IsNil(o.EnableProductDiscount) {
		var ret bool
		return ret
	}
	return *o.EnableProductDiscount
}

// GetEnableProductDiscountOk returns a tuple with the EnableProductDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnableProductDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableProductDiscount) {
		return nil, false
	}
	return o.EnableProductDiscount, true
}

// HasEnableProductDiscount returns a boolean if a field has been set.
func (o *Company) HasEnableProductDiscount() bool {
	if o != nil && !IsNil(o.EnableProductDiscount) {
		return true
	}

	return false
}

// SetEnableProductDiscount gets a reference to the given bool and assigns it to the EnableProductDiscount field.
func (o *Company) SetEnableProductDiscount(v bool) {
	o.EnableProductDiscount = &v
}

// GetCalculateExpenseTaxByAmount returns the CalculateExpenseTaxByAmount field value if set, zero value otherwise.
func (o *Company) GetCalculateExpenseTaxByAmount() bool {
	if o == nil || IsNil(o.CalculateExpenseTaxByAmount) {
		var ret bool
		return ret
	}
	return *o.CalculateExpenseTaxByAmount
}

// GetCalculateExpenseTaxByAmountOk returns a tuple with the CalculateExpenseTaxByAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCalculateExpenseTaxByAmountOk() (*bool, bool) {
	if o == nil || IsNil(o.CalculateExpenseTaxByAmount) {
		return nil, false
	}
	return o.CalculateExpenseTaxByAmount, true
}

// HasCalculateExpenseTaxByAmount returns a boolean if a field has been set.
func (o *Company) HasCalculateExpenseTaxByAmount() bool {
	if o != nil && !IsNil(o.CalculateExpenseTaxByAmount) {
		return true
	}

	return false
}

// SetCalculateExpenseTaxByAmount gets a reference to the given bool and assigns it to the CalculateExpenseTaxByAmount field.
func (o *Company) SetCalculateExpenseTaxByAmount(v bool) {
	o.CalculateExpenseTaxByAmount = &v
}

// GetExpenseInclusiveTaxes returns the ExpenseInclusiveTaxes field value if set, zero value otherwise.
func (o *Company) GetExpenseInclusiveTaxes() bool {
	if o == nil || IsNil(o.ExpenseInclusiveTaxes) {
		var ret bool
		return ret
	}
	return *o.ExpenseInclusiveTaxes
}

// GetExpenseInclusiveTaxesOk returns a tuple with the ExpenseInclusiveTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetExpenseInclusiveTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpenseInclusiveTaxes) {
		return nil, false
	}
	return o.ExpenseInclusiveTaxes, true
}

// HasExpenseInclusiveTaxes returns a boolean if a field has been set.
func (o *Company) HasExpenseInclusiveTaxes() bool {
	if o != nil && !IsNil(o.ExpenseInclusiveTaxes) {
		return true
	}

	return false
}

// SetExpenseInclusiveTaxes gets a reference to the given bool and assigns it to the ExpenseInclusiveTaxes field.
func (o *Company) SetExpenseInclusiveTaxes(v bool) {
	o.ExpenseInclusiveTaxes = &v
}

// GetSessionTimeout returns the SessionTimeout field value if set, zero value otherwise.
func (o *Company) GetSessionTimeout() int32 {
	if o == nil || IsNil(o.SessionTimeout) {
		var ret int32
		return ret
	}
	return *o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSessionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.SessionTimeout) {
		return nil, false
	}
	return o.SessionTimeout, true
}

// HasSessionTimeout returns a boolean if a field has been set.
func (o *Company) HasSessionTimeout() bool {
	if o != nil && !IsNil(o.SessionTimeout) {
		return true
	}

	return false
}

// SetSessionTimeout gets a reference to the given int32 and assigns it to the SessionTimeout field.
func (o *Company) SetSessionTimeout(v int32) {
	o.SessionTimeout = &v
}

// GetOauthPasswordRequired returns the OauthPasswordRequired field value if set, zero value otherwise.
func (o *Company) GetOauthPasswordRequired() bool {
	if o == nil || IsNil(o.OauthPasswordRequired) {
		var ret bool
		return ret
	}
	return *o.OauthPasswordRequired
}

// GetOauthPasswordRequiredOk returns a tuple with the OauthPasswordRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetOauthPasswordRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.OauthPasswordRequired) {
		return nil, false
	}
	return o.OauthPasswordRequired, true
}

// HasOauthPasswordRequired returns a boolean if a field has been set.
func (o *Company) HasOauthPasswordRequired() bool {
	if o != nil && !IsNil(o.OauthPasswordRequired) {
		return true
	}

	return false
}

// SetOauthPasswordRequired gets a reference to the given bool and assigns it to the OauthPasswordRequired field.
func (o *Company) SetOauthPasswordRequired(v bool) {
	o.OauthPasswordRequired = &v
}

// GetInvoiceTaskDatelog returns the InvoiceTaskDatelog field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskDatelog() bool {
	if o == nil || IsNil(o.InvoiceTaskDatelog) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskDatelog
}

// GetInvoiceTaskDatelogOk returns a tuple with the InvoiceTaskDatelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskDatelogOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskDatelog) {
		return nil, false
	}
	return o.InvoiceTaskDatelog, true
}

// HasInvoiceTaskDatelog returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskDatelog() bool {
	if o != nil && !IsNil(o.InvoiceTaskDatelog) {
		return true
	}

	return false
}

// SetInvoiceTaskDatelog gets a reference to the given bool and assigns it to the InvoiceTaskDatelog field.
func (o *Company) SetInvoiceTaskDatelog(v bool) {
	o.InvoiceTaskDatelog = &v
}

// GetDefaultPasswordTimeout returns the DefaultPasswordTimeout field value if set, zero value otherwise.
func (o *Company) GetDefaultPasswordTimeout() int32 {
	if o == nil || IsNil(o.DefaultPasswordTimeout) {
		var ret int32
		return ret
	}
	return *o.DefaultPasswordTimeout
}

// GetDefaultPasswordTimeoutOk returns a tuple with the DefaultPasswordTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetDefaultPasswordTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultPasswordTimeout) {
		return nil, false
	}
	return o.DefaultPasswordTimeout, true
}

// HasDefaultPasswordTimeout returns a boolean if a field has been set.
func (o *Company) HasDefaultPasswordTimeout() bool {
	if o != nil && !IsNil(o.DefaultPasswordTimeout) {
		return true
	}

	return false
}

// SetDefaultPasswordTimeout gets a reference to the given int32 and assigns it to the DefaultPasswordTimeout field.
func (o *Company) SetDefaultPasswordTimeout(v int32) {
	o.DefaultPasswordTimeout = &v
}

// GetShowTaskEndDate returns the ShowTaskEndDate field value if set, zero value otherwise.
func (o *Company) GetShowTaskEndDate() bool {
	if o == nil || IsNil(o.ShowTaskEndDate) {
		var ret bool
		return ret
	}
	return *o.ShowTaskEndDate
}

// GetShowTaskEndDateOk returns a tuple with the ShowTaskEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetShowTaskEndDateOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowTaskEndDate) {
		return nil, false
	}
	return o.ShowTaskEndDate, true
}

// HasShowTaskEndDate returns a boolean if a field has been set.
func (o *Company) HasShowTaskEndDate() bool {
	if o != nil && !IsNil(o.ShowTaskEndDate) {
		return true
	}

	return false
}

// SetShowTaskEndDate gets a reference to the given bool and assigns it to the ShowTaskEndDate field.
func (o *Company) SetShowTaskEndDate(v bool) {
	o.ShowTaskEndDate = &v
}

// GetMarkdownEnabled returns the MarkdownEnabled field value if set, zero value otherwise.
func (o *Company) GetMarkdownEnabled() bool {
	if o == nil || IsNil(o.MarkdownEnabled) {
		var ret bool
		return ret
	}
	return *o.MarkdownEnabled
}

// GetMarkdownEnabledOk returns a tuple with the MarkdownEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetMarkdownEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkdownEnabled) {
		return nil, false
	}
	return o.MarkdownEnabled, true
}

// HasMarkdownEnabled returns a boolean if a field has been set.
func (o *Company) HasMarkdownEnabled() bool {
	if o != nil && !IsNil(o.MarkdownEnabled) {
		return true
	}

	return false
}

// SetMarkdownEnabled gets a reference to the given bool and assigns it to the MarkdownEnabled field.
func (o *Company) SetMarkdownEnabled(v bool) {
	o.MarkdownEnabled = &v
}

// GetReportIncludeDrafts returns the ReportIncludeDrafts field value if set, zero value otherwise.
func (o *Company) GetReportIncludeDrafts() bool {
	if o == nil || IsNil(o.ReportIncludeDrafts) {
		var ret bool
		return ret
	}
	return *o.ReportIncludeDrafts
}

// GetReportIncludeDraftsOk returns a tuple with the ReportIncludeDrafts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetReportIncludeDraftsOk() (*bool, bool) {
	if o == nil || IsNil(o.ReportIncludeDrafts) {
		return nil, false
	}
	return o.ReportIncludeDrafts, true
}

// HasReportIncludeDrafts returns a boolean if a field has been set.
func (o *Company) HasReportIncludeDrafts() bool {
	if o != nil && !IsNil(o.ReportIncludeDrafts) {
		return true
	}

	return false
}

// SetReportIncludeDrafts gets a reference to the given bool and assigns it to the ReportIncludeDrafts field.
func (o *Company) SetReportIncludeDrafts(v bool) {
	o.ReportIncludeDrafts = &v
}

// GetClientRegistrationFields returns the ClientRegistrationFields field value if set, zero value otherwise.
func (o *Company) GetClientRegistrationFields() map[string]any {
	if o == nil || IsNil(o.ClientRegistrationFields) {
		var ret map[string]any
		return ret
	}
	return o.ClientRegistrationFields
}

// GetClientRegistrationFieldsOk returns a tuple with the ClientRegistrationFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetClientRegistrationFieldsOk() (map[string]any, bool) {
	if o == nil || IsNil(o.ClientRegistrationFields) {
		return map[string]any{}, false
	}
	return o.ClientRegistrationFields, true
}

// HasClientRegistrationFields returns a boolean if a field has been set.
func (o *Company) HasClientRegistrationFields() bool {
	if o != nil && !IsNil(o.ClientRegistrationFields) {
		return true
	}

	return false
}

// SetClientRegistrationFields gets a reference to the given map[string]interface{} and assigns it to the ClientRegistrationFields field.
func (o *Company) SetClientRegistrationFields(v map[string]any) {
	o.ClientRegistrationFields = v
}

// GetStopOnUnpaidRecurring returns the StopOnUnpaidRecurring field value if set, zero value otherwise.
func (o *Company) GetStopOnUnpaidRecurring() bool {
	if o == nil || IsNil(o.StopOnUnpaidRecurring) {
		var ret bool
		return ret
	}
	return *o.StopOnUnpaidRecurring
}

// GetStopOnUnpaidRecurringOk returns a tuple with the StopOnUnpaidRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetStopOnUnpaidRecurringOk() (*bool, bool) {
	if o == nil || IsNil(o.StopOnUnpaidRecurring) {
		return nil, false
	}
	return o.StopOnUnpaidRecurring, true
}

// HasStopOnUnpaidRecurring returns a boolean if a field has been set.
func (o *Company) HasStopOnUnpaidRecurring() bool {
	if o != nil && !IsNil(o.StopOnUnpaidRecurring) {
		return true
	}

	return false
}

// SetStopOnUnpaidRecurring gets a reference to the given bool and assigns it to the StopOnUnpaidRecurring field.
func (o *Company) SetStopOnUnpaidRecurring(v bool) {
	o.StopOnUnpaidRecurring = &v
}

// GetUseQuoteTermsOnConversion returns the UseQuoteTermsOnConversion field value if set, zero value otherwise.
func (o *Company) GetUseQuoteTermsOnConversion() bool {
	if o == nil || IsNil(o.UseQuoteTermsOnConversion) {
		var ret bool
		return ret
	}
	return *o.UseQuoteTermsOnConversion
}

// GetUseQuoteTermsOnConversionOk returns a tuple with the UseQuoteTermsOnConversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetUseQuoteTermsOnConversionOk() (*bool, bool) {
	if o == nil || IsNil(o.UseQuoteTermsOnConversion) {
		return nil, false
	}
	return o.UseQuoteTermsOnConversion, true
}

// HasUseQuoteTermsOnConversion returns a boolean if a field has been set.
func (o *Company) HasUseQuoteTermsOnConversion() bool {
	if o != nil && !IsNil(o.UseQuoteTermsOnConversion) {
		return true
	}

	return false
}

// SetUseQuoteTermsOnConversion gets a reference to the given bool and assigns it to the UseQuoteTermsOnConversion field.
func (o *Company) SetUseQuoteTermsOnConversion(v bool) {
	o.UseQuoteTermsOnConversion = &v
}

// GetEnableApplyingPayments returns the EnableApplyingPayments field value if set, zero value otherwise.
func (o *Company) GetEnableApplyingPayments() bool {
	if o == nil || IsNil(o.EnableApplyingPayments) {
		var ret bool
		return ret
	}
	return *o.EnableApplyingPayments
}

// GetEnableApplyingPaymentsOk returns a tuple with the EnableApplyingPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnableApplyingPaymentsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableApplyingPayments) {
		return nil, false
	}
	return o.EnableApplyingPayments, true
}

// HasEnableApplyingPayments returns a boolean if a field has been set.
func (o *Company) HasEnableApplyingPayments() bool {
	if o != nil && !IsNil(o.EnableApplyingPayments) {
		return true
	}

	return false
}

// SetEnableApplyingPayments gets a reference to the given bool and assigns it to the EnableApplyingPayments field.
func (o *Company) SetEnableApplyingPayments(v bool) {
	o.EnableApplyingPayments = &v
}

// GetTrackInventory returns the TrackInventory field value if set, zero value otherwise.
func (o *Company) GetTrackInventory() bool {
	if o == nil || IsNil(o.TrackInventory) {
		var ret bool
		return ret
	}
	return *o.TrackInventory
}

// GetTrackInventoryOk returns a tuple with the TrackInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetTrackInventoryOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackInventory) {
		return nil, false
	}
	return o.TrackInventory, true
}

// HasTrackInventory returns a boolean if a field has been set.
func (o *Company) HasTrackInventory() bool {
	if o != nil && !IsNil(o.TrackInventory) {
		return true
	}

	return false
}

// SetTrackInventory gets a reference to the given bool and assigns it to the TrackInventory field.
func (o *Company) SetTrackInventory(v bool) {
	o.TrackInventory = &v
}

// GetInventoryNotificationThreshold returns the InventoryNotificationThreshold field value if set, zero value otherwise.
func (o *Company) GetInventoryNotificationThreshold() int32 {
	if o == nil || IsNil(o.InventoryNotificationThreshold) {
		var ret int32
		return ret
	}
	return *o.InventoryNotificationThreshold
}

// GetInventoryNotificationThresholdOk returns a tuple with the InventoryNotificationThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInventoryNotificationThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.InventoryNotificationThreshold) {
		return nil, false
	}
	return o.InventoryNotificationThreshold, true
}

// HasInventoryNotificationThreshold returns a boolean if a field has been set.
func (o *Company) HasInventoryNotificationThreshold() bool {
	if o != nil && !IsNil(o.InventoryNotificationThreshold) {
		return true
	}

	return false
}

// SetInventoryNotificationThreshold gets a reference to the given int32 and assigns it to the InventoryNotificationThreshold field.
func (o *Company) SetInventoryNotificationThreshold(v int32) {
	o.InventoryNotificationThreshold = &v
}

// GetStockNotification returns the StockNotification field value if set, zero value otherwise.
func (o *Company) GetStockNotification() bool {
	if o == nil || IsNil(o.StockNotification) {
		var ret bool
		return ret
	}
	return *o.StockNotification
}

// GetStockNotificationOk returns a tuple with the StockNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetStockNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.StockNotification) {
		return nil, false
	}
	return o.StockNotification, true
}

// HasStockNotification returns a boolean if a field has been set.
func (o *Company) HasStockNotification() bool {
	if o != nil && !IsNil(o.StockNotification) {
		return true
	}

	return false
}

// SetStockNotification gets a reference to the given bool and assigns it to the StockNotification field.
func (o *Company) SetStockNotification(v bool) {
	o.StockNotification = &v
}

// GetMatomoUrl returns the MatomoUrl field value if set, zero value otherwise.
func (o *Company) GetMatomoUrl() string {
	if o == nil || IsNil(o.MatomoUrl) {
		var ret string
		return ret
	}
	return *o.MatomoUrl
}

// GetMatomoUrlOk returns a tuple with the MatomoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetMatomoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MatomoUrl) {
		return nil, false
	}
	return o.MatomoUrl, true
}

// HasMatomoUrl returns a boolean if a field has been set.
func (o *Company) HasMatomoUrl() bool {
	if o != nil && !IsNil(o.MatomoUrl) {
		return true
	}

	return false
}

// SetMatomoUrl gets a reference to the given string and assigns it to the MatomoUrl field.
func (o *Company) SetMatomoUrl(v string) {
	o.MatomoUrl = &v
}

// GetMatomoId returns the MatomoId field value if set, zero value otherwise.
func (o *Company) GetMatomoId() string {
	if o == nil || IsNil(o.MatomoId) {
		var ret string
		return ret
	}
	return *o.MatomoId
}

// GetMatomoIdOk returns a tuple with the MatomoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetMatomoIdOk() (*string, bool) {
	if o == nil || IsNil(o.MatomoId) {
		return nil, false
	}
	return o.MatomoId, true
}

// HasMatomoId returns a boolean if a field has been set.
func (o *Company) HasMatomoId() bool {
	if o != nil && !IsNil(o.MatomoId) {
		return true
	}

	return false
}

// SetMatomoId gets a reference to the given string and assigns it to the MatomoId field.
func (o *Company) SetMatomoId(v string) {
	o.MatomoId = &v
}

// GetEnabledExpenseTaxRates returns the EnabledExpenseTaxRates field value if set, zero value otherwise.
func (o *Company) GetEnabledExpenseTaxRates() int32 {
	if o == nil || IsNil(o.EnabledExpenseTaxRates) {
		var ret int32
		return ret
	}
	return *o.EnabledExpenseTaxRates
}

// GetEnabledExpenseTaxRatesOk returns a tuple with the EnabledExpenseTaxRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEnabledExpenseTaxRatesOk() (*int32, bool) {
	if o == nil || IsNil(o.EnabledExpenseTaxRates) {
		return nil, false
	}
	return o.EnabledExpenseTaxRates, true
}

// HasEnabledExpenseTaxRates returns a boolean if a field has been set.
func (o *Company) HasEnabledExpenseTaxRates() bool {
	if o != nil && !IsNil(o.EnabledExpenseTaxRates) {
		return true
	}

	return false
}

// SetEnabledExpenseTaxRates gets a reference to the given int32 and assigns it to the EnabledExpenseTaxRates field.
func (o *Company) SetEnabledExpenseTaxRates(v int32) {
	o.EnabledExpenseTaxRates = &v
}

// GetInvoiceTaskProject returns the InvoiceTaskProject field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskProject() bool {
	if o == nil || IsNil(o.InvoiceTaskProject) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskProject
}

// GetInvoiceTaskProjectOk returns a tuple with the InvoiceTaskProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskProjectOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskProject) {
		return nil, false
	}
	return o.InvoiceTaskProject, true
}

// HasInvoiceTaskProject returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskProject() bool {
	if o != nil && !IsNil(o.InvoiceTaskProject) {
		return true
	}

	return false
}

// SetInvoiceTaskProject gets a reference to the given bool and assigns it to the InvoiceTaskProject field.
func (o *Company) SetInvoiceTaskProject(v bool) {
	o.InvoiceTaskProject = &v
}

// GetReportIncludeDeleted returns the ReportIncludeDeleted field value if set, zero value otherwise.
func (o *Company) GetReportIncludeDeleted() bool {
	if o == nil || IsNil(o.ReportIncludeDeleted) {
		var ret bool
		return ret
	}
	return *o.ReportIncludeDeleted
}

// GetReportIncludeDeletedOk returns a tuple with the ReportIncludeDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetReportIncludeDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.ReportIncludeDeleted) {
		return nil, false
	}
	return o.ReportIncludeDeleted, true
}

// HasReportIncludeDeleted returns a boolean if a field has been set.
func (o *Company) HasReportIncludeDeleted() bool {
	if o != nil && !IsNil(o.ReportIncludeDeleted) {
		return true
	}

	return false
}

// SetReportIncludeDeleted gets a reference to the given bool and assigns it to the ReportIncludeDeleted field.
func (o *Company) SetReportIncludeDeleted(v bool) {
	o.ReportIncludeDeleted = &v
}

// GetInvoiceTaskLock returns the InvoiceTaskLock field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskLock() bool {
	if o == nil || IsNil(o.InvoiceTaskLock) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskLock
}

// GetInvoiceTaskLockOk returns a tuple with the InvoiceTaskLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskLockOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskLock) {
		return nil, false
	}
	return o.InvoiceTaskLock, true
}

// HasInvoiceTaskLock returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskLock() bool {
	if o != nil && !IsNil(o.InvoiceTaskLock) {
		return true
	}

	return false
}

// SetInvoiceTaskLock gets a reference to the given bool and assigns it to the InvoiceTaskLock field.
func (o *Company) SetInvoiceTaskLock(v bool) {
	o.InvoiceTaskLock = &v
}

// GetConvertPaymentCurrency returns the ConvertPaymentCurrency field value if set, zero value otherwise.
func (o *Company) GetConvertPaymentCurrency() bool {
	if o == nil || IsNil(o.ConvertPaymentCurrency) {
		var ret bool
		return ret
	}
	return *o.ConvertPaymentCurrency
}

// GetConvertPaymentCurrencyOk returns a tuple with the ConvertPaymentCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetConvertPaymentCurrencyOk() (*bool, bool) {
	if o == nil || IsNil(o.ConvertPaymentCurrency) {
		return nil, false
	}
	return o.ConvertPaymentCurrency, true
}

// HasConvertPaymentCurrency returns a boolean if a field has been set.
func (o *Company) HasConvertPaymentCurrency() bool {
	if o != nil && !IsNil(o.ConvertPaymentCurrency) {
		return true
	}

	return false
}

// SetConvertPaymentCurrency gets a reference to the given bool and assigns it to the ConvertPaymentCurrency field.
func (o *Company) SetConvertPaymentCurrency(v bool) {
	o.ConvertPaymentCurrency = &v
}

// GetConvertExpenseCurrency returns the ConvertExpenseCurrency field value if set, zero value otherwise.
func (o *Company) GetConvertExpenseCurrency() bool {
	if o == nil || IsNil(o.ConvertExpenseCurrency) {
		var ret bool
		return ret
	}
	return *o.ConvertExpenseCurrency
}

// GetConvertExpenseCurrencyOk returns a tuple with the ConvertExpenseCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetConvertExpenseCurrencyOk() (*bool, bool) {
	if o == nil || IsNil(o.ConvertExpenseCurrency) {
		return nil, false
	}
	return o.ConvertExpenseCurrency, true
}

// HasConvertExpenseCurrency returns a boolean if a field has been set.
func (o *Company) HasConvertExpenseCurrency() bool {
	if o != nil && !IsNil(o.ConvertExpenseCurrency) {
		return true
	}

	return false
}

// SetConvertExpenseCurrency gets a reference to the given bool and assigns it to the ConvertExpenseCurrency field.
func (o *Company) SetConvertExpenseCurrency(v bool) {
	o.ConvertExpenseCurrency = &v
}

// GetNotifyVendorWhenPaid returns the NotifyVendorWhenPaid field value if set, zero value otherwise.
func (o *Company) GetNotifyVendorWhenPaid() bool {
	if o == nil || IsNil(o.NotifyVendorWhenPaid) {
		var ret bool
		return ret
	}
	return *o.NotifyVendorWhenPaid
}

// GetNotifyVendorWhenPaidOk returns a tuple with the NotifyVendorWhenPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetNotifyVendorWhenPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyVendorWhenPaid) {
		return nil, false
	}
	return o.NotifyVendorWhenPaid, true
}

// HasNotifyVendorWhenPaid returns a boolean if a field has been set.
func (o *Company) HasNotifyVendorWhenPaid() bool {
	if o != nil && !IsNil(o.NotifyVendorWhenPaid) {
		return true
	}

	return false
}

// SetNotifyVendorWhenPaid gets a reference to the given bool and assigns it to the NotifyVendorWhenPaid field.
func (o *Company) SetNotifyVendorWhenPaid(v bool) {
	o.NotifyVendorWhenPaid = &v
}

// GetInvoiceTaskHours returns the InvoiceTaskHours field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskHours() bool {
	if o == nil || IsNil(o.InvoiceTaskHours) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskHours
}

// GetInvoiceTaskHoursOk returns a tuple with the InvoiceTaskHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskHoursOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskHours) {
		return nil, false
	}
	return o.InvoiceTaskHours, true
}

// HasInvoiceTaskHours returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskHours() bool {
	if o != nil && !IsNil(o.InvoiceTaskHours) {
		return true
	}

	return false
}

// SetInvoiceTaskHours gets a reference to the given bool and assigns it to the InvoiceTaskHours field.
func (o *Company) SetInvoiceTaskHours(v bool) {
	o.InvoiceTaskHours = &v
}

// GetCalculateTaxes returns the CalculateTaxes field value if set, zero value otherwise.
func (o *Company) GetCalculateTaxes() bool {
	if o == nil || IsNil(o.CalculateTaxes) {
		var ret bool
		return ret
	}
	return *o.CalculateTaxes
}

// GetCalculateTaxesOk returns a tuple with the CalculateTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetCalculateTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.CalculateTaxes) {
		return nil, false
	}
	return o.CalculateTaxes, true
}

// HasCalculateTaxes returns a boolean if a field has been set.
func (o *Company) HasCalculateTaxes() bool {
	if o != nil && !IsNil(o.CalculateTaxes) {
		return true
	}

	return false
}

// SetCalculateTaxes gets a reference to the given bool and assigns it to the CalculateTaxes field.
func (o *Company) SetCalculateTaxes(v bool) {
	o.CalculateTaxes = &v
}

// GetTaxData returns the TaxData field value if set, zero value otherwise.
func (o *Company) GetTaxData() map[string]any {
	if o == nil || IsNil(o.TaxData) {
		var ret map[string]any
		return ret
	}
	return o.TaxData
}

// GetTaxDataOk returns a tuple with the TaxData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetTaxDataOk() (map[string]any, bool) {
	if o == nil || IsNil(o.TaxData) {
		return map[string]any{}, false
	}
	return o.TaxData, true
}

// HasTaxData returns a boolean if a field has been set.
func (o *Company) HasTaxData() bool {
	if o != nil && !IsNil(o.TaxData) {
		return true
	}

	return false
}

// SetTaxData gets a reference to the given map[string]interface{} and assigns it to the TaxData field.
func (o *Company) SetTaxData(v map[string]any) {
	o.TaxData = v
}

// GetEInvoiceCertificate returns the EInvoiceCertificate field value if set, zero value otherwise.
func (o *Company) GetEInvoiceCertificate() string {
	if o == nil || IsNil(o.EInvoiceCertificate) {
		var ret string
		return ret
	}
	return *o.EInvoiceCertificate
}

// GetEInvoiceCertificateOk returns a tuple with the EInvoiceCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEInvoiceCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.EInvoiceCertificate) {
		return nil, false
	}
	return o.EInvoiceCertificate, true
}

// HasEInvoiceCertificate returns a boolean if a field has been set.
func (o *Company) HasEInvoiceCertificate() bool {
	if o != nil && !IsNil(o.EInvoiceCertificate) {
		return true
	}

	return false
}

// SetEInvoiceCertificate gets a reference to the given string and assigns it to the EInvoiceCertificate field.
func (o *Company) SetEInvoiceCertificate(v string) {
	o.EInvoiceCertificate = &v
}

// GetEInvoiceCertificatePassphrase returns the EInvoiceCertificatePassphrase field value if set, zero value otherwise.
func (o *Company) GetEInvoiceCertificatePassphrase() string {
	if o == nil || IsNil(o.EInvoiceCertificatePassphrase) {
		var ret string
		return ret
	}
	return *o.EInvoiceCertificatePassphrase
}

// GetEInvoiceCertificatePassphraseOk returns a tuple with the EInvoiceCertificatePassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEInvoiceCertificatePassphraseOk() (*string, bool) {
	if o == nil || IsNil(o.EInvoiceCertificatePassphrase) {
		return nil, false
	}
	return o.EInvoiceCertificatePassphrase, true
}

// HasEInvoiceCertificatePassphrase returns a boolean if a field has been set.
func (o *Company) HasEInvoiceCertificatePassphrase() bool {
	if o != nil && !IsNil(o.EInvoiceCertificatePassphrase) {
		return true
	}

	return false
}

// SetEInvoiceCertificatePassphrase gets a reference to the given string and assigns it to the EInvoiceCertificatePassphrase field.
func (o *Company) SetEInvoiceCertificatePassphrase(v string) {
	o.EInvoiceCertificatePassphrase = &v
}

// GetOriginTaxData returns the OriginTaxData field value if set, zero value otherwise.
func (o *Company) GetOriginTaxData() map[string]any {
	if o == nil || IsNil(o.OriginTaxData) {
		var ret map[string]any
		return ret
	}
	return o.OriginTaxData
}

// GetOriginTaxDataOk returns a tuple with the OriginTaxData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetOriginTaxDataOk() (map[string]any, bool) {
	if o == nil || IsNil(o.OriginTaxData) {
		return map[string]any{}, false
	}
	return o.OriginTaxData, true
}

// HasOriginTaxData returns a boolean if a field has been set.
func (o *Company) HasOriginTaxData() bool {
	if o != nil && !IsNil(o.OriginTaxData) {
		return true
	}

	return false
}

// SetOriginTaxData gets a reference to the given map[string]interface{} and assigns it to the OriginTaxData field.
func (o *Company) SetOriginTaxData(v map[string]any) {
	o.OriginTaxData = v
}

// GetInvoiceTaskProjectHeader returns the InvoiceTaskProjectHeader field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskProjectHeader() bool {
	if o == nil || IsNil(o.InvoiceTaskProjectHeader) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskProjectHeader
}

// GetInvoiceTaskProjectHeaderOk returns a tuple with the InvoiceTaskProjectHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskProjectHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskProjectHeader) {
		return nil, false
	}
	return o.InvoiceTaskProjectHeader, true
}

// HasInvoiceTaskProjectHeader returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskProjectHeader() bool {
	if o != nil && !IsNil(o.InvoiceTaskProjectHeader) {
		return true
	}

	return false
}

// SetInvoiceTaskProjectHeader gets a reference to the given bool and assigns it to the InvoiceTaskProjectHeader field.
func (o *Company) SetInvoiceTaskProjectHeader(v bool) {
	o.InvoiceTaskProjectHeader = &v
}

// GetInvoiceTaskItemDescription returns the InvoiceTaskItemDescription field value if set, zero value otherwise.
func (o *Company) GetInvoiceTaskItemDescription() bool {
	if o == nil || IsNil(o.InvoiceTaskItemDescription) {
		var ret bool
		return ret
	}
	return *o.InvoiceTaskItemDescription
}

// GetInvoiceTaskItemDescriptionOk returns a tuple with the InvoiceTaskItemDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInvoiceTaskItemDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceTaskItemDescription) {
		return nil, false
	}
	return o.InvoiceTaskItemDescription, true
}

// HasInvoiceTaskItemDescription returns a boolean if a field has been set.
func (o *Company) HasInvoiceTaskItemDescription() bool {
	if o != nil && !IsNil(o.InvoiceTaskItemDescription) {
		return true
	}

	return false
}

// SetInvoiceTaskItemDescription gets a reference to the given bool and assigns it to the InvoiceTaskItemDescription field.
func (o *Company) SetInvoiceTaskItemDescription(v bool) {
	o.InvoiceTaskItemDescription = &v
}

// GetExpenseMailbox returns the ExpenseMailbox field value if set, zero value otherwise.
func (o *Company) GetExpenseMailbox() string {
	if o == nil || IsNil(o.ExpenseMailbox) {
		var ret string
		return ret
	}
	return *o.ExpenseMailbox
}

// GetExpenseMailboxOk returns a tuple with the ExpenseMailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetExpenseMailboxOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseMailbox) {
		return nil, false
	}
	return o.ExpenseMailbox, true
}

// HasExpenseMailbox returns a boolean if a field has been set.
func (o *Company) HasExpenseMailbox() bool {
	if o != nil && !IsNil(o.ExpenseMailbox) {
		return true
	}

	return false
}

// SetExpenseMailbox gets a reference to the given string and assigns it to the ExpenseMailbox field.
func (o *Company) SetExpenseMailbox(v string) {
	o.ExpenseMailbox = &v
}

// GetExpenseMailboxActive returns the ExpenseMailboxActive field value if set, zero value otherwise.
func (o *Company) GetExpenseMailboxActive() bool {
	if o == nil || IsNil(o.ExpenseMailboxActive) {
		var ret bool
		return ret
	}
	return *o.ExpenseMailboxActive
}

// GetExpenseMailboxActiveOk returns a tuple with the ExpenseMailboxActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetExpenseMailboxActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpenseMailboxActive) {
		return nil, false
	}
	return o.ExpenseMailboxActive, true
}

// HasExpenseMailboxActive returns a boolean if a field has been set.
func (o *Company) HasExpenseMailboxActive() bool {
	if o != nil && !IsNil(o.ExpenseMailboxActive) {
		return true
	}

	return false
}

// SetExpenseMailboxActive gets a reference to the given bool and assigns it to the ExpenseMailboxActive field.
func (o *Company) SetExpenseMailboxActive(v bool) {
	o.ExpenseMailboxActive = &v
}

// GetInboundMailboxAllowCompanyUsers returns the InboundMailboxAllowCompanyUsers field value if set, zero value otherwise.
func (o *Company) GetInboundMailboxAllowCompanyUsers() bool {
	if o == nil || IsNil(o.InboundMailboxAllowCompanyUsers) {
		var ret bool
		return ret
	}
	return *o.InboundMailboxAllowCompanyUsers
}

// GetInboundMailboxAllowCompanyUsersOk returns a tuple with the InboundMailboxAllowCompanyUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInboundMailboxAllowCompanyUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.InboundMailboxAllowCompanyUsers) {
		return nil, false
	}
	return o.InboundMailboxAllowCompanyUsers, true
}

// HasInboundMailboxAllowCompanyUsers returns a boolean if a field has been set.
func (o *Company) HasInboundMailboxAllowCompanyUsers() bool {
	if o != nil && !IsNil(o.InboundMailboxAllowCompanyUsers) {
		return true
	}

	return false
}

// SetInboundMailboxAllowCompanyUsers gets a reference to the given bool and assigns it to the InboundMailboxAllowCompanyUsers field.
func (o *Company) SetInboundMailboxAllowCompanyUsers(v bool) {
	o.InboundMailboxAllowCompanyUsers = &v
}

// GetInboundMailboxAllowVendors returns the InboundMailboxAllowVendors field value if set, zero value otherwise.
func (o *Company) GetInboundMailboxAllowVendors() bool {
	if o == nil || IsNil(o.InboundMailboxAllowVendors) {
		var ret bool
		return ret
	}
	return *o.InboundMailboxAllowVendors
}

// GetInboundMailboxAllowVendorsOk returns a tuple with the InboundMailboxAllowVendors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInboundMailboxAllowVendorsOk() (*bool, bool) {
	if o == nil || IsNil(o.InboundMailboxAllowVendors) {
		return nil, false
	}
	return o.InboundMailboxAllowVendors, true
}

// HasInboundMailboxAllowVendors returns a boolean if a field has been set.
func (o *Company) HasInboundMailboxAllowVendors() bool {
	if o != nil && !IsNil(o.InboundMailboxAllowVendors) {
		return true
	}

	return false
}

// SetInboundMailboxAllowVendors gets a reference to the given bool and assigns it to the InboundMailboxAllowVendors field.
func (o *Company) SetInboundMailboxAllowVendors(v bool) {
	o.InboundMailboxAllowVendors = &v
}

// GetInboundMailboxAllowClients returns the InboundMailboxAllowClients field value if set, zero value otherwise.
func (o *Company) GetInboundMailboxAllowClients() bool {
	if o == nil || IsNil(o.InboundMailboxAllowClients) {
		var ret bool
		return ret
	}
	return *o.InboundMailboxAllowClients
}

// GetInboundMailboxAllowClientsOk returns a tuple with the InboundMailboxAllowClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInboundMailboxAllowClientsOk() (*bool, bool) {
	if o == nil || IsNil(o.InboundMailboxAllowClients) {
		return nil, false
	}
	return o.InboundMailboxAllowClients, true
}

// HasInboundMailboxAllowClients returns a boolean if a field has been set.
func (o *Company) HasInboundMailboxAllowClients() bool {
	if o != nil && !IsNil(o.InboundMailboxAllowClients) {
		return true
	}

	return false
}

// SetInboundMailboxAllowClients gets a reference to the given bool and assigns it to the InboundMailboxAllowClients field.
func (o *Company) SetInboundMailboxAllowClients(v bool) {
	o.InboundMailboxAllowClients = &v
}

// GetInboundMailboxAllowUnknown returns the InboundMailboxAllowUnknown field value if set, zero value otherwise.
func (o *Company) GetInboundMailboxAllowUnknown() bool {
	if o == nil || IsNil(o.InboundMailboxAllowUnknown) {
		var ret bool
		return ret
	}
	return *o.InboundMailboxAllowUnknown
}

// GetInboundMailboxAllowUnknownOk returns a tuple with the InboundMailboxAllowUnknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInboundMailboxAllowUnknownOk() (*bool, bool) {
	if o == nil || IsNil(o.InboundMailboxAllowUnknown) {
		return nil, false
	}
	return o.InboundMailboxAllowUnknown, true
}

// HasInboundMailboxAllowUnknown returns a boolean if a field has been set.
func (o *Company) HasInboundMailboxAllowUnknown() bool {
	if o != nil && !IsNil(o.InboundMailboxAllowUnknown) {
		return true
	}

	return false
}

// SetInboundMailboxAllowUnknown gets a reference to the given bool and assigns it to the InboundMailboxAllowUnknown field.
func (o *Company) SetInboundMailboxAllowUnknown(v bool) {
	o.InboundMailboxAllowUnknown = &v
}

// GetInboundMailboxWhitelist returns the InboundMailboxWhitelist field value if set, zero value otherwise.
func (o *Company) GetInboundMailboxWhitelist() string {
	if o == nil || IsNil(o.InboundMailboxWhitelist) {
		var ret string
		return ret
	}
	return *o.InboundMailboxWhitelist
}

// GetInboundMailboxWhitelistOk returns a tuple with the InboundMailboxWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInboundMailboxWhitelistOk() (*string, bool) {
	if o == nil || IsNil(o.InboundMailboxWhitelist) {
		return nil, false
	}
	return o.InboundMailboxWhitelist, true
}

// HasInboundMailboxWhitelist returns a boolean if a field has been set.
func (o *Company) HasInboundMailboxWhitelist() bool {
	if o != nil && !IsNil(o.InboundMailboxWhitelist) {
		return true
	}

	return false
}

// SetInboundMailboxWhitelist gets a reference to the given string and assigns it to the InboundMailboxWhitelist field.
func (o *Company) SetInboundMailboxWhitelist(v string) {
	o.InboundMailboxWhitelist = &v
}

// GetInboundMailboxBlacklist returns the InboundMailboxBlacklist field value if set, zero value otherwise.
func (o *Company) GetInboundMailboxBlacklist() string {
	if o == nil || IsNil(o.InboundMailboxBlacklist) {
		var ret string
		return ret
	}
	return *o.InboundMailboxBlacklist
}

// GetInboundMailboxBlacklistOk returns a tuple with the InboundMailboxBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetInboundMailboxBlacklistOk() (*string, bool) {
	if o == nil || IsNil(o.InboundMailboxBlacklist) {
		return nil, false
	}
	return o.InboundMailboxBlacklist, true
}

// HasInboundMailboxBlacklist returns a boolean if a field has been set.
func (o *Company) HasInboundMailboxBlacklist() bool {
	if o != nil && !IsNil(o.InboundMailboxBlacklist) {
		return true
	}

	return false
}

// SetInboundMailboxBlacklist gets a reference to the given string and assigns it to the InboundMailboxBlacklist field.
func (o *Company) SetInboundMailboxBlacklist(v string) {
	o.InboundMailboxBlacklist = &v
}

// GetSmtpHost returns the SmtpHost field value if set, zero value otherwise.
func (o *Company) GetSmtpHost() string {
	if o == nil || IsNil(o.SmtpHost) {
		var ret string
		return ret
	}
	return *o.SmtpHost
}

// GetSmtpHostOk returns a tuple with the SmtpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSmtpHostOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpHost) {
		return nil, false
	}
	return o.SmtpHost, true
}

// HasSmtpHost returns a boolean if a field has been set.
func (o *Company) HasSmtpHost() bool {
	if o != nil && !IsNil(o.SmtpHost) {
		return true
	}

	return false
}

// SetSmtpHost gets a reference to the given string and assigns it to the SmtpHost field.
func (o *Company) SetSmtpHost(v string) {
	o.SmtpHost = &v
}

// GetSmtpPort returns the SmtpPort field value if set, zero value otherwise.
func (o *Company) GetSmtpPort() int32 {
	if o == nil || IsNil(o.SmtpPort) {
		var ret int32
		return ret
	}
	return *o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSmtpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SmtpPort) {
		return nil, false
	}
	return o.SmtpPort, true
}

// HasSmtpPort returns a boolean if a field has been set.
func (o *Company) HasSmtpPort() bool {
	if o != nil && !IsNil(o.SmtpPort) {
		return true
	}

	return false
}

// SetSmtpPort gets a reference to the given int32 and assigns it to the SmtpPort field.
func (o *Company) SetSmtpPort(v int32) {
	o.SmtpPort = &v
}

// GetSmtpEncryption returns the SmtpEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Company) GetSmtpEncryption() string {
	if o == nil || IsNil(o.SmtpEncryption.Get()) {
		var ret string
		return ret
	}
	return *o.SmtpEncryption.Get()
}

// GetSmtpEncryptionOk returns a tuple with the SmtpEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Company) GetSmtpEncryptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmtpEncryption.Get(), o.SmtpEncryption.IsSet()
}

// HasSmtpEncryption returns a boolean if a field has been set.
func (o *Company) HasSmtpEncryption() bool {
	if o != nil && o.SmtpEncryption.IsSet() {
		return true
	}

	return false
}

// SetSmtpEncryption gets a reference to the given NullableString and assigns it to the SmtpEncryption field.
func (o *Company) SetSmtpEncryption(v string) {
	o.SmtpEncryption.Set(&v)
}

// SetSmtpEncryptionNil sets the value for SmtpEncryption to be an explicit nil
func (o *Company) SetSmtpEncryptionNil() {
	o.SmtpEncryption.Set(nil)
}

// UnsetSmtpEncryption ensures that no value is present for SmtpEncryption, not even an explicit nil
func (o *Company) UnsetSmtpEncryption() {
	o.SmtpEncryption.Unset()
}

// GetSmtpLocalDomain returns the SmtpLocalDomain field value if set, zero value otherwise.
func (o *Company) GetSmtpLocalDomain() string {
	if o == nil || IsNil(o.SmtpLocalDomain) {
		var ret string
		return ret
	}
	return *o.SmtpLocalDomain
}

// GetSmtpLocalDomainOk returns a tuple with the SmtpLocalDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSmtpLocalDomainOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpLocalDomain) {
		return nil, false
	}
	return o.SmtpLocalDomain, true
}

// HasSmtpLocalDomain returns a boolean if a field has been set.
func (o *Company) HasSmtpLocalDomain() bool {
	if o != nil && !IsNil(o.SmtpLocalDomain) {
		return true
	}

	return false
}

// SetSmtpLocalDomain gets a reference to the given string and assigns it to the SmtpLocalDomain field.
func (o *Company) SetSmtpLocalDomain(v string) {
	o.SmtpLocalDomain = &v
}

// GetSmtpVerifyPeer returns the SmtpVerifyPeer field value if set, zero value otherwise.
func (o *Company) GetSmtpVerifyPeer() bool {
	if o == nil || IsNil(o.SmtpVerifyPeer) {
		var ret bool
		return ret
	}
	return *o.SmtpVerifyPeer
}

// GetSmtpVerifyPeerOk returns a tuple with the SmtpVerifyPeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSmtpVerifyPeerOk() (*bool, bool) {
	if o == nil || IsNil(o.SmtpVerifyPeer) {
		return nil, false
	}
	return o.SmtpVerifyPeer, true
}

// HasSmtpVerifyPeer returns a boolean if a field has been set.
func (o *Company) HasSmtpVerifyPeer() bool {
	if o != nil && !IsNil(o.SmtpVerifyPeer) {
		return true
	}

	return false
}

// SetSmtpVerifyPeer gets a reference to the given bool and assigns it to the SmtpVerifyPeer field.
func (o *Company) SetSmtpVerifyPeer(v bool) {
	o.SmtpVerifyPeer = &v
}

// GetEInvoice returns the EInvoice field value if set, zero value otherwise.
func (o *Company) GetEInvoice() map[string]any {
	if o == nil || IsNil(o.EInvoice) {
		var ret map[string]any
		return ret
	}
	return o.EInvoice
}

// GetEInvoiceOk returns a tuple with the EInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetEInvoiceOk() (map[string]any, bool) {
	if o == nil || IsNil(o.EInvoice) {
		return map[string]any{}, false
	}
	return o.EInvoice, true
}

// HasEInvoice returns a boolean if a field has been set.
func (o *Company) HasEInvoice() bool {
	if o != nil && !IsNil(o.EInvoice) {
		return true
	}

	return false
}

// SetEInvoice gets a reference to the given map[string]interface{} and assigns it to the EInvoice field.
func (o *Company) SetEInvoice(v map[string]any) {
	o.EInvoice = v
}

// GetLegalEntityId returns the LegalEntityId field value if set, zero value otherwise.
func (o *Company) GetLegalEntityId() int32 {
	if o == nil || IsNil(o.LegalEntityId) {
		var ret int32
		return ret
	}
	return *o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetLegalEntityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LegalEntityId) {
		return nil, false
	}
	return o.LegalEntityId, true
}

// HasLegalEntityId returns a boolean if a field has been set.
func (o *Company) HasLegalEntityId() bool {
	if o != nil && !IsNil(o.LegalEntityId) {
		return true
	}

	return false
}

// SetLegalEntityId gets a reference to the given int32 and assigns it to the LegalEntityId field.
func (o *Company) SetLegalEntityId(v int32) {
	o.LegalEntityId = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Company) GetSettings() CompanySettings {
	if o == nil || IsNil(o.Settings) {
		var ret CompanySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Company) GetSettingsOk() (*CompanySettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Company) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given CompanySettings and assigns it to the Settings field.
func (o *Company) SetSettings(v CompanySettings) {
	o.Settings = &v
}

func (o Company) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Company) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SizeId) {
		toSerialize["size_id"] = o.SizeId
	}
	if !IsNil(o.IndustryId) {
		toSerialize["industry_id"] = o.IndustryId
	}
	if !IsNil(o.SlackWebhookUrl) {
		toSerialize["slack_webhook_url"] = o.SlackWebhookUrl
	}
	if !IsNil(o.GoogleAnalyticsKey) {
		toSerialize["google_analytics_key"] = o.GoogleAnalyticsKey
	}
	if !IsNil(o.PortalMode) {
		toSerialize["portal_mode"] = o.PortalMode
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.PortalDomain) {
		toSerialize["portal_domain"] = o.PortalDomain
	}
	if !IsNil(o.EnabledTaxRates) {
		toSerialize["enabled_tax_rates"] = o.EnabledTaxRates
	}
	if !IsNil(o.FillProducts) {
		toSerialize["fill_products"] = o.FillProducts
	}
	if !IsNil(o.ConvertProducts) {
		toSerialize["convert_products"] = o.ConvertProducts
	}
	if !IsNil(o.UpdateProducts) {
		toSerialize["update_products"] = o.UpdateProducts
	}
	if !IsNil(o.ShowProductDetails) {
		toSerialize["show_product_details"] = o.ShowProductDetails
	}
	if !IsNil(o.ShowProductCost) {
		toSerialize["show_product_cost"] = o.ShowProductCost
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.EnableProductCost) {
		toSerialize["enable_product_cost"] = o.EnableProductCost
	}
	if !IsNil(o.EnableProductQuantity) {
		toSerialize["enable_product_quantity"] = o.EnableProductQuantity
	}
	if !IsNil(o.DefaultQuantity) {
		toSerialize["default_quantity"] = o.DefaultQuantity
	}
	if !IsNil(o.CustomSurchargeTaxes1) {
		toSerialize["custom_surcharge_taxes1"] = o.CustomSurchargeTaxes1
	}
	if !IsNil(o.CustomSurchargeTaxes2) {
		toSerialize["custom_surcharge_taxes2"] = o.CustomSurchargeTaxes2
	}
	if !IsNil(o.CustomSurchargeTaxes3) {
		toSerialize["custom_surcharge_taxes3"] = o.CustomSurchargeTaxes3
	}
	if o.CustomSurchargeTaxes4 != nil {
		toSerialize["custom_surcharge_taxes4"] = o.CustomSurchargeTaxes4
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.CompanyKey) {
		toSerialize["company_key"] = o.CompanyKey
	}
	if !IsNil(o.ClientCanRegister) {
		toSerialize["client_can_register"] = o.ClientCanRegister
	}
	if !IsNil(o.EnabledModules) {
		toSerialize["enabled_modules"] = o.EnabledModules
	}
	if !IsNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	if !IsNil(o.FirstDayOfWeek) {
		toSerialize["first_day_of_week"] = o.FirstDayOfWeek
	}
	if !IsNil(o.FirstMonthOfYear) {
		toSerialize["first_month_of_year"] = o.FirstMonthOfYear
	}
	if !IsNil(o.EnabledItemTaxRates) {
		toSerialize["enabled_item_tax_rates"] = o.EnabledItemTaxRates
	}
	if !IsNil(o.IsLarge) {
		toSerialize["is_large"] = o.IsLarge
	}
	if !IsNil(o.DefaultAutoBill) {
		toSerialize["default_auto_bill"] = o.DefaultAutoBill
	}
	if !IsNil(o.MarkExpensesInvoiceable) {
		toSerialize["mark_expenses_invoiceable"] = o.MarkExpensesInvoiceable
	}
	if !IsNil(o.MarkExpensesPaid) {
		toSerialize["mark_expenses_paid"] = o.MarkExpensesPaid
	}
	if !IsNil(o.InvoiceExpenseDocuments) {
		toSerialize["invoice_expense_documents"] = o.InvoiceExpenseDocuments
	}
	if !IsNil(o.AutoStartTasks) {
		toSerialize["auto_start_tasks"] = o.AutoStartTasks
	}
	if !IsNil(o.InvoiceTaskTimelog) {
		toSerialize["invoice_task_timelog"] = o.InvoiceTaskTimelog
	}
	if !IsNil(o.InvoiceTaskDocuments) {
		toSerialize["invoice_task_documents"] = o.InvoiceTaskDocuments
	}
	if !IsNil(o.ShowTasksTable) {
		toSerialize["show_tasks_table"] = o.ShowTasksTable
	}
	if !IsNil(o.IsDisabled) {
		toSerialize["is_disabled"] = o.IsDisabled
	}
	if !IsNil(o.DefaultTaskIsDateBased) {
		toSerialize["default_task_is_date_based"] = o.DefaultTaskIsDateBased
	}
	if !IsNil(o.EnableProductDiscount) {
		toSerialize["enable_product_discount"] = o.EnableProductDiscount
	}
	if !IsNil(o.CalculateExpenseTaxByAmount) {
		toSerialize["calculate_expense_tax_by_amount"] = o.CalculateExpenseTaxByAmount
	}
	if !IsNil(o.ExpenseInclusiveTaxes) {
		toSerialize["expense_inclusive_taxes"] = o.ExpenseInclusiveTaxes
	}
	if !IsNil(o.SessionTimeout) {
		toSerialize["session_timeout"] = o.SessionTimeout
	}
	if !IsNil(o.OauthPasswordRequired) {
		toSerialize["oauth_password_required"] = o.OauthPasswordRequired
	}
	if !IsNil(o.InvoiceTaskDatelog) {
		toSerialize["invoice_task_datelog"] = o.InvoiceTaskDatelog
	}
	if !IsNil(o.DefaultPasswordTimeout) {
		toSerialize["default_password_timeout"] = o.DefaultPasswordTimeout
	}
	if !IsNil(o.ShowTaskEndDate) {
		toSerialize["show_task_end_date"] = o.ShowTaskEndDate
	}
	if !IsNil(o.MarkdownEnabled) {
		toSerialize["markdown_enabled"] = o.MarkdownEnabled
	}
	if !IsNil(o.ReportIncludeDrafts) {
		toSerialize["report_include_drafts"] = o.ReportIncludeDrafts
	}
	if !IsNil(o.ClientRegistrationFields) {
		toSerialize["client_registration_fields"] = o.ClientRegistrationFields
	}
	if !IsNil(o.StopOnUnpaidRecurring) {
		toSerialize["stop_on_unpaid_recurring"] = o.StopOnUnpaidRecurring
	}
	if !IsNil(o.UseQuoteTermsOnConversion) {
		toSerialize["use_quote_terms_on_conversion"] = o.UseQuoteTermsOnConversion
	}
	if !IsNil(o.EnableApplyingPayments) {
		toSerialize["enable_applying_payments"] = o.EnableApplyingPayments
	}
	if !IsNil(o.TrackInventory) {
		toSerialize["track_inventory"] = o.TrackInventory
	}
	if !IsNil(o.InventoryNotificationThreshold) {
		toSerialize["inventory_notification_threshold"] = o.InventoryNotificationThreshold
	}
	if !IsNil(o.StockNotification) {
		toSerialize["stock_notification"] = o.StockNotification
	}
	if !IsNil(o.MatomoUrl) {
		toSerialize["matomo_url"] = o.MatomoUrl
	}
	if !IsNil(o.MatomoId) {
		toSerialize["matomo_id"] = o.MatomoId
	}
	if !IsNil(o.EnabledExpenseTaxRates) {
		toSerialize["enabled_expense_tax_rates"] = o.EnabledExpenseTaxRates
	}
	if !IsNil(o.InvoiceTaskProject) {
		toSerialize["invoice_task_project"] = o.InvoiceTaskProject
	}
	if !IsNil(o.ReportIncludeDeleted) {
		toSerialize["report_include_deleted"] = o.ReportIncludeDeleted
	}
	if !IsNil(o.InvoiceTaskLock) {
		toSerialize["invoice_task_lock"] = o.InvoiceTaskLock
	}
	if !IsNil(o.ConvertPaymentCurrency) {
		toSerialize["convert_payment_currency"] = o.ConvertPaymentCurrency
	}
	if !IsNil(o.ConvertExpenseCurrency) {
		toSerialize["convert_expense_currency"] = o.ConvertExpenseCurrency
	}
	if !IsNil(o.NotifyVendorWhenPaid) {
		toSerialize["notify_vendor_when_paid"] = o.NotifyVendorWhenPaid
	}
	if !IsNil(o.InvoiceTaskHours) {
		toSerialize["invoice_task_hours"] = o.InvoiceTaskHours
	}
	if !IsNil(o.CalculateTaxes) {
		toSerialize["calculate_taxes"] = o.CalculateTaxes
	}
	if !IsNil(o.TaxData) {
		toSerialize["tax_data"] = o.TaxData
	}
	if !IsNil(o.EInvoiceCertificate) {
		toSerialize["e_invoice_certificate"] = o.EInvoiceCertificate
	}
	if !IsNil(o.EInvoiceCertificatePassphrase) {
		toSerialize["e_invoice_certificate_passphrase"] = o.EInvoiceCertificatePassphrase
	}
	if !IsNil(o.OriginTaxData) {
		toSerialize["origin_tax_data"] = o.OriginTaxData
	}
	if !IsNil(o.InvoiceTaskProjectHeader) {
		toSerialize["invoice_task_project_header"] = o.InvoiceTaskProjectHeader
	}
	if !IsNil(o.InvoiceTaskItemDescription) {
		toSerialize["invoice_task_item_description"] = o.InvoiceTaskItemDescription
	}
	if !IsNil(o.ExpenseMailbox) {
		toSerialize["expense_mailbox"] = o.ExpenseMailbox
	}
	if !IsNil(o.ExpenseMailboxActive) {
		toSerialize["expense_mailbox_active"] = o.ExpenseMailboxActive
	}
	if !IsNil(o.InboundMailboxAllowCompanyUsers) {
		toSerialize["inbound_mailbox_allow_company_users"] = o.InboundMailboxAllowCompanyUsers
	}
	if !IsNil(o.InboundMailboxAllowVendors) {
		toSerialize["inbound_mailbox_allow_vendors"] = o.InboundMailboxAllowVendors
	}
	if !IsNil(o.InboundMailboxAllowClients) {
		toSerialize["inbound_mailbox_allow_clients"] = o.InboundMailboxAllowClients
	}
	if !IsNil(o.InboundMailboxAllowUnknown) {
		toSerialize["inbound_mailbox_allow_unknown"] = o.InboundMailboxAllowUnknown
	}
	if !IsNil(o.InboundMailboxWhitelist) {
		toSerialize["inbound_mailbox_whitelist"] = o.InboundMailboxWhitelist
	}
	if !IsNil(o.InboundMailboxBlacklist) {
		toSerialize["inbound_mailbox_blacklist"] = o.InboundMailboxBlacklist
	}
	if !IsNil(o.SmtpHost) {
		toSerialize["smtp_host"] = o.SmtpHost
	}
	if !IsNil(o.SmtpPort) {
		toSerialize["smtp_port"] = o.SmtpPort
	}
	if o.SmtpEncryption.IsSet() {
		toSerialize["smtp_encryption"] = o.SmtpEncryption.Get()
	}
	if !IsNil(o.SmtpLocalDomain) {
		toSerialize["smtp_local_domain"] = o.SmtpLocalDomain
	}
	if !IsNil(o.SmtpVerifyPeer) {
		toSerialize["smtp_verify_peer"] = o.SmtpVerifyPeer
	}
	if !IsNil(o.EInvoice) {
		toSerialize["e_invoice"] = o.EInvoice
	}
	if !IsNil(o.LegalEntityId) {
		toSerialize["legal_entity_id"] = o.LegalEntityId
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableCompany struct {
	value *Company
	isSet bool
}

func (v NullableCompany) Get() *Company {
	return v.value
}

func (v *NullableCompany) Set(val *Company) {
	v.value = val
	v.isSet = true
}

func (v NullableCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompany(val *Company) *NullableCompany {
	return &NullableCompany{value: val, isSet: true}
}

func (v NullableCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
