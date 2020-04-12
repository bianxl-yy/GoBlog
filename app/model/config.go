package model

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"log"
	"path"
	"time"
)

const (
//uploadFileSuffix = ".jpg,.png,.gif,.zip,.txt,.doc,.docx,.xls,.xlsx,.ppt,.pptx"
)

var Config *koanf.Koanf

func GetSetting(key string) string {
	return Config.String(key)
}

func GetSettingAll() map[string]interface{} {
	return Config.All()
}

func AddSetting(key string, v string) error {
	return Config.Load(confmap.Provider(map[string]interface{}{
		key: v,
	}, "."), nil)
}

func initConfig() error {

	// 初始化配置
	Config = koanf.New(".")

	// 配置默认值
	errCheck(Config.Load(confmap.Provider(map[string]interface{}{
		"static_dir":  "static",
		"upload_size": 1024 * 1024 * 10,
		"navigators": []*Navigator{
			{
				Model: &Model{
					// TODO: 待添加
					Id: "",
					// TODO: 待添加
					AuthorId:   "",
					CreateTime: time.Now().Unix(),
				},
				Title:       "Home",
				Description: "首页",
				Link:        "/",
				Sort:        0,
			},
			{
				Model: &Model{
					// TODO: 待添加
					Id: "",
					// TODO: 待添加
					AuthorId:   "",
					CreateTime: time.Now().Unix(),
				},
				Title:       "关于",
				Description: "About",
				Link:        "/about",
				Sort:        1,
			},
		},
		//"upload_files": uploadFileSuffix,
	}, "."), nil))
	errCheck(Config.Load(confmap.Provider(map[string]interface{}{
		"article_dir": path.Join(Config.String("storage_dir"), "articles"),
		"page_dir":    path.Join(Config.String("storage_dir"), "pages"),
		"upload_dir":  path.Join(Config.String("static_dir"), "uploads"),
		"log_dir":     path.Join(Config.String("storage_dir"), "log"),
	}, "."), nil))

	// 载入配置文件
	configFile := file.Provider("config.json")
	errCheck(Config.Load(configFile, json.Parser()))

	// 监视配置文件中的更改
	configFile.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("监视配置文件出错: %v", err)
			return
		}
		log.Println("配置已更改, 正在重新加载...")
		if err := Config.Load(configFile, json.Parser()); err != nil {
			log.Println(err)
		}
	})

	return nil
}
