package httpHandler

import (
	h "WordAudio/httpHandler/handlers"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/getAudio", h.GetAudioHandler).Methods("GET")
	router.HandleFunc("/sendAudio", h.SendAudioHandler).Methods("POST")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err := http.ListenAndServe(":8081", handlers.CORS(headersOk, originsOk, methodsOk)(router))
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println("server started")
}
