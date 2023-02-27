package kola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶坎斯科伊JekanskoiBarony struct {
	feud.BaseBarony
}

var BJekanskoi耶坎斯科伊 feud.Barony = &耶坎斯科伊JekanskoiBarony{}

func init() {
    f := BJekanskoi耶坎斯科伊.(*耶坎斯科伊JekanskoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jekanskoi",
		TitleName: "耶坎斯科伊",
		TitleCode: "b_jekanskoi",
	}
}
