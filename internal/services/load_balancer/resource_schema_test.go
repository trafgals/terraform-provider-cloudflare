// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/load_balancer"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestLoadBalancerModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*load_balancer.LoadBalancerModel)(nil)
	schema := load_balancer.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
