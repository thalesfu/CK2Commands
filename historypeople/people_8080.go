package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8080] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8080][138080] = People_138080
	HistoryPeopleMap[8080][168080] = People_168080
	HistoryPeopleMap[8080][188080] = People_188080
	HistoryPeopleMap[8080][248080] = People_248080
	HistoryPeopleMap[8080][478080] = People_478080
}
