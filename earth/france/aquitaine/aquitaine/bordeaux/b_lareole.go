package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉雷奥尔LareoleBarony struct {
	feud.BaseBarony
}

var BLareole拉雷奥尔 feud.Barony = &拉雷奥尔LareoleBarony{}

func init() {
	f := BLareole拉雷奥尔.(*拉雷奥尔LareoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lareole",
		TitleName: "拉雷奥尔",
		TitleCode: "b_lareole",
	}
}
