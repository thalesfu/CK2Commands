package nellore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿登吉AddankiBarony struct {
	feud.BaseBarony
}

var BAddanki阿登吉 feud.Barony = &阿登吉AddankiBarony{}

func init() {
	f := BAddanki阿登吉.(*阿登吉AddankiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "addanki",
		TitleName: "阿登吉",
		TitleCode: "b_addanki",
	}
}
