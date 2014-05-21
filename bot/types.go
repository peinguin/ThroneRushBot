package bot

type Unit struct {
	Id uint64 `json:"id"`
	Amount uint64 `json:"amount"`
}

type Building struct {
	Id uint64 `json:"id"`
    TypeId uint `json:"typeId"`
    Flip bool `json:"flip"`
    Level uint `json:"level"`
    X uint `json:"x"`
    Y uint `json:"y"`
    Completed bool `json:"completed"`
    Volume uint `json:"volume"`
    StateTimestamp uint64 `json:"stateTimestamp"`
    Hitpoints uint64 `json:"hitpoints"`
    CompleteTime uint64 `json:"completeTime"`
}

type Player struct {
	Units []Unit
	Buildings []Building
	Stars uint //interlan game currency
	Level uint
	Builders int //Builder house lvl
	GoldCapacity uint32
	FoodCapacity uint32
	Food uint32
	Gold uint32
	CastleLvl uint
}

type Result struct {
	Ident string `json:"ident"`
	Result map[string]interface{} `json:"result"`
}

type Error struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Call `json:"call"`
}
type Response struct{
	Date float64 `json:"date"`
	Results []Result `json:"results"`
	Error Error `json:"error"`
}

type BuildingDependency struct {
	CastleLvl uint
	Cost uint32
}

type BuildingDependencies []BuildingDependency

type Buildings struct {
	Wall BuildingDependencies
}

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
