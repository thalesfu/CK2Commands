package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科罗斯坚KorostenBarony struct {
	feud.BaseBarony
}

var BKorosten科罗斯坚 feud.Barony = &科罗斯坚KorostenBarony{}

func init() {
	f := BKorosten科罗斯坚.(*科罗斯坚KorostenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korosten",
		TitleName: "科罗斯坚",
		TitleCode: "b_korosten",
	}
}
