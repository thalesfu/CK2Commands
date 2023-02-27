package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔嘎朗特ArgalantBarony struct {
	feud.BaseBarony
}

var BArgalant阿尔嘎朗特 feud.Barony = &阿尔嘎朗特ArgalantBarony{}

func init() {
    f := BArgalant阿尔嘎朗特.(*阿尔嘎朗特ArgalantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argalant",
		TitleName: "阿尔嘎朗特",
		TitleCode: "b_argalant",
	}
}
