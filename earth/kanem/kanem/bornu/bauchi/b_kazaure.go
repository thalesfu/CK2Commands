package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡扎乌雷KazaureBarony struct {
	feud.BaseBarony
}

var BKazaure卡扎乌雷 feud.Barony = &卡扎乌雷KazaureBarony{}

func init() {
	f := BKazaure卡扎乌雷.(*卡扎乌雷KazaureBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazaure",
		TitleName: "卡扎乌雷",
		TitleCode: "b_kazaure",
	}
}
