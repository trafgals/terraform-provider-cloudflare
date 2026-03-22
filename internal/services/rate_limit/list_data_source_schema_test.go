// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limit_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/rate_limit"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestRateLimitsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*rate_limit.RateLimitsDataSourceModel)(nil)
	schema := rate_limit.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
