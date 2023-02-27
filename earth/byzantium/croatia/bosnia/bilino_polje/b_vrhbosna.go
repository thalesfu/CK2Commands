package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗尔赫波斯纳VrhbosnaBarony struct {
	feud.BaseBarony
}

var BVrhbosna弗尔赫波斯纳 feud.Barony = &弗尔赫波斯纳VrhbosnaBarony{}

func init() {
    f := BVrhbosna弗尔赫波斯纳.(*弗尔赫波斯纳VrhbosnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vrhbosna",
		TitleName: "弗尔赫波斯纳",
		TitleCode: "b_vrhbosna",
	}
}
