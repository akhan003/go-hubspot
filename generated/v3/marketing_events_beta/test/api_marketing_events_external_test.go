/*
Marketing Marketing Events

Testing MarketingEventsExternalAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package marketing_events_beta

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_marketing_events_beta_MarketingEventsExternalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MarketingEventsExternalAPIService ExternalComplete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var externalEventId string

		resp, httpRes, err := apiClient.MarketingEventsExternalAPI.ExternalComplete(context.Background(), externalEventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
