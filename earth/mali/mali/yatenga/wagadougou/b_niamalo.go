package wagadougou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼亚马洛NiamaloBarony struct {
	feud.BaseBarony
}

var BNiamalo尼亚马洛 feud.Barony = &尼亚马洛NiamaloBarony{}

func init() {
	f := BNiamalo尼亚马洛.(*尼亚马洛NiamaloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "niamalo",
		TitleName: "尼亚马洛",
		TitleCode: "b_niamalo",
	}
}
