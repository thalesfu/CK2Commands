package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1400][31400] = People_31400
	HistoryPeopleMap[1400][71400] = People_71400
	HistoryPeopleMap[1400][91400] = People_91400
	HistoryPeopleMap[1400][161400] = People_161400
	HistoryPeopleMap[1400][191400] = People_191400
}
