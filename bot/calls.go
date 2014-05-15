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

func GetTstCalls() []byte {
	return getFormattedData(Calls{
		Call{
			Ident: "group_0_body",
			Args: struct{Ts int `json:"ts"`; Id string `json:"id"`}{Ts:1400134235,Id:"15"},
			Name: "battleStartMission",},
		Call{
			Ident:"group_1_body",
			Args: struct{}{},
			Name:"state"}})
}