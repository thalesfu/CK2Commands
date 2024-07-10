package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4456] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4456][34456] = People_34456
	HistoryPeopleMap[4456][74456] = People_74456
}
