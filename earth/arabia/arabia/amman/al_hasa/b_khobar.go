package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡拜尔KhobarBarony struct {
	feud.BaseBarony
}

var BKhobar胡拜尔 feud.Barony = &胡拜尔KhobarBarony{}

func init() {
    f := BKhobar胡拜尔.(*胡拜尔KhobarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khobar",
		TitleName: "胡拜尔",
		TitleCode: "b_khobar",
	}
}
