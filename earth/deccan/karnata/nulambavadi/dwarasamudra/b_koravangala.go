package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科罗伐那加拉KoravangalaBarony struct {
	feud.BaseBarony
}

var BKoravangala科罗伐那加拉 feud.Barony = &科罗伐那加拉KoravangalaBarony{}

func init() {
    f := BKoravangala科罗伐那加拉.(*科罗伐那加拉KoravangalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koravangala",
		TitleName: "科罗伐那加拉",
		TitleCode: "b_koravangala",
	}
}
