package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5709] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5709][125709] = People_125709
	HistoryPeopleMap[5709][145709] = People_145709
}
