package kartaly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基洛夫卡KirovkaBarony struct {
	feud.BaseBarony
}

var BKirovka基洛夫卡 feud.Barony = &基洛夫卡KirovkaBarony{}

func init() {
    f := BKirovka基洛夫卡.(*基洛夫卡KirovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirovka",
		TitleName: "基洛夫卡",
		TitleCode: "b_kirovka",
	}
}
