package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫城MeauxBarony struct {
	feud.BaseBarony
}

var BMeaux莫城 feud.Barony = &莫城MeauxBarony{}

func init() {
    f := BMeaux莫城.(*莫城MeauxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meaux",
		TitleName: "莫城",
		TitleCode: "b_meaux",
	}
}
