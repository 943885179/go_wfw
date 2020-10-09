package handler

import (
	"context"
	"qshapi/proto/file"
	"qshapi/srv/file/server"
)

type Handler struct {}

func (h Handler) UploadFile(ctx context.Context, req *file.FileReq, resp *file.FileResp) error {
	return server.NewFile().UploadFile(req,resp)
}
