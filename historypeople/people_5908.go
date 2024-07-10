package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5908] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5908][125908] = People_125908
	HistoryPeopleMap[5908][145908] = People_145908
}
