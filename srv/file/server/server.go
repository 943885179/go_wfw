package server

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/file"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjuuid"
)

var conf models.APIConfig

const sendHeard = "code_"

func init() {
	if err := mzjinit.Default(&conf); err != nil {
		log.Fatal(err)
	}
}

type IFile interface {
	UploadFile(req *file.FileInfo, resp *file.FileId) error
	GetFile(req *file.FileId, resp *file.FileInfo) error
}

func NewFile() IFile {
	return &fileSrv{}
}

type fileSrv struct {
}

func (*fileSrv) UploadFile(req *file.FileInfo, resp *file.FileId) error {
	db := conf.DbConfig.New()
	defer db.Close()
	if req.Id == 0 {
		req.Id = mzjuuid.WorkerDefault()
	}
	file := models.SysFile{
		ID:         req.Id,
		Name:       req.Name,
		Path:       req.Path,
		Size:       req.Size,
		Sort:       req.Sort,
		FileType:   req.FileType,
		FileSuffix: req.FileSuffix,
	}
	resp.Id = file.ID
	return db.Create(&file).Error
}
func (*fileSrv) GetFile(req *file.FileId, resp *file.FileInfo) error { //获取图片基本信息
	db := conf.DbConfig.New()
	defer db.Close()
	sf := &models.SysFile{}
	if err := db.First(sf, req.Id).Error; err != nil {
		return err
	}
	resp.Id = sf.ID
	resp.Name = sf.Name
	resp.Path = sf.Path
	resp.Size = sf.Size
	resp.Sort = sf.Sort
	resp.FileSuffix = sf.FileSuffix
	resp.FileType = sf.FileType
	return nil
}
