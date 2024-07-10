package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8666][188666] = People_188666
	HistoryPeopleMap[8666][248666] = People_248666
}
