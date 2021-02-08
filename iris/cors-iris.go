package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	app.UseRouter(crs)
	// OR per group of routes:
	// api := app.Party("/api")
	// api.AllowMethods(iris.MethodOptions) <- important for the preflight.
	// api.Use(crs)

	api := app.Party("/api").AllowMethods(iris.MethodOptions)
	api.Post("/mailer", func(ctx iris.Context) {
		var any iris.Map
		err := ctx.ReadJSON(&any)
		if err != nil {
			ctx.StopWithError(iris.StatusBadRequest, err)
			return
		}
		ctx.Application().Logger().Infof("received %#+v", any)

		accessToken := ctx.GetHeader("AccessToken")
		ctx.Application().Logger().Infof("received %#+v", accessToken)

		// var headers iris.Map
		// err1 := ctx.ReadHeaders(headers)
		// if err1 != nil {
		// 	ctx.StopWithError(iris.StatusBadRequest, err1)
		// 	return
		// }
		// ctx.Application().Logger().Infof("ctx.ReadHeaders()=> %#+v", headers)

		ctx.JSON(iris.Map{"message": "ok"})
	})

	api.Post("/send", func(ctx iris.Context) {
		ctx.WriteString("sent")
	})
	api.Put("/send", func(ctx iris.Context) {
		ctx.WriteString("updated")
	})
	api.Delete("/send", func(ctx iris.Context) {
		ctx.WriteString("deleted")
	})

	app.Listen(":8083", iris.WithTunneling)
}
