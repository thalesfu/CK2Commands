package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱泽尔盖LezergueBarony struct {
	feud.BaseBarony
}

var BLezergue莱泽尔盖 feud.Barony = &莱泽尔盖LezergueBarony{}

func init() {
	f := BLezergue莱泽尔盖.(*莱泽尔盖LezergueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lezergue",
		TitleName: "莱泽尔盖",
		TitleCode: "b_lezergue",
	}
}
