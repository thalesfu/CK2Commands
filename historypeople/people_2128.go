package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2128] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2128][32128] = People_32128
	HistoryPeopleMap[2128][72128] = People_72128
	HistoryPeopleMap[2128][142128] = People_142128
	HistoryPeopleMap[2128][252128] = People_252128
	HistoryPeopleMap[2128][462128] = People_462128
}
