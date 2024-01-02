/*
CRM Timeline

Testing TemplatesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package timeline

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_timeline_TemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TemplatesAPIService DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var eventTemplateId string
		var appId int32

		httpRes, err := apiClient.TemplatesAPI.DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId(context.Background(), eventTemplateId, appId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService GetCrmV3TimelineAppIdEventTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.TemplatesAPI.GetCrmV3TimelineAppIdEventTemplates(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService GetCrmV3TimelineAppIdEventTemplatesEventTemplateId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var eventTemplateId string
		var appId int32

		resp, httpRes, err := apiClient.TemplatesAPI.GetCrmV3TimelineAppIdEventTemplatesEventTemplateId(context.Background(), eventTemplateId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService PostCrmV3TimelineAppIdEventTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.TemplatesAPI.PostCrmV3TimelineAppIdEventTemplates(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService PutCrmV3TimelineAppIdEventTemplatesEventTemplateId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var eventTemplateId string
		var appId int32

		resp, httpRes, err := apiClient.TemplatesAPI.PutCrmV3TimelineAppIdEventTemplatesEventTemplateId(context.Background(), eventTemplateId, appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}