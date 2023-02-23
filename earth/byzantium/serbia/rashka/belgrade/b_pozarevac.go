package belgrade

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波扎雷瓦茨PozarevacBarony struct {
	feud.BaseBarony
}

var BPozarevac波扎雷瓦茨 feud.Barony = &波扎雷瓦茨PozarevacBarony{}

func init() {
	f := BPozarevac波扎雷瓦茨.(*波扎雷瓦茨PozarevacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pozarevac",
		TitleName: "波扎雷瓦茨",
		TitleCode: "b_pozarevac",
	}
}
