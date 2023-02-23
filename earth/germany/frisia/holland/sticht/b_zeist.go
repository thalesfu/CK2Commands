package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽斯特ZeistBarony struct {
	feud.BaseBarony
}

var BZeist泽斯特 feud.Barony = &泽斯特ZeistBarony{}

func init() {
	f := BZeist泽斯特.(*泽斯特ZeistBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeist",
		TitleName: "泽斯特",
		TitleCode: "b_zeist",
	}
}
