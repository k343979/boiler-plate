// ログ設定用パッケージ
package logger

import (
	log "github.com/cihub/seelog"
)

var Log log.LoggerInterface

// Set
// loggerの初期設定
func Set() {
	logger, err := log.LoggerFromConfigAsFile("log.xml")
	if err != nil {
		panic(err)
	}
	Log = logger
}