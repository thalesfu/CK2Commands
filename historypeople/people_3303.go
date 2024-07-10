package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3303] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3303][33303] = People_33303
	HistoryPeopleMap[3303][73303] = People_73303
	HistoryPeopleMap[3303][83303] = People_83303
	HistoryPeopleMap[3303][93303] = People_93303
}
