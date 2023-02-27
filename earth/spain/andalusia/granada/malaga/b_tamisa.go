package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔米萨TamisaBarony struct {
	feud.BaseBarony
}

var BTamisa塔米萨 feud.Barony = &塔米萨TamisaBarony{}

func init() {
    f := BTamisa塔米萨.(*塔米萨TamisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamisa",
		TitleName: "塔米萨",
		TitleCode: "b_tamisa",
	}
}
