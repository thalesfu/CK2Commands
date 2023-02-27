package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯KosBarony struct {
	feud.BaseBarony
}

var BKos科斯 feud.Barony = &科斯KosBarony{}

func init() {
    f := BKos科斯.(*科斯KosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kos",
		TitleName: "科斯",
		TitleCode: "b_kos",
	}
}
