package fyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿森斯AssensBarony struct {
	feud.BaseBarony
}

var BAssens阿森斯 feud.Barony = &阿森斯AssensBarony{}

func init() {
	f := BAssens阿森斯.(*阿森斯AssensBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assens",
		TitleName: "阿森斯",
		TitleCode: "b_assens",
	}
}
