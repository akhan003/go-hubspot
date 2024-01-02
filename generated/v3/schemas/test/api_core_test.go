/*
Schemas

Testing CoreAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package schemas

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_schemas_CoreAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CoreAPIService DeleteCrmV3SchemasObjectType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string

		httpRes, err := apiClient.CoreAPI.DeleteCrmV3SchemasObjectType(context.Background(), objectType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreAPIService DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string
		var associationIdentifier string

		httpRes, err := apiClient.CoreAPI.DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier(context.Background(), objectType, associationIdentifier).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreAPIService GetCrmV3Schemas", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CoreAPI.GetCrmV3Schemas(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreAPIService GetCrmV3SchemasObjectType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.CoreAPI.GetCrmV3SchemasObjectType(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreAPIService PatchCrmV3SchemasObjectType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.CoreAPI.PatchCrmV3SchemasObjectType(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreAPIService PostCrmV3Schemas", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CoreAPI.PostCrmV3Schemas(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreAPIService PostCrmV3SchemasObjectTypeAssociations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectType string

		resp, httpRes, err := apiClient.CoreAPI.PostCrmV3SchemasObjectTypeAssociations(context.Background(), objectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
