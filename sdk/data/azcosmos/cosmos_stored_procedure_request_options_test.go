// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"testing"
)

func TestStoredProcedureRequestOptionsToHeaders(t *testing.T) {
	options := &StoredProcedureRequestOptions{}
	options.ConsistencyLevel = ConsistencyLevelSession.ToPtr()
	options.SessionToken = "sessionToken"
	options.EnableScriptLogging = true
	header := options.toHeaders()
	if header == nil {
		t.Fatal("toHeaders should return non-nil")
	}

	headers := *header
	if headers[cosmosHeaderConsistencyLevel] != "Session" {
		t.Errorf("ConsistencyLevel should be Session but got %v", headers[cosmosHeaderConsistencyLevel])
	}
	if headers[cosmosHeaderSessionToken] != "sessionToken" {
		t.Errorf("SessionToken should be sessionToken but got %v", headers[cosmosHeaderSessionToken])
	}
	if headers[cosmosHeaderScriptEnableLogging] != "true" {
		t.Errorf("ScriptEnableLogging should be true but got %v", headers[cosmosHeaderScriptEnableLogging])
	}
}
