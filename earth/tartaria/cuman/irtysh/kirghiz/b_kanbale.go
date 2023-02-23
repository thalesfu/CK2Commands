package kirghiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎巴列KanbaleBarony struct {
	feud.BaseBarony
}

var BKanbale坎巴列 feud.Barony = &坎巴列KanbaleBarony{}

func init() {
	f := BKanbale坎巴列.(*坎巴列KanbaleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanbale",
		TitleName: "坎巴列",
		TitleCode: "b_kanbale",
	}
}
