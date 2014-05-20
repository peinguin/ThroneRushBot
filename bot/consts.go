package bot

const CASTLE_ID        = 1
const MINE_ID          = 2
const TREASURY_ID      = 3
const MILL_ID          = 4
const BARN_ID          = 5
const BARRACKS_ID      = 6
const STAFF_ID         = 7
const BUILDER_HUT_ID   = 8
const FORGE_ID         = 9
const BALLISTA_ID      = 10
const WALL_ID          = 11
const ARCHER_TOWER_ID  = 12
const CANNON_ID        = 13
const THUNDER_TOWER_ID = 14
const ICE_TOWER_ID     = 15
const FIRE_TOWER_ID    = 16
const CLAN_HOUSE_ID    = 17
const DARK_TOWER_ID    = 18
const TAVERN_ID        = 19
const ALCHEMIST_ID     = 20

const GOLD_RESOURCE_ID = 1
const FOOD_RESOURCE_ID = 2

var CAPACITIES = [12]uint32 {0, 5000, 15000, 35000, 75000, 150000, 300000, 600000, 1000000, 2000000, 3000000, 4000000}

var BUILDINGS = Buildings{
	Wall: BuildingDependencies{
		BuildingDependency{CastleLvl: 2,  Cost: 250},
		BuildingDependency{CastleLvl: 2,  Cost: 500},
		BuildingDependency{CastleLvl: 3,  Cost: 1000},
		BuildingDependency{CastleLvl: 4,  Cost: 3000},
		BuildingDependency{CastleLvl: 5,  Cost: 10000},
		BuildingDependency{CastleLvl: 6,  Cost: 25000},
		BuildingDependency{CastleLvl: 7,  Cost: 60000},
		BuildingDependency{CastleLvl: 8,  Cost: 150000},
		BuildingDependency{CastleLvl: 9,  Cost: 400000},
		BuildingDependency{CastleLvl: 10, Cost: 1000000},
		BuildingDependency{CastleLvl: 11, Cost: 2000000}}}
