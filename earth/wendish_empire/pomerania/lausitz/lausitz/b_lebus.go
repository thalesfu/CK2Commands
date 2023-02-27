package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱布斯LebusBarony struct {
	feud.BaseBarony
}

var BLebus莱布斯 feud.Barony = &莱布斯LebusBarony{}

func init() {
    f := BLebus莱布斯.(*莱布斯LebusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lebus",
		TitleName: "莱布斯",
		TitleCode: "b_lebus",
	}
}
