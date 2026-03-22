// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/workers_script"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestWorkersScriptDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_script.WorkersScriptDataSourceModel)(nil)
	schema := workers_script.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
