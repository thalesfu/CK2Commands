package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2666][72666] = People_72666
	HistoryPeopleMap[2666][142666] = People_142666
}
