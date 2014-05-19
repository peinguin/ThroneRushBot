package bot

type Unit struct {
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
	Builders int //Builder house lvl
	GoldCapacity uint32
	FoodCapacity uint32
	Food uint32
	Gold uint32
}

type Result struct {
	Ident string `json:"ident"`
	Result map[string]interface{} `json:"result"`
}

type Responce struct{
	Date float64 `json:"date"`
	Results []Result `json:"results"`
}
