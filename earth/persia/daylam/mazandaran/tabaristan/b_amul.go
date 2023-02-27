package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莫勒AmulBarony struct {
	feud.BaseBarony
}

var BAmul阿莫勒 feud.Barony = &阿莫勒AmulBarony{}

func init() {
    f := BAmul阿莫勒.(*阿莫勒AmulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amul",
		TitleName: "阿莫勒",
		TitleCode: "b_amul",
	}
}
