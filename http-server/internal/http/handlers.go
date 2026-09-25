package http

import (
	netHttp "net/http"
)

func SystemHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	w.WriteHeader(netHttp.StatusOK)
	_,_ = w.Write([]byte("PONG"))
}