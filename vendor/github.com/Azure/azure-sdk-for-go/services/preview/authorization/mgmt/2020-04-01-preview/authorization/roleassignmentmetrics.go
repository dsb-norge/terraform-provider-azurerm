package authorization

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

// RoleAssignmentMetricsClient is the client for the RoleAssignmentMetrics methods of the Authorization service.
type RoleAssignmentMetricsClient struct {
	BaseClient
}

// NewRoleAssignmentMetricsClient creates an instance of the RoleAssignmentMetricsClient client.
func NewRoleAssignmentMetricsClient(subscriptionID string) RoleAssignmentMetricsClient {
	return NewRoleAssignmentMetricsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleAssignmentMetricsClientWithBaseURI creates an instance of the RoleAssignmentMetricsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewRoleAssignmentMetricsClientWithBaseURI(baseURI string, subscriptionID string) RoleAssignmentMetricsClient {
	return RoleAssignmentMetricsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetMetricsForSubscription get role assignment usage metrics for a subscription
func (client RoleAssignmentMetricsClient) GetMetricsForSubscription(ctx context.Context) (result RoleAssignmentMetricsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleAssignmentMetricsClient.GetMetricsForSubscription")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.RoleAssignmentMetricsClient", "GetMetricsForSubscription", err.Error())
	}

	req, err := client.GetMetricsForSubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentMetricsClient", "GetMetricsForSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetMetricsForSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentMetricsClient", "GetMetricsForSubscription", resp, "Failure sending request")
		return
	}

	result, err = client.GetMetricsForSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleAssignmentMetricsClient", "GetMetricsForSubscription", resp, "Failure responding to request")
		return
	}

	return
}

// GetMetricsForSubscriptionPreparer prepares the GetMetricsForSubscription request.
func (client RoleAssignmentMetricsClient) GetMetricsForSubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignmentsUsageMetrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetMetricsForSubscriptionSender sends the GetMetricsForSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client RoleAssignmentMetricsClient) GetMetricsForSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetMetricsForSubscriptionResponder handles the response to the GetMetricsForSubscription request. The method always
// closes the http.Response Body.
func (client RoleAssignmentMetricsClient) GetMetricsForSubscriptionResponder(resp *http.Response) (result RoleAssignmentMetricsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
