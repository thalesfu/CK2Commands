package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[789][20789] = People_20789
	HistoryPeopleMap[789][30789] = People_30789
	HistoryPeopleMap[789][70789] = People_70789
	HistoryPeopleMap[789][260789] = People_260789
}
