package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 秣尼MatliBarony struct {
	feud.BaseBarony
}

var BMatli秣尼 feud.Barony = &秣尼MatliBarony{}

func init() {
	f := BMatli秣尼.(*秣尼MatliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matli",
		TitleName: "秣尼",
		TitleCode: "b_matli",
	}
}
