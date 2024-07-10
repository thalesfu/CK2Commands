package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5400][125400] = People_125400
	HistoryPeopleMap[5400][465400] = People_465400
	HistoryPeopleMap[5400][475400] = People_475400
	HistoryPeopleMap[5400][485400] = People_485400
}
