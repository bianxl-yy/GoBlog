package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

type storage struct {
	dir    string
	locker *sync.Mutex
}

type storageData interface {
	GetKey() string
	GetType() string
}

// 全局使用统一入口读写文件
var storageManage *storage

// Init 设置文件的存放位置以及初始化互斥锁
// 文件存放路径参数是可选的
func (s *storage) Init(dir ...string) error {
	s.dir = Config.String("storage")
	if len(dir) >= 0 {
		s.dir = dir[0]
	}
	s.locker = new(sync.Mutex)
	return nil
}

// Get 读取内容
func (s *storage) Get(v storageData) error {
	file := path.Join(s.dir, v.GetType(), v.GetKey()+".json")
	// 读取内容
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.New("read storage '" + v.GetKey() + "' error")
	}
	// 解析内容
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return errors.New("json decode '" + v.GetKey() + "' error")
	}
	return nil
}

// Add 添加文件
func (s *storage) Write(v storageData) error {
	// 加锁
	s.locker.Lock()
	defer s.locker.Unlock()
	// 编码内容
	bytes, e := json.Marshal(v)
	if e != nil {
		return errors.New("json encode '" + v.GetKey() + "' error")
	}
	// 无脑删文件
	// 讲道理先判断一下文件是否存在会更好
	if err := s.Remove(v); err != nil {
		return err
	}
	// 组合路径
	file := path.Join(s.dir, v.GetType(), v.GetKey()+".json")
	// 写入文件
	e = ioutil.WriteFile(file, bytes, 0777)
	if e != nil {
		return errors.New("write storage '" + v.GetKey() + "' error")
	}
	return nil
}

// Delete 删除内容
func (s *storage) Remove(v storageData) error {
	file := path.Join(s.dir, v.GetType(), v.GetKey()+".json")
	e := os.Remove(file)
	if e != nil {
		return errors.New("delete storage '" + v.GetKey() + "' error")
	}
	return nil
}

// Has 验证内容是否存在
func (s *storage) Has(v storageData) bool {
	file := path.Join(s.dir, v.GetType(), v.GetKey()+".json")
	_, e := os.Stat(file)
	return e == nil
}

func initStorage() error {
	// 创建必要的目录
	err := os.Mkdir(Config.String("storage"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Mkdir(Config.String("log_dir"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Mkdir(Config.String("upload_dir"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Mkdir(Config.String("article_dir"), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Mkdir(Config.String("page_dir"), os.ModePerm)
	if err != nil {
		return err
	}

	err = storageManage.Init()
	if err != nil {
		return err
	}

	return nil
}
