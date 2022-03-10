package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DeletedServersClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type DeletedServersClient struct {
	BaseClient
}

// NewDeletedServersClient creates an instance of the DeletedServersClient client.
func NewDeletedServersClient(subscriptionID string) DeletedServersClient {
	return NewDeletedServersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeletedServersClientWithBaseURI creates an instance of the DeletedServersClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDeletedServersClientWithBaseURI(baseURI string, subscriptionID string) DeletedServersClient {
	return DeletedServersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a deleted server.
// Parameters:
// locationName - the name of the region where the resource is located.
// deletedServerName - the name of the deleted server.
func (client DeletedServersClient) Get(ctx context.Context, locationName string, deletedServerName string) (result DeletedServer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedServersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, locationName, deletedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeletedServersClient) GetPreparer(ctx context.Context, locationName string, deletedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deletedServerName": autorest.Encode("path", deletedServerName),
		"locationName":      autorest.Encode("path", locationName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers/{deletedServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedServersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeletedServersClient) GetResponder(resp *http.Response) (result DeletedServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of all deleted servers in a subscription.
func (client DeletedServersClient) List(ctx context.Context) (result DeletedServerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedServersClient.List")
		defer func() {
			sc := -1
			if result.dslr.Response.Response != nil {
				sc = result.dslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "List", resp, "Failure sending request")
		return
	}

	result.dslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dslr.hasNextLink() && result.dslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DeletedServersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/deletedServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedServersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeletedServersClient) ListResponder(resp *http.Response) (result DeletedServerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeletedServersClient) listNextResults(ctx context.Context, lastResults DeletedServerListResult) (result DeletedServerListResult, err error) {
	req, err := lastResults.deletedServerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.DeletedServersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.DeletedServersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeletedServersClient) ListComplete(ctx context.Context) (result DeletedServerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedServersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByLocation gets a list of deleted servers for a location.
// Parameters:
// locationName - the name of the region where the resource is located.
func (client DeletedServersClient) ListByLocation(ctx context.Context, locationName string) (result DeletedServerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedServersClient.ListByLocation")
		defer func() {
			sc := -1
			if result.dslr.Response.Response != nil {
				sc = result.dslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByLocationNextResults
	req, err := client.ListByLocationPreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.dslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result.dslr, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "ListByLocation", resp, "Failure responding to request")
		return
	}
	if result.dslr.hasNextLink() && result.dslr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client DeletedServersClient) ListByLocationPreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedServersClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client DeletedServersClient) ListByLocationResponder(resp *http.Response) (result DeletedServerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByLocationNextResults retrieves the next set of results, if any.
func (client DeletedServersClient) listByLocationNextResults(ctx context.Context, lastResults DeletedServerListResult) (result DeletedServerListResult, err error) {
	req, err := lastResults.deletedServerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.DeletedServersClient", "listByLocationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.DeletedServersClient", "listByLocationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "listByLocationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByLocationComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeletedServersClient) ListByLocationComplete(ctx context.Context, locationName string) (result DeletedServerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedServersClient.ListByLocation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByLocation(ctx, locationName)
	return
}

// Recover recovers a deleted server.
// Parameters:
// locationName - the name of the region where the resource is located.
// deletedServerName - the name of the deleted server.
func (client DeletedServersClient) Recover(ctx context.Context, locationName string, deletedServerName string) (result DeletedServersRecoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedServersClient.Recover")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RecoverPreparer(ctx, locationName, deletedServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "Recover", nil, "Failure preparing request")
		return
	}

	result, err = client.RecoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DeletedServersClient", "Recover", result.Response(), "Failure sending request")
		return
	}

	return
}

// RecoverPreparer prepares the Recover request.
func (client DeletedServersClient) RecoverPreparer(ctx context.Context, locationName string, deletedServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deletedServerName": autorest.Encode("path", deletedServerName),
		"locationName":      autorest.Encode("path", locationName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/deletedServers/{deletedServerName}/recover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RecoverSender sends the Recover request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedServersClient) RecoverSender(req *http.Request) (future DeletedServersRecoverFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RecoverResponder handles the response to the Recover request. The method always
// closes the http.Response Body.
func (client DeletedServersClient) RecoverResponder(resp *http.Response) (result DeletedServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
