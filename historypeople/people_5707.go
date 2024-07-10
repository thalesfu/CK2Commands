package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5707] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5707][125707] = People_125707
	HistoryPeopleMap[5707][145707] = People_145707
}
