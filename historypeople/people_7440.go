package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7440] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7440][7440] = People_7440
	HistoryPeopleMap[7440][247440] = People_247440
}
