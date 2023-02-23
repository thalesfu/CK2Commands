package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳夫普利翁NauplionBarony struct {
	feud.BaseBarony
}

var BNauplion纳夫普利翁 feud.Barony = &纳夫普利翁NauplionBarony{}

func init() {
	f := BNauplion纳夫普利翁.(*纳夫普利翁NauplionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nauplion",
		TitleName: "纳夫普利翁",
		TitleCode: "b_nauplion",
	}
}
