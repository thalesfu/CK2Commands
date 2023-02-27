package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波斯加PosgaBarony struct {
	feud.BaseBarony
}

var BPosga波斯加 feud.Barony = &波斯加PosgaBarony{}

func init() {
    f := BPosga波斯加.(*波斯加PosgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "posga",
		TitleName: "波斯加",
		TitleCode: "b_posga",
	}
}
