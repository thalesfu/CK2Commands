package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥博利ObolBarony struct {
	feud.BaseBarony
}

var BObol奥博利 feud.Barony = &奥博利ObolBarony{}

func init() {
	f := BObol奥博利.(*奥博利ObolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obol",
		TitleName: "奥博利",
		TitleCode: "b_obol",
	}
}
