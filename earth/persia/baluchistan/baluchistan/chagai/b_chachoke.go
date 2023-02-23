package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 遮朱计ChachokeBarony struct {
	feud.BaseBarony
}

var BChachoke遮朱计 feud.Barony = &遮朱计ChachokeBarony{}

func init() {
	f := BChachoke遮朱计.(*遮朱计ChachokeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chachoke",
		TitleName: "遮朱计",
		TitleCode: "b_chachoke",
	}
}
