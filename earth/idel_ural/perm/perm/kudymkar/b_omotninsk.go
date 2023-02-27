package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥穆特宁斯克OmotninskBarony struct {
	feud.BaseBarony
}

var BOmotninsk奥穆特宁斯克 feud.Barony = &奥穆特宁斯克OmotninskBarony{}

func init() {
    f := BOmotninsk奥穆特宁斯克.(*奥穆特宁斯克OmotninskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "omotninsk",
		TitleName: "奥穆特宁斯克",
		TitleCode: "b_omotninsk",
	}
}
