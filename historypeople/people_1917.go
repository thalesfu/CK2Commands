package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1917] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1917][161917] = People_161917
	HistoryPeopleMap[1917][191917] = People_191917
	HistoryPeopleMap[1917][461917] = People_461917
}
