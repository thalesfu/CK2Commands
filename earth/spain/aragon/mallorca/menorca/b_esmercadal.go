package menorca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯梅卡达尔EsmercadalBarony struct {
	feud.BaseBarony
}

var BEsmercadal埃斯梅卡达尔 feud.Barony = &埃斯梅卡达尔EsmercadalBarony{}

func init() {
	f := BEsmercadal埃斯梅卡达尔.(*埃斯梅卡达尔EsmercadalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esmercadal",
		TitleName: "埃斯梅卡达尔",
		TitleCode: "b_esmercadal",
	}
}
