package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7800][7800] = People_7800
	HistoryPeopleMap[7800][167800] = People_167800
	HistoryPeopleMap[7800][247800] = People_247800
}
