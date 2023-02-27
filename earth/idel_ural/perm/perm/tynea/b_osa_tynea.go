package tynea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特涅阿Osa_tyneaBarony struct {
	feud.BaseBarony
}

var BOsa_tynea特涅阿 feud.Barony = &特涅阿Osa_tyneaBarony{}

func init() {
    f := BOsa_tynea特涅阿.(*特涅阿Osa_tyneaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osa_tynea",
		TitleName: "特涅阿",
		TitleCode: "b_osa_tynea",
	}
}
