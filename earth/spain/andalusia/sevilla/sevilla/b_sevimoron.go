package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫龙SevimoronBarony struct {
	feud.BaseBarony
}

var BSevimoron莫龙 feud.Barony = &莫龙SevimoronBarony{}

func init() {
    f := BSevimoron莫龙.(*莫龙SevimoronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sevimoron",
		TitleName: "莫龙",
		TitleCode: "b_sevimoron",
	}
}
