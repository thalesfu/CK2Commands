package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德拉斯提亚AdrasteaBarony struct {
	feud.BaseBarony
}

var BAdrastea阿德拉斯提亚 feud.Barony = &阿德拉斯提亚AdrasteaBarony{}

func init() {
    f := BAdrastea阿德拉斯提亚.(*阿德拉斯提亚AdrasteaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adrastea",
		TitleName: "阿德拉斯提亚",
		TitleCode: "b_adrastea",
	}
}
