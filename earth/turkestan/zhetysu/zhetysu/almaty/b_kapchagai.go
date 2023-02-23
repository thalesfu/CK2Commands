package almaty

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡普恰盖KapchagaiBarony struct {
	feud.BaseBarony
}

var BKapchagai卡普恰盖 feud.Barony = &卡普恰盖KapchagaiBarony{}

func init() {
	f := BKapchagai卡普恰盖.(*卡普恰盖KapchagaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapchagai",
		TitleName: "卡普恰盖",
		TitleCode: "b_kapchagai",
	}
}
