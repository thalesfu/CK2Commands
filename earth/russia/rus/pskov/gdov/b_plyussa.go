package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普柳萨PlyussaBarony struct {
	feud.BaseBarony
}

var BPlyussa普柳萨 feud.Barony = &普柳萨PlyussaBarony{}

func init() {
    f := BPlyussa普柳萨.(*普柳萨PlyussaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plyussa",
		TitleName: "普柳萨",
		TitleCode: "b_plyussa",
	}
}
