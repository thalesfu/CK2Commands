package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃雷布尼ErebuniBarony struct {
	feud.BaseBarony
}

var BErebuni埃雷布尼 feud.Barony = &埃雷布尼ErebuniBarony{}

func init() {
    f := BErebuni埃雷布尼.(*埃雷布尼ErebuniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erebuni",
		TitleName: "埃雷布尼",
		TitleCode: "b_erebuni",
	}
}
