package handler

import (
	"context"
	"qshapi/proto/file"
	"qshapi/srv/file/server"
)

var sv=server.Server{}
type Handler struct {}
func (h Handler) UploadFile(ctx context.Context, req *file.FileReq, resp *file.FileResp) error {
	return sv.UploadFile(req,resp)
}
