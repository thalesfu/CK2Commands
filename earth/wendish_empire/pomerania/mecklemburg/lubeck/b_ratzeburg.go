package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉策堡RatzeburgBarony struct {
	feud.BaseBarony
}

var BRatzeburg拉策堡 feud.Barony = &拉策堡RatzeburgBarony{}

func init() {
    f := BRatzeburg拉策堡.(*拉策堡RatzeburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ratzeburg",
		TitleName: "拉策堡",
		TitleCode: "b_ratzeburg",
	}
}
