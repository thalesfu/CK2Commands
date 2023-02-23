package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔汉格尔ArkhangelBarony struct {
	feud.BaseBarony
}

var BArkhangel阿尔汉格尔 feud.Barony = &阿尔汉格尔ArkhangelBarony{}

func init() {
	f := BArkhangel阿尔汉格尔.(*阿尔汉格尔ArkhangelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkhangel",
		TitleName: "阿尔汉格尔",
		TitleCode: "b_arkhangel",
	}
}
