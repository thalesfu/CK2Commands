package khasagt_khairkhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴彦乌拉Bayan_uulBarony struct {
	feud.BaseBarony
}

var BBayan_uul巴彦乌拉 feud.Barony = &巴彦乌拉Bayan_uulBarony{}

func init() {
    f := BBayan_uul巴彦乌拉.(*巴彦乌拉Bayan_uulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayan_uul",
		TitleName: "巴彦乌拉",
		TitleCode: "b_bayan_uul",
	}
}
