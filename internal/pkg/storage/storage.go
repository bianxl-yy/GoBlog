package storage

import (
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
	GetKey() (string, error)
	GetType() (string, error)
	Content() ([]byte, error)
	Unmarshal([]byte) error
}

// 全局使用统一入口读写文件
var Manage *storage

// Init 设置文件的存放位置以及初始化互斥锁
// 文件存放路径参数是可选的
func (s *storage) Init(dir string) error {
	s.dir = dir
	s.locker = new(sync.Mutex)
	return nil
}

// Get 读取内容
func (s *storage) Get(v storageData) error {
	// 获取信息
	key, _ := v.GetKey()
	contentType, _ := v.GetType()
	// 组合路径
	file := path.Join(s.dir, contentType+"s", key+".json")
	// 读取内容
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.New("read storage '" + key + "' error")
	}
	// 解析内容
	return v.Unmarshal(bytes)
}

// Add 添加文件
func (s *storage) Write(v storageData) error {
	// 加锁
	s.locker.Lock()
	defer s.locker.Unlock()

	// 获取信息
	key, _ := v.GetKey()
	contentType, _ := v.GetType()
	// 获取内容
	bytes, err := v.Content()
	if err != nil {
		return err
	}
	// 无脑删文件
	// 讲道理先判断一下文件是否存在会更好
	if err := s.Remove(v); err != nil {
		return err
	}
	// 组合路径
	file := path.Join(s.dir, contentType+"s", key+".json")
	// 写入文件
	err = ioutil.WriteFile(file, bytes, 0777)
	if err != nil {
		return errors.New("write storage '" + key + "' error")
	}
	return nil
}

// Delete 删除内容
func (s *storage) Remove(v storageData) error {
	// 获取信息
	key, _ := v.GetKey()
	contentType, _ := v.GetType()
	// 组合路径
	file := path.Join(s.dir, contentType+"s", key+".json")
	// 执行删除
	if err := os.Remove(file); err != nil {
		return errors.New("delete storage '" + key + "' error")
	}
	return nil
}

// Has 根据查询到的文件信息判断内容是否存在
func (s *storage) Has(v storageData) bool {
	// 加锁
	s.locker.Lock()
	defer s.locker.Unlock()

	// 获取信息
	key, _ := v.GetKey()
	contentType, _ := v.GetType()
	// 组合路径
	file := path.Join(s.dir, contentType+"s", key+".json")
	// 获取文件信息
	_, e := os.Stat(file)

	return e == nil
}

func Init(dir string) error {
	// 初始化存储管理
	return Manage.Init(dir)
}
