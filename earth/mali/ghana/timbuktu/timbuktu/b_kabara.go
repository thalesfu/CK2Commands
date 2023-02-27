package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡巴拉KabaraBarony struct {
	feud.BaseBarony
}

var BKabara卡巴拉 feud.Barony = &卡巴拉KabaraBarony{}

func init() {
    f := BKabara卡巴拉.(*卡巴拉KabaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kabara",
		TitleName: "卡巴拉",
		TitleCode: "b_kabara",
	}
}
