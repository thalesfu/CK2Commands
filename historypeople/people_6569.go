package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6569] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6569][6569] = People_6569
	HistoryPeopleMap[6569][166569] = People_166569
}
