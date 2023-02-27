package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎卡兹尼克ZakaznikBarony struct {
	feud.BaseBarony
}

var BZakaznik扎卡兹尼克 feud.Barony = &扎卡兹尼克ZakaznikBarony{}

func init() {
    f := BZakaznik扎卡兹尼克.(*扎卡兹尼克ZakaznikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zakaznik",
		TitleName: "扎卡兹尼克",
		TitleCode: "b_zakaznik",
	}
}
