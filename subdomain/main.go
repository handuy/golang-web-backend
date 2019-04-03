package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
    cookieNameForSessionID = "mycookiesessionnameid"
    sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
)

func secret(ctx iris.Context) {
    // Check if user is authenticated
    if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
        ctx.StatusCode(iris.StatusForbidden)
        return
    }

	// Print secret message
	subdomain := ctx.Subdomain()
    ctx.Writef("Hello %s", subdomain)
}

func login(ctx iris.Context) {
    session := sess.Start(ctx)

    // Authentication goes here
    // ...

    // Set user as authenticated
	session.Set("authenticated", true)
	ctx.WriteString("Login thanh cong")
}

func loginRedirect(ctx iris.Context) {
	ctx.Redirect("http://domain.local:8080/login")
}

func main() {
	app := iris.New()

	/*
	 * Setup static files
	 */

	app.StaticWeb("/assets", "./public/assets")
	app.StaticWeb("/upload_resources", "./public/upload_resources")

	// dashboard.domain.local:8080
	dashboard := app.Subdomain("dashboard")
	dashboard.Get("/", func(ctx iris.Context) {
		ctx.Writef("HEY FROM dashboard")
	})
	dashboard.Get("/secret", secret)
	dashboard.Get("/login", loginRedirect)

	// system.domain.local:8080
	system := app.Subdomain("system")
	system.Get("/", func(ctx iris.Context) {
		ctx.Writef("HEY FROM system")
	})
	system.Get("/secret", secret)
	system.Get("/login", loginRedirect)

	// domain.local:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("HEY FROM frontend /")
	})
	app.Get("/login", login)

	app.Run(iris.Addr("domain.local:8080")) // for beginners: look ../hosts file
}