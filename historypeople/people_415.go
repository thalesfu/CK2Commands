package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[415] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[415][415] = People_415
	HistoryPeopleMap[415][30415] = People_30415
	HistoryPeopleMap[415][70415] = People_70415
	HistoryPeopleMap[415][160415] = People_160415
	HistoryPeopleMap[415][190415] = People_190415
	HistoryPeopleMap[415][260415] = People_260415
}
