package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利亚里CagliariBarony struct {
	feud.BaseBarony
}

var BCagliari卡利亚里 feud.Barony = &卡利亚里CagliariBarony{}

func init() {
    f := BCagliari卡利亚里.(*卡利亚里CagliariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cagliari",
		TitleName: "卡利亚里",
		TitleCode: "b_cagliari",
	}
}
