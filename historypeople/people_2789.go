package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2789][32789] = People_32789
	HistoryPeopleMap[2789][72789] = People_72789
	HistoryPeopleMap[2789][142789] = People_142789
}
