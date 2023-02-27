package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺颜乌拉Noin_ulaBarony struct {
	feud.BaseBarony
}

var BNoin_ula诺颜乌拉 feud.Barony = &诺颜乌拉Noin_ulaBarony{}

func init() {
    f := BNoin_ula诺颜乌拉.(*诺颜乌拉Noin_ulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noin_ula",
		TitleName: "诺颜乌拉",
		TitleCode: "b_noin_ula",
	}
}
