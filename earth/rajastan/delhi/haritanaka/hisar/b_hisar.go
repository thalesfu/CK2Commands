package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 醯娑罗HisarBarony struct {
	feud.BaseBarony
}

var BHisar醯娑罗 feud.Barony = &醯娑罗HisarBarony{}

func init() {
	f := BHisar醯娑罗.(*醯娑罗HisarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hisar",
		TitleName: "醯娑罗",
		TitleCode: "b_hisar",
	}
}
