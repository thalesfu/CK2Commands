package vishera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克奇姆AkchimBarony struct {
	feud.BaseBarony
}

var BAkchim阿克奇姆 feud.Barony = &阿克奇姆AkchimBarony{}

func init() {
    f := BAkchim阿克奇姆.(*阿克奇姆AkchimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akchim",
		TitleName: "阿克奇姆",
		TitleCode: "b_akchim",
	}
}
