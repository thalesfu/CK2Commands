package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉瓦尔丹LavardinBarony struct {
	feud.BaseBarony
}

var BLavardin拉瓦尔丹 feud.Barony = &拉瓦尔丹LavardinBarony{}

func init() {
	f := BLavardin拉瓦尔丹.(*拉瓦尔丹LavardinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lavardin",
		TitleName: "拉瓦尔丹",
		TitleCode: "b_lavardin",
	}
}
