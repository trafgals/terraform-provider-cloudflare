// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/custom_origin_trust_store"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestCustomOriginTrustStoresDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_origin_trust_store.CustomOriginTrustStoresDataSourceModel)(nil)
	schema := custom_origin_trust_store.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
