package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4096] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4096][34096] = People_34096
	HistoryPeopleMap[4096][144096] = People_144096
	HistoryPeopleMap[4096][184096] = People_184096
	HistoryPeopleMap[4096][194096] = People_194096
	HistoryPeopleMap[4096][214096] = People_214096
}
