/*
 * Blog Post endpoints
 *
 * \"Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags\"
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package authors

type CollectionResponseWithTotalBlogAuthorForwardPaging struct {
	Total   int32          `json:"total"`
	Results []BlogAuthor   `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}