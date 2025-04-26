# ClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the client | [optional] [readonly] 
**Contacts** | [**[]ClientContactRequest**](ClientContactRequest.md) | A list of contacts associated with the client | 
**Name** | Pointer to **string** | The name of the client company or organization | [optional] 
**Website** | Pointer to **string** | The website URL of the client company or organization | [optional] 
**PrivateNotes** | Pointer to **string** | Notes that are only visible to the user who created the client | [optional] 
**IndustryId** | Pointer to **float32** | The unique identifier of the industry the client operates in | [optional] 
**SizeId** | Pointer to **float32** | The unique identifier for the size category of the client company or organization | [optional] 
**Address1** | Pointer to **string** | First line of the client&#39;s address | [optional] 
**Address2** | Pointer to **string** | Second line of the client&#39;s address, if needed | [optional] 
**City** | Pointer to **string** | The city the client is located in | [optional] 
**State** | Pointer to **string** | The state, province, or locality the client is located in | [optional] 
**PostalCode** | Pointer to **string** | The postal code or ZIP code of the client | [optional] 
**Phone** | Pointer to **string** | The client&#39;s phone number | [optional] 
**CountryId** | **int32** | Available country types: - 4: Afghanistan (AF/AFG) - 8: Albania (AL/ALB) - 10: Antarctica (AQ/ATA) - 12: Algeria (DZ/DZA) - 16: American Samoa (AS/ASM) - 20: Andorra (AD/AND) - 24: Angola (AO/AGO) - 28: Antigua and Barbuda (AG/ATG) - 31: Azerbaijan (AZ/AZE) - 32: Argentina (AR/ARG) - 36: Australia (AU/AUS) - 40: Austria (AT/AUT) - 44: Bahamas (BS/BHS) - 48: Bahrain (BH/BHR) - 50: Bangladesh (BD/BGD) - 51: Armenia (AM/ARM) - 52: Barbados (BB/BRB) - 56: Belgium (BE/BEL) - 60: Bermuda (BM/BMU) - 64: Bhutan (BT/BTN) - 68: Bolivia, Plurinational State of (BO/BOL) - 70: Bosnia and Herzegovina (BA/BIH) - 72: Botswana (BW/BWA) - 74: Bouvet Island (BV/BVT) - 76: Brazil (BR/BRA) - 84: Belize (BZ/BLZ) - 86: British Indian Ocean Territory (IO/IOT) - 90: Solomon Islands (SB/SLB) - 92: Virgin Islands, British (VG/VGB) - 96: Brunei Darussalam (BN/BRN) - 100: Bulgaria (BG/BGR) - 104: Myanmar (MM/MMR) - 108: Burundi (BI/BDI) - 112: Belarus (BY/BLR) - 116: Cambodia (KH/KHM) - 120: Cameroon (CM/CMR) - 124: Canada (CA/CAN) - 132: Cape Verde (CV/CPV) - 136: Cayman Islands (KY/CYM) - 140: Central African Republic (CF/CAF) - 144: Sri Lanka (LK/LKA) - 148: Chad (TD/TCD) - 152: Chile (CL/CHL) - 156: China (CN/CHN) - 158: Taiwan (TW/TWN) - 162: Christmas Island (CX/CXR) - 166: Cocos (Keeling) Islands (CC/CCK) - 170: Colombia (CO/COL) - 174: Comoros (KM/COM) - 175: Mayotte (YT/MYT) - 178: Congo (CG/COG) - 180: Congo, the Democratic Republic of the (CD/COD) - 184: Cook Islands (CK/COK) - 188: Costa Rica (CR/CRI) - 191: Croatia (HR/HRV) - 192: Cuba (CU/CUB) - 196: Cyprus (CY/CYP) - 203: Czech Republic (CZ/CZE) - 204: Benin (BJ/BEN) - 208: Denmark (DK/DNK) - 212: Dominica (DM/DMA) - 214: Dominican Republic (DO/DOM) - 218: Ecuador (EC/ECU) - 222: El Salvador (SV/SLV) - 226: Equatorial Guinea (GQ/GNQ) - 231: Ethiopia (ET/ETH) - 232: Eritrea (ER/ERI) - 233: Estonia (EE/EST) - 234: Faroe Islands (FO/FRO) - 238: Falkland Islands (Malvinas) (FK/FLK) - 239: South Georgia and the South Sandwich Islands (GS/SGS) - 242: Fiji (FJ/FJI) - 246: Finland (FI/FIN) - 248: Åland Islands (AX/ALA) - 250: France (FR/FRA) - 254: French Guiana (GF/GUF) - 258: French Polynesia (PF/PYF) - 260: French Southern Territories (TF/ATF) - 262: Djibouti (DJ/DJI) - 266: Gabon (GA/GAB) - 268: Georgia (GE/GEO) - 270: Gambia (GM/GMB) - 275: Palestine (PS/PSE) - 276: Germany (DE/DEU) - 288: Ghana (GH/GHA) - 292: Gibraltar (GI/GIB) - 296: Kiribati (KI/KIR) - 300: Greece (GR/GRC) - 304: Greenland (GL/GRL) - 308: Grenada (GD/GRD) - 312: Guadeloupe (GP/GLP) - 316: Guam (GU/GUM) - 320: Guatemala (GT/GTM) - 324: Guinea (GN/GIN) - 328: Guyana (GY/GUY) - 332: Haiti (HT/HTI) - 334: Heard Island and McDonald Islands (HM/HMD) - 336: Holy See (Vatican City State) (VA/VAT) - 340: Honduras (HN/HND) - 344: Hong Kong (HK/HKG) - 348: Hungary (HU/HUN) - 352: Iceland (IS/ISL) - 356: India (IN/IND) - 360: Indonesia (ID/IDN) - 364: Iran, Islamic Republic of (IR/IRN) - 368: Iraq (IQ/IRQ) - 372: Ireland (IE/IRL) - 376: Israel (IL/ISR) - 380: Italy (IT/ITA) - 384: Côte d&#39;Ivoire (CI/CIV) - 388: Jamaica (JM/JAM) - 392: Japan (JP/JPN) - 398: Kazakhstan (KZ/KAZ) - 400: Jordan (JO/JOR) - 404: Kenya (KE/KEN) - 408: Korea, Democratic People&#39;s Republic of (KP/PRK) - 410: Korea, Republic of (KR/KOR) - 414: Kuwait (KW/KWT) - 417: Kyrgyzstan (KG/KGZ) - 418: Lao People&#39;s Democratic Republic (LA/LAO) - 422: Lebanon (LB/LBN) - 426: Lesotho (LS/LSO) - 428: Latvia (LV/LVA) - 430: Liberia (LR/LBR) - 434: Libya (LY/LBY) - 438: Liechtenstein (LI/LIE) - 440: Lithuania (LT/LTU) - 442: Luxembourg (LU/LUX) - 446: Macao (MO/MAC) - 450: Madagascar (MG/MDG) - 454: Malawi (MW/MWI) - 458: Malaysia (MY/MYS) - 462: Maldives (MV/MDV) - 466: Mali (ML/MLI) - 470: Malta (MT/MLT) - 474: Martinique (MQ/MTQ) - 478: Mauritania (MR/MRT) - 480: Mauritius (MU/MUS) - 484: Mexico (MX/MEX) - 492: Monaco (MC/MCO) - 496: Mongolia (MN/MNG) - 498: Moldova, Republic of (MD/MDA) - 499: Montenegro (ME/MNE) - 500: Montserrat (MS/MSR) - 504: Morocco (MA/MAR) - 508: Mozambique (MZ/MOZ) - 512: Oman (OM/OMN) - 516: Namibia (NA/NAM) - 520: Nauru (NR/NRU) - 524: Nepal (NP/NPL) - 528: Netherlands (NL/NLD) - 531: Curaçao (CW/CUW) - 533: Aruba (AW/ABW) - 534: Sint Maarten (Dutch part) (SX/SXM) - 535: Bonaire, Sint Eustatius and Saba (BQ/BES) - 540: New Caledonia (NC/NCL) - 548: Vanuatu (VU/VUT) - 554: New Zealand (NZ/NZL) - 558: Nicaragua (NI/NIC) - 562: Niger (NE/NER) - 566: Nigeria (NG/NGA) - 570: Niue (NU/NIU) - 574: Norfolk Island (NF/NFK) - 578: Norway (NO/NOR) - 580: Northern Mariana Islands (MP/MNP) - 581: United States Minor Outlying Islands (UM/UMI) - 583: Micronesia, Federated States of (FM/FSM) - 584: Marshall Islands (MH/MHL) - 585: Palau (PW/PLW) - 586: Pakistan (PK/PAK) - 591: Panama (PA/PAN) - 598: Papua New Guinea (PG/PNG) - 600: Paraguay (PY/PRY) - 604: Peru (PE/PER) - 608: Philippines (PH/PHL) - 612: Pitcairn (PN/PCN) - 616: Poland (PL/POL) - 620: Portugal (PT/PRT) - 624: Guinea-Bissau (GW/GNB) - 626: Timor-Leste (TL/TLS) - 630: Puerto Rico (PR/PRI) - 634: Qatar (QA/QAT) - 638: Réunion (RE/REU) - 642: Romania (RO/ROU) - 643: Russian Federation (RU/RUS) - 646: Rwanda (RW/RWA) - 652: Saint Barthélemy (BL/BLM) - 654: Saint Helena, Ascension and Tristan da Cunha (SH/SHN) - 659: Saint Kitts and Nevis (KN/KNA) - 660: Anguilla (AI/AIA) - 662: Saint Lucia (LC/LCA) - 663: Saint Martin (French part) (MF/MAF) - 666: Saint Pierre and Miquelon (PM/SPM) - 670: Saint Vincent and the Grenadines (VC/VCT) - 674: San Marino (SM/SMR) - 678: Sao Tome and Principe (ST/STP) - 682: Saudi Arabia (SA/SAU) - 686: Senegal (SN/SEN) - 688: Serbia (RS/SRB) - 690: Seychelles (SC/SYC) - 694: Sierra Leone (SL/SLE) - 702: Singapore (SG/SGP) - 703: Slovakia (SK/SVK) - 704: Viet Nam (VN/VNM) - 705: Slovenia (SI/SVN) - 706: Somalia (SO/SOM) - 710: South Africa (ZA/ZAF) - 716: Zimbabwe (ZW/ZWE) - 724: Spain (ES/ESP) - 728: South Sudan (SS/SSD) - 729: Sudan (SD/SDN) - 732: Western Sahara (EH/ESH) - 740: Suriname (SR/SUR) - 744: Svalbard and Jan Mayen (SJ/SJM) - 748: Swaziland (SZ/SWZ) - 752: Sweden (SE/SWE) - 756: Switzerland (CH/CHE) - 760: Syrian Arab Republic (SY/SYR) - 762: Tajikistan (TJ/TJK) - 764: Thailand (TH/THA) - 768: Togo (TG/TGO) - 772: Tokelau (TK/TKL) - 776: Tonga (TO/TON) - 780: Trinidad and Tobago (TT/TTO) - 784: United Arab Emirates (AE/ARE) - 788: Tunisia (TN/TUN) - 792: Turkey (TR/TUR) - 795: Turkmenistan (TM/TKM) - 796: Turks and Caicos Islands (TC/TCA) - 798: Tuvalu (TV/TUV) - 800: Uganda (UG/UGA) - 804: Ukraine (UA/UKR) - 807: Macedonia, the former Yugoslav Republic of (MK/MKD) - 818: Egypt (EG/EGY) - 826: United Kingdom (GB/GBR) - 831: Guernsey (GG/GGY) - 832: Jersey (JE/JEY) - 833: Isle of Man (IM/IMN) - 834: Tanzania, United Republic of (TZ/TZA) - 840: United States (US/USA) - 850: Virgin Islands, U.S. (VI/VIR) - 854: Burkina Faso (BF/BFA) - 858: Uruguay (UY/URY) - 860: Uzbekistan (UZ/UZB) - 862: Venezuela, Bolivarian Republic of (VE/VEN) - 876: Wallis and Futuna (WF/WLF) - 882: Samoa (WS/WSM) - 887: Yemen (YE/YEM) - 894: Zambia (ZM/ZMB)  | 
**CustomValue1** | Pointer to **string** | A custom field for storing additional information | [optional] 
**CustomValue2** | Pointer to **string** | A custom field for storing additional information | [optional] 
**CustomValue3** | Pointer to **string** | A custom field for storing additional information | [optional] 
**CustomValue4** | Pointer to **string** | A custom field for storing additional information | [optional] 
**VatNumber** | Pointer to **string** | The client&#39;s VAT (Value Added Tax) number, if applicable | [optional] 
**IdNumber** | Pointer to **string** | A unique identification number for the client, such as a tax ID or business registration number | [optional] 
**Number** | Pointer to **string** | A system-assigned unique number for the client, typically used for invoicing purposes | [optional] 
**ShippingAddress1** | Pointer to **string** | First line of the client&#39;s shipping address | [optional] 
**ShippingAddress2** | Pointer to **string** | Second line of the client&#39;s shipping address, if needed | [optional] 
**ShippingCity** | Pointer to **string** | The city of the client&#39;s shipping address | [optional] 
**ShippingState** | Pointer to **string** | The state, province, or locality of the client&#39;s shipping address | [optional] 
**ShippingPostalCode** | Pointer to **string** | The postal code or ZIP code of the client&#39;s shipping address | [optional] 
**ShippingCountryId** | Pointer to **float32** | The unique identifier of the client&#39;s shipping country expressed by the countries ISO number.  Optionally, instead of the shipping_country_id you can pass either the iso_3166_2 or iso_3166_3 country code into the shipping_country_code property.  | [optional] 
**IsDeleted** | Pointer to **bool** | A boolean value indicating whether the client has been deleted or not | [optional] [readonly] 
**GroupSettingsId** | Pointer to **string** | The group settings assigned to the client | [optional] 
**RoutingId** | Pointer to **string** | The routing address id for e-invoicing for this client | [optional] 
**IsTaxExempt** | Pointer to **bool** | Flag which defines if the client is exempt from taxes | [optional] 
**HasValidVatNumber** | Pointer to **bool** | Flag which defines if the client has a valid VAT number | [optional] 
**Classification** | Pointer to **string** | The classification of the client | [optional] 
**Settings** | Pointer to [**ClientSettings**](ClientSettings.md) |  | [optional] 

## Methods

### NewClientRequest

`func NewClientRequest(contacts []ClientContactRequest, countryId int32, ) *ClientRequest`

NewClientRequest instantiates a new ClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRequestWithDefaults

`func NewClientRequestWithDefaults() *ClientRequest`

NewClientRequestWithDefaults instantiates a new ClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContacts

`func (o *ClientRequest) GetContacts() []ClientContactRequest`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ClientRequest) GetContactsOk() (*[]ClientContactRequest, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ClientRequest) SetContacts(v []ClientContactRequest)`

SetContacts sets Contacts field to given value.


### GetName

`func (o *ClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWebsite

`func (o *ClientRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ClientRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ClientRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ClientRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *ClientRequest) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *ClientRequest) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *ClientRequest) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *ClientRequest) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetIndustryId

`func (o *ClientRequest) GetIndustryId() float32`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *ClientRequest) GetIndustryIdOk() (*float32, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *ClientRequest) SetIndustryId(v float32)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *ClientRequest) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetSizeId

`func (o *ClientRequest) GetSizeId() float32`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *ClientRequest) GetSizeIdOk() (*float32, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *ClientRequest) SetSizeId(v float32)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *ClientRequest) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### GetAddress1

`func (o *ClientRequest) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ClientRequest) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ClientRequest) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ClientRequest) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *ClientRequest) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ClientRequest) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ClientRequest) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ClientRequest) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *ClientRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ClientRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ClientRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ClientRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *ClientRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClientRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClientRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClientRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *ClientRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ClientRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ClientRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ClientRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPhone

`func (o *ClientRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ClientRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ClientRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ClientRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCountryId

`func (o *ClientRequest) GetCountryId() int32`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ClientRequest) GetCountryIdOk() (*int32, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ClientRequest) SetCountryId(v int32)`

SetCountryId sets CountryId field to given value.


### GetCustomValue1

`func (o *ClientRequest) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *ClientRequest) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *ClientRequest) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *ClientRequest) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *ClientRequest) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *ClientRequest) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *ClientRequest) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *ClientRequest) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *ClientRequest) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *ClientRequest) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *ClientRequest) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *ClientRequest) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *ClientRequest) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *ClientRequest) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *ClientRequest) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *ClientRequest) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetVatNumber

`func (o *ClientRequest) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *ClientRequest) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *ClientRequest) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *ClientRequest) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetIdNumber

`func (o *ClientRequest) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *ClientRequest) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *ClientRequest) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.

### HasIdNumber

`func (o *ClientRequest) HasIdNumber() bool`

HasIdNumber returns a boolean if a field has been set.

### GetNumber

`func (o *ClientRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ClientRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ClientRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ClientRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetShippingAddress1

`func (o *ClientRequest) GetShippingAddress1() string`

GetShippingAddress1 returns the ShippingAddress1 field if non-nil, zero value otherwise.

### GetShippingAddress1Ok

`func (o *ClientRequest) GetShippingAddress1Ok() (*string, bool)`

GetShippingAddress1Ok returns a tuple with the ShippingAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress1

`func (o *ClientRequest) SetShippingAddress1(v string)`

SetShippingAddress1 sets ShippingAddress1 field to given value.

### HasShippingAddress1

`func (o *ClientRequest) HasShippingAddress1() bool`

HasShippingAddress1 returns a boolean if a field has been set.

### GetShippingAddress2

`func (o *ClientRequest) GetShippingAddress2() string`

GetShippingAddress2 returns the ShippingAddress2 field if non-nil, zero value otherwise.

### GetShippingAddress2Ok

`func (o *ClientRequest) GetShippingAddress2Ok() (*string, bool)`

GetShippingAddress2Ok returns a tuple with the ShippingAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress2

`func (o *ClientRequest) SetShippingAddress2(v string)`

SetShippingAddress2 sets ShippingAddress2 field to given value.

### HasShippingAddress2

`func (o *ClientRequest) HasShippingAddress2() bool`

HasShippingAddress2 returns a boolean if a field has been set.

### GetShippingCity

`func (o *ClientRequest) GetShippingCity() string`

GetShippingCity returns the ShippingCity field if non-nil, zero value otherwise.

### GetShippingCityOk

`func (o *ClientRequest) GetShippingCityOk() (*string, bool)`

GetShippingCityOk returns a tuple with the ShippingCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCity

`func (o *ClientRequest) SetShippingCity(v string)`

SetShippingCity sets ShippingCity field to given value.

### HasShippingCity

`func (o *ClientRequest) HasShippingCity() bool`

HasShippingCity returns a boolean if a field has been set.

### GetShippingState

`func (o *ClientRequest) GetShippingState() string`

GetShippingState returns the ShippingState field if non-nil, zero value otherwise.

### GetShippingStateOk

`func (o *ClientRequest) GetShippingStateOk() (*string, bool)`

GetShippingStateOk returns a tuple with the ShippingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingState

`func (o *ClientRequest) SetShippingState(v string)`

SetShippingState sets ShippingState field to given value.

### HasShippingState

`func (o *ClientRequest) HasShippingState() bool`

HasShippingState returns a boolean if a field has been set.

### GetShippingPostalCode

`func (o *ClientRequest) GetShippingPostalCode() string`

GetShippingPostalCode returns the ShippingPostalCode field if non-nil, zero value otherwise.

### GetShippingPostalCodeOk

`func (o *ClientRequest) GetShippingPostalCodeOk() (*string, bool)`

GetShippingPostalCodeOk returns a tuple with the ShippingPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPostalCode

`func (o *ClientRequest) SetShippingPostalCode(v string)`

SetShippingPostalCode sets ShippingPostalCode field to given value.

### HasShippingPostalCode

`func (o *ClientRequest) HasShippingPostalCode() bool`

HasShippingPostalCode returns a boolean if a field has been set.

### GetShippingCountryId

`func (o *ClientRequest) GetShippingCountryId() float32`

GetShippingCountryId returns the ShippingCountryId field if non-nil, zero value otherwise.

### GetShippingCountryIdOk

`func (o *ClientRequest) GetShippingCountryIdOk() (*float32, bool)`

GetShippingCountryIdOk returns a tuple with the ShippingCountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryId

`func (o *ClientRequest) SetShippingCountryId(v float32)`

SetShippingCountryId sets ShippingCountryId field to given value.

### HasShippingCountryId

`func (o *ClientRequest) HasShippingCountryId() bool`

HasShippingCountryId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *ClientRequest) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ClientRequest) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ClientRequest) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ClientRequest) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetGroupSettingsId

`func (o *ClientRequest) GetGroupSettingsId() string`

GetGroupSettingsId returns the GroupSettingsId field if non-nil, zero value otherwise.

### GetGroupSettingsIdOk

`func (o *ClientRequest) GetGroupSettingsIdOk() (*string, bool)`

GetGroupSettingsIdOk returns a tuple with the GroupSettingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSettingsId

`func (o *ClientRequest) SetGroupSettingsId(v string)`

SetGroupSettingsId sets GroupSettingsId field to given value.

### HasGroupSettingsId

`func (o *ClientRequest) HasGroupSettingsId() bool`

HasGroupSettingsId returns a boolean if a field has been set.

### GetRoutingId

`func (o *ClientRequest) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### GetRoutingIdOk

`func (o *ClientRequest) GetRoutingIdOk() (*string, bool)`

GetRoutingIdOk returns a tuple with the RoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingId

`func (o *ClientRequest) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### HasRoutingId

`func (o *ClientRequest) HasRoutingId() bool`

HasRoutingId returns a boolean if a field has been set.

### GetIsTaxExempt

`func (o *ClientRequest) GetIsTaxExempt() bool`

GetIsTaxExempt returns the IsTaxExempt field if non-nil, zero value otherwise.

### GetIsTaxExemptOk

`func (o *ClientRequest) GetIsTaxExemptOk() (*bool, bool)`

GetIsTaxExemptOk returns a tuple with the IsTaxExempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaxExempt

`func (o *ClientRequest) SetIsTaxExempt(v bool)`

SetIsTaxExempt sets IsTaxExempt field to given value.

### HasIsTaxExempt

`func (o *ClientRequest) HasIsTaxExempt() bool`

HasIsTaxExempt returns a boolean if a field has been set.

### GetHasValidVatNumber

`func (o *ClientRequest) GetHasValidVatNumber() bool`

GetHasValidVatNumber returns the HasValidVatNumber field if non-nil, zero value otherwise.

### GetHasValidVatNumberOk

`func (o *ClientRequest) GetHasValidVatNumberOk() (*bool, bool)`

GetHasValidVatNumberOk returns a tuple with the HasValidVatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValidVatNumber

`func (o *ClientRequest) SetHasValidVatNumber(v bool)`

SetHasValidVatNumber sets HasValidVatNumber field to given value.

### HasHasValidVatNumber

`func (o *ClientRequest) HasHasValidVatNumber() bool`

HasHasValidVatNumber returns a boolean if a field has been set.

### GetClassification

`func (o *ClientRequest) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ClientRequest) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ClientRequest) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ClientRequest) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetSettings

`func (o *ClientRequest) GetSettings() ClientSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ClientRequest) GetSettingsOk() (*ClientSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ClientRequest) SetSettings(v ClientSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ClientRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


