package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提云陀罗城DevendranagarBarony struct {
	feud.BaseBarony
}

var BDevendranagar提云陀罗城 feud.Barony = &提云陀罗城DevendranagarBarony{}

func init() {
    f := BDevendranagar提云陀罗城.(*提云陀罗城DevendranagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "devendranagar",
		TitleName: "提云陀罗城",
		TitleCode: "b_devendranagar",
	}
}
