package memel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德雷韦尔纳DrevernaBarony struct {
	feud.BaseBarony
}

var BDreverna德雷韦尔纳 feud.Barony = &德雷韦尔纳DrevernaBarony{}

func init() {
    f := BDreverna德雷韦尔纳.(*德雷韦尔纳DrevernaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dreverna",
		TitleName: "德雷韦尔纳",
		TitleCode: "b_dreverna",
	}
}
