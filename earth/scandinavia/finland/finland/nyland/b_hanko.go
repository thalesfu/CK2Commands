package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉科HankoBarony struct {
	feud.BaseBarony
}

var BHanko汉科 feud.Barony = &汉科HankoBarony{}

func init() {
    f := BHanko汉科.(*汉科HankoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hanko",
		TitleName: "汉科",
		TitleCode: "b_hanko",
	}
}
