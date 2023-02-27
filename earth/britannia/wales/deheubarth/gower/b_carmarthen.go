package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马森CarmarthenBarony struct {
	feud.BaseBarony
}

var BCarmarthen卡马森 feud.Barony = &卡马森CarmarthenBarony{}

func init() {
    f := BCarmarthen卡马森.(*卡马森CarmarthenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carmarthen",
		TitleName: "卡马森",
		TitleCode: "b_carmarthen",
	}
}
