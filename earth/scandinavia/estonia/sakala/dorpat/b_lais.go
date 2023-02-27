package dorpat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱斯LaisBarony struct {
	feud.BaseBarony
}

var BLais莱斯 feud.Barony = &莱斯LaisBarony{}

func init() {
    f := BLais莱斯.(*莱斯LaisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lais",
		TitleName: "莱斯",
		TitleCode: "b_lais",
	}
}
