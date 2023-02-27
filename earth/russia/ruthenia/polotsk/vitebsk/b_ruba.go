package vitebsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁巴RubaBarony struct {
	feud.BaseBarony
}

var BRuba鲁巴 feud.Barony = &鲁巴RubaBarony{}

func init() {
    f := BRuba鲁巴.(*鲁巴RubaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruba",
		TitleName: "鲁巴",
		TitleCode: "b_ruba",
	}
}
