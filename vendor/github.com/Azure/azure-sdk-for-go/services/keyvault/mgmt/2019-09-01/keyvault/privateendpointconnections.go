package keyvault

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrivateEndpointConnectionsClient is the the Azure management API provides a RESTful set of web services that
// interact with Azure Key Vault.
type PrivateEndpointConnectionsClient struct {
	BaseClient
}

// NewPrivateEndpointConnectionsClient creates an instance of the PrivateEndpointConnectionsClient client.
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return NewPrivateEndpointConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionsClientWithBaseURI creates an instance of the PrivateEndpointConnectionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return PrivateEndpointConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes the specified private endpoint connection associated with the key vault.
// Parameters:
// resourceGroupName - name of the resource group that contains the key vault.
// vaultName - the name of the key vault.
// privateEndpointConnectionName - name of the private endpoint connection associated with the key vault.
func (client PrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (result PrivateEndpointConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: vaultName,
			Constraints: []validation.Constraint{{Target: "vaultName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("keyvault.PrivateEndpointConnectionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, vaultName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateEndpointConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
		"vaultName":                     autorest.Encode("path", vaultName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) DeleteSender(req *http.Request) (future PrivateEndpointConnectionsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = func(client PrivateEndpointConnectionsClient) (pec PrivateEndpointConnection, err error) {
		var done bool
		done, err = future.DoneWithContext(context.Background(), client)
		if err != nil {
			err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsDeleteFuture", "Result", future.Response(), "Polling failure")
			return
		}
		if !done {
			err = azure.NewAsyncOpIncompleteError("keyvault.PrivateEndpointConnectionsDeleteFuture")
			return
		}
		sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
		pec.Response.Response, err = future.GetResult(sender)
		if pec.Response.Response == nil && err == nil {
			err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsDeleteFuture", "Result", nil, "received nil response and error")
		}
		if err == nil && pec.Response.Response.StatusCode != http.StatusNoContent {
			pec, err = client.DeleteResponder(pec.Response.Response)
			if err != nil {
				err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsDeleteFuture", "Result", pec.Response.Response, "Failure responding to request")
			}
		}
		return
	}
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) DeleteResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the specified private endpoint connection associated with the key vault.
// Parameters:
// resourceGroupName - name of the resource group that contains the key vault.
// vaultName - the name of the key vault.
// privateEndpointConnectionName - name of the private endpoint connection associated with the key vault.
func (client PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: vaultName,
			Constraints: []validation.Constraint{{Target: "vaultName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("keyvault.PrivateEndpointConnectionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, vaultName, privateEndpointConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateEndpointConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
		"vaultName":                     autorest.Encode("path", vaultName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) GetResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put updates the specified private endpoint connection associated with the key vault.
// Parameters:
// resourceGroupName - name of the resource group that contains the key vault.
// vaultName - the name of the key vault.
// privateEndpointConnectionName - name of the private endpoint connection associated with the key vault.
// properties - the intended state of private endpoint connection.
func (client PrivateEndpointConnectionsClient) Put(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (result PrivateEndpointConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointConnectionsClient.Put")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: vaultName,
			Constraints: []validation.Constraint{{Target: "vaultName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("keyvault.PrivateEndpointConnectionsClient", "Put", err.Error())
	}

	req, err := client.PutPreparer(ctx, resourceGroupName, vaultName, privateEndpointConnectionName, properties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.PrivateEndpointConnectionsClient", "Put", resp, "Failure responding to request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client PrivateEndpointConnectionsClient) PutPreparer(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties PrivateEndpointConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointConnectionName": autorest.Encode("path", privateEndpointConnectionName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
		"vaultName":                     autorest.Encode("path", vaultName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}", pathParameters),
		autorest.WithJSON(properties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointConnectionsClient) PutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client PrivateEndpointConnectionsClient) PutResponder(resp *http.Response) (result PrivateEndpointConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
