package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4440] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4440][4440] = People_4440
	HistoryPeopleMap[4440][34440] = People_34440
	HistoryPeopleMap[4440][74440] = People_74440
	HistoryPeopleMap[4440][194440] = People_194440
}
