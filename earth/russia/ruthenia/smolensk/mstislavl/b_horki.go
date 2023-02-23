package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔基HorkiBarony struct {
	feud.BaseBarony
}

var BHorki戈尔基 feud.Barony = &戈尔基HorkiBarony{}

func init() {
	f := BHorki戈尔基.(*戈尔基HorkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horki",
		TitleName: "戈尔基",
		TitleCode: "b_horki",
	}
}
