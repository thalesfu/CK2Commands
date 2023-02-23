package oldenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法勒尔VarelBarony struct {
	feud.BaseBarony
}

var BVarel法勒尔 feud.Barony = &法勒尔VarelBarony{}

func init() {
	f := BVarel法勒尔.(*法勒尔VarelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varel",
		TitleName: "法勒尔",
		TitleCode: "b_varel",
	}
}
