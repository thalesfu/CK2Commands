package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林佐斯LindosBarony struct {
	feud.BaseBarony
}

var BLindos林佐斯 feud.Barony = &林佐斯LindosBarony{}

func init() {
    f := BLindos林佐斯.(*林佐斯LindosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lindos",
		TitleName: "林佐斯",
		TitleCode: "b_lindos",
	}
}
