package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔麦克雷南KilmacrenanBarony struct {
	feud.BaseBarony
}

var BKilmacrenan基尔麦克雷南 feud.Barony = &基尔麦克雷南KilmacrenanBarony{}

func init() {
	f := BKilmacrenan基尔麦克雷南.(*基尔麦克雷南KilmacrenanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilmacrenan",
		TitleName: "基尔麦克雷南",
		TitleCode: "b_kilmacrenan",
	}
}
