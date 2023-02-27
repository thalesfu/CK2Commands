package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普翁斯克PlonskBarony struct {
	feud.BaseBarony
}

var BPlonsk普翁斯克 feud.Barony = &普翁斯克PlonskBarony{}

func init() {
    f := BPlonsk普翁斯克.(*普翁斯克PlonskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plonsk",
		TitleName: "普翁斯克",
		TitleCode: "b_plonsk",
	}
}
