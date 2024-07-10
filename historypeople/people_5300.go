package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5300] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5300][125300] = People_125300
	HistoryPeopleMap[5300][475300] = People_475300
	HistoryPeopleMap[5300][485300] = People_485300
}
