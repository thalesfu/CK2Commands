package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5601][125601] = People_125601
	HistoryPeopleMap[5601][145601] = People_145601
	HistoryPeopleMap[5601][205601] = People_205601
	HistoryPeopleMap[5601][465601] = People_465601
}
