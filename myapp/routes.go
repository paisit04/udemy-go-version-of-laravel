package main

import (
	"fmt"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/paisit04/celeritas/mailer"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes
	a.use(a.Middleware.CheckRemember)

	// add routes here
	a.get("/", a.Handler.Home)
	a.App.Routes.Get("/go-page", a.Handler.GoPage)
	a.App.Routes.Get("/jet-page", a.Handler.JetPage)
	a.App.Routes.Get("/sessions", a.Handler.SessionTest)

	a.App.Routes.Get("/users/login", a.Handler.UserLogin)
	a.post("/users/login", a.Handler.PostUserLogin)
	a.App.Routes.Get("/users/logout", a.Handler.Logout)
	a.get("/users/forgot-password", a.Handler.Forgot)
	a.post("/users/forgot-password", a.Handler.PostForgot)
	a.get("/users/reset-password", a.Handler.ResetPasswordForm)
	a.post("/users/reset-password", a.Handler.PostResetPassword)

	a.App.Routes.Get("/form", a.Handler.Form)
	a.App.Routes.Post("/form", a.Handler.PostForm)

	a.get("/json", a.Handler.JSON)
	a.get("/xml", a.Handler.XML)
	a.get("/download-file", a.Handler.DownloadFile)

	a.get("/crypto", a.Handler.TestCrypto)

	a.get("/cache-test", a.Handler.ShowCachePage)
	a.post("/api/save-in-cache", a.Handler.SaveInCache)
	a.post("/api/get-from-cache", a.Handler.GetFromCache)
	a.post("/api/delete-from-cache", a.Handler.DeleteFromCache)
	a.post("/api/empty-cache", a.Handler.EmptyCache)

	a.get("/test-mail", func(w http.ResponseWriter, r *http.Request) {
		// msg := mailer.Message{
		// 	From:        "test@example.com",
		// 	To:          "you@there.com",
		// 	Subject:     "Test Subject - send using channel",
		// 	Template:    "test",
		// 	Attachments: nil,
		// 	Data:        nil,
		// }
		msg := mailer.Message{
			From:        "info@mg.enres.co",
			To:          "paisit.j@gmail.com",
			Subject:     "Test Subject - send using an API",
			Template:    "test",
			Attachments: nil,
			Data:        nil,
		}

		a.App.Mail.Jobs <- msg
		res := <-a.App.Mail.Results
		if res.Error != nil {
			a.App.ErrorLog.Println(res.Error)
		}

		// err := a.App.Mail.SendSMTPMessage(msg)
		// if err != nil {
		// 	a.App.ErrorLog.Println(err)
		// 	return
		// }
		fmt.Fprint(w, "Send mail!")
	})

	a.App.Routes.Get("/create-user", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "Trevor",
			LastName:  "Sawler",
			Email:     "me@here.com",
			Active:    1,
			Password:  "password",
		}

		id, err := a.Models.Users.Insert(u)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d: %s", id, u.FirstName)
	})

	a.App.Routes.Get("/get-all-users", func(w http.ResponseWriter, r *http.Request) {
		users, err := a.Models.Users.GetAll()
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		for _, x := range users {
			fmt.Fprint(w, x.LastName)
		}
	})

	a.App.Routes.Get("/get-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		u, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		fmt.Fprintf(w, "%s %s %s", u.FirstName, u.LastName, u.Email)
	})

	a.App.Routes.Get("/update-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		u.LastName = a.App.RandomString(10)

		validator := a.App.Validator(nil)
		u.LastName = ""

		u.Validate(validator)

		if !validator.Valid() {
			fmt.Fprint(w, "failed validation")
			return
		}

		err = u.Update(*u)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		fmt.Fprintf(w, "updated last name to %s", u.LastName)
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
