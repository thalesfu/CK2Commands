package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基列马雷KilemaryBarony struct {
	feud.BaseBarony
}

var BKilemary基列马雷 feud.Barony = &基列马雷KilemaryBarony{}

func init() {
    f := BKilemary基列马雷.(*基列马雷KilemaryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilemary",
		TitleName: "基列马雷",
		TitleCode: "b_kilemary",
	}
}
