package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
	redirectAfterLogin     = ""
	mainDomain             = "techmaster.local:8080"
)

func secret(ctx iris.Context) {
	// Check if user is authenticated
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	// Print secret message
	subdomain := ctx.Subdomain()
	ctx.Writef("Đây là tin nhắn từ subdomain %s", subdomain)
}

func login(ctx iris.Context) {
	session := sess.Start(ctx)

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Set("authenticated", true)
	ctx.Redirect(redirectAfterLogin)
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)

	// Revoke users authentication
	session.Set("authenticated", false)

	subdomain := ctx.Subdomain()
	ctx.Writef("Logout từ subdomain %s", subdomain)
}

// redirect về domain chính http://domain.local:8080/login để xử lý logic đăng nhập
func dashboardLoginRedirect(ctx iris.Context) {
	redirectAfterLogin = "http://golang.techmaster.local:8080/secret"
	ctx.Redirect("http://" + mainDomain + "/login")
}

// redirect về domain chính http://domain.local:8080/login để xử lý logic đăng nhập
func systemLoginRedirect(ctx iris.Context) {
	redirectAfterLogin = "http://java.techmaster.local:8080/secret"
	ctx.Redirect("http://" + mainDomain + "/login")
}

func main() {
	app := iris.New()

	// golang.techmaster.local:8080
	dashboard := app.Subdomain("golang")
	dashboard.Get("/", func(ctx iris.Context) {
		subdomain := ctx.Subdomain()
		ctx.Writef("Hello %s", subdomain)
	})
	dashboard.Get("/course", func(ctx iris.Context) {
		ctx.Writef("Hello golang course")
	})
	dashboard.Get("/secret", secret)
	dashboard.Get("/login", dashboardLoginRedirect)
	dashboard.Get("/logout", logout)

	// java.techmaster.local:8080
	system := app.Subdomain("java")
	system.Get("/", func(ctx iris.Context) {
		subdomain := ctx.Subdomain()
		ctx.Writef("Hello %s", subdomain)
	})
	system.Get("/course", func(ctx iris.Context) {
		ctx.Writef("Hello java course")
	})
	system.Get("/secret", secret)
	system.Get("/login", systemLoginRedirect)
	system.Get("/logout", logout)

	// techmaster.local:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("Hello Techmaster")
	})
	app.Get("/course", func(ctx iris.Context) {
		ctx.Writef("Hello techmaster courses")
	})
	app.Get("/login", login)
	app.Get("/logout", logout)
	app.Get("/secret", secret)

	app.Run(iris.Addr(mainDomain))
}
