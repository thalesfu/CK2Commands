package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4789][74789] = People_74789
}
