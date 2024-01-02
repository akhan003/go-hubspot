/*
Posts

Testing BlogPostsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package blog_posts

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_blog_posts_BlogPostsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlogPostsAPIService DeleteCmsV3BlogsPostsObjectId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		httpRes, err := apiClient.BlogPostsAPI.DeleteCmsV3BlogsPostsObjectId(context.Background(), objectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService GetCmsV3BlogsPosts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPosts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService GetCmsV3BlogsPostsObjectId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectId(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService GetCmsV3BlogsPostsObjectIdDraft", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectIdDraft(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService GetCmsV3BlogsPostsObjectIdRevisions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisions(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService GetCmsV3BlogsPostsObjectIdRevisionsRevisionId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string
		var revisionId string

		resp, httpRes, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisionsRevisionId(context.Background(), objectId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PatchCmsV3BlogsPostsObjectId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsAPI.PatchCmsV3BlogsPostsObjectId(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PatchCmsV3BlogsPostsObjectIdDraft", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		resp, httpRes, err := apiClient.BlogPostsAPI.PatchCmsV3BlogsPostsObjectIdDraft(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPosts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPosts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsBatchArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchArchive(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsBatchCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsBatchRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchRead(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsBatchUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchUpdate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsClone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsClone(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsMultiLanguageUpdateLanguages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageUpdateLanguages(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsObjectIdDraftPushLive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdDraftPushLive(context.Background(), objectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsObjectIdDraftReset", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string

		httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdDraftReset(context.Background(), objectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string
		var revisionId string

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore(context.Background(), objectId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var objectId string
		var revisionId int64

		resp, httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft(context.Background(), objectId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PostCmsV3BlogsPostsSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsSchedule(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlogPostsAPIService PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.BlogPostsAPI.PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}