package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆斯塔萨里KorsholmBarony struct {
	feud.BaseBarony
}

var BKorsholm穆斯塔萨里 feud.Barony = &穆斯塔萨里KorsholmBarony{}

func init() {
    f := BKorsholm穆斯塔萨里.(*穆斯塔萨里KorsholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korsholm",
		TitleName: "穆斯塔萨里",
		TitleCode: "b_korsholm",
	}
}
