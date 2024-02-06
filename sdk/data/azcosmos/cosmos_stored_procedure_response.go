// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"net/http"

	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// StoredProcedureResponse represents the response from the execution of a stored procedure.
type StoredProcedureResponse struct {
	// The byte content of the operation response.
	Value []byte

	// TODO: change this remove etag and include other interesting stored proc response properties.
	Response

	// SessionToken contains the value from the session token header to be used on session consistency.
	SessionToken string
}

func newStoredProcedureResponse(resp *http.Response) (StoredProcedureResponse, error) {
	response := StoredProcedureResponse{
		Response: newResponse(resp),
	}
	response.SessionToken = resp.Header.Get(cosmosHeaderSessionToken)
	defer resp.Body.Close()
	body, err := azruntime.Payload(resp)
	if err != nil {
		return response, err
	}
	response.Value = body
	return response, nil
}
