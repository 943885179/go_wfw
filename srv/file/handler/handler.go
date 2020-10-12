package handler

import (
	"context"
	"qshapi/proto/file"
	"qshapi/srv/file/server"
)

type Handler struct {}

func (h Handler) GetFile(ctx context.Context, req *file.FileId, resp *file.FileInfo) error {
	return server.NewFile().GetFile(req,resp)
}
func (h Handler) UploadFile(ctx context.Context, req *file.FileInfo, resp *file.FileId) error {
	return server.NewFile().UploadFile(req,resp)
}
