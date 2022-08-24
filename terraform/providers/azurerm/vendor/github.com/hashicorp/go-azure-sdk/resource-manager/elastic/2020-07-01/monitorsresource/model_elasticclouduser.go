package monitorsresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticCloudUser struct {
	ElasticCloudSsoDefaultUrl *string `json:"elasticCloudSsoDefaultUrl,omitempty"`
	EmailAddress              *string `json:"emailAddress,omitempty"`
	Id                        *string `json:"id,omitempty"`
}
