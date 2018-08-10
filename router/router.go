package router

import (
	"github.com/MEIGUOSHU/meiguoshu_api_server/controllers"
	"github.com/kataras/iris"
)

func InitRouter(app *iris.Application) *iris.Application {
	//app.Get("/", func(ctx iris.Context) {
	//	name := ctx.Params().Get("name")
	//	ctx.Writef("Hello %s", name)
	//})

	//依赖注入
	//app.Get("/{to:string}", hero.Handler(controllers.Hello))
	app.Get("/app/{id:int}", controllers.GetApplicationById)
	app.Get("/apps/", controllers.GetApplicationAll)
	app.Post("/app/", controllers.CreateApplication)
	app.Put("/app/{id:int}", controllers.UpdateApplicationById)
	app.Delete("/app/{id:int}", controllers.DeleteApplicationById)

	// However, this one will match /user/john/ and also /user/john/send.
	app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	return app
}
