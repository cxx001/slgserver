package middleware

import (
	"slgserver/constant"
	"slgserver/log"
	"slgserver/net"
	"go.uber.org/zap"
)

func CheckLogin() net.MiddlewareFunc {
	return func(next net.HandlerFunc) net.HandlerFunc {
		return func(req *net.WsMsgReq, rsp *net.WsMsgRsp) {

			_, err := req.Conn.GetProperty("uid")
			if err != nil {
				log.DefaultLog.Warn("connect not found uid",
					zap.String("msgName", req.Body.Name))
				rsp.Body.Code = constant.UserNotInConnect
				req.Conn.Push("account.pleaseLogin", nil)
				return
			}

			next(req, rsp)
		}
	}
}