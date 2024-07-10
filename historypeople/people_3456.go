package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3456] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3456][33456] = People_33456
	HistoryPeopleMap[3456][73456] = People_73456
	HistoryPeopleMap[3456][83456] = People_83456
	HistoryPeopleMap[3456][93456] = People_93456
}
