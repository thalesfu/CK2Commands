package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4682] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4682][74682] = People_74682
	HistoryPeopleMap[4682][204682] = People_204682
}
