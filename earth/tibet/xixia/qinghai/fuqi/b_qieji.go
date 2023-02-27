package fuqi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切吉QiejiBarony struct {
	feud.BaseBarony
}

var BQieji切吉 feud.Barony = &切吉QiejiBarony{}

func init() {
    f := BQieji切吉.(*切吉QiejiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qieji",
		TitleName: "切吉",
		TitleCode: "b_qieji",
	}
}
