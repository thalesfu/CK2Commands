package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7789][7789] = People_7789
	HistoryPeopleMap[7789][167789] = People_167789
	HistoryPeopleMap[7789][247789] = People_247789
}
