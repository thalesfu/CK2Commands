package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨伊迪耶SaldiaBarony struct {
	feud.BaseBarony
}

var BSaldia萨伊迪耶 feud.Barony = &萨伊迪耶SaldiaBarony{}

func init() {
    f := BSaldia萨伊迪耶.(*萨伊迪耶SaldiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saldia",
		TitleName: "萨伊迪耶",
		TitleCode: "b_saldia",
	}
}
