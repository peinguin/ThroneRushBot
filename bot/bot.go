package bot

import(
	"../network"
	"log"
	"net/http"
	"html/template"
	"encoding/json"
	"strconv"
)

type Unit struct {
	Id uint64 `json:"id"`
	Amount uint64 `json:"amount"`
}

type Resource struct {
	Id uint64 `json:"id"`
	Amount uint64 `json:"amount"`
}

type Building struct {
	Id uint64 `json:"id"`
    TypeId int `json:"typeId"`
    Flip bool `json:"flip"`
    Level int `json:"level"`
    X int `json:"x"`
    Y int `json:"y"`
    Completed bool `json:"completed"`
    Volume int `json:"volume"`
    StateTimestamp uint64 `json:"stateTimestamp"`
    Hitpoints uint64 `json:"hitpoints"`
    CompleteTime uint64 `json:"completeTime"`
}

type Player struct {
	Units []Unit
	Buildings []Building
	Stars int //interlan game currency
	Level int
	Resources []Resource
}

type Result struct {
	Ident string `json:"ident"`
	Result map[string]interface{} `json:"result"`
}

type Responce struct{
	Date float64 `json:"date"`
	Results []Result `json:"results"`
}

func decodeJson(encoded_json []byte) *Responce {
	var resp *Responce
	err := json.Unmarshal(encoded_json, &resp)
	if(err != nil){
		log.Fatal("decodeJson", err)
	}

	return resp
}

func initGame(player *Player) {
	resp := decodeJson(network.Post(getStartBoard()))
	
	for _, result := range resp.Results {
		if(result.Ident == "getSelfInfo"){
			user := result.Result["user"].(map[string]interface{})

			player.Level, _ = strconv.Atoi(user["level"].(string))
			player.Stars = int(user["starmoney"].(float64))
			
			player.Units = []Unit{}

			for _, unit := range user["unit"].([]interface{}) {
				player.Units = append(
					player.Units,
					Unit{
						Id: uint64(unit.(map[string]interface{})["id"].(float64)),
						Amount: uint64(unit.(map[string]interface{})["amount"].(float64))})
			}
			player.Resources = []Resource{}
			for _, resource := range user["resource"].([]interface{}) {
				player.Resources = append(
					player.Resources,
					Resource{
						Id: uint64(resource.(map[string]interface{})["id"].(float64)),
						Amount: uint64(resource.(map[string]interface{})["amount"].(float64))})
			}
		}
		if(result.Ident == "getBuildings"){
			player.Buildings = []Building{}
			buildings := result.Result["building"].([]interface{})
			for _, building := range buildings {
				player.Buildings = append(
					player.Buildings,
					Building{
						Id:            uint64(building.(map[string]interface{})["id"].(float64)),
						TypeId:        int(building.(map[string]interface{})["typeId"].(float64)),
						Flip:          building.(map[string]interface{})["flip"].(bool),
						Level:         int(building.(map[string]interface{})["level"].(float64)),
						X:             int(building.(map[string]interface{})["x"].(float64)),
						Y:             int(building.(map[string]interface{})["y"].(float64)),
						Completed:     building.(map[string]interface{})["completed"].(bool),
						Volume:        int(building.(map[string]interface{})["volume"].(float64)),
						StateTimestamp:uint64(building.(map[string]interface{})["stateTimestamp"].(float64)),
						Hitpoints:     uint64(building.(map[string]interface{})["hitpoints"].(float64)),
						CompleteTime:  uint64(building.(map[string]interface{})["completeTime"].(float64))})
			}
		}
	}
}

func Main(){
	var player Player
	initGame(&player)


	http.HandleFunc("/bot", func (w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/bot.html")
		if err != nil {
	        log.Fatal("There was an error:", err)
	    }
    	err = t.Execute(w, nil)
    	if err != nil {
	        log.Fatal("There was an error:", err)
	    }
	})
}
