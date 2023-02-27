package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗饶海吉RozsahegyBarony struct {
	feud.BaseBarony
}

var BRozsahegy罗饶海吉 feud.Barony = &罗饶海吉RozsahegyBarony{}

func init() {
    f := BRozsahegy罗饶海吉.(*罗饶海吉RozsahegyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rozsahegy",
		TitleName: "罗饶海吉",
		TitleCode: "b_rozsahegy",
	}
}
