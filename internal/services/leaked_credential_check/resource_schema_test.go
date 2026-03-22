// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/leaked_credential_check"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestLeakedCredentialCheckModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*leaked_credential_check.LeakedCredentialCheckModel)(nil)
	schema := leaked_credential_check.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
