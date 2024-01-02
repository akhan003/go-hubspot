/*
Products

Testing BasicAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package products

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_products_BasicAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BasicAPIService Archive", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var productId string

		httpRes, err := apiClient.BasicAPI.Archive(context.Background(), productId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService Create", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BasicAPI.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService GetByID", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var productId string

		resp, httpRes, err := apiClient.BasicAPI.GetByID(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService GetPage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.BasicAPI.GetPage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BasicAPIService Update", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var productId string

		resp, httpRes, err := apiClient.BasicAPI.Update(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
