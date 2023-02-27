package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维捷格拉VytegraBarony struct {
	feud.BaseBarony
}

var BVytegra维捷格拉 feud.Barony = &维捷格拉VytegraBarony{}

func init() {
    f := BVytegra维捷格拉.(*维捷格拉VytegraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vytegra",
		TitleName: "维捷格拉",
		TitleCode: "b_vytegra",
	}
}
