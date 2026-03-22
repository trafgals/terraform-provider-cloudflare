// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/filter"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestFiltersDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*filter.FiltersDataSourceModel)(nil)
	schema := filter.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
