package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫尔马斯HarmashiBarony struct {
	feud.BaseBarony
}

var BHarmashi赫尔马斯 feud.Barony = &赫尔马斯HarmashiBarony{}

func init() {
	f := BHarmashi赫尔马斯.(*赫尔马斯HarmashiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harmashi",
		TitleName: "赫尔马斯",
		TitleCode: "b_harmashi",
	}
}
