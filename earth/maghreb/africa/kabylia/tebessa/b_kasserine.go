package tebessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡塞林KasserineBarony struct {
	feud.BaseBarony
}

var BKasserine卡塞林 feud.Barony = &卡塞林KasserineBarony{}

func init() {
    f := BKasserine卡塞林.(*卡塞林KasserineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kasserine",
		TitleName: "卡塞林",
		TitleCode: "b_kasserine",
	}
}
