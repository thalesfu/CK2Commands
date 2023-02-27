package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 先诺SiannoBarony struct {
	feud.BaseBarony
}

var BSianno先诺 feud.Barony = &先诺SiannoBarony{}

func init() {
    f := BSianno先诺.(*先诺SiannoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sianno",
		TitleName: "先诺",
		TitleCode: "b_sianno",
	}
}
