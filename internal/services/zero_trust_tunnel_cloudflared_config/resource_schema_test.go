// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/zero_trust_tunnel_cloudflared_config"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestZeroTrustTunnelCloudflaredConfigModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_tunnel_cloudflared_config.ZeroTrustTunnelCloudflaredConfigModel)(nil)
	schema := zero_trust_tunnel_cloudflared_config.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
