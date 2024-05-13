// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package sample provides functionality for generating sample values of
// the types contained in the internal package for testing purposes.
package sample

import (
	"fmt"
)

const (
	// ServiceURL is the hostname of the service.
	//
	// Example:
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L36
	ServiceURL = "bigquerymigration.googleapis.com"

	// ServiceName is the name of the service.
	//
	// Example:
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L35
	ServiceName = "MigrationService"

	// CreateMethod is the name of the RPC method for creating a resource.
	// The same name is used for the proto RPC method and the Go method.
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2/migrationpb#MigrationServiceClient.CreateMigrationWorkflow
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L41
	CreateMethod = "CreateMigrationWorkflow"

	// CreateRequest is the name of the request for creating a resource.
	// The same name is used for the proto message and the Go type.
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2/migrationpb#CreateMigrationWorkflowRequest
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L110
	CreateRequest = "CreateMigrationWorkflowRequest"

	// GetMethod is the name of the RPC method used to fetch a resource.
	// The same name is used for the proto RPC method and the Go method.
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2/migrationpb#MigrationServiceClient.GetMigrationWorkflow
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L51
	GetMethod = "GetMigrationWorkflow"

	// GetRequest is the name of the request for fetching a resource.
	// The same name is used for the proto message and the Go type.
	//
	// A GetRequest often contains `google.api.resource_reference`, in order to
	// reference the name of the resource (see https://aip.dev/4231#referencing-other-resources).
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2/migrationpb#GetMigrationWorkflowRequest
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L126
	GetRequest = "GetMigrationWorkflowRequest"

	// Resource is the name of the resource returned by a Get or Create request.
	//
	// A resource message often contains a `google.api.resource` option with a
	// type and pattern (see https://aip.dev/4231#resource-messages).
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2/migrationpb#MigrationWorkflow
	// https://github.com/googleapis/googleapis/blob/master/google/cloud/bigquery/migration/v2alpha/migration_entities.proto#L38
	Resource = "MigrationWorkflow"
)

const (

	// ProtoServiceName is the fully qualified name of service.
	//
	// Example:
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L35.
	ProtoServiceName = "google.cloud.bigquery.migration.v2.MigrationService"

	// ProtoPackagePath is the package path of the proto file.
	//
	// Example:
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L17
	ProtoPackagePath = "google.cloud.bigquery.migration.v2"
)

const (
	// GoPackageName is the package name for the auto-generated Go package.
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2
	GoPackageName = "migration"

	// GoPackagePath is the package import path for the auto-generated Go
	// package.
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L28
	GoPackagePath = "cloud.google.com/go/bigquery/migration/apiv2"

	// GoProtoPackageName is the package name of the auto-generated proto
	// package, which is imported by package at GoPackagePath. This name is
	// derived from the value following the ";" `go_package` in the proto file.
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2/migrationpb
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L28
	GoProtoPackageName = "migrationpb"

	// GoProtoPackagePath is the package import path of the auto-generated proto
	// package.  This name is derived from the value before the ";"
	// `go_package` in the proto file.
	//
	// Example:
	// https://pkg.go.dev/cloud.google.com/go/bigquery/migration/apiv2/migrationpb.
	// https://github.com/googleapis/googleapis/blob/f7df662a24c56ecaab79cb7d808fed4d2bb4981d/google/cloud/bigquery/migration/v2/migration_service.proto#L28
	GoProtoPackagePath = "cloud.google.com/go/bigquery/migration/apiv2/migrationpb"
)

// DescriptorInfoTypeName constructs the name format used by g.descInfo.Type.
func DescriptorInfoTypeName(typ string) string {
	return fmt.Sprintf(".%s.%s", ProtoPackagePath, typ)
}
