package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼西亚NikaeaBarony struct {
	feud.BaseBarony
}

var BNikaea尼西亚 feud.Barony = &尼西亚NikaeaBarony{}

func init() {
	f := BNikaea尼西亚.(*尼西亚NikaeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nikaea",
		TitleName: "尼西亚",
		TitleCode: "b_nikaea",
	}
}
