package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷尔米塔TermitaBarony struct {
	feud.BaseBarony
}

var BTermita捷尔米塔 feud.Barony = &捷尔米塔TermitaBarony{}

func init() {
    f := BTermita捷尔米塔.(*捷尔米塔TermitaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "termita",
		TitleName: "捷尔米塔",
		TitleCode: "b_termita",
	}
}
