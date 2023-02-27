package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利斯提伦PolystylonBarony struct {
	feud.BaseBarony
}

var BPolystylon波利斯提伦 feud.Barony = &波利斯提伦PolystylonBarony{}

func init() {
    f := BPolystylon波利斯提伦.(*波利斯提伦PolystylonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polystylon",
		TitleName: "波利斯提伦",
		TitleCode: "b_polystylon",
	}
}
