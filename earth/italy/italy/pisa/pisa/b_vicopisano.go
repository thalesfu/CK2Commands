package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维科皮萨诺VicopisanoBarony struct {
	feud.BaseBarony
}

var BVicopisano维科皮萨诺 feud.Barony = &维科皮萨诺VicopisanoBarony{}

func init() {
    f := BVicopisano维科皮萨诺.(*维科皮萨诺VicopisanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vicopisano",
		TitleName: "维科皮萨诺",
		TitleCode: "b_vicopisano",
	}
}
