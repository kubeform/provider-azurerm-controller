package streamanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// FunctionsClient is the stream Analytics Client
type FunctionsClient struct {
	BaseClient
}

// NewFunctionsClient creates an instance of the FunctionsClient client.
func NewFunctionsClient(subscriptionID string) FunctionsClient {
	return NewFunctionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFunctionsClientWithBaseURI creates an instance of the FunctionsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewFunctionsClientWithBaseURI(baseURI string, subscriptionID string) FunctionsClient {
	return FunctionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrReplace creates a function or replaces an already existing function under an existing streaming job.
// Parameters:
// function - the definition of the function that will be used to create a new function or replace the existing
// one under the streaming job.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// ifMatch - the ETag of the function. Omit this value to always overwrite the current function. Specify the
// last-seen ETag value to prevent accidentally overwriting concurrent changes.
// ifNoneMatch - set to '*' to allow a new function to be created, but to prevent updating an existing
// function. Other values will result in a 412 Pre-condition Failed response.
func (client FunctionsClient) CreateOrReplace(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string, ifNoneMatch string) (result Function, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.CreateOrReplace")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.FunctionsClient", "CreateOrReplace", err.Error())
	}

	req, err := client.CreateOrReplacePreparer(ctx, function, resourceGroupName, jobName, functionName, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrReplaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrReplaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "CreateOrReplace", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrReplacePreparer prepares the CreateOrReplace request.
func (client FunctionsClient) CreateOrReplacePreparer(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithJSON(function),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrReplaceSender sends the CreateOrReplace request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) CreateOrReplaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrReplaceResponder handles the response to the CreateOrReplace request. The method always
// closes the http.Response Body.
func (client FunctionsClient) CreateOrReplaceResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a function from the streaming job.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// functionName - the name of the function.
func (client FunctionsClient) Delete(ctx context.Context, resourceGroupName string, jobName string, functionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.FunctionsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, jobName, functionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FunctionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FunctionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets details about the specified function.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// functionName - the name of the function.
func (client FunctionsClient) Get(ctx context.Context, resourceGroupName string, jobName string, functionName string) (result Function, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.FunctionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, jobName, functionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client FunctionsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FunctionsClient) GetResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStreamingJob lists all of the functions under the specified streaming job.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// selectParameter - the $select OData query parameter. This is a comma-separated list of structural properties
// to include in the response, or "*" to include all properties. By default, all properties are returned except
// diagnostics. Currently only accepts '*' as a valid value.
func (client FunctionsClient) ListByStreamingJob(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result FunctionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.ListByStreamingJob")
		defer func() {
			sc := -1
			if result.flr.Response.Response != nil {
				sc = result.flr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.FunctionsClient", "ListByStreamingJob", err.Error())
	}

	result.fn = client.listByStreamingJobNextResults
	req, err := client.ListByStreamingJobPreparer(ctx, resourceGroupName, jobName, selectParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.flr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure sending request")
		return
	}

	result.flr, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "ListByStreamingJob", resp, "Failure responding to request")
		return
	}
	if result.flr.hasNextLink() && result.flr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByStreamingJobPreparer prepares the ListByStreamingJob request.
func (client FunctionsClient) ListByStreamingJobPreparer(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByStreamingJobSender sends the ListByStreamingJob request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) ListByStreamingJobSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByStreamingJobResponder handles the response to the ListByStreamingJob request. The method always
// closes the http.Response Body.
func (client FunctionsClient) ListByStreamingJobResponder(resp *http.Response) (result FunctionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByStreamingJobNextResults retrieves the next set of results, if any.
func (client FunctionsClient) listByStreamingJobNextResults(ctx context.Context, lastResults FunctionListResult) (result FunctionListResult, err error) {
	req, err := lastResults.functionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "listByStreamingJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByStreamingJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "listByStreamingJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByStreamingJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "listByStreamingJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByStreamingJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client FunctionsClient) ListByStreamingJobComplete(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result FunctionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.ListByStreamingJob")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByStreamingJob(ctx, resourceGroupName, jobName, selectParameter)
	return
}

// RetrieveDefaultDefinition retrieves the default definition of a function based on the parameters specified.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// functionRetrieveDefaultDefinitionParameters - parameters used to specify the type of function to retrieve
// the default definition for.
func (client FunctionsClient) RetrieveDefaultDefinition(ctx context.Context, resourceGroupName string, jobName string, functionName string, functionRetrieveDefaultDefinitionParameters *BasicFunctionRetrieveDefaultDefinitionParameters) (result Function, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.RetrieveDefaultDefinition")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", err.Error())
	}

	req, err := client.RetrieveDefaultDefinitionPreparer(ctx, resourceGroupName, jobName, functionName, functionRetrieveDefaultDefinitionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", nil, "Failure preparing request")
		return
	}

	resp, err := client.RetrieveDefaultDefinitionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", resp, "Failure sending request")
		return
	}

	result, err = client.RetrieveDefaultDefinitionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "RetrieveDefaultDefinition", resp, "Failure responding to request")
		return
	}

	return
}

// RetrieveDefaultDefinitionPreparer prepares the RetrieveDefaultDefinition request.
func (client FunctionsClient) RetrieveDefaultDefinitionPreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string, functionRetrieveDefaultDefinitionParameters *BasicFunctionRetrieveDefaultDefinitionParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}/RetrieveDefaultDefinition", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if functionRetrieveDefaultDefinitionParameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(functionRetrieveDefaultDefinitionParameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RetrieveDefaultDefinitionSender sends the RetrieveDefaultDefinition request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) RetrieveDefaultDefinitionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RetrieveDefaultDefinitionResponder handles the response to the RetrieveDefaultDefinition request. The method always
// closes the http.Response Body.
func (client FunctionsClient) RetrieveDefaultDefinitionResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Test tests if the information provided for a function is valid. This can range from testing the connection to the
// underlying web service behind the function or making sure the function code provided is syntactically correct.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// function - if the function specified does not already exist, this parameter must contain the full function
// definition intended to be tested. If the function specified already exists, this parameter can be left null
// to test the existing function as is or if specified, the properties specified will overwrite the
// corresponding properties in the existing function (exactly like a PATCH operation) and the resulting
// function will be tested.
func (client FunctionsClient) Test(ctx context.Context, resourceGroupName string, jobName string, functionName string, function *Function) (result FunctionsTestFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.Test")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.FunctionsClient", "Test", err.Error())
	}

	req, err := client.TestPreparer(ctx, resourceGroupName, jobName, functionName, function)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Test", nil, "Failure preparing request")
		return
	}

	result, err = client.TestSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Test", result.Response(), "Failure sending request")
		return
	}

	return
}

// TestPreparer prepares the Test request.
func (client FunctionsClient) TestPreparer(ctx context.Context, resourceGroupName string, jobName string, functionName string, function *Function) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}/test", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if function != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(function))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TestSender sends the Test request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) TestSender(req *http.Request) (future FunctionsTestFuture, err error) {
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

// TestResponder handles the response to the Test request. The method always
// closes the http.Response Body.
func (client FunctionsClient) TestResponder(resp *http.Response) (result ResourceTestStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing function under an existing streaming job. This can be used to partially update (ie.
// update one or two properties) a function without affecting the rest the job or function definition.
// Parameters:
// function - a function object. The properties specified here will overwrite the corresponding properties in
// the existing function (ie. Those properties will be updated). Any properties that are set to null here will
// mean that the corresponding property in the existing function will remain the same and not change as a
// result of this PATCH operation.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// jobName - the name of the streaming job.
// functionName - the name of the function.
// ifMatch - the ETag of the function. Omit this value to always overwrite the current function. Specify the
// last-seen ETag value to prevent accidentally overwriting concurrent changes.
func (client FunctionsClient) Update(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string) (result Function, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FunctionsClient.Update")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("streamanalytics.FunctionsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, function, resourceGroupName, jobName, functionName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "streamanalytics.FunctionsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FunctionsClient) UpdatePreparer(ctx context.Context, function Function, resourceGroupName string, jobName string, functionName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"functionName":      autorest.Encode("path", functionName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", pathParameters),
		autorest.WithJSON(function),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FunctionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FunctionsClient) UpdateResponder(resp *http.Response) (result Function, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
