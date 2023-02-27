package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯梅纳AsmenaBarony struct {
	feud.BaseBarony
}

var BAsmena阿斯梅纳 feud.Barony = &阿斯梅纳AsmenaBarony{}

func init() {
    f := BAsmena阿斯梅纳.(*阿斯梅纳AsmenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asmena",
		TitleName: "阿斯梅纳",
		TitleCode: "b_asmena",
	}
}
