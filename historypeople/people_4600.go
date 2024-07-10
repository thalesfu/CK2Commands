package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4600][74600] = People_74600
	HistoryPeopleMap[4600][144600] = People_144600
	HistoryPeopleMap[4600][204600] = People_204600
	HistoryPeopleMap[4600][214600] = People_214600
	HistoryPeopleMap[4600][454600] = People_454600
}
