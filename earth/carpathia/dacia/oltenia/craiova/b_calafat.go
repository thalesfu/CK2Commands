package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉法特CalafatBarony struct {
	feud.BaseBarony
}

var BCalafat卡拉法特 feud.Barony = &卡拉法特CalafatBarony{}

func init() {
	f := BCalafat卡拉法特.(*卡拉法特CalafatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calafat",
		TitleName: "卡拉法特",
		TitleCode: "b_calafat",
	}
}
