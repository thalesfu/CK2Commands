package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7890] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7890][7890] = People_7890
	HistoryPeopleMap[7890][167890] = People_167890
	HistoryPeopleMap[7890][247890] = People_247890
}
