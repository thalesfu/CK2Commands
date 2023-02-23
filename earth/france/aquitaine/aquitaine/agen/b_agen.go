package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿让AgenBarony struct {
	feud.BaseBarony
}

var BAgen阿让 feud.Barony = &阿让AgenBarony{}

func init() {
	f := BAgen阿让.(*阿让AgenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agen",
		TitleName: "阿让",
		TitleCode: "b_agen",
	}
}
