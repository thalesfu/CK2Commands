package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达乌斯索坦DarussoltanBarony struct {
	feud.BaseBarony
}

var BDarussoltan达乌斯索坦 feud.Barony = &达乌斯索坦DarussoltanBarony{}

func init() {
	f := BDarussoltan达乌斯索坦.(*达乌斯索坦DarussoltanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darussoltan",
		TitleName: "达乌斯索坦",
		TitleCode: "b_darussoltan",
	}
}
