package steiermark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁宾尼察RubinichaBarony struct {
	feud.BaseBarony
}

var BRubinicha鲁宾尼察 feud.Barony = &鲁宾尼察RubinichaBarony{}

func init() {
    f := BRubinicha鲁宾尼察.(*鲁宾尼察RubinichaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rubinicha",
		TitleName: "鲁宾尼察",
		TitleCode: "b_rubinicha",
	}
}
