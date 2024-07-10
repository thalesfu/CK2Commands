package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4800][34800] = People_34800
	HistoryPeopleMap[4800][74800] = People_74800
	HistoryPeopleMap[4800][204800] = People_204800
	HistoryPeopleMap[4800][214800] = People_214800
}
