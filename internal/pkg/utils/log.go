package utils

import (
	setting "github.com/bianxl-yy/GoBlog/internal/app/system/setting/base"
	"io/ioutil"
	"os"
	"path"
)

// LogErrors logs error bytes to tmp/log directory.
func LogError(bytes []byte) {
	//dir := App.Config().String("app.log_dir")
	dir := setting.Config.String("log_dir")
	file := path.Join(dir, DateInt64(Now(), "MMDDHHmmss.log"))
	ioutil.WriteFile(file, bytes, os.ModePerm)
}
