package bot

import "log"

func parseResources (player *Player, resources []interface{}){
	for _, resource := range resources {
		var Id uint
		var Amount uint32
		Id = uint(resource.(map[string]interface{})["id"].(float64))
		Amount = uint32(resource.(map[string]interface{})["amount"].(float64))
		if(Id == GOLD_RESOURCE_ID){
			player.Gold = Amount
			log.Print("Gold - ", Amount)
		}
		if(Id == FOOD_RESOURCE_ID){
			player.Food = Amount
			log.Print("Food - ", Amount)
		}
	}
}

func parseUnits (player *Player, units []interface{}){
	player.Units = []Unit{}
	for _, unit := range units {
		player.Units = append(
			player.Units,
			Unit{
				Id: uint64(unit.(map[string]interface{})["id"].(float64)),
				Amount: uint64(unit.(map[string]interface{})["amount"].(float64))})
	}
}
func parseBuildings(player *Player, buildings []interface{}){
	player.Buildings = []Building{}
	for _, building := range buildings {
		var typeId uint
		var level uint
		typeId = uint(building.(map[string]interface{})["typeId"].(float64))
		level  = uint(building.(map[string]interface{})["level"].(float64))
		if(typeId == BARN_ID) {
			player.FoodCapacity += CAPACITIES[level]
		}
		if(typeId == TREASURY_ID) {
			player.GoldCapacity += CAPACITIES[level]
		}
		if(typeId == CASTLE_ID){
			player.CastleLvl = level
		}
		player.Buildings = append(
			player.Buildings,
			Building{
				Id:            uint64(building.(map[string]interface{})["id"].(float64)),
				TypeId:        typeId,
				Flip:          building.(map[string]interface{})["flip"].(bool),
				Level:         level,
				X:             uint(building.(map[string]interface{})["x"].(float64)),
				Y:             uint(building.(map[string]interface{})["y"].(float64)),
				Completed:     building.(map[string]interface{})["completed"].(bool),
				Volume:        uint(building.(map[string]interface{})["volume"].(float64)),
				StateTimestamp:uint64(building.(map[string]interface{})["stateTimestamp"].(float64)),
				Hitpoints:     uint64(building.(map[string]interface{})["hitpoints"].(float64)),
				CompleteTime:  uint64(building.(map[string]interface{})["completeTime"].(float64))})
	}
}
