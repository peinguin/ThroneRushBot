package bot

import(
	"../network"
	"log"
	"net/http"
	"html/template"
	"encoding/json"
	"strconv"
)

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
				var Id uint
				var Amount uint32
				Id = uint(resource.(map[string]interface{})["id"].(float64))
				Amount = uint32(resource.(map[string]interface{})["amount"].(float64))
				if(Id == GOLD_RESOURCE_ID){
					player.Gold = Amount
				}
				if(Id == FOOD_RESOURCE_ID){
					player.Food = Amount
				}
			}
		}
		player.GoldCapacity = 0
		player.FoodCapacity = 0
		if(result.Ident == "getBuildings"){
			player.Buildings = []Building{}
			buildings := result.Result["building"].([]interface{})
			for _, building := range buildings {
				var typeId int
				typeId = int(building.(map[string]interface{})["typeId"].(float64))
				if(typeId == BARN_ID) {
					player.FoodCapacity += CAPACITIES[BARN_ID]
				}
				if(typeId == TREASURY_ID) {
					player.FoodCapacity += CAPACITIES[BARN_ID]
				}
				player.Buildings = append(
					player.Buildings,
					Building{
						Id:            uint64(building.(map[string]interface{})["id"].(float64)),
						TypeId:        typeId,
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

func collectFood(player *Player) {
	for _, building := rande player.Buildings {
		if(building.TypeId == MILL_ID){
			resp := decodeJson(network.Post(collectResource(building.Id)))
		}
	}
}

func collectGold(player *Player) {

}

func resourcesCollector(player *Player) {
	ticker := time.NewTicker(time.Minute * 5)
    go func() {
        for t := range ticker.C {
			if(player.FoodCapacity > player.Food){
				log.Print("Collect Food")
				collectFood(player)
			}
			if(player.GoldCapacity > player.Gold){
				log.Print("Collect Gold")
				collectGold(player)
			}
        }
    }()
}

func Main(){
	var player Player
	initGame(&player)
	resourcesCollector(&player)

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
