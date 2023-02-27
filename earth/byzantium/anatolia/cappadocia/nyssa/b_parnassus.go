package nyssa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕耳那索斯ParnassusBarony struct {
	feud.BaseBarony
}

var BParnassus帕耳那索斯 feud.Barony = &帕耳那索斯ParnassusBarony{}

func init() {
    f := BParnassus帕耳那索斯.(*帕耳那索斯ParnassusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parnassus",
		TitleName: "帕耳那索斯",
		TitleCode: "b_parnassus",
	}
}
