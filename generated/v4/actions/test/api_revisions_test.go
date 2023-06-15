/*
Custom Workflow Actions

Testing RevisionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package actions

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_actions_RevisionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RevisionsApiService RevisionsGetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var revisionId string
		var appId int32

		resp, httpRes, err := apiClient.RevisionsApi.RevisionsGetByID(context.Background(), definitionId, revisionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RevisionsApiService RevisionsGetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var definitionId string
		var appId int32

		resp, httpRes, err := apiClient.RevisionsApi.RevisionsGetPage(context.Background(), definitionId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}