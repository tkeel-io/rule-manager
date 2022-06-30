package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	// metrics label.
	MetricsLabelTenant = "tenant_id"

	// metrics name.
	MetricsNameRuleNum = "rule_num"

	// metrics name.
	MetricsNameRuleMax = "rule_max"
)

var CollectorRuleNumber = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: MetricsNameRuleNum,
		Help: "rule num.",
	},
	[]string{MetricsLabelTenant},
)

/*
var CollectorRuleMax = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: MetricsNameRuleMax,
		Help: "rule max.",
	},
	[]string{MetricsLabelTenant},
)
var Metrics = []prometheus.Collector{CollectorRuleMax, CollectorRuleNumber}
*/
var Metrics = []prometheus.Collector{CollectorRuleNumber}