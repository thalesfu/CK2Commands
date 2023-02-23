package syrte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡德尔AssidrBarony struct {
	feud.BaseBarony
}

var BAssidr锡德尔 feud.Barony = &锡德尔AssidrBarony{}

func init() {
	f := BAssidr锡德尔.(*锡德尔AssidrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assidr",
		TitleName: "锡德尔",
		TitleCode: "b_assidr",
	}
}
