package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡菲利Caerphilly_gowerBarony struct {
	feud.BaseBarony
}

var BCaerphilly_gower卡菲利 feud.Barony = &卡菲利Caerphilly_gowerBarony{}

func init() {
    f := BCaerphilly_gower卡菲利.(*卡菲利Caerphilly_gowerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caerphilly_gower",
		TitleName: "卡菲利",
		TitleCode: "b_caerphilly_gower",
	}
}
