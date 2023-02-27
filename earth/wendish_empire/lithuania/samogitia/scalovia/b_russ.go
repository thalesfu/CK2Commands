package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁斯RussBarony struct {
	feud.BaseBarony
}

var BRuss鲁斯 feud.Barony = &鲁斯RussBarony{}

func init() {
    f := BRuss鲁斯.(*鲁斯RussBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "russ",
		TitleName: "鲁斯",
		TitleCode: "b_russ",
	}
}
