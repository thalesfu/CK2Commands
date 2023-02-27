package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波沙PoshaBarony struct {
	feud.BaseBarony
}

var BPosha波沙 feud.Barony = &波沙PoshaBarony{}

func init() {
    f := BPosha波沙.(*波沙PoshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "posha",
		TitleName: "波沙",
		TitleCode: "b_posha",
	}
}
