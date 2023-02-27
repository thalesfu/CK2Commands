package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约特瓦塔YotvataBarony struct {
	feud.BaseBarony
}

var BYotvata约特瓦塔 feud.Barony = &约特瓦塔YotvataBarony{}

func init() {
    f := BYotvata约特瓦塔.(*约特瓦塔YotvataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yotvata",
		TitleName: "约特瓦塔",
		TitleCode: "b_yotvata",
	}
}
