/*
功能：日志便利封装
说明：
*/
package xlog

import (
	"database/sql"
	"errors"
	"x/xlog/async_hook"
	"x/xlog/local_file_hook"
	"x/xlog/mysql_hook"
	"x/xlog/promrus_hook"

	"github.com/sirupsen/logrus"
)

func StandardLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

var AllLevels = []logrus.Level{logrus.TraceLevel, logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}

func Level2Levels(level logrus.Level) []logrus.Level {
	levels := make([]logrus.Level, 0)
	for _, l := range logrus.AllLevels {
		if level >= l {
			levels = append(levels, l)
		}
	}

	return levels
}

// 一些hooks
// 邮件hook 		https://github.com/zbindenren/logrus_mail
// slackrushook 	https://github.com/johntdyer/slackrus
// 钉钉hook			https://github.com/dandans-dan/dingrus
// amqp-hook		https://github.com/vladoatanasov/logrus_amqp
// logstash-hook	https://github.com/bshuster-repo/logrus-logstash-hook
// Mongodb Hooks	https://github.com/weekface/mgorus
// redis-hook		https://github.com/rogierlommers/logrus-redis-hook
// influxdb-hook	https://github.com/abramovic/logrus_influxdb
// ElasticSearch Hook	https://github.com/sohlich/elogrus
// mysql 			https://github.com/LyricTian/logrus-mysql-hook
// promrus			https://github.com/weaveworks/promrus

func AddLocalFileHook(dir, filename string, filterKey string, filterValues map[string]struct{}, level logrus.Level) (func(), error) {

	fileHook := local_file_hook.NewLocalFileHook(dir, filename, Level2Levels(level))
	asyncHook := async_hook.NewAsyncHookWithHook(fileHook, func(entry *logrus.Entry) bool {
		if filterValues == nil || filterKey == "" {
			return true
		}

		if val, bfind := entry.Data[filterKey]; bfind {
			if strval, ok := val.(string); ok {
				if _, bexist := filterValues[strval]; bexist {
					return true
				}
			}
		}

		return false
	})

	StandardLogger().AddHook(asyncHook)

	return func() {
		asyncHook.Close()
		fileHook.Close()
	}, nil
}

// addr like ":8080"，只支持counter
func AddPromrusHook(name, help string, labels []string, filterKey string, filterValues map[string]struct{}, level logrus.Level) (func(), error) {
	if filterKey == "" || filterValues == nil {
		return nil, errors.New("promrus filter key and value must set")
	}

	promrushook := promrus_hook.MustNewPrometheusHook(name, help, labels, filterKey, filterValues, Level2Levels(level))
	StandardLogger().AddHook(promrushook)

	// Expose Prometheus metrics via HTTP, as you usually would:
	// go http.ListenAndServe(addr, promhttp.Handler())
	// server := &http.Server{
	// 	Addr:    addr,
	// 	Handler: promhttp.Handler(),
	// }
	// go server.ListenAndServe()

	return func() {}, nil
}

func AddMysqlHook(db *sql.DB, tableName string, extraItems []*mysql_hook.ExecExtraItem, filterKey string, filterValues map[string]struct{}, level logrus.Level) (func(), error) {
	myhook := mysql_hook.NewWithExtra(db, tableName, extraItems, mysql_hook.SetLevels(Level2Levels(level)...))
	asyncHook := async_hook.NewAsyncHookWithHook(myhook, func(entry *logrus.Entry) bool {
		if filterValues == nil || filterKey == "" {
			return true
		}

		if val, bfind := entry.Data[filterKey]; bfind {
			if strval, ok := val.(string); ok {
				if _, bexist := filterValues[strval]; bexist {
					return true
				}
			}
		}

		return false
	})

	StandardLogger().AddHook(asyncHook)

	return func() {
		asyncHook.Close()
		myhook.Flush()
	}, nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
