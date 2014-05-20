package bot

import(
	"../network"
	"log"
	"net/http"
	"html/template"
	"encoding/json"
	"strconv"
	"time"
)

func decodeJson(encoded_json []byte) *Responce {
	var resp *Responce
	err := json.Unmarshal(encoded_json, &resp)
	if(err != nil){
		log.Fatal("decodeJson", err)
	}

	return resp
}

func initGame(playerChan chan Player) {
	var playerStruct Player
	var player *Player
	playerStruct = Player{}
	player = &playerStruct
	resp := decodeJson(network.Post(getStartBoard()))
	
	for _, result := range resp.Results {
		if(result.Ident == "getSelfInfo"){
			user := result.Result["user"].(map[string]interface{})

			level, _ := strconv.Atoi(user["level"].(string))
			player.Level = uint(level)
			player.Stars = uint(user["starmoney"].(float64))

			parseUnits(player, user["unit"].([]interface{}))
			parseResources(player, user["resource"].([]interface{}))
		}
		player.GoldCapacity = 0
		player.FoodCapacity = 0
		if(result.Ident == "getBuildings"){
			parseBuildings(player, result.Result["building"].([]interface{}))
		}
	}
	playerChan <- playerStruct
}

func processCollectRequest(player *Player, resp *Responce){
	for _, result := range resp.Results {
		if(result.Ident == "group_1_body"){
			parseResources(player, result.Result["resource"].([]interface{}))
			parseBuildings(player, result.Result["building"].([]interface{}))
		}
	}
}

func collectFood(player *Player) *Responce{
	var resp *Responce
	for _, building := range player.Buildings {
		if(building.TypeId == MILL_ID){
			resp = decodeJson(network.Post(collectResource(building.Id)))
		}
	}
	return resp
}

func collectGold(player *Player) *Responce{
	var resp *Responce
	for _, building := range player.Buildings {
		if(building.TypeId == MINE_ID){
			resp = decodeJson(network.Post(collectResource(building.Id)))
		}
	}
	return resp
}

func resourcesCollector(playerChan chan Player) {
	var resp *Responce
	var playerStruct Player
	var player *Player

	resp = nil

	playerStruct = <- playerChan
	player = &playerStruct

	if(player.FoodCapacity > player.Food){
		log.Print("Collect Food")
		resp = collectFood(player)
	}
	if(player.GoldCapacity > player.Gold){
		log.Print("Collect Gold")
		resp = collectGold(player)
	}
	if(resp != nil){
		processCollectRequest(player, resp)
		resp = nil
	}
	playerChan <- playerStruct

    time.Sleep(time.Minute * 10)
	resourcesCollector(playerChan)
}

func builder(playerChan chan Player){
	var playerStruct Player
	var player *Player
	var buildingToUpgrade *Building
	var resp *Responce

	buildingToUpgrade = nil
	playerStruct = <- playerChan
	player = &playerStruct

	for _, building := range player.Buildings {
		if(
			building.TypeId == WALL_ID &&
				len(BUILDINGS.Wall) > int(building.Level - 1) &&
			player.CastleLvl >= BUILDINGS.Wall[building.Level - 1].CastleLvl &&
			player.Gold >= BUILDINGS.Wall[building.Level - 1].Cost){
			log.Print("Upgrade Wall")
			buildingToUpgrade = &building
			break
		}
	}

	if(buildingToUpgrade != nil){
		resp = decodeJson(network.Post(upgradeBuilding(buildingToUpgrade.Id)))
		processCollectRequest(player, resp)
	}

	playerChan <- playerStruct

	if(buildingToUpgrade != nil){
		time.Sleep(time.Second * time.Duration(BUILDINGS.Wall[buildingToUpgrade.Level - 1].Time))
	}else{
		time.Sleep(time.Hour)
	}

	builder(playerChan)
}

func Main(){
	var player = make(chan Player)
	go resourcesCollector(player)
	go builder(player)
	initGame(player)


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
