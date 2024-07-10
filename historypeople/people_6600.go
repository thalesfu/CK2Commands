package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6600][106600] = People_106600
	HistoryPeopleMap[6600][166600] = People_166600
	HistoryPeopleMap[6600][206600] = People_206600
}
