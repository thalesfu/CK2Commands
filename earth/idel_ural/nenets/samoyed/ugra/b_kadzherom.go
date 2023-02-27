package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡杰罗姆KadzheromBarony struct {
	feud.BaseBarony
}

var BKadzherom卡杰罗姆 feud.Barony = &卡杰罗姆KadzheromBarony{}

func init() {
    f := BKadzherom卡杰罗姆.(*卡杰罗姆KadzheromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadzherom",
		TitleName: "卡杰罗姆",
		TitleCode: "b_kadzherom",
	}
}
