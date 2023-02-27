package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴甫洛BurahBarony struct {
	feud.BaseBarony
}

var BBurah巴甫洛 feud.Barony = &巴甫洛BurahBarony{}

func init() {
    f := BBurah巴甫洛.(*巴甫洛BurahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burah",
		TitleName: "巴甫洛",
		TitleCode: "b_burah",
	}
}
