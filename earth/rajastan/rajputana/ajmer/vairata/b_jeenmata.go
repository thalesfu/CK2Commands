package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耆那摩多JeenmataBarony struct {
	feud.BaseBarony
}

var BJeenmata耆那摩多 feud.Barony = &耆那摩多JeenmataBarony{}

func init() {
    f := BJeenmata耆那摩多.(*耆那摩多JeenmataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jeenmata",
		TitleName: "耆那摩多",
		TitleCode: "b_jeenmata",
	}
}
