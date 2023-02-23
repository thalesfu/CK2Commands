package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃兹伏尔EidsvollBarony struct {
	feud.BaseBarony
}

var BEidsvoll埃兹伏尔 feud.Barony = &埃兹伏尔EidsvollBarony{}

func init() {
	f := BEidsvoll埃兹伏尔.(*埃兹伏尔EidsvollBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eidsvoll",
		TitleName: "埃兹伏尔",
		TitleCode: "b_eidsvoll",
	}
}
