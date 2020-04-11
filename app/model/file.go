package model

import (
	"github.com/bianxl-yy/GoBlog/app/utils"
	"github.com/rs/xid"
	"os"
	"path"
	"strconv"
)

var (
	files     []*File
	fileMaxId int
)

type File struct {
	*Model
	Name string
	Url  string
	Size int64
	Type string
	Hits int
}

func (f *File) CreateFile() (*File, error) {
	id := xid.New()
	f.Id = id.String()
	f.CreateTime = utils.Now()
	f.EditTime = f.CreateTime
	f.UpdateTime = f.CreateTime
	// TODO: 待处理
	// 添加文件索引
	return f, nil
}

func CreateFilePath(dir string, f *File) string {
	os.MkdirAll(dir, os.ModePerm)
	name := utils.DateInt64(utils.Now(), "YYYYMMDDHHmmss")
	name += strconv.Itoa(Storage.TimeInc(10)) + path.Ext(f.Name)
	return path.Join(dir, name)
}

func GetFileList(page, size int) ([]*File, *utils.Pager) {
	pager := utils.NewPager(page, size, len(files))
	f := make([]*File, 0)
	if page > pager.Pages || len(files) < 1 {
		return f, pager
	}
	for i := pager.Begin; i <= pager.End; i++ {
		f = files[pager.Begin-1 : pager.End]
	}
	return f, pager
}

func RemoveFile(id int) {
	for i, f2 := range files {
		if id == f2.Id {
			files = append(files[:i], files[i+1:]...)
			os.Remove(f2.Url)
		}
	}
	go SyncFiles()
}
