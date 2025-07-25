/*
GCOM API

Testing InstancesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gcom

import (
	"context"
	"testing"

	openapiclient "github.com/grafana/grafana-com-public-clients/go/gcom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_gcom_InstancesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InstancesAPIService DeleteInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.DeleteInstance(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService DeleteInstancePlugin", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var pluginSlugOrId string

		resp, httpRes, err := apiClient.InstancesAPI.DeleteInstancePlugin(context.Background(), instanceId, pluginSlugOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService DeleteInstanceServiceAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var serviceAccountId string

		httpRes, err := apiClient.InstancesAPI.DeleteInstanceServiceAccount(context.Background(), instanceId, serviceAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService DeleteInstanceServiceAccountToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var serviceAccountId string
		var tokenId string

		httpRes, err := apiClient.InstancesAPI.DeleteInstanceServiceAccountToken(context.Background(), instanceId, serviceAccountId, tokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetConnections", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.GetConnections(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.GetInstance(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetInstancePlugin", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var pluginSlugOrId string

		resp, httpRes, err := apiClient.InstancesAPI.GetInstancePlugin(context.Background(), instanceId, pluginSlugOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetInstancePlugins", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.GetInstancePlugins(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetInstanceServiceAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var serviceAccountId string

		resp, httpRes, err := apiClient.InstancesAPI.GetInstanceServiceAccount(context.Background(), instanceId, serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetInstanceServiceAccountTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var serviceAccountId string

		resp, httpRes, err := apiClient.InstancesAPI.GetInstanceServiceAccountTokens(context.Background(), instanceId, serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetInstanceUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.GetInstanceUsers(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService GetInstances", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InstancesAPI.GetInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService PostInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.PostInstance(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService PostInstancePlugin", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var pluginSlugOrId string

		resp, httpRes, err := apiClient.InstancesAPI.PostInstancePlugin(context.Background(), instanceId, pluginSlugOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService PostInstancePlugins", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.PostInstancePlugins(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService PostInstanceServiceAccountTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string
		var serviceAccountId string

		resp, httpRes, err := apiClient.InstancesAPI.PostInstanceServiceAccountTokens(context.Background(), instanceId, serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService PostInstanceServiceAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var instanceId string

		resp, httpRes, err := apiClient.InstancesAPI.PostInstanceServiceAccounts(context.Background(), instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstancesAPIService PostInstances", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InstancesAPI.PostInstances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
