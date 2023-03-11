package terebovl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫克列科夫MoklekovBarony struct {
	feud.BaseBarony
}

var BMoklekov莫克列科夫 feud.Barony = &莫克列科夫MoklekovBarony{}

func init() {
    f := BMoklekov莫克列科夫.(*莫克列科夫MoklekovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moklekov",
		TitleName: "莫克列科夫",
		TitleCode: "b_moklekov",
	}
}
