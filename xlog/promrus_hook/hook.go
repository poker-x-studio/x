package promrus_hook

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

// PrometheusHook exposes Prometheus counters for each of logrus' log levels.
type PrometheusHook struct {
	counterVec *prometheus.CounterVec

	fileterKey   string
	filterValues map[string]struct{}
	levels       []logrus.Level
}

// NewPrometheusHook creates a new instance of PrometheusHook which exposes Prometheus counters for various log levels.
// Contrarily to MustNewPrometheusHook, it returns an error to the caller in case of issue.
// Use NewPrometheusHook if you want more control. Use MustNewPrometheusHook if you want a less verbose hook creation.
func NewPrometheusHook(name, help string, labels []string, fileterKey string, filterValues map[string]struct{}, levels []logrus.Level) (*PrometheusHook, error) {
	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: help,
	}, labels)

	// Initialise counters for all supported levels:
	for lv := range filterValues {
		counterVec.WithLabelValues(lv)
	}
	// Try to unregister the counter vector, in case already registered for some reason,
	// e.g. double initialisation/configuration done by mistake by the end-user.
	prometheus.Unregister(counterVec)
	// Try to register the counter vector:
	err := prometheus.Register(counterVec)
	if err != nil {
		return nil, err
	}
	return &PrometheusHook{
		counterVec: counterVec,

		fileterKey:   fileterKey,
		filterValues: filterValues,
		levels:       levels,
	}, nil
}

// MustNewPrometheusHook creates a new instance of PrometheusHook which exposes Prometheus counters for various log levels.
// Contrarily to NewPrometheusHook, it does not return any error to the caller, but panics instead.
// Use MustNewPrometheusHook if you want a less verbose hook creation. Use NewPrometheusHook if you want more control.
func MustNewPrometheusHook(name, help string, labels []string, fileterKey string, filterValues map[string]struct{}, levels []logrus.Level) *PrometheusHook {
	hook, err := NewPrometheusHook(name, help, labels, fileterKey, filterValues, levels)
	if err != nil {
		panic(err)
	}
	return hook
}

// Fire increments the appropriate Prometheus counter depending on the entry's log level.
// logger方式. filterValue in filterValues
// 增加 logger.WithField(fileterKey, filterValue).WithField("add", 3.1415926).Info("xxxxx")
// 自增 logger.WithField(fileterKey, filterValue).Info("xxxxx")
func (hook *PrometheusHook) Fire(entry *logrus.Entry) error {
	if val, bfind := entry.Data[hook.fileterKey]; bfind {
		if strval, ok := val.(string); ok {
			if _, bexist := hook.filterValues[strval]; bexist {
				strlv, ok := entry.Data["add"].(string)
				if ok && len(strval) > 0 {
					addval, _ := strconv.ParseFloat(strlv, 64)
					hook.counterVec.WithLabelValues(strval).Add(addval)
				} else {
					hook.counterVec.WithLabelValues(strval).Inc()
				}
			}
		}
	}

	return nil
}

// Levels returns all supported log levels, i.e.: Debug, Info, Warn and Error, as
// there is no point incrementing a counter just before exiting/panicking.
func (hook *PrometheusHook) Levels() []logrus.Level {
	return hook.levels
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
