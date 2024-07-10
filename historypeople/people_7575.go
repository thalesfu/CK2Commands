package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7575] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7575][7575] = People_7575
	HistoryPeopleMap[7575][167575] = People_167575
	HistoryPeopleMap[7575][247575] = People_247575
}
