package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔瓦拉TaywaraBarony struct {
	feud.BaseBarony
}

var BTaywara塔瓦拉 feud.Barony = &塔瓦拉TaywaraBarony{}

func init() {
	f := BTaywara塔瓦拉.(*塔瓦拉TaywaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taywara",
		TitleName: "塔瓦拉",
		TitleCode: "b_taywara",
	}
}
