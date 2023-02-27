package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴加罗科塔BajrakotBarony struct {
	feud.BaseBarony
}

var BBajrakot巴加罗科塔 feud.Barony = &巴加罗科塔BajrakotBarony{}

func init() {
    f := BBajrakot巴加罗科塔.(*巴加罗科塔BajrakotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bajrakot",
		TitleName: "巴加罗科塔",
		TitleCode: "b_bajrakot",
	}
}
