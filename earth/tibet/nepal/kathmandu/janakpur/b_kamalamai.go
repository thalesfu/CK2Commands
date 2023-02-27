package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马拉迈KamalamaiBarony struct {
	feud.BaseBarony
}

var BKamalamai卡马拉迈 feud.Barony = &卡马拉迈KamalamaiBarony{}

func init() {
    f := BKamalamai卡马拉迈.(*卡马拉迈KamalamaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamalamai",
		TitleName: "卡马拉迈",
		TitleCode: "b_kamalamai",
	}
}
