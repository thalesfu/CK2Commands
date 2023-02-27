package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿温达卡AondakaBarony struct {
	feud.BaseBarony
}

var BAondaka阿温达卡 feud.Barony = &阿温达卡AondakaBarony{}

func init() {
    f := BAondaka阿温达卡.(*阿温达卡AondakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aondaka",
		TitleName: "阿温达卡",
		TitleCode: "b_aondaka",
	}
}
