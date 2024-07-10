package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6809] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6809][6809] = People_6809
	HistoryPeopleMap[6809][166809] = People_166809
	HistoryPeopleMap[6809][206809] = People_206809
}
