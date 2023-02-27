package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济苏法勒DhisufalBarony struct {
	feud.BaseBarony
}

var BDhisufal济苏法勒 feud.Barony = &济苏法勒DhisufalBarony{}

func init() {
    f := BDhisufal济苏法勒.(*济苏法勒DhisufalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhisufal",
		TitleName: "济苏法勒",
		TitleCode: "b_dhisufal",
	}
}
