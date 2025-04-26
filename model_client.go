package openapi

import (
	"encoding/json"
)

// checks if the Client type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Client{}

// Client struct for Client
type Client struct {
	// The unique identifier of the client
	Id *string `json:"id,omitempty"`
	// The unique identifier of the user who created the client
	UserId *string `json:"user_id,omitempty"`
	// The unique identifier of the user who has been assigned the client
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The name of the client company or organization
	Name *string `json:"name,omitempty"`
	// The website URL of the client company or organization
	Website *string `json:"website,omitempty"`
	// Notes that are only visible to the user who created the client
	PrivateNotes *string `json:"private_notes,omitempty"`
	// A unique hash value for the client
	ClientHash *string `json:"client_hash,omitempty"`
	// The unique identifier of the industry the client operates in
	IndustryId *float32 `json:"industry_id,omitempty"`
	// The unique identifier for the size category of the client company or organization
	SizeId *float32 `json:"size_id,omitempty"`
	// First line of the client's address
	Address1 *string `json:"address1,omitempty"`
	// Second line of the client's address, if needed
	Address2 *string `json:"address2,omitempty"`
	// The city the client is located in
	City *string `json:"city,omitempty"`
	// The state, province, or locality the client is located in
	State *string `json:"state,omitempty"`
	// The postal code or ZIP code of the client
	PostalCode *string `json:"postal_code,omitempty"`
	// The client's phone number
	Phone *string `json:"phone,omitempty"`
	// Available country types: - 4: Afghanistan (AF/AFG) - 8: Albania (AL/ALB) - 10: Antarctica (AQ/ATA) - 12: Algeria (DZ/DZA) - 16: American Samoa (AS/ASM) - 20: Andorra (AD/AND) - 24: Angola (AO/AGO) - 28: Antigua and Barbuda (AG/ATG) - 31: Azerbaijan (AZ/AZE) - 32: Argentina (AR/ARG) - 36: Australia (AU/AUS) - 40: Austria (AT/AUT) - 44: Bahamas (BS/BHS) - 48: Bahrain (BH/BHR) - 50: Bangladesh (BD/BGD) - 51: Armenia (AM/ARM) - 52: Barbados (BB/BRB) - 56: Belgium (BE/BEL) - 60: Bermuda (BM/BMU) - 64: Bhutan (BT/BTN) - 68: Bolivia, Plurinational State of (BO/BOL) - 70: Bosnia and Herzegovina (BA/BIH) - 72: Botswana (BW/BWA) - 74: Bouvet Island (BV/BVT) - 76: Brazil (BR/BRA) - 84: Belize (BZ/BLZ) - 86: British Indian Ocean Territory (IO/IOT) - 90: Solomon Islands (SB/SLB) - 92: Virgin Islands, British (VG/VGB) - 96: Brunei Darussalam (BN/BRN) - 100: Bulgaria (BG/BGR) - 104: Myanmar (MM/MMR) - 108: Burundi (BI/BDI) - 112: Belarus (BY/BLR) - 116: Cambodia (KH/KHM) - 120: Cameroon (CM/CMR) - 124: Canada (CA/CAN) - 132: Cape Verde (CV/CPV) - 136: Cayman Islands (KY/CYM) - 140: Central African Republic (CF/CAF) - 144: Sri Lanka (LK/LKA) - 148: Chad (TD/TCD) - 152: Chile (CL/CHL) - 156: China (CN/CHN) - 158: Taiwan (TW/TWN) - 162: Christmas Island (CX/CXR) - 166: Cocos (Keeling) Islands (CC/CCK) - 170: Colombia (CO/COL) - 174: Comoros (KM/COM) - 175: Mayotte (YT/MYT) - 178: Congo (CG/COG) - 180: Congo, the Democratic Republic of the (CD/COD) - 184: Cook Islands (CK/COK) - 188: Costa Rica (CR/CRI) - 191: Croatia (HR/HRV) - 192: Cuba (CU/CUB) - 196: Cyprus (CY/CYP) - 203: Czech Republic (CZ/CZE) - 204: Benin (BJ/BEN) - 208: Denmark (DK/DNK) - 212: Dominica (DM/DMA) - 214: Dominican Republic (DO/DOM) - 218: Ecuador (EC/ECU) - 222: El Salvador (SV/SLV) - 226: Equatorial Guinea (GQ/GNQ) - 231: Ethiopia (ET/ETH) - 232: Eritrea (ER/ERI) - 233: Estonia (EE/EST) - 234: Faroe Islands (FO/FRO) - 238: Falkland Islands (Malvinas) (FK/FLK) - 239: South Georgia and the South Sandwich Islands (GS/SGS) - 242: Fiji (FJ/FJI) - 246: Finland (FI/FIN) - 248: Åland Islands (AX/ALA) - 250: France (FR/FRA) - 254: French Guiana (GF/GUF) - 258: French Polynesia (PF/PYF) - 260: French Southern Territories (TF/ATF) - 262: Djibouti (DJ/DJI) - 266: Gabon (GA/GAB) - 268: Georgia (GE/GEO) - 270: Gambia (GM/GMB) - 275: Palestine (PS/PSE) - 276: Germany (DE/DEU) - 288: Ghana (GH/GHA) - 292: Gibraltar (GI/GIB) - 296: Kiribati (KI/KIR) - 300: Greece (GR/GRC) - 304: Greenland (GL/GRL) - 308: Grenada (GD/GRD) - 312: Guadeloupe (GP/GLP) - 316: Guam (GU/GUM) - 320: Guatemala (GT/GTM) - 324: Guinea (GN/GIN) - 328: Guyana (GY/GUY) - 332: Haiti (HT/HTI) - 334: Heard Island and McDonald Islands (HM/HMD) - 336: Holy See (Vatican City State) (VA/VAT) - 340: Honduras (HN/HND) - 344: Hong Kong (HK/HKG) - 348: Hungary (HU/HUN) - 352: Iceland (IS/ISL) - 356: India (IN/IND) - 360: Indonesia (ID/IDN) - 364: Iran, Islamic Republic of (IR/IRN) - 368: Iraq (IQ/IRQ) - 372: Ireland (IE/IRL) - 376: Israel (IL/ISR) - 380: Italy (IT/ITA) - 384: Côte d'Ivoire (CI/CIV) - 388: Jamaica (JM/JAM) - 392: Japan (JP/JPN) - 398: Kazakhstan (KZ/KAZ) - 400: Jordan (JO/JOR) - 404: Kenya (KE/KEN) - 408: Korea, Democratic People's Republic of (KP/PRK) - 410: Korea, Republic of (KR/KOR) - 414: Kuwait (KW/KWT) - 417: Kyrgyzstan (KG/KGZ) - 418: Lao People's Democratic Republic (LA/LAO) - 422: Lebanon (LB/LBN) - 426: Lesotho (LS/LSO) - 428: Latvia (LV/LVA) - 430: Liberia (LR/LBR) - 434: Libya (LY/LBY) - 438: Liechtenstein (LI/LIE) - 440: Lithuania (LT/LTU) - 442: Luxembourg (LU/LUX) - 446: Macao (MO/MAC) - 450: Madagascar (MG/MDG) - 454: Malawi (MW/MWI) - 458: Malaysia (MY/MYS) - 462: Maldives (MV/MDV) - 466: Mali (ML/MLI) - 470: Malta (MT/MLT) - 474: Martinique (MQ/MTQ) - 478: Mauritania (MR/MRT) - 480: Mauritius (MU/MUS) - 484: Mexico (MX/MEX) - 492: Monaco (MC/MCO) - 496: Mongolia (MN/MNG) - 498: Moldova, Republic of (MD/MDA) - 499: Montenegro (ME/MNE) - 500: Montserrat (MS/MSR) - 504: Morocco (MA/MAR) - 508: Mozambique (MZ/MOZ) - 512: Oman (OM/OMN) - 516: Namibia (NA/NAM) - 520: Nauru (NR/NRU) - 524: Nepal (NP/NPL) - 528: Netherlands (NL/NLD) - 531: Curaçao (CW/CUW) - 533: Aruba (AW/ABW) - 534: Sint Maarten (Dutch part) (SX/SXM) - 535: Bonaire, Sint Eustatius and Saba (BQ/BES) - 540: New Caledonia (NC/NCL) - 548: Vanuatu (VU/VUT) - 554: New Zealand (NZ/NZL) - 558: Nicaragua (NI/NIC) - 562: Niger (NE/NER) - 566: Nigeria (NG/NGA) - 570: Niue (NU/NIU) - 574: Norfolk Island (NF/NFK) - 578: Norway (NO/NOR) - 580: Northern Mariana Islands (MP/MNP) - 581: United States Minor Outlying Islands (UM/UMI) - 583: Micronesia, Federated States of (FM/FSM) - 584: Marshall Islands (MH/MHL) - 585: Palau (PW/PLW) - 586: Pakistan (PK/PAK) - 591: Panama (PA/PAN) - 598: Papua New Guinea (PG/PNG) - 600: Paraguay (PY/PRY) - 604: Peru (PE/PER) - 608: Philippines (PH/PHL) - 612: Pitcairn (PN/PCN) - 616: Poland (PL/POL) - 620: Portugal (PT/PRT) - 624: Guinea-Bissau (GW/GNB) - 626: Timor-Leste (TL/TLS) - 630: Puerto Rico (PR/PRI) - 634: Qatar (QA/QAT) - 638: Réunion (RE/REU) - 642: Romania (RO/ROU) - 643: Russian Federation (RU/RUS) - 646: Rwanda (RW/RWA) - 652: Saint Barthélemy (BL/BLM) - 654: Saint Helena, Ascension and Tristan da Cunha (SH/SHN) - 659: Saint Kitts and Nevis (KN/KNA) - 660: Anguilla (AI/AIA) - 662: Saint Lucia (LC/LCA) - 663: Saint Martin (French part) (MF/MAF) - 666: Saint Pierre and Miquelon (PM/SPM) - 670: Saint Vincent and the Grenadines (VC/VCT) - 674: San Marino (SM/SMR) - 678: Sao Tome and Principe (ST/STP) - 682: Saudi Arabia (SA/SAU) - 686: Senegal (SN/SEN) - 688: Serbia (RS/SRB) - 690: Seychelles (SC/SYC) - 694: Sierra Leone (SL/SLE) - 702: Singapore (SG/SGP) - 703: Slovakia (SK/SVK) - 704: Viet Nam (VN/VNM) - 705: Slovenia (SI/SVN) - 706: Somalia (SO/SOM) - 710: South Africa (ZA/ZAF) - 716: Zimbabwe (ZW/ZWE) - 724: Spain (ES/ESP) - 728: South Sudan (SS/SSD) - 729: Sudan (SD/SDN) - 732: Western Sahara (EH/ESH) - 740: Suriname (SR/SUR) - 744: Svalbard and Jan Mayen (SJ/SJM) - 748: Swaziland (SZ/SWZ) - 752: Sweden (SE/SWE) - 756: Switzerland (CH/CHE) - 760: Syrian Arab Republic (SY/SYR) - 762: Tajikistan (TJ/TJK) - 764: Thailand (TH/THA) - 768: Togo (TG/TGO) - 772: Tokelau (TK/TKL) - 776: Tonga (TO/TON) - 780: Trinidad and Tobago (TT/TTO) - 784: United Arab Emirates (AE/ARE) - 788: Tunisia (TN/TUN) - 792: Turkey (TR/TUR) - 795: Turkmenistan (TM/TKM) - 796: Turks and Caicos Islands (TC/TCA) - 798: Tuvalu (TV/TUV) - 800: Uganda (UG/UGA) - 804: Ukraine (UA/UKR) - 807: Macedonia, the former Yugoslav Republic of (MK/MKD) - 818: Egypt (EG/EGY) - 826: United Kingdom (GB/GBR) - 831: Guernsey (GG/GGY) - 832: Jersey (JE/JEY) - 833: Isle of Man (IM/IMN) - 834: Tanzania, United Republic of (TZ/TZA) - 840: United States (US/USA) - 850: Virgin Islands, U.S. (VI/VIR) - 854: Burkina Faso (BF/BFA) - 858: Uruguay (UY/URY) - 860: Uzbekistan (UZ/UZB) - 862: Venezuela, Bolivarian Republic of (VE/VEN) - 876: Wallis and Futuna (WF/WLF) - 882: Samoa (WS/WSM) - 887: Yemen (YE/YEM) - 894: Zambia (ZM/ZMB)
	CountryId *int32 `json:"country_id,omitempty"`
	// A custom field for storing additional information
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A custom field for storing additional information
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A custom field for storing additional information
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A custom field for storing additional information
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The client's VAT (Value Added Tax) number, if applicable
	VatNumber *string `json:"vat_number,omitempty"`
	// A unique identification number for the client, such as a tax ID or business registration number
	IdNumber *string `json:"id_number,omitempty"`
	// A system-assigned unique number for the client, typically used for invoicing purposes
	Number *string `json:"number,omitempty"`
	// First line of the client's shipping address
	ShippingAddress1 *string `json:"shipping_address1,omitempty"`
	// Second line of the client's shipping address, if needed
	ShippingAddress2 *string `json:"shipping_address2,omitempty"`
	// The city of the client's shipping address
	ShippingCity *string `json:"shipping_city,omitempty"`
	// The state, province, or locality of the client's shipping address
	ShippingState *string `json:"shipping_state,omitempty"`
	// The postal code or ZIP code of the client's shipping address
	ShippingPostalCode *string `json:"shipping_postal_code,omitempty"`
	// The unique identifier of the country for the client's shipping address
	ShippingCountryId *float32 `json:"shipping_country_id,omitempty"`
	// A boolean value indicating whether the client has been deleted or not
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The outstanding balance the client owes
	Balance *float32 `json:"balance,omitempty"`
	// The total amount the client has paid to date
	PaidToDate *float32 `json:"paid_to_date,omitempty"`
	// The available credit balance for the client to use on future purchases
	CreditBalance *float32 `json:"credit_balance,omitempty"`
	// The timestamp of the client's last login
	LastLogin *float32 `json:"last_login,omitempty"`
	// The timestamp when the client was created
	CreatedAt *float32 `json:"created_at,omitempty"`
	// The timestamp when the client was last updated
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// The group settings assigned to the client
	GroupSettingsId *string `json:"group_settings_id,omitempty"`
	// The routing address id for e-invoicing for this client
	RoutingId *string `json:"routing_id,omitempty"`
	// Flag which defines if the client is exempt from taxes
	IsTaxExempt *bool `json:"is_tax_exempt,omitempty"`
	// Flag which defines if the client has a valid VAT number
	HasValidVatNumber *bool `json:"has_valid_vat_number,omitempty"`
	// Defines the payment balance the client has on file (pre payments / over payments / unapplied amounts)
	PaymentBalance *float32        `json:"payment_balance,omitempty"`
	Settings       *ClientSettings `json:"settings,omitempty"`
	Contacts       []ClientContact `json:"contacts,omitempty"`
}

// NewClient instantiates a new Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClient() *Client {
	this := Client{}
	return &this
}

// NewClientWithDefaults instantiates a new Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWithDefaults() *Client {
	this := Client{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Client) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Client) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Client) SetId(v string) {
	o.Id = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *Client) GetContacts() []ClientContact {
	if o == nil || IsNil(o.Contacts) {
		var ret []ClientContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetContactsOk() ([]ClientContact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *Client) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []ClientContact and assigns it to the Contacts field.
func (o *Client) SetContacts(v []ClientContact) {
	o.Contacts = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Client) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Client) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Client) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Client) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Client) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Client) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Client) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Client) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Client) SetName(v string) {
	o.Name = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *Client) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *Client) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *Client) SetWebsite(v string) {
	o.Website = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *Client) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *Client) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *Client) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetClientHash returns the ClientHash field value if set, zero value otherwise.
func (o *Client) GetClientHash() string {
	if o == nil || IsNil(o.ClientHash) {
		var ret string
		return ret
	}
	return *o.ClientHash
}

// GetClientHashOk returns a tuple with the ClientHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientHashOk() (*string, bool) {
	if o == nil || IsNil(o.ClientHash) {
		return nil, false
	}
	return o.ClientHash, true
}

// HasClientHash returns a boolean if a field has been set.
func (o *Client) HasClientHash() bool {
	if o != nil && !IsNil(o.ClientHash) {
		return true
	}

	return false
}

// SetClientHash gets a reference to the given string and assigns it to the ClientHash field.
func (o *Client) SetClientHash(v string) {
	o.ClientHash = &v
}

// GetIndustryId returns the IndustryId field value if set, zero value otherwise.
func (o *Client) GetIndustryId() float32 {
	if o == nil || IsNil(o.IndustryId) {
		var ret float32
		return ret
	}
	return *o.IndustryId
}

// GetIndustryIdOk returns a tuple with the IndustryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIndustryIdOk() (*float32, bool) {
	if o == nil || IsNil(o.IndustryId) {
		return nil, false
	}
	return o.IndustryId, true
}

// HasIndustryId returns a boolean if a field has been set.
func (o *Client) HasIndustryId() bool {
	if o != nil && !IsNil(o.IndustryId) {
		return true
	}

	return false
}

// SetIndustryId gets a reference to the given float32 and assigns it to the IndustryId field.
func (o *Client) SetIndustryId(v float32) {
	o.IndustryId = &v
}

// GetSizeId returns the SizeId field value if set, zero value otherwise.
func (o *Client) GetSizeId() float32 {
	if o == nil || IsNil(o.SizeId) {
		var ret float32
		return ret
	}
	return *o.SizeId
}

// GetSizeIdOk returns a tuple with the SizeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSizeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SizeId) {
		return nil, false
	}
	return o.SizeId, true
}

// HasSizeId returns a boolean if a field has been set.
func (o *Client) HasSizeId() bool {
	if o != nil && !IsNil(o.SizeId) {
		return true
	}

	return false
}

// SetSizeId gets a reference to the given float32 and assigns it to the SizeId field.
func (o *Client) SetSizeId(v float32) {
	o.SizeId = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *Client) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *Client) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *Client) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *Client) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *Client) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *Client) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Client) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Client) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Client) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Client) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Client) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Client) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Client) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Client) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Client) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Client) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Client) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Client) SetPhone(v string) {
	o.Phone = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *Client) GetCountryId() int32 {
	if o == nil || IsNil(o.CountryId) {
		var ret int32
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCountryIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *Client) HasCountryId() bool {
	if o != nil && !IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given int32 and assigns it to the CountryId field.
func (o *Client) SetCountryId(v int32) {
	o.CountryId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *Client) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Client) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *Client) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *Client) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Client) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *Client) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *Client) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Client) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *Client) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *Client) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Client) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *Client) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *Client) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *Client) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *Client) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise.
func (o *Client) GetIdNumber() string {
	if o == nil || IsNil(o.IdNumber) {
		var ret string
		return ret
	}
	return *o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIdNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IdNumber) {
		return nil, false
	}
	return o.IdNumber, true
}

// HasIdNumber returns a boolean if a field has been set.
func (o *Client) HasIdNumber() bool {
	if o != nil && !IsNil(o.IdNumber) {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given string and assigns it to the IdNumber field.
func (o *Client) SetIdNumber(v string) {
	o.IdNumber = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Client) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Client) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Client) SetNumber(v string) {
	o.Number = &v
}

// GetShippingAddress1 returns the ShippingAddress1 field value if set, zero value otherwise.
func (o *Client) GetShippingAddress1() string {
	if o == nil || IsNil(o.ShippingAddress1) {
		var ret string
		return ret
	}
	return *o.ShippingAddress1
}

// GetShippingAddress1Ok returns a tuple with the ShippingAddress1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetShippingAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.ShippingAddress1) {
		return nil, false
	}
	return o.ShippingAddress1, true
}

// HasShippingAddress1 returns a boolean if a field has been set.
func (o *Client) HasShippingAddress1() bool {
	if o != nil && !IsNil(o.ShippingAddress1) {
		return true
	}

	return false
}

// SetShippingAddress1 gets a reference to the given string and assigns it to the ShippingAddress1 field.
func (o *Client) SetShippingAddress1(v string) {
	o.ShippingAddress1 = &v
}

// GetShippingAddress2 returns the ShippingAddress2 field value if set, zero value otherwise.
func (o *Client) GetShippingAddress2() string {
	if o == nil || IsNil(o.ShippingAddress2) {
		var ret string
		return ret
	}
	return *o.ShippingAddress2
}

// GetShippingAddress2Ok returns a tuple with the ShippingAddress2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetShippingAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.ShippingAddress2) {
		return nil, false
	}
	return o.ShippingAddress2, true
}

// HasShippingAddress2 returns a boolean if a field has been set.
func (o *Client) HasShippingAddress2() bool {
	if o != nil && !IsNil(o.ShippingAddress2) {
		return true
	}

	return false
}

// SetShippingAddress2 gets a reference to the given string and assigns it to the ShippingAddress2 field.
func (o *Client) SetShippingAddress2(v string) {
	o.ShippingAddress2 = &v
}

// GetShippingCity returns the ShippingCity field value if set, zero value otherwise.
func (o *Client) GetShippingCity() string {
	if o == nil || IsNil(o.ShippingCity) {
		var ret string
		return ret
	}
	return *o.ShippingCity
}

// GetShippingCityOk returns a tuple with the ShippingCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetShippingCityOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingCity) {
		return nil, false
	}
	return o.ShippingCity, true
}

// HasShippingCity returns a boolean if a field has been set.
func (o *Client) HasShippingCity() bool {
	if o != nil && !IsNil(o.ShippingCity) {
		return true
	}

	return false
}

// SetShippingCity gets a reference to the given string and assigns it to the ShippingCity field.
func (o *Client) SetShippingCity(v string) {
	o.ShippingCity = &v
}

// GetShippingState returns the ShippingState field value if set, zero value otherwise.
func (o *Client) GetShippingState() string {
	if o == nil || IsNil(o.ShippingState) {
		var ret string
		return ret
	}
	return *o.ShippingState
}

// GetShippingStateOk returns a tuple with the ShippingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetShippingStateOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingState) {
		return nil, false
	}
	return o.ShippingState, true
}

// HasShippingState returns a boolean if a field has been set.
func (o *Client) HasShippingState() bool {
	if o != nil && !IsNil(o.ShippingState) {
		return true
	}

	return false
}

// SetShippingState gets a reference to the given string and assigns it to the ShippingState field.
func (o *Client) SetShippingState(v string) {
	o.ShippingState = &v
}

// GetShippingPostalCode returns the ShippingPostalCode field value if set, zero value otherwise.
func (o *Client) GetShippingPostalCode() string {
	if o == nil || IsNil(o.ShippingPostalCode) {
		var ret string
		return ret
	}
	return *o.ShippingPostalCode
}

// GetShippingPostalCodeOk returns a tuple with the ShippingPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetShippingPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingPostalCode) {
		return nil, false
	}
	return o.ShippingPostalCode, true
}

// HasShippingPostalCode returns a boolean if a field has been set.
func (o *Client) HasShippingPostalCode() bool {
	if o != nil && !IsNil(o.ShippingPostalCode) {
		return true
	}

	return false
}

// SetShippingPostalCode gets a reference to the given string and assigns it to the ShippingPostalCode field.
func (o *Client) SetShippingPostalCode(v string) {
	o.ShippingPostalCode = &v
}

// GetShippingCountryId returns the ShippingCountryId field value if set, zero value otherwise.
func (o *Client) GetShippingCountryId() float32 {
	if o == nil || IsNil(o.ShippingCountryId) {
		var ret float32
		return ret
	}
	return *o.ShippingCountryId
}

// GetShippingCountryIdOk returns a tuple with the ShippingCountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetShippingCountryIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ShippingCountryId) {
		return nil, false
	}
	return o.ShippingCountryId, true
}

// HasShippingCountryId returns a boolean if a field has been set.
func (o *Client) HasShippingCountryId() bool {
	if o != nil && !IsNil(o.ShippingCountryId) {
		return true
	}

	return false
}

// SetShippingCountryId gets a reference to the given float32 and assigns it to the ShippingCountryId field.
func (o *Client) SetShippingCountryId(v float32) {
	o.ShippingCountryId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Client) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Client) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Client) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Client) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Client) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *Client) SetBalance(v float32) {
	o.Balance = &v
}

// GetPaidToDate returns the PaidToDate field value if set, zero value otherwise.
func (o *Client) GetPaidToDate() float32 {
	if o == nil || IsNil(o.PaidToDate) {
		var ret float32
		return ret
	}
	return *o.PaidToDate
}

// GetPaidToDateOk returns a tuple with the PaidToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPaidToDateOk() (*float32, bool) {
	if o == nil || IsNil(o.PaidToDate) {
		return nil, false
	}
	return o.PaidToDate, true
}

// HasPaidToDate returns a boolean if a field has been set.
func (o *Client) HasPaidToDate() bool {
	if o != nil && !IsNil(o.PaidToDate) {
		return true
	}

	return false
}

// SetPaidToDate gets a reference to the given float32 and assigns it to the PaidToDate field.
func (o *Client) SetPaidToDate(v float32) {
	o.PaidToDate = &v
}

// GetCreditBalance returns the CreditBalance field value if set, zero value otherwise.
func (o *Client) GetCreditBalance() float32 {
	if o == nil || IsNil(o.CreditBalance) {
		var ret float32
		return ret
	}
	return *o.CreditBalance
}

// GetCreditBalanceOk returns a tuple with the CreditBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCreditBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditBalance) {
		return nil, false
	}
	return o.CreditBalance, true
}

// HasCreditBalance returns a boolean if a field has been set.
func (o *Client) HasCreditBalance() bool {
	if o != nil && !IsNil(o.CreditBalance) {
		return true
	}

	return false
}

// SetCreditBalance gets a reference to the given float32 and assigns it to the CreditBalance field.
func (o *Client) SetCreditBalance(v float32) {
	o.CreditBalance = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *Client) GetLastLogin() float32 {
	if o == nil || IsNil(o.LastLogin) {
		var ret float32
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetLastLoginOk() (*float32, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *Client) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given float32 and assigns it to the LastLogin field.
func (o *Client) SetLastLogin(v float32) {
	o.LastLogin = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Client) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Client) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *Client) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Client) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Client) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Client) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetGroupSettingsId returns the GroupSettingsId field value if set, zero value otherwise.
func (o *Client) GetGroupSettingsId() string {
	if o == nil || IsNil(o.GroupSettingsId) {
		var ret string
		return ret
	}
	return *o.GroupSettingsId
}

// GetGroupSettingsIdOk returns a tuple with the GroupSettingsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetGroupSettingsIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupSettingsId) {
		return nil, false
	}
	return o.GroupSettingsId, true
}

// HasGroupSettingsId returns a boolean if a field has been set.
func (o *Client) HasGroupSettingsId() bool {
	if o != nil && !IsNil(o.GroupSettingsId) {
		return true
	}

	return false
}

// SetGroupSettingsId gets a reference to the given string and assigns it to the GroupSettingsId field.
func (o *Client) SetGroupSettingsId(v string) {
	o.GroupSettingsId = &v
}

// GetRoutingId returns the RoutingId field value if set, zero value otherwise.
func (o *Client) GetRoutingId() string {
	if o == nil || IsNil(o.RoutingId) {
		var ret string
		return ret
	}
	return *o.RoutingId
}

// GetRoutingIdOk returns a tuple with the RoutingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRoutingIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingId) {
		return nil, false
	}
	return o.RoutingId, true
}

// HasRoutingId returns a boolean if a field has been set.
func (o *Client) HasRoutingId() bool {
	if o != nil && !IsNil(o.RoutingId) {
		return true
	}

	return false
}

// SetRoutingId gets a reference to the given string and assigns it to the RoutingId field.
func (o *Client) SetRoutingId(v string) {
	o.RoutingId = &v
}

// GetIsTaxExempt returns the IsTaxExempt field value if set, zero value otherwise.
func (o *Client) GetIsTaxExempt() bool {
	if o == nil || IsNil(o.IsTaxExempt) {
		var ret bool
		return ret
	}
	return *o.IsTaxExempt
}

// GetIsTaxExemptOk returns a tuple with the IsTaxExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIsTaxExemptOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTaxExempt) {
		return nil, false
	}
	return o.IsTaxExempt, true
}

// HasIsTaxExempt returns a boolean if a field has been set.
func (o *Client) HasIsTaxExempt() bool {
	if o != nil && !IsNil(o.IsTaxExempt) {
		return true
	}

	return false
}

// SetIsTaxExempt gets a reference to the given bool and assigns it to the IsTaxExempt field.
func (o *Client) SetIsTaxExempt(v bool) {
	o.IsTaxExempt = &v
}

// GetHasValidVatNumber returns the HasValidVatNumber field value if set, zero value otherwise.
func (o *Client) GetHasValidVatNumber() bool {
	if o == nil || IsNil(o.HasValidVatNumber) {
		var ret bool
		return ret
	}
	return *o.HasValidVatNumber
}

// GetHasValidVatNumberOk returns a tuple with the HasValidVatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetHasValidVatNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.HasValidVatNumber) {
		return nil, false
	}
	return o.HasValidVatNumber, true
}

// HasHasValidVatNumber returns a boolean if a field has been set.
func (o *Client) HasHasValidVatNumber() bool {
	if o != nil && !IsNil(o.HasValidVatNumber) {
		return true
	}

	return false
}

// SetHasValidVatNumber gets a reference to the given bool and assigns it to the HasValidVatNumber field.
func (o *Client) SetHasValidVatNumber(v bool) {
	o.HasValidVatNumber = &v
}

// GetPaymentBalance returns the PaymentBalance field value if set, zero value otherwise.
func (o *Client) GetPaymentBalance() float32 {
	if o == nil || IsNil(o.PaymentBalance) {
		var ret float32
		return ret
	}
	return *o.PaymentBalance
}

// GetPaymentBalanceOk returns a tuple with the PaymentBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPaymentBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.PaymentBalance) {
		return nil, false
	}
	return o.PaymentBalance, true
}

// HasPaymentBalance returns a boolean if a field has been set.
func (o *Client) HasPaymentBalance() bool {
	if o != nil && !IsNil(o.PaymentBalance) {
		return true
	}

	return false
}

// SetPaymentBalance gets a reference to the given float32 and assigns it to the PaymentBalance field.
func (o *Client) SetPaymentBalance(v float32) {
	o.PaymentBalance = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Client) GetSettings() ClientSettings {
	if o == nil || IsNil(o.Settings) {
		var ret ClientSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSettingsOk() (*ClientSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Client) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given ClientSettings and assigns it to the Settings field.
func (o *Client) SetSettings(v ClientSettings) {
	o.Settings = &v
}

func (o Client) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Client) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.PrivateNotes) {
		toSerialize["private_notes"] = o.PrivateNotes
	}
	if !IsNil(o.ClientHash) {
		toSerialize["client_hash"] = o.ClientHash
	}
	if !IsNil(o.IndustryId) {
		toSerialize["industry_id"] = o.IndustryId
	}
	if !IsNil(o.SizeId) {
		toSerialize["size_id"] = o.SizeId
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
	if !IsNil(o.CountryId) {
		toSerialize["country_id"] = o.CountryId
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
	if !IsNil(o.VatNumber) {
		toSerialize["vat_number"] = o.VatNumber
	}
	if !IsNil(o.IdNumber) {
		toSerialize["id_number"] = o.IdNumber
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.ShippingAddress1) {
		toSerialize["shipping_address1"] = o.ShippingAddress1
	}
	if !IsNil(o.ShippingAddress2) {
		toSerialize["shipping_address2"] = o.ShippingAddress2
	}
	if !IsNil(o.ShippingCity) {
		toSerialize["shipping_city"] = o.ShippingCity
	}
	if !IsNil(o.ShippingState) {
		toSerialize["shipping_state"] = o.ShippingState
	}
	if !IsNil(o.ShippingPostalCode) {
		toSerialize["shipping_postal_code"] = o.ShippingPostalCode
	}
	if !IsNil(o.ShippingCountryId) {
		toSerialize["shipping_country_id"] = o.ShippingCountryId
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.PaidToDate) {
		toSerialize["paid_to_date"] = o.PaidToDate
	}
	if !IsNil(o.CreditBalance) {
		toSerialize["credit_balance"] = o.CreditBalance
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.GroupSettingsId) {
		toSerialize["group_settings_id"] = o.GroupSettingsId
	}
	if !IsNil(o.RoutingId) {
		toSerialize["routing_id"] = o.RoutingId
	}
	if !IsNil(o.IsTaxExempt) {
		toSerialize["is_tax_exempt"] = o.IsTaxExempt
	}
	if !IsNil(o.HasValidVatNumber) {
		toSerialize["has_valid_vat_number"] = o.HasValidVatNumber
	}
	if !IsNil(o.PaymentBalance) {
		toSerialize["payment_balance"] = o.PaymentBalance
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableClient struct {
	value *Client
	isSet bool
}

func (v NullableClient) Get() *Client {
	return v.value
}

func (v *NullableClient) Set(val *Client) {
	v.value = val
	v.isSet = true
}

func (v NullableClient) IsSet() bool {
	return v.isSet
}

func (v *NullableClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClient(val *Client) *NullableClient {
	return &NullableClient{value: val, isSet: true}
}

func (v NullableClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
