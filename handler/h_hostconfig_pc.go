package handler

import "fit"

type HostConfigController struct {
	PCController
}

func (c HostConfigController) Get(w *fit.Response, r *fit.Request, p fit.Params) {
	c.LoadView(w, "/pc/v_host_config.html")
}
