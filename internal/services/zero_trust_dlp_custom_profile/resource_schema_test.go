// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/zero_trust_dlp_custom_profile"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestZeroTrustDLPCustomProfileModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dlp_custom_profile.ZeroTrustDLPCustomProfileModel)(nil)
	schema := zero_trust_dlp_custom_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
