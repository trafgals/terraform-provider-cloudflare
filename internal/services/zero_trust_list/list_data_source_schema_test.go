// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/zero_trust_list"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestZeroTrustListsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_list.ZeroTrustListsDataSourceModel)(nil)
	schema := zero_trust_list.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
