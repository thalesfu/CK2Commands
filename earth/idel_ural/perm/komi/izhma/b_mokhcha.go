package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫赫恰MokhchaBarony struct {
	feud.BaseBarony
}

var BMokhcha莫赫恰 feud.Barony = &莫赫恰MokhchaBarony{}

func init() {
    f := BMokhcha莫赫恰.(*莫赫恰MokhchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mokhcha",
		TitleName: "莫赫恰",
		TitleCode: "b_mokhcha",
	}
}
