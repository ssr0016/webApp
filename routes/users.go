package routes

import (
	"net/http"

	"github.com/ssr0016/webapp/models"
	"github.com/ssr0016/webapp/sessions"
	"github.com/ssr0016/webapp/utils"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	message := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "register.html", struct {
		Message string
	}{
		Message: message,
	})
}
func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")
	_, err := models.NewUser(user)
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["MESSAGE"] = "successfully registered!"
	session.Save(r, w)
	http.Redirect(w, r, "/register", 302)
}
