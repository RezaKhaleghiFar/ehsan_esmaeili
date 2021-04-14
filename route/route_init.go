package route

import (
	"ehsan_esmaeili/route/client_route"
//	"ehsan_esmaeili/route/panel_route"

	"github.com/julienschmidt/httprouter"
)

func RouterInit(r *httprouter.Router) {
	client_route.ClientInit(r)
}
