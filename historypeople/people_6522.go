package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6522] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6522][106522] = People_106522
	HistoryPeopleMap[6522][166522] = People_166522
}
