package sarasvata_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃罗卢KheraluBarony struct {
	feud.BaseBarony
}

var BKheralu诃罗卢 feud.Barony = &诃罗卢KheraluBarony{}

func init() {
    f := BKheralu诃罗卢.(*诃罗卢KheraluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kheralu",
		TitleName: "诃罗卢",
		TitleCode: "b_kheralu",
	}
}
