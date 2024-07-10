package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5222] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5222][125222] = People_125222
	HistoryPeopleMap[5222][145222] = People_145222
	HistoryPeopleMap[5222][155222] = People_155222
	HistoryPeopleMap[5222][205222] = People_205222
	HistoryPeopleMap[5222][235222] = People_235222
}
