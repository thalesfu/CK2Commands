package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3999][73999] = People_73999
}
