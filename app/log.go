package app

import (
	"github.com/bianxl-yy/GoBlog/app/model"
	"github.com/bianxl-yy/GoBlog/app/utils"
	"io/ioutil"
	"os"
	"path"
)

// LogErrors logs error bytes to tmp/log directory.
func LogError(bytes []byte) {
	//dir := App.Config().String("app.log_dir")
	dir := model.Config.String("log_dir")
	file := path.Join(dir, utils.DateInt64(utils.Now(), "MMDDHHmmss.log"))
	ioutil.WriteFile(file, bytes, os.ModePerm)
}
