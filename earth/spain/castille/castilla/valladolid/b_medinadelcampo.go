package valladolid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎波城MedinadelcampoBarony struct {
	feud.BaseBarony
}

var BMedinadelcampo坎波城 feud.Barony = &坎波城MedinadelcampoBarony{}

func init() {
    f := BMedinadelcampo坎波城.(*坎波城MedinadelcampoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medinadelcampo",
		TitleName: "坎波城",
		TitleCode: "b_medinadelcampo",
	}
}
