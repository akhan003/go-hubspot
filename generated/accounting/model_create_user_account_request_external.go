/*
 * Accounting Extension
 *
 * These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package accounting

// Information about the account in your external account system.
type CreateUserAccountRequestExternal struct {
	// The id of the account in your system.
	AccountId string `json:"accountId"`
	// The name of the account in your system. This is normally the name visible to your users.
	AccountName string `json:"accountName"`
	// The default currency that this account uses.
	CurrencyCode string `json:"currencyCode"`
}