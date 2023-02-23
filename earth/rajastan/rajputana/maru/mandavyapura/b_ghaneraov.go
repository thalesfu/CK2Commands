package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 嘎内拉罗GhaneraovBarony struct {
	feud.BaseBarony
}

var BGhaneraov嘎内拉罗 feud.Barony = &嘎内拉罗GhaneraovBarony{}

func init() {
	f := BGhaneraov嘎内拉罗.(*嘎内拉罗GhaneraovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghaneraov",
		TitleName: "嘎内拉罗",
		TitleCode: "b_ghaneraov",
	}
}
