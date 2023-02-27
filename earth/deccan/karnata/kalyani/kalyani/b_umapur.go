package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤摩补罗UmapurBarony struct {
	feud.BaseBarony
}

var BUmapur尤摩补罗 feud.Barony = &尤摩补罗UmapurBarony{}

func init() {
    f := BUmapur尤摩补罗.(*尤摩补罗UmapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "umapur",
		TitleName: "尤摩补罗",
		TitleCode: "b_umapur",
	}
}
