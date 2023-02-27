package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 康斯坦察ConstantiaBarony struct {
	feud.BaseBarony
}

var BConstantia康斯坦察 feud.Barony = &康斯坦察ConstantiaBarony{}

func init() {
    f := BConstantia康斯坦察.(*康斯坦察ConstantiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "constantia",
		TitleName: "康斯坦察",
		TitleCode: "b_constantia",
	}
}
