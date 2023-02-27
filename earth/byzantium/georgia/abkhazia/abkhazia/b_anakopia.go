package abkhazia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿纳科皮亚AnakopiaBarony struct {
	feud.BaseBarony
}

var BAnakopia阿纳科皮亚 feud.Barony = &阿纳科皮亚AnakopiaBarony{}

func init() {
    f := BAnakopia阿纳科皮亚.(*阿纳科皮亚AnakopiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anakopia",
		TitleName: "阿纳科皮亚",
		TitleCode: "b_anakopia",
	}
}
