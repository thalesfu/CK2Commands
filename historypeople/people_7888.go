package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7888] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7888][7888] = People_7888
	HistoryPeopleMap[7888][167888] = People_167888
	HistoryPeopleMap[7888][247888] = People_247888
}
