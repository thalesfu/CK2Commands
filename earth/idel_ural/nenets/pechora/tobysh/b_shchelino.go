package tobysh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢利诺ShchelinoBarony struct {
	feud.BaseBarony
}

var BShchelino谢利诺 feud.Barony = &谢利诺ShchelinoBarony{}

func init() {
    f := BShchelino谢利诺.(*谢利诺ShchelinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shchelino",
		TitleName: "谢利诺",
		TitleCode: "b_shchelino",
	}
}
