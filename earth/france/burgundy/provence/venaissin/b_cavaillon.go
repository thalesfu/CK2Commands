package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡瓦永CavaillonBarony struct {
	feud.BaseBarony
}

var BCavaillon卡瓦永 feud.Barony = &卡瓦永CavaillonBarony{}

func init() {
    f := BCavaillon卡瓦永.(*卡瓦永CavaillonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cavaillon",
		TitleName: "卡瓦永",
		TitleCode: "b_cavaillon",
	}
}
