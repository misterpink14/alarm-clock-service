package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// ----- USERS -----
	router.POST("/users", UsersPOST)
	router.GET("/users/:user_id", UsersGET)

	// ----- ALARMS -----
	router.GET("/users/:user_id/alarms", AlarmsGET)
	router.POST("/users/:user_id/alarms", AlarmsPOST)
	router.GET("/users/:user_id/alarms/:alarm_id", AlarmsGET)

	log.Fatal(http.ListenAndServe(":8080", router))
}
