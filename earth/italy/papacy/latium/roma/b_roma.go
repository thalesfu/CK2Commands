package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗马RomaBarony struct {
	feud.BaseBarony
}

var BRoma罗马 feud.Barony = &罗马RomaBarony{}

func init() {
	f := BRoma罗马.(*罗马RomaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roma",
		TitleName: "罗马",
		TitleCode: "b_roma",
	}
}
