package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒拜卡AshshubaykahBarony struct {
	feud.BaseBarony
}

var BAshshubaykah舒拜卡 feud.Barony = &舒拜卡AshshubaykahBarony{}

func init() {
    f := BAshshubaykah舒拜卡.(*舒拜卡AshshubaykahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashshubaykah",
		TitleName: "舒拜卡",
		TitleCode: "b_ashshubaykah",
	}
}
