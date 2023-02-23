package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔朱拉TadjouraBarony struct {
	feud.BaseBarony
}

var BTadjoura塔朱拉 feud.Barony = &塔朱拉TadjouraBarony{}

func init() {
	f := BTadjoura塔朱拉.(*塔朱拉TadjouraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tadjoura",
		TitleName: "塔朱拉",
		TitleCode: "b_tadjoura",
	}
}
