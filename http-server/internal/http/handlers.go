package http

import (
	netHttp "net/http"
)

func SystemHandler(w netHttp.ResponseWriter, r *netHttp.Request) {
	panic("some panic")
	w.WriteHeader(netHttp.StatusOK)
	_,_ = w.Write([]byte("PONG"))
}