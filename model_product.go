package openapi

import (
	"encoding/json"
)

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product struct for Product
type Product struct {
	// The hashed product ID.
	Id *string `json:"id,omitempty"`
	// The hashed ID of the user that created this product.
	UserId *string `json:"user_id,omitempty"`
	// The hashed ID of the user assigned to this product.
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The hashed ID of the project that this product is associated with.
	ProjectId *string `json:"project_id,omitempty"`
	// The hashed ID of the vendor that this product is associated with.
	VendorId *string `json:"vendor_id,omitempty"`
	// Custom value field 1.
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// Custom value field 2.
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// Custom value field 3.
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// Custom value field 4.
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The product key.
	ProductKey *string `json:"product_key,omitempty"`
	// Notes about the product.
	Notes *string `json:"notes,omitempty"`
	// The cost of the product. (Your purchase price for this product)
	Cost *float64 `json:"cost,omitempty"`
	// The price of the product that you are charging.
	Price *float64 `json:"price,omitempty"`
	// The quantity of the product. (used as a default)
	Quantity *float64 `json:"quantity,omitempty"`
	// The name of tax 1.
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The rate of tax 1.
	TaxRate1 *float64 `json:"tax_rate1,omitempty"`
	// The name of tax 2.
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The rate of tax 2.
	TaxRate2 *float64 `json:"tax_rate2,omitempty"`
	// The name of tax 3.
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The rate of tax 3.
	TaxRate3 *float64 `json:"tax_rate3,omitempty"`
	// The timestamp when the product was archived.
	ArchivedAt *int32 `json:"archived_at,omitempty"`
	// The timestamp when the product was created.
	CreatedAt *int32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// Boolean flag determining if the product has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The quantity of the product that is currently in stock
	InStockQuantity *int32 `json:"in_stock_quantity,omitempty"`
	// Indicates whether stock notifications are enabled for this product
	StockNotification *bool `json:"stock_notification,omitempty"`
	// The minimum quantity threshold for which stock notifications will be triggered
	StockNotificationThreshold *int32 `json:"stock_notification_threshold,omitempty"`
	// The maximum quantity that can be ordered for this product
	MaxQuantity *int32 `json:"max_quantity,omitempty"`
	// The URL of the product image
	ProductImage *string `json:"product_image,omitempty"`
	// The tax category id for this product.'  The following constants are available (default = '1')  ``` PRODUCT_TYPE_PHYSICAL = '1' PRODUCT_TYPE_SERVICE = '2' PRODUCT_TYPE_DIGITAL = '3' PRODUCT_TYPE_SHIPPING = '4' PRODUCT_TYPE_EXEMPT = '5' PRODUCT_TYPE_REDUCED_TAX = '6' PRODUCT_TYPE_OVERRIDE_TAX = '7' PRODUCT_TYPE_ZERO_RATED = '8' PRODUCT_TYPE_REVERSE_TAX = '9' ```
	TaxId *string `json:"tax_id,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct() *Product {
	this := Product{}
	var inStockQuantity int32 = 0
	this.InStockQuantity = &inStockQuantity
	var stockNotification bool = true
	this.StockNotification = &stockNotification
	var stockNotificationThreshold int32 = 0
	this.StockNotificationThreshold = &stockNotificationThreshold
	var taxId string = "1"
	this.TaxId = &taxId
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	var inStockQuantity int32 = 0
	this.InStockQuantity = &inStockQuantity
	var stockNotification bool = true
	this.StockNotification = &stockNotification
	var stockNotificationThreshold int32 = 0
	this.StockNotificationThreshold = &stockNotificationThreshold
	var taxId string = "1"
	this.TaxId = &taxId
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Product) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Product) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Product) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Product) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Product) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Product) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Product) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Product) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *Product) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *Product) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *Product) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *Product) SetVendorId(v string) {
	o.VendorId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *Product) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Product) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *Product) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *Product) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Product) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *Product) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *Product) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Product) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *Product) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *Product) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Product) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *Product) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *Product) GetProductKey() string {
	if o == nil || IsNil(o.ProductKey) {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProductKey) {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *Product) HasProductKey() bool {
	if o != nil && !IsNil(o.ProductKey) {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *Product) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Product) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Product) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Product) SetNotes(v string) {
	o.Notes = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *Product) GetCost() float64 {
	if o == nil || IsNil(o.Cost) {
		var ret float64
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCostOk() (*float64, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *Product) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float64 and assigns it to the Cost field.
func (o *Product) SetCost(v float64) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Product) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Product) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *Product) SetPrice(v float64) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Product) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Product) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *Product) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *Product) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *Product) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *Product) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *Product) GetTaxRate1() float64 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float64
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTaxRate1Ok() (*float64, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *Product) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float64 and assigns it to the TaxRate1 field.
func (o *Product) SetTaxRate1(v float64) {
	o.TaxRate1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *Product) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *Product) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *Product) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *Product) GetTaxRate2() float64 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float64
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTaxRate2Ok() (*float64, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *Product) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float64 and assigns it to the TaxRate2 field.
func (o *Product) SetTaxRate2(v float64) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *Product) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *Product) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *Product) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *Product) GetTaxRate3() float64 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float64
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTaxRate3Ok() (*float64, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *Product) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float64 and assigns it to the TaxRate3 field.
func (o *Product) SetTaxRate3(v float64) {
	o.TaxRate3 = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Product) GetArchivedAt() int32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret int32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetArchivedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Product) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given int32 and assigns it to the ArchivedAt field.
func (o *Product) SetArchivedAt(v int32) {
	o.ArchivedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Product) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Product) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *Product) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Product) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Product) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *Product) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Product) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Product) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Product) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetInStockQuantity returns the InStockQuantity field value if set, zero value otherwise.
func (o *Product) GetInStockQuantity() int32 {
	if o == nil || IsNil(o.InStockQuantity) {
		var ret int32
		return ret
	}
	return *o.InStockQuantity
}

// GetInStockQuantityOk returns a tuple with the InStockQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetInStockQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.InStockQuantity) {
		return nil, false
	}
	return o.InStockQuantity, true
}

// HasInStockQuantity returns a boolean if a field has been set.
func (o *Product) HasInStockQuantity() bool {
	if o != nil && !IsNil(o.InStockQuantity) {
		return true
	}

	return false
}

// SetInStockQuantity gets a reference to the given int32 and assigns it to the InStockQuantity field.
func (o *Product) SetInStockQuantity(v int32) {
	o.InStockQuantity = &v
}

// GetStockNotification returns the StockNotification field value if set, zero value otherwise.
func (o *Product) GetStockNotification() bool {
	if o == nil || IsNil(o.StockNotification) {
		var ret bool
		return ret
	}
	return *o.StockNotification
}

// GetStockNotificationOk returns a tuple with the StockNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStockNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.StockNotification) {
		return nil, false
	}
	return o.StockNotification, true
}

// HasStockNotification returns a boolean if a field has been set.
func (o *Product) HasStockNotification() bool {
	if o != nil && !IsNil(o.StockNotification) {
		return true
	}

	return false
}

// SetStockNotification gets a reference to the given bool and assigns it to the StockNotification field.
func (o *Product) SetStockNotification(v bool) {
	o.StockNotification = &v
}

// GetStockNotificationThreshold returns the StockNotificationThreshold field value if set, zero value otherwise.
func (o *Product) GetStockNotificationThreshold() int32 {
	if o == nil || IsNil(o.StockNotificationThreshold) {
		var ret int32
		return ret
	}
	return *o.StockNotificationThreshold
}

// GetStockNotificationThresholdOk returns a tuple with the StockNotificationThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetStockNotificationThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.StockNotificationThreshold) {
		return nil, false
	}
	return o.StockNotificationThreshold, true
}

// HasStockNotificationThreshold returns a boolean if a field has been set.
func (o *Product) HasStockNotificationThreshold() bool {
	if o != nil && !IsNil(o.StockNotificationThreshold) {
		return true
	}

	return false
}

// SetStockNotificationThreshold gets a reference to the given int32 and assigns it to the StockNotificationThreshold field.
func (o *Product) SetStockNotificationThreshold(v int32) {
	o.StockNotificationThreshold = &v
}

// GetMaxQuantity returns the MaxQuantity field value if set, zero value otherwise.
func (o *Product) GetMaxQuantity() int32 {
	if o == nil || IsNil(o.MaxQuantity) {
		var ret int32
		return ret
	}
	return *o.MaxQuantity
}

// GetMaxQuantityOk returns a tuple with the MaxQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetMaxQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxQuantity) {
		return nil, false
	}
	return o.MaxQuantity, true
}

// HasMaxQuantity returns a boolean if a field has been set.
func (o *Product) HasMaxQuantity() bool {
	if o != nil && !IsNil(o.MaxQuantity) {
		return true
	}

	return false
}

// SetMaxQuantity gets a reference to the given int32 and assigns it to the MaxQuantity field.
func (o *Product) SetMaxQuantity(v int32) {
	o.MaxQuantity = &v
}

// GetProductImage returns the ProductImage field value if set, zero value otherwise.
func (o *Product) GetProductImage() string {
	if o == nil || IsNil(o.ProductImage) {
		var ret string
		return ret
	}
	return *o.ProductImage
}

// GetProductImageOk returns a tuple with the ProductImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetProductImageOk() (*string, bool) {
	if o == nil || IsNil(o.ProductImage) {
		return nil, false
	}
	return o.ProductImage, true
}

// HasProductImage returns a boolean if a field has been set.
func (o *Product) HasProductImage() bool {
	if o != nil && !IsNil(o.ProductImage) {
		return true
	}

	return false
}

// SetProductImage gets a reference to the given string and assigns it to the ProductImage field.
func (o *Product) SetProductImage(v string) {
	o.ProductImage = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *Product) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *Product) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *Product) SetTaxId(v string) {
	o.TaxId = &v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Product) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
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
	if !IsNil(o.ProductKey) {
		toSerialize["product_key"] = o.ProductKey
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.TaxName1) {
		toSerialize["tax_name1"] = o.TaxName1
	}
	if !IsNil(o.TaxRate1) {
		toSerialize["tax_rate1"] = o.TaxRate1
	}
	if !IsNil(o.TaxName2) {
		toSerialize["tax_name2"] = o.TaxName2
	}
	if !IsNil(o.TaxRate2) {
		toSerialize["tax_rate2"] = o.TaxRate2
	}
	if !IsNil(o.TaxName3) {
		toSerialize["tax_name3"] = o.TaxName3
	}
	if !IsNil(o.TaxRate3) {
		toSerialize["tax_rate3"] = o.TaxRate3
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.InStockQuantity) {
		toSerialize["in_stock_quantity"] = o.InStockQuantity
	}
	if !IsNil(o.StockNotification) {
		toSerialize["stock_notification"] = o.StockNotification
	}
	if !IsNil(o.StockNotificationThreshold) {
		toSerialize["stock_notification_threshold"] = o.StockNotificationThreshold
	}
	if !IsNil(o.MaxQuantity) {
		toSerialize["max_quantity"] = o.MaxQuantity
	}
	if !IsNil(o.ProductImage) {
		toSerialize["product_image"] = o.ProductImage
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
