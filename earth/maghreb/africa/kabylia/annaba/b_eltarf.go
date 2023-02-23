package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔里夫EltarfBarony struct {
	feud.BaseBarony
}

var BEltarf塔里夫 feud.Barony = &塔里夫EltarfBarony{}

func init() {
	f := BEltarf塔里夫.(*塔里夫EltarfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eltarf",
		TitleName: "塔里夫",
		TitleCode: "b_eltarf",
	}
}
