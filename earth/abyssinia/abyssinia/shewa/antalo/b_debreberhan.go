package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德卜勒伯尔汉DebreberhanBarony struct {
	feud.BaseBarony
}

var BDebreberhan德卜勒伯尔汉 feud.Barony = &德卜勒伯尔汉DebreberhanBarony{}

func init() {
    f := BDebreberhan德卜勒伯尔汉.(*德卜勒伯尔汉DebreberhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debreberhan",
		TitleName: "德卜勒伯尔汉",
		TitleCode: "b_debreberhan",
	}
}
