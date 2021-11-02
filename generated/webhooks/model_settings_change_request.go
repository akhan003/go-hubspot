/*
 * Webhooks API
 *
 * Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package webhooks

// New or updated webhook settings for an app.
type SettingsChangeRequest struct {
	// A publicly available URL for Hubspot to call where event payloads will be delivered. See [link-so-some-doc](#) for details about the format of these event payloads.
	TargetUrl  string              `json:"targetUrl"`
	Throttling *ThrottlingSettings `json:"throttling"`
}