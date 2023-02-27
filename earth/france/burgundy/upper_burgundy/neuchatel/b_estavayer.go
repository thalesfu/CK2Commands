package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯塔瓦耶EstavayerBarony struct {
	feud.BaseBarony
}

var BEstavayer埃斯塔瓦耶 feud.Barony = &埃斯塔瓦耶EstavayerBarony{}

func init() {
    f := BEstavayer埃斯塔瓦耶.(*埃斯塔瓦耶EstavayerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "estavayer",
		TitleName: "埃斯塔瓦耶",
		TitleCode: "b_estavayer",
	}
}
