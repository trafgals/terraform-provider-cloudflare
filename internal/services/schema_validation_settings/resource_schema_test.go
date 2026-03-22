// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_settings_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/schema_validation_settings"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestSchemaValidationSettingsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*schema_validation_settings.SchemaValidationSettingsModel)(nil)
	schema := schema_validation_settings.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
