package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3440] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3440][73440] = People_73440
	HistoryPeopleMap[3440][83440] = People_83440
	HistoryPeopleMap[3440][93440] = People_93440
}
