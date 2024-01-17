package routes

import (
	"net/http"

	"github.com/ssr0016/webapp/utils"
)

func adminGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "admin.html", nil)
}
