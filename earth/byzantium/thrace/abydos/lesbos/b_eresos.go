package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃雷索斯EresosBarony struct {
	feud.BaseBarony
}

var BEresos埃雷索斯 feud.Barony = &埃雷索斯EresosBarony{}

func init() {
	f := BEresos埃雷索斯.(*埃雷索斯EresosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eresos",
		TitleName: "埃雷索斯",
		TitleCode: "b_eresos",
	}
}
