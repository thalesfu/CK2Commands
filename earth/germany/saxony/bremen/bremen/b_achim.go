package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿希姆AchimBarony struct {
	feud.BaseBarony
}

var BAchim阿希姆 feud.Barony = &阿希姆AchimBarony{}

func init() {
    f := BAchim阿希姆.(*阿希姆AchimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "achim",
		TitleName: "阿希姆",
		TitleCode: "b_achim",
	}
}
