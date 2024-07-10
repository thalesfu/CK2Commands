package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6180] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6180][6180] = People_6180
	HistoryPeopleMap[6180][146180] = People_146180
	HistoryPeopleMap[6180][166180] = People_166180
}
