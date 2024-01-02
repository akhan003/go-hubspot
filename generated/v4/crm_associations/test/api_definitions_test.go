/*
CrmPublicAssociationsServiceV4

Testing DefinitionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package crm_associations

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_crm_associations_DefinitionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefinitionsAPIService DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fromObjectType string
		var toObjectType string
		var associationTypeId int32

		httpRes, err := apiClient.DefinitionsAPI.DeleteCrmV4AssociationsFromObjectTypeToObjectTypeLabelsAssociationTypeIdArchive(context.Background(), fromObjectType, toObjectType, associationTypeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefinitionsAPIService GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fromObjectType string
		var toObjectType string

		resp, httpRes, err := apiClient.DefinitionsAPI.GetCrmV4AssociationsFromObjectTypeToObjectTypeLabelsGetAll(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefinitionsAPIService PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fromObjectType string
		var toObjectType string

		resp, httpRes, err := apiClient.DefinitionsAPI.PostCrmV4AssociationsFromObjectTypeToObjectTypeLabelsCreate(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefinitionsAPIService PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fromObjectType string
		var toObjectType string

		httpRes, err := apiClient.DefinitionsAPI.PutCrmV4AssociationsFromObjectTypeToObjectTypeLabelsUpdate(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}