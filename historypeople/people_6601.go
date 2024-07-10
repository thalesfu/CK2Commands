package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6601][106601] = People_106601
	HistoryPeopleMap[6601][166601] = People_166601
	HistoryPeopleMap[6601][206601] = People_206601
}
