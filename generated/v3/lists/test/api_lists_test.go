/*
Lists

Testing ListsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package lists

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lists_ListsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ListsAPIService DeleteCrmV3ListsListIdRemove", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		httpRes, err := apiClient.ListsAPI.DeleteCrmV3ListsListIdRemove(context.Background(), listId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService GetCrmV3ListsGetAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ListsAPI.GetCrmV3ListsGetAll(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService GetCrmV3ListsListIdGetById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.ListsAPI.GetCrmV3ListsListIdGetById(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService GetCrmV3ListsObjectTypeIdObjectTypeIdNameListNameGetByName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listName string
		var objectTypeId string

		resp, httpRes, err := apiClient.ListsAPI.GetCrmV3ListsObjectTypeIdObjectTypeIdNameListNameGetByName(context.Background(), listName, objectTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService PostCrmV3ListsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ListsAPI.PostCrmV3ListsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService PostCrmV3ListsSearchDoSearch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ListsAPI.PostCrmV3ListsSearchDoSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService PutCrmV3ListsListIdRestoreRestore", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		httpRes, err := apiClient.ListsAPI.PutCrmV3ListsListIdRestoreRestore(context.Background(), listId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService PutCrmV3ListsListIdUpdateListFiltersUpdateListFilters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.ListsAPI.PutCrmV3ListsListIdUpdateListFiltersUpdateListFilters(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService UpdateName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId int32

		resp, httpRes, err := apiClient.ListsAPI.UpdateName(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
