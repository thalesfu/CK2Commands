package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕贝克LubeckBarony struct {
	feud.BaseBarony
}

var BLubeck吕贝克 feud.Barony = &吕贝克LubeckBarony{}

func init() {
    f := BLubeck吕贝克.(*吕贝克LubeckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lubeck",
		TitleName: "吕贝克",
		TitleCode: "b_lubeck",
	}
}
