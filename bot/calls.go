package bot

import (
	"log"
	"encoding/json"
)

type Call struct {
	Ident string      `json:"ident"`
	Args interface {} `json:"args"`
	Name string       `json:"name"`
}

type Calls []Call

type Request struct {
	Calls Calls         `json:"calls"`
	Session interface{} `json:"session"`
}

func getFormattedData(calls Calls) []byte {
    data, err := json.Marshal(Request{Calls:calls, Session: nil})
    if(err != nil){
		log.Fatal(err)
	}
    return data
}

func getStartBoard() []byte {
	return getFormattedData(Calls{
		Call{Name:"getSelfInfo",Args:struct{}{},Ident:"getSelfInfo"},
		Call{Name:"getBuildings",Args:struct{}{},Ident:"getBuildings"}})
}
