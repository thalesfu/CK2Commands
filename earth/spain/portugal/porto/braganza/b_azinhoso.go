package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿齐尼奥苏AzinhosoBarony struct {
	feud.BaseBarony
}

var BAzinhoso阿齐尼奥苏 feud.Barony = &阿齐尼奥苏AzinhosoBarony{}

func init() {
    f := BAzinhoso阿齐尼奥苏.(*阿齐尼奥苏AzinhosoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azinhoso",
		TitleName: "阿齐尼奥苏",
		TitleCode: "b_azinhoso",
	}
}
