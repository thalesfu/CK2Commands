package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉基什卡RackiskaBarony struct {
	feud.BaseBarony
}

var BRackiska拉基什卡 feud.Barony = &拉基什卡RackiskaBarony{}

func init() {
    f := BRackiska拉基什卡.(*拉基什卡RackiskaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rackiska",
		TitleName: "拉基什卡",
		TitleCode: "b_rackiska",
	}
}
