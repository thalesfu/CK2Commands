package lucca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮斯托亚PistoiaBarony struct {
	feud.BaseBarony
}

var BPistoia皮斯托亚 feud.Barony = &皮斯托亚PistoiaBarony{}

func init() {
	f := BPistoia皮斯托亚.(*皮斯托亚PistoiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pistoia",
		TitleName: "皮斯托亚",
		TitleCode: "b_pistoia",
	}
}
