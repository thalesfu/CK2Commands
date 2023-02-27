package kharga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴里斯BarisBarony struct {
	feud.BaseBarony
}

var BBaris巴里斯 feud.Barony = &巴里斯BarisBarony{}

func init() {
    f := BBaris巴里斯.(*巴里斯BarisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baris",
		TitleName: "巴里斯",
		TitleCode: "b_baris",
	}
}
