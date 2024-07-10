package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7647] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7647][7647] = People_7647
	HistoryPeopleMap[7647][167647] = People_167647
	HistoryPeopleMap[7647][247647] = People_247647
}
