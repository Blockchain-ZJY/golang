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

// 获取真实的视频流媒体播放地址
func PlayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewPlayLogic(r.Context(), svcCtx)
		resp, err := l.Play(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
