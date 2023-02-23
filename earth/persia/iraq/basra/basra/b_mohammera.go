package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫哈梅拉MohammeraBarony struct {
	feud.BaseBarony
}

var BMohammera莫哈梅拉 feud.Barony = &莫哈梅拉MohammeraBarony{}

func init() {
	f := BMohammera莫哈梅拉.(*莫哈梅拉MohammeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mohammera",
		TitleName: "莫哈梅拉",
		TitleCode: "b_mohammera",
	}
}
