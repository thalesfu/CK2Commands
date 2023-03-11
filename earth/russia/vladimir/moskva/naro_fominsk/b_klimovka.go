package naro_fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克利莫夫卡KlimovkaBarony struct {
	feud.BaseBarony
}

var BKlimovka克利莫夫卡 feud.Barony = &克利莫夫卡KlimovkaBarony{}

func init() {
    f := BKlimovka克利莫夫卡.(*克利莫夫卡KlimovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klimovka",
		TitleName: "克利莫夫卡",
		TitleCode: "b_klimovka",
	}
}
