package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5256] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5256][125256] = People_125256
	HistoryPeopleMap[5256][145256] = People_145256
	HistoryPeopleMap[5256][155256] = People_155256
	HistoryPeopleMap[5256][205256] = People_205256
}
