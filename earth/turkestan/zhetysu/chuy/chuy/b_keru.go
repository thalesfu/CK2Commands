package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁KeruBarony struct {
	feud.BaseBarony
}

var BKeru克鲁 feud.Barony = &克鲁KeruBarony{}

func init() {
    f := BKeru克鲁.(*克鲁KeruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "keru",
		TitleName: "克鲁",
		TitleCode: "b_keru",
	}
}
