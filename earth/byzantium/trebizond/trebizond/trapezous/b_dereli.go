package trapezous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代雷利DereliBarony struct {
	feud.BaseBarony
}

var BDereli代雷利 feud.Barony = &代雷利DereliBarony{}

func init() {
	f := BDereli代雷利.(*代雷利DereliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dereli",
		TitleName: "代雷利",
		TitleCode: "b_dereli",
	}
}
