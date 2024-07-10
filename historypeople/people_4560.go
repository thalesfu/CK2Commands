package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4560] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4560][34560] = People_34560
	HistoryPeopleMap[4560][74560] = People_74560
	HistoryPeopleMap[4560][144560] = People_144560
	HistoryPeopleMap[4560][204560] = People_204560
	HistoryPeopleMap[4560][454560] = People_454560
}
