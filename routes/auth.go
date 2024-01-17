package routes

import (
	"fmt"
	"net/http"

	"github.com/ssr0016/webapp/auth"
	"github.com/ssr0016/webapp/models"
	"github.com/ssr0016/webapp/sessions"
	"github.com/ssr0016/webapp/utils"
)

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "login.html", struct {
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	user, err := auth.Signin(email, password)
	checkErrAuthenticate(err, w, r, user)
}

func checkErrAuthenticate(err error, w http.ResponseWriter, r *http.Request, user models.User) {
	session, _ := sessions.Store.Get(r, "session")
	if err != nil {
		switch err {
		case auth.ErrEmptyFields,
			auth.ErrEmailNotFound,
			models.ErrInvalidEmail,
			auth.ErrInvalidPassword:
			session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
			session.Values["ALERT"] = "danger"
			session.Save(r, w)
			http.Redirect(w, r, "/login", 302)
			return
		default:
			utils.InternalServerError(w)
			return
		}
	}
	session.Values["USERID"] = user.Id
	session.Save(r, w)
	http.Redirect(w, r, "/admin", 302)
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "USERID")
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
