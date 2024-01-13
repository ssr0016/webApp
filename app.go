package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ssr0016/webapp/models"
	"github.com/ssr0016/webapp/routes"
	"github.com/ssr0016/webapp/utils"
)

const PORT = ":8080"

func main() {
	models.TestConnection()
	fmt.Printf("Listening Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
