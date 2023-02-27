package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉曾RadzynBarony struct {
	feud.BaseBarony
}

var BRadzyn拉曾 feud.Barony = &拉曾RadzynBarony{}

func init() {
    f := BRadzyn拉曾.(*拉曾RadzynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "radzyn",
		TitleName: "拉曾",
		TitleCode: "b_radzyn",
	}
}
