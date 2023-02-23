package varadzin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢德布雷格LudbregBarony struct {
	feud.BaseBarony
}

var BLudbreg卢德布雷格 feud.Barony = &卢德布雷格LudbregBarony{}

func init() {
	f := BLudbreg卢德布雷格.(*卢德布雷格LudbregBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ludbreg",
		TitleName: "卢德布雷格",
		TitleCode: "b_ludbreg",
	}
}
