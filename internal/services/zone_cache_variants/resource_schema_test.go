// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_variants_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/zone_cache_variants"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestZoneCacheVariantsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_cache_variants.ZoneCacheVariantsModel)(nil)
	schema := zone_cache_variants.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
