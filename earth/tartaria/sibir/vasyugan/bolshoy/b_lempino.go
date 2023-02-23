package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 连皮诺LempinoBarony struct {
	feud.BaseBarony
}

var BLempino连皮诺 feud.Barony = &连皮诺LempinoBarony{}

func init() {
	f := BLempino连皮诺.(*连皮诺LempinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lempino",
		TitleName: "连皮诺",
		TitleCode: "b_lempino",
	}
}
