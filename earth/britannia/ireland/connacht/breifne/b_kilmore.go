package breifne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔莫雷KilmoreBarony struct {
	feud.BaseBarony
}

var BKilmore基尔莫雷 feud.Barony = &基尔莫雷KilmoreBarony{}

func init() {
    f := BKilmore基尔莫雷.(*基尔莫雷KilmoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilmore",
		TitleName: "基尔莫雷",
		TitleCode: "b_kilmore",
	}
}
