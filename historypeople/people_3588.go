package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3588] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3588][33588] = People_33588
	HistoryPeopleMap[3588][73588] = People_73588
	HistoryPeopleMap[3588][83588] = People_83588
}
