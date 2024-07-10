package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7400][7400] = People_7400
	HistoryPeopleMap[7400][167400] = People_167400
	HistoryPeopleMap[7400][217400] = People_217400
	HistoryPeopleMap[7400][247400] = People_247400
}
