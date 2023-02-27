package dasapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀舍补罗DasapuraBarony struct {
	feud.BaseBarony
}

var BDasapura陀舍补罗 feud.Barony = &陀舍补罗DasapuraBarony{}

func init() {
    f := BDasapura陀舍补罗.(*陀舍补罗DasapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dasapura",
		TitleName: "陀舍补罗",
		TitleCode: "b_dasapura",
	}
}
