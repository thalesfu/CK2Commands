package alcantara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞克拉温CeclavinBarony struct {
	feud.BaseBarony
}

var BCeclavin塞克拉温 feud.Barony = &塞克拉温CeclavinBarony{}

func init() {
    f := BCeclavin塞克拉温.(*塞克拉温CeclavinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ceclavin",
		TitleName: "塞克拉温",
		TitleCode: "b_ceclavin",
	}
}
