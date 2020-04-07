package model

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"log"
	"path"
)

const (
	uploadFileSuffix = ".jpg,.png,.gif,.zip,.txt,.doc,.docx,.xls,.xlsx,.ppt,.pptx"
)

var Config *koanf.Koanf

func init() {

	// 初始化配置
	Config = koanf.New(".")

	// 配置默认值
	Config.Load(confmap.Provider(map[string]interface{}{
		"static_dir":   "static",
		"tmp_dir":      "tmp",
		"upload_size":  1024 * 1024 * 10,
		"upload_files": uploadFileSuffix,
	}, "."), nil)
	Config.Load(confmap.Provider(map[string]interface{}{
		"log_dir":     path.Join(Config.String("tmp_dir"), "log"),
		"tmp_storage": path.Join(Config.String("tmp_dir"), "data"),
		"upload_dir":  path.Join(Config.String("static_dir"), "upload"),
	}, "."), nil)

	// 载入配置文件
	configFile := file.Provider("config.toml")
	Config.Load(configFile, toml.Parser())

	// 监视配置文件中的更改
	configFile.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("监视配置文件出错: %v", err)
			return
		}
		log.Println("配置已更改, 正在重新加载...")
		Config.Load(configFile, toml.Parser())
	})
}
