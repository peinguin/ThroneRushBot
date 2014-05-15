package main

import (
	"./network"
	"./bot"
	"./utils"
	"net/http"
	"log"
)

func main() {
	network.Init(USER_ID, SESSION_ID)
	original(USER_ID)
	bot.Main()

	t := utils.Template("index")
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {	
    	err := t.Execute(w, nil)
    	if err != nil {
	        log.Fatal("There was an error:", err)
	    }
	})

	http.ListenAndServe(":8080", nil )
}