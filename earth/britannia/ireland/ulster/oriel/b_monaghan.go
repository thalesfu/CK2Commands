package oriel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫纳亨MonaghanBarony struct {
	feud.BaseBarony
}

var BMonaghan莫纳亨 feud.Barony = &莫纳亨MonaghanBarony{}

func init() {
	f := BMonaghan莫纳亨.(*莫纳亨MonaghanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monaghan",
		TitleName: "莫纳亨",
		TitleCode: "b_monaghan",
	}
}
