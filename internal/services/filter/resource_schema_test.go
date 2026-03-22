// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/filter"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestFilterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*filter.FilterModel)(nil)
	schema := filter.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
