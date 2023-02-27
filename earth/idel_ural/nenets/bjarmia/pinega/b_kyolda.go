package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基奥尔达KyoldaBarony struct {
	feud.BaseBarony
}

var BKyolda基奥尔达 feud.Barony = &基奥尔达KyoldaBarony{}

func init() {
    f := BKyolda基奥尔达.(*基奥尔达KyoldaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyolda",
		TitleName: "基奥尔达",
		TitleCode: "b_kyolda",
	}
}
