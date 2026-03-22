// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/pages_project"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestPagesProjectsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pages_project.PagesProjectsDataSourceModel)(nil)
	schema := pages_project.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
