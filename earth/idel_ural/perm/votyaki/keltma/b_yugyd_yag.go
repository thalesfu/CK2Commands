package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤格德亚格Yugyd_yagBarony struct {
	feud.BaseBarony
}

var BYugyd_yag尤格德亚格 feud.Barony = &尤格德亚格Yugyd_yagBarony{}

func init() {
    f := BYugyd_yag尤格德亚格.(*尤格德亚格Yugyd_yagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yugyd_yag",
		TitleName: "尤格德亚格",
		TitleCode: "b_yugyd_yag",
	}
}
