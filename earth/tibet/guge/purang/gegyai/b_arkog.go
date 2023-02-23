package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔过ArkogBarony struct {
	feud.BaseBarony
}

var BArkog阿尔过 feud.Barony = &阿尔过ArkogBarony{}

func init() {
	f := BArkog阿尔过.(*阿尔过ArkogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arkog",
		TitleName: "阿尔过",
		TitleCode: "b_arkog",
	}
}
