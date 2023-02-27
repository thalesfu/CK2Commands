package orbetello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮蒂利亚诺PitiglianoBarony struct {
	feud.BaseBarony
}

var BPitigliano皮蒂利亚诺 feud.Barony = &皮蒂利亚诺PitiglianoBarony{}

func init() {
    f := BPitigliano皮蒂利亚诺.(*皮蒂利亚诺PitiglianoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pitigliano",
		TitleName: "皮蒂利亚诺",
		TitleCode: "b_pitigliano",
	}
}
