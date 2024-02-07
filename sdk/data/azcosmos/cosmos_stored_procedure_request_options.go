// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

// StoredProcedureRequestOptions includes options for executing stored procedures.
type StoredProcedureRequestOptions struct {
	// SessionToken to be used when using Session consistency on the account.
	// When working with Session consistency, each new write request to Azure Cosmos DB is assigned a new SessionToken.
	// The client instance will use this token internally with each read/query request to ensure that the set consistency level is maintained.
	// In some scenarios you need to manage this Session yourself: Consider a web application with multiple nodes, each node will have its own client instance.
	// If you wanted these nodes to participate in the same session (to be able read your own writes consistently across web tiers),
	// you would have to send the SessionToken from the response of the write action on one node to the client tier, using a cookie or some other mechanism, and have that token flow back to the web tier for subsequent reads.
	// If you are using a round-robin load balancer which does not maintain session affinity between requests, such as the Azure Load Balancer,the read could potentially land on a different node to the write request, where the session was created.
	SessionToken string

	// ConsistencyLevel overrides the account defined consistency level for this operation.
	// Consistency can only be relaxed.
	ConsistencyLevel *ConsistencyLevel

	// EnableScriptLogging is used to enable/disable logging in JavaScript stored procedures.
	// By default script logging is disabled. The log can also be accessible in response
	// header (x-ms-documentdb-script-log-results).
	EnableScriptLogging bool
}

func (options *StoredProcedureRequestOptions) toHeaders() *map[string]string {
	headers := make(map[string]string)

	if options.ConsistencyLevel != nil {
		headers[cosmosHeaderConsistencyLevel] = string(*options.ConsistencyLevel)
	}

	if options.SessionToken != "" {
		headers[cosmosHeaderSessionToken] = options.SessionToken
	}

	if options.EnableScriptLogging {
		headers[cosmosHeaderScriptEnableLogging] = "true"
	}

	return &headers
}
