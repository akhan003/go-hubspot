/*
 * Line Items
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package line_items

type Filter struct {
	Value        string `json:"value,omitempty"`
	PropertyName string `json:"propertyName"`
	// null
	Operator string `json:"operator"`
}