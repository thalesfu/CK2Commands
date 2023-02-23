package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼萨NisaBarony struct {
	feud.BaseBarony
}

var BNisa尼萨 feud.Barony = &尼萨NisaBarony{}

func init() {
	f := BNisa尼萨.(*尼萨NisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nisa",
		TitleName: "尼萨",
		TitleCode: "b_nisa",
	}
}
