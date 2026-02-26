// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pilipili/app/video/api/internal/logic/public"
	"pilipili/app/video/api/internal/svc"
	"pilipili/app/video/api/internal/types"
)

// 首页推荐视频流
func FeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewFeedLogic(r.Context(), svcCtx)
		resp, err := l.Feed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
