package kassala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多卡DokaBarony struct {
	feud.BaseBarony
}

var BDoka多卡 feud.Barony = &多卡DokaBarony{}

func init() {
    f := BDoka多卡.(*多卡DokaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doka",
		TitleName: "多卡",
		TitleCode: "b_doka",
	}
}
