package http

import "pembiayaan/src/definitions"

type Http struct {
	*definitions.AppContext
}

type IHttp interface {
	Launch()
}

func NewHttp(appContext *definitions.AppContext) IHttp {
	return &Http{
		appContext,
	}
}
