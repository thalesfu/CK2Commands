package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2628] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2628][32628] = People_32628
	HistoryPeopleMap[2628][72628] = People_72628
	HistoryPeopleMap[2628][142628] = People_142628
}
