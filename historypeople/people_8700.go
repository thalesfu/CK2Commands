package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8700][168700] = People_168700
	HistoryPeopleMap[8700][188700] = People_188700
	HistoryPeopleMap[8700][248700] = People_248700
}
