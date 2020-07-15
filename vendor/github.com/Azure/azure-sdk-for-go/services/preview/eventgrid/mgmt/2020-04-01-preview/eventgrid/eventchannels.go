package eventgrid

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EventChannelsClient is the azure EventGrid Management Client
type EventChannelsClient struct {
	BaseClient
}

// NewEventChannelsClient creates an instance of the EventChannelsClient client.
func NewEventChannelsClient(subscriptionID string) EventChannelsClient {
	return NewEventChannelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEventChannelsClientWithBaseURI creates an instance of the EventChannelsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEventChannelsClientWithBaseURI(baseURI string, subscriptionID string) EventChannelsClient {
	return EventChannelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate asynchronously creates a new event channel with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerNamespaceName - name of the partner namespace.
// eventChannelName - name of the event channel.
// eventChannelInfo - eventChannel information.
func (client EventChannelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, eventChannelInfo EventChannel) (result EventChannel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventChannelsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, partnerNamespaceName, eventChannelName, eventChannelInfo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EventChannelsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string, eventChannelInfo EventChannel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventChannelName":     autorest.Encode("path", eventChannelName),
		"partnerNamespaceName": autorest.Encode("path", partnerNamespaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels/{eventChannelName}", pathParameters),
		autorest.WithJSON(eventChannelInfo),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EventChannelsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EventChannelsClient) CreateOrUpdateResponder(resp *http.Response) (result EventChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete existing event channel.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerNamespaceName - name of the partner namespace.
// eventChannelName - name of the event channel.
func (client EventChannelsClient) Delete(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string) (result EventChannelsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventChannelsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, partnerNamespaceName, eventChannelName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EventChannelsClient) DeletePreparer(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventChannelName":     autorest.Encode("path", eventChannelName),
		"partnerNamespaceName": autorest.Encode("path", partnerNamespaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels/{eventChannelName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EventChannelsClient) DeleteSender(req *http.Request) (future EventChannelsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EventChannelsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get properties of an event channel.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerNamespaceName - name of the partner namespace.
// eventChannelName - name of the event channel.
func (client EventChannelsClient) Get(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string) (result EventChannel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventChannelsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, partnerNamespaceName, eventChannelName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EventChannelsClient) GetPreparer(ctx context.Context, resourceGroupName string, partnerNamespaceName string, eventChannelName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventChannelName":     autorest.Encode("path", eventChannelName),
		"partnerNamespaceName": autorest.Encode("path", partnerNamespaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels/{eventChannelName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EventChannelsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EventChannelsClient) GetResponder(resp *http.Response) (result EventChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPartnerNamespace list all the event channels in a partner namespace.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// partnerNamespaceName - name of the partner namespace.
// filter - the query used to filter the search results using OData syntax. Filtering is permitted on the
// 'name' property only and with limited number of OData operations. These operations are: the 'contains'
// function as well as the following logical operations: not, and, or, eq (for equal), and ne (for not equal).
// No arithmetic operations are supported. The following is a valid filter example: $filter=contains(namE,
// 'PATTERN') and name ne 'PATTERN-1'. The following is not a valid filter example: $filter=location eq
// 'westus'.
// top - the number of results to return per page for the list operation. Valid range for top parameter is 1 to
// 100. If not specified, the default number of results to be returned is 20 items per page.
func (client EventChannelsClient) ListByPartnerNamespace(ctx context.Context, resourceGroupName string, partnerNamespaceName string, filter string, top *int32) (result EventChannelsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventChannelsClient.ListByPartnerNamespace")
		defer func() {
			sc := -1
			if result.eclr.Response.Response != nil {
				sc = result.eclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPartnerNamespaceNextResults
	req, err := client.ListByPartnerNamespacePreparer(ctx, resourceGroupName, partnerNamespaceName, filter, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "ListByPartnerNamespace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPartnerNamespaceSender(req)
	if err != nil {
		result.eclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "ListByPartnerNamespace", resp, "Failure sending request")
		return
	}

	result.eclr, err = client.ListByPartnerNamespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "ListByPartnerNamespace", resp, "Failure responding to request")
	}

	return
}

// ListByPartnerNamespacePreparer prepares the ListByPartnerNamespace request.
func (client EventChannelsClient) ListByPartnerNamespacePreparer(ctx context.Context, resourceGroupName string, partnerNamespaceName string, filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"partnerNamespaceName": autorest.Encode("path", partnerNamespaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-04-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/eventChannels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPartnerNamespaceSender sends the ListByPartnerNamespace request. The method will close the
// http.Response Body if it receives an error.
func (client EventChannelsClient) ListByPartnerNamespaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPartnerNamespaceResponder handles the response to the ListByPartnerNamespace request. The method always
// closes the http.Response Body.
func (client EventChannelsClient) ListByPartnerNamespaceResponder(resp *http.Response) (result EventChannelsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPartnerNamespaceNextResults retrieves the next set of results, if any.
func (client EventChannelsClient) listByPartnerNamespaceNextResults(ctx context.Context, lastResults EventChannelsListResult) (result EventChannelsListResult, err error) {
	req, err := lastResults.eventChannelsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "listByPartnerNamespaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPartnerNamespaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "listByPartnerNamespaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPartnerNamespaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "eventgrid.EventChannelsClient", "listByPartnerNamespaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPartnerNamespaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client EventChannelsClient) ListByPartnerNamespaceComplete(ctx context.Context, resourceGroupName string, partnerNamespaceName string, filter string, top *int32) (result EventChannelsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventChannelsClient.ListByPartnerNamespace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPartnerNamespace(ctx, resourceGroupName, partnerNamespaceName, filter, top)
	return
}
