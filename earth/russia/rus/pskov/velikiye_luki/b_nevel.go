package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅韦尔NevelBarony struct {
	feud.BaseBarony
}

var BNevel涅韦尔 feud.Barony = &涅韦尔NevelBarony{}

func init() {
    f := BNevel涅韦尔.(*涅韦尔NevelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nevel",
		TitleName: "涅韦尔",
		TitleCode: "b_nevel",
	}
}
