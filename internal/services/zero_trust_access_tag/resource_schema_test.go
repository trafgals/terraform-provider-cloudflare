// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/zero_trust_access_tag"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestZeroTrustAccessTagModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_access_tag.ZeroTrustAccessTagModel)(nil)
	schema := zero_trust_access_tag.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
