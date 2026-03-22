// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/address_map"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestAddressMapModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*address_map.AddressMapModel)(nil)
	schema := address_map.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
