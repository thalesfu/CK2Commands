package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[99] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[99][99] = People_99
	HistoryPeopleMap[99][30099] = People_30099
	HistoryPeopleMap[99][70099] = People_70099
	HistoryPeopleMap[99][160099] = People_160099
	HistoryPeopleMap[99][170099] = People_170099
	HistoryPeopleMap[99][180099] = People_180099
	HistoryPeopleMap[99][190099] = People_190099
	HistoryPeopleMap[99][200099] = People_200099
	HistoryPeopleMap[99][260099] = People_260099
	HistoryPeopleMap[99][470099] = People_470099
}
