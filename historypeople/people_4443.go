package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4443] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4443][4443] = People_4443
	HistoryPeopleMap[4443][34443] = People_34443
	HistoryPeopleMap[4443][74443] = People_74443
	HistoryPeopleMap[4443][194443] = People_194443
}
