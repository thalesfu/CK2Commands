package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃雷克瓦VorykvaBarony struct {
	feud.BaseBarony
}

var BVorykva沃雷克瓦 feud.Barony = &沃雷克瓦VorykvaBarony{}

func init() {
    f := BVorykva沃雷克瓦.(*沃雷克瓦VorykvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vorykva",
		TitleName: "沃雷克瓦",
		TitleCode: "b_vorykva",
	}
}
