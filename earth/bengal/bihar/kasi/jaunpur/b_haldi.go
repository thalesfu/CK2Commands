package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃罗提HaldiBarony struct {
	feud.BaseBarony
}

var BHaldi诃罗提 feud.Barony = &诃罗提HaldiBarony{}

func init() {
	f := BHaldi诃罗提.(*诃罗提HaldiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haldi",
		TitleName: "诃罗提",
		TitleCode: "b_haldi",
	}
}
