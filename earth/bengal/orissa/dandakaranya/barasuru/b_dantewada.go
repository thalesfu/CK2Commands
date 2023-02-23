package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达恩泰瓦达DantewadaBarony struct {
	feud.BaseBarony
}

var BDantewada达恩泰瓦达 feud.Barony = &达恩泰瓦达DantewadaBarony{}

func init() {
	f := BDantewada达恩泰瓦达.(*达恩泰瓦达DantewadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dantewada",
		TitleName: "达恩泰瓦达",
		TitleCode: "b_dantewada",
	}
}
