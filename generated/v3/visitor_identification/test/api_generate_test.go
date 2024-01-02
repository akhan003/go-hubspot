/*
Conversations Visitor Identification

Testing GenerateAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package visitor_identification

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_visitor_identification_GenerateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GenerateAPIService PostConversationsV3VisitorIdentificationTokensCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GenerateAPI.PostConversationsV3VisitorIdentificationTokensCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
