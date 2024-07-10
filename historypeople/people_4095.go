package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4095] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4095][34095] = People_34095
	HistoryPeopleMap[4095][144095] = People_144095
	HistoryPeopleMap[4095][184095] = People_184095
	HistoryPeopleMap[4095][194095] = People_194095
	HistoryPeopleMap[4095][214095] = People_214095
}
