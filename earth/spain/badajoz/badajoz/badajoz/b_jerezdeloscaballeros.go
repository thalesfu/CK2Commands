package badajoz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫雷斯德洛斯卡瓦列罗斯JerezdeloscaballerosBarony struct {
	feud.BaseBarony
}

var BJerezdeloscaballeros赫雷斯德洛斯卡瓦列罗斯 feud.Barony = &赫雷斯德洛斯卡瓦列罗斯JerezdeloscaballerosBarony{}

func init() {
	f := BJerezdeloscaballeros赫雷斯德洛斯卡瓦列罗斯.(*赫雷斯德洛斯卡瓦列罗斯JerezdeloscaballerosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerezdeloscaballeros",
		TitleName: "赫雷斯德洛斯卡瓦列罗斯",
		TitleCode: "b_jerezdeloscaballeros",
	}
}
