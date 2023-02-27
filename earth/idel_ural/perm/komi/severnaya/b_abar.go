package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴尔AbarBarony struct {
	feud.BaseBarony
}

var BAbar阿巴尔 feud.Barony = &阿巴尔AbarBarony{}

func init() {
    f := BAbar阿巴尔.(*阿巴尔AbarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abar",
		TitleName: "阿巴尔",
		TitleCode: "b_abar",
	}
}
