package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 第戎DijonBarony struct {
	feud.BaseBarony
}

var BDijon第戎 feud.Barony = &第戎DijonBarony{}

func init() {
    f := BDijon第戎.(*第戎DijonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dijon",
		TitleName: "第戎",
		TitleCode: "b_dijon",
	}
}
