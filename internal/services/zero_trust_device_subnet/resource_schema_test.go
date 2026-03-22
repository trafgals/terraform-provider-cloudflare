// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_subnet_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/zero_trust_device_subnet"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestZeroTrustDeviceSubnetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_device_subnet.ZeroTrustDeviceSubnetModel)(nil)
	schema := zero_trust_device_subnet.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
