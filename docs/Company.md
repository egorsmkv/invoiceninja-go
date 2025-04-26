# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique hashed identifier for the company | [optional] 
**SizeId** | Pointer to **string** | The unique identifier representing the company&#39;s size category | [optional] 
**IndustryId** | Pointer to **string** | The unique identifier representing the company&#39;s industry category | [optional] 
**SlackWebhookUrl** | Pointer to **string** | The URL for the company&#39;s Slack webhook notifications | [optional] 
**GoogleAnalyticsKey** | Pointer to **string** | The company&#39;s Google Analytics tracking ID | [optional] 
**PortalMode** | Pointer to **string** | The mode determining how client-facing URLs are structured (e.g., subdomain, domain, or iframe) | [optional] 
**Subdomain** | Pointer to **string** | The subdomain prefix for the company&#39;s domain (e.g., &#39;acme&#39; in acme.domain.com) | [optional] 
**PortalDomain** | Pointer to **string** | The fully qualified domain used for client-facing URLs | [optional] 
**EnabledTaxRates** | Pointer to **int32** | The number of tax rates used per entity | [optional] 
**FillProducts** | Pointer to **bool** | A flag determining whether to auto-fill product descriptions based on the product key | [optional] 
**ConvertProducts** | Pointer to **bool** | A flag determining whether to convert products between different types or units | [optional] 
**UpdateProducts** | Pointer to **bool** | A flag determining whether to update product descriptions when the description changes | [optional] 
**ShowProductDetails** | Pointer to **bool** | A flag determining whether to display product details in the user interface | [optional] 
**ShowProductCost** | Pointer to **bool** | A flag determining whether to display product cost is shown in the user interface | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | A mapping of custom fields for various objects within the company  **note**  The custom fields object is formatted as follows    key : value   entity(integer) : label|type    for example:   \&quot;company1\&quot;: \&quot;Number|single_line_text\&quot;,    This defines the first company custom field, with label Number which has a custom field type of a single text line    Supported entity types    invoice   product   client   contact   task   user   project   vendor    expense   payment         Supported input types values    single_line_text   date   switch    For text areas, you only need to supply the label ie \&quot;TextArea\&quot;, the | and type values are not required.   For the dropdown the data format is as follows:    label|your,drop,down,values,in,comma,separated,format  | [optional] 
**EnableProductCost** | Pointer to **bool** | A flag determining whether to show or hide the product cost field in the user interface | [optional] 
**EnableProductQuantity** | Pointer to **bool** | A flag determining whether to show or hide the product quantity field in the user interface | [optional] 
**DefaultQuantity** | Pointer to **bool** | A flag determining whether to use a default quantity for products | [optional] 
**CustomSurchargeTaxes1** | Pointer to **bool** | A flag determining whether to apply taxes on custom surcharge amounts for the first custom surcharge field | [optional] 
**CustomSurchargeTaxes2** | Pointer to **bool** | A flag determining whether to apply taxes on custom surcharge amounts for the second custom surcharge field | [optional] 
**CustomSurchargeTaxes3** | Pointer to **bool** | A flag determining whether to apply taxes on custom surcharge amounts for the third custom surcharge field | [optional] 
**CustomSurchargeTaxes4** | Pointer to **interface{}** | A flag determining whether to apply taxes on custom surcharge amounts for the fourth custom | [optional] 
**Logo** | Pointer to ***os.File** | The company logo file in binary format | [optional] 
**CompanyKey** | Pointer to **string** | The static company key hash used to identify the Company | [optional] [readonly] 
**ClientCanRegister** | Pointer to **bool** | A flag determining whether clients can register for the client portal | [optional] 
**EnabledModules** | Pointer to **int32** | Bitmask representation of the modules that are enabled in the application  &#x60;&#x60;&#x60; self::ENTITY_RECURRING_INVOICE &#x3D;&gt; 1, self::ENTITY_CREDIT &#x3D;&gt; 2, self::ENTITY_QUOTE &#x3D;&gt; 4, self::ENTITY_TASK &#x3D;&gt; 8, self::ENTITY_EXPENSE &#x3D;&gt; 16, self::ENTITY_PROJECT &#x3D;&gt; 32, self::ENTITY_VENDOR &#x3D;&gt; 64, self::ENTITY_TICKET &#x3D;&gt; 128, self::ENTITY_PROPOSAL &#x3D;&gt; 256, self::ENTITY_RECURRING_EXPENSE &#x3D;&gt; 512, self::ENTITY_RECURRING_TASK &#x3D;&gt; 1024, self::ENTITY_RECURRING_QUOTE &#x3D;&gt; 2048, &#x60;&#x60;&#x60;  The default per_page value is 20.  | [optional] 
**Db** | Pointer to **string** |  | [optional] [readonly] 
**FirstDayOfWeek** | Pointer to **string** | The first day of the week for the company | [optional] 
**FirstMonthOfYear** | Pointer to **string** | The first month for the company financial year | [optional] 
**EnabledItemTaxRates** | Pointer to **int32** | The number of tax rates used per item | [optional] 
**IsLarge** | Pointer to **bool** | A flag determining whether the company is considered large | [optional] 
**DefaultAutoBill** | Pointer to **string** | A flag determining whether to auto-bill clients by default  values:  - always - Always auto bill - disabled - Never auto bill - optin - Allow the client to select their auto bill status with the default being disabled - optout -Allow the client to select their auto bill status with the default being enabled  | [optional] 
**MarkExpensesInvoiceable** | Pointer to **bool** | A flag determining whether to mark expenses as invoiceable by default | [optional] 
**MarkExpensesPaid** | Pointer to **bool** | A flag determining whether to mark expenses as paid by default | [optional] 
**InvoiceExpenseDocuments** | Pointer to **bool** | A flag determining whether to include expense documents on invoices by default | [optional] 
**AutoStartTasks** | Pointer to **bool** | A flag determining whether to auto-start tasks by default | [optional] 
**InvoiceTaskTimelog** | Pointer to **bool** | A flag determining whether to include task time logs on invoices by default | [optional] 
**InvoiceTaskDocuments** | Pointer to **bool** | A flag determining whether to include task documents on invoices by default | [optional] 
**ShowTasksTable** | Pointer to **bool** | A flag determining whether to show the tasks table on invoices by default | [optional] 
**IsDisabled** | Pointer to **bool** | A flag determining whether the company is disabled | [optional] 
**DefaultTaskIsDateBased** | Pointer to **bool** | A flag determining whether to default tasks to be date-based | [optional] 
**EnableProductDiscount** | Pointer to **bool** | A flag determining whether to show or hide the product discount field in the user interface | [optional] 
**CalculateExpenseTaxByAmount** | Pointer to **bool** | A flag determining whether to calculate expense taxes by amount | [optional] 
**ExpenseInclusiveTaxes** | Pointer to **bool** | A flag determining whether to include taxes in the expense amount | [optional] 
**SessionTimeout** | Pointer to **int32** | The session timeout for the company | [optional] 
**OauthPasswordRequired** | Pointer to **bool** | A flag determining whether to require a password for &#x60;dangerous&#x60; actions when using OAuth | [optional] 
**InvoiceTaskDatelog** | Pointer to **bool** | A flag determining whether to include task date logs on invoices by default | [optional] 
**DefaultPasswordTimeout** | Pointer to **int32** | The default password timeout for the company | [optional] 
**ShowTaskEndDate** | Pointer to **bool** | A flag determining whether to show the task end date on invoices by default | [optional] 
**MarkdownEnabled** | Pointer to **bool** | A flag determining whether markdown is enabled for the company | [optional] 
**ReportIncludeDrafts** | Pointer to **bool** | A flag determining whether to include draft invoices in reports | [optional] 
**ClientRegistrationFields** | Pointer to **map[string]interface{}** | The client registration fields for the company | [optional] 
**StopOnUnpaidRecurring** | Pointer to **bool** | A flag determining whether to stop recurring invoices when they are unpaid | [optional] 
**UseQuoteTermsOnConversion** | Pointer to **bool** | A flag determining whether to use quote terms on conversion to an invoice | [optional] 
**EnableApplyingPayments** | Pointer to **bool** | A flag determining whether to enable applying payments to invoices | [optional] 
**TrackInventory** | Pointer to **bool** | A flag determining whether to track inventory for the company | [optional] 
**InventoryNotificationThreshold** | Pointer to **int32** | The inventory notification threshold for the company | [optional] 
**StockNotification** | Pointer to **bool** | A flag determining whether to send stock notifications for the company | [optional] 
**MatomoUrl** | Pointer to **string** | The Matomo URL for the company | [optional] 
**MatomoId** | Pointer to **string** | The Matomo ID for the company | [optional] 
**EnabledExpenseTaxRates** | Pointer to **int32** | The number of tax rates used per expense | [optional] 
**InvoiceTaskProject** | Pointer to **bool** | A flag determining whether to include the project on invoices by default | [optional] 
**ReportIncludeDeleted** | Pointer to **bool** | A flag determining whether to include deleted invoices in reports | [optional] 
**InvoiceTaskLock** | Pointer to **bool** | A flag determining whether to lock tasks when invoiced | [optional] 
**ConvertPaymentCurrency** | Pointer to **bool** | A flag determining whether to convert the payment currency | [optional] 
**ConvertExpenseCurrency** | Pointer to **bool** | A flag determining whether to convert the expense currency | [optional] 
**NotifyVendorWhenPaid** | Pointer to **bool** | A flag determining whether to notify the vendor when an expense is paid | [optional] 
**InvoiceTaskHours** | Pointer to **bool** | A flag determining whether to include the task hours on invoices by default | [optional] 
**CalculateTaxes** | Pointer to **bool** | A flag determining whether to calculate taxes for the company | [optional] 
**TaxData** | Pointer to **map[string]interface{}** | The tax data for the company | [optional] 
**EInvoiceCertificate** | Pointer to **string** | The e-invoice certificate for the company | [optional] 
**EInvoiceCertificatePassphrase** | Pointer to **string** | The e-invoice certificate passphrase for the company | [optional] 
**OriginTaxData** | Pointer to **map[string]interface{}** | The origin tax data for the company | [optional] 
**InvoiceTaskProjectHeader** | Pointer to **bool** | A flag determining whether to include the project header on invoices by default | [optional] 
**InvoiceTaskItemDescription** | Pointer to **bool** | A flag determining whether to include the item description on invoices by default | [optional] 
**ExpenseMailbox** | Pointer to **string** | The email address for the expense mailbox | [optional] 
**ExpenseMailboxActive** | Pointer to **bool** | Whether the expense mailbox is active | [optional] 
**InboundMailboxAllowCompanyUsers** | Pointer to **bool** | Whether company users are allowed to use the inbound mailbox | [optional] 
**InboundMailboxAllowVendors** | Pointer to **bool** | Whether vendors are allowed to use the inbound mailbox | [optional] 
**InboundMailboxAllowClients** | Pointer to **bool** | Whether clients are allowed to use the inbound mailbox | [optional] 
**InboundMailboxAllowUnknown** | Pointer to **bool** | Whether unknown senders are allowed to use the inbound mailbox | [optional] 
**InboundMailboxWhitelist** | Pointer to **string** | Comma-separated list of whitelisted email addresses for the inbound mailbox | [optional] 
**InboundMailboxBlacklist** | Pointer to **string** | Comma-separated list of blacklisted email addresses for the inbound mailbox | [optional] 
**SmtpHost** | Pointer to **string** | The SMTP host for sending emails | [optional] 
**SmtpPort** | Pointer to **int32** | The SMTP port for sending emails | [optional] 
**SmtpEncryption** | Pointer to **NullableString** | The encryption method for SMTP | [optional] 
**SmtpLocalDomain** | Pointer to **string** | The local domain for SMTP | [optional] 
**SmtpVerifyPeer** | Pointer to **bool** | Whether to verify the SMTP peer | [optional] 
**EInvoice** | Pointer to **map[string]interface{}** | E-invoice settings for the company | [optional] 
**LegalEntityId** | Pointer to **int32** | The ID of the legal entity associated with the company | [optional] 
**Settings** | Pointer to [**CompanySettings**](CompanySettings.md) |  | [optional] 

## Methods

### NewCompany

`func NewCompany() *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Company) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSizeId

`func (o *Company) GetSizeId() string`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *Company) GetSizeIdOk() (*string, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *Company) SetSizeId(v string)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *Company) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### GetIndustryId

`func (o *Company) GetIndustryId() string`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *Company) GetIndustryIdOk() (*string, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *Company) SetIndustryId(v string)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *Company) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetSlackWebhookUrl

`func (o *Company) GetSlackWebhookUrl() string`

GetSlackWebhookUrl returns the SlackWebhookUrl field if non-nil, zero value otherwise.

### GetSlackWebhookUrlOk

`func (o *Company) GetSlackWebhookUrlOk() (*string, bool)`

GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackWebhookUrl

`func (o *Company) SetSlackWebhookUrl(v string)`

SetSlackWebhookUrl sets SlackWebhookUrl field to given value.

### HasSlackWebhookUrl

`func (o *Company) HasSlackWebhookUrl() bool`

HasSlackWebhookUrl returns a boolean if a field has been set.

### GetGoogleAnalyticsKey

`func (o *Company) GetGoogleAnalyticsKey() string`

GetGoogleAnalyticsKey returns the GoogleAnalyticsKey field if non-nil, zero value otherwise.

### GetGoogleAnalyticsKeyOk

`func (o *Company) GetGoogleAnalyticsKeyOk() (*string, bool)`

GetGoogleAnalyticsKeyOk returns a tuple with the GoogleAnalyticsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalyticsKey

`func (o *Company) SetGoogleAnalyticsKey(v string)`

SetGoogleAnalyticsKey sets GoogleAnalyticsKey field to given value.

### HasGoogleAnalyticsKey

`func (o *Company) HasGoogleAnalyticsKey() bool`

HasGoogleAnalyticsKey returns a boolean if a field has been set.

### GetPortalMode

`func (o *Company) GetPortalMode() string`

GetPortalMode returns the PortalMode field if non-nil, zero value otherwise.

### GetPortalModeOk

`func (o *Company) GetPortalModeOk() (*string, bool)`

GetPortalModeOk returns a tuple with the PortalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalMode

`func (o *Company) SetPortalMode(v string)`

SetPortalMode sets PortalMode field to given value.

### HasPortalMode

`func (o *Company) HasPortalMode() bool`

HasPortalMode returns a boolean if a field has been set.

### GetSubdomain

`func (o *Company) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Company) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Company) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Company) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetPortalDomain

`func (o *Company) GetPortalDomain() string`

GetPortalDomain returns the PortalDomain field if non-nil, zero value otherwise.

### GetPortalDomainOk

`func (o *Company) GetPortalDomainOk() (*string, bool)`

GetPortalDomainOk returns a tuple with the PortalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalDomain

`func (o *Company) SetPortalDomain(v string)`

SetPortalDomain sets PortalDomain field to given value.

### HasPortalDomain

`func (o *Company) HasPortalDomain() bool`

HasPortalDomain returns a boolean if a field has been set.

### GetEnabledTaxRates

`func (o *Company) GetEnabledTaxRates() int32`

GetEnabledTaxRates returns the EnabledTaxRates field if non-nil, zero value otherwise.

### GetEnabledTaxRatesOk

`func (o *Company) GetEnabledTaxRatesOk() (*int32, bool)`

GetEnabledTaxRatesOk returns a tuple with the EnabledTaxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledTaxRates

`func (o *Company) SetEnabledTaxRates(v int32)`

SetEnabledTaxRates sets EnabledTaxRates field to given value.

### HasEnabledTaxRates

`func (o *Company) HasEnabledTaxRates() bool`

HasEnabledTaxRates returns a boolean if a field has been set.

### GetFillProducts

`func (o *Company) GetFillProducts() bool`

GetFillProducts returns the FillProducts field if non-nil, zero value otherwise.

### GetFillProductsOk

`func (o *Company) GetFillProductsOk() (*bool, bool)`

GetFillProductsOk returns a tuple with the FillProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillProducts

`func (o *Company) SetFillProducts(v bool)`

SetFillProducts sets FillProducts field to given value.

### HasFillProducts

`func (o *Company) HasFillProducts() bool`

HasFillProducts returns a boolean if a field has been set.

### GetConvertProducts

`func (o *Company) GetConvertProducts() bool`

GetConvertProducts returns the ConvertProducts field if non-nil, zero value otherwise.

### GetConvertProductsOk

`func (o *Company) GetConvertProductsOk() (*bool, bool)`

GetConvertProductsOk returns a tuple with the ConvertProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertProducts

`func (o *Company) SetConvertProducts(v bool)`

SetConvertProducts sets ConvertProducts field to given value.

### HasConvertProducts

`func (o *Company) HasConvertProducts() bool`

HasConvertProducts returns a boolean if a field has been set.

### GetUpdateProducts

`func (o *Company) GetUpdateProducts() bool`

GetUpdateProducts returns the UpdateProducts field if non-nil, zero value otherwise.

### GetUpdateProductsOk

`func (o *Company) GetUpdateProductsOk() (*bool, bool)`

GetUpdateProductsOk returns a tuple with the UpdateProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateProducts

`func (o *Company) SetUpdateProducts(v bool)`

SetUpdateProducts sets UpdateProducts field to given value.

### HasUpdateProducts

`func (o *Company) HasUpdateProducts() bool`

HasUpdateProducts returns a boolean if a field has been set.

### GetShowProductDetails

`func (o *Company) GetShowProductDetails() bool`

GetShowProductDetails returns the ShowProductDetails field if non-nil, zero value otherwise.

### GetShowProductDetailsOk

`func (o *Company) GetShowProductDetailsOk() (*bool, bool)`

GetShowProductDetailsOk returns a tuple with the ShowProductDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowProductDetails

`func (o *Company) SetShowProductDetails(v bool)`

SetShowProductDetails sets ShowProductDetails field to given value.

### HasShowProductDetails

`func (o *Company) HasShowProductDetails() bool`

HasShowProductDetails returns a boolean if a field has been set.

### GetShowProductCost

`func (o *Company) GetShowProductCost() bool`

GetShowProductCost returns the ShowProductCost field if non-nil, zero value otherwise.

### GetShowProductCostOk

`func (o *Company) GetShowProductCostOk() (*bool, bool)`

GetShowProductCostOk returns a tuple with the ShowProductCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowProductCost

`func (o *Company) SetShowProductCost(v bool)`

SetShowProductCost sets ShowProductCost field to given value.

### HasShowProductCost

`func (o *Company) HasShowProductCost() bool`

HasShowProductCost returns a boolean if a field has been set.

### GetCustomFields

`func (o *Company) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Company) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Company) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Company) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEnableProductCost

`func (o *Company) GetEnableProductCost() bool`

GetEnableProductCost returns the EnableProductCost field if non-nil, zero value otherwise.

### GetEnableProductCostOk

`func (o *Company) GetEnableProductCostOk() (*bool, bool)`

GetEnableProductCostOk returns a tuple with the EnableProductCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductCost

`func (o *Company) SetEnableProductCost(v bool)`

SetEnableProductCost sets EnableProductCost field to given value.

### HasEnableProductCost

`func (o *Company) HasEnableProductCost() bool`

HasEnableProductCost returns a boolean if a field has been set.

### GetEnableProductQuantity

`func (o *Company) GetEnableProductQuantity() bool`

GetEnableProductQuantity returns the EnableProductQuantity field if non-nil, zero value otherwise.

### GetEnableProductQuantityOk

`func (o *Company) GetEnableProductQuantityOk() (*bool, bool)`

GetEnableProductQuantityOk returns a tuple with the EnableProductQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductQuantity

`func (o *Company) SetEnableProductQuantity(v bool)`

SetEnableProductQuantity sets EnableProductQuantity field to given value.

### HasEnableProductQuantity

`func (o *Company) HasEnableProductQuantity() bool`

HasEnableProductQuantity returns a boolean if a field has been set.

### GetDefaultQuantity

`func (o *Company) GetDefaultQuantity() bool`

GetDefaultQuantity returns the DefaultQuantity field if non-nil, zero value otherwise.

### GetDefaultQuantityOk

`func (o *Company) GetDefaultQuantityOk() (*bool, bool)`

GetDefaultQuantityOk returns a tuple with the DefaultQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuantity

`func (o *Company) SetDefaultQuantity(v bool)`

SetDefaultQuantity sets DefaultQuantity field to given value.

### HasDefaultQuantity

`func (o *Company) HasDefaultQuantity() bool`

HasDefaultQuantity returns a boolean if a field has been set.

### GetCustomSurchargeTaxes1

`func (o *Company) GetCustomSurchargeTaxes1() bool`

GetCustomSurchargeTaxes1 returns the CustomSurchargeTaxes1 field if non-nil, zero value otherwise.

### GetCustomSurchargeTaxes1Ok

`func (o *Company) GetCustomSurchargeTaxes1Ok() (*bool, bool)`

GetCustomSurchargeTaxes1Ok returns a tuple with the CustomSurchargeTaxes1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTaxes1

`func (o *Company) SetCustomSurchargeTaxes1(v bool)`

SetCustomSurchargeTaxes1 sets CustomSurchargeTaxes1 field to given value.

### HasCustomSurchargeTaxes1

`func (o *Company) HasCustomSurchargeTaxes1() bool`

HasCustomSurchargeTaxes1 returns a boolean if a field has been set.

### GetCustomSurchargeTaxes2

`func (o *Company) GetCustomSurchargeTaxes2() bool`

GetCustomSurchargeTaxes2 returns the CustomSurchargeTaxes2 field if non-nil, zero value otherwise.

### GetCustomSurchargeTaxes2Ok

`func (o *Company) GetCustomSurchargeTaxes2Ok() (*bool, bool)`

GetCustomSurchargeTaxes2Ok returns a tuple with the CustomSurchargeTaxes2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTaxes2

`func (o *Company) SetCustomSurchargeTaxes2(v bool)`

SetCustomSurchargeTaxes2 sets CustomSurchargeTaxes2 field to given value.

### HasCustomSurchargeTaxes2

`func (o *Company) HasCustomSurchargeTaxes2() bool`

HasCustomSurchargeTaxes2 returns a boolean if a field has been set.

### GetCustomSurchargeTaxes3

`func (o *Company) GetCustomSurchargeTaxes3() bool`

GetCustomSurchargeTaxes3 returns the CustomSurchargeTaxes3 field if non-nil, zero value otherwise.

### GetCustomSurchargeTaxes3Ok

`func (o *Company) GetCustomSurchargeTaxes3Ok() (*bool, bool)`

GetCustomSurchargeTaxes3Ok returns a tuple with the CustomSurchargeTaxes3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTaxes3

`func (o *Company) SetCustomSurchargeTaxes3(v bool)`

SetCustomSurchargeTaxes3 sets CustomSurchargeTaxes3 field to given value.

### HasCustomSurchargeTaxes3

`func (o *Company) HasCustomSurchargeTaxes3() bool`

HasCustomSurchargeTaxes3 returns a boolean if a field has been set.

### GetCustomSurchargeTaxes4

`func (o *Company) GetCustomSurchargeTaxes4() interface{}`

GetCustomSurchargeTaxes4 returns the CustomSurchargeTaxes4 field if non-nil, zero value otherwise.

### GetCustomSurchargeTaxes4Ok

`func (o *Company) GetCustomSurchargeTaxes4Ok() (*interface{}, bool)`

GetCustomSurchargeTaxes4Ok returns a tuple with the CustomSurchargeTaxes4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTaxes4

`func (o *Company) SetCustomSurchargeTaxes4(v interface{})`

SetCustomSurchargeTaxes4 sets CustomSurchargeTaxes4 field to given value.

### HasCustomSurchargeTaxes4

`func (o *Company) HasCustomSurchargeTaxes4() bool`

HasCustomSurchargeTaxes4 returns a boolean if a field has been set.

### SetCustomSurchargeTaxes4Nil

`func (o *Company) SetCustomSurchargeTaxes4Nil(b bool)`

 SetCustomSurchargeTaxes4Nil sets the value for CustomSurchargeTaxes4 to be an explicit nil

### UnsetCustomSurchargeTaxes4
`func (o *Company) UnsetCustomSurchargeTaxes4()`

UnsetCustomSurchargeTaxes4 ensures that no value is present for CustomSurchargeTaxes4, not even an explicit nil
### GetLogo

`func (o *Company) GetLogo() *os.File`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Company) GetLogoOk() (**os.File, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Company) SetLogo(v *os.File)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Company) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetCompanyKey

`func (o *Company) GetCompanyKey() string`

GetCompanyKey returns the CompanyKey field if non-nil, zero value otherwise.

### GetCompanyKeyOk

`func (o *Company) GetCompanyKeyOk() (*string, bool)`

GetCompanyKeyOk returns a tuple with the CompanyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyKey

`func (o *Company) SetCompanyKey(v string)`

SetCompanyKey sets CompanyKey field to given value.

### HasCompanyKey

`func (o *Company) HasCompanyKey() bool`

HasCompanyKey returns a boolean if a field has been set.

### GetClientCanRegister

`func (o *Company) GetClientCanRegister() bool`

GetClientCanRegister returns the ClientCanRegister field if non-nil, zero value otherwise.

### GetClientCanRegisterOk

`func (o *Company) GetClientCanRegisterOk() (*bool, bool)`

GetClientCanRegisterOk returns a tuple with the ClientCanRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCanRegister

`func (o *Company) SetClientCanRegister(v bool)`

SetClientCanRegister sets ClientCanRegister field to given value.

### HasClientCanRegister

`func (o *Company) HasClientCanRegister() bool`

HasClientCanRegister returns a boolean if a field has been set.

### GetEnabledModules

`func (o *Company) GetEnabledModules() int32`

GetEnabledModules returns the EnabledModules field if non-nil, zero value otherwise.

### GetEnabledModulesOk

`func (o *Company) GetEnabledModulesOk() (*int32, bool)`

GetEnabledModulesOk returns a tuple with the EnabledModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledModules

`func (o *Company) SetEnabledModules(v int32)`

SetEnabledModules sets EnabledModules field to given value.

### HasEnabledModules

`func (o *Company) HasEnabledModules() bool`

HasEnabledModules returns a boolean if a field has been set.

### GetDb

`func (o *Company) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *Company) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *Company) SetDb(v string)`

SetDb sets Db field to given value.

### HasDb

`func (o *Company) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetFirstDayOfWeek

`func (o *Company) GetFirstDayOfWeek() string`

GetFirstDayOfWeek returns the FirstDayOfWeek field if non-nil, zero value otherwise.

### GetFirstDayOfWeekOk

`func (o *Company) GetFirstDayOfWeekOk() (*string, bool)`

GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDayOfWeek

`func (o *Company) SetFirstDayOfWeek(v string)`

SetFirstDayOfWeek sets FirstDayOfWeek field to given value.

### HasFirstDayOfWeek

`func (o *Company) HasFirstDayOfWeek() bool`

HasFirstDayOfWeek returns a boolean if a field has been set.

### GetFirstMonthOfYear

`func (o *Company) GetFirstMonthOfYear() string`

GetFirstMonthOfYear returns the FirstMonthOfYear field if non-nil, zero value otherwise.

### GetFirstMonthOfYearOk

`func (o *Company) GetFirstMonthOfYearOk() (*string, bool)`

GetFirstMonthOfYearOk returns a tuple with the FirstMonthOfYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMonthOfYear

`func (o *Company) SetFirstMonthOfYear(v string)`

SetFirstMonthOfYear sets FirstMonthOfYear field to given value.

### HasFirstMonthOfYear

`func (o *Company) HasFirstMonthOfYear() bool`

HasFirstMonthOfYear returns a boolean if a field has been set.

### GetEnabledItemTaxRates

`func (o *Company) GetEnabledItemTaxRates() int32`

GetEnabledItemTaxRates returns the EnabledItemTaxRates field if non-nil, zero value otherwise.

### GetEnabledItemTaxRatesOk

`func (o *Company) GetEnabledItemTaxRatesOk() (*int32, bool)`

GetEnabledItemTaxRatesOk returns a tuple with the EnabledItemTaxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledItemTaxRates

`func (o *Company) SetEnabledItemTaxRates(v int32)`

SetEnabledItemTaxRates sets EnabledItemTaxRates field to given value.

### HasEnabledItemTaxRates

`func (o *Company) HasEnabledItemTaxRates() bool`

HasEnabledItemTaxRates returns a boolean if a field has been set.

### GetIsLarge

`func (o *Company) GetIsLarge() bool`

GetIsLarge returns the IsLarge field if non-nil, zero value otherwise.

### GetIsLargeOk

`func (o *Company) GetIsLargeOk() (*bool, bool)`

GetIsLargeOk returns a tuple with the IsLarge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLarge

`func (o *Company) SetIsLarge(v bool)`

SetIsLarge sets IsLarge field to given value.

### HasIsLarge

`func (o *Company) HasIsLarge() bool`

HasIsLarge returns a boolean if a field has been set.

### GetDefaultAutoBill

`func (o *Company) GetDefaultAutoBill() string`

GetDefaultAutoBill returns the DefaultAutoBill field if non-nil, zero value otherwise.

### GetDefaultAutoBillOk

`func (o *Company) GetDefaultAutoBillOk() (*string, bool)`

GetDefaultAutoBillOk returns a tuple with the DefaultAutoBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAutoBill

`func (o *Company) SetDefaultAutoBill(v string)`

SetDefaultAutoBill sets DefaultAutoBill field to given value.

### HasDefaultAutoBill

`func (o *Company) HasDefaultAutoBill() bool`

HasDefaultAutoBill returns a boolean if a field has been set.

### GetMarkExpensesInvoiceable

`func (o *Company) GetMarkExpensesInvoiceable() bool`

GetMarkExpensesInvoiceable returns the MarkExpensesInvoiceable field if non-nil, zero value otherwise.

### GetMarkExpensesInvoiceableOk

`func (o *Company) GetMarkExpensesInvoiceableOk() (*bool, bool)`

GetMarkExpensesInvoiceableOk returns a tuple with the MarkExpensesInvoiceable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkExpensesInvoiceable

`func (o *Company) SetMarkExpensesInvoiceable(v bool)`

SetMarkExpensesInvoiceable sets MarkExpensesInvoiceable field to given value.

### HasMarkExpensesInvoiceable

`func (o *Company) HasMarkExpensesInvoiceable() bool`

HasMarkExpensesInvoiceable returns a boolean if a field has been set.

### GetMarkExpensesPaid

`func (o *Company) GetMarkExpensesPaid() bool`

GetMarkExpensesPaid returns the MarkExpensesPaid field if non-nil, zero value otherwise.

### GetMarkExpensesPaidOk

`func (o *Company) GetMarkExpensesPaidOk() (*bool, bool)`

GetMarkExpensesPaidOk returns a tuple with the MarkExpensesPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkExpensesPaid

`func (o *Company) SetMarkExpensesPaid(v bool)`

SetMarkExpensesPaid sets MarkExpensesPaid field to given value.

### HasMarkExpensesPaid

`func (o *Company) HasMarkExpensesPaid() bool`

HasMarkExpensesPaid returns a boolean if a field has been set.

### GetInvoiceExpenseDocuments

`func (o *Company) GetInvoiceExpenseDocuments() bool`

GetInvoiceExpenseDocuments returns the InvoiceExpenseDocuments field if non-nil, zero value otherwise.

### GetInvoiceExpenseDocumentsOk

`func (o *Company) GetInvoiceExpenseDocumentsOk() (*bool, bool)`

GetInvoiceExpenseDocumentsOk returns a tuple with the InvoiceExpenseDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceExpenseDocuments

`func (o *Company) SetInvoiceExpenseDocuments(v bool)`

SetInvoiceExpenseDocuments sets InvoiceExpenseDocuments field to given value.

### HasInvoiceExpenseDocuments

`func (o *Company) HasInvoiceExpenseDocuments() bool`

HasInvoiceExpenseDocuments returns a boolean if a field has been set.

### GetAutoStartTasks

`func (o *Company) GetAutoStartTasks() bool`

GetAutoStartTasks returns the AutoStartTasks field if non-nil, zero value otherwise.

### GetAutoStartTasksOk

`func (o *Company) GetAutoStartTasksOk() (*bool, bool)`

GetAutoStartTasksOk returns a tuple with the AutoStartTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStartTasks

`func (o *Company) SetAutoStartTasks(v bool)`

SetAutoStartTasks sets AutoStartTasks field to given value.

### HasAutoStartTasks

`func (o *Company) HasAutoStartTasks() bool`

HasAutoStartTasks returns a boolean if a field has been set.

### GetInvoiceTaskTimelog

`func (o *Company) GetInvoiceTaskTimelog() bool`

GetInvoiceTaskTimelog returns the InvoiceTaskTimelog field if non-nil, zero value otherwise.

### GetInvoiceTaskTimelogOk

`func (o *Company) GetInvoiceTaskTimelogOk() (*bool, bool)`

GetInvoiceTaskTimelogOk returns a tuple with the InvoiceTaskTimelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskTimelog

`func (o *Company) SetInvoiceTaskTimelog(v bool)`

SetInvoiceTaskTimelog sets InvoiceTaskTimelog field to given value.

### HasInvoiceTaskTimelog

`func (o *Company) HasInvoiceTaskTimelog() bool`

HasInvoiceTaskTimelog returns a boolean if a field has been set.

### GetInvoiceTaskDocuments

`func (o *Company) GetInvoiceTaskDocuments() bool`

GetInvoiceTaskDocuments returns the InvoiceTaskDocuments field if non-nil, zero value otherwise.

### GetInvoiceTaskDocumentsOk

`func (o *Company) GetInvoiceTaskDocumentsOk() (*bool, bool)`

GetInvoiceTaskDocumentsOk returns a tuple with the InvoiceTaskDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskDocuments

`func (o *Company) SetInvoiceTaskDocuments(v bool)`

SetInvoiceTaskDocuments sets InvoiceTaskDocuments field to given value.

### HasInvoiceTaskDocuments

`func (o *Company) HasInvoiceTaskDocuments() bool`

HasInvoiceTaskDocuments returns a boolean if a field has been set.

### GetShowTasksTable

`func (o *Company) GetShowTasksTable() bool`

GetShowTasksTable returns the ShowTasksTable field if non-nil, zero value otherwise.

### GetShowTasksTableOk

`func (o *Company) GetShowTasksTableOk() (*bool, bool)`

GetShowTasksTableOk returns a tuple with the ShowTasksTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTasksTable

`func (o *Company) SetShowTasksTable(v bool)`

SetShowTasksTable sets ShowTasksTable field to given value.

### HasShowTasksTable

`func (o *Company) HasShowTasksTable() bool`

HasShowTasksTable returns a boolean if a field has been set.

### GetIsDisabled

`func (o *Company) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *Company) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *Company) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *Company) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetDefaultTaskIsDateBased

`func (o *Company) GetDefaultTaskIsDateBased() bool`

GetDefaultTaskIsDateBased returns the DefaultTaskIsDateBased field if non-nil, zero value otherwise.

### GetDefaultTaskIsDateBasedOk

`func (o *Company) GetDefaultTaskIsDateBasedOk() (*bool, bool)`

GetDefaultTaskIsDateBasedOk returns a tuple with the DefaultTaskIsDateBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaskIsDateBased

`func (o *Company) SetDefaultTaskIsDateBased(v bool)`

SetDefaultTaskIsDateBased sets DefaultTaskIsDateBased field to given value.

### HasDefaultTaskIsDateBased

`func (o *Company) HasDefaultTaskIsDateBased() bool`

HasDefaultTaskIsDateBased returns a boolean if a field has been set.

### GetEnableProductDiscount

`func (o *Company) GetEnableProductDiscount() bool`

GetEnableProductDiscount returns the EnableProductDiscount field if non-nil, zero value otherwise.

### GetEnableProductDiscountOk

`func (o *Company) GetEnableProductDiscountOk() (*bool, bool)`

GetEnableProductDiscountOk returns a tuple with the EnableProductDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductDiscount

`func (o *Company) SetEnableProductDiscount(v bool)`

SetEnableProductDiscount sets EnableProductDiscount field to given value.

### HasEnableProductDiscount

`func (o *Company) HasEnableProductDiscount() bool`

HasEnableProductDiscount returns a boolean if a field has been set.

### GetCalculateExpenseTaxByAmount

`func (o *Company) GetCalculateExpenseTaxByAmount() bool`

GetCalculateExpenseTaxByAmount returns the CalculateExpenseTaxByAmount field if non-nil, zero value otherwise.

### GetCalculateExpenseTaxByAmountOk

`func (o *Company) GetCalculateExpenseTaxByAmountOk() (*bool, bool)`

GetCalculateExpenseTaxByAmountOk returns a tuple with the CalculateExpenseTaxByAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateExpenseTaxByAmount

`func (o *Company) SetCalculateExpenseTaxByAmount(v bool)`

SetCalculateExpenseTaxByAmount sets CalculateExpenseTaxByAmount field to given value.

### HasCalculateExpenseTaxByAmount

`func (o *Company) HasCalculateExpenseTaxByAmount() bool`

HasCalculateExpenseTaxByAmount returns a boolean if a field has been set.

### GetExpenseInclusiveTaxes

`func (o *Company) GetExpenseInclusiveTaxes() bool`

GetExpenseInclusiveTaxes returns the ExpenseInclusiveTaxes field if non-nil, zero value otherwise.

### GetExpenseInclusiveTaxesOk

`func (o *Company) GetExpenseInclusiveTaxesOk() (*bool, bool)`

GetExpenseInclusiveTaxesOk returns a tuple with the ExpenseInclusiveTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseInclusiveTaxes

`func (o *Company) SetExpenseInclusiveTaxes(v bool)`

SetExpenseInclusiveTaxes sets ExpenseInclusiveTaxes field to given value.

### HasExpenseInclusiveTaxes

`func (o *Company) HasExpenseInclusiveTaxes() bool`

HasExpenseInclusiveTaxes returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *Company) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *Company) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *Company) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *Company) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetOauthPasswordRequired

`func (o *Company) GetOauthPasswordRequired() bool`

GetOauthPasswordRequired returns the OauthPasswordRequired field if non-nil, zero value otherwise.

### GetOauthPasswordRequiredOk

`func (o *Company) GetOauthPasswordRequiredOk() (*bool, bool)`

GetOauthPasswordRequiredOk returns a tuple with the OauthPasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthPasswordRequired

`func (o *Company) SetOauthPasswordRequired(v bool)`

SetOauthPasswordRequired sets OauthPasswordRequired field to given value.

### HasOauthPasswordRequired

`func (o *Company) HasOauthPasswordRequired() bool`

HasOauthPasswordRequired returns a boolean if a field has been set.

### GetInvoiceTaskDatelog

`func (o *Company) GetInvoiceTaskDatelog() bool`

GetInvoiceTaskDatelog returns the InvoiceTaskDatelog field if non-nil, zero value otherwise.

### GetInvoiceTaskDatelogOk

`func (o *Company) GetInvoiceTaskDatelogOk() (*bool, bool)`

GetInvoiceTaskDatelogOk returns a tuple with the InvoiceTaskDatelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskDatelog

`func (o *Company) SetInvoiceTaskDatelog(v bool)`

SetInvoiceTaskDatelog sets InvoiceTaskDatelog field to given value.

### HasInvoiceTaskDatelog

`func (o *Company) HasInvoiceTaskDatelog() bool`

HasInvoiceTaskDatelog returns a boolean if a field has been set.

### GetDefaultPasswordTimeout

`func (o *Company) GetDefaultPasswordTimeout() int32`

GetDefaultPasswordTimeout returns the DefaultPasswordTimeout field if non-nil, zero value otherwise.

### GetDefaultPasswordTimeoutOk

`func (o *Company) GetDefaultPasswordTimeoutOk() (*int32, bool)`

GetDefaultPasswordTimeoutOk returns a tuple with the DefaultPasswordTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordTimeout

`func (o *Company) SetDefaultPasswordTimeout(v int32)`

SetDefaultPasswordTimeout sets DefaultPasswordTimeout field to given value.

### HasDefaultPasswordTimeout

`func (o *Company) HasDefaultPasswordTimeout() bool`

HasDefaultPasswordTimeout returns a boolean if a field has been set.

### GetShowTaskEndDate

`func (o *Company) GetShowTaskEndDate() bool`

GetShowTaskEndDate returns the ShowTaskEndDate field if non-nil, zero value otherwise.

### GetShowTaskEndDateOk

`func (o *Company) GetShowTaskEndDateOk() (*bool, bool)`

GetShowTaskEndDateOk returns a tuple with the ShowTaskEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTaskEndDate

`func (o *Company) SetShowTaskEndDate(v bool)`

SetShowTaskEndDate sets ShowTaskEndDate field to given value.

### HasShowTaskEndDate

`func (o *Company) HasShowTaskEndDate() bool`

HasShowTaskEndDate returns a boolean if a field has been set.

### GetMarkdownEnabled

`func (o *Company) GetMarkdownEnabled() bool`

GetMarkdownEnabled returns the MarkdownEnabled field if non-nil, zero value otherwise.

### GetMarkdownEnabledOk

`func (o *Company) GetMarkdownEnabledOk() (*bool, bool)`

GetMarkdownEnabledOk returns a tuple with the MarkdownEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdownEnabled

`func (o *Company) SetMarkdownEnabled(v bool)`

SetMarkdownEnabled sets MarkdownEnabled field to given value.

### HasMarkdownEnabled

`func (o *Company) HasMarkdownEnabled() bool`

HasMarkdownEnabled returns a boolean if a field has been set.

### GetReportIncludeDrafts

`func (o *Company) GetReportIncludeDrafts() bool`

GetReportIncludeDrafts returns the ReportIncludeDrafts field if non-nil, zero value otherwise.

### GetReportIncludeDraftsOk

`func (o *Company) GetReportIncludeDraftsOk() (*bool, bool)`

GetReportIncludeDraftsOk returns a tuple with the ReportIncludeDrafts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIncludeDrafts

`func (o *Company) SetReportIncludeDrafts(v bool)`

SetReportIncludeDrafts sets ReportIncludeDrafts field to given value.

### HasReportIncludeDrafts

`func (o *Company) HasReportIncludeDrafts() bool`

HasReportIncludeDrafts returns a boolean if a field has been set.

### GetClientRegistrationFields

`func (o *Company) GetClientRegistrationFields() map[string]interface{}`

GetClientRegistrationFields returns the ClientRegistrationFields field if non-nil, zero value otherwise.

### GetClientRegistrationFieldsOk

`func (o *Company) GetClientRegistrationFieldsOk() (*map[string]interface{}, bool)`

GetClientRegistrationFieldsOk returns a tuple with the ClientRegistrationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRegistrationFields

`func (o *Company) SetClientRegistrationFields(v map[string]interface{})`

SetClientRegistrationFields sets ClientRegistrationFields field to given value.

### HasClientRegistrationFields

`func (o *Company) HasClientRegistrationFields() bool`

HasClientRegistrationFields returns a boolean if a field has been set.

### GetStopOnUnpaidRecurring

`func (o *Company) GetStopOnUnpaidRecurring() bool`

GetStopOnUnpaidRecurring returns the StopOnUnpaidRecurring field if non-nil, zero value otherwise.

### GetStopOnUnpaidRecurringOk

`func (o *Company) GetStopOnUnpaidRecurringOk() (*bool, bool)`

GetStopOnUnpaidRecurringOk returns a tuple with the StopOnUnpaidRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnUnpaidRecurring

`func (o *Company) SetStopOnUnpaidRecurring(v bool)`

SetStopOnUnpaidRecurring sets StopOnUnpaidRecurring field to given value.

### HasStopOnUnpaidRecurring

`func (o *Company) HasStopOnUnpaidRecurring() bool`

HasStopOnUnpaidRecurring returns a boolean if a field has been set.

### GetUseQuoteTermsOnConversion

`func (o *Company) GetUseQuoteTermsOnConversion() bool`

GetUseQuoteTermsOnConversion returns the UseQuoteTermsOnConversion field if non-nil, zero value otherwise.

### GetUseQuoteTermsOnConversionOk

`func (o *Company) GetUseQuoteTermsOnConversionOk() (*bool, bool)`

GetUseQuoteTermsOnConversionOk returns a tuple with the UseQuoteTermsOnConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseQuoteTermsOnConversion

`func (o *Company) SetUseQuoteTermsOnConversion(v bool)`

SetUseQuoteTermsOnConversion sets UseQuoteTermsOnConversion field to given value.

### HasUseQuoteTermsOnConversion

`func (o *Company) HasUseQuoteTermsOnConversion() bool`

HasUseQuoteTermsOnConversion returns a boolean if a field has been set.

### GetEnableApplyingPayments

`func (o *Company) GetEnableApplyingPayments() bool`

GetEnableApplyingPayments returns the EnableApplyingPayments field if non-nil, zero value otherwise.

### GetEnableApplyingPaymentsOk

`func (o *Company) GetEnableApplyingPaymentsOk() (*bool, bool)`

GetEnableApplyingPaymentsOk returns a tuple with the EnableApplyingPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApplyingPayments

`func (o *Company) SetEnableApplyingPayments(v bool)`

SetEnableApplyingPayments sets EnableApplyingPayments field to given value.

### HasEnableApplyingPayments

`func (o *Company) HasEnableApplyingPayments() bool`

HasEnableApplyingPayments returns a boolean if a field has been set.

### GetTrackInventory

`func (o *Company) GetTrackInventory() bool`

GetTrackInventory returns the TrackInventory field if non-nil, zero value otherwise.

### GetTrackInventoryOk

`func (o *Company) GetTrackInventoryOk() (*bool, bool)`

GetTrackInventoryOk returns a tuple with the TrackInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackInventory

`func (o *Company) SetTrackInventory(v bool)`

SetTrackInventory sets TrackInventory field to given value.

### HasTrackInventory

`func (o *Company) HasTrackInventory() bool`

HasTrackInventory returns a boolean if a field has been set.

### GetInventoryNotificationThreshold

`func (o *Company) GetInventoryNotificationThreshold() int32`

GetInventoryNotificationThreshold returns the InventoryNotificationThreshold field if non-nil, zero value otherwise.

### GetInventoryNotificationThresholdOk

`func (o *Company) GetInventoryNotificationThresholdOk() (*int32, bool)`

GetInventoryNotificationThresholdOk returns a tuple with the InventoryNotificationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryNotificationThreshold

`func (o *Company) SetInventoryNotificationThreshold(v int32)`

SetInventoryNotificationThreshold sets InventoryNotificationThreshold field to given value.

### HasInventoryNotificationThreshold

`func (o *Company) HasInventoryNotificationThreshold() bool`

HasInventoryNotificationThreshold returns a boolean if a field has been set.

### GetStockNotification

`func (o *Company) GetStockNotification() bool`

GetStockNotification returns the StockNotification field if non-nil, zero value otherwise.

### GetStockNotificationOk

`func (o *Company) GetStockNotificationOk() (*bool, bool)`

GetStockNotificationOk returns a tuple with the StockNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNotification

`func (o *Company) SetStockNotification(v bool)`

SetStockNotification sets StockNotification field to given value.

### HasStockNotification

`func (o *Company) HasStockNotification() bool`

HasStockNotification returns a boolean if a field has been set.

### GetMatomoUrl

`func (o *Company) GetMatomoUrl() string`

GetMatomoUrl returns the MatomoUrl field if non-nil, zero value otherwise.

### GetMatomoUrlOk

`func (o *Company) GetMatomoUrlOk() (*string, bool)`

GetMatomoUrlOk returns a tuple with the MatomoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatomoUrl

`func (o *Company) SetMatomoUrl(v string)`

SetMatomoUrl sets MatomoUrl field to given value.

### HasMatomoUrl

`func (o *Company) HasMatomoUrl() bool`

HasMatomoUrl returns a boolean if a field has been set.

### GetMatomoId

`func (o *Company) GetMatomoId() string`

GetMatomoId returns the MatomoId field if non-nil, zero value otherwise.

### GetMatomoIdOk

`func (o *Company) GetMatomoIdOk() (*string, bool)`

GetMatomoIdOk returns a tuple with the MatomoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatomoId

`func (o *Company) SetMatomoId(v string)`

SetMatomoId sets MatomoId field to given value.

### HasMatomoId

`func (o *Company) HasMatomoId() bool`

HasMatomoId returns a boolean if a field has been set.

### GetEnabledExpenseTaxRates

`func (o *Company) GetEnabledExpenseTaxRates() int32`

GetEnabledExpenseTaxRates returns the EnabledExpenseTaxRates field if non-nil, zero value otherwise.

### GetEnabledExpenseTaxRatesOk

`func (o *Company) GetEnabledExpenseTaxRatesOk() (*int32, bool)`

GetEnabledExpenseTaxRatesOk returns a tuple with the EnabledExpenseTaxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledExpenseTaxRates

`func (o *Company) SetEnabledExpenseTaxRates(v int32)`

SetEnabledExpenseTaxRates sets EnabledExpenseTaxRates field to given value.

### HasEnabledExpenseTaxRates

`func (o *Company) HasEnabledExpenseTaxRates() bool`

HasEnabledExpenseTaxRates returns a boolean if a field has been set.

### GetInvoiceTaskProject

`func (o *Company) GetInvoiceTaskProject() bool`

GetInvoiceTaskProject returns the InvoiceTaskProject field if non-nil, zero value otherwise.

### GetInvoiceTaskProjectOk

`func (o *Company) GetInvoiceTaskProjectOk() (*bool, bool)`

GetInvoiceTaskProjectOk returns a tuple with the InvoiceTaskProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskProject

`func (o *Company) SetInvoiceTaskProject(v bool)`

SetInvoiceTaskProject sets InvoiceTaskProject field to given value.

### HasInvoiceTaskProject

`func (o *Company) HasInvoiceTaskProject() bool`

HasInvoiceTaskProject returns a boolean if a field has been set.

### GetReportIncludeDeleted

`func (o *Company) GetReportIncludeDeleted() bool`

GetReportIncludeDeleted returns the ReportIncludeDeleted field if non-nil, zero value otherwise.

### GetReportIncludeDeletedOk

`func (o *Company) GetReportIncludeDeletedOk() (*bool, bool)`

GetReportIncludeDeletedOk returns a tuple with the ReportIncludeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIncludeDeleted

`func (o *Company) SetReportIncludeDeleted(v bool)`

SetReportIncludeDeleted sets ReportIncludeDeleted field to given value.

### HasReportIncludeDeleted

`func (o *Company) HasReportIncludeDeleted() bool`

HasReportIncludeDeleted returns a boolean if a field has been set.

### GetInvoiceTaskLock

`func (o *Company) GetInvoiceTaskLock() bool`

GetInvoiceTaskLock returns the InvoiceTaskLock field if non-nil, zero value otherwise.

### GetInvoiceTaskLockOk

`func (o *Company) GetInvoiceTaskLockOk() (*bool, bool)`

GetInvoiceTaskLockOk returns a tuple with the InvoiceTaskLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskLock

`func (o *Company) SetInvoiceTaskLock(v bool)`

SetInvoiceTaskLock sets InvoiceTaskLock field to given value.

### HasInvoiceTaskLock

`func (o *Company) HasInvoiceTaskLock() bool`

HasInvoiceTaskLock returns a boolean if a field has been set.

### GetConvertPaymentCurrency

`func (o *Company) GetConvertPaymentCurrency() bool`

GetConvertPaymentCurrency returns the ConvertPaymentCurrency field if non-nil, zero value otherwise.

### GetConvertPaymentCurrencyOk

`func (o *Company) GetConvertPaymentCurrencyOk() (*bool, bool)`

GetConvertPaymentCurrencyOk returns a tuple with the ConvertPaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertPaymentCurrency

`func (o *Company) SetConvertPaymentCurrency(v bool)`

SetConvertPaymentCurrency sets ConvertPaymentCurrency field to given value.

### HasConvertPaymentCurrency

`func (o *Company) HasConvertPaymentCurrency() bool`

HasConvertPaymentCurrency returns a boolean if a field has been set.

### GetConvertExpenseCurrency

`func (o *Company) GetConvertExpenseCurrency() bool`

GetConvertExpenseCurrency returns the ConvertExpenseCurrency field if non-nil, zero value otherwise.

### GetConvertExpenseCurrencyOk

`func (o *Company) GetConvertExpenseCurrencyOk() (*bool, bool)`

GetConvertExpenseCurrencyOk returns a tuple with the ConvertExpenseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertExpenseCurrency

`func (o *Company) SetConvertExpenseCurrency(v bool)`

SetConvertExpenseCurrency sets ConvertExpenseCurrency field to given value.

### HasConvertExpenseCurrency

`func (o *Company) HasConvertExpenseCurrency() bool`

HasConvertExpenseCurrency returns a boolean if a field has been set.

### GetNotifyVendorWhenPaid

`func (o *Company) GetNotifyVendorWhenPaid() bool`

GetNotifyVendorWhenPaid returns the NotifyVendorWhenPaid field if non-nil, zero value otherwise.

### GetNotifyVendorWhenPaidOk

`func (o *Company) GetNotifyVendorWhenPaidOk() (*bool, bool)`

GetNotifyVendorWhenPaidOk returns a tuple with the NotifyVendorWhenPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyVendorWhenPaid

`func (o *Company) SetNotifyVendorWhenPaid(v bool)`

SetNotifyVendorWhenPaid sets NotifyVendorWhenPaid field to given value.

### HasNotifyVendorWhenPaid

`func (o *Company) HasNotifyVendorWhenPaid() bool`

HasNotifyVendorWhenPaid returns a boolean if a field has been set.

### GetInvoiceTaskHours

`func (o *Company) GetInvoiceTaskHours() bool`

GetInvoiceTaskHours returns the InvoiceTaskHours field if non-nil, zero value otherwise.

### GetInvoiceTaskHoursOk

`func (o *Company) GetInvoiceTaskHoursOk() (*bool, bool)`

GetInvoiceTaskHoursOk returns a tuple with the InvoiceTaskHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskHours

`func (o *Company) SetInvoiceTaskHours(v bool)`

SetInvoiceTaskHours sets InvoiceTaskHours field to given value.

### HasInvoiceTaskHours

`func (o *Company) HasInvoiceTaskHours() bool`

HasInvoiceTaskHours returns a boolean if a field has been set.

### GetCalculateTaxes

`func (o *Company) GetCalculateTaxes() bool`

GetCalculateTaxes returns the CalculateTaxes field if non-nil, zero value otherwise.

### GetCalculateTaxesOk

`func (o *Company) GetCalculateTaxesOk() (*bool, bool)`

GetCalculateTaxesOk returns a tuple with the CalculateTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateTaxes

`func (o *Company) SetCalculateTaxes(v bool)`

SetCalculateTaxes sets CalculateTaxes field to given value.

### HasCalculateTaxes

`func (o *Company) HasCalculateTaxes() bool`

HasCalculateTaxes returns a boolean if a field has been set.

### GetTaxData

`func (o *Company) GetTaxData() map[string]interface{}`

GetTaxData returns the TaxData field if non-nil, zero value otherwise.

### GetTaxDataOk

`func (o *Company) GetTaxDataOk() (*map[string]interface{}, bool)`

GetTaxDataOk returns a tuple with the TaxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxData

`func (o *Company) SetTaxData(v map[string]interface{})`

SetTaxData sets TaxData field to given value.

### HasTaxData

`func (o *Company) HasTaxData() bool`

HasTaxData returns a boolean if a field has been set.

### GetEInvoiceCertificate

`func (o *Company) GetEInvoiceCertificate() string`

GetEInvoiceCertificate returns the EInvoiceCertificate field if non-nil, zero value otherwise.

### GetEInvoiceCertificateOk

`func (o *Company) GetEInvoiceCertificateOk() (*string, bool)`

GetEInvoiceCertificateOk returns a tuple with the EInvoiceCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoiceCertificate

`func (o *Company) SetEInvoiceCertificate(v string)`

SetEInvoiceCertificate sets EInvoiceCertificate field to given value.

### HasEInvoiceCertificate

`func (o *Company) HasEInvoiceCertificate() bool`

HasEInvoiceCertificate returns a boolean if a field has been set.

### GetEInvoiceCertificatePassphrase

`func (o *Company) GetEInvoiceCertificatePassphrase() string`

GetEInvoiceCertificatePassphrase returns the EInvoiceCertificatePassphrase field if non-nil, zero value otherwise.

### GetEInvoiceCertificatePassphraseOk

`func (o *Company) GetEInvoiceCertificatePassphraseOk() (*string, bool)`

GetEInvoiceCertificatePassphraseOk returns a tuple with the EInvoiceCertificatePassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoiceCertificatePassphrase

`func (o *Company) SetEInvoiceCertificatePassphrase(v string)`

SetEInvoiceCertificatePassphrase sets EInvoiceCertificatePassphrase field to given value.

### HasEInvoiceCertificatePassphrase

`func (o *Company) HasEInvoiceCertificatePassphrase() bool`

HasEInvoiceCertificatePassphrase returns a boolean if a field has been set.

### GetOriginTaxData

`func (o *Company) GetOriginTaxData() map[string]interface{}`

GetOriginTaxData returns the OriginTaxData field if non-nil, zero value otherwise.

### GetOriginTaxDataOk

`func (o *Company) GetOriginTaxDataOk() (*map[string]interface{}, bool)`

GetOriginTaxDataOk returns a tuple with the OriginTaxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginTaxData

`func (o *Company) SetOriginTaxData(v map[string]interface{})`

SetOriginTaxData sets OriginTaxData field to given value.

### HasOriginTaxData

`func (o *Company) HasOriginTaxData() bool`

HasOriginTaxData returns a boolean if a field has been set.

### GetInvoiceTaskProjectHeader

`func (o *Company) GetInvoiceTaskProjectHeader() bool`

GetInvoiceTaskProjectHeader returns the InvoiceTaskProjectHeader field if non-nil, zero value otherwise.

### GetInvoiceTaskProjectHeaderOk

`func (o *Company) GetInvoiceTaskProjectHeaderOk() (*bool, bool)`

GetInvoiceTaskProjectHeaderOk returns a tuple with the InvoiceTaskProjectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskProjectHeader

`func (o *Company) SetInvoiceTaskProjectHeader(v bool)`

SetInvoiceTaskProjectHeader sets InvoiceTaskProjectHeader field to given value.

### HasInvoiceTaskProjectHeader

`func (o *Company) HasInvoiceTaskProjectHeader() bool`

HasInvoiceTaskProjectHeader returns a boolean if a field has been set.

### GetInvoiceTaskItemDescription

`func (o *Company) GetInvoiceTaskItemDescription() bool`

GetInvoiceTaskItemDescription returns the InvoiceTaskItemDescription field if non-nil, zero value otherwise.

### GetInvoiceTaskItemDescriptionOk

`func (o *Company) GetInvoiceTaskItemDescriptionOk() (*bool, bool)`

GetInvoiceTaskItemDescriptionOk returns a tuple with the InvoiceTaskItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTaskItemDescription

`func (o *Company) SetInvoiceTaskItemDescription(v bool)`

SetInvoiceTaskItemDescription sets InvoiceTaskItemDescription field to given value.

### HasInvoiceTaskItemDescription

`func (o *Company) HasInvoiceTaskItemDescription() bool`

HasInvoiceTaskItemDescription returns a boolean if a field has been set.

### GetExpenseMailbox

`func (o *Company) GetExpenseMailbox() string`

GetExpenseMailbox returns the ExpenseMailbox field if non-nil, zero value otherwise.

### GetExpenseMailboxOk

`func (o *Company) GetExpenseMailboxOk() (*string, bool)`

GetExpenseMailboxOk returns a tuple with the ExpenseMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseMailbox

`func (o *Company) SetExpenseMailbox(v string)`

SetExpenseMailbox sets ExpenseMailbox field to given value.

### HasExpenseMailbox

`func (o *Company) HasExpenseMailbox() bool`

HasExpenseMailbox returns a boolean if a field has been set.

### GetExpenseMailboxActive

`func (o *Company) GetExpenseMailboxActive() bool`

GetExpenseMailboxActive returns the ExpenseMailboxActive field if non-nil, zero value otherwise.

### GetExpenseMailboxActiveOk

`func (o *Company) GetExpenseMailboxActiveOk() (*bool, bool)`

GetExpenseMailboxActiveOk returns a tuple with the ExpenseMailboxActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseMailboxActive

`func (o *Company) SetExpenseMailboxActive(v bool)`

SetExpenseMailboxActive sets ExpenseMailboxActive field to given value.

### HasExpenseMailboxActive

`func (o *Company) HasExpenseMailboxActive() bool`

HasExpenseMailboxActive returns a boolean if a field has been set.

### GetInboundMailboxAllowCompanyUsers

`func (o *Company) GetInboundMailboxAllowCompanyUsers() bool`

GetInboundMailboxAllowCompanyUsers returns the InboundMailboxAllowCompanyUsers field if non-nil, zero value otherwise.

### GetInboundMailboxAllowCompanyUsersOk

`func (o *Company) GetInboundMailboxAllowCompanyUsersOk() (*bool, bool)`

GetInboundMailboxAllowCompanyUsersOk returns a tuple with the InboundMailboxAllowCompanyUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMailboxAllowCompanyUsers

`func (o *Company) SetInboundMailboxAllowCompanyUsers(v bool)`

SetInboundMailboxAllowCompanyUsers sets InboundMailboxAllowCompanyUsers field to given value.

### HasInboundMailboxAllowCompanyUsers

`func (o *Company) HasInboundMailboxAllowCompanyUsers() bool`

HasInboundMailboxAllowCompanyUsers returns a boolean if a field has been set.

### GetInboundMailboxAllowVendors

`func (o *Company) GetInboundMailboxAllowVendors() bool`

GetInboundMailboxAllowVendors returns the InboundMailboxAllowVendors field if non-nil, zero value otherwise.

### GetInboundMailboxAllowVendorsOk

`func (o *Company) GetInboundMailboxAllowVendorsOk() (*bool, bool)`

GetInboundMailboxAllowVendorsOk returns a tuple with the InboundMailboxAllowVendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMailboxAllowVendors

`func (o *Company) SetInboundMailboxAllowVendors(v bool)`

SetInboundMailboxAllowVendors sets InboundMailboxAllowVendors field to given value.

### HasInboundMailboxAllowVendors

`func (o *Company) HasInboundMailboxAllowVendors() bool`

HasInboundMailboxAllowVendors returns a boolean if a field has been set.

### GetInboundMailboxAllowClients

`func (o *Company) GetInboundMailboxAllowClients() bool`

GetInboundMailboxAllowClients returns the InboundMailboxAllowClients field if non-nil, zero value otherwise.

### GetInboundMailboxAllowClientsOk

`func (o *Company) GetInboundMailboxAllowClientsOk() (*bool, bool)`

GetInboundMailboxAllowClientsOk returns a tuple with the InboundMailboxAllowClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMailboxAllowClients

`func (o *Company) SetInboundMailboxAllowClients(v bool)`

SetInboundMailboxAllowClients sets InboundMailboxAllowClients field to given value.

### HasInboundMailboxAllowClients

`func (o *Company) HasInboundMailboxAllowClients() bool`

HasInboundMailboxAllowClients returns a boolean if a field has been set.

### GetInboundMailboxAllowUnknown

`func (o *Company) GetInboundMailboxAllowUnknown() bool`

GetInboundMailboxAllowUnknown returns the InboundMailboxAllowUnknown field if non-nil, zero value otherwise.

### GetInboundMailboxAllowUnknownOk

`func (o *Company) GetInboundMailboxAllowUnknownOk() (*bool, bool)`

GetInboundMailboxAllowUnknownOk returns a tuple with the InboundMailboxAllowUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMailboxAllowUnknown

`func (o *Company) SetInboundMailboxAllowUnknown(v bool)`

SetInboundMailboxAllowUnknown sets InboundMailboxAllowUnknown field to given value.

### HasInboundMailboxAllowUnknown

`func (o *Company) HasInboundMailboxAllowUnknown() bool`

HasInboundMailboxAllowUnknown returns a boolean if a field has been set.

### GetInboundMailboxWhitelist

`func (o *Company) GetInboundMailboxWhitelist() string`

GetInboundMailboxWhitelist returns the InboundMailboxWhitelist field if non-nil, zero value otherwise.

### GetInboundMailboxWhitelistOk

`func (o *Company) GetInboundMailboxWhitelistOk() (*string, bool)`

GetInboundMailboxWhitelistOk returns a tuple with the InboundMailboxWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMailboxWhitelist

`func (o *Company) SetInboundMailboxWhitelist(v string)`

SetInboundMailboxWhitelist sets InboundMailboxWhitelist field to given value.

### HasInboundMailboxWhitelist

`func (o *Company) HasInboundMailboxWhitelist() bool`

HasInboundMailboxWhitelist returns a boolean if a field has been set.

### GetInboundMailboxBlacklist

`func (o *Company) GetInboundMailboxBlacklist() string`

GetInboundMailboxBlacklist returns the InboundMailboxBlacklist field if non-nil, zero value otherwise.

### GetInboundMailboxBlacklistOk

`func (o *Company) GetInboundMailboxBlacklistOk() (*string, bool)`

GetInboundMailboxBlacklistOk returns a tuple with the InboundMailboxBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMailboxBlacklist

`func (o *Company) SetInboundMailboxBlacklist(v string)`

SetInboundMailboxBlacklist sets InboundMailboxBlacklist field to given value.

### HasInboundMailboxBlacklist

`func (o *Company) HasInboundMailboxBlacklist() bool`

HasInboundMailboxBlacklist returns a boolean if a field has been set.

### GetSmtpHost

`func (o *Company) GetSmtpHost() string`

GetSmtpHost returns the SmtpHost field if non-nil, zero value otherwise.

### GetSmtpHostOk

`func (o *Company) GetSmtpHostOk() (*string, bool)`

GetSmtpHostOk returns a tuple with the SmtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpHost

`func (o *Company) SetSmtpHost(v string)`

SetSmtpHost sets SmtpHost field to given value.

### HasSmtpHost

`func (o *Company) HasSmtpHost() bool`

HasSmtpHost returns a boolean if a field has been set.

### GetSmtpPort

`func (o *Company) GetSmtpPort() int32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *Company) GetSmtpPortOk() (*int32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *Company) SetSmtpPort(v int32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *Company) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpEncryption

`func (o *Company) GetSmtpEncryption() string`

GetSmtpEncryption returns the SmtpEncryption field if non-nil, zero value otherwise.

### GetSmtpEncryptionOk

`func (o *Company) GetSmtpEncryptionOk() (*string, bool)`

GetSmtpEncryptionOk returns a tuple with the SmtpEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEncryption

`func (o *Company) SetSmtpEncryption(v string)`

SetSmtpEncryption sets SmtpEncryption field to given value.

### HasSmtpEncryption

`func (o *Company) HasSmtpEncryption() bool`

HasSmtpEncryption returns a boolean if a field has been set.

### SetSmtpEncryptionNil

`func (o *Company) SetSmtpEncryptionNil(b bool)`

 SetSmtpEncryptionNil sets the value for SmtpEncryption to be an explicit nil

### UnsetSmtpEncryption
`func (o *Company) UnsetSmtpEncryption()`

UnsetSmtpEncryption ensures that no value is present for SmtpEncryption, not even an explicit nil
### GetSmtpLocalDomain

`func (o *Company) GetSmtpLocalDomain() string`

GetSmtpLocalDomain returns the SmtpLocalDomain field if non-nil, zero value otherwise.

### GetSmtpLocalDomainOk

`func (o *Company) GetSmtpLocalDomainOk() (*string, bool)`

GetSmtpLocalDomainOk returns a tuple with the SmtpLocalDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpLocalDomain

`func (o *Company) SetSmtpLocalDomain(v string)`

SetSmtpLocalDomain sets SmtpLocalDomain field to given value.

### HasSmtpLocalDomain

`func (o *Company) HasSmtpLocalDomain() bool`

HasSmtpLocalDomain returns a boolean if a field has been set.

### GetSmtpVerifyPeer

`func (o *Company) GetSmtpVerifyPeer() bool`

GetSmtpVerifyPeer returns the SmtpVerifyPeer field if non-nil, zero value otherwise.

### GetSmtpVerifyPeerOk

`func (o *Company) GetSmtpVerifyPeerOk() (*bool, bool)`

GetSmtpVerifyPeerOk returns a tuple with the SmtpVerifyPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpVerifyPeer

`func (o *Company) SetSmtpVerifyPeer(v bool)`

SetSmtpVerifyPeer sets SmtpVerifyPeer field to given value.

### HasSmtpVerifyPeer

`func (o *Company) HasSmtpVerifyPeer() bool`

HasSmtpVerifyPeer returns a boolean if a field has been set.

### GetEInvoice

`func (o *Company) GetEInvoice() map[string]interface{}`

GetEInvoice returns the EInvoice field if non-nil, zero value otherwise.

### GetEInvoiceOk

`func (o *Company) GetEInvoiceOk() (*map[string]interface{}, bool)`

GetEInvoiceOk returns a tuple with the EInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInvoice

`func (o *Company) SetEInvoice(v map[string]interface{})`

SetEInvoice sets EInvoice field to given value.

### HasEInvoice

`func (o *Company) HasEInvoice() bool`

HasEInvoice returns a boolean if a field has been set.

### GetLegalEntityId

`func (o *Company) GetLegalEntityId() int32`

GetLegalEntityId returns the LegalEntityId field if non-nil, zero value otherwise.

### GetLegalEntityIdOk

`func (o *Company) GetLegalEntityIdOk() (*int32, bool)`

GetLegalEntityIdOk returns a tuple with the LegalEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalEntityId

`func (o *Company) SetLegalEntityId(v int32)`

SetLegalEntityId sets LegalEntityId field to given value.

### HasLegalEntityId

`func (o *Company) HasLegalEntityId() bool`

HasLegalEntityId returns a boolean if a field has been set.

### GetSettings

`func (o *Company) GetSettings() CompanySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Company) GetSettingsOk() (*CompanySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Company) SetSettings(v CompanySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Company) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


