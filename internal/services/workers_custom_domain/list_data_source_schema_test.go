// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/workers_custom_domain"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestWorkersCustomDomainsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*workers_custom_domain.WorkersCustomDomainsDataSourceModel)(nil)
	schema := workers_custom_domain.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
