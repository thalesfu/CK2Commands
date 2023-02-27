package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维塞乌ViseuBarony struct {
	feud.BaseBarony
}

var BViseu维塞乌 feud.Barony = &维塞乌ViseuBarony{}

func init() {
    f := BViseu维塞乌.(*维塞乌ViseuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viseu",
		TitleName: "维塞乌",
		TitleCode: "b_viseu",
	}
}
