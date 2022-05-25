package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	// metrics label.
	MetricsLabelTenant = "tenant_id"

	// metrics name.
	MetricsNameRuleNum = "core_rule_num"
)

var CollectorRuleNumber = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: MetricsNameRuleNum,
		Help: "rule num.",
	},
	[]string{MetricsLabelTenant},
)
