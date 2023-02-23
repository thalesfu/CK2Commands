package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫托利MotalBarony struct {
	feud.BaseBarony
}

var BMotal莫托利 feud.Barony = &莫托利MotalBarony{}

func init() {
	f := BMotal莫托利.(*莫托利MotalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "motal",
		TitleName: "莫托利",
		TitleCode: "b_motal",
	}
}
