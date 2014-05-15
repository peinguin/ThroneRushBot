package bot

import(
	"../network"
	"log"
	"net/http"
	"html/template"
)

func Main(){
	log.Print(network.Post(GetTstCalls()))

	http.HandleFunc("/bot", func (w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/bot.html")
		if err != nil {
	        log.Fatal("There was an error:", err)
	    }
    	err = t.Execute(w, nil)
    	if err != nil {
	        log.Fatal("There was an error:", err)
	    }
	})
}