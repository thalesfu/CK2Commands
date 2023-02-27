package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿普莱ApuleBarony struct {
	feud.BaseBarony
}

var BApule阿普莱 feud.Barony = &阿普莱ApuleBarony{}

func init() {
    f := BApule阿普莱.(*阿普莱ApuleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apule",
		TitleName: "阿普莱",
		TitleCode: "b_apule",
	}
}
