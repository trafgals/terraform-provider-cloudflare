// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/d1_database"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestD1DatabasesDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*d1_database.D1DatabasesDataSourceModel)(nil)
	schema := d1_database.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
