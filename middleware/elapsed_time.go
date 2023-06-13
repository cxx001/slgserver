package middleware

import (
	"fmt"
	"time"

	"slgserver/log"
	"slgserver/net"
	"go.uber.org/zap"
)

func ElapsedTime() net.MiddlewareFunc {
	return func(next net.HandlerFunc) net.HandlerFunc {
		return func(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
			bt := time.Now().UnixNano()
			next(req, rsp)
			et := time.Now().UnixNano()
			diff := (et - bt)/int64(time.Millisecond)

			log.DefaultLog.Info("ElapsedTime:",
				zap.String("msgName", req.Body.Name),
				zap.String("cost", fmt.Sprintf("%dms", diff)))
		}
	}
}