package component

type Player struct {
	Name        string
	Type        PlayerType
	Status      Status
	Weapon      Weapon
	ArmorOnArm  ArmorOnArm
	ArmorOnBody ArmorOnBody
	Parameter   Parameter
}

func NewPlayer(playerType PlayerType) *Player {
	return &Player{
		Type:        playerType,
		Status:      *InitStatusByPlayerType(playerType),
		Weapon:      EmptyWeapon,
		ArmorOnArm:  EmptyArmorOnArm,
		ArmorOnBody: EmptyArmorOnBody,
		Parameter:   *NewParameter(),
	}
}

type PlayerType struct {
	Name     string
	BasicHP  int
	BasicATK int
	BasicDEF int
	BasicMP  int
	BasicAGI int
}

var (
	PlayerTypeWarrior = PlayerType{
		Name:     "WARRIOR",
		BasicHP:  500,
		BasicATK: 100,
		BasicDEF: 0,
		BasicMP:  0,
		BasicAGI: 20,
	}
	PlayerTypeWizard = PlayerType{
		Name:     "WIZARD",
		BasicHP:  200,
		BasicATK: 50,
		BasicDEF: 0,
		BasicMP:  100,
		BasicAGI: 10,
	}
)

func InitStatusByPlayerType(player PlayerType) *Status {
	return &Status{
		LEVEL{
			StatusInfo{
				ACTUAL: InitialLEVEL(),
				MIN:    1,
				MAX:    100,
			},
		},
		HP{
			StatusInfo{
				ACTUAL: InitialHP(player),
				MIN:    1,
				MAX:    999,
			},
		},
		ATK{
			StatusInfo{
				ACTUAL: InitialATK(player),
				MIN:    0,
				MAX:    100,
			},
		},
		DEF{
			StatusInfo{
				ACTUAL: InitialDEF(player),
				MIN:    0,
				MAX:    100,
			},
		},
		MP{
			StatusInfo{
				ACTUAL: InitialMP(player),
				MIN:    0,
				MAX:    100,
			},
		},
		AGI{
			StatusInfo{
				ACTUAL: InitialAGI(player),
				MIN:    0,
				MAX:    100,
			},
		},
	}
}

type Condition string

const (
	ConditionHealthy  Condition = "HEALTHY"
	ConditionDead     Condition = "DEAD"
	ConditionPoisoned Condition = "POISONED"
	ConditionBurned   Condition = "BURNED"
)

type Parameter struct {
	HasWeapon      bool
	HasArmorOnArm  bool
	HasArmorOnBody bool
	Condition      Condition
}

func NewParameter() *Parameter {
	return &Parameter{
		HasWeapon:      false,
		HasArmorOnArm:  false,
		HasArmorOnBody: false,
		Condition:      ConditionHealthy,
	}
}
