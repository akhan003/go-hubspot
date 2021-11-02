/*
 * CRM Imports
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package imports

// An object detailing a summary of the import record outputs
type PublicImportMetadata struct {
	// The lists containing the imported objects.
	ObjectLists []PublicObjectListRecord `json:"objectLists"`
	// Summarized outcomes of each row a developer attempted to import into HubSpot.
	Counters map[string]int32 `json:"counters"`
	// The IDs of files uploaded in the File Manager API.
	FileIds []string `json:"fileIds"`
}