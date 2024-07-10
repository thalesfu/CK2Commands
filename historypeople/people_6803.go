package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6803] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6803][6803] = People_6803
	HistoryPeopleMap[6803][166803] = People_166803
	HistoryPeopleMap[6803][206803] = People_206803
}
