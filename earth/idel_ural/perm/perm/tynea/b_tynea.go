package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特涅阿TyneaBarony struct {
	feud.BaseBarony
}

var BTynea特涅阿 feud.Barony = &特涅阿TyneaBarony{}

func init() {
    f := BTynea特涅阿.(*特涅阿TyneaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tynea",
		TitleName: "特涅阿",
		TitleCode: "b_tynea",
	}
}
