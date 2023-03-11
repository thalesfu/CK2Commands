package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔恰托夫KourchatovBarony struct {
	feud.BaseBarony
}

var BKourchatov库尔恰托夫 feud.Barony = &库尔恰托夫KourchatovBarony{}

func init() {
    f := BKourchatov库尔恰托夫.(*库尔恰托夫KourchatovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kourchatov",
		TitleName: "库尔恰托夫",
		TitleCode: "b_kourchatov",
	}
}
