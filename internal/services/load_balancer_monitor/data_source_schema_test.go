// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/load_balancer_monitor"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestLoadBalancerMonitorDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*load_balancer_monitor.LoadBalancerMonitorDataSourceModel)(nil)
	schema := load_balancer_monitor.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
