package handler

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"qshapi/proto/send"
	"qshapi/srv/send/server"
)

var sv=server.Server{}



type Handler struct {
	
}

func (h Handler) SendCode(ctx context.Context, req *send.SendCodeReq, resp *send.SendCodeResp) error {
	switch req.SendType {
	case send.SendType_PHONE:
		return sv.SendCodePhone(req.EmailOrPhone,resp)
	case send.SendType_EMAIL:
		return sv.SendCodeEmail(req.EmailOrPhone,resp)
	default:
		return errors.New("暂不支持该类型信息发送")
	}
}

func (h Handler) CodeVerify(ctx context.Context, req *send.CodeVerifyReq, resp *send.CodeVerifyResp) error {
	return sv.CodeVerify(req,resp)
}

func (h Handler) Send(ctx context.Context, req *send.SendReq, empty *empty.Empty) error {
	switch req.SendType {
	case send.SendType_PHONE:
		return sv.SendPhone(req.Msg,req.EmailOrPhone)
	case send.SendType_EMAIL:
		return sv.SendEmail(req.Msg,req.EmailOrPhone)
	default:
		return errors.New("暂不支持该类型信息发送")
	}
}

func (h Handler) SendAll(ctx context.Context, req *send.SendAllReq, empty *empty.Empty) error {
	switch req.SendType {
	case send.SendType_PHONE:
		return sv.SendPhone(req.Msg,req.EmailOrPhone...)
	case send.SendType_EMAIL:
		return sv.SendPhone(req.Msg,req.EmailOrPhone...)
	default:
		return errors.New("暂不支持该类型信息发送")
	}
}
