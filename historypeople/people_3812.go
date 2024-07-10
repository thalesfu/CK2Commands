package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3812] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3812][3812] = People_3812
	HistoryPeopleMap[3812][73812] = People_73812
	HistoryPeopleMap[3812][93812] = People_93812
}
