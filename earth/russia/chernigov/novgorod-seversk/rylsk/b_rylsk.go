package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷利斯克RylskBarony struct {
	feud.BaseBarony
}

var BRylsk雷利斯克 feud.Barony = &雷利斯克RylskBarony{}

func init() {
    f := BRylsk雷利斯克.(*雷利斯克RylskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rylsk",
		TitleName: "雷利斯克",
		TitleCode: "b_rylsk",
	}
}
