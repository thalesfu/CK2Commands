package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马蒙托沃MamontovoBarony struct {
	feud.BaseBarony
}

var BMamontovo马蒙托沃 feud.Barony = &马蒙托沃MamontovoBarony{}

func init() {
	f := BMamontovo马蒙托沃.(*马蒙托沃MamontovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamontovo",
		TitleName: "马蒙托沃",
		TitleCode: "b_mamontovo",
	}
}
