package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦赫坦VakhtanBarony struct {
	feud.BaseBarony
}

var BVakhtan瓦赫坦 feud.Barony = &瓦赫坦VakhtanBarony{}

func init() {
    f := BVakhtan瓦赫坦.(*瓦赫坦VakhtanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vakhtan",
		TitleName: "瓦赫坦",
		TitleCode: "b_vakhtan",
	}
}
