package sakmara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普鲁迪PrudyBarony struct {
	feud.BaseBarony
}

var BPrudy普鲁迪 feud.Barony = &普鲁迪PrudyBarony{}

func init() {
    f := BPrudy普鲁迪.(*普鲁迪PrudyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prudy",
		TitleName: "普鲁迪",
		TitleCode: "b_prudy",
	}
}
