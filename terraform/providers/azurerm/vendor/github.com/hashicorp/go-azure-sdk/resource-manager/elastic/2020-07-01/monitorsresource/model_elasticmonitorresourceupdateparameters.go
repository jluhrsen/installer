package monitorsresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticMonitorResourceUpdateParameters struct {
	Tags *map[string]string `json:"tags,omitempty"`
}
