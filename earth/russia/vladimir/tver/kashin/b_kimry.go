package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基姆雷KimryBarony struct {
	feud.BaseBarony
}

var BKimry基姆雷 feud.Barony = &基姆雷KimryBarony{}

func init() {
	f := BKimry基姆雷.(*基姆雷KimryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kimry",
		TitleName: "基姆雷",
		TitleCode: "b_kimry",
	}
}
