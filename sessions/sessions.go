package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("S3CR3TK3Y"))

func Flash(r *http.Request, w http.ResponseWriter) string {
	var message string = ""
	session, _ := Store.Get(r, "session")
	unTypedMessage := session.Values["MESSAGE"]
	message, ok := unTypedMessage.(string)
	if !ok {
		return ""
	}
	session.Values["MESSAGE"] = ""
	session.Save(r, w)
	return message
}
