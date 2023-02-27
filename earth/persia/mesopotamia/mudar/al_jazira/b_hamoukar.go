package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈默卡尔HamoukarBarony struct {
	feud.BaseBarony
}

var BHamoukar哈默卡尔 feud.Barony = &哈默卡尔HamoukarBarony{}

func init() {
    f := BHamoukar哈默卡尔.(*哈默卡尔HamoukarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamoukar",
		TitleName: "哈默卡尔",
		TitleCode: "b_hamoukar",
	}
}
