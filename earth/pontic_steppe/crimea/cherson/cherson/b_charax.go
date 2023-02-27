package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉克斯CharaxBarony struct {
	feud.BaseBarony
}

var BCharax哈拉克斯 feud.Barony = &哈拉克斯CharaxBarony{}

func init() {
    f := BCharax哈拉克斯.(*哈拉克斯CharaxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charax",
		TitleName: "哈拉克斯",
		TitleCode: "b_charax",
	}
}
