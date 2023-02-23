package sthanisvara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑诃兰那城SaharanpurBarony struct {
	feud.BaseBarony
}

var BSaharanpur娑诃兰那城 feud.Barony = &娑诃兰那城SaharanpurBarony{}

func init() {
	f := BSaharanpur娑诃兰那城.(*娑诃兰那城SaharanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saharanpur",
		TitleName: "娑诃兰那城",
		TitleCode: "b_saharanpur",
	}
}
