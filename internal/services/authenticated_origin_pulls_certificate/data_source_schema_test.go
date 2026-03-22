// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/authenticated_origin_pulls_certificate"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestAuthenticatedOriginPullsCertificateDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*authenticated_origin_pulls_certificate.AuthenticatedOriginPullsCertificateDataSourceModel)(nil)
	schema := authenticated_origin_pulls_certificate.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
