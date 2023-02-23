package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗杭爱ZuunkhangaiBarony struct {
	feud.BaseBarony
}

var BZuunkhangai宗杭爱 feud.Barony = &宗杭爱ZuunkhangaiBarony{}

func init() {
	f := BZuunkhangai宗杭爱.(*宗杭爱ZuunkhangaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zuunkhangai",
		TitleName: "宗杭爱",
		TitleCode: "b_zuunkhangai",
	}
}
