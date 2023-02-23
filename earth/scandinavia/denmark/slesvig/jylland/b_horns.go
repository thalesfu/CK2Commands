package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍恩斯HornsBarony struct {
	feud.BaseBarony
}

var BHorns霍恩斯 feud.Barony = &霍恩斯HornsBarony{}

func init() {
	f := BHorns霍恩斯.(*霍恩斯HornsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horns",
		TitleName: "霍恩斯",
		TitleCode: "b_horns",
	}
}
