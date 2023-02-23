package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔诺瓦特_沃日PolnovatvozhBarony struct {
	feud.BaseBarony
}

var BPolnovatvozh波尔诺瓦特_沃日 feud.Barony = &波尔诺瓦特_沃日PolnovatvozhBarony{}

func init() {
	f := BPolnovatvozh波尔诺瓦特_沃日.(*波尔诺瓦特_沃日PolnovatvozhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polnovatvozh",
		TitleName: "波尔诺瓦特_沃日",
		TitleCode: "b_polnovatvozh",
	}
}
