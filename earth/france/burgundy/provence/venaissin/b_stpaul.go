package venaissin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣保罗StpaulBarony struct {
	feud.BaseBarony
}

var BStpaul圣保罗 feud.Barony = &圣保罗StpaulBarony{}

func init() {
	f := BStpaul圣保罗.(*圣保罗StpaulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stpaul",
		TitleName: "圣保罗",
		TitleCode: "b_stpaul",
	}
}
