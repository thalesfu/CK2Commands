package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7050] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7050][127050] = People_127050
	HistoryPeopleMap[7050][247050] = People_247050
	HistoryPeopleMap[7050][487050] = People_487050
}
