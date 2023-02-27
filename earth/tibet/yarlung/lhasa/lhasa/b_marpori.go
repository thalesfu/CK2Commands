package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛布日MarporiBarony struct {
	feud.BaseBarony
}

var BMarpori玛布日 feud.Barony = &玛布日MarporiBarony{}

func init() {
    f := BMarpori玛布日.(*玛布日MarporiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marpori",
		TitleName: "玛布日",
		TitleCode: "b_marpori",
	}
}
