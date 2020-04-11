package model

import (
	"log"
)

type Model struct {
	Id         string `json:"id"`
	AuthorId   string `json:"author_id"`
	CreateTime int64  `json:"create_time"`
	EditTime   int64  `json:"edit_time"`
	UpdateTime int64  `json:"update_time"`
}

func init() {
	// 初始化配置文件
	errCheck(initConfig())
	// 初始化程序数据
	errCheck(initStorage())
}

func errCheck(err error) {
	log.Panic(err)
}
