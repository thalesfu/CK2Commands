package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4461] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4461][34461] = People_34461
	HistoryPeopleMap[4461][74461] = People_74461
}
