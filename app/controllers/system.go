package controllers

import (
	"github.com/phachon/wmqx/app"
	"github.com/phachon/wmqx/app/service"
	"github.com/phachon/wmqx/container"
	"github.com/valyala/fasthttp"
	"time"
)

type SystemController struct {
	BaseController
}

// return SystemController
func NewSystemController() *SystemController {
	return &SystemController{}
}

func (this *SystemController) Index(ctx *fasthttp.RequestCtx) {
	//if !this.AccessToken(ctx) {
	//	this.jsonError(ctx, "token error", nil)
	//	return
	//}

	data := map[string]string{
		"version": app.Version,
	}

	this.jsonSuccess(ctx, "welcome use wmqx!", data)
}

// reload wmqx
func (this *SystemController) Reload(ctx *fasthttp.RequestCtx) {
	if !this.AccessToken(ctx) {
		this.jsonError(ctx, "token error", nil)
		return
	}

	app.Log.Info("wmqx start reload all message")

	service.MQ.StopAllConsumer()

	// wait all consumer stop
	for {
		if len(container.Ctx.ConsumerProcess.ProcessMessages) != 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		app.Log.Info("wmqx stop all consumer success!")
		break
	}

	err := container.Ctx.InitExchanges()
	if err != nil {
		app.Log.Error("wmqx reload error: " + err.Error())
		this.jsonError(ctx, "Reload error: "+err.Error(), nil)
		return
	}
	app.Log.Info("wmqx reload init exchange success!")

	this.jsonSuccess(ctx, "success", nil)
}
