package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾因提卜尼耶Ain_et_tibniyaBarony struct {
	feud.BaseBarony
}

var BAin_et_tibniya艾因提卜尼耶 feud.Barony = &艾因提卜尼耶Ain_et_tibniyaBarony{}

func init() {
    f := BAin_et_tibniya艾因提卜尼耶.(*艾因提卜尼耶Ain_et_tibniyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ain_et_tibniya",
		TitleName: "艾因提卜尼耶",
		TitleCode: "b_ain_et_tibniya",
	}
}
