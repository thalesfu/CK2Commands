package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比雷埃夫斯PiraeausBarony struct {
	feud.BaseBarony
}

var BPiraeaus比雷埃夫斯 feud.Barony = &比雷埃夫斯PiraeausBarony{}

func init() {
    f := BPiraeaus比雷埃夫斯.(*比雷埃夫斯PiraeausBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piraeaus",
		TitleName: "比雷埃夫斯",
		TitleCode: "b_piraeaus",
	}
}
