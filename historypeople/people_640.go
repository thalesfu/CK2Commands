package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[640] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[640][20640] = People_20640
	HistoryPeopleMap[640][30640] = People_30640
	HistoryPeopleMap[640][70640] = People_70640
	HistoryPeopleMap[640][110640] = People_110640
	HistoryPeopleMap[640][180640] = People_180640
	HistoryPeopleMap[640][260640] = People_260640
	HistoryPeopleMap[640][450640] = People_450640
}
