package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库斯塔奈KostanayBarony struct {
	feud.BaseBarony
}

var BKostanay库斯塔奈 feud.Barony = &库斯塔奈KostanayBarony{}

func init() {
	f := BKostanay库斯塔奈.(*库斯塔奈KostanayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kostanay",
		TitleName: "库斯塔奈",
		TitleCode: "b_kostanay",
	}
}
