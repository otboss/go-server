package main

import (
	"controller"
	"controller/middleware"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("listening on port 8080")
	http.DefaultClient.Timeout = time.Second * 10
	// Routes:
	http.HandleFunc(
		"/",
		middleware.PanicHandler(
			middleware.GetRequest(
				middleware.RequestLimiter(
					controller.Root,
				),
			),
		),
	)
	// UNCOMMENT FOR HTTPS
	// if !utils.FileExists("./ssl/key.pem") || !utils.FileExists("./ssl/cert.pem") {
	// 	log.Fatalln("key.pem or cert.pem file missing. please place these files in the ssl folder.")
	// }
	// err := http.ListenAndServeTLS(":8080", "ssl/cert.pem", "ssl/key.pem", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServeTLS: ", err)
	// }
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
