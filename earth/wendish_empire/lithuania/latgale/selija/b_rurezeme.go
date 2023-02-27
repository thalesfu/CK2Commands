package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔泽梅RurezemeBarony struct {
	feud.BaseBarony
}

var BRurezeme库尔泽梅 feud.Barony = &库尔泽梅RurezemeBarony{}

func init() {
    f := BRurezeme库尔泽梅.(*库尔泽梅RurezemeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rurezeme",
		TitleName: "库尔泽梅",
		TitleCode: "b_rurezeme",
	}
}
