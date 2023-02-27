package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔加BalgaBarony struct {
	feud.BaseBarony
}

var BBalga巴尔加 feud.Barony = &巴尔加BalgaBarony{}

func init() {
    f := BBalga巴尔加.(*巴尔加BalgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "balga",
		TitleName: "巴尔加",
		TitleCode: "b_balga",
	}
}
