package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀什真PostyenBarony struct {
	feud.BaseBarony
}

var BPostyen珀什真 feud.Barony = &珀什真PostyenBarony{}

func init() {
	f := BPostyen珀什真.(*珀什真PostyenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "postyen",
		TitleName: "珀什真",
		TitleCode: "b_postyen",
	}
}
