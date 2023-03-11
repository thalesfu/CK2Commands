package desht_i_kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫斯皮诺MospyneBarony struct {
	feud.BaseBarony
}

var BMospyne莫斯皮诺 feud.Barony = &莫斯皮诺MospyneBarony{}

func init() {
    f := BMospyne莫斯皮诺.(*莫斯皮诺MospyneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mospyne",
		TitleName: "莫斯皮诺",
		TitleCode: "b_mospyne",
	}
}
