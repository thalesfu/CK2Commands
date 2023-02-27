package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维雷达ViredaBarony struct {
	feud.BaseBarony
}

var BVireda维雷达 feud.Barony = &维雷达ViredaBarony{}

func init() {
    f := BVireda维雷达.(*维雷达ViredaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vireda",
		TitleName: "维雷达",
		TitleCode: "b_vireda",
	}
}
