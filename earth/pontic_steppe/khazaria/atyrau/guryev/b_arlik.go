package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾里克ArlikBarony struct {
	feud.BaseBarony
}

var BArlik艾里克 feud.Barony = &艾里克ArlikBarony{}

func init() {
    f := BArlik艾里克.(*艾里克ArlikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arlik",
		TitleName: "艾里克",
		TitleCode: "b_arlik",
	}
}
