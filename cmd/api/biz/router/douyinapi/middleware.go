// Code generated by hertz generator.

package Douyinapi

import (
	"context"
	"douyin/cmd/api/biz/mw"
	"douyin/pkg/errno"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/requestid"
	"go.opentelemetry.io/otel/trace"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		//deal with panic
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"code":    errno.ServiceErr.ErrCode,
					"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		//generate requestId (maybe for tracing)
		requestid.New(requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
			traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
			return traceID
		})),
		//zip to accelerate web
		gzip.Gzip(gzip.DefaultCompression),
	}
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func __ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messageMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _chatrecordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _chatMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _sendmessageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _comment_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _favorite_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _feedvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list4Mw() []app.HandlerFunc {
	// your code...
	return nil
}