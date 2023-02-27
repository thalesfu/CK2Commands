package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗萨KrossaBarony struct {
	feud.BaseBarony
}

var BKrossa克罗萨 feud.Barony = &克罗萨KrossaBarony{}

func init() {
    f := BKrossa克罗萨.(*克罗萨KrossaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krossa",
		TitleName: "克罗萨",
		TitleCode: "b_krossa",
	}
}
