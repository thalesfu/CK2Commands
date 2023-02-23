package nasikya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 剌怛那揭陀RatangadBarony struct {
	feud.BaseBarony
}

var BRatangad剌怛那揭陀 feud.Barony = &剌怛那揭陀RatangadBarony{}

func init() {
	f := BRatangad剌怛那揭陀.(*剌怛那揭陀RatangadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ratangad",
		TitleName: "剌怛那揭陀",
		TitleCode: "b_ratangad",
	}
}
