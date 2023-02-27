package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗兰科潘FrankopanBarony struct {
	feud.BaseBarony
}

var BFrankopan弗兰科潘 feud.Barony = &弗兰科潘FrankopanBarony{}

func init() {
    f := BFrankopan弗兰科潘.(*弗兰科潘FrankopanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "frankopan",
		TitleName: "弗兰科潘",
		TitleCode: "b_frankopan",
	}
}
