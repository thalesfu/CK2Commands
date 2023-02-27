package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮翁比诺PiombinoBarony struct {
	feud.BaseBarony
}

var BPiombino皮翁比诺 feud.Barony = &皮翁比诺PiombinoBarony{}

func init() {
    f := BPiombino皮翁比诺.(*皮翁比诺PiombinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piombino",
		TitleName: "皮翁比诺",
		TitleCode: "b_piombino",
	}
}
