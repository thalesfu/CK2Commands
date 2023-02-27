package galich_mersky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏萨尼诺SusaninoBarony struct {
	feud.BaseBarony
}

var BSusanino苏萨尼诺 feud.Barony = &苏萨尼诺SusaninoBarony{}

func init() {
    f := BSusanino苏萨尼诺.(*苏萨尼诺SusaninoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "susanino",
		TitleName: "苏萨尼诺",
		TitleCode: "b_susanino",
	}
}
