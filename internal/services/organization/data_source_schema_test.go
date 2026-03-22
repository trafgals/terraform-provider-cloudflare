// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/organization"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestOrganizationDataSourceModelSchemaParity(t *testing.T) {
	t.Skip("FIXME: unexpected model/schema parity issues")
	t.Parallel()
	model := (*organization.OrganizationDataSourceModel)(nil)
	schema := organization.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
