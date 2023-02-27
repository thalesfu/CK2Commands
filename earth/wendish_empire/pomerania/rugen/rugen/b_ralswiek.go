package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉尔斯维克RalswiekBarony struct {
	feud.BaseBarony
}

var BRalswiek拉尔斯维克 feud.Barony = &拉尔斯维克RalswiekBarony{}

func init() {
    f := BRalswiek拉尔斯维克.(*拉尔斯维克RalswiekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ralswiek",
		TitleName: "拉尔斯维克",
		TitleCode: "b_ralswiek",
	}
}
