package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃悉帝军持HastikundiBarony struct {
	feud.BaseBarony
}

var BHastikundi诃悉帝军持 feud.Barony = &诃悉帝军持HastikundiBarony{}

func init() {
	f := BHastikundi诃悉帝军持.(*诃悉帝军持HastikundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hastikundi",
		TitleName: "诃悉帝军持",
		TitleCode: "b_hastikundi",
	}
}
