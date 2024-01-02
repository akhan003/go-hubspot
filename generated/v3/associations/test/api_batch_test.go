/*
Associations

Testing BatchAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package associations

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_associations_BatchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BatchAPIService BatchArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fromObjectType string
		var toObjectType string

		httpRes, err := apiClient.BatchAPI.BatchArchive(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService BatchCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fromObjectType string
		var toObjectType string

		resp, httpRes, err := apiClient.BatchAPI.BatchCreate(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BatchAPIService BatchRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fromObjectType string
		var toObjectType string

		resp, httpRes, err := apiClient.BatchAPI.BatchRead(context.Background(), fromObjectType, toObjectType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
