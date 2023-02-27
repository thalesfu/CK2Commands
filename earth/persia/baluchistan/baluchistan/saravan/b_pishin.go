package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮欣PishinBarony struct {
	feud.BaseBarony
}

var BPishin皮欣 feud.Barony = &皮欣PishinBarony{}

func init() {
    f := BPishin皮欣.(*皮欣PishinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pishin",
		TitleName: "皮欣",
		TitleCode: "b_pishin",
	}
}
