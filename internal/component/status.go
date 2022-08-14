package component

type Status struct {
	LEVEL LEVEL
	HP    HP
	ATK   ATK
	DEF   DEF
	MP    MP
	AGI   AGI
}

type LEVEL struct {
	StatusInfo
}

type HP struct {
	StatusInfo
}

type ATK struct {
	StatusInfo
}

type DEF struct {
	StatusInfo
}

type MP struct {
	StatusInfo
}

type AGI struct {
	StatusInfo
}

func (s Status) GetLEVEL() int {
	return s.LEVEL.StatusInfo.ACTUAL
}

func (s Status) GetHP() int {
	return s.HP.StatusInfo.ACTUAL
}

func (s Status) GetATK() int {
	return s.ATK.StatusInfo.ACTUAL
}

func (s Status) GetDEF() int {
	return s.DEF.StatusInfo.ACTUAL
}

func (s Status) GetMP() int {
	return s.MP.StatusInfo.ACTUAL
}

func (s Status) GetAGI() int {
	return s.AGI.StatusInfo.ACTUAL
}

type StatusIface interface {
	Validate()
	Add()
	Sub()
	Mul()
}

type StatusInfo struct {
	MIN    int
	MAX    int
	ACTUAL int
}

func (sc StatusInfo) Validate() int {
	if sc.ACTUAL < sc.MIN {
		return sc.MIN
	}
	if sc.MAX < sc.ACTUAL {
		return sc.MAX
	}

	return sc.ACTUAL
}

func (sc StatusInfo) Add(a int) int {
	return sc.ACTUAL + a
}

func (sc StatusInfo) Sub(s int) int {
	return sc.ACTUAL - s
}

func (sc StatusInfo) Mul(m int) int {
	return sc.ACTUAL * m
}

func InitialLEVEL() int {
	return 1
}

func InitialHP(p PlayerType) int {
	level := InitialLEVEL()
	base := p.BasicHP
	return base*level/100 + level + 10
}

func InitialATK(p PlayerType) int {
	return p.BasicATK
}

func InitialDEF(p PlayerType) int {
	return p.BasicDEF
}

func InitialMP(p PlayerType) int {
	return p.BasicMP
}

func InitialAGI(p PlayerType) int {
	return p.BasicAGI
}
