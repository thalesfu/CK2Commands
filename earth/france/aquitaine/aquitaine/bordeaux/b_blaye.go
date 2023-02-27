package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布莱BlayeBarony struct {
	feud.BaseBarony
}

var BBlaye布莱 feud.Barony = &布莱BlayeBarony{}

func init() {
    f := BBlaye布莱.(*布莱BlayeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blaye",
		TitleName: "布莱",
		TitleCode: "b_blaye",
	}
}
