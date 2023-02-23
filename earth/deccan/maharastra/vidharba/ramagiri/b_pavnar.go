package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普拉纳PavnarBarony struct {
	feud.BaseBarony
}

var BPavnar普拉纳 feud.Barony = &普拉纳PavnarBarony{}

func init() {
	f := BPavnar普拉纳.(*普拉纳PavnarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pavnar",
		TitleName: "普拉纳",
		TitleCode: "b_pavnar",
	}
}
