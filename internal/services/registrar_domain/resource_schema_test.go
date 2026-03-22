// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/registrar_domain"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestRegistrarDomainModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*registrar_domain.RegistrarDomainModel)(nil)
	schema := registrar_domain.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
