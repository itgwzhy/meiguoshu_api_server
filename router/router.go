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

	//application
	app.Get("/app/{id:int}", controllers.GetOneApplication)
	app.Get("/apps", controllers.GetAllApplication)
	app.Post("/app", controllers.CreateApplication)
	app.Put("/app/{id:int}", controllers.UpdateApplication)
	app.Delete("/app/{id:int}", controllers.DeleteApplicationById)

	//consultation
	app.Get("/consultation/{id:int}", controllers.GetOneConsultation)
	app.Get("/consultations", controllers.GetAllConsultation)
	app.Post("/consultation", controllers.CreateConsultation)
	app.Put("/consultation/{id:int}", controllers.UpdateConsultation)
	app.Delete("/consultation/{id:int}", controllers.DeleteConsultationById)

	//contact
	app.Get("/contact/{id:int}", controllers.GetOneContact)
	app.Get("/contacts", controllers.GetAllContact)
	app.Post("/contact", controllers.CreateContact)
	app.Put("/contact/{id:int}", controllers.UpdateContact)
	app.Delete("/contact/{id:int}", controllers.DeleteContactById)

	//link
	app.Get("/link/{id:int}", controllers.GetOneLink)
	app.Get("/links", controllers.GetAllLink)
	app.Post("/link", controllers.CreateLink)
	app.Put("/link/{id:int}", controllers.UpdateLink)
	app.Delete("/link/{id:int}", controllers.DeleteLinkById)

	//news
	app.Get("/news/{id:int}", controllers.GetOneNews)
	app.Get("/newss", controllers.GetAllNews)
	app.Post("/news", controllers.CreateNews)
	app.Put("/news/{id:int}", controllers.UpdateNews)
	app.Delete("/news/{id:int}", controllers.DeleteNewsById)

	//product
	app.Get("/product/{id:int}", controllers.GetOneProduct)
	app.Get("/products", controllers.GetAllProduct)
	app.Post("/product", controllers.CreateProduct)
	app.Put("/product/{id:int}", controllers.UpdateProduct)
	app.Delete("/product/{id:int}", controllers.DeleteProductById)


	return app
}
