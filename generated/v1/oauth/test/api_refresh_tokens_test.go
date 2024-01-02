/*

Testing RefreshTokensAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package oauth

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_oauth_RefreshTokensAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RefreshTokensAPIService DeleteOauthV1RefreshTokensTokenArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		httpRes, err := apiClient.RefreshTokensAPI.DeleteOauthV1RefreshTokensTokenArchive(context.Background(), token).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RefreshTokensAPIService GetOauthV1RefreshTokensTokenGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.RefreshTokensAPI.GetOauthV1RefreshTokensTokenGet(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
