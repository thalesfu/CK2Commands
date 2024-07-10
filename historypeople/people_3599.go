package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3599] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3599][33599] = People_33599
	HistoryPeopleMap[3599][73599] = People_73599
}
