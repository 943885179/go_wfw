package handler

import (
	"context"
	"qshapi/proto/dbmodel"
	"qshapi/srv/file/server"
)

type Handler struct{}

func (h Handler) GetFile(ctx context.Context, req *dbmodel.Id, resp *dbmodel.SysFile) error {
	return server.NewFile().GetFile(req, resp)
}
func (h Handler) UploadFile(ctx context.Context, req *dbmodel.SysFile, resp *dbmodel.Id) error {
	return server.NewFile().UploadFile(req, resp)
}
