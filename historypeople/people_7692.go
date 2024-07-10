package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7692] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7692][7692] = People_7692
	HistoryPeopleMap[7692][167692] = People_167692
	HistoryPeopleMap[7692][247692] = People_247692
}
