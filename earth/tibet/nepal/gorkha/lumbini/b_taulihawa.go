package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 道利赫瓦TaulihawaBarony struct {
	feud.BaseBarony
}

var BTaulihawa道利赫瓦 feud.Barony = &道利赫瓦TaulihawaBarony{}

func init() {
	f := BTaulihawa道利赫瓦.(*道利赫瓦TaulihawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taulihawa",
		TitleName: "道利赫瓦",
		TitleCode: "b_taulihawa",
	}
}
