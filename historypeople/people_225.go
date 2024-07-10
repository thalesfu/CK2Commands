package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[225] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[225][225] = People_225
	HistoryPeopleMap[225][30225] = People_30225
	HistoryPeopleMap[225][70225] = People_70225
	HistoryPeopleMap[225][170225] = People_170225
	HistoryPeopleMap[225][180225] = People_180225
	HistoryPeopleMap[225][190225] = People_190225
	HistoryPeopleMap[225][200225] = People_200225
	HistoryPeopleMap[225][260225] = People_260225
}
