package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿本贝格AbenbergBarony struct {
	feud.BaseBarony
}

var BAbenberg阿本贝格 feud.Barony = &阿本贝格AbenbergBarony{}

func init() {
    f := BAbenberg阿本贝格.(*阿本贝格AbenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abenberg",
		TitleName: "阿本贝格",
		TitleCode: "b_abenberg",
	}
}
