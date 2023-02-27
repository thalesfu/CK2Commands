package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 登卡纳尔DhenkanalBarony struct {
	feud.BaseBarony
}

var BDhenkanal登卡纳尔 feud.Barony = &登卡纳尔DhenkanalBarony{}

func init() {
    f := BDhenkanal登卡纳尔.(*登卡纳尔DhenkanalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhenkanal",
		TitleName: "登卡纳尔",
		TitleCode: "b_dhenkanal",
	}
}
