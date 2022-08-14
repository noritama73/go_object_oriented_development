package component

type Weapon struct {
	NAME string
	ATK  int
}

type ArmorOnArm struct {
	NAME   string
	DEF    int
	WEIGHT int
}

type ArmorOnBody struct {
	NAME   string
	DEF    int
	WEIGHT int
}

var (
	EmptyWeapon = Weapon{
		NAME: "EMPTY",
		ATK:  0,
	}
	EmptyArmorOnArm = ArmorOnArm{
		NAME:   "EMPTY",
		DEF:    0,
		WEIGHT: 0,
	}
	EmptyArmorOnBody = ArmorOnBody{
		NAME:   "EMPTY",
		DEF:    0,
		WEIGHT: 0,
	}
)
