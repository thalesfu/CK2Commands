package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡特那QatnaBarony struct {
	feud.BaseBarony
}

var BQatna卡特那 feud.Barony = &卡特那QatnaBarony{}

func init() {
    f := BQatna卡特那.(*卡特那QatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qatna",
		TitleName: "卡特那",
		TitleCode: "b_qatna",
	}
}
