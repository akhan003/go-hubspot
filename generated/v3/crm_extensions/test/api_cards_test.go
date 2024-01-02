/*
CRM cards

Testing CardsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package crm_extensions

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_crm_extensions_CardsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CardsAPIService CardsArchive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var cardId string

		httpRes, err := apiClient.CardsAPI.CardsArchive(context.Background(), appId, cardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsAPIService CardsCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.CardsAPI.CardsCreate(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsAPIService CardsGetAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32

		resp, httpRes, err := apiClient.CardsAPI.CardsGetAll(context.Background(), appId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsAPIService CardsGetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var cardId string

		resp, httpRes, err := apiClient.CardsAPI.CardsGetByID(context.Background(), appId, cardId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CardsAPIService CardsUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var appId int32
		var cardId string

		resp, httpRes, err := apiClient.CardsAPI.CardsUpdate(context.Background(), appId, cardId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
