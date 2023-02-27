package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赫穆德_梅克捷布MakhmudmektebBarony struct {
	feud.BaseBarony
}

var BMakhmudmekteb马赫穆德_梅克捷布 feud.Barony = &马赫穆德_梅克捷布MakhmudmektebBarony{}

func init() {
    f := BMakhmudmekteb马赫穆德_梅克捷布.(*马赫穆德_梅克捷布MakhmudmektebBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makhmudmekteb",
		TitleName: "马赫穆德_梅克捷布",
		TitleCode: "b_makhmudmekteb",
	}
}
