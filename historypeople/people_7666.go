package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7666][7666] = People_7666
	HistoryPeopleMap[7666][167666] = People_167666
	HistoryPeopleMap[7666][247666] = People_247666
	HistoryPeopleMap[7666][457666] = People_457666
}
