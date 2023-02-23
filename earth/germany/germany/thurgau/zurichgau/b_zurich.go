package zurichgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏黎世ZurichBarony struct {
	feud.BaseBarony
}

var BZurich苏黎世 feud.Barony = &苏黎世ZurichBarony{}

func init() {
	f := BZurich苏黎世.(*苏黎世ZurichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zurich",
		TitleName: "苏黎世",
		TitleCode: "b_zurich",
	}
}
