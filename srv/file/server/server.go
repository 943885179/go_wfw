package server

import (
	"github.com/micro/go-micro/v2/util/log"
	"qshapi/models"
	"qshapi/proto/file"
	"qshapi/utils/mzjinit"
	"qshapi/utils/mzjuuid"
)

var conf models.APIConfig
const sendHeard="code_"

type Server struct {

}

func init(){
	if err:=mzjinit.Default(&conf);err != nil {
		log.Fatal(err)
	}
}
func (s Server) UploadFile(req *file.FileReq, resp *file.FileResp) error {
	db:=conf.DbConfig.New()
	defer db.Close()
	file:=models.SysFile{
		ID: mzjuuid.WorkerDefault(),
		Path: req.Path,
		Size: req.Size,
		Sort: req.Sort,
		FileType: int32(req.FileType),
		FileSuffix: req.FileSuffix,
	}
	resp.Id=file.ID
	return db.Create(&file).Error
}
