package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨雷尔SarelBarony struct {
	feud.BaseBarony
}

var BSarel萨雷尔 feud.Barony = &萨雷尔SarelBarony{}

func init() {
    f := BSarel萨雷尔.(*萨雷尔SarelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarel",
		TitleName: "萨雷尔",
		TitleCode: "b_sarel",
	}
}
