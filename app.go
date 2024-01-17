package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ssr0016/webapp/models"
	"github.com/ssr0016/webapp/routes"
	"github.com/ssr0016/webapp/sessions"
	"github.com/ssr0016/webapp/utils"
)

const PORT = ":8080"

func main() {
	models.TestConnection()

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	fmt.Println("Not Port")
	// 	os.Exit(1)
	// }
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = ":8080"
	// }
	fmt.Printf("Listening Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	sessions.SessionOptions("localhost", "/", 3600, true)
	http.Handle("/", r)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	log.Fatal(http.ListenAndServe(PORT, nil))
}
