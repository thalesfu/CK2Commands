package tamralipti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩帝利阿波罗MeltraiparaBarony struct {
	feud.BaseBarony
}

var BMeltraipara摩帝利阿波罗 feud.Barony = &摩帝利阿波罗MeltraiparaBarony{}

func init() {
    f := BMeltraipara摩帝利阿波罗.(*摩帝利阿波罗MeltraiparaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meltraipara",
		TitleName: "摩帝利阿波罗",
		TitleCode: "b_meltraipara",
	}
}
