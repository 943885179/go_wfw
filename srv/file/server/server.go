package server

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/dbmodel"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjstruct"
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
	UploadFile(req *dbmodel.SysFile, resp *dbmodel.Id) error
	GetFile(req *dbmodel.Id, resp *dbmodel.SysFile) error
}

func NewFile() IFile {
	return &fileSrv{}
}

type fileSrv struct {
}

func (*fileSrv) UploadFile(req *dbmodel.SysFile, resp *dbmodel.Id) error {
	db := conf.DbConfig.New()
	file := &models.SysFile{}
	mzjstruct.CopyStruct(req, file)
	if len(req.Id) == 0 {
		req.Id = mzjuuid.WorkerDefaultStr(conf.WorkerId)
	}
	resp.Id = file.Id
	return db.Create(file).Error
}
func (*fileSrv) GetFile(req *dbmodel.Id, resp *dbmodel.SysFile) error { //获取图片基本信息
	db := conf.DbConfig.New().Find(&models.SysFile{})
	if err := db.First(resp, req.Id).Error; err != nil {
		return err
	}
	return nil
}
