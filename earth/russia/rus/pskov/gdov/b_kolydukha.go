package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科雷杜哈KolydukhaBarony struct {
	feud.BaseBarony
}

var BKolydukha科雷杜哈 feud.Barony = &科雷杜哈KolydukhaBarony{}

func init() {
    f := BKolydukha科雷杜哈.(*科雷杜哈KolydukhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kolydukha",
		TitleName: "科雷杜哈",
		TitleCode: "b_kolydukha",
	}
}
