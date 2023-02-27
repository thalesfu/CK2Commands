package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普廖斯PlyosBarony struct {
	feud.BaseBarony
}

var BPlyos普廖斯 feud.Barony = &普廖斯PlyosBarony{}

func init() {
    f := BPlyos普廖斯.(*普廖斯PlyosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plyos",
		TitleName: "普廖斯",
		TitleCode: "b_plyos",
	}
}
