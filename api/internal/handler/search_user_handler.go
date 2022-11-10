package handler

import (
	"net/http"

	"api/internal/logic"
	"api/internal/svc"
	"api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /api/user/search  searchUser
//

//

//
// Parameters:
//  + name: body
//    require: true
//    in: query
//    type: UserSearchReq
//
// Responses:
//  200: UserInfoReply
//  401: SimpleMsg
//  500: SimpleMsg

func searchUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchUserLogic(r.Context(), svcCtx)
		resp, err := l.SearchUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
