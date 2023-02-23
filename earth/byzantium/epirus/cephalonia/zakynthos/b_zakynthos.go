package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎金索斯ZakynthosBarony struct {
	feud.BaseBarony
}

var BZakynthos扎金索斯 feud.Barony = &扎金索斯ZakynthosBarony{}

func init() {
	f := BZakynthos扎金索斯.(*扎金索斯ZakynthosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zakynthos",
		TitleName: "扎金索斯",
		TitleCode: "b_zakynthos",
	}
}
