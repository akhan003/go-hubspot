/*
Feedback Submissions

Testing BasicAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package feedback_submissions

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_feedback_submissions_BasicAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BasicAPIService DeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId string

		httpRes, err := apiClient.BasicAPI.DeleteCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdArchive(context.Background(), feedbackSubmissionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService GetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdGetById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId string

		resp, httpRes, err := apiClient.BasicAPI.GetCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdGetById(context.Background(), feedbackSubmissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService GetCrmV3ObjectsFeedbackSubmissionsGetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BasicAPI.GetCrmV3ObjectsFeedbackSubmissionsGetPage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService PatchCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var feedbackSubmissionId string

		resp, httpRes, err := apiClient.BasicAPI.PatchCrmV3ObjectsFeedbackSubmissionsFeedbackSubmissionIdUpdate(context.Background(), feedbackSubmissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService PostCrmV3ObjectsFeedbackSubmissionsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BasicAPI.PostCrmV3ObjectsFeedbackSubmissionsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}