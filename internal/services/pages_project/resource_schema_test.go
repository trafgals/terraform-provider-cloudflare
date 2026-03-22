// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/pages_project"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestPagesProjectModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pages_project.PagesProjectModel)(nil)
	schema := pages_project.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
