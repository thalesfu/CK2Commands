package aj_bogd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 南里庄NanlizhuangBarony struct {
	feud.BaseBarony
}

var BNanlizhuang南里庄 feud.Barony = &南里庄NanlizhuangBarony{}

func init() {
    f := BNanlizhuang南里庄.(*南里庄NanlizhuangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nanlizhuang",
		TitleName: "南里庄",
		TitleCode: "b_nanlizhuang",
	}
}
