// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/worker"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestWorkerDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*worker.WorkerDataSourceModel)(nil)
	schema := worker.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
