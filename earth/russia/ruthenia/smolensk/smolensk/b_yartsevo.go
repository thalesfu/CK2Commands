package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚尔采沃YartsevoBarony struct {
	feud.BaseBarony
}

var BYartsevo亚尔采沃 feud.Barony = &亚尔采沃YartsevoBarony{}

func init() {
	f := BYartsevo亚尔采沃.(*亚尔采沃YartsevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yartsevo",
		TitleName: "亚尔采沃",
		TitleCode: "b_yartsevo",
	}
}
