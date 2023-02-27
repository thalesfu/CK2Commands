package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾希施泰特EichstadtBarony struct {
	feud.BaseBarony
}

var BEichstadt艾希施泰特 feud.Barony = &艾希施泰特EichstadtBarony{}

func init() {
    f := BEichstadt艾希施泰特.(*艾希施泰特EichstadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eichstadt",
		TitleName: "艾希施泰特",
		TitleCode: "b_eichstadt",
	}
}
