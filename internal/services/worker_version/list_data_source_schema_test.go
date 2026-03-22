// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/worker_version"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestWorkerVersionsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*worker_version.WorkerVersionsDataSourceModel)(nil)
	schema := worker_version.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
