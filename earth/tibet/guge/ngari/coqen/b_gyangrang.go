package coqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 江让GyangrangBarony struct {
	feud.BaseBarony
}

var BGyangrang江让 feud.Barony = &江让GyangrangBarony{}

func init() {
    f := BGyangrang江让.(*江让GyangrangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyangrang",
		TitleName: "江让",
		TitleCode: "b_gyangrang",
	}
}
