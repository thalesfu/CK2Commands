package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8601] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8601][8601] = People_8601
	HistoryPeopleMap[8601][168601] = People_168601
	HistoryPeopleMap[8601][188601] = People_188601
	HistoryPeopleMap[8601][248601] = People_248601
}
