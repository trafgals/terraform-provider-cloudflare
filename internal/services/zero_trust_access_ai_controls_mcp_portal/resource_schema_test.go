// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_portal_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/zero_trust_access_ai_controls_mcp_portal"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestZeroTrustAccessAIControlsMcpPortalModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_ai_controls_mcp_portal.ZeroTrustAccessAIControlsMcpPortalModel)(nil)
	schema := zero_trust_access_ai_controls_mcp_portal.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
