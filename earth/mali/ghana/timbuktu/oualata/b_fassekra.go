package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法瑟卡拉FassekraBarony struct {
	feud.BaseBarony
}

var BFassekra法瑟卡拉 feud.Barony = &法瑟卡拉FassekraBarony{}

func init() {
    f := BFassekra法瑟卡拉.(*法瑟卡拉FassekraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fassekra",
		TitleName: "法瑟卡拉",
		TitleCode: "b_fassekra",
	}
}
