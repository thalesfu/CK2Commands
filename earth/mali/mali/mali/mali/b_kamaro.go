package mali

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马罗KamaroBarony struct {
	feud.BaseBarony
}

var BKamaro卡马罗 feud.Barony = &卡马罗KamaroBarony{}

func init() {
    f := BKamaro卡马罗.(*卡马罗KamaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamaro",
		TitleName: "卡马罗",
		TitleCode: "b_kamaro",
	}
}
