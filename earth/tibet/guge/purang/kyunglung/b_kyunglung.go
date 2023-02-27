package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琼隆KyunglungBarony struct {
	feud.BaseBarony
}

var BKyunglung琼隆 feud.Barony = &琼隆KyunglungBarony{}

func init() {
    f := BKyunglung琼隆.(*琼隆KyunglungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyunglung",
		TitleName: "琼隆",
		TitleCode: "b_kyunglung",
	}
}
