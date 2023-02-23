package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内卢楼NelloreBarony struct {
	feud.BaseBarony
}

var BNellore内卢楼 feud.Barony = &内卢楼NelloreBarony{}

func init() {
	f := BNellore内卢楼.(*内卢楼NelloreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nellore",
		TitleName: "内卢楼",
		TitleCode: "b_nellore",
	}
}
