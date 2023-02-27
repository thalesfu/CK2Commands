package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥莱德阿米拉Oled_amiraBarony struct {
	feud.BaseBarony
}

var BOled_amira奥莱德阿米拉 feud.Barony = &奥莱德阿米拉Oled_amiraBarony{}

func init() {
    f := BOled_amira奥莱德阿米拉.(*奥莱德阿米拉Oled_amiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oled_amira",
		TitleName: "奥莱德阿米拉",
		TitleCode: "b_oled_amira",
	}
}
