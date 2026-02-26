// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package private

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pilipili/app/video/api/internal/logic/private"
	"pilipili/app/video/api/internal/svc"
	"pilipili/app/video/api/internal/types"
)

// UP主投稿视频
func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := private.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
