package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7071] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7071][127071] = People_127071
	HistoryPeopleMap[7071][247071] = People_247071
	HistoryPeopleMap[7071][487071] = People_487071
}
