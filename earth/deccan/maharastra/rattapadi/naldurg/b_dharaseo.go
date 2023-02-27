package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达拉斯DharaseoBarony struct {
	feud.BaseBarony
}

var BDharaseo达拉斯 feud.Barony = &达拉斯DharaseoBarony{}

func init() {
    f := BDharaseo达拉斯.(*达拉斯DharaseoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharaseo",
		TitleName: "达拉斯",
		TitleCode: "b_dharaseo",
	}
}
