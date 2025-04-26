# Go API client for openapi

---
  ![Invoice Ninja](https://invoicing.co/images/new_logo.png)
  ## Introduction
  Welcome to the Invoice Ninja API documentation, your comprehensive guide to integrating Invoice Ninja's powerful features into your applications. Whether you're building a custom client, automating workflows, or integrating with other systems, our API provides the tools you need to streamline your invoicing and billing processes.
  ### What is Invoice Ninja?
  Invoice Ninja is a robust source-available platform designed to simplify invoicing, billing, and payment management for freelancers, small businesses, and enterprises alike. With a user-friendly interface, customizable templates, and a suite of powerful features, Invoice Ninja empowers businesses to create professional invoices, track expenses, manage clients, and get paid faster.
  ### Why use the Invoice Ninja API?
  The Invoice Ninja API allows developers to extend the functionality of Invoice Ninja by programmatically accessing and manipulating data within their Invoice Ninja accounts. With the API, you can automate repetitive tasks, integrate with third-party services, and build custom solutions tailored to your specific business needs.
  ### Getting Started
  To get started with the Invoice Ninja API, you'll need an active Invoice Ninja account (or your own self hosted installation) and API credentials. If you haven't already done so, sign up for an account at Invoice Ninja and generate your API keys from the settings section.

  Once you have your API credentials, you can start exploring the API endpoints, authentication methods, request and response formats, and more using the documentation provided here.
  ### Explore the Documentation  
  This documentation is organized into sections to help you navigate and understand the various aspects of the Invoice Ninja API:

  - Authentication: Learn how to authenticate your requests to the API using API tokens.
  - Endpoints: Explore the available API endpoints for managing invoices, clients, payments, expenses, and more.
  - Request and Response Formats: Understand the structure of API requests and responses, including parameters, headers, and payloads.
  - Error Handling: Learn about error codes, status messages, and best practices for handling errors gracefully.
  - Code Examples: Find code examples and tutorials to help you get started with integrating the Invoice Ninja API into your applications.      
  ### Need Help?     

  If you have any questions, encounter any issues, or need assistance with using the Invoice Ninja API, don't hesitate to reach out to our support team or join our community forums. We're here to help you succeed with Invoice Ninja and make the most of our API.    

  Let's start building together!
  ### Endpoints
  
  <div style=\"background-color: #2D394E; color: #fff padding: 20px; border-radius: 5px; border: 4px solid #212A3B; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);\">
      <p style=\"padding:10px; color: #DBDBDB;\"\">Production: https://invoicing.co</p>
      <p style=\"padding:10px; color: #DBDBDB;\">Demo: https://demo.invoiceninja.com</p>
  </div>

  ### Client Libraries

  PHP SDK can be found [here](https://github.com/invoiceninja/sdk-php)
  ### Authentication:

  Invoice Ninja uses API tokens to authenticate requests. You can view and manage your API keys in Settings > Account Management > Integrations > API tokens

  API requests must be made over HTTPS. Calls made to HTTP will fail.
  ### Errors:

  Invoice Ninja uses standard HTTP response codes to indicate the success or failure of a request. below is a table of standard status codes and responses

  | Status Code | Explanation                                                                 |
  |-------------|-----------------------------------------------------------------------------|
  | 200         | OK: The request has succeeded. The information returned with the response is dependent on the method used in the request. |
  | 301         | Moved Permanently: This and all future requests should be directed to the given URI. |
  | 303         | See Other: The response to the request can be found under another URI using the GET method. |
  | 400         | Bad Request: The server cannot or will not process the request due to an apparent client error. |
  | 401         | Unauthorized: Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided. |
  | 403         | Forbidden: The request was valid, but the server is refusing action. |
  | 404         | Not Found: The requested resource could not be found but may be available in the future. |
  | 405         | Method Not Allowed: A request method is not supported for the requested resource. |
  | 409         | Conflict: Indicates that the request could not be processed because of conflict in the request. |
  | 422         | Unprocessable Entity: The request was well-formed but was unable to be followed due to semantic errors. |
  | 429         | Too Many Requests: The user has sent too many requests in a given amount of time (\"rate limiting\"). |
  | 500         | Internal Server Error: A generic error message, given when an unexpected condition was encountered and no more specific message is suitable. |
  ### Pagination

  When using index routes to retrieve lists of data, by default we limit the number of records returned to 20. You can using standard pagination to paginate results, ie: ?per_page=50


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 5.11.48
- Package version: 1.0.0
- Generator version: 7.13.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.invoiceninja.com](https://www.invoiceninja.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://demo.invoiceninja.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActivitiesAPI* | [**GetActivities**](docs/ActivitiesAPI.md#getactivities) | **Get** /api/v1/activities | Returns a list of activities
*ActivitiesAPI* | [**GetActivityHistoricalEntityPdf**](docs/ActivitiesAPI.md#getactivityhistoricalentitypdf) | **Get** /api/v1/activities/download_entity/{activity_id} | Returns a PDF for the given activity
*AuthAPI* | [**PostLogin**](docs/AuthAPI.md#postlogin) | **Post** /api/v1/login | Login
*BankIntegrationsAPI* | [**BulkBankIntegrations**](docs/BankIntegrationsAPI.md#bulkbankintegrations) | **Post** /api/v1/bank_integrations/bulk | Performs bulk actions on an array of bank_integrations
*BankIntegrationsAPI* | [**DeleteBankIntegration**](docs/BankIntegrationsAPI.md#deletebankintegration) | **Delete** /api/v1/bank_integrations/{id} | Deletes a bank_integration
*BankIntegrationsAPI* | [**EditBankIntegration**](docs/BankIntegrationsAPI.md#editbankintegration) | **Get** /api/v1/bank_integrations/{id}/edit | Shows a bank_integration for editing
*BankIntegrationsAPI* | [**GetAccountTransactions**](docs/BankIntegrationsAPI.md#getaccounttransactions) | **Post** /api/v1/bank_integrations/get_transactions/account_id | Retrieve transactions for a account
*BankIntegrationsAPI* | [**GetBankIntegrations**](docs/BankIntegrationsAPI.md#getbankintegrations) | **Get** /api/v1/bank_integrations | Returns a list of Bank Integrations
*BankIntegrationsAPI* | [**GetBankIntegrationsCreate**](docs/BankIntegrationsAPI.md#getbankintegrationscreate) | **Get** /api/v1/bank_integrations/create | Gets a new blank bank_integration object
*BankIntegrationsAPI* | [**GetRefreshAccounts**](docs/BankIntegrationsAPI.md#getrefreshaccounts) | **Post** /api/v1/bank_integrations/refresh_accounts | Gets the list of accounts from the remote server
*BankIntegrationsAPI* | [**GetRemoveAccount**](docs/BankIntegrationsAPI.md#getremoveaccount) | **Post** /api/v1/bank_integrations/remove_account/account_id | Removes an account from the integration
*BankIntegrationsAPI* | [**ShowBankIntegration**](docs/BankIntegrationsAPI.md#showbankintegration) | **Get** /api/v1/bank_integrations/{id} | Shows a bank_integration
*BankIntegrationsAPI* | [**StoreBankIntegration**](docs/BankIntegrationsAPI.md#storebankintegration) | **Post** /api/v1/bank_integrations | Adds a bank_integration
*BankIntegrationsAPI* | [**UpdateBankIntegration**](docs/BankIntegrationsAPI.md#updatebankintegration) | **Put** /api/v1/bank_integrations/{id} | Updates a bank_integration
*BankTransactionRulesAPI* | [**BulkBankTransactionRules**](docs/BankTransactionRulesAPI.md#bulkbanktransactionrules) | **Post** /api/v1/bank_transation_rules/bulk | Performs bulk actions on an array of bank_transation rules
*BankTransactionRulesAPI* | [**DeleteBankTransactionRule**](docs/BankTransactionRulesAPI.md#deletebanktransactionrule) | **Delete** /api/v1/bank_transaction_rules/{id} | Deletes a bank_transaction rule
*BankTransactionRulesAPI* | [**EditBankTransactionRule**](docs/BankTransactionRulesAPI.md#editbanktransactionrule) | **Get** /api/v1/bank_transaction_rules/{id}/edit | Shows a bank_transaction for editing
*BankTransactionRulesAPI* | [**GetBankTransactionRules**](docs/BankTransactionRulesAPI.md#getbanktransactionrules) | **Get** /api/v1/bank_transaction_rules | Gets a list of bank_transaction_rules
*BankTransactionRulesAPI* | [**GetBankTransactionRulesCreate**](docs/BankTransactionRulesAPI.md#getbanktransactionrulescreate) | **Get** /api/v1/bank_transaction_rules/create | Gets a new blank bank_transaction rule object
*BankTransactionRulesAPI* | [**ShowBankTransactionRule**](docs/BankTransactionRulesAPI.md#showbanktransactionrule) | **Get** /api/v1/bank_transaction_rules/{id} | Shows a bank_transaction
*BankTransactionRulesAPI* | [**StoreBankTransactionRule**](docs/BankTransactionRulesAPI.md#storebanktransactionrule) | **Post** /api/v1/bank_transaction_rules | Adds a bank_transaction rule
*BankTransactionRulesAPI* | [**UpdateBankTransactionRule**](docs/BankTransactionRulesAPI.md#updatebanktransactionrule) | **Put** /api/v1/bank_transaction_rules/{id} | Updates a bank_transaction Rule
*BankTransactionsAPI* | [**BulkBankTransactions**](docs/BankTransactionsAPI.md#bulkbanktransactions) | **Post** /api/v1/bank_transations/bulk | Bulk actions
*BankTransactionsAPI* | [**DeleteBankTransaction**](docs/BankTransactionsAPI.md#deletebanktransaction) | **Delete** /api/v1/bank_transactions/{id} | Deletes a bank_transaction
*BankTransactionsAPI* | [**EditBankTransaction**](docs/BankTransactionsAPI.md#editbanktransaction) | **Get** /api/v1/bank_transactions/{id}/edit | Shows a bank_transaction for editing
*BankTransactionsAPI* | [**GetBankTransactions**](docs/BankTransactionsAPI.md#getbanktransactions) | **Get** /api/v1/bank_transactions | Gets a list of bank_transactions
*BankTransactionsAPI* | [**GetBankTransactionsCreate**](docs/BankTransactionsAPI.md#getbanktransactionscreate) | **Get** /api/v1/bank_transactions/create | Gets a new blank bank_transaction object
*BankTransactionsAPI* | [**MatchBankTransactions**](docs/BankTransactionsAPI.md#matchbanktransactions) | **Post** /api/v1/bank_transations/match | Match transactions
*BankTransactionsAPI* | [**ShowBankTransaction**](docs/BankTransactionsAPI.md#showbanktransaction) | **Get** /api/v1/bank_transactions/{id} | Shows a bank_transaction
*BankTransactionsAPI* | [**StoreBankTransaction**](docs/BankTransactionsAPI.md#storebanktransaction) | **Post** /api/v1/bank_transactions | Adds a bank_transaction
*BankTransactionsAPI* | [**UpdateBankTransaction**](docs/BankTransactionsAPI.md#updatebanktransaction) | **Put** /api/v1/bank_transactions/{id} | Updates a bank_transaction
*ChartsAPI* | [**GetChartTotals**](docs/ChartsAPI.md#getcharttotals) | **Post** /api/v1/charts/totals | Get chart data
*ClaimLicenseAPI* | [**GetClaimLicense**](docs/ClaimLicenseAPI.md#getclaimlicense) | **Get** /api/v1/claim_license | Attempts to claim a white label license
*ClientGatewayTokensAPI* | [**DeleteClientGatewayToken**](docs/ClientGatewayTokensAPI.md#deleteclientgatewaytoken) | **Delete** /api/v1/client_gateway_tokens/{id} | Deletes a client
*ClientGatewayTokensAPI* | [**EditClientGatewayToken**](docs/ClientGatewayTokensAPI.md#editclientgatewaytoken) | **Get** /api/v1/client_gateway_tokens/{id}/edit | Shows a client payment token for editting
*ClientGatewayTokensAPI* | [**GetClientGatewayTokens**](docs/ClientGatewayTokensAPI.md#getclientgatewaytokens) | **Get** /api/v1/client_gateway_tokens | List of client payment tokens
*ClientGatewayTokensAPI* | [**GetClientGatewayTokensCreate**](docs/ClientGatewayTokensAPI.md#getclientgatewaytokenscreate) | **Get** /api/v1/client_gateway_tokens/create | Gets a new blank client payment token object
*ClientGatewayTokensAPI* | [**ShowClientGatewayToken**](docs/ClientGatewayTokensAPI.md#showclientgatewaytoken) | **Get** /api/v1/client_gateway_tokens/{id} | Shows a client payment token
*ClientGatewayTokensAPI* | [**StoreClientGatewayToken**](docs/ClientGatewayTokensAPI.md#storeclientgatewaytoken) | **Post** /api/v1/client_gateway_tokens | Adds a client payment token
*ClientGatewayTokensAPI* | [**UpdateClientGatewayToken**](docs/ClientGatewayTokensAPI.md#updateclientgatewaytoken) | **Put** /api/v1/client_gateway_tokens/{id} | Updates a client payment token
*ClientsAPI* | [**BulkClients**](docs/ClientsAPI.md#bulkclients) | **Post** /api/v1/clients/bulk | Bulk client actions
*ClientsAPI* | [**ClientStatement**](docs/ClientsAPI.md#clientstatement) | **Post** /api/v1/client_statement | Client statement PDF
*ClientsAPI* | [**DeleteClient**](docs/ClientsAPI.md#deleteclient) | **Delete** /api/v1/clients/{id} | Delete client
*ClientsAPI* | [**EditClient**](docs/ClientsAPI.md#editclient) | **Get** /api/v1/clients/{id}/edit | Edit Client
*ClientsAPI* | [**GetClients**](docs/ClientsAPI.md#getclients) | **Get** /api/v1/clients | List clients 
*ClientsAPI* | [**GetClientsCreate**](docs/ClientsAPI.md#getclientscreate) | **Get** /api/v1/clients/create | Blank Client
*ClientsAPI* | [**MergeClient**](docs/ClientsAPI.md#mergeclient) | **Post** /api/v1/clients/{id}/{mergeable_client_hashed_id}/merge | Merge client
*ClientsAPI* | [**PurgeClient**](docs/ClientsAPI.md#purgeclient) | **Post** /api/v1/clients/{id}/purge | Purge client
*ClientsAPI* | [**ReactivateEmail**](docs/ClientsAPI.md#reactivateemail) | **Post** /api/v1/reactivate_email/{bounce_id} | Removes email suppression of a user in the system
*ClientsAPI* | [**ShowClient**](docs/ClientsAPI.md#showclient) | **Get** /api/v1/clients/{id} | Show client
*ClientsAPI* | [**StoreClient**](docs/ClientsAPI.md#storeclient) | **Post** /api/v1/clients | Create client
*ClientsAPI* | [**UpdateClient**](docs/ClientsAPI.md#updateclient) | **Put** /api/v1/clients/{id} | Update client
*ClientsAPI* | [**UpdateClientTaxData**](docs/ClientsAPI.md#updateclienttaxdata) | **Post** /api/v1/clients/{client}/updateTaxData | Update tax data
*ClientsAPI* | [**UploadClient**](docs/ClientsAPI.md#uploadclient) | **Post** /api/v1/clients/{id}/upload | Add client document
*CompaniesAPI* | [**DeleteCompany**](docs/CompaniesAPI.md#deletecompany) | **Delete** /api/v1/companies/{id} | Deletes a company
*CompaniesAPI* | [**EditCompany**](docs/CompaniesAPI.md#editcompany) | **Get** /api/v1/companies/{id}/edit | Shows an company for editting
*CompaniesAPI* | [**GetCompanies**](docs/CompaniesAPI.md#getcompanies) | **Get** /api/v1/companies | Gets a list of companies
*CompaniesAPI* | [**GetCompaniesCreate**](docs/CompaniesAPI.md#getcompaniescreate) | **Get** /api/v1/companies/create | Gets a new blank company object
*CompaniesAPI* | [**SetDefaultCompany**](docs/CompaniesAPI.md#setdefaultcompany) | **Post** /api/v1/companies/{company}/default | Sets the company as the default company.
*CompaniesAPI* | [**ShowCompany**](docs/CompaniesAPI.md#showcompany) | **Get** /api/v1/companies/{id} | Shows an company
*CompaniesAPI* | [**ShowCurrentCompany**](docs/CompaniesAPI.md#showcurrentcompany) | **Post** /api/v1/companies/current | Returns the current comapny
*CompaniesAPI* | [**StoreCompany**](docs/CompaniesAPI.md#storecompany) | **Post** /api/v1/companies | Adds a company
*CompaniesAPI* | [**UpdateCompany**](docs/CompaniesAPI.md#updatecompany) | **Put** /api/v1/companies/{id} | Updates an company
*CompaniesAPI* | [**UploadCompanies**](docs/CompaniesAPI.md#uploadcompanies) | **Post** /api/v1/companies/{id}/upload | Uploads a document to a company
*CompanyGatewaysAPI* | [**BulkCompanyGateways**](docs/CompanyGatewaysAPI.md#bulkcompanygateways) | **Post** /api/v1/company_gateways/bulk | Performs bulk actions on an array of company_gateways
*CompanyGatewaysAPI* | [**DeleteCompanyGateway**](docs/CompanyGatewaysAPI.md#deletecompanygateway) | **Delete** /api/v1/company_gateways/{id} | Deletes a CompanyGateway
*CompanyGatewaysAPI* | [**EditCompanyGateway**](docs/CompanyGatewaysAPI.md#editcompanygateway) | **Get** /api/v1/company_gateways/{id}/edit | Shows an CompanyGateway for editting
*CompanyGatewaysAPI* | [**GetCompanyGateways**](docs/CompanyGatewaysAPI.md#getcompanygateways) | **Get** /api/v1/company_gateways | Gets a list of company_gateways
*CompanyGatewaysAPI* | [**GetCompanyGatewaysCreate**](docs/CompanyGatewaysAPI.md#getcompanygatewayscreate) | **Get** /api/v1/company_gateways/create | Gets a new blank CompanyGateway object
*CompanyGatewaysAPI* | [**ShowCompanyGateway**](docs/CompanyGatewaysAPI.md#showcompanygateway) | **Get** /api/v1/company_gateways/{id} | Shows an CompanyGateway
*CompanyGatewaysAPI* | [**StoreCompanyGateway**](docs/CompanyGatewaysAPI.md#storecompanygateway) | **Post** /api/v1/company_gateways | Adds a CompanyGateway
*CompanyGatewaysAPI* | [**UpdateCompanyGateway**](docs/CompanyGatewaysAPI.md#updatecompanygateway) | **Put** /api/v1/company_gateways/{id} | Updates an CompanyGateway
*CompanyLedgerAPI* | [**GetCompanyLedger**](docs/CompanyLedgerAPI.md#getcompanyledger) | **Get** /api/v1/company_ledger | Gets a list of company_ledger
*CompanyUserAPI* | [**UpdateCompanyUser**](docs/CompanyUserAPI.md#updatecompanyuser) | **Post** /api/v1/company_users | Update a company user record
*ConnectedAccountAPI* | [**ConnectedAccount**](docs/ConnectedAccountAPI.md#connectedaccount) | **Post** /api/v1/connected_account | Connect an oauth user to an existing user
*CreditsAPI* | [**BulkCredits**](docs/CreditsAPI.md#bulkcredits) | **Post** /api/v1/credits/bulk | Bulk credit actions
*CreditsAPI* | [**DeleteCredit**](docs/CreditsAPI.md#deletecredit) | **Delete** /api/v1/credits/{id} | Delete credit
*CreditsAPI* | [**EditCredit**](docs/CreditsAPI.md#editcredit) | **Get** /api/v1/credits/{id}/edit | Edit credit
*CreditsAPI* | [**GetCredits**](docs/CreditsAPI.md#getcredits) | **Get** /api/v1/credits | List credits
*CreditsAPI* | [**GetCreditsCreate**](docs/CreditsAPI.md#getcreditscreate) | **Get** /api/v1/credits/create | Blank credit
*CreditsAPI* | [**ShowCredit**](docs/CreditsAPI.md#showcredit) | **Get** /api/v1/credits/{id} | Show credit
*CreditsAPI* | [**StoreCredit**](docs/CreditsAPI.md#storecredit) | **Post** /api/v1/credits | Create credit
*CreditsAPI* | [**UpdateCredit**](docs/CreditsAPI.md#updatecredit) | **Put** /api/v1/credits/{id} | Update credit
*CreditsAPI* | [**UploadCredits**](docs/CreditsAPI.md#uploadcredits) | **Post** /api/v1/credits/{id}/upload | Upload a credit document
*DesignsAPI* | [**BulkDesigns**](docs/DesignsAPI.md#bulkdesigns) | **Post** /api/v1/designs/bulk | Performs bulk actions on an array of designs
*DesignsAPI* | [**DeleteDesign**](docs/DesignsAPI.md#deletedesign) | **Delete** /api/v1/designs/{id} | Deletes a design
*DesignsAPI* | [**EditDesign**](docs/DesignsAPI.md#editdesign) | **Get** /api/v1/designs/{id}/edit | Shows a design for editting
*DesignsAPI* | [**GetDesigns**](docs/DesignsAPI.md#getdesigns) | **Get** /api/v1/designs | Gets a list of designs
*DesignsAPI* | [**GetDesignsCreate**](docs/DesignsAPI.md#getdesignscreate) | **Get** /api/v1/designs/create | Gets a new blank design object
*DesignsAPI* | [**ShowDesign**](docs/DesignsAPI.md#showdesign) | **Get** /api/v1/designs/{id} | Shows a design
*DesignsAPI* | [**StoreDesign**](docs/DesignsAPI.md#storedesign) | **Post** /api/v1/designs | Adds a design
*DesignsAPI* | [**UpdateDesign**](docs/DesignsAPI.md#updatedesign) | **Put** /api/v1/designs/{id} | Updates a design
*DocumentsAPI* | [**GetDocuments**](docs/DocumentsAPI.md#getdocuments) | **Get** /api/v1/documents | Gets a list of documents
*EmailsAPI* | [**SendEmailTemplate**](docs/EmailsAPI.md#sendemailtemplate) | **Post** /api/v1/emails | Sends an email for an entity
*ExpenseAPI* | [**UploadExpense**](docs/ExpenseAPI.md#uploadexpense) | **Post** /api/v1/expenses/{id}/upload | Uploads a document to a expense
*ExpenseCategoriesAPI* | [**BulkExpenseCategorys**](docs/ExpenseCategoriesAPI.md#bulkexpensecategorys) | **Post** /api/v1/expense_categories/bulk | Performs bulk actions on an array of ExpenseCategorys
*ExpenseCategoriesAPI* | [**DeleteExpenseCategory**](docs/ExpenseCategoriesAPI.md#deleteexpensecategory) | **Delete** /api/v1/expense_categories/{id} | Deletes a ExpenseCategory
*ExpenseCategoriesAPI* | [**EditExpenseCategory**](docs/ExpenseCategoriesAPI.md#editexpensecategory) | **Get** /api/v1/expense_categories/{id}/edit | Shows a Expens Category for editting
*ExpenseCategoriesAPI* | [**GetExpenseCategoryCreate**](docs/ExpenseCategoriesAPI.md#getexpensecategorycreate) | **Get** /api/v1/expense_categories/create | Gets a new blank Expens Category object
*ExpenseCategoriesAPI* | [**GetExpenseCategorys**](docs/ExpenseCategoriesAPI.md#getexpensecategorys) | **Get** /api/v1/expense_categories | Gets a list of expense_categories
*ExpenseCategoriesAPI* | [**ShowExpenseCategory**](docs/ExpenseCategoriesAPI.md#showexpensecategory) | **Get** /api/v1/expense_categories/{id} | Shows a Expens Category
*ExpenseCategoriesAPI* | [**StoreExpenseCategory**](docs/ExpenseCategoriesAPI.md#storeexpensecategory) | **Post** /api/v1/expense_categories | Adds a expense category
*ExpenseCategoriesAPI* | [**UpdateExpenseCategory**](docs/ExpenseCategoriesAPI.md#updateexpensecategory) | **Put** /api/v1/expense_categories/{id} | Updates a tax rate
*ExpensesAPI* | [**BulkExpenses**](docs/ExpensesAPI.md#bulkexpenses) | **Post** /api/v1/expenses/bulk | Performs bulk actions on an array of expenses
*ExpensesAPI* | [**DeleteExpense**](docs/ExpensesAPI.md#deleteexpense) | **Delete** /api/v1/expenses/{id} | Deletes a expense
*ExpensesAPI* | [**EditExpense**](docs/ExpensesAPI.md#editexpense) | **Get** /api/v1/expenses/{id}/edit | Shows a expense for editing
*ExpensesAPI* | [**GetExpenses**](docs/ExpensesAPI.md#getexpenses) | **Get** /api/v1/expenses | Gets a list of expenses
*ExpensesAPI* | [**GetExpensesCreate**](docs/ExpensesAPI.md#getexpensescreate) | **Get** /api/v1/expenses/create | Gets a new blank expense object
*ExpensesAPI* | [**ShowExpense**](docs/ExpensesAPI.md#showexpense) | **Get** /api/v1/expenses/{id} | Shows a expense
*ExpensesAPI* | [**StoreExpense**](docs/ExpensesAPI.md#storeexpense) | **Post** /api/v1/expenses | Adds an expense
*ExpensesAPI* | [**UpdateExpense**](docs/ExpensesAPI.md#updateexpense) | **Put** /api/v1/expenses/{id} | Updates a expense
*ExportAPI* | [**GetExport**](docs/ExportAPI.md#getexport) | **Post** /api/v1/export | Export data from the system
*GroupSettingsAPI* | [**BulkGroupSettings**](docs/GroupSettingsAPI.md#bulkgroupsettings) | **Post** /api/v1/group_settings/bulk | Performs bulk actions on an array of group_settings
*GroupSettingsAPI* | [**DeleteGroupSetting**](docs/GroupSettingsAPI.md#deletegroupsetting) | **Delete** /api/v1/group_settings/{id} | Deletes a GroupSetting
*GroupSettingsAPI* | [**EditGroupSetting**](docs/GroupSettingsAPI.md#editgroupsetting) | **Get** /api/v1/group_settings/{id}/edit | Shows an GroupSetting for editting
*GroupSettingsAPI* | [**GetGroupSettings**](docs/GroupSettingsAPI.md#getgroupsettings) | **Get** /api/v1/group_settings | Gets a list of group_settings
*GroupSettingsAPI* | [**GetGroupSettingsCreate**](docs/GroupSettingsAPI.md#getgroupsettingscreate) | **Get** /api/v1/group_settings/create | Gets a new blank GroupSetting object
*GroupSettingsAPI* | [**ShowGroupSetting**](docs/GroupSettingsAPI.md#showgroupsetting) | **Get** /api/v1/group_settings/{id} | Shows an GroupSetting
*GroupSettingsAPI* | [**StoreGroupSetting**](docs/GroupSettingsAPI.md#storegroupsetting) | **Post** /api/v1/group_settings | Adds a GroupSetting
*GroupSettingsAPI* | [**UpdateGroupSetting**](docs/GroupSettingsAPI.md#updategroupsetting) | **Put** /api/v1/group_settings/{id} | Updates an GroupSetting
*GroupSettingsAPI* | [**UploadGroupSetting**](docs/GroupSettingsAPI.md#uploadgroupsetting) | **Post** /api/v1/group_settings/{id}/upload | Uploads a document to a group setting
*HealthCheckAPI* | [**GetHealthCheck**](docs/HealthCheckAPI.md#gethealthcheck) | **Get** /api/v1/health_check | Attempts to get a health check from the API
*ImportAPI* | [**GetImportJson**](docs/ImportAPI.md#getimportjson) | **Post** /api/v1/import_json | Import data from the system
*ImportsAPI* | [**Preimport**](docs/ImportsAPI.md#preimport) | **Post** /api/v1/preimport | Pre Import checks - returns a reference to the job and the headers of the CSV
*InvoicesAPI* | [**ActionInvoice**](docs/InvoicesAPI.md#actioninvoice) | **Get** /api/v1/invoices/{id}/{action} | Custom invoice action
*InvoicesAPI* | [**ApiV1InvoiceInvitationKeyDownloadGet**](docs/InvoicesAPI.md#apiv1invoiceinvitationkeydownloadget) | **Get** /api/v1/invoice/{invitation_key}/download | Download invoice PDF
*InvoicesAPI* | [**ApiV1InvoicesCreateGet**](docs/InvoicesAPI.md#apiv1invoicescreateget) | **Get** /api/v1/invoices/create | Blank invoice
*InvoicesAPI* | [**ApiV1InvoicesIdDeliveryNoteGet**](docs/InvoicesAPI.md#apiv1invoicesiddeliverynoteget) | **Get** /api/v1/invoices/{id}/delivery_note | Download delivery note
*InvoicesAPI* | [**ApiV1InvoicesIdUploadPost**](docs/InvoicesAPI.md#apiv1invoicesiduploadpost) | **Post** /api/v1/invoices/{id}/upload | Add invoice document
*InvoicesAPI* | [**BulkInvoices**](docs/InvoicesAPI.md#bulkinvoices) | **Post** /api/v1/invoices/bulk | Bulk invoice actions
*InvoicesAPI* | [**DeleteInvoice**](docs/InvoicesAPI.md#deleteinvoice) | **Delete** /api/v1/invoices/{id} | Delete invoice
*InvoicesAPI* | [**EditInvoice**](docs/InvoicesAPI.md#editinvoice) | **Get** /api/v1/invoices/{id}/edit | Edit invoice
*InvoicesAPI* | [**GetInvoices**](docs/InvoicesAPI.md#getinvoices) | **Get** /api/v1/invoices | List invoices
*InvoicesAPI* | [**ShowInvoice**](docs/InvoicesAPI.md#showinvoice) | **Get** /api/v1/invoices/{id} | Show invoice
*InvoicesAPI* | [**StoreInvoice**](docs/InvoicesAPI.md#storeinvoice) | **Post** /api/v1/invoices | Create invoice
*InvoicesAPI* | [**UpdateInvoice**](docs/InvoicesAPI.md#updateinvoice) | **Put** /api/v1/invoices/{id} | Update invoice
*LocationsAPI* | [**BulkLocations**](docs/LocationsAPI.md#bulklocations) | **Post** /api/v1/locations/bulk | Bulk location actions
*LocationsAPI* | [**DeleteLocation**](docs/LocationsAPI.md#deletelocation) | **Delete** /api/v1/locations/{id} | Delete location
*LocationsAPI* | [**GetLocations**](docs/LocationsAPI.md#getlocations) | **Get** /api/v1/locations | List locations
*LocationsAPI* | [**GetLocationsCreate**](docs/LocationsAPI.md#getlocationscreate) | **Get** /api/v1/locations/create | Blank Location
*LocationsAPI* | [**ShowLocation**](docs/LocationsAPI.md#showlocation) | **Get** /api/v1/locations/{id} | Show location
*LocationsAPI* | [**StoreLocation**](docs/LocationsAPI.md#storelocation) | **Post** /api/v1/locations | Create location
*LocationsAPI* | [**UpdateLocation**](docs/LocationsAPI.md#updatelocation) | **Put** /api/v1/locations/{id} | Update location
*LogoutAPI* | [**GetLogout**](docs/LogoutAPI.md#getlogout) | **Post** /api/v1/logout | Logs the user out of their current session
*MigrationAPI* | [**PostPurgeCompany**](docs/MigrationAPI.md#postpurgecompany) | **Post** /api/v1/migration/purge/{company} | Attempts to purge a company record and all its child records
*MigrationAPI* | [**PostPurgeCompanySaveSettings**](docs/MigrationAPI.md#postpurgecompanysavesettings) | **Post** /api/v1/migration/purge_save_settings/{company} | Attempts to purge a companies child records but save the company record and its settings
*MigrationAPI* | [**PostStartMigration**](docs/MigrationAPI.md#poststartmigration) | **Post** /api/v1/migration/start | Starts the migration from previous version of Invoice Ninja
*OneTimeTokenAPI* | [**OneTimeToken**](docs/OneTimeTokenAPI.md#onetimetoken) | **Post** /api/v1/one_time_token | Attempts to create a one time token
*PaymentTermsAPI* | [**BulkPaymentTerms**](docs/PaymentTermsAPI.md#bulkpaymentterms) | **Post** /api/v1/payment_terms/bulk | Performs bulk actions on an array of payment terms
*PaymentTermsAPI* | [**EditPaymentTerms**](docs/PaymentTermsAPI.md#editpaymentterms) | **Get** /api/v1/payment_terms/{id}/edit | Shows an Payment Term for editting
*PaymentTermsAPI* | [**GetPaymentTerms**](docs/PaymentTermsAPI.md#getpaymentterms) | **Get** /api/v1/payment_terms | Gets a list of payment terms
*PaymentTermsAPI* | [**GetPaymentTermsCreate**](docs/PaymentTermsAPI.md#getpaymenttermscreate) | **Get** /api/v1/payment_terms/create | Gets a new blank PaymentTerm object
*PaymentTermsAPI* | [**ShowPaymentTerm**](docs/PaymentTermsAPI.md#showpaymentterm) | **Get** /api/v1/payment_terms/{id} | Shows a Payment Term
*PaymentTermsAPI* | [**StorePaymentTerm**](docs/PaymentTermsAPI.md#storepaymentterm) | **Post** /api/v1/payment_terms | Adds a Payment
*PaymentTermsAPI* | [**UpdatePaymentTerm**](docs/PaymentTermsAPI.md#updatepaymentterm) | **Put** /api/v1/payment_terms/{id} | Updates a Payment Term
*PaymentTermssAPI* | [**DeletePaymentTerm**](docs/PaymentTermssAPI.md#deletepaymentterm) | **Delete** /api/v1/payment_terms/{id} | Deletes a Payment Term
*PaymentsAPI* | [**ActionPayment**](docs/PaymentsAPI.md#actionpayment) | **Get** /api/v1/payments/{id}/{action} | Custom payment actions
*PaymentsAPI* | [**BulkPayments**](docs/PaymentsAPI.md#bulkpayments) | **Post** /api/v1/payments/bulk | Bulk payment actions
*PaymentsAPI* | [**DeletePayment**](docs/PaymentsAPI.md#deletepayment) | **Delete** /api/v1/payments/{id} | Delete payment
*PaymentsAPI* | [**EditPayment**](docs/PaymentsAPI.md#editpayment) | **Get** /api/v1/payments/{id}/edit | Edit payment
*PaymentsAPI* | [**GetPayments**](docs/PaymentsAPI.md#getpayments) | **Get** /api/v1/payments | List payments
*PaymentsAPI* | [**GetPaymentsCreate**](docs/PaymentsAPI.md#getpaymentscreate) | **Get** /api/v1/payments/create | Blank payment
*PaymentsAPI* | [**ShowPayment**](docs/PaymentsAPI.md#showpayment) | **Get** /api/v1/payments/{id} | Show payment
*PaymentsAPI* | [**StorePayment**](docs/PaymentsAPI.md#storepayment) | **Post** /api/v1/payments | Create payment
*PaymentsAPI* | [**StoreRefund**](docs/PaymentsAPI.md#storerefund) | **Post** /api/v1/payments/refund | Refund payment
*PaymentsAPI* | [**UpdatePayment**](docs/PaymentsAPI.md#updatepayment) | **Put** /api/v1/payments/{id} | Update payment
*PaymentsAPI* | [**UploadPayment**](docs/PaymentsAPI.md#uploadpayment) | **Post** /api/v1/payments/{id}/upload | Upload a payment document
*PingAPI* | [**GetPing**](docs/PingAPI.md#getping) | **Get** /api/v1/ping | Attempts to ping the API
*PostmarkAPI* | [**ConfirmApplePurchase**](docs/PostmarkAPI.md#confirmapplepurchase) | **Post** /api/v1/apple/confirm_purchase | Processing webhooks from Apple for in app purchases
*PostmarkAPI* | [**PostmarkWebhook**](docs/PostmarkAPI.md#postmarkwebhook) | **Post** /api/v1/postmark_webhook | Processing webhooks from PostMark
*PostmarkAPI* | [**ProcessAppleWebhook**](docs/PostmarkAPI.md#processapplewebhook) | **Post** /api/v1/apple/process_webhook | Processing event webhooks from Apple for in purchase / subscription status update
*PreviewAPI* | [**GetPreview**](docs/PreviewAPI.md#getpreview) | **Post** /api/v1/preview | Returns a pdf preview
*PreviewAPI* | [**GetPreviewPurchaseOrder**](docs/PreviewAPI.md#getpreviewpurchaseorder) | **Post** /api/v1/preview/purchase_order | Returns a pdf preview for purchase order
*ProductsAPI* | [**BulkProducts**](docs/ProductsAPI.md#bulkproducts) | **Post** /api/v1/products/bulk | Bulk product actions
*ProductsAPI* | [**DeleteProduct**](docs/ProductsAPI.md#deleteproduct) | **Delete** /api/v1/products/{id} | Delete product
*ProductsAPI* | [**EditProduct**](docs/ProductsAPI.md#editproduct) | **Get** /api/v1/products/{id}/edit | Edit product
*ProductsAPI* | [**GetProducts**](docs/ProductsAPI.md#getproducts) | **Get** /api/v1/products | List products
*ProductsAPI* | [**GetProductsCreate**](docs/ProductsAPI.md#getproductscreate) | **Get** /api/v1/products/create | Blank product
*ProductsAPI* | [**ShowProduct**](docs/ProductsAPI.md#showproduct) | **Get** /api/v1/products/{id} | Show product
*ProductsAPI* | [**StoreProduct**](docs/ProductsAPI.md#storeproduct) | **Post** /api/v1/products | Create Product
*ProductsAPI* | [**UpdateProduct**](docs/ProductsAPI.md#updateproduct) | **Put** /api/v1/products/{id} | Update product
*ProductsAPI* | [**UploadProduct**](docs/ProductsAPI.md#uploadproduct) | **Post** /api/v1/products/{id}/upload | Add product document
*ProjectsAPI* | [**BulkProjects**](docs/ProjectsAPI.md#bulkprojects) | **Post** /api/v1/projects/bulk | Bulk project actions
*ProjectsAPI* | [**DeleteProject**](docs/ProjectsAPI.md#deleteproject) | **Delete** /api/v1/projects/{id} | Delete project
*ProjectsAPI* | [**EditProject**](docs/ProjectsAPI.md#editproject) | **Get** /api/v1/projects/{id}/edit | Edit project
*ProjectsAPI* | [**GetProjects**](docs/ProjectsAPI.md#getprojects) | **Get** /api/v1/projects | List projects
*ProjectsAPI* | [**GetProjectsCreate**](docs/ProjectsAPI.md#getprojectscreate) | **Get** /api/v1/projects/create | Blank project
*ProjectsAPI* | [**ShowProject**](docs/ProjectsAPI.md#showproject) | **Get** /api/v1/projects/{id} | Show project
*ProjectsAPI* | [**StoreProject**](docs/ProjectsAPI.md#storeproject) | **Post** /api/v1/projects | Create project
*ProjectsAPI* | [**UpdateProject**](docs/ProjectsAPI.md#updateproject) | **Put** /api/v1/projects/{id} | Update project
*ProjectsAPI* | [**UploadProject**](docs/ProjectsAPI.md#uploadproject) | **Post** /api/v1/projects/{id}/upload | Uploads a project document
*PurchaseOrdersAPI* | [**ActionPurchaseOrder**](docs/PurchaseOrdersAPI.md#actionpurchaseorder) | **Get** /api/v1/purchase_orders/{id}/{action} | Custom purchase order actions
*PurchaseOrdersAPI* | [**BulkPurchaseOrderss**](docs/PurchaseOrdersAPI.md#bulkpurchaseorderss) | **Post** /api/v1/purchase_orders/bulk | Bulk purchase order action
*PurchaseOrdersAPI* | [**DeletePurchaseOrder**](docs/PurchaseOrdersAPI.md#deletepurchaseorder) | **Delete** /api/v1/purchase_order/{id} | Delete purchase order
*PurchaseOrdersAPI* | [**DownloadPurchaseOrder**](docs/PurchaseOrdersAPI.md#downloadpurchaseorder) | **Get** /api/v1/purchase_order/{invitation_key}/download | Download a purchase order PDF
*PurchaseOrdersAPI* | [**EditPurchaseOrder**](docs/PurchaseOrdersAPI.md#editpurchaseorder) | **Get** /api/v1/purchase_orders/{id}/edit | Edit purchase order
*PurchaseOrdersAPI* | [**GetPurchaseOrderCreate**](docs/PurchaseOrdersAPI.md#getpurchaseordercreate) | **Get** /api/v1/purchase_orders/create | Blank purchase order
*PurchaseOrdersAPI* | [**GetPurchaseOrders**](docs/PurchaseOrdersAPI.md#getpurchaseorders) | **Get** /api/v1/purchase_orders | List purchase orders
*PurchaseOrdersAPI* | [**ShowPurchaseOrder**](docs/PurchaseOrdersAPI.md#showpurchaseorder) | **Get** /api/v1/purchase_orders/{id} | Show purchase order
*PurchaseOrdersAPI* | [**StorePurchaseOrder**](docs/PurchaseOrdersAPI.md#storepurchaseorder) | **Post** /api/v1/purchase_orders | Create purchase order
*PurchaseOrdersAPI* | [**UpdatePurchaseOrder**](docs/PurchaseOrdersAPI.md#updatepurchaseorder) | **Put** /api/v1/purchase_order/{id} | Update purchase order
*PurchaseOrdersAPI* | [**UploadPurchaseOrder**](docs/PurchaseOrdersAPI.md#uploadpurchaseorder) | **Post** /api/v1/purchase_orders/{id}/upload | Uploads a purchase order document
*QuotesAPI* | [**ActionQuote**](docs/QuotesAPI.md#actionquote) | **Get** /api/v1/quotes/{id}/{action} | Performs a custom action on an Quote
*QuotesAPI* | [**BulkQuotes**](docs/QuotesAPI.md#bulkquotes) | **Post** /api/v1/quotes/bulk | Bulk quote actions
*QuotesAPI* | [**DeleteQuote**](docs/QuotesAPI.md#deletequote) | **Delete** /api/v1/quotes/{id} | Delete quote
*QuotesAPI* | [**DownloadCredit**](docs/QuotesAPI.md#downloadcredit) | **Get** /api/v1/credit/{invitation_key}/download | Download quote PDF
*QuotesAPI* | [**DownloadQuote**](docs/QuotesAPI.md#downloadquote) | **Get** /api/v1/quote/{invitation_key}/download | Download quote PDF
*QuotesAPI* | [**EditQuote**](docs/QuotesAPI.md#editquote) | **Get** /api/v1/quotes/{id}/edit | Edit quote
*QuotesAPI* | [**GetQuotes**](docs/QuotesAPI.md#getquotes) | **Get** /api/v1/quotes | List quotes
*QuotesAPI* | [**GetQuotesCreate**](docs/QuotesAPI.md#getquotescreate) | **Get** /api/v1/quotes/create | Blank quote
*QuotesAPI* | [**ShowQuote**](docs/QuotesAPI.md#showquote) | **Get** /api/v1/quotes/{id} | Show quote
*QuotesAPI* | [**StoreQuote**](docs/QuotesAPI.md#storequote) | **Post** /api/v1/quotes | Create quote
*QuotesAPI* | [**UpdateQuote**](docs/QuotesAPI.md#updatequote) | **Put** /api/v1/quotes/{id} | Update quote
*QuotesAPI* | [**UploadQuote**](docs/QuotesAPI.md#uploadquote) | **Post** /api/v1/quotes/{id}/upload | Upload a quote document
*RecurringExpenseAPI* | [**UploadRecurringExpense**](docs/RecurringExpenseAPI.md#uploadrecurringexpense) | **Post** /api/v1/recurring_expenses/{id}/upload | Uploads a document to a recurring_expense
*RecurringExpensesAPI* | [**BulkRecurringExpenses**](docs/RecurringExpensesAPI.md#bulkrecurringexpenses) | **Post** /api/v1/recurring_expenses/bulk | Performs bulk actions on an array of recurring_expenses
*RecurringExpensesAPI* | [**DeleteRecurringExpense**](docs/RecurringExpensesAPI.md#deleterecurringexpense) | **Delete** /api/v1/recurring_expenses/{id} | Deletes a recurring expense
*RecurringExpensesAPI* | [**EditRecurringExpense**](docs/RecurringExpensesAPI.md#editrecurringexpense) | **Get** /api/v1/recurring_expenses/{id}/edit | Shows a recurring expense for editting
*RecurringExpensesAPI* | [**GetRecurringExpenses**](docs/RecurringExpensesAPI.md#getrecurringexpenses) | **Get** /api/v1/recurring_expenses | Gets a list of recurring_expenses
*RecurringExpensesAPI* | [**GetRecurringExpensesCreate**](docs/RecurringExpensesAPI.md#getrecurringexpensescreate) | **Get** /api/v1/recurring_expenses/create | Gets a new blank recurring expense object
*RecurringExpensesAPI* | [**ShowRecurringExpense**](docs/RecurringExpensesAPI.md#showrecurringexpense) | **Get** /api/v1/recurring_expenses/{id} | Shows a recurring expense
*RecurringExpensesAPI* | [**StoreRecurringExpense**](docs/RecurringExpensesAPI.md#storerecurringexpense) | **Post** /api/v1/recurring_expenses | Adds a recurring expense
*RecurringExpensesAPI* | [**UpdateRecurringExpense**](docs/RecurringExpensesAPI.md#updaterecurringexpense) | **Put** /api/v1/recurring_expenses/{id} | Updates a recurring expense
*RecurringInvoicesAPI* | [**ActionRecurringInvoice**](docs/RecurringInvoicesAPI.md#actionrecurringinvoice) | **Get** /api/v1/recurring_invoices/{id}/{action} | Custom recurring invoice action
*RecurringInvoicesAPI* | [**BulkRecurringInvoices**](docs/RecurringInvoicesAPI.md#bulkrecurringinvoices) | **Post** /api/v1/recurring_invoices/bulk | Bulk recurring invoice actions
*RecurringInvoicesAPI* | [**DeleteRecurringInvoice**](docs/RecurringInvoicesAPI.md#deleterecurringinvoice) | **Delete** /api/v1/recurring_invoices/{id} | Delete recurring invoice
*RecurringInvoicesAPI* | [**DownloadRecurringInvoice**](docs/RecurringInvoicesAPI.md#downloadrecurringinvoice) | **Get** /api/v1/recurring_invoice/{invitation_key}/download | Download recurring invoice PDF
*RecurringInvoicesAPI* | [**EditRecurringInvoice**](docs/RecurringInvoicesAPI.md#editrecurringinvoice) | **Get** /api/v1/recurring_invoices/{id}/edit | Edit recurring invoice
*RecurringInvoicesAPI* | [**GetRecurringInvoices**](docs/RecurringInvoicesAPI.md#getrecurringinvoices) | **Get** /api/v1/recurring_invoices | List recurring invoices
*RecurringInvoicesAPI* | [**GetRecurringInvoicesCreate**](docs/RecurringInvoicesAPI.md#getrecurringinvoicescreate) | **Get** /api/v1/recurring_invoices/create | Blank recurring invoice
*RecurringInvoicesAPI* | [**ShowRecurringInvoice**](docs/RecurringInvoicesAPI.md#showrecurringinvoice) | **Get** /api/v1/recurring_invoices/{id} | Show recurring invoice
*RecurringInvoicesAPI* | [**StoreRecurringInvoice**](docs/RecurringInvoicesAPI.md#storerecurringinvoice) | **Post** /api/v1/recurring_invoices | Create recurring invoice
*RecurringInvoicesAPI* | [**UpdateRecurringInvoice**](docs/RecurringInvoicesAPI.md#updaterecurringinvoice) | **Put** /api/v1/recurring_invoices/{id} | Update recurring invoice
*RecurringInvoicesAPI* | [**UploadRecurringInvoice**](docs/RecurringInvoicesAPI.md#uploadrecurringinvoice) | **Post** /api/v1/recurring_invoices/{id}/upload | Add recurring invoice document
*RecurringQuotesAPI* | [**ActionRecurringQuote**](docs/RecurringQuotesAPI.md#actionrecurringquote) | **Get** /api/v1/recurring_quotes/{id}/{action} | Performs a custom action on an RecurringQuote
*RecurringQuotesAPI* | [**BulkRecurringQuotes**](docs/RecurringQuotesAPI.md#bulkrecurringquotes) | **Post** /api/v1/recurring_quotes/bulk | Performs bulk actions on an array of recurring_quotes
*RecurringQuotesAPI* | [**DeleteRecurringQuote**](docs/RecurringQuotesAPI.md#deleterecurringquote) | **Delete** /api/v1/recurring_quotes/{id} | Deletes a RecurringQuote
*RecurringQuotesAPI* | [**EditRecurringQuote**](docs/RecurringQuotesAPI.md#editrecurringquote) | **Get** /api/v1/recurring_quotes/{id}/edit | Shows an RecurringQuote for editting
*RecurringQuotesAPI* | [**GetRecurringQuotes**](docs/RecurringQuotesAPI.md#getrecurringquotes) | **Get** /api/v1/recurring_quotes | Gets a list of recurring_quotes
*RecurringQuotesAPI* | [**GetRecurringQuotesCreate**](docs/RecurringQuotesAPI.md#getrecurringquotescreate) | **Get** /api/v1/recurring_quotes/create | Gets a new blank RecurringQuote object
*RecurringQuotesAPI* | [**ShowRecurringQuote**](docs/RecurringQuotesAPI.md#showrecurringquote) | **Get** /api/v1/recurring_quotes/{id} | Shows an RecurringQuote
*RecurringQuotesAPI* | [**StoreRecurringQuote**](docs/RecurringQuotesAPI.md#storerecurringquote) | **Post** /api/v1/recurring_quotes | Adds a RecurringQuote
*RecurringQuotesAPI* | [**UpdateRecurringQuote**](docs/RecurringQuotesAPI.md#updaterecurringquote) | **Put** /api/v1/recurring_quotes/{id} | Updates an RecurringQuote
*RefreshAPI* | [**Refresh**](docs/RefreshAPI.md#refresh) | **Post** /api/v1/refresh | Refresh data by timestamp
*ReportsAPI* | [**GetClientReport**](docs/ReportsAPI.md#getclientreport) | **Post** /api/v1/reports/clients | Client reports
*ReportsAPI* | [**GetContactReport**](docs/ReportsAPI.md#getcontactreport) | **Post** /api/v1/reports/contacts | Contact reports
*ReportsAPI* | [**GetCreditReport**](docs/ReportsAPI.md#getcreditreport) | **Post** /api/v1/reports/credit | Credit reports
*ReportsAPI* | [**GetDocumentReport**](docs/ReportsAPI.md#getdocumentreport) | **Post** /api/v1/reports/documents | Document reports
*ReportsAPI* | [**GetExpenseReport**](docs/ReportsAPI.md#getexpensereport) | **Post** /api/v1/reports/expense | Expense reports
*ReportsAPI* | [**GetInvoiceItemReport**](docs/ReportsAPI.md#getinvoiceitemreport) | **Post** /api/v1/reports/invoice_items | Invoice item reports
*ReportsAPI* | [**GetInvoiceReport**](docs/ReportsAPI.md#getinvoicereport) | **Post** /api/v1/reports/invoices | Invoice reports
*ReportsAPI* | [**GetPaymentReport**](docs/ReportsAPI.md#getpaymentreport) | **Post** /api/v1/reports/payments | Payment reports
*ReportsAPI* | [**GetProductReport**](docs/ReportsAPI.md#getproductreport) | **Post** /api/v1/reports/products | Product reports
*ReportsAPI* | [**GetProductSalesReport**](docs/ReportsAPI.md#getproductsalesreport) | **Post** /api/v1/reports/product_sales | Product Salesreports
*ReportsAPI* | [**GetProfitLossReport**](docs/ReportsAPI.md#getprofitlossreport) | **Post** /api/v1/reports/profitloss | Profit loss reports
*ReportsAPI* | [**GetQuoteItemReport**](docs/ReportsAPI.md#getquoteitemreport) | **Post** /api/v1/reports/quote_items | Quote item reports
*ReportsAPI* | [**GetQuoteReport**](docs/ReportsAPI.md#getquotereport) | **Post** /api/v1/reports/quotes | Quote reports
*ReportsAPI* | [**GetRecurringInvoiceReport**](docs/ReportsAPI.md#getrecurringinvoicereport) | **Post** /api/v1/reports/recurring_invoices | Recurring Invoice reports
*ReportsAPI* | [**GetTaskReport**](docs/ReportsAPI.md#gettaskreport) | **Post** /api/v1/reports/tasks | Task reports
*StaticsAPI* | [**GetStatics**](docs/StaticsAPI.md#getstatics) | **Get** /api/v1/statics | Gets a list of statics
*SubscriptionsAPI* | [**BulkSubscriptions**](docs/SubscriptionsAPI.md#bulksubscriptions) | **Post** /api/v1/subscriptions/bulk | Performs bulk actions on an array of subscriptions
*SubscriptionsAPI* | [**DeleteSubscription**](docs/SubscriptionsAPI.md#deletesubscription) | **Delete** /api/v1/subscriptions/{id} | Deletes a subscriptions
*SubscriptionsAPI* | [**EditSubscription**](docs/SubscriptionsAPI.md#editsubscription) | **Get** /api/v1/subscriptions/{id}/edit | Shows an subscriptions for editting
*SubscriptionsAPI* | [**GetSubscriptions**](docs/SubscriptionsAPI.md#getsubscriptions) | **Get** /api/v1/subscriptions | Gets a list of subscriptions
*SubscriptionsAPI* | [**GetSubscriptionsCreate**](docs/SubscriptionsAPI.md#getsubscriptionscreate) | **Get** /api/v1/subscriptions/create | Gets a new blank subscriptions object
*SubscriptionsAPI* | [**ShowSubscription**](docs/SubscriptionsAPI.md#showsubscription) | **Get** /api/v1/subscriptions/{id} | Shows an subscriptions
*SubscriptionsAPI* | [**StoreSubscription**](docs/SubscriptionsAPI.md#storesubscription) | **Post** /api/v1/subscriptions | Adds a subscriptions
*SubscriptionsAPI* | [**UpdateSubscription**](docs/SubscriptionsAPI.md#updatesubscription) | **Put** /api/v1/subscriptions/{id} | Updates an subscriptions
*SupportAPI* | [**SupportMessage**](docs/SupportAPI.md#supportmessage) | **Post** /api/v1/support/messages/send | Sends a support message to Invoice Ninja team
*SystemLogsAPI* | [**GetSystemLogs**](docs/SystemLogsAPI.md#getsystemlogs) | **Get** /api/v1/system_logs | Gets a list of system logs
*SystemLogsAPI* | [**ShowSystemLogs**](docs/SystemLogsAPI.md#showsystemlogs) | **Get** /api/v1/system_logs/{id} | Shows a system_logs
*TaskSchedulersAPI* | [**BulkTaskSchedulerActions**](docs/TaskSchedulersAPI.md#bulktaskscheduleractions) | **Post** /api/v1/task_schedulers/bulk | Performs bulk actions on an array of task_schedulers
*TaskSchedulersAPI* | [**CreateTaskScheduler**](docs/TaskSchedulersAPI.md#createtaskscheduler) | **Post** /api/v1/task_schedulers/ | Create task scheduler with job 
*TaskSchedulersAPI* | [**DestroyTaskScheduler**](docs/TaskSchedulersAPI.md#destroytaskscheduler) | **Delete** /api/v1/task_schedulers/{id} | Destroy Task Scheduler
*TaskSchedulersAPI* | [**GetTaskScheduler**](docs/TaskSchedulersAPI.md#gettaskscheduler) | **Get** /api/v1/task_schedulers/create | Gets a new blank scheduler object
*TaskSchedulersAPI* | [**GetTaskSchedulers**](docs/TaskSchedulersAPI.md#gettaskschedulers) | **Get** /api/v1/task_schedulers/ | Task Scheduler Index
*TaskSchedulersAPI* | [**ShowTaskScheduler**](docs/TaskSchedulersAPI.md#showtaskscheduler) | **Get** /api/v1/task_schedulers/{id} | Show given scheduler
*TaskSchedulersAPI* | [**UpdateTaskScheduler**](docs/TaskSchedulersAPI.md#updatetaskscheduler) | **Put** /api/v1/task_schedulers/{id} | Update task scheduler 
*TaskStatusAPI* | [**BulkTaskStatuss**](docs/TaskStatusAPI.md#bulktaskstatuss) | **Post** /api/v1/task_statuses/bulk | Performs bulk actions on an array of task statuses
*TaskStatusAPI* | [**EditTaskStatuss**](docs/TaskStatusAPI.md#edittaskstatuss) | **Get** /api/v1/task_statuses/{id}/edit | Shows an TaskStatusfor editting
*TaskStatusAPI* | [**GetTaskStatuses**](docs/TaskStatusAPI.md#gettaskstatuses) | **Get** /api/v1/task_statuses | Gets a list of task statuses
*TaskStatusAPI* | [**GetTaskStatussCreate**](docs/TaskStatusAPI.md#gettaskstatusscreate) | **Get** /api/v1/task_statuses/create | Gets a new blank TaskStatus object
*TaskStatusAPI* | [**ShowTaskStatus**](docs/TaskStatusAPI.md#showtaskstatus) | **Get** /api/v1/task_statuses/{id} | Shows a TaskStatus Term
*TaskStatusAPI* | [**StoreTaskStatus**](docs/TaskStatusAPI.md#storetaskstatus) | **Post** /api/v1/task_statuses | Adds a TaskStatus
*TaskStatusAPI* | [**UpdateTaskStatus**](docs/TaskStatusAPI.md#updatetaskstatus) | **Put** /api/v1/task_statuses/{id} | Updates a TaskStatus Term
*TaskStatussAPI* | [**DeleteTaskStatus**](docs/TaskStatussAPI.md#deletetaskstatus) | **Delete** /api/v1/task_statuses/{id} | Deletes a TaskStatus Term
*TasksAPI* | [**BulkTasks**](docs/TasksAPI.md#bulktasks) | **Post** /api/v1/tasks/bulk | Bulk task actions
*TasksAPI* | [**DeleteTask**](docs/TasksAPI.md#deletetask) | **Delete** /api/v1/tasks/{id} | Delete task
*TasksAPI* | [**EditTask**](docs/TasksAPI.md#edittask) | **Get** /api/v1/tasks/{id}/edit | Edit task
*TasksAPI* | [**GetTasks**](docs/TasksAPI.md#gettasks) | **Get** /api/v1/tasks | List tasks
*TasksAPI* | [**GetTasksCreate**](docs/TasksAPI.md#gettaskscreate) | **Get** /api/v1/tasks/create | Blank task
*TasksAPI* | [**ShowTask**](docs/TasksAPI.md#showtask) | **Get** /api/v1/tasks/{id} | Show task
*TasksAPI* | [**SortTasks**](docs/TasksAPI.md#sorttasks) | **Post** /api/v1/tasks/sort | Sort tasks on KanBan
*TasksAPI* | [**StoreTask**](docs/TasksAPI.md#storetask) | **Post** /api/v1/tasks | Create task
*TasksAPI* | [**UpdateTask**](docs/TasksAPI.md#updatetask) | **Put** /api/v1/tasks/{id} | Update task
*TasksAPI* | [**UploadTask**](docs/TasksAPI.md#uploadtask) | **Post** /api/v1/tasks/{id}/upload | Uploads a task document
*TaxRatesAPI* | [**BulkTaxRates**](docs/TaxRatesAPI.md#bulktaxrates) | **Post** /api/v1/tax_rates/bulk | Performs bulk actions on an array of TaxRates
*TaxRatesAPI* | [**DeleteTaxRate**](docs/TaxRatesAPI.md#deletetaxrate) | **Delete** /api/v1/tax_rates/{id} | Deletes a TaxRate
*TaxRatesAPI* | [**EditTaxRate**](docs/TaxRatesAPI.md#edittaxrate) | **Get** /api/v1/tax_rates/{id}/edit | Shows a Tax Rate for editting
*TaxRatesAPI* | [**GetTaxRateCreate**](docs/TaxRatesAPI.md#gettaxratecreate) | **Get** /api/v1/tax_rates/create | Gets a new blank Tax Rate object
*TaxRatesAPI* | [**GetTaxRates**](docs/TaxRatesAPI.md#gettaxrates) | **Get** /api/v1/tax_rates | Gets a list of tax_rates
*TaxRatesAPI* | [**ShowTaxRate**](docs/TaxRatesAPI.md#showtaxrate) | **Get** /api/v1/tax_rates/{id} | Shows a Tax Rate
*TaxRatesAPI* | [**UpdateTaxRate**](docs/TaxRatesAPI.md#updatetaxrate) | **Put** /api/v1/tax_rates/{id} | Updates a tax rate
*TemplatesAPI* | [**GetShowTemplate**](docs/TemplatesAPI.md#getshowtemplate) | **Post** /api/v1/templates | Returns a entity template with the template variables replaced with the Entities
*TokensAPI* | [**BulkTokens**](docs/TokensAPI.md#bulktokens) | **Post** /api/v1/tokens/bulk | Performs bulk actions on an array of tokens
*TokensAPI* | [**DeleteToken**](docs/TokensAPI.md#deletetoken) | **Delete** /api/v1/tokens/{id} | Deletes a token
*TokensAPI* | [**EditToken**](docs/TokensAPI.md#edittoken) | **Get** /api/v1/tokens/{id}/edit | Shows a token for editting
*TokensAPI* | [**GetTokens**](docs/TokensAPI.md#gettokens) | **Get** /api/v1/tokens | Gets a list of company tokens
*TokensAPI* | [**GetTokensCreate**](docs/TokensAPI.md#gettokenscreate) | **Get** /api/v1/tokens/create | Gets a new blank token object
*TokensAPI* | [**ShowToken**](docs/TokensAPI.md#showtoken) | **Get** /api/v1/tokens/{id} | Shows a token
*TokensAPI* | [**StoreToken**](docs/TokensAPI.md#storetoken) | **Post** /api/v1/tokens | Adds a token
*TokensAPI* | [**UpdateToken**](docs/TokensAPI.md#updatetoken) | **Put** /api/v1/tokens/{id} | Updates a token
*UpdateAPI* | [**SelfUpdate**](docs/UpdateAPI.md#selfupdate) | **Post** /api/v1/self-update | Performs a system update
*UsersAPI* | [**BulkUsers**](docs/UsersAPI.md#bulkusers) | **Post** /api/v1/users/bulk | Performs bulk actions on an array of users
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /api/v1/users/{id} | Deletes a User
*UsersAPI* | [**DetachUser**](docs/UsersAPI.md#detachuser) | **Delete** /api/v1/users/{user}/detach_from_company | Detach an existing user to a company
*UsersAPI* | [**EditUser**](docs/UsersAPI.md#edituser) | **Get** /api/v1/users/{id}/edit | Shows an User for editting
*UsersAPI* | [**GetUsers**](docs/UsersAPI.md#getusers) | **Get** /api/v1/users | Gets a list of users
*UsersAPI* | [**GetUsersCreate**](docs/UsersAPI.md#getuserscreate) | **Get** /api/v1/users/create | Gets a new blank User object
*UsersAPI* | [**InviteUser**](docs/UsersAPI.md#inviteuser) | **Post** /api/v1/users/{user}/invite | Reconfirm an existing user to a company
*UsersAPI* | [**InviteUserReconfirm**](docs/UsersAPI.md#inviteuserreconfirm) | **Post** /api/v1/users/{user}/reconfirm | Reconfirm an existing user to a company
*UsersAPI* | [**ShowUser**](docs/UsersAPI.md#showuser) | **Get** /api/v1/users/{id} | Shows an User
*UsersAPI* | [**StoreUser**](docs/UsersAPI.md#storeuser) | **Post** /api/v1/users | Adds a User
*UsersAPI* | [**UpdateUser**](docs/UsersAPI.md#updateuser) | **Put** /api/v1/users/{id} | Updates an User
*VendorsAPI* | [**BulkVendors**](docs/VendorsAPI.md#bulkvendors) | **Post** /api/v1/vendors/bulk | Bulk vendor actions
*VendorsAPI* | [**DeleteVendor**](docs/VendorsAPI.md#deletevendor) | **Delete** /api/v1/vendors/{id} | Delete vendor
*VendorsAPI* | [**EditVendor**](docs/VendorsAPI.md#editvendor) | **Get** /api/v1/vendors/{id}/edit | Edit vendor
*VendorsAPI* | [**GetVendors**](docs/VendorsAPI.md#getvendors) | **Get** /api/v1/vendors | List vendors
*VendorsAPI* | [**GetVendorsCreate**](docs/VendorsAPI.md#getvendorscreate) | **Get** /api/v1/vendors/create | Blank vendor
*VendorsAPI* | [**ShowVendor**](docs/VendorsAPI.md#showvendor) | **Get** /api/v1/vendors/{id} | Show vendor
*VendorsAPI* | [**StoreVendor**](docs/VendorsAPI.md#storevendor) | **Post** /api/v1/vendors | Create vendor
*VendorsAPI* | [**UpdateVendor**](docs/VendorsAPI.md#updatevendor) | **Put** /api/v1/vendors/{id} | Update vendor
*VendorsAPI* | [**UploadVendor**](docs/VendorsAPI.md#uploadvendor) | **Post** /api/v1/vendors/{id}/upload | Uploads a vendor document
*WebcronAPI* | [**Webcron**](docs/WebcronAPI.md#webcron) | **Get** /api/v1/webcron | Executes the task scheduler via a webcron service
*WebhooksAPI* | [**BulkWebhooks**](docs/WebhooksAPI.md#bulkwebhooks) | **Post** /api/v1/webhooks/bulk | Performs bulk actions on an array of Webhooks
*WebhooksAPI* | [**DeleteWebhook**](docs/WebhooksAPI.md#deletewebhook) | **Delete** /api/v1/webhooks/{id} | Deletes a Webhook
*WebhooksAPI* | [**EditWebhook**](docs/WebhooksAPI.md#editwebhook) | **Get** /api/v1/webhooks/{id}/edit | Shows a Webhook for editting
*WebhooksAPI* | [**GetWebhooks**](docs/WebhooksAPI.md#getwebhooks) | **Get** /api/v1/webhooks | Gets a list of Webhooks
*WebhooksAPI* | [**GetWebhooksCreate**](docs/WebhooksAPI.md#getwebhookscreate) | **Get** /api/v1/webhooks/create | Gets a new blank Webhook object
*WebhooksAPI* | [**ShowWebhook**](docs/WebhooksAPI.md#showwebhook) | **Get** /api/v1/webhooks/{id} | Shows a Webhook
*WebhooksAPI* | [**StoreWebhook**](docs/WebhooksAPI.md#storewebhook) | **Post** /api/v1/webhooks | Adds a Webhook
*WebhooksAPI* | [**UpdateWebhook**](docs/WebhooksAPI.md#updatewebhook) | **Put** /api/v1/webhooks/{id} | Updates a Webhook
*YodleeAPI* | [**YodleeRefreshWebhook**](docs/YodleeAPI.md#yodleerefreshwebhook) | **Post** /api/v1/yodlee/refresh | Yodlee Webhook


## Documentation For Models

 - [Account](docs/Account.md)
 - [Activity](docs/Activity.md)
 - [AuthenticationError](docs/AuthenticationError.md)
 - [AuthorizationError](docs/AuthorizationError.md)
 - [BTRules](docs/BTRules.md)
 - [BankIntegration](docs/BankIntegration.md)
 - [BankTransaction](docs/BankTransaction.md)
 - [BankTransactionRule](docs/BankTransactionRule.md)
 - [BulkInvoicesRequest](docs/BulkInvoicesRequest.md)
 - [BulkQuotesRequest](docs/BulkQuotesRequest.md)
 - [BulkRecurringInvoicesRequest](docs/BulkRecurringInvoicesRequest.md)
 - [Client](docs/Client.md)
 - [ClientContact](docs/ClientContact.md)
 - [ClientContactRequest](docs/ClientContactRequest.md)
 - [ClientGatewayToken](docs/ClientGatewayToken.md)
 - [ClientRequest](docs/ClientRequest.md)
 - [ClientSettings](docs/ClientSettings.md)
 - [ClientStatementRequest](docs/ClientStatementRequest.md)
 - [Company](docs/Company.md)
 - [CompanyGateway](docs/CompanyGateway.md)
 - [CompanyLedger](docs/CompanyLedger.md)
 - [CompanySettings](docs/CompanySettings.md)
 - [CompanyToken](docs/CompanyToken.md)
 - [CompanyUser](docs/CompanyUser.md)
 - [CompanyUserRef](docs/CompanyUserRef.md)
 - [Credit](docs/Credit.md)
 - [CreditPaymentable](docs/CreditPaymentable.md)
 - [Design](docs/Design.md)
 - [Document](docs/Document.md)
 - [Error](docs/Error.md)
 - [Expense](docs/Expense.md)
 - [ExpenseCategory](docs/ExpenseCategory.md)
 - [FeesAndLimits](docs/FeesAndLimits.md)
 - [FillableInvoice](docs/FillableInvoice.md)
 - [GenericBulkAction](docs/GenericBulkAction.md)
 - [GenericReportSchema](docs/GenericReportSchema.md)
 - [GetActivities200Response](docs/GetActivities200Response.md)
 - [GetBankIntegrations200Response](docs/GetBankIntegrations200Response.md)
 - [GetBankTransactionRules200Response](docs/GetBankTransactionRules200Response.md)
 - [GetBankTransactions200Response](docs/GetBankTransactions200Response.md)
 - [GetClientGatewayTokens200Response](docs/GetClientGatewayTokens200Response.md)
 - [GetCompanies200Response](docs/GetCompanies200Response.md)
 - [GetCompanyGateways200Response](docs/GetCompanyGateways200Response.md)
 - [GetCompanyLedger200Response](docs/GetCompanyLedger200Response.md)
 - [GetCredits200Response](docs/GetCredits200Response.md)
 - [GetDesigns200Response](docs/GetDesigns200Response.md)
 - [GetDocuments200Response](docs/GetDocuments200Response.md)
 - [GetExpenseCategorys200Response](docs/GetExpenseCategorys200Response.md)
 - [GetExpenses200Response](docs/GetExpenses200Response.md)
 - [GetGroupSettings200Response](docs/GetGroupSettings200Response.md)
 - [GetInvoices200Response](docs/GetInvoices200Response.md)
 - [GetLocations200Response](docs/GetLocations200Response.md)
 - [GetPaymentTerms200Response](docs/GetPaymentTerms200Response.md)
 - [GetPayments200Response](docs/GetPayments200Response.md)
 - [GetProducts200Response](docs/GetProducts200Response.md)
 - [GetProjects200Response](docs/GetProjects200Response.md)
 - [GetPurchaseOrders200Response](docs/GetPurchaseOrders200Response.md)
 - [GetQuotes200Response](docs/GetQuotes200Response.md)
 - [GetRecurringExpenses200Response](docs/GetRecurringExpenses200Response.md)
 - [GetRecurringInvoices200Response](docs/GetRecurringInvoices200Response.md)
 - [GetRecurringQuotes200Response](docs/GetRecurringQuotes200Response.md)
 - [GetShowTemplateRequest](docs/GetShowTemplateRequest.md)
 - [GetSubscriptions200Response](docs/GetSubscriptions200Response.md)
 - [GetSystemLogs200Response](docs/GetSystemLogs200Response.md)
 - [GetTaskStatuses200Response](docs/GetTaskStatuses200Response.md)
 - [GetTasks200Response](docs/GetTasks200Response.md)
 - [GetTaxRates200Response](docs/GetTaxRates200Response.md)
 - [GetTokens200Response](docs/GetTokens200Response.md)
 - [GetUsers200Response](docs/GetUsers200Response.md)
 - [GetVendors200Response](docs/GetVendors200Response.md)
 - [GetWebhooks200Response](docs/GetWebhooks200Response.md)
 - [GroupSetting](docs/GroupSetting.md)
 - [InvalidInputError](docs/InvalidInputError.md)
 - [Invoice](docs/Invoice.md)
 - [InvoiceInvitation](docs/InvoiceInvitation.md)
 - [InvoiceInvitationRequest](docs/InvoiceInvitationRequest.md)
 - [InvoiceItem](docs/InvoiceItem.md)
 - [InvoicePaymentable](docs/InvoicePaymentable.md)
 - [InvoiceRequest](docs/InvoiceRequest.md)
 - [Location](docs/Location.md)
 - [LocationRequest](docs/LocationRequest.md)
 - [Meta](docs/Meta.md)
 - [Pagination](docs/Pagination.md)
 - [Payment](docs/Payment.md)
 - [PaymentRequest](docs/PaymentRequest.md)
 - [PaymentTerm](docs/PaymentTerm.md)
 - [PaymentType](docs/PaymentType.md)
 - [Paymentable](docs/Paymentable.md)
 - [PostLoginRequest](docs/PostLoginRequest.md)
 - [Product](docs/Product.md)
 - [ProductBulkAction](docs/ProductBulkAction.md)
 - [ProductRequest](docs/ProductRequest.md)
 - [Project](docs/Project.md)
 - [PurchaseOrder](docs/PurchaseOrder.md)
 - [Quote](docs/Quote.md)
 - [RateLimiterError](docs/RateLimiterError.md)
 - [RecurringExpense](docs/RecurringExpense.md)
 - [RecurringInvoice](docs/RecurringInvoice.md)
 - [RecurringInvoiceRequest](docs/RecurringInvoiceRequest.md)
 - [RecurringQuote](docs/RecurringQuote.md)
 - [SendEmailTemplateRequest](docs/SendEmailTemplateRequest.md)
 - [Subscription](docs/Subscription.md)
 - [SystemLog](docs/SystemLog.md)
 - [Task](docs/Task.md)
 - [TaskSchedulerSchema](docs/TaskSchedulerSchema.md)
 - [TaskStatus](docs/TaskStatus.md)
 - [TaxRate](docs/TaxRate.md)
 - [Template](docs/Template.md)
 - [UpdateJobForASchedulerSchema](docs/UpdateJobForASchedulerSchema.md)
 - [UpdateTaskSchedulerSchema](docs/UpdateTaskSchedulerSchema.md)
 - [User](docs/User.md)
 - [UserRef](docs/UserRef.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorErrors](docs/ValidationErrorErrors.md)
 - [Vendor](docs/Vendor.md)
 - [VendorContact](docs/VendorContact.md)
 - [Webhook](docs/Webhook.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-TOKEN
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ApiKeyAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"ApiKeyAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

contact@invoiceninja.com

## Dev

```
go install github.com/dkorunic/betteralign/cmd/betteralign@latest
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.5
go install golang.org/x/tools/cmd/deadcode@latest

go install mvdan.cc/gofumpt@latest
```

```
golangci-lint run -c .golangci.yml ./...
betteralign -apply ./...
nilaway ./...
deadcode ./...

gofmt -r 'interface{} -> any' -w .

gofumpt -l -w .
```
