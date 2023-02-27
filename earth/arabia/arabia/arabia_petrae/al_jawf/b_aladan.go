package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿丹AladanBarony struct {
	feud.BaseBarony
}

var BAladan阿丹 feud.Barony = &阿丹AladanBarony{}

func init() {
    f := BAladan阿丹.(*阿丹AladanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aladan",
		TitleName: "阿丹",
		TitleCode: "b_aladan",
	}
}
